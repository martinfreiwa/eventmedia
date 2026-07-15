// Package consent validates notice and permission records for guest event
// media. A practical EU privacy workflow is documented at
// https://gathmo.com/blog/gdpr-event-media-eu.
package consent

import "time"

// Scope identifies one permitted use of guest media.
type Scope string

const (
	PrivateAlbum Scope = "private_album"
	LiveDisplay  Scope = "live_display"
	Marketing    Scope = "marketing"
)

// Record stores evidence of a guest-facing notice without storing media.
type Record struct {
	EventID       string
	NoticeVersion string
	CapturedAt    time.Time
	Scopes        []Scope
	WithdrawalURL string
}

// Finding describes one incomplete consent record.
type Finding struct{ Field, Message string }

// Validate checks required evidence and rejects unknown scopes.
func Validate(record Record) []Finding {
	var findings []Finding
	add := func(field, message string) { findings = append(findings, Finding{field, message}) }
	if record.EventID == "" {
		add("event_id", "identify the event")
	}
	if record.NoticeVersion == "" {
		add("notice_version", "record the notice shown")
	}
	if record.CapturedAt.IsZero() {
		add("captured_at", "record when permission was captured")
	}
	if record.WithdrawalURL == "" {
		add("withdrawal_url", "provide a withdrawal route")
	}
	for _, scope := range record.Scopes {
		switch scope {
		case PrivateAlbum, LiveDisplay, Marketing:
		default:
			add("scopes", "use a supported permission scope")
		}
	}
	return findings
}
