# ConstructionMaterial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeSymbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Required** | **int32** | The number of units required. | 
**Fulfilled** | **int32** | The number of units fulfilled toward the required amount. | 

## Methods

### NewConstructionMaterial

`func NewConstructionMaterial(tradeSymbol TradeSymbol, required int32, fulfilled int32, ) *ConstructionMaterial`

NewConstructionMaterial instantiates a new ConstructionMaterial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructionMaterialWithDefaults

`func NewConstructionMaterialWithDefaults() *ConstructionMaterial`

NewConstructionMaterialWithDefaults instantiates a new ConstructionMaterial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeSymbol

`func (o *ConstructionMaterial) GetTradeSymbol() TradeSymbol`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *ConstructionMaterial) GetTradeSymbolOk() (*TradeSymbol, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *ConstructionMaterial) SetTradeSymbol(v TradeSymbol)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetRequired

`func (o *ConstructionMaterial) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ConstructionMaterial) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ConstructionMaterial) SetRequired(v int32)`

SetRequired sets Required field to given value.


### GetFulfilled

`func (o *ConstructionMaterial) GetFulfilled() int32`

GetFulfilled returns the Fulfilled field if non-nil, zero value otherwise.

### GetFulfilledOk

`func (o *ConstructionMaterial) GetFulfilledOk() (*int32, bool)`

GetFulfilledOk returns a tuple with the Fulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilled

`func (o *ConstructionMaterial) SetFulfilled(v int32)`

SetFulfilled sets Fulfilled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


