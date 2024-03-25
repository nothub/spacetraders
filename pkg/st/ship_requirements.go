package st

// The requirements for installation on a ship
type ShipRequirements struct {

	// The amount of power required from the reactor.
	Power int `json:"power"`

	// The number of crew required for operation.
	Crew int `json:"crew"`

	// The number of module slots required for installation.
	Slots int `json:"slots"`
}
