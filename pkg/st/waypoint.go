package st

type Waypoint struct {

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

	// Relative position of the waypoint on the system's x axis.
	// This is not an absolute position in the universe.
	//
	// (required)
	X int `json:"x"`

	// Relative position of the waypoint on the system's y axis.
	// This is not an absolute position in the universe.
	//
	// (required)
	Y int `json:"y"`

	// Waypoints that orbit this waypoint.
	//
	// (required)
	Orbitals []struct {

		// The symbol of the orbiting waypoint.
		//
		//   >= 1 characters
		//
		// (required)
		Symbol string `json:"symbol"`
	} `json:"orbitals"`

	// The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint.
	// Otherwise this value is undefined.
	//
	//   >= 1 characters
	Orbits string `json:"orbits"`

	// The faction that controls the waypoint.
	Faction struct {

		// The symbol of the faction.
		//
		// (required)
		Symbol FactionSymbol `json:"symbol"`
	} `json:"faction"`

	// The traits of the waypoint.
	//
	// (required)
	Traits []WaypointTrait `json:"traits"`

	// The modifiers of the waypoint.
	Modifiers []struct {

		// The unique identifier of the modifier.
		//
		// (required)
		Symbol WaypointModifierSymbol `json:"symbol"`

		// The name of the trait.
		//
		// (required)
		Name string `json:"name"`

		// A description of the trait.
		//
		// (required)
		Description string `json:"description"`
	} `json:"modifiers"`

	// The chart of a system or waypoint, which makes the location visible to other agents.
	Chart Chart `json:"chart"`

	// True if the waypoint is under construction.
	//
	// (required)
	IsUnderConstruction bool `json:"isUnderConstruction"`
}

// A waypoint is a location that ships can travel to such as a Planet, Moon or Space Station.
type WaypointTrait struct {

	// The unique identifier of the trait.
	//
	// (required)
	Symbol WaypointTraitSymbol `json:"symbol"`

	// The name of the trait.
	//
	// (required)
	Name string `json:"name"`

	// A description of the trait.
	//
	// (required)
	Description string `json:"description"`
}

type WaypointModifierSymbol string

const (
	WaypointModifierStripped      WaypointModifierSymbol = "STRIPPED"
	WaypointModifierUnstable      WaypointModifierSymbol = "UNSTABLE"
	WaypointModifierRadiationLeak WaypointModifierSymbol = "RADIATION_LEAK"
	WaypointModifierCriticalLimit WaypointModifierSymbol = "CRITICAL_LIMIT"
	WaypointModifierCivilUnrest   WaypointModifierSymbol = "CIVIL_UNREST"
)

type WaypointTraitSymbol string

