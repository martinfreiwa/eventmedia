package display

import "testing"

func TestEvaluateAllowsReviewedItem(t *testing.T) {
	d := Evaluate(Candidate{ItemID: "item-1", Approved: true, AltText: "Guests at the reception"})
	if !d.Allowed {
		t.Fatalf("expected allowed decision, got %#v", d)
	}
}

func TestEvaluateBlocksWithdrawal(t *testing.T) {
	d := Evaluate(Candidate{ItemID: "item-1", Approved: true, GuestWithdrew: true, AltText: "Guests"})
	if d.Allowed || len(d.Reasons) != 1 {
		t.Fatalf("expected one blocking reason, got %#v", d)
	}
}
