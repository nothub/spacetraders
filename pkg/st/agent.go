package st

import (
	"strconv"
)

// Agent details.
type Agent struct {

	// Account ID that is tied to this agent.
	// Only included on your own agent.
	//
	//   >= 1 characters
	AccountId string `json:"accountId"`

	// Symbol of the agent.
	//
	//   >= 3 characters && <= 14 characters
	//
	// (required)
	Symbol string `json:"symbol"`

	// The headquarters of the agent.
	//
	//   >= 1 characters
	//
	// (required)
	Headquarters string `json:"headquarters"`

	// The number of credits the agent has available.
	// Credits can be negative if funds have been overdrawn.
	//
	//   >= -9007199254740991 && <= 9007199254740991
	//
	// (required)
	Credits int64 `json:"credits"`

	// The faction the agent started with.
	//
	//   >= 1 characters
	//
	// (required)
	StartingFaction string `json:"startingFaction"`

	// How many ships are owned by the agent.
	//
	// (required)
	ShipCount int `json:"shipCount"`
}

func GetAgent() (agent Agent, err error) {
	var dto struct {
		Data Agent `json:"data"`
	}

	err = get(BaseUrl+"/my/agent", nil, &dto)
	if err != nil {
		return dto.Data, err
	}

	return dto.Data, nil
}

func GetPublicAgent(agentSymbol string) (agent Agent, err error) {
	var dto struct {
		Data Agent `json:"data"`
	}

	err = get(BaseUrl+"/agents/"+agentSymbol, nil, &dto)
	if err != nil {
		return dto.Data, err
	}

	return dto.Data, nil
}

// limit - How many entries to return per page
//
//	>= 1 && <= 20
//
// page - What entry offset to request (Index starts at 1)
//
//	>= 1
func ListAgents(limit int, page int) (agents []Agent, meta Meta, err error) {
	var dto struct {
		Data []Agent `json:"data"`
		Meta Meta    `json:"meta"`
	}

	err = get(BaseUrl+"/agents", map[string]string{
		"limit": strconv.Itoa(limit),
		"page":  strconv.Itoa(page),
	}, &dto)
	if err != nil {
		return dto.Data, dto.Meta, err
	}

	return dto.Data, dto.Meta, nil
}
