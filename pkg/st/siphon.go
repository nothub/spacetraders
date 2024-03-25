package st

type Siphon struct {

	// Symbol of the ship that executed the siphon.
	//
	//   >= 1 characters
	//
	// (required)
	ShipSymbol string `json:"shipSymbol"`

	// A yield from the siphon operation.
	//
	// (required)
	Yield struct {

		// The good's symbol.
		//
		// (required)
		Symbol TradeGoodSymbol `json:"symbol"`

		// The number of units siphoned that were placed into the ship's cargo hold.
		//
		// (required)
		Units int `json:"units"`
	} `json:"yield"`
}
