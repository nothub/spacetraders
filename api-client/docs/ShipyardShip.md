# ShipyardShip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ShipType**](ShipType.md) |  | [optional] 
**Name** | **string** |  | 
**Description** | **string** |  | 
**PurchasePrice** | **int32** |  | 
**Frame** | [**ShipFrame**](ShipFrame.md) |  | 
**Reactor** | [**ShipReactor**](ShipReactor.md) |  | 
**Engine** | [**ShipEngine**](ShipEngine.md) |  | 
**Modules** | [**[]ShipModule**](ShipModule.md) |  | 
**Mounts** | [**[]ShipMount**](ShipMount.md) |  | 

## Methods

### NewShipyardShip

`func NewShipyardShip(name string, description string, purchasePrice int32, frame ShipFrame, reactor ShipReactor, engine ShipEngine, modules []ShipModule, mounts []ShipMount, ) *ShipyardShip`

NewShipyardShip instantiates a new ShipyardShip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipyardShipWithDefaults

`func NewShipyardShipWithDefaults() *ShipyardShip`

NewShipyardShipWithDefaults instantiates a new ShipyardShip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShipyardShip) GetType() ShipType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipyardShip) GetTypeOk() (*ShipType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipyardShip) SetType(v ShipType)`

SetType sets Type field to given value.

### HasType

`func (o *ShipyardShip) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ShipyardShip) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipyardShip) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipyardShip) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ShipyardShip) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShipyardShip) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShipyardShip) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPurchasePrice

`func (o *ShipyardShip) GetPurchasePrice() int32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *ShipyardShip) GetPurchasePriceOk() (*int32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *ShipyardShip) SetPurchasePrice(v int32)`

SetPurchasePrice sets PurchasePrice field to given value.


### GetFrame

`func (o *ShipyardShip) GetFrame() ShipFrame`

GetFrame returns the Frame field if non-nil, zero value otherwise.

### GetFrameOk

`func (o *ShipyardShip) GetFrameOk() (*ShipFrame, bool)`

GetFrameOk returns a tuple with the Frame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrame

`func (o *ShipyardShip) SetFrame(v ShipFrame)`

SetFrame sets Frame field to given value.


### GetReactor

`func (o *ShipyardShip) GetReactor() ShipReactor`

GetReactor returns the Reactor field if non-nil, zero value otherwise.

### GetReactorOk

`func (o *ShipyardShip) GetReactorOk() (*ShipReactor, bool)`

GetReactorOk returns a tuple with the Reactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactor

`func (o *ShipyardShip) SetReactor(v ShipReactor)`

SetReactor sets Reactor field to given value.


### GetEngine

`func (o *ShipyardShip) GetEngine() ShipEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ShipyardShip) GetEngineOk() (*ShipEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ShipyardShip) SetEngine(v ShipEngine)`

SetEngine sets Engine field to given value.


### GetModules

`func (o *ShipyardShip) GetModules() []ShipModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ShipyardShip) GetModulesOk() (*[]ShipModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ShipyardShip) SetModules(v []ShipModule)`

SetModules sets Modules field to given value.


### GetMounts

`func (o *ShipyardShip) GetMounts() []ShipMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *ShipyardShip) GetMountsOk() (*[]ShipMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *ShipyardShip) SetMounts(v []ShipMount)`

SetMounts sets Mounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


