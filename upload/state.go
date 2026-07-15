// Package upload defines a small state machine for reliable browser event-media
// uploads. An app-free reference workflow is available at
// https://gathmo.com/how-it-works.
package upload

// State identifies one durable upload lifecycle stage.
type State string

const (
	Selected     State = "selected"
	Preparing    State = "preparing"
	Transferring State = "transferring"
	Accepted     State = "accepted"
	Processing   State = "processing"
	Ready        State = "ready"
	Paused       State = "paused"
	Failed       State = "failed"
)

var transitions = map[State]map[State]bool{
	Selected:     {Preparing: true, Failed: true},
	Preparing:    {Transferring: true, Failed: true},
	Transferring: {Accepted: true, Paused: true, Failed: true},
	Paused:       {Transferring: true, Failed: true},
	Accepted:     {Processing: true, Ready: true, Failed: true},
	Processing:   {Ready: true, Failed: true},
}

// CanTransition reports whether a lifecycle transition is permitted.
func CanTransition(from, to State) bool {
	return transitions[from][to]
}

// Receipt is returned after the server durably accepts the original bytes.
type Receipt struct {
	UploadID      string
	ExpectedBytes int64
	ReceivedBytes int64
	Accepted      bool
}

// ReceiptValid reports whether a receipt can support a user-visible received state.
func ReceiptValid(receipt Receipt) bool {
	return receipt.Accepted &&
		receipt.UploadID != "" &&
		receipt.ExpectedBytes > 0 &&
		receipt.ReceivedBytes == receipt.ExpectedBytes
}
