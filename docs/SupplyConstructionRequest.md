# SupplyConstructionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipSymbol** | **string** | Symbol of the ship to use. | 
**TradeSymbol** | **string** | The symbol of the good to supply. | 
**Units** | **int32** | Amount of units to supply. | 

## Methods

### NewSupplyConstructionRequest

`func NewSupplyConstructionRequest(shipSymbol string, tradeSymbol string, units int32, ) *SupplyConstructionRequest`

NewSupplyConstructionRequest instantiates a new SupplyConstructionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplyConstructionRequestWithDefaults

`func NewSupplyConstructionRequestWithDefaults() *SupplyConstructionRequest`

NewSupplyConstructionRequestWithDefaults instantiates a new SupplyConstructionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipSymbol

`func (o *SupplyConstructionRequest) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *SupplyConstructionRequest) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *SupplyConstructionRequest) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetTradeSymbol

`func (o *SupplyConstructionRequest) GetTradeSymbol() string`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *SupplyConstructionRequest) GetTradeSymbolOk() (*string, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *SupplyConstructionRequest) SetTradeSymbol(v string)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetUnits

`func (o *SupplyConstructionRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SupplyConstructionRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *SupplyConstructionRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


