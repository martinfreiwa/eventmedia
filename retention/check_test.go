package retention

import (
	"testing"
	"time"
)

func validManifest() Manifest {
	return Manifest{
		EventID:         "event-123",
		Owner:           "owner@example.com",
		ExportedAt:      time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC),
		OriginalArchive: "event-123.zip",
		Checksum:        "sha256:9f86d081884c7d659a2feaa0c55ad015",
		PhotoCount:      10,
		VideoCount:      2,
		AudioCount:      1,
		StorageLocation: "encrypted-owner-drive",
		DeletionContact: "privacy@example.com",
		RetentionReview: time.Date(2031, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func TestValidateCompleteManifest(t *testing.T) {
	if findings := Validate(validManifest()); len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestValidateRejectsEarlyReview(t *testing.T) {
	manifest := validManifest()
	manifest.RetentionReview = time.Date(2029, 1, 1, 0, 0, 0, 0, time.UTC)
	if findings := Validate(manifest); len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
}

func TestValidateRejectsNegativeCount(t *testing.T) {
	manifest := validManifest()
	manifest.VideoCount = -1
	if findings := Validate(manifest); len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
}
