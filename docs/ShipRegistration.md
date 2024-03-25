# ShipRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The agent&#39;s registered name of the ship | 
**FactionSymbol** | **string** | The symbol of the faction the ship is registered with | 
**Role** | [**ShipRole**](ShipRole.md) |  | 

## Methods

### NewShipRegistration

`func NewShipRegistration(name string, factionSymbol string, role ShipRole, ) *ShipRegistration`

NewShipRegistration instantiates a new ShipRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipRegistrationWithDefaults

`func NewShipRegistrationWithDefaults() *ShipRegistration`

NewShipRegistrationWithDefaults instantiates a new ShipRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShipRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipRegistration) SetName(v string)`

SetName sets Name field to given value.


### GetFactionSymbol

`func (o *ShipRegistration) GetFactionSymbol() string`

GetFactionSymbol returns the FactionSymbol field if non-nil, zero value otherwise.

### GetFactionSymbolOk

`func (o *ShipRegistration) GetFactionSymbolOk() (*string, bool)`

GetFactionSymbolOk returns a tuple with the FactionSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionSymbol

`func (o *ShipRegistration) SetFactionSymbol(v string)`

SetFactionSymbol sets FactionSymbol field to given value.


### GetRole

`func (o *ShipRegistration) GetRole() ShipRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ShipRegistration) GetRoleOk() (*ShipRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ShipRegistration) SetRole(v ShipRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


