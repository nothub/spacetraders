# Survey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | **string** | A unique signature for the location of this survey. This signature is verified when attempting an extraction using this survey. | 
**Symbol** | **string** | The symbol of the waypoint that this survey is for. | 
**Deposits** | [**[]SurveyDeposit**](SurveyDeposit.md) | A list of deposits that can be found at this location. A ship will extract one of these deposits when using this survey in an extraction request. If multiple deposits of the same type are present, the chance of extracting that deposit is increased. | 
**Expiration** | **time.Time** | The date and time when the survey expires. After this date and time, the survey will no longer be available for extraction. | 
**Size** | **string** | The size of the deposit. This value indicates how much can be extracted from the survey before it is exhausted. | 

## Methods

### NewSurvey

`func NewSurvey(signature string, symbol string, deposits []SurveyDeposit, expiration time.Time, size string, ) *Survey`

NewSurvey instantiates a new Survey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyWithDefaults

`func NewSurveyWithDefaults() *Survey`

NewSurveyWithDefaults instantiates a new Survey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *Survey) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Survey) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Survey) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetSymbol

`func (o *Survey) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Survey) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Survey) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDeposits

`func (o *Survey) GetDeposits() []SurveyDeposit`

GetDeposits returns the Deposits field if non-nil, zero value otherwise.

### GetDepositsOk

`func (o *Survey) GetDepositsOk() (*[]SurveyDeposit, bool)`

GetDepositsOk returns a tuple with the Deposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposits

`func (o *Survey) SetDeposits(v []SurveyDeposit)`

SetDeposits sets Deposits field to given value.


### GetExpiration

`func (o *Survey) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Survey) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Survey) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### GetSize

`func (o *Survey) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Survey) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Survey) SetSize(v string)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


