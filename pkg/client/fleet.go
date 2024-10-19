package st

import "strconv"

func ListShips(limit int, page int) (ships []Ship, meta Meta, err error) {
	var dto struct {
		Data []Ship `json:"data"`
		Meta Meta   `json:"meta"`
	}

	err = get(BaseUrl+"/my/ships", map[string]string{
		"limit": strconv.Itoa(limit),
		"page":  strconv.Itoa(page),
	}, &dto)
	if err != nil {
		return ships, meta, err
	}

	return dto.Data, dto.Meta, nil
}

// Purchase a ship from a Shipyard. In order to use this function, a ship under
// your agent's ownership must be in a waypoint that has the Shipyard trait,
// and the Shipyard must sell the type of the desired ship.
//
// Shipyards typically offer ship types, which are predefined templates of
// ships that have dedicated roles. A template comes with a preset of an
// engine, a reactor, and a frame. It may also include a few modules and
// mounts.
func PurchaseShip(shipType ShipType, waypointSymbol string) (agent Agent, ship Ship, transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Agent       Agent       `json:"agent"`
			Ship        Ship        `json:"ship"`
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships", nil, struct {
		ShipType       string `json:"shipType"`
		WaypointSymbol string `json:"waypointSymbol"`
	}{
		ShipType:       string(shipType),
		WaypointSymbol: waypointSymbol,
	}, &dto)
	if err != nil {
		return agent, ship, transaction, err
	}

	return dto.Data.Agent, dto.Data.Ship, dto.Data.Transaction, nil
}

func GetShip(shipSymbol string) (ship Ship, err error) {
	var dto struct {
		Data Ship `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol, nil, &dto)
	if err != nil {
		return ship, err
	}

	return dto.Data, nil
}

func GetShipCargo(shipSymbol string) (cargo ShipCargo, err error) {
	var dto struct {
		Data ShipCargo `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol+"/cargo", nil, &dto)
	if err != nil {
		return cargo, err
	}

	return dto.Data, nil
}

// Attempt to move your ship into orbit at its current location. The request
// will only succeed if your ship is capable of moving into orbit at the time
// of the request.
//
// Orbiting ships are able to do actions that require the ship to be above
// surface such as navigating or extracting, but cannot access elements in
// their current waypoint, such as the market or a shipyard.
//
// The endpoint is idempotent - successive calls will succeed even if the ship
// is already in orbit.
func OrbitShip(shipSymbol string) (nav ShipNav, err error) {
	var dto struct {
		Data struct {
			Nav ShipNav `json:"nav"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/orbit", nil, nil, &dto)
	if err != nil {
		return nav, err
	}

	return dto.Data.Nav, nil
}

// Attempt to refine the raw materials on your ship. The request will only
// succeed if your ship is capable of refining at the time of the request. In
// order to be able to refine, a ship must have goods that can be refined and
// have installed a Refinery module that can refine it.
//
// When refining, 30 basic goods will be converted into 10 processed goods.
func ShipRefine(shipSymbol string, produce RefineYield) (
	cargo ShipCargo,
	cooldown Cooldown,
	produced []TradeGoodStack,
	consumed []TradeGoodStack,
	err error,
) {
	var dto struct {
		Data struct {
			Cargo    ShipCargo        `json:"cargo"`
			Cooldown Cooldown         `json:"cooldown"`
			Produced []TradeGoodStack `json:"produced"`
			Consumed []TradeGoodStack `json:"consumed"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/refine", nil, struct {
		Produce string `json:"produce"`
	}{
		Produce: string(produce),
	}, &dto)
	if err != nil {
		return cargo, cooldown, produced, consumed, err
	}

	return dto.Data.Cargo, dto.Data.Cooldown, dto.Data.Produced, dto.Data.Consumed, nil
}

