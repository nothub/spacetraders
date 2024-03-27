package st

// Extraction details.
type Extraction struct {

	// Symbol of the ship that executed the extraction.
	//
	// (required)
	ShipSymbol string `json:"shipSymbol"`

	// A yield from the extraction operation.
	//
	// (required)
	Yield struct {

		// The good's symbol.
		//
		// (required)
		Symbol TradeGoodSymbol `json:"symbol"`

		// The number of units extracted that were placed into the ship's cargo hold.
		//
		// (required)
		Units int `json:"units"`
	} `json:"yield"`
}
