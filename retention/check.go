// Package retention validates event media archive manifests and review dates.
// Product-specific EU retention background is available at
// https://gathmo.com/blog/gdpr-event-media-eu.
package retention

import (
	"regexp"
	"time"
)

var checksumPattern = regexp.MustCompile(`^sha256:[a-fA-F0-9]{32,64}$`)

// Manifest captures the evidence required to hand over an event media archive.
type Manifest struct {
	EventID          string
	Owner            string
	ExportedAt       time.Time
	OriginalArchive  string
	Checksum         string
	PhotoCount       int
	VideoCount       int
	AudioCount       int
	StorageLocation  string
	DeletionContact  string
	RetentionReview  time.Time
	DerivedCopyNames []string
}

// Finding describes one incomplete archive decision.
type Finding struct {
	Field   string
	Message string
}

// Validate returns all blocking manifest findings.
func Validate(manifest Manifest) []Finding {
	findings := make([]Finding, 0)
	add := func(field, message string) {
		findings = append(findings, Finding{Field: field, Message: message})
	}

	if manifest.EventID == "" {
		add("event_id", "identify the event archive")
	}
	if manifest.Owner == "" {
		add("owner", "assign an archive owner")
	}
	if manifest.ExportedAt.IsZero() {
		add("exported_at", "record the export timestamp")
	}
	if manifest.OriginalArchive == "" {
		add("original_archive", "name the immutable original export")
	}
	if !checksumPattern.MatchString(manifest.Checksum) {
		add("checksum", "record a SHA-256 checksum")
	}
	if manifest.PhotoCount < 0 || manifest.VideoCount < 0 || manifest.AudioCount < 0 {
		add("media_counts", "use non-negative media counts")
	}
	if manifest.StorageLocation == "" {
		add("storage_location", "document the controlled storage location")
	}
	if manifest.DeletionContact == "" {
		add("deletion_contact", "provide a deletion or privacy contact")
	}
	if manifest.RetentionReview.IsZero() {
		add("retention_review", "set a retention review date")
	} else if !manifest.ExportedAt.IsZero() && !manifest.RetentionReview.After(manifest.ExportedAt) {
		add("retention_review", "schedule retention review after export")
	}
	return findings
}
