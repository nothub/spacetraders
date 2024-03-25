# ShipRefine201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Produced** | [**[]ShipRefine201ResponseDataProducedInner**](ShipRefine201ResponseDataProducedInner.md) | Goods that were produced by this refining process. | 
**Consumed** | [**[]ShipRefine201ResponseDataProducedInner**](ShipRefine201ResponseDataProducedInner.md) | Goods that were consumed during this refining process. | 

## Methods

### NewShipRefine201ResponseData

`func NewShipRefine201ResponseData(cargo ShipCargo, cooldown Cooldown, produced []ShipRefine201ResponseDataProducedInner, consumed []ShipRefine201ResponseDataProducedInner, ) *ShipRefine201ResponseData`

NewShipRefine201ResponseData instantiates a new ShipRefine201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipRefine201ResponseDataWithDefaults

`func NewShipRefine201ResponseDataWithDefaults() *ShipRefine201ResponseData`

NewShipRefine201ResponseDataWithDefaults instantiates a new ShipRefine201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCargo

`func (o *ShipRefine201ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *ShipRefine201ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *ShipRefine201ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetCooldown

`func (o *ShipRefine201ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *ShipRefine201ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *ShipRefine201ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetProduced

`func (o *ShipRefine201ResponseData) GetProduced() []ShipRefine201ResponseDataProducedInner`

GetProduced returns the Produced field if non-nil, zero value otherwise.

### GetProducedOk

`func (o *ShipRefine201ResponseData) GetProducedOk() (*[]ShipRefine201ResponseDataProducedInner, bool)`

GetProducedOk returns a tuple with the Produced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduced

`func (o *ShipRefine201ResponseData) SetProduced(v []ShipRefine201ResponseDataProducedInner)`

SetProduced sets Produced field to given value.


### GetConsumed

`func (o *ShipRefine201ResponseData) GetConsumed() []ShipRefine201ResponseDataProducedInner`

GetConsumed returns the Consumed field if non-nil, zero value otherwise.

### GetConsumedOk

`func (o *ShipRefine201ResponseData) GetConsumedOk() (*[]ShipRefine201ResponseDataProducedInner, bool)`

GetConsumedOk returns a tuple with the Consumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumed

`func (o *ShipRefine201ResponseData) SetConsumed(v []ShipRefine201ResponseDataProducedInner)`

SetConsumed sets Consumed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


