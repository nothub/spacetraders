package st

// A module can be installed in a ship and provides a set
// of capabilities such as storage space or quarters for crew.
// Module installations are permanent.
type ShipModule struct {

	// The symbol of the module.
	//
	// (required)
	Symbol ShipModuleSymbol `json:"symbol"`

	// Modules that provide capacity, such as cargo hold or crew quarters
	// will show this value to denote how much of a bonus the module grants.
	//
	//   >= 0
	//
	// (required)
	Capacity int `json:"capacity"`

	// Modules that have a range will such as a sensor array show this value
	// to denote how far can the module reach with its capabilities.
	//
	//   >= 0
	//
	// (required)
	Range int `json:"range"`

	// Name of this module.
	//
	// (required)
	Name string `json:"name"`

	// Description of this module.
	//
	// (required)
	Description string `json:"description"`

	// The requirements for installation on a ship
	//
	// (required)
	Requirements ShipRequirements `json:"requirements"`
}

type ShipModuleSymbol string

const (
	ShipModuleMineralProcessor1 ShipModuleSymbol = "MODULE_MINERAL_PROCESSOR_I"
	ShipModuleGasProcessor1     ShipModuleSymbol = "MODULE_GAS_PROCESSOR_I"
	ShipModuleCargoHold1        ShipModuleSymbol = "MODULE_CARGO_HOLD_I"
	ShipModuleCargoHold2        ShipModuleSymbol = "MODULE_CARGO_HOLD_II"
	ShipModuleCargoHold3        ShipModuleSymbol = "MODULE_CARGO_HOLD_III"
	ShipModuleCrewQuarters1     ShipModuleSymbol = "MODULE_CREW_QUARTERS_I"
	ShipModuleEnvoyQuarters1    ShipModuleSymbol = "MODULE_ENVOY_QUARTERS_I"
	ShipModulePassengerCabin1   ShipModuleSymbol = "MODULE_PASSENGER_CABIN_I"
	ShipModuleMicroRefinery1    ShipModuleSymbol = "MODULE_MICRO_REFINERY_I"
	ShipModuleOreRefinery1      ShipModuleSymbol = "MODULE_ORE_REFINERY_I"
	ShipModuleFuelRefinery1     ShipModuleSymbol = "MODULE_FUEL_REFINERY_I"
	ShipModuleScienceLab1       ShipModuleSymbol = "MODULE_SCIENCE_LAB_I"
	ShipModuleJumpDrive1        ShipModuleSymbol = "MODULE_JUMP_DRIVE_I"
	ShipModuleJumpDrive2        ShipModuleSymbol = "MODULE_JUMP_DRIVE_II"
	ShipModuleJumpDrive3        ShipModuleSymbol = "MODULE_JUMP_DRIVE_III"
	ShipModuleWarpDrive1        ShipModuleSymbol = "MODULE_WARP_DRIVE_I"
	ShipModuleWarpDrive2        ShipModuleSymbol = "MODULE_WARP_DRIVE_II"
	ShipModuleWarpDrive3        ShipModuleSymbol = "MODULE_WARP_DRIVE_III"
	ShipModuleShieldGenerator1  ShipModuleSymbol = "MODULE_SHIELD_GENERATOR_I"
	ShipModuleShieldGenerator2  ShipModuleSymbol = "MODULE_SHIELD_GENERATOR_II"
)
