# GetFactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Faction**](Faction.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetFactions200Response

`func NewGetFactions200Response(data []Faction, meta Meta, ) *GetFactions200Response`

NewGetFactions200Response instantiates a new GetFactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFactions200ResponseWithDefaults

`func NewGetFactions200ResponseWithDefaults() *GetFactions200Response`

NewGetFactions200ResponseWithDefaults instantiates a new GetFactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetFactions200Response) GetData() []Faction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetFactions200Response) GetDataOk() (*[]Faction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetFactions200Response) SetData(v []Faction)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetFactions200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetFactions200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetFactions200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