const (
	WaypointTraitUncharted             WaypointTraitSymbol = "UNCHARTED"
	WaypointTraitUnderConstruction     WaypointTraitSymbol = "UNDER_CONSTRUCTION"
	WaypointTraitMarketplace           WaypointTraitSymbol = "MARKETPLACE"
	WaypointTraitShipyard              WaypointTraitSymbol = "SHIPYARD"
	WaypointTraitOutpost               WaypointTraitSymbol = "OUTPOST"
	WaypointTraitScatteredSettlements  WaypointTraitSymbol = "SCATTERED_SETTLEMENTS"
	WaypointTraitSprawlingCities       WaypointTraitSymbol = "SPRAWLING_CITIES"
	WaypointTraitMegaStructures        WaypointTraitSymbol = "MEGA_STRUCTURES"
	WaypointTraitPirateBase            WaypointTraitSymbol = "PIRATE_BASE"
	WaypointTraitOvercrowded           WaypointTraitSymbol = "OVERCROWDED"
	WaypointTraitHighTech              WaypointTraitSymbol = "HIGH_TECH"
	WaypointTraitCorrupt               WaypointTraitSymbol = "CORRUPT"
	WaypointTraitBureaucratic          WaypointTraitSymbol = "BUREAUCRATIC"
	WaypointTraitTradingHub            WaypointTraitSymbol = "TRADING_HUB"
	WaypointTraitIndustrial            WaypointTraitSymbol = "INDUSTRIAL"
	WaypointTraitBlackMarket           WaypointTraitSymbol = "BLACK_MARKET"
	WaypointTraitResearchFacility      WaypointTraitSymbol = "RESEARCH_FACILITY"
	WaypointTraitMilitaryBase          WaypointTraitSymbol = "MILITARY_BASE"
	WaypointTraitSurveillanceOutpost   WaypointTraitSymbol = "SURVEILLANCE_OUTPOST"
	WaypointTraitExplorationOutpost    WaypointTraitSymbol = "EXPLORATION_OUTPOST"
	WaypointTraitMineralDeposits       WaypointTraitSymbol = "MINERAL_DEPOSITS"
	WaypointTraitCommonMetalDeposits   WaypointTraitSymbol = "COMMON_METAL_DEPOSITS"
	WaypointTraitPreciousMetalDeposits WaypointTraitSymbol = "PRECIOUS_METAL_DEPOSITS"
	WaypointTraitRareMetalDeposits     WaypointTraitSymbol = "RARE_METAL_DEPOSITS"
	WaypointTraitMethanePools          WaypointTraitSymbol = "METHANE_POOLS"
	WaypointTraitIceCrystals           WaypointTraitSymbol = "ICE_CRYSTALS"
	WaypointTraitExplosiveGases        WaypointTraitSymbol = "EXPLOSIVE_GASES"
	WaypointTraitStrongMagnetosphere   WaypointTraitSymbol = "STRONG_MAGNETOSPHERE"
	WaypointTraitVibrantAuroras        WaypointTraitSymbol = "VIBRANT_AURORAS"
	WaypointTraitSaltFlats             WaypointTraitSymbol = "SALT_FLATS"
	WaypointTraitCanyons               WaypointTraitSymbol = "CANYONS"
	WaypointTraitPerpetualDaylight     WaypointTraitSymbol = "PERPETUAL_DAYLIGHT"
	WaypointTraitPerpetualOvercast     WaypointTraitSymbol = "PERPETUAL_OVERCAST"
	WaypointTraitDrySeabeds            WaypointTraitSymbol = "DRY_SEABEDS"
	WaypointTraitMagmaSeas             WaypointTraitSymbol = "MAGMA_SEAS"
	WaypointTraitSupervolcanoes        WaypointTraitSymbol = "SUPERVOLCANOES"
	WaypointTraitAshClouds             WaypointTraitSymbol = "ASH_CLOUDS"
	WaypointTraitVastRuins             WaypointTraitSymbol = "VAST_RUINS"
	WaypointTraitMutatedFlora          WaypointTraitSymbol = "MUTATED_FLORA"
	WaypointTraitTerraformed           WaypointTraitSymbol = "TERRAFORMED"
	WaypointTraitExtremeTemperatures   WaypointTraitSymbol = "EXTREME_TEMPERATURES"
	WaypointTraitExtremePressure       WaypointTraitSymbol = "EXTREME_PRESSURE"
	WaypointTraitDiverseLife           WaypointTraitSymbol = "DIVERSE_LIFE"
	WaypointTraitScarceLife            WaypointTraitSymbol = "SCARCE_LIFE"
	WaypointTraitFossils               WaypointTraitSymbol = "FOSSILS"
	WaypointTraitWeakGravity           WaypointTraitSymbol = "WEAK_GRAVITY"
	WaypointTraitStrongGravity         WaypointTraitSymbol = "STRONG_GRAVITY"
	WaypointTraitCrushingGravity       WaypointTraitSymbol = "CRUSHING_GRAVITY"
	WaypointTraitToxicAtmosphere       WaypointTraitSymbol = "TOXIC_ATMOSPHERE"
	WaypointTraitCorrosiveAtmosphere   WaypointTraitSymbol = "CORROSIVE_ATMOSPHERE"
	WaypointTraitBreathableAtmosphere  WaypointTraitSymbol = "BREATHABLE_ATMOSPHERE"
	WaypointTraitThinAtmosphere        WaypointTraitSymbol = "THIN_ATMOSPHERE"
	WaypointTraitJovian                WaypointTraitSymbol = "JOVIAN"
	WaypointTraitRocky                 WaypointTraitSymbol = "ROCKY"
	WaypointTraitVolcanic              WaypointTraitSymbol = "VOLCANIC"
	WaypointTraitFrozen                WaypointTraitSymbol = "FROZEN"
	WaypointTraitSwamp                 WaypointTraitSymbol = "SWAMP"
	WaypointTraitBarren                WaypointTraitSymbol = "BARREN"
	WaypointTraitTemperate             WaypointTraitSymbol = "TEMPERATE"
	WaypointTraitJungle                WaypointTraitSymbol = "JUNGLE"
	WaypointTraitOcean                 WaypointTraitSymbol = "OCEAN"
	WaypointTraitRadioactive           WaypointTraitSymbol = "RADIOACTIVE"
	WaypointTraitMicroGravityAnomalies WaypointTraitSymbol = "MICRO_GRAVITY_ANOMALIES"
	WaypointTraitDebrisCluster         WaypointTraitSymbol = "DEBRIS_CLUSTER"
	WaypointTraitDeepCraters           WaypointTraitSymbol = "DEEP_CRATERS"
	WaypointTraitShallowCraters        WaypointTraitSymbol = "SHALLOW_CRATERS"
	WaypointTraitUnstableComposition   WaypointTraitSymbol = "UNSTABLE_COMPOSITION"
	WaypointTraitHollowedInterior      WaypointTraitSymbol = "HOLLOWED_INTERIOR"
	WaypointTraitStripped              WaypointTraitSymbol = "STRIPPED"
)

// The type of waypoint.
type WaypointType string

const (
	WaypointTypePlanet                WaypointType = "PLANET"
	WaypointTypeGasGiant              WaypointType = "GAS_GIANT"
	WaypointTypeMoon                  WaypointType = "MOON"
	WaypointTypeOrbitalStation        WaypointType = "ORBITAL_STATION"
	WaypointTypeJumpGate              WaypointType = "JUMP_GATE"
	WaypointTypeAsteroidField         WaypointType = "ASTEROID_FIELD"
	WaypointTypeAsteroid              WaypointType = "ASTEROID"
	WaypointTypeEngineeredAsteroid    WaypointType = "ENGINEERED_ASTEROID"
	WaypointTypeAsteroidBase          WaypointType = "ASTEROID_BASE"
	WaypointTypeNebula                WaypointType = "NEBULA"
	WaypointTypeDebrisField           WaypointType = "DEBRIS_FIELD"
	WaypointTypeGravityWell           WaypointType = "GRAVITY_WELL"
	WaypointTypeArtificialGravityWell WaypointType = "ARTIFICIAL_GRAVITY_WELL"
	WaypointTypeFuelStation           WaypointType = "FUEL_STATION"
)
