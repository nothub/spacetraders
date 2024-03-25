# Faction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**FactionSymbol**](FactionSymbol.md) |  | 
**Name** | **string** | Name of the faction. | 
**Description** | **string** | Description of the faction. | 
**Headquarters** | **string** | The waypoint in which the faction&#39;s HQ is located in. | 
**Traits** | [**[]FactionTrait**](FactionTrait.md) | List of traits that define this faction. | 
**IsRecruiting** | **bool** | Whether or not the faction is currently recruiting new agents. | 

## Methods

### NewFaction

`func NewFaction(symbol FactionSymbol, name string, description string, headquarters string, traits []FactionTrait, isRecruiting bool, ) *Faction`

NewFaction instantiates a new Faction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactionWithDefaults

`func NewFactionWithDefaults() *Faction`

NewFactionWithDefaults instantiates a new Faction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Faction) GetSymbol() FactionSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Faction) GetSymbolOk() (*FactionSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Faction) SetSymbol(v FactionSymbol)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *Faction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Faction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Faction) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Faction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Faction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Faction) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHeadquarters

`func (o *Faction) GetHeadquarters() string`

GetHeadquarters returns the Headquarters field if non-nil, zero value otherwise.

### GetHeadquartersOk

`func (o *Faction) GetHeadquartersOk() (*string, bool)`

GetHeadquartersOk returns a tuple with the Headquarters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadquarters

`func (o *Faction) SetHeadquarters(v string)`

SetHeadquarters sets Headquarters field to given value.


### GetTraits

`func (o *Faction) GetTraits() []FactionTrait`

GetTraits returns the Traits field if non-nil, zero value otherwise.

### GetTraitsOk

`func (o *Faction) GetTraitsOk() (*[]FactionTrait, bool)`

GetTraitsOk returns a tuple with the Traits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraits

`func (o *Faction) SetTraits(v []FactionTrait)`

SetTraits sets Traits field to given value.


### GetIsRecruiting

`func (o *Faction) GetIsRecruiting() bool`

GetIsRecruiting returns the IsRecruiting field if non-nil, zero value otherwise.

### GetIsRecruitingOk

`func (o *Faction) GetIsRecruitingOk() (*bool, bool)`

GetIsRecruitingOk returns a tuple with the IsRecruiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecruiting

`func (o *Faction) SetIsRecruiting(v bool)`

SetIsRecruiting sets IsRecruiting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


