# CreateShipWaypointScan201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Waypoints** | [**[]ScannedWaypoint**](ScannedWaypoint.md) | List of scanned waypoints. | 

## Methods

### NewCreateShipWaypointScan201ResponseData

`func NewCreateShipWaypointScan201ResponseData(cooldown Cooldown, waypoints []ScannedWaypoint, ) *CreateShipWaypointScan201ResponseData`

NewCreateShipWaypointScan201ResponseData instantiates a new CreateShipWaypointScan201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShipWaypointScan201ResponseDataWithDefaults

`func NewCreateShipWaypointScan201ResponseDataWithDefaults() *CreateShipWaypointScan201ResponseData`

NewCreateShipWaypointScan201ResponseDataWithDefaults instantiates a new CreateShipWaypointScan201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *CreateShipWaypointScan201ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *CreateShipWaypointScan201ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *CreateShipWaypointScan201ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetWaypoints

`func (o *CreateShipWaypointScan201ResponseData) GetWaypoints() []ScannedWaypoint`

GetWaypoints returns the Waypoints field if non-nil, zero value otherwise.

### GetWaypointsOk

`func (o *CreateShipWaypointScan201ResponseData) GetWaypointsOk() (*[]ScannedWaypoint, bool)`

GetWaypointsOk returns a tuple with the Waypoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypoints

`func (o *CreateShipWaypointScan201ResponseData) SetWaypoints(v []ScannedWaypoint)`

SetWaypoints sets Waypoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


