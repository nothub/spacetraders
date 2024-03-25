# ShipNav

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemSymbol** | **string** | The symbol of the system. | 
**WaypointSymbol** | **string** | The symbol of the waypoint. | 
**Route** | [**ShipNavRoute**](ShipNavRoute.md) |  | 
**Status** | [**ShipNavStatus**](ShipNavStatus.md) |  | 
**FlightMode** | [**ShipNavFlightMode**](ShipNavFlightMode.md) |  | [default to SHIPNAVFLIGHTMODE_CRUISE]

## Methods

### NewShipNav

`func NewShipNav(systemSymbol string, waypointSymbol string, route ShipNavRoute, status ShipNavStatus, flightMode ShipNavFlightMode, ) *ShipNav`

NewShipNav instantiates a new ShipNav object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipNavWithDefaults

`func NewShipNavWithDefaults() *ShipNav`

NewShipNavWithDefaults instantiates a new ShipNav object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemSymbol

`func (o *ShipNav) GetSystemSymbol() string`

GetSystemSymbol returns the SystemSymbol field if non-nil, zero value otherwise.

### GetSystemSymbolOk

`func (o *ShipNav) GetSystemSymbolOk() (*string, bool)`

GetSystemSymbolOk returns a tuple with the SystemSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemSymbol

`func (o *ShipNav) SetSystemSymbol(v string)`

SetSystemSymbol sets SystemSymbol field to given value.


### GetWaypointSymbol

`func (o *ShipNav) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *ShipNav) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *ShipNav) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.


### GetRoute

`func (o *ShipNav) GetRoute() ShipNavRoute`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *ShipNav) GetRouteOk() (*ShipNavRoute, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *ShipNav) SetRoute(v ShipNavRoute)`

SetRoute sets Route field to given value.


### GetStatus

`func (o *ShipNav) GetStatus() ShipNavStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipNav) GetStatusOk() (*ShipNavStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipNav) SetStatus(v ShipNavStatus)`

SetStatus sets Status field to given value.


### GetFlightMode

`func (o *ShipNav) GetFlightMode() ShipNavFlightMode`

GetFlightMode returns the FlightMode field if non-nil, zero value otherwise.

### GetFlightModeOk

`func (o *ShipNav) GetFlightModeOk() (*ShipNavFlightMode, bool)`

GetFlightModeOk returns a tuple with the FlightMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightMode

`func (o *ShipNav) SetFlightMode(v ShipNavFlightMode)`

SetFlightMode sets FlightMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


