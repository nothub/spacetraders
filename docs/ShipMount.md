# ShipMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbo of this mount. | 
**Name** | **string** | Name of this mount. | 
**Description** | Pointer to **string** | Description of this mount. | [optional] 
**Strength** | Pointer to **int32** | Mounts that have this value, such as mining lasers, denote how powerful this mount&#39;s capabilities are. | [optional] 
**Deposits** | Pointer to **[]string** | Mounts that have this value denote what goods can be produced from using the mount. | [optional] 
**Requirements** | [**ShipRequirements**](ShipRequirements.md) |  | 

## Methods

### NewShipMount

`func NewShipMount(symbol string, name string, requirements ShipRequirements, ) *ShipMount`

NewShipMount instantiates a new ShipMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipMountWithDefaults

`func NewShipMountWithDefaults() *ShipMount`

NewShipMountWithDefaults instantiates a new ShipMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipMount) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipMount) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipMount) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ShipMount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipMount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipMount) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipMount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipMount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipMount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShipMount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStrength

`func (o *ShipMount) GetStrength() int32`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *ShipMount) GetStrengthOk() (*int32, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *ShipMount) SetStrength(v int32)`

SetStrength sets Strength field to given value.

### HasStrength

`func (o *ShipMount) HasStrength() bool`

HasStrength returns a boolean if a field has been set.

### GetDeposits

`func (o *ShipMount) GetDeposits() []string`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *ShipMount) GetDepositsOk() (*[]string, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *ShipMount) SetDeposits(v []string)`

SetDeposits sets Deposits field to given value.

### HasDeposits

`func (o *ShipMount) HasDeposits() bool`

HasDeposits returns a boolean if a field has been set.

### GetRequirements

`func (o *ShipMount) GetRequirements() ShipRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ShipMount) GetRequirementsOk() (*ShipRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ShipMount) SetRequirements(v ShipRequirements)`

SetRequirements sets Requirements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


