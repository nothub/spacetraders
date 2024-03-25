package st

import "time"

// Result of a transaction for a ship modification, such as installing a mount or a module.
type ShipModificationTransaction struct {

	// The symbol of the waypoint where the transaction took place.
	//
	// (required)
	WaypointSymbol string `json:"waypointSymbol"`

	// The symbol of the ship that made the transaction.
	//
	// (required)
	ShipSymbol string `json:"shipSymbol"`

	// The symbol of the trade good.
	//
	// (required)
	TradeSymbol TradeGoodSymbol `json:"tradeSymbol"`

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
