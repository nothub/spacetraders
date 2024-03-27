package st

// The ship's crew service and maintain the ship's systems and equipment.
type ShipCrew struct {

	// The current number of crew members on the ship.
	//
	// (required)
	Current int `json:"current"`

	// The minimum number of crew members required to maintain the ship.
	//
	// (required)
	Required int `json:"required"`

	// The maximum number of crew members the ship can support.
	//
	// (required)
	Capacity int `json:"capacity"`

	// The rotation of crew shifts.
	// A stricter shift improves the ship's performance.
	// A more relaxed shift improves the crew's morale.
	//
	// (required)
	Rotation ShipCrewRotation `json:"rotation"`

	// A rough measure of the crew's morale.
	// A higher morale means the crew is happier and more productive.
	// A lower morale means the ship is more prone to accidents.
	//
	//   >= 0 && <= 100
	//
	// (required)
	Morale int `json:"morale"`

	// The amount of credits per crew member paid per hour.
	// Wages are paid when a ship docks at a civilized waypoint.
	//
	//   >= 0
	//
	// (required)
	Wages int `json:"wages"`
}

type ShipCrewRotation string

const (
	ShipCrewRotationStrict  ShipCrewRotation = "STRICT"
	ShipCrewRotationRelaxed ShipCrewRotation = "RELAXED"
)
