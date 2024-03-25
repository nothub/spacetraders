# SiphonYield

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Units** | **int32** | The number of units siphoned that were placed into the ship&#39;s cargo hold. | 

## Methods

### NewSiphonYield

`func NewSiphonYield(symbol TradeSymbol, units int32, ) *SiphonYield`

NewSiphonYield instantiates a new SiphonYield object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiphonYieldWithDefaults

`func NewSiphonYieldWithDefaults() *SiphonYield`

NewSiphonYieldWithDefaults instantiates a new SiphonYield object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SiphonYield) GetSymbol() TradeSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SiphonYield) GetSymbolOk() (*TradeSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SiphonYield) SetSymbol(v TradeSymbol)`

SetSymbol sets Symbol field to given value.


### GetUnits

`func (o *SiphonYield) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SiphonYield) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *SiphonYield) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


