# PatchShipNavRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlightMode** | Pointer to [**ShipNavFlightMode**](ShipNavFlightMode.md) |  | [optional] [default to SHIPNAVFLIGHTMODE_CRUISE]

## Methods

### NewPatchShipNavRequest

`func NewPatchShipNavRequest() *PatchShipNavRequest`

NewPatchShipNavRequest instantiates a new PatchShipNavRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchShipNavRequestWithDefaults

`func NewPatchShipNavRequestWithDefaults() *PatchShipNavRequest`

NewPatchShipNavRequestWithDefaults instantiates a new PatchShipNavRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlightMode

`func (o *PatchShipNavRequest) GetFlightMode() ShipNavFlightMode`

GetFlightMode returns the FlightMode field if non-nil, zero value otherwise.

### GetFlightModeOk

`func (o *PatchShipNavRequest) GetFlightModeOk() (*ShipNavFlightMode, bool)`

GetFlightModeOk returns a tuple with the FlightMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlightMode

`func (o *PatchShipNavRequest) SetFlightMode(v ShipNavFlightMode)`

SetFlightMode sets FlightMode field to given value.

### HasFlightMode

`func (o *PatchShipNavRequest) HasFlightMode() bool`

HasFlightMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


