# ScannedWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the waypoint. | 
**Type** | [**WaypointType**](WaypointType.md) |  | 
**SystemSymbol** | **string** | The symbol of the system. | 
**X** | **int32** | Position in the universe in the x axis. | 
**Y** | **int32** | Position in the universe in the y axis. | 
**Orbitals** | [**[]WaypointOrbital**](WaypointOrbital.md) | List of waypoints that orbit this waypoint. | 
**Faction** | Pointer to [**WaypointFaction**](WaypointFaction.md) |  | [optional] 
**Traits** | [**[]WaypointTrait**](WaypointTrait.md) | The traits of the waypoint. | 
**Chart** | Pointer to [**Chart**](Chart.md) |  | [optional] 

## Methods

### NewScannedWaypoint

`func NewScannedWaypoint(symbol string, type_ WaypointType, systemSymbol string, x int32, y int32, orbitals []WaypointOrbital, traits []WaypointTrait, ) *ScannedWaypoint`

NewScannedWaypoint instantiates a new ScannedWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScannedWaypointWithDefaults

`func NewScannedWaypointWithDefaults() *ScannedWaypoint`

NewScannedWaypointWithDefaults instantiates a new ScannedWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ScannedWaypoint) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ScannedWaypoint) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ScannedWaypoint) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *ScannedWaypoint) GetType() WaypointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScannedWaypoint) GetTypeOk() (*WaypointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScannedWaypoint) SetType(v WaypointType)`

SetType sets Type field to given value.


### GetSystemSymbol

`func (o *ScannedWaypoint) GetSystemSymbol() string`

GetSystemSymbol returns the SystemSymbol field if non-nil, zero value otherwise.

### GetSystemSymbolOk

`func (o *ScannedWaypoint) GetSystemSymbolOk() (*string, bool)`

GetSystemSymbolOk returns a tuple with the SystemSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSymbol

`func (o *ScannedWaypoint) SetSystemSymbol(v string)`

SetSystemSymbol sets SystemSymbol field to given value.


### GetX

`func (o *ScannedWaypoint) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ScannedWaypoint) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ScannedWaypoint) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *ScannedWaypoint) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ScannedWaypoint) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ScannedWaypoint) SetY(v int32)`

SetY sets Y field to given value.


### GetOrbitals

`func (o *ScannedWaypoint) GetOrbitals() []WaypointOrbital`

GetOrbitals returns the Orbitals field if non-nil, zero value otherwise.

### GetOrbitalsOk

`func (o *ScannedWaypoint) GetOrbitalsOk() (*[]WaypointOrbital, bool)`

GetOrbitalsOk returns a tuple with the Orbitals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbitals

`func (o *ScannedWaypoint) SetOrbitals(v []WaypointOrbital)`

SetOrbitals sets Orbitals field to given value.


### GetFaction

`func (o *ScannedWaypoint) GetFaction() WaypointFaction`

GetFaction returns the Faction field if non-nil, zero value otherwise.

### GetFactionOk

`func (o *ScannedWaypoint) GetFactionOk() (*WaypointFaction, bool)`

GetFactionOk returns a tuple with the Faction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaction

`func (o *ScannedWaypoint) SetFaction(v WaypointFaction)`

SetFaction sets Faction field to given value.

### HasFaction

`func (o *ScannedWaypoint) HasFaction() bool`

HasFaction returns a boolean if a field has been set.

### GetTraits

`func (o *ScannedWaypoint) GetTraits() []WaypointTrait`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *ScannedWaypoint) GetTraitsOk() (*[]WaypointTrait, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *ScannedWaypoint) SetTraits(v []WaypointTrait)`

SetTraits sets Traits field to given value.


### GetChart

`func (o *ScannedWaypoint) GetChart() Chart`

GetChart returns the Chart field if non-nil, zero value otherwise.

### GetChartOk

`func (o *ScannedWaypoint) GetChartOk() (*Chart, bool)`

GetChartOk returns a tuple with the Chart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChart

`func (o *ScannedWaypoint) SetChart(v Chart)`

SetChart sets Chart field to given value.

### HasChart

`func (o *ScannedWaypoint) HasChart() bool`

HasChart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


