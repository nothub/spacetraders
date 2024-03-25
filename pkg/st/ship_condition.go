package st

// The repairable condition of a component.
// A value of 0 indicates the component needs significant repairs,
// while a value of 1 indicates the component is in near perfect condition.
// As the condition of a component is repaired, the overall integrity of the component decreases.
//
//	>= 0 && <= 1
type ShipComponentCondition float32

// The overall integrity of the component, which determines the performance of the component.
// A value of 0 indicates that the component is almost completely degraded,
// while a value of 1 indicates that the component is in near perfect condition.
// The integrity of the component is non-repairable, and represents permanent wear over time.
//
//	>= 0 && <= 1
type ShipComponentIntegrity float32

// An event that represents damage or wear to a ship's reactor, frame, or engine, reducing the condition of the ship.
type ShipConditionEvent struct {

	// (required)
	Symbol ShipConditionEventSymbol `json:"symbol"`

	// (required)
	Component ShipConditionEventComponent `json:"component"`

	// The name of the event.
	//
	// (required)
	Name string `json:"name"`

	// A description of the event.
	//
	// (required)
	Description string `json:"description"`
}

type ShipConditionEventSymbol string

const (
	ShipConditionEventReactorOverload                  ShipConditionEventSymbol = "REACTOR_OVERLOAD"
	ShipConditionEventEnergySpikeFromMineral           ShipConditionEventSymbol = "ENERGY_SPIKE_FROM_MINERAL"
	ShipConditionEventSolarFlareInterference           ShipConditionEventSymbol = "SOLAR_FLARE_INTERFERENCE"
	ShipConditionEventCoolantLeak                      ShipConditionEventSymbol = "COOLANT_LEAK"
	ShipConditionEventPowerDistributionFluctuation     ShipConditionEventSymbol = "POWER_DISTRIBUTION_FLUCTUATION"
	ShipConditionEventMagneticFieldDisruption          ShipConditionEventSymbol = "MAGNETIC_FIELD_DISRUPTION"
	ShipConditionEventHullMicrometeoriteStrikes        ShipConditionEventSymbol = "HULL_MICROMETEORITE_STRIKES"
	ShipConditionEventStructuralStressFractures        ShipConditionEventSymbol = "STRUCTURAL_STRESS_FRACTURES"
	ShipConditionEventCorrosiveMineralContamination    ShipConditionEventSymbol = "CORROSIVE_MINERAL_CONTAMINATION"
	ShipConditionEventThermalExpansionMismatch         ShipConditionEventSymbol = "THERMAL_EXPANSION_MISMATCH"
	ShipConditionEventVibrationDamageFromDrilling      ShipConditionEventSymbol = "VIBRATION_DAMAGE_FROM_DRILLING"
	ShipConditionEventElectromagneticFieldInterference ShipConditionEventSymbol = "ELECTROMAGNETIC_FIELD_INTERFERENCE"
	ShipConditionEventImpactWithExtractedDebris        ShipConditionEventSymbol = "IMPACT_WITH_EXTRACTED_DEBRIS"
	ShipConditionEventFuelEfficiencyDegradation        ShipConditionEventSymbol = "FUEL_EFFICIENCY_DEGRADATION"
	ShipConditionEventCoolantSystemAgeing              ShipConditionEventSymbol = "COOLANT_SYSTEM_AGEING"
	ShipConditionEventDustMicroabrasions               ShipConditionEventSymbol = "DUST_MICROABRASIONS"
	ShipConditionEventThrusterNozzleWear               ShipConditionEventSymbol = "THRUSTER_NOZZLE_WEAR"
	ShipConditionEventExhaustPortClogging              ShipConditionEventSymbol = "EXHAUST_PORT_CLOGGING"
	ShipConditionEventBearingLubricationFade           ShipConditionEventSymbol = "BEARING_LUBRICATION_FADE"
	ShipConditionEventSensorCalibrationDrift           ShipConditionEventSymbol = "SENSOR_CALIBRATION_DRIFT"
	ShipConditionEventHullMicrometeoriteDamage         ShipConditionEventSymbol = "HULL_MICROMETEORITE_DAMAGE"
	ShipConditionEventSpaceDebrisCollision             ShipConditionEventSymbol = "SPACE_DEBRIS_COLLISION"
	ShipConditionEventThermalStress                    ShipConditionEventSymbol = "THERMAL_STRESS"
	ShipConditionEventVibrationOverload                ShipConditionEventSymbol = "VIBRATION_OVERLOAD"
	ShipConditionEventPressureDifferentialStress       ShipConditionEventSymbol = "PRESSURE_DIFFERENTIAL_STRESS"
	ShipConditionEventElectromagneticSurgeEffects      ShipConditionEventSymbol = "ELECTROMAGNETIC_SURGE_EFFECTS"
	ShipConditionEventAtmosphericEntryHeat             ShipConditionEventSymbol = "ATMOSPHERIC_ENTRY_HEAT"
)

type ShipConditionEventComponent string

const (
	ShipConditionEventComponentFrame   ShipConditionEventComponent = "FRAME"
	ShipConditionEventComponentReactor ShipConditionEventComponent = "REACTOR"
	ShipConditionEventComponentEngine  ShipConditionEventComponent = "ENGINE"
)
