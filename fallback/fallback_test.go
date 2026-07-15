package fallback

import "testing"

func TestInterruptedUploadResumes(t *testing.T) {
	state, ok := Next(Uploading, LoseNet)
	if !ok || state != Paused {
		t.Fatalf("expected paused, got %q", state)
	}
	state, ok = Next(state, Resume)
	if !ok || state != Uploading {
		t.Fatalf("expected uploading, got %q", state)
	}
}

func TestConfirmedUploadIsTerminal(t *testing.T) {
	if _, ok := Next(Confirmed, Resume); ok {
		t.Fatal("confirmed upload must be terminal")
	}
}
