# ShipCargo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | **int32** | The max number of items that can be stored in the cargo hold. | 
**Units** | **int32** | The number of items currently stored in the cargo hold. | 
**Inventory** | [**[]ShipCargoItem**](ShipCargoItem.md) | The items currently in the cargo hold. | 

## Methods

### NewShipCargo

`func NewShipCargo(capacity int32, units int32, inventory []ShipCargoItem, ) *ShipCargo`

NewShipCargo instantiates a new ShipCargo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipCargoWithDefaults

`func NewShipCargoWithDefaults() *ShipCargo`

NewShipCargoWithDefaults instantiates a new ShipCargo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *ShipCargo) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ShipCargo) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ShipCargo) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.


### GetUnits

`func (o *ShipCargo) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ShipCargo) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ShipCargo) SetUnits(v int32)`

SetUnits sets Units field to given value.


### GetInventory

`func (o *ShipCargo) GetInventory() []ShipCargoItem`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ShipCargo) GetInventoryOk() (*[]ShipCargoItem, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ShipCargo) SetInventory(v []ShipCargoItem)`

SetInventory sets Inventory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


