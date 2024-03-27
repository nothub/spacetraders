package st

type ScannedShip struct {
	Symbol       string           `json:"symbol"`
	Registration ShipRegistration `json:"registration"`
	Nav          ShipNav          `json:"nav"`
	Frame        struct {
		Symbol string `json:"symbol"`
	} `json:"frame"`
	Reactor struct {
		Symbol string `json:"symbol"`
	} `json:"reactor"`
	Engine struct {
		Symbol string `json:"symbol"`
	} `json:"engine"`
	Mounts []struct {
		Symbol string `json:"symbol"`
	} `json:"mounts"`
}

type ScannedSystem struct {
	Symbol       string     `json:"symbol"`
	SectorSymbol string     `json:"sectorSymbol"`
	Type         SystemType `json:"type"`
	X            int        `json:"x"`
	Y            int        `json:"y"`
	Distance     int        `json:"distance"`
}

type ScannedWaypoint struct {
	Symbol       string       `json:"symbol"`
	Type         WaypointType `json:"type"`
	SystemSymbol string       `json:"systemSymbol"`
	X            int          `json:"x"`
	Y            int          `json:"y"`
	Orbitals     []struct {
		Symbol string `json:"symbol"`
	} `json:"orbitals"`
	Faction struct {
		Symbol FactionSymbol `json:"symbol"`
	} `json:"faction"`
	Traits []WaypointTrait `json:"traits"`
	Chart  Chart           `json:"chart"`
}
