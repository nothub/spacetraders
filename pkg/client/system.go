package st

import (
	"strconv"
	"strings"
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

func ListWaypointsInSystem(
	systemSymbol string,
	waypointTraits []WaypointTraitSymbol,
	waypointType WaypointType,
	limit int,
	page int,
) (
	waypoints []Waypoint,
	err error,
) {
	var dto struct {
		Data []Waypoint `json:"data"`
		Meta Meta       `json:"meta"`
	}

	traits := make([]string, len(waypointTraits))
	for _, t := range waypointTraits {
		traits = append(traits, string(t))
	}
	err = get(BaseUrl+"/systems/"+systemSymbol+"/waypoints", map[string]string{
		"limit":  strconv.Itoa(limit),
		"page":   strconv.Itoa(page),
		"traits": strings.Join(traits, ","),
		"type":   string(waypointType),
	}, &dto)
	if err != nil {
		return waypoints, err
	}

	return dto.Data, nil
}

func GetWaypoint(systemSymbol string, waypointSymbol string) (waypoint Waypoint, err error) {
	var dto struct {
		Data Waypoint `json:"data"`
		Meta Meta     `json:"meta"`
	}

	err = get(BaseUrl+"/systems/"+systemSymbol+"/waypoints/"+waypointSymbol, nil, &dto)
	if err != nil {
		return waypoint, err
	}

	return dto.Data, nil
}

func GetMarket(systemSymbol string, waypointSymbol string) (market Market, err error) {
	var dto struct {
		Data Market `json:"data"`
		Meta Meta   `json:"meta"`
	}

	err = get(BaseUrl+"/systems/"+systemSymbol+"/waypoints/"+waypointSymbol+"/market", nil, &dto)
	if err != nil {
		return market, err
	}

	return dto.Data, nil
}

func GetShipyard(systemSymbol string, waypointSymbol string) (shipyard Shipyard, err error) {
	var dto struct {
		Data Shipyard `json:"data"`
		Meta Meta     `json:"meta"`
	}

	err = get(BaseUrl+"/systems/"+systemSymbol+"/waypoints/"+waypointSymbol+"/shipyard", nil, &dto)
	if err != nil {
		return shipyard, err
	}

	return dto.Data, nil
}

func GetJumpGate(systemSymbol string, waypointSymbol string) (jumpGate JumpGate, err error) {
	var dto struct {
		Data JumpGate `json:"data"`
		Meta Meta     `json:"meta"`
	}

	err = get(BaseUrl+"/systems/"+systemSymbol+"/waypoints/"+waypointSymbol+"/jump-gate", nil, &dto)
	if err != nil {
		return jumpGate, err
	}

	return dto.Data, nil
}

func GetConstructionSite(systemSymbol string, waypointSymbol string) (construction Construction, err error) {
	var dto struct {
		Data Construction `json:"data"`
		Meta Meta         `json:"meta"`
	}

	err = get(BaseUrl+"/systems/"+systemSymbol+"/waypoints/"+waypointSymbol+"/construction", nil, &dto)
	if err != nil {
		return construction, err
	}

	return dto.Data, nil
}

// Supply a construction site with the specified good. Requires a waypoint that
// is under construction.
//
// The good must be in your ship's cargo. The good will be removed from your
// ship's cargo and added to the construction site's materials.
func SupplyConstructionSite(
	systemSymbol string,
	waypointSymbol string,
	shipSymbol string,
	tradeSymbol TradeGoodSymbol,
	units int,
) (
	construction Construction,
	cargo ShipCargo,
	err error,
) {
	var dto struct {
		Data struct {
			Construction Construction `json:"construction"`
			Cargo        ShipCargo    `json:"cargo"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/systems/"+systemSymbol+"/waypoints/"+waypointSymbol+"/construction/supply",
		nil,
		struct {
			ShipSymbol  string `json:"shipSymbol"`
			TradeSymbol string `json:"tradeSymbol"`
			Units       int    `json:"units"`
		}{
			ShipSymbol:  shipSymbol,
			TradeSymbol: string(tradeSymbol),
			Units:       units,
		},
		&dto)
	if err != nil {
		return construction, cargo, err
	}

	return dto.Data.Construction, dto.Data.Cargo, nil
}
