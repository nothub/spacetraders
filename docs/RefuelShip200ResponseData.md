# RefuelShip200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**Agent**](Agent.md) |  | 
**Fuel** | [**ShipFuel**](ShipFuel.md) |  | 
**Transaction** | [**MarketTransaction**](MarketTransaction.md) |  | 

## Methods

### NewRefuelShip200ResponseData

`func NewRefuelShip200ResponseData(agent Agent, fuel ShipFuel, transaction MarketTransaction, ) *RefuelShip200ResponseData`

NewRefuelShip200ResponseData instantiates a new RefuelShip200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefuelShip200ResponseDataWithDefaults

`func NewRefuelShip200ResponseDataWithDefaults() *RefuelShip200ResponseData`

NewRefuelShip200ResponseDataWithDefaults instantiates a new RefuelShip200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *RefuelShip200ResponseData) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *RefuelShip200ResponseData) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *RefuelShip200ResponseData) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetFuel

`func (o *RefuelShip200ResponseData) GetFuel() ShipFuel`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *RefuelShip200ResponseData) GetFuelOk() (*ShipFuel, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *RefuelShip200ResponseData) SetFuel(v ShipFuel)`

SetFuel sets Fuel field to given value.


### GetTransaction

`func (o *RefuelShip200ResponseData) GetTransaction() MarketTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *RefuelShip200ResponseData) GetTransactionOk() (*MarketTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *RefuelShip200ResponseData) SetTransaction(v MarketTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


