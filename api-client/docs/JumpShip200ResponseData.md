# JumpShip200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Nav** | Pointer to [**ShipNav**](ShipNav.md) |  | [optional] 

## Methods

### NewJumpShip200ResponseData

`func NewJumpShip200ResponseData(cooldown Cooldown, ) *JumpShip200ResponseData`

NewJumpShip200ResponseData instantiates a new JumpShip200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpShip200ResponseDataWithDefaults

`func NewJumpShip200ResponseDataWithDefaults() *JumpShip200ResponseData`

NewJumpShip200ResponseDataWithDefaults instantiates a new JumpShip200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *JumpShip200ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *JumpShip200ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *JumpShip200ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetNav

`func (o *JumpShip200ResponseData) GetNav() ShipNav`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *JumpShip200ResponseData) GetNavOk() (*ShipNav, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *JumpShip200ResponseData) SetNav(v ShipNav)`

SetNav sets Nav field to given value.

### HasNav

`func (o *JumpShip200ResponseData) HasNav() bool`

HasNav returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


