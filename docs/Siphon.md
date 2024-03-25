# Siphon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipSymbol** | **string** | Symbol of the ship that executed the siphon. | 
**Yield** | [**SiphonYield**](SiphonYield.md) |  | 

## Methods

### NewSiphon

`func NewSiphon(shipSymbol string, yield SiphonYield, ) *Siphon`

NewSiphon instantiates a new Siphon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiphonWithDefaults

`func NewSiphonWithDefaults() *Siphon`

NewSiphonWithDefaults instantiates a new Siphon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipSymbol

`func (o *Siphon) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *Siphon) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *Siphon) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetYield

`func (o *Siphon) GetYield() SiphonYield`

GetYield returns the Yield field if non-nil, zero value otherwise.

### GetYieldOk

`func (o *Siphon) GetYieldOk() (*SiphonYield, bool)`

GetYieldOk returns a tuple with the Yield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYield

`func (o *Siphon) SetYield(v SiphonYield)`

SetYield sets Yield field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


