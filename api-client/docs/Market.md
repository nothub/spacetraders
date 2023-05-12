# Market

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the market. The symbol is the same as the waypoint where the market is located. | 
**Exports** | [**[]TradeGood**](TradeGood.md) | The list of goods that are exported from this market. | 
**Imports** | [**[]TradeGood**](TradeGood.md) | The list of goods that are sought as imports in this market. | 
**Exchange** | [**[]TradeGood**](TradeGood.md) | The list of goods that are bought and sold between agents at this market. | 
**Transactions** | Pointer to [**[]MarketTransaction**](MarketTransaction.md) | The list of recent transactions at this market. Visible only when a ship is present at the market. | [optional] 
**TradeGoods** | Pointer to [**[]MarketTradeGood**](MarketTradeGood.md) | The list of goods that are traded at this market. Visible only when a ship is present at the market. | [optional] 

## Methods

### NewMarket

`func NewMarket(symbol string, exports []TradeGood, imports []TradeGood, exchange []TradeGood, ) *Market`

NewMarket instantiates a new Market object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketWithDefaults

`func NewMarketWithDefaults() *Market`

NewMarketWithDefaults instantiates a new Market object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Market) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Market) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Market) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetExports

`func (o *Market) GetExports() []TradeGood`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *Market) GetExportsOk() (*[]TradeGood, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *Market) SetExports(v []TradeGood)`

SetExports sets Exports field to given value.


### GetImports

`func (o *Market) GetImports() []TradeGood`

GetImports returns the Imports field if non-nil, zero value otherwise.

### GetImportsOk

`func (o *Market) GetImportsOk() (*[]TradeGood, bool)`

GetImportsOk returns a tuple with the Imports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImports

`func (o *Market) SetImports(v []TradeGood)`

SetImports sets Imports field to given value.


### GetExchange

`func (o *Market) GetExchange() []TradeGood`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Market) GetExchangeOk() (*[]TradeGood, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Market) SetExchange(v []TradeGood)`

SetExchange sets Exchange field to given value.


### GetTransactions

`func (o *Market) GetTransactions() []MarketTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Market) GetTransactionsOk() (*[]MarketTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Market) SetTransactions(v []MarketTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Market) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetTradeGoods

`func (o *Market) GetTradeGoods() []MarketTradeGood`

GetTradeGoods returns the TradeGoods field if non-nil, zero value otherwise.

### GetTradeGoodsOk

`func (o *Market) GetTradeGoodsOk() (*[]MarketTradeGood, bool)`

GetTradeGoodsOk returns a tuple with the TradeGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeGoods

`func (o *Market) SetTradeGoods(v []MarketTradeGood)`

SetTradeGoods sets TradeGoods field to given value.

### HasTradeGoods

`func (o *Market) HasTradeGoods() bool`

HasTradeGoods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


