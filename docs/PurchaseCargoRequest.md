# PurchaseCargoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Units** | **int32** | Amounts of units to purchase. | 

## Methods

### NewPurchaseCargoRequest

`func NewPurchaseCargoRequest(symbol TradeSymbol, units int32, ) *PurchaseCargoRequest`

NewPurchaseCargoRequest instantiates a new PurchaseCargoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseCargoRequestWithDefaults

`func NewPurchaseCargoRequestWithDefaults() *PurchaseCargoRequest`

NewPurchaseCargoRequestWithDefaults instantiates a new PurchaseCargoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *PurchaseCargoRequest) GetSymbol() TradeSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *PurchaseCargoRequest) GetSymbolOk() (*TradeSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *PurchaseCargoRequest) SetSymbol(v TradeSymbol)`

SetSymbol sets Symbol field to given value.


### GetUnits

`func (o *PurchaseCargoRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *PurchaseCargoRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *PurchaseCargoRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


