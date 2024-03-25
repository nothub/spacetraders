# TransferCargoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TradeSymbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Units** | **int32** | Amount of units to transfer. | 
**ShipSymbol** | **string** | The symbol of the ship to transfer to. | 

## Methods

### NewTransferCargoRequest

`func NewTransferCargoRequest(tradeSymbol TradeSymbol, units int32, shipSymbol string, ) *TransferCargoRequest`

NewTransferCargoRequest instantiates a new TransferCargoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferCargoRequestWithDefaults

`func NewTransferCargoRequestWithDefaults() *TransferCargoRequest`

NewTransferCargoRequestWithDefaults instantiates a new TransferCargoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTradeSymbol

`func (o *TransferCargoRequest) GetTradeSymbol() TradeSymbol`

GetTradeSymbol returns the TradeSymbol field if non-nil, zero value otherwise.

### GetTradeSymbolOk

`func (o *TransferCargoRequest) GetTradeSymbolOk() (*TradeSymbol, bool)`

GetTradeSymbolOk returns a tuple with the TradeSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeSymbol

`func (o *TransferCargoRequest) SetTradeSymbol(v TradeSymbol)`

SetTradeSymbol sets TradeSymbol field to given value.


### GetUnits

`func (o *TransferCargoRequest) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *TransferCargoRequest) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *TransferCargoRequest) SetUnits(v int32)`

SetUnits sets Units field to given value.


### GetShipSymbol

`func (o *TransferCargoRequest) GetShipSymbol() string`

GetShipSymbol returns the ShipSymbol field if non-nil, zero value otherwise.

### GetShipSymbolOk

`func (o *TransferCargoRequest) GetShipSymbolOk() (*string, bool)`

GetShipSymbolOk returns a tuple with the ShipSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSymbol

`func (o *TransferCargoRequest) SetShipSymbol(v string)`

SetShipSymbol sets ShipSymbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


