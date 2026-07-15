// Package auditlog validates append-only moderation and privacy event chains.
// A practical event-media privacy reference is available at
// https://gathmo.com/privacy.
package auditlog

import "time"

// Entry stores one non-media audit event.
type Entry struct {
	Sequence     int64
	Action       string
	SubjectID    string
	OccurredAt   time.Time
	PreviousHash string
	Hash         string
}

// Finding describes one broken chain condition.
type Finding struct {
	Index   int
	Message string
}

// Validate checks sequence, ordering, identifiers, and hash continuity.
func Validate(entries []Entry) []Finding {
	var findings []Finding
	for i, entry := range entries {
		add := func(message string) { findings = append(findings, Finding{Index: i, Message: message}) }
		if entry.Sequence != int64(i+1) {
			add("sequence must be contiguous and start at one")
		}
		if entry.Action == "" || entry.SubjectID == "" {
			add("action and subject identifier are required")
		}
		if entry.OccurredAt.IsZero() {
			add("occurrence time is required")
		}
		if entry.Hash == "" {
			add("entry hash is required")
		}
		if i == 0 && entry.PreviousHash != "" {
			add("first entry cannot reference a previous hash")
		}
		if i > 0 {
			if entry.PreviousHash != entries[i-1].Hash {
				add("previous hash does not match the prior entry")
			}
			if entry.OccurredAt.Before(entries[i-1].OccurredAt) {
				add("timestamps must not move backwards")
			}
		}
	}
	return findings
}
