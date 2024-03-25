# ShipCargoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | [**TradeSymbol**](TradeSymbol.md) |  | 
**Name** | **string** | The name of the cargo item type. | 
**Description** | **string** | The description of the cargo item type. | 
**Units** | **int32** | The number of units of the cargo item. | 

## Methods

### NewShipCargoItem

`func NewShipCargoItem(symbol TradeSymbol, name string, description string, units int32, ) *ShipCargoItem`

NewShipCargoItem instantiates a new ShipCargoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipCargoItemWithDefaults

`func NewShipCargoItemWithDefaults() *ShipCargoItem`

NewShipCargoItemWithDefaults instantiates a new ShipCargoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ShipCargoItem) GetSymbol() TradeSymbol`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ShipCargoItem) GetSymbolOk() (*TradeSymbol, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ShipCargoItem) SetSymbol(v TradeSymbol)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *ShipCargoItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipCargoItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipCargoItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipCargoItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipCargoItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipCargoItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnits

`func (o *ShipCargoItem) GetUnits() int32`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ShipCargoItem) GetUnitsOk() (*int32, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ShipCargoItem) SetUnits(v int32)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


