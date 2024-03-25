package st

import (
	"net/url"
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

type systemsDto struct {
	Data []System `json:"data"`
	Meta Meta     `json:"meta"`
}

func ListSystems(limit int, page int) (systems []System, meta Meta, err error) {
	var dto systemsDto

	u, err := url.Parse(BaseUrl + "/systems")
	if err != nil {
		return dto.Data, dto.Meta, err
	}

	u.Query().Set("limit", strconv.Itoa(limit))
	u.Query().Set("page", strconv.Itoa(page))

	err = get(u.String(), &dto)
	if err != nil {
		return dto.Data, dto.Meta, err
	}

	return dto.Data, dto.Meta, nil
}
