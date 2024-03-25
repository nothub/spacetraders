# CreateSurvey201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Surveys** | [**[]Survey**](Survey.md) | Surveys created by this action. | 

## Methods

### NewCreateSurvey201ResponseData

`func NewCreateSurvey201ResponseData(cooldown Cooldown, surveys []Survey, ) *CreateSurvey201ResponseData`

NewCreateSurvey201ResponseData instantiates a new CreateSurvey201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSurvey201ResponseDataWithDefaults

`func NewCreateSurvey201ResponseDataWithDefaults() *CreateSurvey201ResponseData`

NewCreateSurvey201ResponseDataWithDefaults instantiates a new CreateSurvey201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCooldown

`func (o *CreateSurvey201ResponseData) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *CreateSurvey201ResponseData) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *CreateSurvey201ResponseData) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetSurveys

`func (o *CreateSurvey201ResponseData) GetSurveys() []Survey`

GetSurveys returns the Surveys field if non-nil, zero value otherwise.

### GetSurveysOk

`func (o *CreateSurvey201ResponseData) GetSurveysOk() (*[]Survey, bool)`

GetSurveysOk returns a tuple with the Surveys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveys

`func (o *CreateSurvey201ResponseData) SetSurveys(v []Survey)`

SetSurveys sets Surveys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


