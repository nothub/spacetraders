# CreateShipSystemScan201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Systems** | [**[]ScannedSystem**](ScannedSystem.md) | List of scanned systems. | 

## Methods

### NewCreateShipSystemScan201ResponseData

`func NewCreateShipSystemScan201ResponseData(cooldown Cooldown, systems []ScannedSystem, ) *CreateShipSystemScan201ResponseData`

NewCreateShipSystemScan201ResponseData instantiates a new CreateShipSystemScan201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShipSystemScan201ResponseDataWithDefaults

`func NewCreateShipSystemScan201ResponseDataWithDefaults() *CreateShipSystemScan201ResponseData`

NewCreateShipSystemScan201ResponseDataWithDefaults instantiates a new CreateShipSystemScan201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *CreateShipSystemScan201ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *CreateShipSystemScan201ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *CreateShipSystemScan201ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetSystems

`func (o *CreateShipSystemScan201ResponseData) GetSystems() []ScannedSystem`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *CreateShipSystemScan201ResponseData) GetSystemsOk() (*[]ScannedSystem, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *CreateShipSystemScan201ResponseData) SetSystems(v []ScannedSystem)`

SetSystems sets Systems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


