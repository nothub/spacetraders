# ShipModificationTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaypointSymbol** | **string** | The symbol of the waypoint where the transaction took place. | 
**ShipSymbol** | **string** | The symbol of the ship that made the transaction. | 
**TradeSymbol** | **string** | The symbol of the trade good. | 
**TotalPrice** | **int32** | The total price of the transaction. | 
**Timestamp** | **time.Time** | The timestamp of the transaction. | 

## Methods

### NewShipModificationTransaction

`func NewShipModificationTransaction(waypointSymbol string, shipSymbol string, tradeSymbol string, totalPrice int32, timestamp time.Time, ) *ShipModificationTransaction`

NewShipModificationTransaction instantiates a new ShipModificationTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipModificationTransactionWithDefaults

`func NewShipModificationTransactionWithDefaults() *ShipModificationTransaction`

NewShipModificationTransactionWithDefaults instantiates a new ShipModificationTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaypointSymbol

`func (o *ShipModificationTransaction) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *ShipModificationTransaction) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *ShipModificationTransaction) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.


### GetShipSymbol

`func (o *ShipModificationTransaction) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *ShipModificationTransaction) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *ShipModificationTransaction) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetTradeSymbol

`func (o *ShipModificationTransaction) GetTradeSymbol() string`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *ShipModificationTransaction) GetTradeSymbolOk() (*string, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *ShipModificationTransaction) SetTradeSymbol(v string)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetTotalPrice

`func (o *ShipModificationTransaction) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *ShipModificationTransaction) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *ShipModificationTransaction) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.


### GetTimestamp

`func (o *ShipModificationTransaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShipModificationTransaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShipModificationTransaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


