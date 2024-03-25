# ContractTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | **time.Time** | The deadline for the contract. | 
**Payment** | [**ContractPayment**](ContractPayment.md) |  | 
**Deliver** | Pointer to [**[]ContractDeliverGood**](ContractDeliverGood.md) | The cargo that needs to be delivered to fulfill the contract. | [optional] 

## Methods

### NewContractTerms

`func NewContractTerms(deadline time.Time, payment ContractPayment, ) *ContractTerms`

NewContractTerms instantiates a new ContractTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractTermsWithDefaults

`func NewContractTermsWithDefaults() *ContractTerms`

NewContractTermsWithDefaults instantiates a new ContractTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadline

`func (o *ContractTerms) GetDeadline() time.Time`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *ContractTerms) GetDeadlineOk() (*time.Time, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *ContractTerms) SetDeadline(v time.Time)`

SetDeadline sets Deadline field to given value.


### GetPayment

`func (o *ContractTerms) GetPayment() ContractPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *ContractTerms) GetPaymentOk() (*ContractPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *ContractTerms) SetPayment(v ContractPayment)`

SetPayment sets Payment field to given value.


### GetDeliver

`func (o *ContractTerms) GetDeliver() []ContractDeliverGood`

GetDeliver returns the Deliver field if non-nil, zero value otherwise.

### GetDeliverOk

`func (o *ContractTerms) GetDeliverOk() (*[]ContractDeliverGood, bool)`

GetDeliverOk returns a tuple with the Deliver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliver

`func (o *ContractTerms) SetDeliver(v []ContractDeliverGood)`

SetDeliver sets Deliver field to given value.

### HasDeliver

`func (o *ContractTerms) HasDeliver() bool`

HasDeliver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


