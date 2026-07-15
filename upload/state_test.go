package upload

import "testing"

func TestExpectedTransitions(t *testing.T) {
	if !CanTransition(Transferring, Accepted) {
		t.Fatal("expected transferring to accepted transition")
	}
	if CanTransition(Selected, Ready) {
		t.Fatal("selected must not transition directly to ready")
	}
}

func TestReceiptValid(t *testing.T) {
	receipt := Receipt{UploadID: "upload-1", ExpectedBytes: 100, ReceivedBytes: 100, Accepted: true}
	if !ReceiptValid(receipt) {
		t.Fatal("expected complete accepted receipt to be valid")
	}
}

func TestReceiptRejectsOptimisticCompletion(t *testing.T) {
	receipt := Receipt{UploadID: "upload-1", ExpectedBytes: 100, ReceivedBytes: 100, Accepted: false}
	if ReceiptValid(receipt) {
		t.Fatal("receipt must not be valid before server acceptance")
	}
}
