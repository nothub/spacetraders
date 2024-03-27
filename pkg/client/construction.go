package st

// The construction details of a waypoint.
type Construction struct {

	// The symbol of the waypoint.
	//
	// (required)
	Symbol string `json:"symbol"`

	// The materials required to construct the waypoint.
	//
	// (required)
	Materials []ConstructionMaterial `json:"materials"`

	// Whether the waypoint has been constructed.
	//
	// (required)
	IsComplete bool `json:"isComplete"`
}

// The details of the required construction materials for a given waypoint under construction.
type ConstructionMaterial struct {

	// The good's symbol.
	//
	// (required)
	TradeSymbol TradeGoodSymbol `json:"tradeSymbol"`

	// The number of units required.
	//
	// (required)
	Required int `json:"required"`

	// The number of units fulfilled toward the required amount.
	//
	// (required)
	Fulfilled int `json:"fulfilled"`
}
