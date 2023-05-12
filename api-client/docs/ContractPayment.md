# ContractPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnAccepted** | **int32** | The amount of credits received up front for accepting the contract. | 
**OnFulfilled** | **int32** | The amount of credits received when the contract is fulfilled. | 

## Methods

### NewContractPayment

`func NewContractPayment(onAccepted int32, onFulfilled int32, ) *ContractPayment`

NewContractPayment instantiates a new ContractPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractPaymentWithDefaults

`func NewContractPaymentWithDefaults() *ContractPayment`

NewContractPaymentWithDefaults instantiates a new ContractPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnAccepted

`func (o *ContractPayment) GetOnAccepted() int32`

GetOnAccepted returns the OnAccepted field if non-nil, zero value otherwise.

### GetOnAcceptedOk

`func (o *ContractPayment) GetOnAcceptedOk() (*int32, bool)`

GetOnAcceptedOk returns a tuple with the OnAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnAccepted

`func (o *ContractPayment) SetOnAccepted(v int32)`

SetOnAccepted sets OnAccepted field to given value.


### GetOnFulfilled

`func (o *ContractPayment) GetOnFulfilled() int32`

GetOnFulfilled returns the OnFulfilled field if non-nil, zero value otherwise.

### GetOnFulfilledOk

`func (o *ContractPayment) GetOnFulfilledOk() (*int32, bool)`

GetOnFulfilledOk returns a tuple with the OnFulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFulfilled

`func (o *ContractPayment) SetOnFulfilled(v int32)`

SetOnFulfilled sets OnFulfilled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


