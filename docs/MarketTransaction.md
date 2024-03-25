# MarketTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaypointSymbol** | **string** | The symbol of the waypoint. | 
**ShipSymbol** | **string** | The symbol of the ship that made the transaction. | 
**TradeSymbol** | **string** | The symbol of the trade good. | 
**Type** | **string** | The type of transaction. | 
**Units** | **int32** | The number of units of the transaction. | 
**PricePerUnit** | **int32** | The price per unit of the transaction. | 
**TotalPrice** | **int32** | The total price of the transaction. | 
**Timestamp** | **time.Time** | The timestamp of the transaction. | 

## Methods

### NewMarketTransaction

`func NewMarketTransaction(waypointSymbol string, shipSymbol string, tradeSymbol string, type_ string, units int32, pricePerUnit int32, totalPrice int32, timestamp time.Time, ) *MarketTransaction`

NewMarketTransaction instantiates a new MarketTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTransactionWithDefaults

`func NewMarketTransactionWithDefaults() *MarketTransaction`

NewMarketTransactionWithDefaults instantiates a new MarketTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaypointSymbol

`func (o *MarketTransaction) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *MarketTransaction) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *MarketTransaction) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.


### GetShipSymbol

`func (o *MarketTransaction) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *MarketTransaction) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *MarketTransaction) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetTradeSymbol

`func (o *MarketTransaction) GetTradeSymbol() string`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *MarketTransaction) GetTradeSymbolOk() (*string, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *MarketTransaction) SetTradeSymbol(v string)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetType

`func (o *MarketTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetUnits

`func (o *MarketTransaction) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *MarketTransaction) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *MarketTransaction) SetUnits(v int32)`

SetUnits sets Units field to given value.


### GetPricePerUnit

`func (o *MarketTransaction) GetPricePerUnit() int32`

GetPricePerUnit returns the PricePerUnit field if non-nil, zero value otherwise.

### GetPricePerUnitOk

`func (o *MarketTransaction) GetPricePerUnitOk() (*int32, bool)`

GetPricePerUnitOk returns a tuple with the PricePerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerUnit

`func (o *MarketTransaction) SetPricePerUnit(v int32)`

SetPricePerUnit sets PricePerUnit field to given value.


### GetTotalPrice

`func (o *MarketTransaction) GetTotalPrice() int32`

GetTotalPrice returns the TotalPrice field if non-nil, zero value otherwise.

### GetTotalPriceOk

`func (o *MarketTransaction) GetTotalPriceOk() (*int32, bool)`

GetTotalPriceOk returns a tuple with the TotalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrice

`func (o *MarketTransaction) SetTotalPrice(v int32)`

SetTotalPrice sets TotalPrice field to given value.


### GetTimestamp

`func (o *MarketTransaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MarketTransaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MarketTransaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


