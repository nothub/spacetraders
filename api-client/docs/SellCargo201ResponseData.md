# SellCargo201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**Agent**](Agent.md) |  | 
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Transaction** | [**MarketTransaction**](MarketTransaction.md) |  | 

## Methods

### NewSellCargo201ResponseData

`func NewSellCargo201ResponseData(agent Agent, cargo ShipCargo, transaction MarketTransaction, ) *SellCargo201ResponseData`

NewSellCargo201ResponseData instantiates a new SellCargo201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellCargo201ResponseDataWithDefaults

`func NewSellCargo201ResponseDataWithDefaults() *SellCargo201ResponseData`

NewSellCargo201ResponseDataWithDefaults instantiates a new SellCargo201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *SellCargo201ResponseData) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *SellCargo201ResponseData) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *SellCargo201ResponseData) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetCargo

`func (o *SellCargo201ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *SellCargo201ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *SellCargo201ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetTransaction

`func (o *SellCargo201ResponseData) GetTransaction() MarketTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *SellCargo201ResponseData) GetTransactionOk() (*MarketTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *SellCargo201ResponseData) SetTransaction(v MarketTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


