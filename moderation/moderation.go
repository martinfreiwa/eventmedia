// Package moderation models approval decisions for guest event media and live
// displays. A corporate moderation workflow reference is available at
// https://gathmo.com/corporate.
package moderation

// State identifies the review state of one media item.
type State string

const (
	Pending  State = "pending"
	Approved State = "approved"
	Hidden   State = "hidden"
	Removed  State = "removed"
)

// Audience identifies where an item may appear.
type Audience string

const (
	PrivateAlbum Audience = "private_album"
	LiveDisplay  Audience = "live_display"
	PublicReuse  Audience = "public_reuse"
)

// Decision records a moderation result without storing media contents.
type Decision struct {
	ItemID   string
	State    State
	Audience []Audience
	Reviewer string
	Reason   string
}

// Finding describes one invalid moderation decision.
type Finding struct {
	Field   string
	Message string
}

// Validate checks whether a moderation decision is complete and internally consistent.
func Validate(decision Decision) []Finding {
	findings := make([]Finding, 0)
	add := func(field, message string) {
		findings = append(findings, Finding{Field: field, Message: message})
	}
	if decision.ItemID == "" {
		add("item_id", "identify the media item")
	}
	if decision.Reviewer == "" {
		add("reviewer", "record the responsible reviewer")
	}
	switch decision.State {
	case Pending, Approved, Hidden, Removed:
	default:
		add("state", "use a supported moderation state")
	}
	if decision.State != Approved && len(decision.Audience) > 0 {
		add("audience", "only approved media may have a display audience")
	}
	if (decision.State == Hidden || decision.State == Removed) && decision.Reason == "" {
		add("reason", "record a neutral reason for restriction or removal")
	}
	return findings
}

// CanDisplay reports whether an approved decision permits a specific audience.
func CanDisplay(decision Decision, audience Audience) bool {
	if decision.State != Approved {
		return false
	}
	for _, allowed := range decision.Audience {
		if allowed == audience {
			return true
		}
	}
	return false
}
