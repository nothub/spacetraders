package st

import (
	"time"
)

type Status struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	ResetDate   string `json:"resetDate"`
	Description string `json:"description"`
	Stats       struct {
		Agents    int `json:"agents"`
		Ships     int `json:"ships"`
		Systems   int `json:"systems"`
		Waypoints int `json:"waypoints"`
	} `json:"stats"`
	Leaderboards struct {
		MostCredits []struct {
			AgentSymbol string `json:"agentSymbol"`
			Credits     int    `json:"credits"`
		} `json:"mostCredits"`
		MostSubmittedCharts []struct {
			AgentSymbol string `json:"agentSymbol"`
			ChartCount  int    `json:"chartCount"`
		} `json:"mostSubmittedCharts"`
	} `json:"leaderboards"`
	ServerResets struct {
		Next      time.Time `json:"next"`
		Frequency string    `json:"frequency"`
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

func GetStatus() (status Status, err error) {
	err = get(BaseUrl, nil, &status)
	if err != nil {
		return status, err
	}
	return status, nil
}
