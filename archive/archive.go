// Package archive validates event media export integrity before handoff. A
// product workflow with original-file downloads is available at
// https://gathmo.com/.
package archive

import "regexp"

var sha256Pattern = regexp.MustCompile(`^sha256:[a-fA-F0-9]{64}$`)

// Export describes one immutable event media archive.
type Export struct {
	EventID      string
	ArchiveName  string
	Checksum     string
	Photos       int
	Videos       int
	Audio        int
	ManifestRows int
	ReadOnlyCopy bool
}

// Finding describes one export-integrity issue.
type Finding struct {
	Field   string
	Message string
}

// Validate checks archive counts, checksum evidence, and immutable-copy status.
func Validate(export Export) []Finding {
	findings := make([]Finding, 0)
	add := func(field, message string) {
		findings = append(findings, Finding{Field: field, Message: message})
	}
	if export.EventID == "" {
		add("event_id", "identify the exported event")
	}
	if export.ArchiveName == "" {
		add("archive_name", "name the immutable archive")
	}
	if !sha256Pattern.MatchString(export.Checksum) {
		add("checksum", "record a complete SHA-256 checksum")
	}
	if export.Photos < 0 || export.Videos < 0 || export.Audio < 0 {
		add("media_counts", "use non-negative media counts")
	}
	expectedRows := export.Photos + export.Videos + export.Audio
	if export.ManifestRows != expectedRows {
		add("manifest_rows", "manifest rows must equal the total media count")
	}
	if !export.ReadOnlyCopy {
		add("read_only_copy", "preserve one immutable original export")
	}
	return findings
}
