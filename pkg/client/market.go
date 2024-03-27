package st

import "time"

type Market struct {

	// The symbol of the market. The symbol is the same as the waypoint where the market is located.
	//
	// (required)
	Symbol string `json:"symbol"`

	// The list of goods that are exported from this market.
	//
	// (required)
	Exports []TradeGood `json:"exports"`

	// The list of goods that are sought as imports in this market.
	//
	// (required)
	Imports []TradeGood `json:"imports"`

	// The list of goods that are bought and sold between agents at this market.
	//
	// (required)
	Exchange []TradeGood `json:"exchange"`

	// The list of recent transactions at this market. Visible only when a ship is present at the market.
	Transactions []MarketTransaction `json:"transactions"`

	// The list of goods that are traded at this market. Visible only when a ship is present at the market.
	TradeGoods []MarketTradeGood `json:"tradeGoods"`
}

// A good that can be traded for other goods or currency.
type TradeGood struct {

	// The good's symbol.
	//
	// (required)
	Symbol TradeGoodSymbol `json:"symbol"`

	// The name of the good.
	//
	// (required)
	Name string `json:"name"`

	// The description of the good.
	//
	// (required)
	Description string `json:"description"`
}

// The good's symbol.
type TradeGoodSymbol string

const (
	TradeGoodPreciousStones          TradeGoodSymbol = "PRECIOUS_STONES"
	TradeGoodQuartzSand              TradeGoodSymbol = "QUARTZ_SAND"
	TradeGoodSiliconCrystals         TradeGoodSymbol = "SILICON_CRYSTALS"
	TradeGoodAmmoniaIce              TradeGoodSymbol = "AMMONIA_ICE"
	TradeGoodLiquidHydrogen          TradeGoodSymbol = "LIQUID_HYDROGEN"
	TradeGoodLiquidNitrogen          TradeGoodSymbol = "LIQUID_NITROGEN"
	TradeGoodIceWater                TradeGoodSymbol = "ICE_WATER"
	TradeGoodExoticMatter            TradeGoodSymbol = "EXOTIC_MATTER"
	TradeGoodAdvancedCircuitry       TradeGoodSymbol = "ADVANCED_CIRCUITRY"
	TradeGoodGravitonEmitters        TradeGoodSymbol = "GRAVITON_EMITTERS"
	TradeGoodIron                    TradeGoodSymbol = "IRON"
	TradeGoodIronOre                 TradeGoodSymbol = "IRON_ORE"
	TradeGoodCopper                  TradeGoodSymbol = "COPPER"
	TradeGoodCopperOre               TradeGoodSymbol = "COPPER_ORE"
	TradeGoodAluminum                TradeGoodSymbol = "ALUMINUM"
	TradeGoodAluminumOre             TradeGoodSymbol = "ALUMINUM_ORE"
	TradeGoodSilver                  TradeGoodSymbol = "SILVER"
	TradeGoodSilverOre               TradeGoodSymbol = "SILVER_ORE"
	TradeGoodGold                    TradeGoodSymbol = "GOLD"
	TradeGoodGoldOre                 TradeGoodSymbol = "GOLD_ORE"
	TradeGoodPlatinum                TradeGoodSymbol = "PLATINUM"
	TradeGoodPlatinumOre             TradeGoodSymbol = "PLATINUM_ORE"
	TradeGoodDiamonds                TradeGoodSymbol = "DIAMONDS"
	TradeGoodUranite                 TradeGoodSymbol = "URANITE"
	TradeGoodUraniteOre              TradeGoodSymbol = "URANITE_ORE"
	TradeGoodMeritium                TradeGoodSymbol = "MERITIUM"
	TradeGoodMeritiumOre             TradeGoodSymbol = "MERITIUM_ORE"
	TradeGoodHydrocarbon             TradeGoodSymbol = "HYDROCARBON"
	TradeGoodAntimatter              TradeGoodSymbol = "ANTIMATTER"
	TradeGoodFabMats                 TradeGoodSymbol = "FAB_MATS"
	TradeGoodFertilizers             TradeGoodSymbol = "FERTILIZERS"
	TradeGoodFabrics                 TradeGoodSymbol = "FABRICS"
	TradeGoodFood                    TradeGoodSymbol = "FOOD"
	TradeGoodJewelry                 TradeGoodSymbol = "JEWELRY"
	TradeGoodMachinery               TradeGoodSymbol = "MACHINERY"
	TradeGoodFirearms                TradeGoodSymbol = "FIREARMS"
	TradeGoodAssaultRifles           TradeGoodSymbol = "ASSAULT_RIFLES"
	TradeGoodMilitaryEquipment       TradeGoodSymbol = "MILITARY_EQUIPMENT"
	TradeGoodExplosives              TradeGoodSymbol = "EXPLOSIVES"
	TradeGoodLabInstruments          TradeGoodSymbol = "LAB_INSTRUMENTS"
	TradeGoodAmmunition              TradeGoodSymbol = "AMMUNITION"
	TradeGoodElectronics             TradeGoodSymbol = "ELECTRONICS"
	TradeGoodShipPlating             TradeGoodSymbol = "SHIP_PLATING"
	TradeGoodShipParts               TradeGoodSymbol = "SHIP_PARTS"
	TradeGoodEquipment               TradeGoodSymbol = "EQUIPMENT"
	TradeGoodFuel                    TradeGoodSymbol = "FUEL"
	TradeGoodMedicine                TradeGoodSymbol = "MEDICINE"
	TradeGoodDrugs                   TradeGoodSymbol = "DRUGS"
	TradeGoodClothing                TradeGoodSymbol = "CLOTHING"
	TradeGoodMicroprocessors         TradeGoodSymbol = "MICROPROCESSORS"
	TradeGoodPlastics                TradeGoodSymbol = "PLASTICS"
	TradeGoodPolynucleotides         TradeGoodSymbol = "POLYNUCLEOTIDES"
	TradeGoodBiocomposites           TradeGoodSymbol = "BIOCOMPOSITES"
	TradeGoodQuantumStabilizers      TradeGoodSymbol = "QUANTUM_STABILIZERS"
	TradeGoodNanobots                TradeGoodSymbol = "NANOBOTS"
	TradeGoodAiMainframes            TradeGoodSymbol = "AI_MAINFRAMES"
	TradeGoodQuantumDrives           TradeGoodSymbol = "QUANTUM_DRIVES"
	TradeGoodRoboticDrones           TradeGoodSymbol = "ROBOTIC_DRONES"
	TradeGoodCyberImplants           TradeGoodSymbol = "CYBER_IMPLANTS"
	TradeGoodGeneTherapeutics        TradeGoodSymbol = "GENE_THERAPEUTICS"
	TradeGoodNeuralChips             TradeGoodSymbol = "NEURAL_CHIPS"
	TradeGoodMoodRegulators          TradeGoodSymbol = "MOOD_REGULATORS"
	TradeGoodViralAgents             TradeGoodSymbol = "VIRAL_AGENTS"
	TradeGoodMicroFusionGenerators   TradeGoodSymbol = "MICRO_FUSION_GENERATORS"
	TradeGoodSupergrains             TradeGoodSymbol = "SUPERGRAINS"
	TradeGoodLaserRifles             TradeGoodSymbol = "LASER_RIFLES"
	TradeGoodHolographics            TradeGoodSymbol = "HOLOGRAPHICS"
	TradeGoodShipSalvage             TradeGoodSymbol = "SHIP_SALVAGE"
	TradeGoodRelicTech               TradeGoodSymbol = "RELIC_TECH"
	TradeGoodNovelLifeforms          TradeGoodSymbol = "NOVEL_LIFEFORMS"
	TradeGoodBotanicalSpecimens      TradeGoodSymbol = "BOTANICAL_SPECIMENS"
	TradeGoodCulturalArtifacts       TradeGoodSymbol = "CULTURAL_ARTIFACTS"
	TradeGoodFrameProbe              TradeGoodSymbol = "FRAME_PROBE"
	TradeGoodFrameDrone              TradeGoodSymbol = "FRAME_DRONE"
	TradeGoodFrameInterceptor        TradeGoodSymbol = "FRAME_INTERCEPTOR"
	TradeGoodFrameRacer              TradeGoodSymbol = "FRAME_RACER"
	TradeGoodFrameFighter            TradeGoodSymbol = "FRAME_FIGHTER"
	TradeGoodFrameFrigate            TradeGoodSymbol = "FRAME_FRIGATE"
	TradeGoodFrameShuttle            TradeGoodSymbol = "FRAME_SHUTTLE"
	TradeGoodFrameExplorer           TradeGoodSymbol = "FRAME_EXPLORER"
	TradeGoodFrameMiner              TradeGoodSymbol = "FRAME_MINER"
	TradeGoodFrameLightFreighter     TradeGoodSymbol = "FRAME_LIGHT_FREIGHTER"
	TradeGoodFrameHeavyFreighter     TradeGoodSymbol = "FRAME_HEAVY_FREIGHTER"
	TradeGoodFrameTransport          TradeGoodSymbol = "FRAME_TRANSPORT"
	TradeGoodFrameDestroyer          TradeGoodSymbol = "FRAME_DESTROYER"
	TradeGoodFrameCruiser            TradeGoodSymbol = "FRAME_CRUISER"
	TradeGoodFrameCarrier            TradeGoodSymbol = "FRAME_CARRIER"
	TradeGoodReactorSolar1           TradeGoodSymbol = "REACTOR_SOLAR_I"
	TradeGoodReactorFusion1          TradeGoodSymbol = "REACTOR_FUSION_I"
	TradeGoodReactorFission1         TradeGoodSymbol = "REACTOR_FISSION_I"
	TradeGoodReactorChemical1        TradeGoodSymbol = "REACTOR_CHEMICAL_I"
	TradeGoodReactorAntimatter1      TradeGoodSymbol = "REACTOR_ANTIMATTER_I"
	TradeGoodEngineImpulseDrive1     TradeGoodSymbol = "ENGINE_IMPULSE_DRIVE_I"
	TradeGoodEngineIonDrive1         TradeGoodSymbol = "ENGINE_ION_DRIVE_I"
	TradeGoodEngineIonDrive2         TradeGoodSymbol = "ENGINE_ION_DRIVE_II"
	TradeGoodEngineHyperDrive1       TradeGoodSymbol = "ENGINE_HYPER_DRIVE_I"
	TradeGoodModuleMineralProcessor1 TradeGoodSymbol = "MODULE_MINERAL_PROCESSOR_I"
	TradeGoodModuleGasProcessor1     TradeGoodSymbol = "MODULE_GAS_PROCESSOR_I"
	TradeGoodModuleCargoHold1        TradeGoodSymbol = "MODULE_CARGO_HOLD_I"
	TradeGoodModuleCargoHold2        TradeGoodSymbol = "MODULE_CARGO_HOLD_II"
	TradeGoodModuleCargoHold3        TradeGoodSymbol = "MODULE_CARGO_HOLD_III"
	TradeGoodModuleCrewQuarters1     TradeGoodSymbol = "MODULE_CREW_QUARTERS_I"
	TradeGoodModuleEnvoyQuarters1    TradeGoodSymbol = "MODULE_ENVOY_QUARTERS_I"
	TradeGoodModulePassengerCabin1   TradeGoodSymbol = "MODULE_PASSENGER_CABIN_I"
	TradeGoodModuleMicroRefinery1    TradeGoodSymbol = "MODULE_MICRO_REFINERY_I"
	TradeGoodModuleScienceLab1       TradeGoodSymbol = "MODULE_SCIENCE_LAB_I"
	TradeGoodModuleJumpDrive1        TradeGoodSymbol = "MODULE_JUMP_DRIVE_I"
	TradeGoodModuleJumpDrive2        TradeGoodSymbol = "MODULE_JUMP_DRIVE_II"
	TradeGoodModuleJumpDrive3        TradeGoodSymbol = "MODULE_JUMP_DRIVE_III"
	TradeGoodModuleWarpDrive1        TradeGoodSymbol = "MODULE_WARP_DRIVE_I"
	TradeGoodModuleWarpDrive2        TradeGoodSymbol = "MODULE_WARP_DRIVE_II"
	TradeGoodModuleWarpDrive3        TradeGoodSymbol = "MODULE_WARP_DRIVE_III"
	TradeGoodModuleShieldGenerator1  TradeGoodSymbol = "MODULE_SHIELD_GENERATOR_I"
	TradeGoodModuleShieldGenerator2  TradeGoodSymbol = "MODULE_SHIELD_GENERATOR_II"
	TradeGoodModuleOreRefinery1      TradeGoodSymbol = "MODULE_ORE_REFINERY_I"
	TradeGoodModuleFuelRefinery1     TradeGoodSymbol = "MODULE_FUEL_REFINERY_I"
	TradeGoodMountGasSiphon1         TradeGoodSymbol = "MOUNT_GAS_SIPHON_I"
	TradeGoodMountGasSiphon2         TradeGoodSymbol = "MOUNT_GAS_SIPHON_II"
	TradeGoodMountGasSiphon3         TradeGoodSymbol = "MOUNT_GAS_SIPHON_III"
	TradeGoodMountSurveyor1          TradeGoodSymbol = "MOUNT_SURVEYOR_I"
	TradeGoodMountSurveyor2          TradeGoodSymbol = "MOUNT_SURVEYOR_II"
	TradeGoodMountSurveyor3          TradeGoodSymbol = "MOUNT_SURVEYOR_III"
	TradeGoodMountSensorArray1       TradeGoodSymbol = "MOUNT_SENSOR_ARRAY_I"
	TradeGoodMountSensorArray2       TradeGoodSymbol = "MOUNT_SENSOR_ARRAY_II"
	TradeGoodMountSensorArray3       TradeGoodSymbol = "MOUNT_SENSOR_ARRAY_III"
	TradeGoodMountMiningLaser1       TradeGoodSymbol = "MOUNT_MINING_LASER_I"
	TradeGoodMountMiningLaser2       TradeGoodSymbol = "MOUNT_MINING_LASER_II"
	TradeGoodMountMiningLaser3       TradeGoodSymbol = "MOUNT_MINING_LASER_III"
	TradeGoodMountLaserCannon1       TradeGoodSymbol = "MOUNT_LASER_CANNON_I"
	TradeGoodMountMissileLauncher1   TradeGoodSymbol = "MOUNT_MISSILE_LAUNCHER_I"
	TradeGoodMountTurret1            TradeGoodSymbol = "MOUNT_TURRET_I"
	TradeGoodShipProbe               TradeGoodSymbol = "SHIP_PROBE"
	TradeGoodShipMiningDrone         TradeGoodSymbol = "SHIP_MINING_DRONE"
	TradeGoodShipSiphonDrone         TradeGoodSymbol = "SHIP_SIPHON_DRONE"
	TradeGoodShipInterceptor         TradeGoodSymbol = "SHIP_INTERCEPTOR"
	TradeGoodShipLightHauler         TradeGoodSymbol = "SHIP_LIGHT_HAULER"
	TradeGoodShipCommandFrigate      TradeGoodSymbol = "SHIP_COMMAND_FRIGATE"
	TradeGoodShipExplorer            TradeGoodSymbol = "SHIP_EXPLORER"
	TradeGoodShipHeavyFreighter      TradeGoodSymbol = "SHIP_HEAVY_FREIGHTER"
	TradeGoodShipLightShuttle        TradeGoodSymbol = "SHIP_LIGHT_SHUTTLE"
	TradeGoodShipOreHound            TradeGoodSymbol = "SHIP_ORE_HOUND"
	TradeGoodShipRefiningFreighter   TradeGoodSymbol = "SHIP_REFINING_FREIGHTER"
	TradeGoodShipSurveyor            TradeGoodSymbol = "SHIP_SURVEYOR"
)

