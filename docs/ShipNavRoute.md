# ShipNavRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**ShipNavRouteWaypoint**](ShipNavRouteWaypoint.md) |  | 
**Origin** | [**ShipNavRouteWaypoint**](ShipNavRouteWaypoint.md) |  | 
**DepartureTime** | **time.Time** | The date time of the ship&#39;s departure. | 
**Arrival** | **time.Time** | The date time of the ship&#39;s arrival. If the ship is in-transit, this is the expected time of arrival. | 

## Methods

### NewShipNavRoute

`func NewShipNavRoute(destination ShipNavRouteWaypoint, origin ShipNavRouteWaypoint, departureTime time.Time, arrival time.Time, ) *ShipNavRoute`

NewShipNavRoute instantiates a new ShipNavRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipNavRouteWithDefaults

`func NewShipNavRouteWithDefaults() *ShipNavRoute`

NewShipNavRouteWithDefaults instantiates a new ShipNavRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *ShipNavRoute) GetDestination() ShipNavRouteWaypoint`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ShipNavRoute) GetDestinationOk() (*ShipNavRouteWaypoint, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ShipNavRoute) SetDestination(v ShipNavRouteWaypoint)`

SetDestination sets Destination field to given value.


### GetOrigin

`func (o *ShipNavRoute) GetOrigin() ShipNavRouteWaypoint`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ShipNavRoute) GetOriginOk() (*ShipNavRouteWaypoint, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ShipNavRoute) SetOrigin(v ShipNavRouteWaypoint)`

SetOrigin sets Origin field to given value.


### GetDepartureTime

`func (o *ShipNavRoute) GetDepartureTime() time.Time`

GetDepartureTime returns the DepartureTime field if non-nil, zero value otherwise.

### GetDepartureTimeOk

`func (o *ShipNavRoute) GetDepartureTimeOk() (*time.Time, bool)`

GetDepartureTimeOk returns a tuple with the DepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureTime

`func (o *ShipNavRoute) SetDepartureTime(v time.Time)`

SetDepartureTime sets DepartureTime field to given value.


### GetArrival

`func (o *ShipNavRoute) GetArrival() time.Time`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *ShipNavRoute) GetArrivalOk() (*time.Time, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *ShipNavRoute) SetArrival(v time.Time)`

SetArrival sets Arrival field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


