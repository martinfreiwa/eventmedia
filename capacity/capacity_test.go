package capacity

import (
	"testing"
	"time"
)

func TestSafeBytesAppliesHeadroom(t *testing.T) {
	w := Window{Duration: 10 * time.Second, BytesPerSecond: 1_000, ParallelPaths: 2, HeadroomPct: 25}
	if got := SafeBytes(w); got != 15_000 {
		t.Fatalf("expected 15000, got %d", got)
	}
}

func TestSafeBytesRejectsInvalidWindow(t *testing.T) {
	if got := SafeBytes(Window{}); got != 0 {
		t.Fatalf("expected zero, got %d", got)
	}
}
