# PurchaseShipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipType** | [**ShipType**](ShipType.md) |  | 
**WaypointSymbol** | **string** | The symbol of the waypoint you want to purchase the ship at. | 

## Methods

### NewPurchaseShipRequest

`func NewPurchaseShipRequest(shipType ShipType, waypointSymbol string, ) *PurchaseShipRequest`

NewPurchaseShipRequest instantiates a new PurchaseShipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseShipRequestWithDefaults

`func NewPurchaseShipRequestWithDefaults() *PurchaseShipRequest`

NewPurchaseShipRequestWithDefaults instantiates a new PurchaseShipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipType

`func (o *PurchaseShipRequest) GetShipType() ShipType`

GetShipType returns the ShipType field if non-nil, zero value otherwise.

### GetShipTypeOk

`func (o *PurchaseShipRequest) GetShipTypeOk() (*ShipType, bool)`

GetShipTypeOk returns a tuple with the ShipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipType

`func (o *PurchaseShipRequest) SetShipType(v ShipType)`

SetShipType sets ShipType field to given value.


### GetWaypointSymbol

`func (o *PurchaseShipRequest) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *PurchaseShipRequest) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *PurchaseShipRequest) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


