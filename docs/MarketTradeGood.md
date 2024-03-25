# MarketTradeGood

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Type** | **string** | The type of trade good (export, import, or exchange). | 
**TradeVolume** | **int32** | This is the maximum number of units that can be purchased or sold at this market in a single trade for this good. Trade volume also gives an indication of price volatility. A market with a low trade volume will have large price swings, while high trade volume will be more resilient to price changes. | 
**Supply** | [**SupplyLevel**](SupplyLevel.md) |  | 
**Activity** | Pointer to [**ActivityLevel**](ActivityLevel.md) |  | [optional] 
**PurchasePrice** | **int32** | The price at which this good can be purchased from the market. | 
**SellPrice** | **int32** | The price at which this good can be sold to the market. | 

## Methods

### NewMarketTradeGood

`func NewMarketTradeGood(symbol TradeSymbol, type_ string, tradeVolume int32, supply SupplyLevel, purchasePrice int32, sellPrice int32, ) *MarketTradeGood`

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

`func (o *MarketTradeGood) GetSymbol() TradeSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarketTradeGood) GetSymbolOk() (*TradeSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarketTradeGood) SetSymbol(v TradeSymbol)`

SetSymbol sets Symbol field to given value.


### GetType

`func (o *MarketTradeGood) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketTradeGood) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketTradeGood) SetType(v string)`

SetType sets Type field to given value.


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

`func (o *MarketTradeGood) GetSupply() SupplyLevel`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *MarketTradeGood) GetSupplyOk() (*SupplyLevel, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *MarketTradeGood) SetSupply(v SupplyLevel)`

SetSupply sets Supply field to given value.


### GetActivity

`func (o *MarketTradeGood) GetActivity() ActivityLevel`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MarketTradeGood) GetActivityOk() (*ActivityLevel, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *MarketTradeGood) SetActivity(v ActivityLevel)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *MarketTradeGood) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

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


