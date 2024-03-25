# RefuelShipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | Pointer to **int32** | The amount of fuel to fill in the ship&#39;s tanks. When not specified, the ship will be refueled to its maximum fuel capacity. If the amount specified is greater than the ship&#39;s remaining capacity, the ship will only be refueled to its maximum fuel capacity. The amount specified is not in market units but in ship fuel units. | [optional] 
**FromCargo** | Pointer to **bool** | Wether to use the FUEL thats in your cargo or not. Default: false | [optional] 

## Methods

### NewRefuelShipRequest

`func NewRefuelShipRequest() *RefuelShipRequest`

NewRefuelShipRequest instantiates a new RefuelShipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefuelShipRequestWithDefaults

`func NewRefuelShipRequestWithDefaults() *RefuelShipRequest`

NewRefuelShipRequestWithDefaults instantiates a new RefuelShipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *RefuelShipRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *RefuelShipRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *RefuelShipRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *RefuelShipRequest) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetFromCargo

`func (o *RefuelShipRequest) GetFromCargo() bool`

GetFromCargo returns the FromCargo field if non-nil, zero value otherwise.

### GetFromCargoOk

`func (o *RefuelShipRequest) GetFromCargoOk() (*bool, bool)`

GetFromCargoOk returns a tuple with the FromCargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCargo

`func (o *RefuelShipRequest) SetFromCargo(v bool)`

SetFromCargo sets FromCargo field to given value.

### HasFromCargo

`func (o *RefuelShipRequest) HasFromCargo() bool`

HasFromCargo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


