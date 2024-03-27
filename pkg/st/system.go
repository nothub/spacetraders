package st

import (
	"strconv"
)

type System struct {
	Symbol       string     `json:"symbol"`
	SectorSymbol string     `json:"sectorSymbol"`
	Type         SystemType `json:"type"`
	X            int        `json:"x"`
	Y            int        `json:"y"`
	Waypoints    []struct {
		Symbol   string       `json:"symbol"`
		Type     WaypointType `json:"type"`
		X        int          `json:"x"`
		Y        int          `json:"y"`
		Orbitals []struct {
			Symbol string `json:"symbol"`
		} `json:"orbitals"`
		Orbits string `json:"orbits"`
	} `json:"waypoints"`
	Factions []SystemFaction `json:"factions"`
}

type SystemFaction struct {
	Symbol FactionSymbol `json:"symbol"`
}

type SystemSymbol string

// The type of system.
type SystemType string

const (
	SystemTypeNeutronStar SystemType = "NEUTRON_STAR"
	SystemTypeRedStar     SystemType = "RED_STAR"
	SystemTypeOrangeStar  SystemType = "ORANGE_STAR"
	SystemTypeBlueStar    SystemType = "BLUE_STAR"
	SystemTypeYoungStar   SystemType = "YOUNG_STAR"
	SystemTypeWhiteDwarf  SystemType = "WHITE_DWARF"
	SystemTypeBlackHole   SystemType = "BLACK_HOLE"
	SystemTypeHypergiant  SystemType = "HYPERGIANT"
	SystemTypeNebula      SystemType = "NEBULA"
	SystemTypeUnstable    SystemType = "UNSTABLE"
)

type SystemWaypoint struct {
	Symbol   string     `json:"symbol"`
	Type     SystemType `json:"type"`
	X        int        `json:"x"`
	Y        int        `json:"y"`
	Orbitals []struct {

		// The symbol of the orbiting waypoint.
		//
		//   >= 1 characters
		//
		// (required)
		Symbol string `json:"symbol"`
	} `json:"orbitals"`
	Orbits string `json:"orbits"`
}

type ConnectedSystem struct {

	// The symbol of the system.
	//
	//   >= 1 characters
	//
	// (required)
	Symbol string `json:"symbol"`

	// The sector of this system.
	//
	//   >= 1 characters
	//
	// (required)
	SectorSymbol string `json:"sectorSymbol"`

	// The type of system.
	//
	// (required)
	Type SystemType `json:"type"`

	// The symbol of the faction that owns the connected jump gate in the system.
	FactionSymbol string `json:"factionSymbol"`

	// Position in the universe in the x axis.
	//
	// (required)
	X int `json:"x"`

	// Position in the universe in the y axis.
	//
	// (required)
	Y int `json:"y"`

	// The distance of this system to the connected Jump Gate.
	//
	// (required)
	Distance int `json:"distance"`
}

func ListSystems(limit int, page int) (systems []System, meta Meta, err error) {
	var dto struct {
		Data []System `json:"data"`
		Meta Meta     `json:"meta"`
	}

	err = get(BaseUrl+"/systems", map[string]string{
		"limit": strconv.Itoa(limit),
		"page":  strconv.Itoa(page),
	}, &dto)
	if err != nil {
		return systems, meta, err
	}

	return dto.Data, dto.Meta, nil
}

func GetSystem(systemSymbol string) (systems System, err error) {
	var dto struct {
		Data System `json:"data"`
		Meta Meta   `json:"meta"`
	}

	err = get(BaseUrl+"/systems/"+systemSymbol, nil, &dto)
	if err != nil {
		return systems, err
	}

	return dto.Data, nil
}

func ListWaypointsInSystem() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/32186cf59e324-list-waypoints-in-system
	return nil
}

func GetWaypoint() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/58e66f2fa8c82-get-waypoint
	return nil
}

func GetMarket() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/a4fed7a0221e0-get-market
	return nil
}

func GetShipyard() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/460fe70c0e4c2-get-shipyard
	return nil
}

func GetJumpGate() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/decd101af6414-get-jump-gate
	return nil
}

func GetConstructionSite() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/c4db8d0c02144-get-construction-site
	return nil
}

func SupplyConstructionSite() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/c78f7d4f260ef-supply-construction-site
	return nil
}
