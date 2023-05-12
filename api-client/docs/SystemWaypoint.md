# SystemWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**Type** | [**WaypointType**](WaypointType.md) |  | 
**X** | **int32** |  | 
**Y** | **int32** |  | 

## Methods

### NewSystemWaypoint

`func NewSystemWaypoint(symbol string, type_ WaypointType, x int32, y int32, ) *SystemWaypoint`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


