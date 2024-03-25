# ShipFuelConsumed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | The amount of fuel consumed by the most recent transit or action. | 
**Timestamp** | **time.Time** | The time at which the fuel was consumed. | 

## Methods

### NewShipFuelConsumed

`func NewShipFuelConsumed(amount int32, timestamp time.Time, ) *ShipFuelConsumed`

NewShipFuelConsumed instantiates a new ShipFuelConsumed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipFuelConsumedWithDefaults

`func NewShipFuelConsumedWithDefaults() *ShipFuelConsumed`

NewShipFuelConsumedWithDefaults instantiates a new ShipFuelConsumed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ShipFuelConsumed) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ShipFuelConsumed) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ShipFuelConsumed) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetTimestamp

`func (o *ShipFuelConsumed) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ShipFuelConsumed) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ShipFuelConsumed) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


