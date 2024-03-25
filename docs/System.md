# System

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the system. | 
**SectorSymbol** | **string** | The symbol of the sector. | 
**Type** | [**SystemType**](SystemType.md) |  | 
**X** | **int32** | Relative position of the system in the sector in the x axis. | 
**Y** | **int32** | Relative position of the system in the sector in the y axis. | 
**Waypoints** | [**[]SystemWaypoint**](SystemWaypoint.md) | Waypoints in this system. | 
**Factions** | [**[]SystemFaction**](SystemFaction.md) | Factions that control this system. | 

## Methods

### NewSystem

`func NewSystem(symbol string, sectorSymbol string, type_ SystemType, x int32, y int32, waypoints []SystemWaypoint, factions []SystemFaction, ) *System`

NewSystem instantiates a new System object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemWithDefaults

`func NewSystemWithDefaults() *System`

NewSystemWithDefaults instantiates a new System object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *System) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *System) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *System) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSectorSymbol

`func (o *System) GetSectorSymbol() string`

GetSectorSymbol returns the SectorSymbol field if non-nil, zero value otherwise.

### GetSectorSymbolOk

`func (o *System) GetSectorSymbolOk() (*string, bool)`

GetSectorSymbolOk returns a tuple with the SectorSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSymbol

`func (o *System) SetSectorSymbol(v string)`

SetSectorSymbol sets SectorSymbol field to given value.


### GetType

`func (o *System) GetType() SystemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *System) GetTypeOk() (*SystemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *System) SetType(v SystemType)`

SetType sets Type field to given value.


### GetX

`func (o *System) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *System) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *System) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *System) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *System) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *System) SetY(v int32)`

SetY sets Y field to given value.


### GetWaypoints

`func (o *System) GetWaypoints() []SystemWaypoint`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *System) GetWaypointsOk() (*[]SystemWaypoint, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *System) SetWaypoints(v []SystemWaypoint)`

SetWaypoints sets Waypoints field to given value.


### GetFactions

`func (o *System) GetFactions() []SystemFaction`

GetFactions returns the Factions field if non-nil, zero value otherwise.

### GetFactionsOk

`func (o *System) GetFactionsOk() (*[]SystemFaction, bool)`

GetFactionsOk returns a tuple with the Factions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactions

`func (o *System) SetFactions(v []SystemFaction)`

SetFactions sets Factions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


