package consent

import (
	"testing"
	"time"
)

func TestValidateCompleteRecord(t *testing.T) {
	r := Record{EventID: "event-1", NoticeVersion: "v2", CapturedAt: time.Now(), Scopes: []Scope{PrivateAlbum}, WithdrawalURL: "https://example.com/remove"}
	if got := Validate(r); len(got) != 0 {
		t.Fatalf("expected no findings, got %#v", got)
	}
}

func TestValidateRejectsUnknownScope(t *testing.T) {
	r := Record{EventID: "event-1", NoticeVersion: "v2", CapturedAt: time.Now(), Scopes: []Scope{"unknown"}, WithdrawalURL: "https://example.com/remove"}
	if got := Validate(r); len(got) != 1 {
		t.Fatalf("expected one finding, got %#v", got)
	}
}
