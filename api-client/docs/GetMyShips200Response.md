# GetMyShips200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Ship**](Ship.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetMyShips200Response

`func NewGetMyShips200Response(data []Ship, meta Meta, ) *GetMyShips200Response`

NewGetMyShips200Response instantiates a new GetMyShips200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyShips200ResponseWithDefaults

`func NewGetMyShips200ResponseWithDefaults() *GetMyShips200Response`

NewGetMyShips200ResponseWithDefaults instantiates a new GetMyShips200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetMyShips200Response) GetData() []Ship`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetMyShips200Response) GetDataOk() (*[]Ship, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetMyShips200Response) SetData(v []Ship)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetMyShips200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMyShips200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMyShips200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


