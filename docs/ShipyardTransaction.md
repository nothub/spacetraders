# ShipyardTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaypointSymbol** | **string** | The symbol of the waypoint. | 
**ShipSymbol** | **string** | The symbol of the ship that was the subject of the transaction. | 
**ShipType** | **string** | The symbol of the ship that was the subject of the transaction. | 
**Price** | **int32** | The price of the transaction. | 
**AgentSymbol** | **string** | The symbol of the agent that made the transaction. | 
**Timestamp** | **time.Time** | The timestamp of the transaction. | 

## Methods

### NewShipyardTransaction

`func NewShipyardTransaction(waypointSymbol string, shipSymbol string, shipType string, price int32, agentSymbol string, timestamp time.Time, ) *ShipyardTransaction`

NewShipyardTransaction instantiates a new ShipyardTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipyardTransactionWithDefaults

`func NewShipyardTransactionWithDefaults() *ShipyardTransaction`

NewShipyardTransactionWithDefaults instantiates a new ShipyardTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaypointSymbol

`func (o *ShipyardTransaction) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *ShipyardTransaction) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *ShipyardTransaction) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.


### GetShipSymbol

`func (o *ShipyardTransaction) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *ShipyardTransaction) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *ShipyardTransaction) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetShipType

`func (o *ShipyardTransaction) GetShipType() string`

GetShipType returns the ShipType field if non-nil, zero value otherwise.

### GetShipTypeOk

`func (o *ShipyardTransaction) GetShipTypeOk() (*string, bool)`

GetShipTypeOk returns a tuple with the ShipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipType

`func (o *ShipyardTransaction) SetShipType(v string)`

SetShipType sets ShipType field to given value.


### GetPrice

`func (o *ShipyardTransaction) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ShipyardTransaction) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ShipyardTransaction) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetAgentSymbol

`func (o *ShipyardTransaction) GetAgentSymbol() string`

GetAgentSymbol returns the AgentSymbol field if non-nil, zero value otherwise.

### GetAgentSymbolOk

`func (o *ShipyardTransaction) GetAgentSymbolOk() (*string, bool)`

GetAgentSymbolOk returns a tuple with the AgentSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSymbol

`func (o *ShipyardTransaction) SetAgentSymbol(v string)`

SetAgentSymbol sets AgentSymbol field to given value.


### GetTimestamp

`func (o *ShipyardTransaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShipyardTransaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShipyardTransaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


