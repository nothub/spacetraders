package st

// The engine determines how quickly a ship travels between waypoints.
type ShipEngine struct {

	// The symbol of the engine.
	//
	// (required)
	Symbol ShipEngineSymbol `json:"symbol"`

	// The name of the engine.
	//
	// (required)
	Name string `json:"name"`

	// The description of the engine.
	//
	// (required)
	Description string `json:"description"`

	// The repairable condition of a component.
	// A value of 0 indicates the component needs significant repairs,
	// while a value of 1 indicates the component is in near perfect condition.
	// As the condition of a component is repaired, the overall integrity of the component decreases.
	//
	//   >= 0 && <= 1
	//
	// (required)
	Condition float32 `json:"condition"`

	// The overall integrity of the component, which determines the performance of the component.
	// A value of 0 indicates that the component is almost completely degraded,
	// while a value of 1 indicates that the component is in near perfect condition.
	// The integrity of the component is non-repairable, and represents permanent wear over time.
	//
	//   >= 0 && <= 1
	//
	// (required)
	Integrity float32 `json:"integrity"`

	// The speed stat of this engine.
	// The higher the speed, the faster a ship can travel from one point to another.
	// Reduces the time of arrival when navigating the ship.
	//
	//   >= 1
	//
	// (required)
	Speed int `json:"speed"`

	// The requirements for installation on a ship
	//
	// (required)
	Requirements ShipRequirements `json:"requirements"`
}

type ShipEngineSymbol string

const (
	ShipEngineImpulseDrive1 ShipEngineSymbol = "ENGINE_IMPULSE_DRIVE_I"
	ShipEngineIonDrive1     ShipEngineSymbol = "ENGINE_ION_DRIVE_I"
	ShipEngineIonDrive2     ShipEngineSymbol = "ENGINE_ION_DRIVE_II"
	ShipEngineHyperDrive1   ShipEngineSymbol = "ENGINE_HYPER_DRIVE_I"
)
