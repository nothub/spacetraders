# ShipCrew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **int32** | The current number of crew members on the ship. | 
**Required** | **int32** | The minimum number of crew members required to maintain the ship. | 
**Capacity** | **int32** | The maximum number of crew members the ship can support. | 
**Rotation** | **string** | The rotation of crew shifts. A stricter shift improves the ship&#39;s performance. A more relaxed shift improves the crew&#39;s morale. | [default to "STRICT"]
**Morale** | **int32** | A rough measure of the crew&#39;s morale. A higher morale means the crew is happier and more productive. A lower morale means the ship is more prone to accidents. | 
**Wages** | **int32** | The amount of credits per crew member paid per hour. Wages are paid when a ship docks at a civilized waypoint. | 

## Methods

### NewShipCrew

`func NewShipCrew(current int32, required int32, capacity int32, rotation string, morale int32, wages int32, ) *ShipCrew`

NewShipCrew instantiates a new ShipCrew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipCrewWithDefaults

`func NewShipCrewWithDefaults() *ShipCrew`

NewShipCrewWithDefaults instantiates a new ShipCrew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *ShipCrew) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ShipCrew) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ShipCrew) SetCurrent(v int32)`

SetCurrent sets Current field to given value.


### GetRequired

`func (o *ShipCrew) GetRequired() int32`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ShipCrew) GetRequiredOk() (*int32, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ShipCrew) SetRequired(v int32)`

SetRequired sets Required field to given value.


### GetCapacity

`func (o *ShipCrew) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ShipCrew) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ShipCrew) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetRotation

`func (o *ShipCrew) GetRotation() string`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *ShipCrew) GetRotationOk() (*string, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *ShipCrew) SetRotation(v string)`

SetRotation sets Rotation field to given value.


### GetMorale

`func (o *ShipCrew) GetMorale() int32`

GetMorale returns the Morale field if non-nil, zero value otherwise.

### GetMoraleOk

`func (o *ShipCrew) GetMoraleOk() (*int32, bool)`

GetMoraleOk returns a tuple with the Morale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMorale

`func (o *ShipCrew) SetMorale(v int32)`

SetMorale sets Morale field to given value.


### GetWages

`func (o *ShipCrew) GetWages() int32`

GetWages returns the Wages field if non-nil, zero value otherwise.

### GetWagesOk

`func (o *ShipCrew) GetWagesOk() (*int32, bool)`

GetWagesOk returns a tuple with the Wages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWages

`func (o *ShipCrew) SetWages(v int32)`

SetWages sets Wages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


