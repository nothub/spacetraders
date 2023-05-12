# ShipFuel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **int32** | The current amount of fuel in the ship&#39;s tanks. | 
**Capacity** | **int32** | The maximum amount of fuel the ship&#39;s tanks can hold. | 
**Consumed** | Pointer to [**ShipFuelConsumed**](ShipFuelConsumed.md) |  | [optional] 

## Methods

### NewShipFuel

`func NewShipFuel(current int32, capacity int32, ) *ShipFuel`

NewShipFuel instantiates a new ShipFuel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipFuelWithDefaults

`func NewShipFuelWithDefaults() *ShipFuel`

NewShipFuelWithDefaults instantiates a new ShipFuel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *ShipFuel) GetCurrent() int32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *ShipFuel) GetCurrentOk() (*int32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *ShipFuel) SetCurrent(v int32)`

SetCurrent sets Current field to given value.


### GetCapacity

`func (o *ShipFuel) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ShipFuel) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ShipFuel) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetConsumed

`func (o *ShipFuel) GetConsumed() ShipFuelConsumed`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *ShipFuel) GetConsumedOk() (*ShipFuelConsumed, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *ShipFuel) SetConsumed(v ShipFuelConsumed)`

SetConsumed sets Consumed field to given value.

### HasConsumed

`func (o *ShipFuel) HasConsumed() bool`

HasConsumed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


