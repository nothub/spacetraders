# ShipRefine200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Produced** | [**[]ShipRefine200ResponseDataProducedInner**](ShipRefine200ResponseDataProducedInner.md) |  | 
**Consumed** | [**[]ShipRefine200ResponseDataProducedInner**](ShipRefine200ResponseDataProducedInner.md) |  | 

## Methods

### NewShipRefine200ResponseData

`func NewShipRefine200ResponseData(cargo ShipCargo, cooldown Cooldown, produced []ShipRefine200ResponseDataProducedInner, consumed []ShipRefine200ResponseDataProducedInner, ) *ShipRefine200ResponseData`

NewShipRefine200ResponseData instantiates a new ShipRefine200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipRefine200ResponseDataWithDefaults

`func NewShipRefine200ResponseDataWithDefaults() *ShipRefine200ResponseData`

NewShipRefine200ResponseDataWithDefaults instantiates a new ShipRefine200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCargo

`func (o *ShipRefine200ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *ShipRefine200ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *ShipRefine200ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetCooldown

`func (o *ShipRefine200ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *ShipRefine200ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *ShipRefine200ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetProduced

`func (o *ShipRefine200ResponseData) GetProduced() []ShipRefine200ResponseDataProducedInner`

GetProduced returns the Produced field if non-nil, zero value otherwise.

### GetProducedOk

`func (o *ShipRefine200ResponseData) GetProducedOk() (*[]ShipRefine200ResponseDataProducedInner, bool)`

GetProducedOk returns a tuple with the Produced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduced

`func (o *ShipRefine200ResponseData) SetProduced(v []ShipRefine200ResponseDataProducedInner)`

SetProduced sets Produced field to given value.


### GetConsumed

`func (o *ShipRefine200ResponseData) GetConsumed() []ShipRefine200ResponseDataProducedInner`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *ShipRefine200ResponseData) GetConsumedOk() (*[]ShipRefine200ResponseDataProducedInner, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *ShipRefine200ResponseData) SetConsumed(v []ShipRefine200ResponseDataProducedInner)`

SetConsumed sets Consumed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


