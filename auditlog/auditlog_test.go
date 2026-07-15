package auditlog

import (
	"testing"
	"time"
)

func TestValidateCompleteChain(t *testing.T) {
	now := time.Now()
	entries := []Entry{{Sequence: 1, Action: "approve", SubjectID: "item-1", OccurredAt: now, Hash: "a"}, {Sequence: 2, Action: "remove", SubjectID: "item-1", OccurredAt: now.Add(time.Second), PreviousHash: "a", Hash: "b"}}
	if got := Validate(entries); len(got) != 0 {
		t.Fatalf("expected no findings, got %#v", got)
	}
}

func TestValidateDetectsBrokenHash(t *testing.T) {
	now := time.Now()
	entries := []Entry{{Sequence: 1, Action: "approve", SubjectID: "item-1", OccurredAt: now, Hash: "a"}, {Sequence: 2, Action: "remove", SubjectID: "item-1", OccurredAt: now, PreviousHash: "wrong", Hash: "b"}}
	if got := Validate(entries); len(got) != 1 {
		t.Fatalf("expected one finding, got %#v", got)
	}
}