// Result of a transaction with a market.
type MarketTransaction struct {

	// The symbol of the waypoint.
	//
	//   >= 1 characters
	//
	// (required)
	WaypointSymbol string `json:"waypointSymbol"`

	// The symbol of the ship that made the transaction.
	//
	// (required)
	ShipSymbol string `json:"shipSymbol"`

	// The symbol of the trade good.
	//
	// (required)
	TradeSymbol TradeGoodSymbol `json:"tradeSymbol"`

	// The type of transaction.
	//
	// (required)
	Type TransactionType `json:"type"`

	// The number of units of the transaction.
	//
	//   >= 0
	//
	// (required)
	Units int `json:"units"`

	// The price per unit of the transaction.
	//
	//   >= 0
	//
	// (required)
	PricePerUnit int `json:"pricePerUnit"`

	// The total price of the transaction.
	//
	//   >= 0
	//
	// (required)
	TotalPrice int `json:"totalPrice"`

	// The timestamp of the transaction.
	//
	// (required)
	Timestamp time.Time `json:"timestamp"`
}

type TransactionType string

const (
	TransactionTypePurchase TransactionType = "PURCHASE"
	TransactionTypeSell     TransactionType = "SELL"
)

type MarketTradeGood struct {

	// The good's symbol.
	//
	// (required)
	Symbol TradeGoodSymbol `json:"symbol"`

	// The type of trade good (export, import, or exchange).
	//
	// (required)
	Type TradeGoodType `json:"type"`

	// This is the maximum number of units that can be purchased
	// or sold at this market in a single trade for this good.
	// Trade volume also gives an indication of price volatility.
	// A market with a low trade volume will have large price swings,
	// while high trade volume will be more resilient to price changes.
	//
	//   >= 1
	//
	// (required)
	TradeVolume int `json:"tradeVolume"`

	// The supply level of a trade good.
	//
	// (required)
	Supply SupplyLevel `json:"supply"`

	// The activity level of a trade good.
	// If the good is an import, this represents how strong consumption is.
	// If the good is an export, this represents how strong the production is for the good.
	// When activity is strong, consumption or production is near maximum capacity.
	// When activity is weak, consumption or production is near minimum capacity.
	Activity ActivityLevel `json:"activity"`

	// The price at which this good can be purchased from the market.
	//
	//   >= 0
	//
	// (required)
	PurchasePrice int `json:"purchasePrice"`

	// The price at which this good can be sold to the market.
	//
	//   >= 0
	//
	// (required)
	SellPrice int `json:"sellPrice"`
}

type TradeGoodType string

const (
	TradeGoodTypeExport   TradeGoodType = "EXPORT"
	TradeGoodTypeImport   TradeGoodType = "IMPORT"
	TradeGoodTypeExchange TradeGoodType = "EXCHANGE"
)

// The supply level of a trade good.
type SupplyLevel string

const (
	SupplyLevelScarce   SupplyLevel = "SCARCE"
	SupplyLevelLimited  SupplyLevel = "LIMITED"
	SupplyLevelModerate SupplyLevel = "MODERATE"
	SupplyLevelHigh     SupplyLevel = "HIGH"
	SupplyLevelAbundant SupplyLevel = "ABUNDANT"
)

type ActivityLevel string

const (
	ActivityLevelWeak       ActivityLevel = "WEAK"
	ActivityLevelGrowing    ActivityLevel = "GROWING"
	ActivityLevelStrong     ActivityLevel = "STRONG"
	ActivityLevelRestricted ActivityLevel = "RESTRICTED"
)
