# GetStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The current status of the game server. | 
**Version** | **string** | The current version of the API. | 
**ResetDate** | **string** | The date when the game server was last reset. | 
**Description** | **string** |  | 
**Stats** | [**GetStatus200ResponseStats**](GetStatus200ResponseStats.md) |  | 
**Leaderboards** | [**GetStatus200ResponseLeaderboards**](GetStatus200ResponseLeaderboards.md) |  | 
**ServerResets** | [**GetStatus200ResponseServerResets**](GetStatus200ResponseServerResets.md) |  | 
**Announcements** | [**[]GetStatus200ResponseAnnouncementsInner**](GetStatus200ResponseAnnouncementsInner.md) |  | 
**Links** | [**[]GetStatus200ResponseLinksInner**](GetStatus200ResponseLinksInner.md) |  | 

## Methods

### NewGetStatus200Response

`func NewGetStatus200Response(status string, version string, resetDate string, description string, stats GetStatus200ResponseStats, leaderboards GetStatus200ResponseLeaderboards, serverResets GetStatus200ResponseServerResets, announcements []GetStatus200ResponseAnnouncementsInner, links []GetStatus200ResponseLinksInner, ) *GetStatus200Response`

NewGetStatus200Response instantiates a new GetStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatus200ResponseWithDefaults

`func NewGetStatus200ResponseWithDefaults() *GetStatus200Response`

NewGetStatus200ResponseWithDefaults instantiates a new GetStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVersion

`func (o *GetStatus200Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetStatus200Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetStatus200Response) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetResetDate

`func (o *GetStatus200Response) GetResetDate() string`

GetResetDate returns the ResetDate field if non-nil, zero value otherwise.

### GetResetDateOk

`func (o *GetStatus200Response) GetResetDateOk() (*string, bool)`

GetResetDateOk returns a tuple with the ResetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetDate

`func (o *GetStatus200Response) SetResetDate(v string)`

SetResetDate sets ResetDate field to given value.


### GetDescription

`func (o *GetStatus200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetStatus200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetStatus200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStats

`func (o *GetStatus200Response) GetStats() GetStatus200ResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetStatus200Response) GetStatsOk() (*GetStatus200ResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetStatus200Response) SetStats(v GetStatus200ResponseStats)`

SetStats sets Stats field to given value.


### GetLeaderboards

`func (o *GetStatus200Response) GetLeaderboards() GetStatus200ResponseLeaderboards`

GetLeaderboards returns the Leaderboards field if non-nil, zero value otherwise.

### GetLeaderboardsOk

`func (o *GetStatus200Response) GetLeaderboardsOk() (*GetStatus200ResponseLeaderboards, bool)`

GetLeaderboardsOk returns a tuple with the Leaderboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboards

`func (o *GetStatus200Response) SetLeaderboards(v GetStatus200ResponseLeaderboards)`

SetLeaderboards sets Leaderboards field to given value.


### GetServerResets

`func (o *GetStatus200Response) GetServerResets() GetStatus200ResponseServerResets`

GetServerResets returns the ServerResets field if non-nil, zero value otherwise.

### GetServerResetsOk

`func (o *GetStatus200Response) GetServerResetsOk() (*GetStatus200ResponseServerResets, bool)`

GetServerResetsOk returns a tuple with the ServerResets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerResets

`func (o *GetStatus200Response) SetServerResets(v GetStatus200ResponseServerResets)`

SetServerResets sets ServerResets field to given value.


### GetAnnouncements

`func (o *GetStatus200Response) GetAnnouncements() []GetStatus200ResponseAnnouncementsInner`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *GetStatus200Response) GetAnnouncementsOk() (*[]GetStatus200ResponseAnnouncementsInner, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *GetStatus200Response) SetAnnouncements(v []GetStatus200ResponseAnnouncementsInner)`

SetAnnouncements sets Announcements field to given value.


### GetLinks

`func (o *GetStatus200Response) GetLinks() []GetStatus200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetStatus200Response) GetLinksOk() (*[]GetStatus200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetStatus200Response) SetLinks(v []GetStatus200ResponseLinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


