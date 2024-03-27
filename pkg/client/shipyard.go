package st

import "time"

type Shipyard struct {

	// The symbol of the shipyard. The symbol is the same as the waypoint where
	// the shipyard is located.
	Symbol string `json:"symbol"`

	// The list of ship types available for purchase at this shipyard.
	ShipTypes []struct {
		Type ShipType `json:"type"`
	} `json:"shipTypes"`

	// The list of recent transactions at this shipyard.
	Transactions []ShipyardTransaction `json:"transactions"`

	// The ships that are currently available for purchase at the shipyard.
	Ships []ShipyardShip `json:"ships"`

	// The fee to modify a ship at this shipyard. This includes installing or
	// removing modules and mounts on a ship. In the case of mounts, the fee is
	// a flat rate per mount. In the case of modules, the fee is per slot the
	// module occupies.
	ModificationsFee int `json:"modificationsFee"`
}

type ShipyardShip struct {
	Type        ShipType    `json:"type"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Supply      SupplyLevel `json:"supply"`

	// The activity level of a trade good. If the good is an import, this
	// represents how strong consumption is. If the good is an export, this
	// represents how strong the production is for the good. When activity is
	// strong, consumption or production is near maximum capacity. When
	// activity is weak, consumption or production is near minimum capacity.
	Activity      ActivityLevel `json:"activity"`
	PurchasePrice int           `json:"purchasePrice"`

	// The frame of the ship. The frame determines the number of modules and
	// mounting points of the ship, as well as base fuel capacity. As the
	// condition of the frame takes more wear, the ship will become more
	// sluggish and less maneuverable.
	Frame ShipFrame `json:"frame"`

	// The reactor of the ship. The reactor is responsible for powering the
	// ship's systems and weapons.
	Reactor ShipReactor `json:"reactor"`

	// The engine determines how quickly a ship travels between waypoints.
	Engine ShipEngine `json:"engine"`

	// A module can be installed in a ship and provides a set of capabilities
	// such as storage space or quarters for crew. Module installations are
	// permanent.
	Modules []ShipModule `json:"modules"`

	// A mount is installed on the exterier of a ship.
	Mounts []ShipMount `json:"mounts"`

	Crew struct {
		Required int `json:"required"`
		Capacity int `json:"capacity"`
	} `json:"crew"`
}

type ShipyardTransaction struct {
	WaypointSymbol string    `json:"waypointSymbol"`
	ShipSymbol     string    `json:"shipSymbol"`
	ShipType       ShipType  `json:"shipType"`
	Price          int       `json:"price"`
	AgentSymbol    string    `json:"agentSymbol"`
	Timestamp      time.Time `json:"timestamp"`
}
