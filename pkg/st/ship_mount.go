package st

// A mount is installed on the exterier of a ship.
type ShipMount struct {

	// Symbo of this mount.
	//
	// (required)
	Symbol ShipMountSymbol `json:"symbol"`

	// Name of this mount.
	//
	// (required)
	Name string `json:"name"`

	// Description of this mount.
	Description string `json:"description"`

	// Mounts that have this value, such as mining lasers,
	// denote how powerful this mount's capabilities are.
	//
	//   >= 0
	Strength int `json:"strength"`

	// Mounts that have this value denote what goods can be produced from using the mount.
	Deposits []ShipMountYield `json:"deposits"`

	// The requirements for installation on a ship
	//
	// (required)
	Requirements ShipRequirements `json:"requirements"`
}

type ShipMountSymbol string

const (
	ShipMountMountGasSiphon1       ShipMountSymbol = "MOUNT_GAS_SIPHON_I"
	ShipMountMountGasSiphon2       ShipMountSymbol = "MOUNT_GAS_SIPHON_II"
	ShipMountMountGasSiphon3       ShipMountSymbol = "MOUNT_GAS_SIPHON_III"
	ShipMountMountSurveyor1        ShipMountSymbol = "MOUNT_SURVEYOR_I"
	ShipMountMountSurveyor2        ShipMountSymbol = "MOUNT_SURVEYOR_II"
	ShipMountMountSurveyor3        ShipMountSymbol = "MOUNT_SURVEYOR_III"
	ShipMountMountSensorArray1     ShipMountSymbol = "MOUNT_SENSOR_ARRAY_I"
	ShipMountMountSensorArray2     ShipMountSymbol = "MOUNT_SENSOR_ARRAY_II"
	ShipMountMountSensorArray3     ShipMountSymbol = "MOUNT_SENSOR_ARRAY_III"
	ShipMountMountMiningLaser1     ShipMountSymbol = "MOUNT_MINING_LASER_I"
	ShipMountMountMiningLaser2     ShipMountSymbol = "MOUNT_MINING_LASER_II"
	ShipMountMountMiningLaser3     ShipMountSymbol = "MOUNT_MINING_LASER_III"
	ShipMountMountLaserCannon1     ShipMountSymbol = "MOUNT_LASER_CANNON_I"
	ShipMountMountMissileLauncher1 ShipMountSymbol = "MOUNT_MISSILE_LAUNCHER_I"
	ShipMountMountTurret1          ShipMountSymbol = "MOUNT_TURRET_I"
)

type ShipMountYield string

const (
	ShipMountYieldQuartzSand      = ShipMountYield(TradeGoodQuartzSand)
	ShipMountYieldSiliconCrystals = ShipMountYield(TradeGoodSiliconCrystals)
	ShipMountYieldPreciousStones  = ShipMountYield(TradeGoodPreciousStones)
	ShipMountYieldIceWater        = ShipMountYield(TradeGoodIceWater)
	ShipMountYieldAmmoniaIce      = ShipMountYield(TradeGoodAmmoniaIce)
	ShipMountYieldIronOre         = ShipMountYield(TradeGoodIronOre)
	ShipMountYieldCopperOre       = ShipMountYield(TradeGoodCopperOre)
	ShipMountYieldSilverOre       = ShipMountYield(TradeGoodSilverOre)
	ShipMountYieldAluminumOre     = ShipMountYield(TradeGoodAluminumOre)
	ShipMountYieldGoldOre         = ShipMountYield(TradeGoodGoldOre)
	ShipMountYieldPlatinumOre     = ShipMountYield(TradeGoodPlatinumOre)
	ShipMountYieldDiamonds        = ShipMountYield(TradeGoodDiamonds)
	ShipMountYieldUraniteOre      = ShipMountYield(TradeGoodUraniteOre)
	ShipMountYieldMeritiumOre     = ShipMountYield(TradeGoodMeritiumOre)
)
