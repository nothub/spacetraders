package st

// The public registration information of the ship
type ShipRegistration struct {

	// The agent's registered name of the ship
	//
	//   >= 1 characters
	//
	// (required)
	Name string `json:"name"`

	// The symbol of the faction the ship is registered with
	//
	//   >= 1 characters
	//
	// (required)
	FactionSymbol string `json:"factionSymbol"`

	// The registered role of the ship
	//
	// (required)
	Role ShipRole `json:"role"`
}
