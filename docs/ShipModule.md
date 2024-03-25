# ShipModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the module. | 
**Capacity** | Pointer to **int32** | Modules that provide capacity, such as cargo hold or crew quarters will show this value to denote how much of a bonus the module grants. | [optional] 
**Range** | Pointer to **int32** | Modules that have a range will such as a sensor array show this value to denote how far can the module reach with its capabilities. | [optional] 
**Name** | **string** | Name of this module. | 
**Description** | **string** | Description of this module. | 
**Requirements** | [**ShipRequirements**](ShipRequirements.md) |  | 

## Methods

### NewShipModule

`func NewShipModule(symbol string, name string, description string, requirements ShipRequirements, ) *ShipModule`

NewShipModule instantiates a new ShipModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipModuleWithDefaults

`func NewShipModuleWithDefaults() *ShipModule`

NewShipModuleWithDefaults instantiates a new ShipModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipModule) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipModule) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipModule) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetCapacity

`func (o *ShipModule) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ShipModule) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ShipModule) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ShipModule) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetRange

`func (o *ShipModule) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ShipModule) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ShipModule) SetRange(v int32)`

SetRange sets Range field to given value.

### HasRange

`func (o *ShipModule) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetName

`func (o *ShipModule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipModule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipModule) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipModule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipModule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipModule) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRequirements

`func (o *ShipModule) GetRequirements() ShipRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ShipModule) GetRequirementsOk() (*ShipRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ShipModule) SetRequirements(v ShipRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


