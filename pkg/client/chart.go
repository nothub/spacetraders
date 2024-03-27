package st

import "time"

// The chart of a system or waypoint, which makes the location visible to other agents.
type Chart struct {

	// The symbol of the waypoint.
	//
	//   >= 1 characters
	WaypointSymbol string `json:"waypointSymbol"`

	// The agent that submitted the chart for this waypoint.
	SubmittedBy string `json:"submittedBy"`

	// The time the chart for this waypoint was submitted.
	SubmittedOn time.Time `json:"submittedOn"`
}
