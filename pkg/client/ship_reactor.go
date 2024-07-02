package st

// The reactor of the ship. The reactor is responsible for powering the ship's systems and weapons.
type ShipReactor struct {

	// Symbol of the reactor.
	//
	// (required)
	Symbol ShipReactorSymbol `json:"symbol"`

	// Name of the reactor.
	//
	// (required)
	Name string `json:"name"`

	// Description of the reactor.
	//
	// (required)
	Description string `json:"description"`

	// The repairable condition of a component. A value of 0 indicates the
	// component needs significant repairs, while a value of 1 indicates the
	// component is in near perfect condition. As the condition of a component
	// is repaired, the overall integrity of the component decreases.
	//
	//   >= 0 && <= 1
	//
	// (required)
	Condition float32 `json:"condition"`

	// The overall integrity of the component, which determines the performance
	// of the component. A value of 0 indicates that the component is almost
	// completely degraded, while a value of 1 indicates that the component is
	// in near perfect condition. The integrity of the component is
	// non-repairable, and represents permanent wear over time.
	//
	//   >= 0 && <= 1
	//
	// (required)
	Integrity float32 `json:"integrity"`

	// The amount of power provided by this reactor. The more power a reactor
	// provides to the ship, the lower the cooldown it gets when using a module
	// or mount that taxes the ship's power.
	//
	//   >= 1
	//
	// (required)
	PowerOutput int `json:"powerOutput"`

	// The requirements for installation on a ship
	//
	// (required)
	Requirements ShipRequirements `json:"requirements"`
}

type ShipReactorSymbol string

const (
	ShipReactorReactorSolar1      ShipReactorSymbol = "REACTOR_SOLAR_I"
	ShipReactorReactorFusion1     ShipReactorSymbol = "REACTOR_FUSION_I"
	ShipReactorReactorFission1    ShipReactorSymbol = "REACTOR_FISSION_I"
	ShipReactorReactorChemical1   ShipReactorSymbol = "REACTOR_CHEMICAL_I"
	ShipReactorReactorAntimatter1 ShipReactorSymbol = "REACTOR_ANTIMATTER_I"
)
