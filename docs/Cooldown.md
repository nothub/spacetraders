# Cooldown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipSymbol** | **string** | The symbol of the ship that is on cooldown | 
**TotalSeconds** | **int32** | The total duration of the cooldown in seconds | 
**RemainingSeconds** | **int32** | The remaining duration of the cooldown in seconds | 
**Expiration** | Pointer to **time.Time** | The date and time when the cooldown expires in ISO 8601 format | [optional] 

## Methods

### NewCooldown

`func NewCooldown(shipSymbol string, totalSeconds int32, remainingSeconds int32, ) *Cooldown`

NewCooldown instantiates a new Cooldown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCooldownWithDefaults

`func NewCooldownWithDefaults() *Cooldown`

NewCooldownWithDefaults instantiates a new Cooldown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipSymbol

`func (o *Cooldown) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *Cooldown) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *Cooldown) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetTotalSeconds

`func (o *Cooldown) GetTotalSeconds() int32`

GetTotalSeconds returns the TotalSeconds field if non-nil, zero value otherwise.

### GetTotalSecondsOk

`func (o *Cooldown) GetTotalSecondsOk() (*int32, bool)`

GetTotalSecondsOk returns a tuple with the TotalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeconds

`func (o *Cooldown) SetTotalSeconds(v int32)`

SetTotalSeconds sets TotalSeconds field to given value.


### GetRemainingSeconds

`func (o *Cooldown) GetRemainingSeconds() int32`

GetRemainingSeconds returns the RemainingSeconds field if non-nil, zero value otherwise.

### GetRemainingSecondsOk

`func (o *Cooldown) GetRemainingSecondsOk() (*int32, bool)`

GetRemainingSecondsOk returns a tuple with the RemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingSeconds

`func (o *Cooldown) SetRemainingSeconds(v int32)`

SetRemainingSeconds sets RemainingSeconds field to given value.


### GetExpiration

`func (o *Cooldown) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Cooldown) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Cooldown) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *Cooldown) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


