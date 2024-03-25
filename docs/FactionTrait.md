# FactionTrait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**FactionTraitSymbol**](FactionTraitSymbol.md) |  | 
**Name** | **string** | The name of the trait. | 
**Description** | **string** | A description of the trait. | 

## Methods

### NewFactionTrait

`func NewFactionTrait(symbol FactionTraitSymbol, name string, description string, ) *FactionTrait`

NewFactionTrait instantiates a new FactionTrait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactionTraitWithDefaults

`func NewFactionTraitWithDefaults() *FactionTrait`

NewFactionTraitWithDefaults instantiates a new FactionTrait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FactionTrait) GetSymbol() FactionTraitSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FactionTrait) GetSymbolOk() (*FactionTraitSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FactionTrait) SetSymbol(v FactionTraitSymbol)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *FactionTrait) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FactionTrait) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FactionTrait) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FactionTrait) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FactionTrait) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FactionTrait) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


