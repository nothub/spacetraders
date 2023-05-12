# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Symbol** | **string** |  | 
**Headquarters** | **string** | The headquarters of the agent. | 
**Credits** | **int32** | The number of credits the agent has available. Credits can be negative if funds have been overdrawn. | 

## Methods

### NewAgent

`func NewAgent(accountId string, symbol string, headquarters string, credits int32, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Agent) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Agent) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Agent) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSymbol

`func (o *Agent) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Agent) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Agent) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetHeadquarters

`func (o *Agent) GetHeadquarters() string`

GetHeadquarters returns the Headquarters field if non-nil, zero value otherwise.

### GetHeadquartersOk

`func (o *Agent) GetHeadquartersOk() (*string, bool)`

GetHeadquartersOk returns a tuple with the Headquarters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadquarters

`func (o *Agent) SetHeadquarters(v string)`

SetHeadquarters sets Headquarters field to given value.


### GetCredits

`func (o *Agent) GetCredits() int32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Agent) GetCreditsOk() (*int32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Agent) SetCredits(v int32)`

SetCredits sets Credits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


