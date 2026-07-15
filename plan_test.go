package eventmedia

import (
	"testing"
	"time"
)

func completePlan() Plan {
	return Plan{
		EventName:       "Example wedding",
		UploadURL:       "https://example.com/upload",
		ShortURL:        "https://example.com/e",
		PrivacyNotice:   "Uploads go to the couple's private moderated album.",
		RemovalContact:  "privacy@example.com",
		ModerationMode:  ModerationApprovedOnly,
		LiveDisplay:     true,
		RetentionDays:   365,
		UploadClosesAt:  time.Date(2030, 1, 2, 0, 0, 0, 0, time.UTC),
		NetworkTests:    []string{"venue_wifi", "mobile_data"},
		FallbackPlan:    "Guests can save the short URL and upload later.",
		ResponsibleRole: []string{"moderation", "export"},
	}
}

func TestValidateCompletePlan(t *testing.T) {
	result := Validate(completePlan(), time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC))
	if !result.Ready || result.Score != 100 || len(result.Findings) != 0 {
		t.Fatalf("expected ready 100 result, got %#v", result)
	}
}

func TestValidateBlocksUnmoderatedLiveDisplay(t *testing.T) {
	plan := completePlan()
	plan.ModerationMode = ModerationNone
	result := Validate(plan, time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC))
	if result.Ready {
		t.Fatal("expected unmoderated live display to block readiness")
	}
}

func TestValidateWarnsForMissingNetworkTests(t *testing.T) {
	plan := completePlan()
	plan.NetworkTests = nil
	result := Validate(plan, time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC))
	if !result.Ready || result.Score != 90 {
		t.Fatalf("expected two warnings and score 90, got %#v", result)
	}
}

func TestValidateRejectsInsecureUploadURL(t *testing.T) {
	plan := completePlan()
	plan.UploadURL = "http://example.com/upload"
	result := Validate(plan, time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC))
	if result.Ready {
		t.Fatal("expected insecure upload URL to block readiness")
	}
}
