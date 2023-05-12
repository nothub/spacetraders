# Waypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Type** | [**WaypointType**](WaypointType.md) |  | 
**SystemSymbol** | **string** |  | 
**X** | **int32** |  | 
**Y** | **int32** |  | 
**Orbitals** | [**[]WaypointOrbital**](WaypointOrbital.md) |  | 
**Faction** | Pointer to [**WaypointFaction**](WaypointFaction.md) |  | [optional] 
**Traits** | [**[]WaypointTrait**](WaypointTrait.md) | The traits of the waypoint. | 
**Chart** | Pointer to [**Chart**](Chart.md) |  | [optional] 

## Methods

### NewWaypoint

`func NewWaypoint(symbol string, type_ WaypointType, systemSymbol string, x int32, y int32, orbitals []WaypointOrbital, traits []WaypointTrait, ) *Waypoint`

NewWaypoint instantiates a new Waypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWaypointWithDefaults

`func NewWaypointWithDefaults() *Waypoint`

NewWaypointWithDefaults instantiates a new Waypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Waypoint) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Waypoint) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Waypoint) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *Waypoint) GetType() WaypointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Waypoint) GetTypeOk() (*WaypointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Waypoint) SetType(v WaypointType)`

SetType sets Type field to given value.


### GetSystemSymbol

`func (o *Waypoint) GetSystemSymbol() string`

GetSystemSymbol returns the SystemSymbol field if non-nil, zero value otherwise.

### GetSystemSymbolOk

`func (o *Waypoint) GetSystemSymbolOk() (*string, bool)`

GetSystemSymbolOk returns a tuple with the SystemSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSymbol

`func (o *Waypoint) SetSystemSymbol(v string)`

SetSystemSymbol sets SystemSymbol field to given value.


### GetX

`func (o *Waypoint) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *Waypoint) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *Waypoint) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *Waypoint) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *Waypoint) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *Waypoint) SetY(v int32)`

SetY sets Y field to given value.


### GetOrbitals

`func (o *Waypoint) GetOrbitals() []WaypointOrbital`

GetOrbitals returns the Orbitals field if non-nil, zero value otherwise.

### GetOrbitalsOk

`func (o *Waypoint) GetOrbitalsOk() (*[]WaypointOrbital, bool)`

GetOrbitalsOk returns a tuple with the Orbitals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbitals

`func (o *Waypoint) SetOrbitals(v []WaypointOrbital)`

SetOrbitals sets Orbitals field to given value.


### GetFaction

`func (o *Waypoint) GetFaction() WaypointFaction`

GetFaction returns the Faction field if non-nil, zero value otherwise.

### GetFactionOk

`func (o *Waypoint) GetFactionOk() (*WaypointFaction, bool)`

GetFactionOk returns a tuple with the Faction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaction

`func (o *Waypoint) SetFaction(v WaypointFaction)`

SetFaction sets Faction field to given value.

### HasFaction

`func (o *Waypoint) HasFaction() bool`

HasFaction returns a boolean if a field has been set.

### GetTraits

`func (o *Waypoint) GetTraits() []WaypointTrait`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *Waypoint) GetTraitsOk() (*[]WaypointTrait, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *Waypoint) SetTraits(v []WaypointTrait)`

SetTraits sets Traits field to given value.


### GetChart

`func (o *Waypoint) GetChart() Chart`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *Waypoint) GetChartOk() (*Chart, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *Waypoint) SetChart(v Chart)`

SetChart sets Chart field to given value.

### HasChart

`func (o *Waypoint) HasChart() bool`

HasChart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


