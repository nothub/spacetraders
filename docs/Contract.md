# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the contract. | 
**FactionSymbol** | **string** | The symbol of the faction that this contract is for. | 
**Type** | **string** | Type of contract. | 
**Terms** | [**ContractTerms**](ContractTerms.md) |  | 
**Accepted** | **bool** | Whether the contract has been accepted by the agent | [default to false]
**Fulfilled** | **bool** | Whether the contract has been fulfilled | [default to false]
**Expiration** | **time.Time** | Deprecated in favor of deadlineToAccept | 
**DeadlineToAccept** | Pointer to **time.Time** | The time at which the contract is no longer available to be accepted | [optional] 

## Methods

### NewContract

`func NewContract(id string, factionSymbol string, type_ string, terms ContractTerms, accepted bool, fulfilled bool, expiration time.Time, ) *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contract) SetId(v string)`

SetId sets Id field to given value.


### GetFactionSymbol

`func (o *Contract) GetFactionSymbol() string`

GetFactionSymbol returns the FactionSymbol field if non-nil, zero value otherwise.

### GetFactionSymbolOk

`func (o *Contract) GetFactionSymbolOk() (*string, bool)`

GetFactionSymbolOk returns a tuple with the FactionSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionSymbol

`func (o *Contract) SetFactionSymbol(v string)`

SetFactionSymbol sets FactionSymbol field to given value.


### GetType

`func (o *Contract) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contract) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contract) SetType(v string)`

SetType sets Type field to given value.


### GetTerms

`func (o *Contract) GetTerms() ContractTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Contract) GetTermsOk() (*ContractTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Contract) SetTerms(v ContractTerms)`

SetTerms sets Terms field to given value.


### GetAccepted

`func (o *Contract) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *Contract) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *Contract) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.


### GetFulfilled

`func (o *Contract) GetFulfilled() bool`

GetFulfilled returns the Fulfilled field if non-nil, zero value otherwise.

### GetFulfilledOk

`func (o *Contract) GetFulfilledOk() (*bool, bool)`

GetFulfilledOk returns a tuple with the Fulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfilled

`func (o *Contract) SetFulfilled(v bool)`

SetFulfilled sets Fulfilled field to given value.


### GetExpiration

`func (o *Contract) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *Contract) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *Contract) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### GetDeadlineToAccept

`func (o *Contract) GetDeadlineToAccept() time.Time`

GetDeadlineToAccept returns the DeadlineToAccept field if non-nil, zero value otherwise.

### GetDeadlineToAcceptOk

`func (o *Contract) GetDeadlineToAcceptOk() (*time.Time, bool)`

GetDeadlineToAcceptOk returns a tuple with the DeadlineToAccept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineToAccept

`func (o *Contract) SetDeadlineToAccept(v time.Time)`

SetDeadlineToAccept sets DeadlineToAccept field to given value.

### HasDeadlineToAccept

`func (o *Contract) HasDeadlineToAccept() bool`

HasDeadlineToAccept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


