# ShipEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the engine. | 
**Name** | **string** | The name of the engine. | 
**Description** | **string** | The description of the engine. | 
**Condition** | **float64** | The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases. | 
**Integrity** | **float64** | The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time. | 
**Speed** | **int32** | The speed stat of this engine. The higher the speed, the faster a ship can travel from one point to another. Reduces the time of arrival when navigating the ship. | 
**Requirements** | [**ShipRequirements**](ShipRequirements.md) |  | 

## Methods

### NewShipEngine

`func NewShipEngine(symbol string, name string, description string, condition float64, integrity float64, speed int32, requirements ShipRequirements, ) *ShipEngine`

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

`func (o *ShipEngine) GetCondition() float64`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ShipEngine) GetConditionOk() (*float64, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ShipEngine) SetCondition(v float64)`

SetCondition sets Condition field to given value.


### GetIntegrity

`func (o *ShipEngine) GetIntegrity() float64`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *ShipEngine) GetIntegrityOk() (*float64, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *ShipEngine) SetIntegrity(v float64)`

SetIntegrity sets Integrity field to given value.


### GetSpeed

`func (o *ShipEngine) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ShipEngine) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ShipEngine) SetSpeed(v int32)`

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


