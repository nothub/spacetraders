# SiphonResources201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Siphon** | [**Siphon**](Siphon.md) |  | 
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Events** | [**[]ExtractResources201ResponseDataEventsInner**](ExtractResources201ResponseDataEventsInner.md) |  | 

## Methods

### NewSiphonResources201ResponseData

`func NewSiphonResources201ResponseData(cooldown Cooldown, siphon Siphon, cargo ShipCargo, events []ExtractResources201ResponseDataEventsInner, ) *SiphonResources201ResponseData`

NewSiphonResources201ResponseData instantiates a new SiphonResources201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiphonResources201ResponseDataWithDefaults

`func NewSiphonResources201ResponseDataWithDefaults() *SiphonResources201ResponseData`

NewSiphonResources201ResponseDataWithDefaults instantiates a new SiphonResources201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *SiphonResources201ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *SiphonResources201ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *SiphonResources201ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetSiphon

`func (o *SiphonResources201ResponseData) GetSiphon() Siphon`

GetSiphon returns the Siphon field if non-nil, zero value otherwise.

### GetSiphonOk

`func (o *SiphonResources201ResponseData) GetSiphonOk() (*Siphon, bool)`

GetSiphonOk returns a tuple with the Siphon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiphon

`func (o *SiphonResources201ResponseData) SetSiphon(v Siphon)`

SetSiphon sets Siphon field to given value.


### GetCargo

`func (o *SiphonResources201ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *SiphonResources201ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *SiphonResources201ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetEvents

`func (o *SiphonResources201ResponseData) GetEvents() []ExtractResources201ResponseDataEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SiphonResources201ResponseData) GetEventsOk() (*[]ExtractResources201ResponseDataEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SiphonResources201ResponseData) SetEvents(v []ExtractResources201ResponseDataEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


