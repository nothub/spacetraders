package st

import "time"

// A cooldown is a period of time in which a ship cannot perform certain actions.
type Cooldown struct {

	// The symbol of the ship that is on cooldown
	//
	//   >= 1 characters
	//
	// (required)
	ShipSymbol string `json:"shipSymbol"`

	// The total duration of the cooldown in seconds
	//
	//   >= 0
	//
	// (required)
	TotalSeconds int `json:"totalSeconds"`

	// The remaining duration of the cooldown in seconds
	//
	//   >= 0
	//
	// (required)
	RemainingSeconds int `json:"remainingSeconds"`

	// The date and time when the cooldown expires in ISO 8601 format
	Expiration time.Time `json:"expiration"`
}
