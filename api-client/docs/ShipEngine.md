# ShipEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Condition** | Pointer to **int32** | Condition is a range of 0 to 100 where 0 is completely worn out and 100 is brand new. | [optional] 
**Speed** | **float32** |  | 
**Requirements** | [**ShipRequirements**](ShipRequirements.md) |  | 

## Methods

### NewShipEngine

`func NewShipEngine(symbol string, name string, description string, speed float32, requirements ShipRequirements, ) *ShipEngine`

NewShipEngine instantiates a new ShipEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipEngineWithDefaults

`func NewShipEngineWithDefaults() *ShipEngine`

NewShipEngineWithDefaults instantiates a new ShipEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipEngine) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipEngine) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipEngine) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ShipEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipEngine) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipEngine) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipEngine) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipEngine) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCondition

`func (o *ShipEngine) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ShipEngine) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ShipEngine) SetCondition(v int32)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ShipEngine) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetSpeed

`func (o *ShipEngine) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ShipEngine) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ShipEngine) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.


### GetRequirements

`func (o *ShipEngine) GetRequirements() ShipRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ShipEngine) GetRequirementsOk() (*ShipRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ShipEngine) SetRequirements(v ShipRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


