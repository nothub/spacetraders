# Shipyard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the shipyard. The symbol is the same as the waypoint where the shipyard is located. | 
**ShipTypes** | [**[]ShipyardShipTypesInner**](ShipyardShipTypesInner.md) | The list of ship types available for purchase at this shipyard. | 
**Transactions** | Pointer to [**[]ShipyardTransaction**](ShipyardTransaction.md) | The list of recent transactions at this shipyard. | [optional] 
**Ships** | Pointer to [**[]ShipyardShip**](ShipyardShip.md) | The ships that are currently available for purchase at the shipyard. | [optional] 
**ModificationsFee** | **int32** | The fee to modify a ship at this shipyard. This includes installing or removing modules and mounts on a ship. In the case of mounts, the fee is a flat rate per mount. In the case of modules, the fee is per slot the module occupies. | 

## Methods

### NewShipyard

`func NewShipyard(symbol string, shipTypes []ShipyardShipTypesInner, modificationsFee int32, ) *Shipyard`

NewShipyard instantiates a new Shipyard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipyardWithDefaults

`func NewShipyardWithDefaults() *Shipyard`

NewShipyardWithDefaults instantiates a new Shipyard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Shipyard) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Shipyard) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Shipyard) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetShipTypes

`func (o *Shipyard) GetShipTypes() []ShipyardShipTypesInner`

GetShipTypes returns the ShipTypes field if non-nil, zero value otherwise.

### GetShipTypesOk

`func (o *Shipyard) GetShipTypesOk() (*[]ShipyardShipTypesInner, bool)`

GetShipTypesOk returns a tuple with the ShipTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipTypes

`func (o *Shipyard) SetShipTypes(v []ShipyardShipTypesInner)`

SetShipTypes sets ShipTypes field to given value.


### GetTransactions

`func (o *Shipyard) GetTransactions() []ShipyardTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Shipyard) GetTransactionsOk() (*[]ShipyardTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Shipyard) SetTransactions(v []ShipyardTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Shipyard) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetShips

`func (o *Shipyard) GetShips() []ShipyardShip`

GetShips returns the Ships field if non-nil, zero value otherwise.

### GetShipsOk

`func (o *Shipyard) GetShipsOk() (*[]ShipyardShip, bool)`

GetShipsOk returns a tuple with the Ships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShips

`func (o *Shipyard) SetShips(v []ShipyardShip)`

SetShips sets Ships field to given value.

### HasShips

`func (o *Shipyard) HasShips() bool`

HasShips returns a boolean if a field has been set.

### GetModificationsFee

`func (o *Shipyard) GetModificationsFee() int32`

GetModificationsFee returns the ModificationsFee field if non-nil, zero value otherwise.

### GetModificationsFeeOk

`func (o *Shipyard) GetModificationsFeeOk() (*int32, bool)`

GetModificationsFeeOk returns a tuple with the ModificationsFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationsFee

`func (o *Shipyard) SetModificationsFee(v int32)`

SetModificationsFee sets ModificationsFee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


