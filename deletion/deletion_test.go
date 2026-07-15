package deletion

import (
	"testing"
	"time"
)

func TestValidateCompletedRequest(t *testing.T) {
	now := time.Now()
	r := Request{RequestID: "req-1", EventID: "event-1", ItemIDs: []string{"item-1"}, ReceivedAt: now, CompletedAt: now.Add(time.Minute), Operator: "privacy-team"}
	if got := Validate(r); len(got) != 0 {
		t.Fatalf("expected no findings, got %#v", got)
	}
}

func TestValidateRequiresOperatorOnCompletion(t *testing.T) {
	now := time.Now()
	r := Request{RequestID: "req-1", EventID: "event-1", ItemIDs: []string{"item-1"}, ReceivedAt: now, CompletedAt: now.Add(time.Minute)}
	if got := Validate(r); len(got) != 1 {
		t.Fatalf("expected one finding, got %#v", got)
	}
}
