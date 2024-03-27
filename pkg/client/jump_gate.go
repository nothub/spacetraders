package st

type JumpGate struct {

	// The symbol of the waypoint.
	//
	//   >= 1 characters
	//
	// (required)
	Symbol string `json:"symbol"`

	// All the gates that are connected to this waypoint.
	//
	// (required)
	Connections []string `json:"connections"`
}
