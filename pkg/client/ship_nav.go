package st

import "time"

// The navigation information of the ship.
type ShipNav struct {

	// The symbol of the system.
	//
	//   >= 1 characters
	//
	// (required)
	SystemSymbol string `json:"systemSymbol"`

	// The symbol of the waypoint.
	//
	// (required)
	WaypointSymbol string `json:"waypointSymbol"`

	// The routing information for the ship's most recent transit or current location.
	//
	// (required)
	Route ShipNavRoute `json:"route"`

	// The current status of the ship
	//
	// (required)
	Status ShipNavStatus `json:"status"`

	// The ship's set speed when traveling between waypoints or systems.
	//
	// (required)
	FlightMode ShipNavFlightMode `json:"flightMode"`
}

// The routing information for the ship's most recent transit or current location.
type ShipNavRoute struct {

	// The destination or departure of a ships nav route.
	//
	// (required)
	Destination ShipNavRouteWaypoint `json:"destination"`

	// The destination or departure of a ships nav route.
	//
	// (required)
	Origin ShipNavRouteWaypoint `json:"origin"`

	// The date time of the ship's departure.
	//
	// (required)
	DepartureTime time.Time `json:"departureTime"`

	// The date time of the ship's arrival.
	// If the ship is in-transit, this is the expected time of arrival.
	//
	// (required)
	Arrival time.Time `json:"arrival"`
}

type ShipNavRouteWaypoint struct {

	// The symbol of the waypoint.
	//
	//   >= 1 characters
	//
	// (required)
	Symbol string `json:"symbol"`

	// The type of waypoint.
	//
	// (required)
	Type WaypointType `json:"type"`

	// The symbol of the system.
	//
	//   >= 1 characters
	//
	// (required)
	SystemSymbol string `json:"systemSymbol"`

	// Position in the universe in the x axis.
	//
	// (required)
	X int `json:"x"`

	// Position in the universe in the y axis.
	//
	// (required)
	Y int `json:"y"`
}

// The current status of the ship
type ShipNavStatus string

const (
	ShipNavStatusInTransit ShipNavStatus = "IN_TRANSIT"
	ShipNavStatusInOrbit   ShipNavStatus = "IN_ORBIT"
	ShipNavStatusDocked    ShipNavStatus = "DOCKED"
)

// The ship's set speed when traveling between waypoints or systems.
type ShipNavFlightMode string

const (
	ShipNavFlightModeDrift   ShipNavFlightMode = "DRIFT"
	ShipNavFlightModeStealth ShipNavFlightMode = "STEALTH"
	ShipNavFlightModeCruise  ShipNavFlightMode = "CRUISE"
	ShipNavFlightModeBurn    ShipNavFlightMode = "BURN"
)
