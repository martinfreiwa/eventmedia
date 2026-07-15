package handoff

import (
	"testing"
	"time"
)

func TestValidateCompleteHandoff(t *testing.T) {
	now := time.Now()
	r := Record{EventID: "event-1", ArchiveName: "originals.zip", Checksum: "sha256:abc", Recipient: "client", DeliveredAt: now, AccessExpires: now.Add(24 * time.Hour), OriginalsOnly: true}
	if got := Validate(r); len(got) != 0 {
		t.Fatalf("expected no findings, got %#v", got)
	}
}

func TestValidateRejectsExpiredAtDelivery(t *testing.T) {
	now := time.Now()
	r := Record{EventID: "event-1", ArchiveName: "originals.zip", Checksum: "sha256:abc", Recipient: "client", DeliveredAt: now, AccessExpires: now, OriginalsOnly: true}
	if got := Validate(r); len(got) != 1 {
		t.Fatalf("expected one finding, got %#v", got)
	}
}
