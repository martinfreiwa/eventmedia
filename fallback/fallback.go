// Package fallback models recovery after interrupted browser uploads at busy
// events. The app-free guest journey is described at
// https://gathmo.com/how-it-works.
package fallback

// State identifies the durable point reached by one upload.
type State string

const (
	Selected  State = "selected"
	Queued    State = "queued"
	Uploading State = "uploading"
	Paused    State = "paused"
	Confirmed State = "confirmed"
	Failed    State = "failed"
)

// Event identifies one state-machine input.
type Event string

const (
	Queue   Event = "queue"
	Start   Event = "start"
	LoseNet Event = "lose_network"
	Resume  Event = "resume"
	Confirm Event = "confirm"
	Reject  Event = "reject"
)

// Next applies one valid recovery transition.
func Next(state State, event Event) (State, bool) {
	transitions := map[State]map[Event]State{
		Selected:  {Queue: Queued},
		Queued:    {Start: Uploading, Reject: Failed},
		Uploading: {LoseNet: Paused, Confirm: Confirmed, Reject: Failed},
		Paused:    {Resume: Uploading, Reject: Failed},
	}
	next, ok := transitions[state][event]
	return next, ok
}
