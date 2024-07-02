package st

// Ship cargo details.
type ShipCargo struct {

	// The max number of items that can be stored in the cargo hold.
	//
	//   >= 0
	//
	// (required)
	Capacity int `json:"capacity"`

	// The number of items currently stored in the cargo hold.
	//
	//   >= 0
	//
	// (required)
	Units int `json:"units"`

	// The items currently in the cargo hold.
	//
	// (required)
	Inventory []ShipCargoItem `json:"inventory"`
}

// The type of cargo item and the number of units.
type ShipCargoItem struct {

	// The good's symbol.
	//
	// (required)
	Symbol TradeGoodSymbol `json:"symbol"`

	// The name of the cargo item type.
	//
	// (required)
	Name string `json:"name"`

	// The description of the cargo item type.
	//
	// (required)
	Description string `json:"description"`

	// The number of units of the cargo item.
	//
	//   >= 1
	//
	// (required)
	Units int `json:"units"`
}
