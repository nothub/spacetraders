package st

// Ship details.
type Ship struct {

	// The globally unique identifier of the ship in the following format:
	//
	//   [AGENT_SYMBOL]-[HEX_ID]
	//
	// (required)
	Symbol string `json:"symbol"`

	// The public registration information of the ship
	//
	// (required)
	Registration ShipRegistration `json:"registration"`

	// The navigation information of the ship.
	//
	// (required)
	Nav ShipNav `json:"nav"`

	// The ship's crew service and maintain the ship's systems and equipment.
	//
	// (required)
	Crew ShipCrew `json:"crew"`

	// The frame of the ship. The frame determines the number of modules and
	// mounting points of the ship, as well as base fuel capacity. As the
	// condition of the frame takes more wear, the ship will become more
	// sluggish and less maneuverable.
	//
	// (required)
	Frame ShipFrame `json:"frame"`

	// The reactor of the ship. The reactor is responsible for powering the
	// ship's systems and weapons.
	//
	// (required)
	Reactor ShipReactor `json:"reactor"`

	// The engine determines how quickly a ship travels between waypoints.
	//
	// (required)
	Engine ShipEngine `json:"engine"`

	// A cooldown is a period of time in which a ship cannot perform certain
	// actions.
	//
	// (required)
	Cooldown Cooldown `json:"cooldown"`

	// Modules installed in this ship.
	//
	// (required)
	Modules []ShipModule `json:"modules"`

	// Mounts installed in this ship.
	//
	// (required)
	Mounts []ShipMount `json:"mounts"`

	// Ship cargo details.
	//
	// (required)
	Cargo ShipCargo `json:"cargo"`

	// Details of the ship's fuel tanks including how much fuel was consumed
	// during the last transit or action.
	//
	// (required)
	Fuel ShipFuel `json:"fuel"`
}
