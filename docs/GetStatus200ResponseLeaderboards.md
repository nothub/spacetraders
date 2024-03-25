# GetStatus200ResponseLeaderboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MostCredits** | [**[]GetStatus200ResponseLeaderboardsMostCreditsInner**](GetStatus200ResponseLeaderboardsMostCreditsInner.md) | Top agents with the most credits. | 
**MostSubmittedCharts** | [**[]GetStatus200ResponseLeaderboardsMostSubmittedChartsInner**](GetStatus200ResponseLeaderboardsMostSubmittedChartsInner.md) | Top agents with the most charted submitted. | 

## Methods

### NewGetStatus200ResponseLeaderboards

`func NewGetStatus200ResponseLeaderboards(mostCredits []GetStatus200ResponseLeaderboardsMostCreditsInner, mostSubmittedCharts []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner, ) *GetStatus200ResponseLeaderboards`

NewGetStatus200ResponseLeaderboards instantiates a new GetStatus200ResponseLeaderboards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatus200ResponseLeaderboardsWithDefaults

`func NewGetStatus200ResponseLeaderboardsWithDefaults() *GetStatus200ResponseLeaderboards`

NewGetStatus200ResponseLeaderboardsWithDefaults instantiates a new GetStatus200ResponseLeaderboards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMostCredits

`func (o *GetStatus200ResponseLeaderboards) GetMostCredits() []GetStatus200ResponseLeaderboardsMostCreditsInner`

GetMostCredits returns the MostCredits field if non-nil, zero value otherwise.

### GetMostCreditsOk

`func (o *GetStatus200ResponseLeaderboards) GetMostCreditsOk() (*[]GetStatus200ResponseLeaderboardsMostCreditsInner, bool)`

GetMostCreditsOk returns a tuple with the MostCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostCredits

`func (o *GetStatus200ResponseLeaderboards) SetMostCredits(v []GetStatus200ResponseLeaderboardsMostCreditsInner)`

SetMostCredits sets MostCredits field to given value.


### GetMostSubmittedCharts

`func (o *GetStatus200ResponseLeaderboards) GetMostSubmittedCharts() []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner`

GetMostSubmittedCharts returns the MostSubmittedCharts field if non-nil, zero value otherwise.

### GetMostSubmittedChartsOk

`func (o *GetStatus200ResponseLeaderboards) GetMostSubmittedChartsOk() (*[]GetStatus200ResponseLeaderboardsMostSubmittedChartsInner, bool)`

GetMostSubmittedChartsOk returns a tuple with the MostSubmittedCharts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostSubmittedCharts

`func (o *GetStatus200ResponseLeaderboards) SetMostSubmittedCharts(v []GetStatus200ResponseLeaderboardsMostSubmittedChartsInner)`

SetMostSubmittedCharts sets MostSubmittedCharts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