// Command a ship to chart the waypoint at its current location.
//
// Most waypoints in the universe are uncharted by default. These waypoints
// have their traits hidden until they have been charted by a ship.
//
// Charting a waypoint will record your agent as the one who created the chart,
// and all other agents would also be able to see the waypoint's traits.
func CreateChart(shipSymbol string) (chart Chart, waypoint Waypoint, err error) {
	var dto struct {
		Data struct {
			Chart    Chart    `json:"chart"`
			Waypoint Waypoint `json:"waypoint"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/orbit", nil, nil, &dto)
	if err != nil {
		return chart, waypoint, err
	}

	return dto.Data.Chart, dto.Data.Waypoint, nil
}

func GetShipCooldown(shipSymbol string) (cooldown Cooldown, err error) {
	var dto struct {
		Data Cooldown `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol+"/cooldown", nil, &dto)
	if err != nil {
		return cooldown, err
	}

	return dto.Data, nil
}

// Attempt to dock your ship at its current location. Docking will only succeed
// if your ship is capable of docking at the time of the request.
//
// Docked ships can access elements in their current location, such as the
// market or a shipyard, but cannot do actions that require the ship to be
// above surface such as navigating or extracting.
//
// The endpoint is idempotent - successive calls will succeed even if the ship
// is already docked.
func DockShip(shipSymbol string) (nav ShipNav, err error) {
	var dto struct {
		Data struct {
			Nav ShipNav `json:"nav"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/dock", nil, nil, &dto)
	if err != nil {
		return nav, err
	}

	return dto.Data.Nav, nil
}

// Create surveys on a waypoint that can be extracted such as asteroid fields.
// A survey focuses on specific types of deposits from the extracted location.
// When ships extract using this survey, they are guaranteed to procure a high
// amount of one of the goods in the survey.
//
// In order to use a survey, send the entire survey details in the body of the
// extract request.
//
// Each survey may have multiple deposits, and if a symbol shows up more than
// once, that indicates a higher chance of extracting that resource.
//
// Your ship will enter a cooldown after surveying in which it is unable to
// perform certain actions. Surveys will eventually expire after a period of
// time or will be exhausted after being extracted several times based on the
// survey's size. Multiple ships can use the same survey for extraction.
//
// A ship must have the Surveyor mount installed in order to use this function.
func CreateSurvey(shipSymbol string) (cooldown Cooldown, surveys []Survey, err error) {
	var dto struct {
		Data struct {
			Cooldown Cooldown `json:"cooldown"`
			Surveys  []Survey `json:"surveys"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/orbit", nil, nil, &dto)
	if err != nil {
		return cooldown, surveys, err
	}

	return dto.Data.Cooldown, dto.Data.Surveys, nil
}

func ExtractResources() (err error) {
	// TODO(#2): https://spacetraders.stoplight.io/docs/spacetraders/b3931d097608d-extract-resources
	return nil
}

// Siphon gases, such as hydrocarbon, from gas giants.
//
// The ship must be in orbit to be able to siphon and must have siphon mounts
// and a gas processor installed.
func SiphonResources(shipSymbol string) (cooldown Cooldown, siphon Siphon, cargo ShipCargo, events []ShipConditionEventSymbol, err error) {
	var dto struct {
		Data struct {
			Cooldown Cooldown                   `json:"cooldown"`
			Siphon   Siphon                     `json:"siphon"`
			Cargo    ShipCargo                  `json:"cargo"`
			Events   []ShipConditionEventSymbol `json:"events"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/siphon", nil, nil, &dto)
	if err != nil {
		return cooldown, siphon, cargo, events, err
	}

	return dto.Data.Cooldown, dto.Data.Siphon, dto.Data.Cargo, dto.Data.Events, nil
}

func ExtractResourcesWithSurvey() (err error) {
	// TODO(#3): https://spacetraders.stoplight.io/docs/spacetraders/cdf110a7af0ea-extract-resources-with-survey
	return nil
}

func JettisonCargo() (err error) {
	// TODO(#4): https://spacetraders.stoplight.io/docs/spacetraders/3b0f8b69f56ac-jettison-cargo
	return nil
}

func JumpShip() (err error) {
	// TODO(#5): https://spacetraders.stoplight.io/docs/spacetraders/19f0dd2d633de-jump-ship
	return nil
}

func NavigateShip() (err error) {
	// TODO(#6): https://spacetraders.stoplight.io/docs/spacetraders/c766b84253edc-navigate-ship
	return nil
}

func PatchShipNav() (err error) {
	// TODO(#7): https://spacetraders.stoplight.io/docs/spacetraders/34a305032ec79-patch-ship-nav
	return nil
}

func GetShipNav(shipSymbol string) (nav ShipNav, err error) {
	var dto struct {
		Data ShipNav `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol+"/nav", nil, &dto)
	if err != nil {
		return nav, err
	}

	return dto.Data, nil
}

func WarpShip() (err error) {
	// TODO(#8): https://spacetraders.stoplight.io/docs/spacetraders/faaf6603fc732-warp-ship
	return nil
}

func SellCargo() (err error) {
	// TODO(#9): https://spacetraders.stoplight.io/docs/spacetraders/b8ed791381b41-sell-cargo
	return nil
}

// Scan for nearby systems, retrieving information on the systems' distance
// from the ship and their waypoints. Requires a ship to have the "Sensor Array"
// mount installed to use.
//
// The ship will enter a cooldown after using this function, during which it
// cannot execute certain actions.
func ScanSystems(shipSymbol string) (cooldown Cooldown, systems []ScannedSystem, err error) {
	var dto struct {
		Data struct {
			Cooldown Cooldown        `json:"cooldown"`
			Systems  []ScannedSystem `json:"systems"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/scan/systems", nil, nil, &dto)
	if err != nil {
		return cooldown, systems, err
	}

	return dto.Data.Cooldown, dto.Data.Systems, nil
}

// Scan for nearby waypoints, retrieving detailed information on each waypoint
// in range. Scanning uncharted waypoints will allow you to ignore their
// uncharted state and will list the waypoints' traits.
//
// Requires a ship to have the "Sensor Array" mount installed to use.
//
// The ship will enter a cooldown after using this function, during which it
// cannot execute certain actions.
func ScanWaypoints(shipSymbol string) (cooldown Cooldown, waypoints []ScannedWaypoint, err error) {
	var dto struct {
		Data struct {
			Cooldown  Cooldown          `json:"cooldown"`
			Waypoints []ScannedWaypoint `json:"waypoints"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/scan/waypoints", nil, nil, &dto)
	if err != nil {
		return cooldown, waypoints, err
	}

	return dto.Data.Cooldown, dto.Data.Waypoints, nil
}

// Scan for nearby ships, retrieving information for all ships in range.
//
// Requires a ship to have the "Sensor Array" mount installed to use.
//
// The ship will enter a cooldown after using this function, during which it
// cannot execute certain actions.
func ScanShips(shipSymbol string) (cooldown Cooldown, ships []ScannedShip, err error) {
	var dto struct {
		Data struct {
			Cooldown Cooldown      `json:"cooldown"`
			Ships    []ScannedShip `json:"ships"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/scan/ships", nil, nil, &dto)
	if err != nil {
		return cooldown, ships, err
	}

	return dto.Data.Cooldown, dto.Data.Ships, nil
}

func RefuelShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/1bfb58c5239dd-refuel-ship
	return nil
}

func PurchaseCargo() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/45acbf7dc3005-purchase-cargo
	return nil
}

func TransferCargo() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/78b22e13e1ea1-transfer-cargo
	return nil
}

// Negotiate a new contract with the HQ.
//
// In order to negotiate a new contract, an agent must not have ongoing or
// offered contracts over the allowed maximum amount. Currently the maximum
// contracts an agent can have at a time is 1.
//
// Once a contract is negotiated, it is added to the list of contracts offered
// to the agent, which the agent can then accept.
//
// The ship must be present at any waypoint with a faction present to negotiate
// a contract with that faction.
func NegotiateContract(shipSymbol string) (contract Contract, err error) {
	var dto struct {
		Data struct {
			Contract Contract `json:"contract"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/negotiate/contract", nil, nil, &dto)
	if err != nil {
		return contract, err
	}

	return dto.Data.Contract, nil
}

// Get the mounts installed on a ship.
func GetMounts(shipSymbol string) (mounts []ShipMount, err error) {
	var dto struct {
		Data []ShipMount `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol+"/mounts", nil, &dto)
	if err != nil {
		return mounts, err
	}

	return dto.Data, nil
}

// Install a mount on a ship.
//
// In order to install a mount, the ship must be docked and located in a
// waypoint that has a Shipyard trait. The ship also must have the mount to
// install in its cargo hold.
//
// An installation fee will be deduced by the Shipyard for installing the mount
// on the ship.
func InstallMount(shipSymbol string, mountSymbol ShipMountSymbol) (agent Agent, mounts []ShipMount, cargo ShipCargo, transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Agent       Agent       `json:"agent"`
			Mounts      []ShipMount `json:"mounts"`
			Cargo       ShipCargo   `json:"cargo"`
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/mounts/install", nil, struct {
		Symbol string `json:"symbol"`
	}{
		Symbol: string(mountSymbol),
	}, &dto)
	if err != nil {
		return agent, mounts, cargo, transaction, err
	}

	return dto.Data.Agent, dto.Data.Mounts, dto.Data.Cargo, dto.Data.Transaction, nil
}

// Remove a mount from a ship.
//
// The ship must be docked in a waypoint that has the Shipyard trait, and must
// have the desired mount that it wish to remove installed.
//
// A removal fee will be deduced from the agent by the Shipyard.
func RemoveMount(shipSymbol string, mountSymbol ShipMountSymbol) (agent Agent, mounts []ShipMount, cargo ShipCargo, transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Agent       Agent       `json:"agent"`
			Mounts      []ShipMount `json:"mounts"`
			Cargo       ShipCargo   `json:"cargo"`
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/mounts/remove", nil, struct {
		Symbol string `json:"symbol"`
	}{
		Symbol: string(mountSymbol),
	}, &dto)
	if err != nil {
		return agent, mounts, cargo, transaction, err
	}

	return dto.Data.Agent, dto.Data.Mounts, dto.Data.Cargo, dto.Data.Transaction, nil
}

// Get the amount of value that will be returned when scrapping a ship.
func GetScrapShip(shipSymbol string) (transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol+"/scrap", nil, &dto)
	if err != nil {
		return transaction, err
	}

	return dto.Data.Transaction, nil
}

// Scrap a ship, removing it from the game and returning a portion of the
// ship's value to the agent. The ship must be docked in a waypoint that has
// the Shipyard trait in order to use this function. To preview the amount of
// value that will be returned, use the Get Ship action.
func ScrapShip(shipSymbol string) (agent Agent, transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Agent       Agent       `json:"agent"`
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/scrap", nil, nil, &dto)
	if err != nil {
		return agent, transaction, err
	}

	return dto.Data.Agent, dto.Data.Transaction, nil
}

// Get the cost of repairing a ship.
func GetRepairShip(shipSymbol string) (transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = get(BaseUrl+"/my/ships/"+shipSymbol+"/repair", nil, &dto)
	if err != nil {
		return transaction, err
	}

	return dto.Data.Transaction, nil
}

// Repair a ship, restoring the ship to maximum condition. The ship must be
// docked at a waypoint that has the Shipyard trait in order to use this
// function. To preview the cost of repairing the ship, use the Get action.
func RepairShip(shipSymbol string) (agent Agent, ship Ship, transaction Transaction, err error) {
	var dto struct {
		Data struct {
			Agent       Agent       `json:"agent"`
			Ship        Ship        `json:"ship"`
			Transaction Transaction `json:"transaction"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/ships/"+shipSymbol+"/repair", nil, nil, &dto)
	if err != nil {
		return agent, ship, transaction, err
	}

	return dto.Data.Agent, dto.Data.Ship, dto.Data.Transaction, nil
}
