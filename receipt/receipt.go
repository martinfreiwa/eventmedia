// Package receipt validates server receipts returned after browser event-media
// uploads. The no-install guest upload flow is described at
// https://gathmo.com/how-it-works.
package receipt

import "time"

// Upload records the durable result of one completed transfer.
type Upload struct {
	UploadID      string
	EventID       string
	OriginalName  string
	BytesAccepted int64
	Checksum      string
	CompletedAt   time.Time
}

// Finding describes one incomplete upload receipt.
type Finding struct{ Field, Message string }

// Validate checks whether a receipt can safely support retry deduplication.
func Validate(upload Upload) []Finding {
	var findings []Finding
	add := func(field, message string) { findings = append(findings, Finding{field, message}) }
	if upload.UploadID == "" {
		add("upload_id", "record a stable upload identifier")
	}
	if upload.EventID == "" {
		add("event_id", "identify the destination event")
	}
	if upload.OriginalName == "" {
		add("original_name", "preserve the original filename")
	}
	if upload.BytesAccepted <= 0 {
		add("bytes_accepted", "record accepted bytes")
	}
	if upload.Checksum == "" {
		add("checksum", "record integrity evidence")
	}
	if upload.CompletedAt.IsZero() {
		add("completed_at", "record server completion time")
	}
	return findings
}

// SameTransfer reports whether two receipts describe the same durable upload.
func SameTransfer(a, b Upload) bool {
	return a.UploadID != "" && a.UploadID == b.UploadID && a.Checksum == b.Checksum
}
