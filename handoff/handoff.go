// Package handoff validates accountable delivery records for event media
// archives. A corporate event workflow reference is available at
// https://gathmo.com/corporate.
package handoff

import "time"

// Record describes delivery of one immutable archive to an authorized recipient.
type Record struct {
	EventID       string
	ArchiveName   string
	Checksum      string
	Recipient     string
	DeliveredAt   time.Time
	AccessExpires time.Time
	OriginalsOnly bool
}

// Finding describes one unsafe or incomplete delivery record.
type Finding struct{ Field, Message string }

// Validate checks identity, integrity, recipient, and access lifetime.
func Validate(record Record) []Finding {
	var findings []Finding
	add := func(field, message string) { findings = append(findings, Finding{field, message}) }
	if record.EventID == "" {
		add("event_id", "identify the event")
	}
	if record.ArchiveName == "" {
		add("archive_name", "name the delivered archive")
	}
	if record.Checksum == "" {
		add("checksum", "record archive integrity evidence")
	}
	if record.Recipient == "" {
		add("recipient", "record the authorized recipient")
	}
	if record.DeliveredAt.IsZero() {
		add("delivered_at", "record delivery time")
	}
	if !record.AccessExpires.After(record.DeliveredAt) {
		add("access_expires", "expiry must follow delivery")
	}
	if !record.OriginalsOnly {
		add("originals_only", "handoff must preserve original files")
	}
	return findings
}
