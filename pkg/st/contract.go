package st

import "time"

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

func ListContracts() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/b5d513949b11a-list-contracts
	return nil
}

func GetContract() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/2889d8b056533-get-contract
	return nil
}

func AcceptContract() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/7dbc359629250-accept-contract
	return nil
}

func DeliverCargoToContract() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/8f89f3b4a246e-deliver-cargo-to-contract
	return nil
}

func FulfillContract() (err error) {
	// TODO: https://spacetraders.stoplight.io/docs/spacetraders/d4ff41c101af0-fulfill-contract
	return nil
}
