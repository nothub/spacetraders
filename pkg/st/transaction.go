package st

import "time"

// Result of a transaction.
type Transaction struct {

	// The symbol of the waypoint.
	//
	//   >= 1 characters
	//
	// (required)
	WaypointSymbol string `json:"waypointSymbol"`

	// The symbol of the ship.
	//
	// (required)
	ShipSymbol string `json:"shipSymbol"`

	// The total price of the transaction.
	//
	//   >= 0
	//
	// (required)
	TotalPrice int `json:"totalPrice"`

	// The timestamp of the transaction.
	//
	// (required)
	Timestamp time.Time `json:"timestamp"`
}

// Result of a repair transaction.
type RepairTransaction Transaction

// Result of a scrap transaction.
type ScrapTransaction Transaction
