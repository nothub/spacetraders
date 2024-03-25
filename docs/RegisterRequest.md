# RegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Faction** | [**FactionSymbol**](FactionSymbol.md) |  | 
**Symbol** | **string** | Your desired agent symbol. This will be a unique name used to represent your agent, and will be the prefix for your ships. | 
**Email** | Pointer to **string** | Your email address. This is used if you reserved your call sign between resets. | [optional] 

## Methods

### NewRegisterRequest

`func NewRegisterRequest(faction FactionSymbol, symbol string, ) *RegisterRequest`

NewRegisterRequest instantiates a new RegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterRequestWithDefaults

`func NewRegisterRequestWithDefaults() *RegisterRequest`

NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaction

`func (o *RegisterRequest) GetFaction() FactionSymbol`

GetFaction returns the Faction field if non-nil, zero value otherwise.

### GetFactionOk

`func (o *RegisterRequest) GetFactionOk() (*FactionSymbol, bool)`

GetFactionOk returns a tuple with the Faction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaction

`func (o *RegisterRequest) SetFaction(v FactionSymbol)`

SetFaction sets Faction field to given value.


### GetSymbol

`func (o *RegisterRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *RegisterRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *RegisterRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetEmail

`func (o *RegisterRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RegisterRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


