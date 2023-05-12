# GetContracts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Contract**](Contract.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetContracts200Response

`func NewGetContracts200Response(data []Contract, meta Meta, ) *GetContracts200Response`

NewGetContracts200Response instantiates a new GetContracts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContracts200ResponseWithDefaults

`func NewGetContracts200ResponseWithDefaults() *GetContracts200Response`

NewGetContracts200ResponseWithDefaults instantiates a new GetContracts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetContracts200Response) GetData() []Contract`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetContracts200Response) GetDataOk() (*[]Contract, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetContracts200Response) SetData(v []Contract)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetContracts200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetContracts200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetContracts200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


