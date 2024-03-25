# ShipReactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol of the reactor. | 
**Name** | **string** | Name of the reactor. | 
**Description** | **string** | Description of the reactor. | 
**Condition** | **float64** | The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases. | 
**Integrity** | **float64** | The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time. | 
**PowerOutput** | **int32** | The amount of power provided by this reactor. The more power a reactor provides to the ship, the lower the cooldown it gets when using a module or mount that taxes the ship&#39;s power. | 
**Requirements** | [**ShipRequirements**](ShipRequirements.md) |  | 

## Methods

### NewShipReactor

`func NewShipReactor(symbol string, name string, description string, condition float64, integrity float64, powerOutput int32, requirements ShipRequirements, ) *ShipReactor`

NewShipReactor instantiates a new ShipReactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipReactorWithDefaults

`func NewShipReactorWithDefaults() *ShipReactor`

NewShipReactorWithDefaults instantiates a new ShipReactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipReactor) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipReactor) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipReactor) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ShipReactor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipReactor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipReactor) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipReactor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipReactor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipReactor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCondition

`func (o *ShipReactor) GetCondition() float64`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ShipReactor) GetConditionOk() (*float64, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ShipReactor) SetCondition(v float64)`

SetCondition sets Condition field to given value.


### GetIntegrity

`func (o *ShipReactor) GetIntegrity() float64`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *ShipReactor) GetIntegrityOk() (*float64, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *ShipReactor) SetIntegrity(v float64)`

SetIntegrity sets Integrity field to given value.


### GetPowerOutput

`func (o *ShipReactor) GetPowerOutput() int32`

GetPowerOutput returns the PowerOutput field if non-nil, zero value otherwise.

### GetPowerOutputOk

`func (o *ShipReactor) GetPowerOutputOk() (*int32, bool)`

GetPowerOutputOk returns a tuple with the PowerOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutput

`func (o *ShipReactor) SetPowerOutput(v int32)`

SetPowerOutput sets PowerOutput field to given value.


### GetRequirements

`func (o *ShipReactor) GetRequirements() ShipRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ShipReactor) GetRequirementsOk() (*ShipRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ShipReactor) SetRequirements(v ShipRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


