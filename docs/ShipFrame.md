# ShipFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol of the frame. | 
**Name** | **string** | Name of the frame. | 
**Description** | **string** | Description of the frame. | 
**Condition** | **float64** | The repairable condition of a component. A value of 0 indicates the component needs significant repairs, while a value of 1 indicates the component is in near perfect condition. As the condition of a component is repaired, the overall integrity of the component decreases. | 
**Integrity** | **float64** | The overall integrity of the component, which determines the performance of the component. A value of 0 indicates that the component is almost completely degraded, while a value of 1 indicates that the component is in near perfect condition. The integrity of the component is non-repairable, and represents permanent wear over time. | 
**ModuleSlots** | **int32** | The amount of slots that can be dedicated to modules installed in the ship. Each installed module take up a number of slots, and once there are no more slots, no new modules can be installed. | 
**MountingPoints** | **int32** | The amount of slots that can be dedicated to mounts installed in the ship. Each installed mount takes up a number of points, and once there are no more points remaining, no new mounts can be installed. | 
**FuelCapacity** | **int32** | The maximum amount of fuel that can be stored in this ship. When refueling, the ship will be refueled to this amount. | 
**Requirements** | [**ShipRequirements**](ShipRequirements.md) |  | 

## Methods

### NewShipFrame

`func NewShipFrame(symbol string, name string, description string, condition float64, integrity float64, moduleSlots int32, mountingPoints int32, fuelCapacity int32, requirements ShipRequirements, ) *ShipFrame`

NewShipFrame instantiates a new ShipFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipFrameWithDefaults

`func NewShipFrameWithDefaults() *ShipFrame`

NewShipFrameWithDefaults instantiates a new ShipFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipFrame) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipFrame) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipFrame) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ShipFrame) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipFrame) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipFrame) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipFrame) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipFrame) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipFrame) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCondition

`func (o *ShipFrame) GetCondition() float64`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ShipFrame) GetConditionOk() (*float64, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ShipFrame) SetCondition(v float64)`

SetCondition sets Condition field to given value.


### GetIntegrity

`func (o *ShipFrame) GetIntegrity() float64`

GetIntegrity returns the Integrity field if non-nil, zero value otherwise.

### GetIntegrityOk

`func (o *ShipFrame) GetIntegrityOk() (*float64, bool)`

GetIntegrityOk returns a tuple with the Integrity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrity

`func (o *ShipFrame) SetIntegrity(v float64)`

SetIntegrity sets Integrity field to given value.


### GetModuleSlots

`func (o *ShipFrame) GetModuleSlots() int32`

GetModuleSlots returns the ModuleSlots field if non-nil, zero value otherwise.

### GetModuleSlotsOk

`func (o *ShipFrame) GetModuleSlotsOk() (*int32, bool)`

GetModuleSlotsOk returns a tuple with the ModuleSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleSlots

`func (o *ShipFrame) SetModuleSlots(v int32)`

SetModuleSlots sets ModuleSlots field to given value.


### GetMountingPoints

`func (o *ShipFrame) GetMountingPoints() int32`

GetMountingPoints returns the MountingPoints field if non-nil, zero value otherwise.

### GetMountingPointsOk

`func (o *ShipFrame) GetMountingPointsOk() (*int32, bool)`

GetMountingPointsOk returns a tuple with the MountingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountingPoints

`func (o *ShipFrame) SetMountingPoints(v int32)`

SetMountingPoints sets MountingPoints field to given value.


### GetFuelCapacity

`func (o *ShipFrame) GetFuelCapacity() int32`

GetFuelCapacity returns the FuelCapacity field if non-nil, zero value otherwise.

### GetFuelCapacityOk

`func (o *ShipFrame) GetFuelCapacityOk() (*int32, bool)`

GetFuelCapacityOk returns a tuple with the FuelCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelCapacity

`func (o *ShipFrame) SetFuelCapacity(v int32)`

SetFuelCapacity sets FuelCapacity field to given value.


### GetRequirements

`func (o *ShipFrame) GetRequirements() ShipRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ShipFrame) GetRequirementsOk() (*ShipRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ShipFrame) SetRequirements(v ShipRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


