# ContractDeliverGood

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeSymbol** | **string** | The symbol of the trade good to deliver. | 
**DestinationSymbol** | **string** | The destination where goods need to be delivered. | 
**UnitsRequired** | **int32** | The number of units that need to be delivered on this contract. | 
**UnitsFulfilled** | **int32** | The number of units fulfilled on this contract. | 

## Methods

### NewContractDeliverGood

`func NewContractDeliverGood(tradeSymbol string, destinationSymbol string, unitsRequired int32, unitsFulfilled int32, ) *ContractDeliverGood`

NewContractDeliverGood instantiates a new ContractDeliverGood object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractDeliverGoodWithDefaults

`func NewContractDeliverGoodWithDefaults() *ContractDeliverGood`

NewContractDeliverGoodWithDefaults instantiates a new ContractDeliverGood object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeSymbol

`func (o *ContractDeliverGood) GetTradeSymbol() string`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *ContractDeliverGood) GetTradeSymbolOk() (*string, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *ContractDeliverGood) SetTradeSymbol(v string)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetDestinationSymbol

`func (o *ContractDeliverGood) GetDestinationSymbol() string`

GetDestinationSymbol returns the DestinationSymbol field if non-nil, zero value otherwise.

### GetDestinationSymbolOk

`func (o *ContractDeliverGood) GetDestinationSymbolOk() (*string, bool)`

GetDestinationSymbolOk returns a tuple with the DestinationSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSymbol

`func (o *ContractDeliverGood) SetDestinationSymbol(v string)`

SetDestinationSymbol sets DestinationSymbol field to given value.


### GetUnitsRequired

`func (o *ContractDeliverGood) GetUnitsRequired() int32`

GetUnitsRequired returns the UnitsRequired field if non-nil, zero value otherwise.

### GetUnitsRequiredOk

`func (o *ContractDeliverGood) GetUnitsRequiredOk() (*int32, bool)`

GetUnitsRequiredOk returns a tuple with the UnitsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsRequired

`func (o *ContractDeliverGood) SetUnitsRequired(v int32)`

SetUnitsRequired sets UnitsRequired field to given value.


### GetUnitsFulfilled

`func (o *ContractDeliverGood) GetUnitsFulfilled() int32`

GetUnitsFulfilled returns the UnitsFulfilled field if non-nil, zero value otherwise.

### GetUnitsFulfilledOk

`func (o *ContractDeliverGood) GetUnitsFulfilledOk() (*int32, bool)`

GetUnitsFulfilledOk returns a tuple with the UnitsFulfilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsFulfilled

`func (o *ContractDeliverGood) SetUnitsFulfilled(v int32)`

SetUnitsFulfilled sets UnitsFulfilled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


