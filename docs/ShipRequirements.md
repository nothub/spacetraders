# ShipRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Power** | Pointer to **int32** | The amount of power required from the reactor. | [optional] 
**Crew** | Pointer to **int32** | The number of crew required for operation. | [optional] 
**Slots** | Pointer to **int32** | The number of module slots required for installation. | [optional] 

## Methods

### NewShipRequirements

`func NewShipRequirements() *ShipRequirements`

NewShipRequirements instantiates a new ShipRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipRequirementsWithDefaults

`func NewShipRequirementsWithDefaults() *ShipRequirements`

NewShipRequirementsWithDefaults instantiates a new ShipRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPower

`func (o *ShipRequirements) GetPower() int32`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *ShipRequirements) GetPowerOk() (*int32, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *ShipRequirements) SetPower(v int32)`

SetPower sets Power field to given value.

### HasPower

`func (o *ShipRequirements) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetCrew

`func (o *ShipRequirements) GetCrew() int32`

GetCrew returns the Crew field if non-nil, zero value otherwise.

### GetCrewOk

`func (o *ShipRequirements) GetCrewOk() (*int32, bool)`

GetCrewOk returns a tuple with the Crew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrew

`func (o *ShipRequirements) SetCrew(v int32)`

SetCrew sets Crew field to given value.

### HasCrew

`func (o *ShipRequirements) HasCrew() bool`

HasCrew returns a boolean if a field has been set.

### GetSlots

`func (o *ShipRequirements) GetSlots() int32`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *ShipRequirements) GetSlotsOk() (*int32, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *ShipRequirements) SetSlots(v int32)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *ShipRequirements) HasSlots() bool`

HasSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


