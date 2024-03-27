package st

import (
	"time"
)

type Status struct {

	// The current status of the game server.
	Status string `json:"status"`

	// The current version of the API.
	Version string `json:"version"`

	// The date when the game server was last reset.
	ResetDate time.Time `json:"resetDate"`

	Description string `json:"description"`

	Stats struct {

		// Number of registered agents in the game.
		Agents int `json:"agents"`

		// Total number of ships in the game.
		Ships int `json:"ships"`

		// Total number of systems in the game.
		Systems int `json:"systems"`

		// Total number of waypoints in the game.
		Waypoints int `json:"waypoints"`
	} `json:"stats"`

	Leaderboards struct {

		// Top agents with the most credits.
		MostCredits []struct {

			// Symbol of the agent.
			AgentSymbol string `json:"agentSymbol"`

			// Amount of credits.
			//
			//   >= -9007199254740991 && <= 9007199254740991
			Credits int `json:"credits"`
		} `json:"mostCredits"`

		// Top agents with the most charted submitted.
		MostSubmittedCharts []struct {

			// Symbol of the agent.
			AgentSymbol string `json:"agentSymbol"`

			// Amount of charts done by the agent.
			ChartCount int `json:"chartCount"`
		} `json:"mostSubmittedCharts"`
	} `json:"leaderboards"`

	ServerResets struct {

		// The date and time when the game server will reset.
		Next time.Time `json:"next"`

		// How often we intend to reset the game server.
		Frequency string `json:"frequency"`
	} `json:"serverResets"`

	Announcements []struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"announcements"`

	Links []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"links"`
}

// GetStatus returns the status of the game server. This also includes a few
// global elements, such as announcements, server reset dates and leaderboards.
//
// (Authorization not required)
func GetStatus() (status Status, err error) {
	err = get(BaseUrl, nil, &status)
	if err != nil {
		return status, err
	}
	return status, nil
}
