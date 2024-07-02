package st

// The frame of the ship.
// The frame determines the number of modules and mounting points of the ship, as well as base fuel capacity.
// As the condition of the frame takes more wear, the ship will become more sluggish and less maneuverable.
type ShipFrame struct {

	// Symbol of the frame.
	//
	// (required)
	Symbol ShipFrameSymbol `json:"symbol"`

	// Name of the frame.
	//
	// (required)
	Name string `json:"name"`

	// Description of the frame.
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

	// The amount of slots that can be dedicated to modules installed in the ship.
	// Each installed module take up a number of slots,
	// and once there are no more slots,
	// no new modules can be installed.
	//
	//   >= 0
	//
	// (required)
	ModuleSlots int `json:"moduleSlots"`

	// The amount of slots that can be dedicated to mounts installed in the ship.
	// Each installed mount takes up a number of points,
	// and once there are no more points remaining,
	// no new mounts can be installed.
	//
	//   >= 0
	//
	// (required)
	MountingPoints int `json:"mountingPoints"`

	// The maximum amount of fuel that can be stored in this ship.
	// When refueling, the ship will be refueled to this amount.
	//
	//   >= 0
	//
	// (required)
	FuelCapacity int `json:"fuelCapacity"`

	// The requirements for installation on a ship
	//
	// (required)
	Requirements ShipRequirements `json:"requirements"`
}

type ShipFrameSymbol string

const (
	ShipFrameProbe          ShipFrameSymbol = "FRAME_PROBE"
	ShipFrameDrone          ShipFrameSymbol = "FRAME_DRONE"
	ShipFrameInterceptor    ShipFrameSymbol = "FRAME_INTERCEPTOR"
	ShipFrameRacer          ShipFrameSymbol = "FRAME_RACER"
	ShipFrameFighter        ShipFrameSymbol = "FRAME_FIGHTER"
	ShipFrameFrigate        ShipFrameSymbol = "FRAME_FRIGATE"
	ShipFrameShuttle        ShipFrameSymbol = "FRAME_SHUTTLE"
	ShipFrameExplorer       ShipFrameSymbol = "FRAME_EXPLORER"
	ShipFrameMiner          ShipFrameSymbol = "FRAME_MINER"
	ShipFrameLightFreighter ShipFrameSymbol = "FRAME_LIGHT_FREIGHTER"
	ShipFrameHeavyFreighter ShipFrameSymbol = "FRAME_HEAVY_FREIGHTER"
	ShipFrameTransport      ShipFrameSymbol = "FRAME_TRANSPORT"
	ShipFrameDestroyer      ShipFrameSymbol = "FRAME_DESTROYER"
	ShipFrameCruiser        ShipFrameSymbol = "FRAME_CRUISER"
	ShipFrameCarrier        ShipFrameSymbol = "FRAME_CARRIER"
)
