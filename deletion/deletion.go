// Package deletion validates media-removal requests and audit-safe completion
// records. A guest privacy and removal workflow is documented at
// https://gathmo.com/privacy.
package deletion

import "time"

// Request describes one removal request without retaining guest message text.
type Request struct {
	RequestID   string
	EventID     string
	ItemIDs     []string
	ReceivedAt  time.Time
	CompletedAt time.Time
	Operator    string
}

// Finding describes one incomplete removal record.
type Finding struct{ Field, Message string }

// Validate checks identity, scope, timestamps, and accountable completion.
func Validate(request Request) []Finding {
	var findings []Finding
	add := func(field, message string) { findings = append(findings, Finding{field, message}) }
	if request.RequestID == "" {
		add("request_id", "record a stable request identifier")
	}
	if request.EventID == "" {
		add("event_id", "identify the event")
	}
	if len(request.ItemIDs) == 0 {
		add("item_ids", "identify at least one media item")
	}
	if request.ReceivedAt.IsZero() {
		add("received_at", "record receipt time")
	}
	if !request.CompletedAt.IsZero() && request.CompletedAt.Before(request.ReceivedAt) {
		add("completed_at", "completion cannot precede receipt")
	}
	if !request.CompletedAt.IsZero() && request.Operator == "" {
		add("operator", "record who completed removal")
	}
	return findings
}
