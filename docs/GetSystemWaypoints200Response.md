# GetSystemWaypoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Waypoint**](Waypoint.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewGetSystemWaypoints200Response

`func NewGetSystemWaypoints200Response(data []Waypoint, meta Meta, ) *GetSystemWaypoints200Response`

NewGetSystemWaypoints200Response instantiates a new GetSystemWaypoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSystemWaypoints200ResponseWithDefaults

`func NewGetSystemWaypoints200ResponseWithDefaults() *GetSystemWaypoints200Response`

NewGetSystemWaypoints200ResponseWithDefaults instantiates a new GetSystemWaypoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSystemWaypoints200Response) GetData() []Waypoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSystemWaypoints200Response) GetDataOk() (*[]Waypoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSystemWaypoints200Response) SetData(v []Waypoint)`

SetData sets Data field to given value.


### GetMeta

`func (o *GetSystemWaypoints200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSystemWaypoints200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSystemWaypoints200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


