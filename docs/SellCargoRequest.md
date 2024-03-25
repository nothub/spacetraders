# SellCargoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Units** | **int32** | Amounts of units to sell of the selected good. | 

## Methods

### NewSellCargoRequest

`func NewSellCargoRequest(symbol TradeSymbol, units int32, ) *SellCargoRequest`

NewSellCargoRequest instantiates a new SellCargoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSellCargoRequestWithDefaults

`func NewSellCargoRequestWithDefaults() *SellCargoRequest`

NewSellCargoRequestWithDefaults instantiates a new SellCargoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SellCargoRequest) GetSymbol() TradeSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SellCargoRequest) GetSymbolOk() (*TradeSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SellCargoRequest) SetSymbol(v TradeSymbol)`

SetSymbol sets Symbol field to given value.


### GetUnits

`func (o *SellCargoRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SellCargoRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *SellCargoRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


