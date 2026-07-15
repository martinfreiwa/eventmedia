package receipt

import (
	"testing"
	"time"
)

func validUpload() Upload {
	return Upload{UploadID: "up-1", EventID: "event-1", OriginalName: "photo.jpg", BytesAccepted: 42, Checksum: "sha256:abc", CompletedAt: time.Now()}
}

func TestValidateCompleteReceipt(t *testing.T) {
	if got := Validate(validUpload()); len(got) != 0 {
		t.Fatalf("expected no findings, got %#v", got)
	}
}

func TestSameTransferRequiresMatchingChecksum(t *testing.T) {
	a, b := validUpload(), validUpload()
	b.Checksum = "sha256:def"
	if SameTransfer(a, b) {
		t.Fatal("different checksums must not deduplicate")
	}
}
