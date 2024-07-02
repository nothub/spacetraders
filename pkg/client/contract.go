package st

import (
	"strconv"
	"time"
)

// Contract details.
type Contract struct {

	// ID of the contract.
	//
	//   >= 1 characters
	//
	// (required)
	Id string `json:"id"`

	// The symbol of the faction that this contract is for.
	//
	//   >= 1 characters
	//
	// (required)
	FactionSymbol string `json:"factionSymbol"`

	// Type of contract.
	//
	// (required)
	Type ContractType `json:"type"`

	// The terms to fulfill the contract.
	//
	// (required)
	Terms ContractTerms `json:"terms"`

	// Whether the contract has been accepted by the agent
	//
	// (required)
	Accepted bool `json:"accepted"`

	// Whether the contract has been fulfilled
	//
	// (required)
	Fulfilled bool `json:"fulfilled"`

	// Deprecated in favor of deadlineToAccept
	//
	// (deprecated) (required)
	Expiration time.Time `json:"expiration"`

	// The time at which the contract is no longer available to be accepted
	DeadlineToAccept time.Time `json:"deadlineToAccept"`
}

// Type of contract.
type ContractType string

const (
	ContractTypeProcurement ContractType = "PROCUREMENT"
	ContractTypeTransport   ContractType = "TRANSPORT"
	ContractTypeShuttle     ContractType = "SHUTTLE"
)

// Payments for the contract.
type ContractPayment struct {

	// The amount of credits received up front for accepting the contract.
	//
	// (required)
	OnAccepted int `json:"onAccepted"`

	// The amount of credits received when the contract is fulfilled.
	//
	// (required)
	OnFulfilled int `json:"onFulfilled"`
}

// The details of a delivery contract. Includes the type of good, units needed, and the destination.
type ContractDeliverGood struct {

	// The symbol of the trade good to deliver.
	//
	//   >= 1 characters
	//
	// (required)
	TradeSymbol TradeGoodSymbol `json:"tradeSymbol"`

	// The destination where goods need to be delivered.
	//
	//   >= 1 characters
	//
	// (required)
	DestinationSymbol string `json:"destinationSymbol"`

	// The number of units that need to be delivered on this contract.
	//
	// (required)
	UnitsRequired int `json:"unitsRequired"`

	// The number of units fulfilled on this contract.
	//
	// (required)
	UnitsFulfilled int `json:"unitsFulfilled"`
}

// The terms to fulfill the contract.
type ContractTerms struct {

	// The deadline for the contract.
	//
	// (required)
	Deadline time.Time `json:"deadline"`

	// Payments for the contract.
	//
	// (required)
	Payment ContractPayment `json:"payment"`

	// The cargo that needs to be delivered to fulfill the contract.
	Deliver []ContractDeliverGood `json:"deliver"`
}

func ListContracts(limit int, page int) (contracts []Contract, meta Meta, err error) {
	var dto struct {
		Data []Contract `json:"data"`
		Meta Meta       `json:"meta"`
	}

	err = get(BaseUrl+"/my/contracts", map[string]string{
		"limit": strconv.Itoa(limit),
		"page":  strconv.Itoa(page),
	}, &dto)
	if err != nil {
		return contracts, meta, err
	}

	return dto.Data, dto.Meta, nil
}

func GetContract(contractId string) (contract Contract, err error) {
	var dto struct {
		Data Contract `json:"data"`
	}

	err = get(BaseUrl+"/my/contracts/"+contractId, nil, &dto)
	if err != nil {
		return contract, err
	}

	return dto.Data, nil
}

// You can only accept contracts that were offered to you, were not accepted
// yet, and whose deadlines has not passed yet.
func AcceptContract(contractId string) (agent Agent, contract Contract, err error) {
	var dto struct {
		Data struct {
			Agent    Agent    `json:"agent"`
			Contract Contract `json:"contract"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/contracts/"+contractId+"/accept", nil, nil, &dto)
	if err != nil {
		return agent, contract, err
	}

	return dto.Data.Agent, dto.Data.Contract, nil
}

func DeliverCargoToContract(contractId string, shipSymbol string, tradeSymbol string, units int) (contract Contract, cargo ShipCargo, err error) {
	var dto struct {
		Data struct {
			Contract Contract  `json:"contract"`
			Cargo    ShipCargo `json:"cargo"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/contracts/"+contractId+"/deliver", nil, struct {
		ShipSymbol  string `json:"shipSymbol"`
		TradeSymbol string `json:"tradeSymbol"`
		Units       int    `json:"units"`
	}{
		ShipSymbol:  shipSymbol,
		TradeSymbol: tradeSymbol,
		Units:       units,
	}, &dto)
	if err != nil {
		return contract, cargo, err
	}

	return dto.Data.Contract, dto.Data.Cargo, nil
}

// Can only be used on contracts that have all of their delivery terms
// fulfilled.
func FulfillContract(contractId string) (agent Agent, contract Contract, err error) {
	var dto struct {
		Data struct {
			Agent    Agent    `json:"agent"`
			Contract Contract `json:"contract"`
		} `json:"data"`
	}

	err = post(BaseUrl+"/my/contracts/"+contractId+"/fulfill", nil, nil, &dto)
	if err != nil {
		return agent, contract, err
	}

	return dto.Data.Agent, dto.Data.Contract, nil
}
