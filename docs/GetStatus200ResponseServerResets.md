# GetStatus200ResponseServerResets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | **string** | The date and time when the game server will reset. | 
**Frequency** | **string** | How often we intend to reset the game server. | 

## Methods

### NewGetStatus200ResponseServerResets

`func NewGetStatus200ResponseServerResets(next string, frequency string, ) *GetStatus200ResponseServerResets`

NewGetStatus200ResponseServerResets instantiates a new GetStatus200ResponseServerResets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatus200ResponseServerResetsWithDefaults

`func NewGetStatus200ResponseServerResetsWithDefaults() *GetStatus200ResponseServerResets`

NewGetStatus200ResponseServerResetsWithDefaults instantiates a new GetStatus200ResponseServerResets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *GetStatus200ResponseServerResets) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetStatus200ResponseServerResets) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetStatus200ResponseServerResets) SetNext(v string)`

SetNext sets Next field to given value.


### GetFrequency

`func (o *GetStatus200ResponseServerResets) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetStatus200ResponseServerResets) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetStatus200ResponseServerResets) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


