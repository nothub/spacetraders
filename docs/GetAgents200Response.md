# GetAgents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Agent**](Agent.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetAgents200Response

`func NewGetAgents200Response(data []Agent, meta Meta, ) *GetAgents200Response`

NewGetAgents200Response instantiates a new GetAgents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAgents200ResponseWithDefaults

`func NewGetAgents200ResponseWithDefaults() *GetAgents200Response`

NewGetAgents200ResponseWithDefaults instantiates a new GetAgents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAgents200Response) GetData() []Agent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAgents200Response) GetDataOk() (*[]Agent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAgents200Response) SetData(v []Agent)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetAgents200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAgents200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAgents200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


