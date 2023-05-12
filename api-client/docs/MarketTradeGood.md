# MarketTradeGood

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the trade good. | 
**TradeVolume** | **int32** | The typical volume flowing through the market for this type of good. The larger the trade volume, the more stable prices will be. | 
**Supply** | **string** | A rough estimate of the total supply of this good in the marketplace. | 
**PurchasePrice** | **int32** | The price at which this good can be purchased from the market. | 
**SellPrice** | **int32** | The price at which this good can be sold to the market. | 

## Methods

### NewMarketTradeGood

`func NewMarketTradeGood(symbol string, tradeVolume int32, supply string, purchasePrice int32, sellPrice int32, ) *MarketTradeGood`

NewMarketTradeGood instantiates a new MarketTradeGood object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTradeGoodWithDefaults

`func NewMarketTradeGoodWithDefaults() *MarketTradeGood`

NewMarketTradeGoodWithDefaults instantiates a new MarketTradeGood object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MarketTradeGood) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarketTradeGood) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarketTradeGood) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTradeVolume

`func (o *MarketTradeGood) GetTradeVolume() int32`

GetTradeVolume returns the TradeVolume field if non-nil, zero value otherwise.

### GetTradeVolumeOk

`func (o *MarketTradeGood) GetTradeVolumeOk() (*int32, bool)`

GetTradeVolumeOk returns a tuple with the TradeVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVolume

`func (o *MarketTradeGood) SetTradeVolume(v int32)`

SetTradeVolume sets TradeVolume field to given value.


### GetSupply

`func (o *MarketTradeGood) GetSupply() string`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *MarketTradeGood) GetSupplyOk() (*string, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *MarketTradeGood) SetSupply(v string)`

SetSupply sets Supply field to given value.


### GetPurchasePrice

`func (o *MarketTradeGood) GetPurchasePrice() int32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *MarketTradeGood) GetPurchasePriceOk() (*int32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *MarketTradeGood) SetPurchasePrice(v int32)`

SetPurchasePrice sets PurchasePrice field to given value.


### GetSellPrice

`func (o *MarketTradeGood) GetSellPrice() int32`

GetSellPrice returns the SellPrice field if non-nil, zero value otherwise.

### GetSellPriceOk

`func (o *MarketTradeGood) GetSellPriceOk() (*int32, bool)`

GetSellPriceOk returns a tuple with the SellPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellPrice

`func (o *MarketTradeGood) SetSellPrice(v int32)`

SetSellPrice sets SellPrice field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


