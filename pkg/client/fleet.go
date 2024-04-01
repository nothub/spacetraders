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

func PurchaseShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/403855e2e99ad-purchase-ship
	return nil
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

func OrbitShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/08777d60b6197-orbit-ship
	return nil
}

func ShipRefine() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/c42b57743a49f-ship-refine
	return nil
}

func CreateChart() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/177f127c7f888-create-chart
	return nil
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

func DockShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/a1061ae6545d5-dock-ship
	return nil
}

func CreateSurvey() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/6b7cb030c3b91-create-survey
	return nil
}

func ExtractResources() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/b3931d097608d-extract-resources
	return nil
}

func SiphonResources() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/f6c0d7877c43a-siphon-resources
	return nil
}

func ExtractResourcesWithSurvey() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/cdf110a7af0ea-extract-resources-with-survey
	return nil
}

func JettisonCargo() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/3b0f8b69f56ac-jettison-cargo
	return nil
}

func JumpShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/19f0dd2d633de-jump-ship
	return nil
}

func NavigateShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/c766b84253edc-navigate-ship
	return nil
}

func PatchShipNav() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/34a305032ec79-patch-ship-nav
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
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/faaf6603fc732-warp-ship
	return nil
}

func SellCargo() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/b8ed791381b41-sell-cargo
	return nil
}

func ScanSystems() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/d3358a9202901-scan-systems
	return nil
}

func ScanWaypoints() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/23dbc0fed17ec-scan-waypoints
	return nil
}

func ScanShips() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/74da68b7c32a7-scan-ships
	return nil
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

func NegotiateContract() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/1582bafa95003-negotiate-contract
	return nil
}

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

func InstallMount() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/266f3d0591399-install-mount
	return nil
}

func RemoveMount() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/9380132527c1d-remove-mount
	return nil
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

func ScrapShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/76039bfbf0cdb-scrap-ship
	return nil
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

func RepairShip() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/54a08ae25a5da-repair-ship
	return nil
}
