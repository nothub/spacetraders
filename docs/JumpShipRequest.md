# JumpShipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaypointSymbol** | **string** | The symbol of the waypoint to jump to. The destination must be a connected waypoint. | 

## Methods

### NewJumpShipRequest

`func NewJumpShipRequest(waypointSymbol string, ) *JumpShipRequest`

NewJumpShipRequest instantiates a new JumpShipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpShipRequestWithDefaults

`func NewJumpShipRequestWithDefaults() *JumpShipRequest`

NewJumpShipRequestWithDefaults instantiates a new JumpShipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaypointSymbol

`func (o *JumpShipRequest) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *JumpShipRequest) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *JumpShipRequest) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


