# CreateShipShipScan201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Ships** | [**[]ScannedShip**](ScannedShip.md) | List of scanned ships. | 

## Methods

### NewCreateShipShipScan201ResponseData

`func NewCreateShipShipScan201ResponseData(cooldown Cooldown, ships []ScannedShip, ) *CreateShipShipScan201ResponseData`

NewCreateShipShipScan201ResponseData instantiates a new CreateShipShipScan201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShipShipScan201ResponseDataWithDefaults

`func NewCreateShipShipScan201ResponseDataWithDefaults() *CreateShipShipScan201ResponseData`

NewCreateShipShipScan201ResponseDataWithDefaults instantiates a new CreateShipShipScan201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *CreateShipShipScan201ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *CreateShipShipScan201ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *CreateShipShipScan201ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetShips

`func (o *CreateShipShipScan201ResponseData) GetShips() []ScannedShip`

GetShips returns the Ships field if non-nil, zero value otherwise.

### GetShipsOk

`func (o *CreateShipShipScan201ResponseData) GetShipsOk() (*[]ScannedShip, bool)`

GetShipsOk returns a tuple with the Ships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShips

`func (o *CreateShipShipScan201ResponseData) SetShips(v []ScannedShip)`

SetShips sets Ships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


