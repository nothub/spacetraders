# ExtractionYield

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Units** | **int32** | The number of units extracted that were placed into the ship&#39;s cargo hold. | 

## Methods

### NewExtractionYield

`func NewExtractionYield(symbol TradeSymbol, units int32, ) *ExtractionYield`

NewExtractionYield instantiates a new ExtractionYield object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionYieldWithDefaults

`func NewExtractionYieldWithDefaults() *ExtractionYield`

NewExtractionYieldWithDefaults instantiates a new ExtractionYield object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ExtractionYield) GetSymbol() TradeSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ExtractionYield) GetSymbolOk() (*TradeSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ExtractionYield) SetSymbol(v TradeSymbol)`

SetSymbol sets Symbol field to given value.


### GetUnits

`func (o *ExtractionYield) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ExtractionYield) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ExtractionYield) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


