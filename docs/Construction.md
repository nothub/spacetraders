# Construction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the waypoint. | 
**Materials** | [**[]ConstructionMaterial**](ConstructionMaterial.md) | The materials required to construct the waypoint. | 
**IsComplete** | **bool** | Whether the waypoint has been constructed. | 

## Methods

### NewConstruction

`func NewConstruction(symbol string, materials []ConstructionMaterial, isComplete bool, ) *Construction`

NewConstruction instantiates a new Construction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstructionWithDefaults

`func NewConstructionWithDefaults() *Construction`

NewConstructionWithDefaults instantiates a new Construction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Construction) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Construction) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Construction) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetMaterials

`func (o *Construction) GetMaterials() []ConstructionMaterial`

GetMaterials returns the Materials field if non-nil, zero value otherwise.

### GetMaterialsOk

`func (o *Construction) GetMaterialsOk() (*[]ConstructionMaterial, bool)`

GetMaterialsOk returns a tuple with the Materials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterials

`func (o *Construction) SetMaterials(v []ConstructionMaterial)`

SetMaterials sets Materials field to given value.


### GetIsComplete

`func (o *Construction) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *Construction) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *Construction) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


