# DeliverContractRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipSymbol** | **string** | Symbol of a ship located in the destination to deliver a contract and that has a good to deliver in its cargo. | 
**TradeSymbol** | **string** | The symbol of the good to deliver. | 
**Units** | **int32** | Amount of units to deliver. | 

## Methods

### NewDeliverContractRequest

`func NewDeliverContractRequest(shipSymbol string, tradeSymbol string, units int32, ) *DeliverContractRequest`

NewDeliverContractRequest instantiates a new DeliverContractRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverContractRequestWithDefaults

`func NewDeliverContractRequestWithDefaults() *DeliverContractRequest`

NewDeliverContractRequestWithDefaults instantiates a new DeliverContractRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipSymbol

`func (o *DeliverContractRequest) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *DeliverContractRequest) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *DeliverContractRequest) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.


### GetTradeSymbol

`func (o *DeliverContractRequest) GetTradeSymbol() string`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *DeliverContractRequest) GetTradeSymbolOk() (*string, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *DeliverContractRequest) SetTradeSymbol(v string)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetUnits

`func (o *DeliverContractRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *DeliverContractRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *DeliverContractRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


