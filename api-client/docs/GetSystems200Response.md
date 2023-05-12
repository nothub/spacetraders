# GetSystems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]System**](System.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetSystems200Response

`func NewGetSystems200Response(data []System, meta Meta, ) *GetSystems200Response`

NewGetSystems200Response instantiates a new GetSystems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSystems200ResponseWithDefaults

`func NewGetSystems200ResponseWithDefaults() *GetSystems200Response`

NewGetSystems200ResponseWithDefaults instantiates a new GetSystems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSystems200Response) GetData() []System`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSystems200Response) GetDataOk() (*[]System, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSystems200Response) SetData(v []System)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSystems200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSystems200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSystems200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


