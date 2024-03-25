package st

import "time"

// Details of the ship's fuel tanks including how much fuel was consumed during the last transit or action.
type ShipFuel struct {

	// The current amount of fuel in the ship's tanks.
	//
	//   >= 0
	//
	// (required)
	Current int `json:"current"`

	// The maximum amount of fuel the ship's tanks can hold.
	//
	//   >= 0
	//
	// (required)
	Capacity int `json:"capacity"`

	// An object that only shows up when an action has consumed fuel in the process.
	// Shows the fuel consumption data.
	Consumed struct {

		// The amount of fuel consumed by the most recent transit or action.
		//
		//   >= 0
		//
		// (required)
		Amount int `json:"amount"`

		// The time at which the fuel was consumed.
		//
		// (required)
		Timestamp time.Time `json:"timestamp"`
	} `json:"consumed"`
}
