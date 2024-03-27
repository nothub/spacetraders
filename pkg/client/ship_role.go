package st

// The registered role of the ship
type ShipRole string

const (
	ShipRoleFabricator  ShipRole = "FABRICATOR"
	ShipRoleHarvester   ShipRole = "HARVESTER"
	ShipRoleHauler      ShipRole = "HAULER"
	ShipRoleInterceptor ShipRole = "INTERCEPTOR"
	ShipRoleExcavator   ShipRole = "EXCAVATOR"
	ShipRoleTransport   ShipRole = "TRANSPORT"
	ShipRoleRepair      ShipRole = "REPAIR"
	ShipRoleSurveyor    ShipRole = "SURVEYOR"
	ShipRoleCommand     ShipRole = "COMMAND"
	ShipRoleCarrier     ShipRole = "CARRIER"
	ShipRolePatrol      ShipRole = "PATROL"
	ShipRoleSatellite   ShipRole = "SATELLITE"
	ShipRoleExplorer    ShipRole = "EXPLORER"
	ShipRoleRefinery    ShipRole = "REFINERY"
)
