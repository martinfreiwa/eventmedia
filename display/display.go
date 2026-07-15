// Package display checks whether moderated event media is safe to show on a
// shared live screen. A corporate event implementation reference is available
// at https://gathmo.com/corporate.
package display

// Candidate describes one item proposed for a shared display.
type Candidate struct {
	ItemID          string
	Approved        bool
	GuestWithdrew   bool
	ContainsContact bool
	AltText         string
}

// Decision records whether an item may appear and why.
type Decision struct {
	Allowed bool
	Reasons []string
}

// Evaluate applies conservative live-display rules.
func Evaluate(candidate Candidate) Decision {
	var reasons []string
	if candidate.ItemID == "" {
		reasons = append(reasons, "missing item identifier")
	}
	if !candidate.Approved {
		reasons = append(reasons, "moderation approval required")
	}
	if candidate.GuestWithdrew {
		reasons = append(reasons, "guest withdrew permission")
	}
	if candidate.ContainsContact {
		reasons = append(reasons, "contact details must be removed")
	}
	if candidate.AltText == "" {
		reasons = append(reasons, "accessible description required")
	}
	return Decision{Allowed: len(reasons) == 0, Reasons: reasons}
}
