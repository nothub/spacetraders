# ShipNavRouteWaypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the waypoint. | 
**Type** | [**WaypointType**](WaypointType.md) |  | 
**SystemSymbol** | **string** | The symbol of the system. | 
**X** | **int32** | Position in the universe in the x axis. | 
**Y** | **int32** | Position in the universe in the y axis. | 

## Methods

### NewShipNavRouteWaypoint

`func NewShipNavRouteWaypoint(symbol string, type_ WaypointType, systemSymbol string, x int32, y int32, ) *ShipNavRouteWaypoint`

NewShipNavRouteWaypoint instantiates a new ShipNavRouteWaypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipNavRouteWaypointWithDefaults

`func NewShipNavRouteWaypointWithDefaults() *ShipNavRouteWaypoint`

NewShipNavRouteWaypointWithDefaults instantiates a new ShipNavRouteWaypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipNavRouteWaypoint) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipNavRouteWaypoint) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipNavRouteWaypoint) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *ShipNavRouteWaypoint) GetType() WaypointType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipNavRouteWaypoint) GetTypeOk() (*WaypointType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipNavRouteWaypoint) SetType(v WaypointType)`

SetType sets Type field to given value.


### GetSystemSymbol

`func (o *ShipNavRouteWaypoint) GetSystemSymbol() string`

GetSystemSymbol returns the SystemSymbol field if non-nil, zero value otherwise.

### GetSystemSymbolOk

`func (o *ShipNavRouteWaypoint) GetSystemSymbolOk() (*string, bool)`

GetSystemSymbolOk returns a tuple with the SystemSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSymbol

`func (o *ShipNavRouteWaypoint) SetSystemSymbol(v string)`

SetSystemSymbol sets SystemSymbol field to given value.


### GetX

`func (o *ShipNavRouteWaypoint) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ShipNavRouteWaypoint) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ShipNavRouteWaypoint) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *ShipNavRouteWaypoint) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ShipNavRouteWaypoint) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ShipNavRouteWaypoint) SetY(v int32)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


