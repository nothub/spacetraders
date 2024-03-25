# SystemWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the waypoint. | 
**Type** | [**WaypointType**](WaypointType.md) |  | 
**X** | **int32** | Relative position of the waypoint on the system&#39;s x axis. This is not an absolute position in the universe. | 
**Y** | **int32** | Relative position of the waypoint on the system&#39;s y axis. This is not an absolute position in the universe. | 
**Orbitals** | [**[]WaypointOrbital**](WaypointOrbital.md) | Waypoints that orbit this waypoint. | 
**Orbits** | Pointer to **string** | The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint. Otherwise this value is undefined. | [optional] 

## Methods

### NewSystemWaypoint

`func NewSystemWaypoint(symbol string, type_ WaypointType, x int32, y int32, orbitals []WaypointOrbital, ) *SystemWaypoint`

NewSystemWaypoint instantiates a new SystemWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWaypointWithDefaults

`func NewSystemWaypointWithDefaults() *SystemWaypoint`

NewSystemWaypointWithDefaults instantiates a new SystemWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SystemWaypoint) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SystemWaypoint) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SystemWaypoint) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *SystemWaypoint) GetType() WaypointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemWaypoint) GetTypeOk() (*WaypointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemWaypoint) SetType(v WaypointType)`

SetType sets Type field to given value.


### GetX

`func (o *SystemWaypoint) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *SystemWaypoint) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *SystemWaypoint) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *SystemWaypoint) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *SystemWaypoint) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *SystemWaypoint) SetY(v int32)`

SetY sets Y field to given value.


### GetOrbitals

`func (o *SystemWaypoint) GetOrbitals() []WaypointOrbital`

GetOrbitals returns the Orbitals field if non-nil, zero value otherwise.

### GetOrbitalsOk

`func (o *SystemWaypoint) GetOrbitalsOk() (*[]WaypointOrbital, bool)`

GetOrbitalsOk returns a tuple with the Orbitals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbitals

`func (o *SystemWaypoint) SetOrbitals(v []WaypointOrbital)`

SetOrbitals sets Orbitals field to given value.


### GetOrbits

`func (o *SystemWaypoint) GetOrbits() string`

GetOrbits returns the Orbits field if non-nil, zero value otherwise.

### GetOrbitsOk

`func (o *SystemWaypoint) GetOrbitsOk() (*string, bool)`

GetOrbitsOk returns a tuple with the Orbits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrbits

`func (o *SystemWaypoint) SetOrbits(v string)`

SetOrbits sets Orbits field to given value.

### HasOrbits

`func (o *SystemWaypoint) HasOrbits() bool`

HasOrbits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


