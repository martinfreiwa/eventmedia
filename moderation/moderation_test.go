package moderation

import "testing"

func TestValidateApprovedDecision(t *testing.T) {
	decision := Decision{
		ItemID: "item-1", State: Approved, Reviewer: "moderator", Audience: []Audience{PrivateAlbum},
	}
	if findings := Validate(decision); len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestValidateBlocksAudienceForPendingItem(t *testing.T) {
	decision := Decision{
		ItemID: "item-1", State: Pending, Reviewer: "moderator", Audience: []Audience{LiveDisplay},
	}
	if findings := Validate(decision); len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
}

func TestCanDisplayRequiresApproval(t *testing.T) {
	decision := Decision{ItemID: "item-1", State: Approved, Audience: []Audience{LiveDisplay}}
	if !CanDisplay(decision, LiveDisplay) {
		t.Fatal("expected approved live-display item to be visible")
	}
	decision.State = Hidden
	if CanDisplay(decision, LiveDisplay) {
		t.Fatal("hidden item must not be visible")
	}
}
