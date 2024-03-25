# ScannedShip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The globally unique identifier of the ship. | 
**Registration** | [**ShipRegistration**](ShipRegistration.md) |  | 
**Nav** | [**ShipNav**](ShipNav.md) |  | 
**Frame** | Pointer to [**ScannedShipFrame**](ScannedShipFrame.md) |  | [optional] 
**Reactor** | Pointer to [**ScannedShipReactor**](ScannedShipReactor.md) |  | [optional] 
**Engine** | [**ScannedShipEngine**](ScannedShipEngine.md) |  | 
**Mounts** | Pointer to [**[]ScannedShipMountsInner**](ScannedShipMountsInner.md) | List of mounts installed in the ship. | [optional] 

## Methods

### NewScannedShip

`func NewScannedShip(symbol string, registration ShipRegistration, nav ShipNav, engine ScannedShipEngine, ) *ScannedShip`

NewScannedShip instantiates a new ScannedShip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScannedShipWithDefaults

`func NewScannedShipWithDefaults() *ScannedShip`

NewScannedShipWithDefaults instantiates a new ScannedShip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ScannedShip) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ScannedShip) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ScannedShip) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetRegistration

`func (o *ScannedShip) GetRegistration() ShipRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *ScannedShip) GetRegistrationOk() (*ShipRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *ScannedShip) SetRegistration(v ShipRegistration)`

SetRegistration sets Registration field to given value.


### GetNav

`func (o *ScannedShip) GetNav() ShipNav`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *ScannedShip) GetNavOk() (*ShipNav, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *ScannedShip) SetNav(v ShipNav)`

SetNav sets Nav field to given value.


### GetFrame

`func (o *ScannedShip) GetFrame() ScannedShipFrame`

GetFrame returns the Frame field if non-nil, zero value otherwise.

### GetFrameOk

`func (o *ScannedShip) GetFrameOk() (*ScannedShipFrame, bool)`

GetFrameOk returns a tuple with the Frame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrame

`func (o *ScannedShip) SetFrame(v ScannedShipFrame)`

SetFrame sets Frame field to given value.

### HasFrame

`func (o *ScannedShip) HasFrame() bool`

HasFrame returns a boolean if a field has been set.

### GetReactor

`func (o *ScannedShip) GetReactor() ScannedShipReactor`

GetReactor returns the Reactor field if non-nil, zero value otherwise.

### GetReactorOk

`func (o *ScannedShip) GetReactorOk() (*ScannedShipReactor, bool)`

GetReactorOk returns a tuple with the Reactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactor

`func (o *ScannedShip) SetReactor(v ScannedShipReactor)`

SetReactor sets Reactor field to given value.

### HasReactor

`func (o *ScannedShip) HasReactor() bool`

HasReactor returns a boolean if a field has been set.

### GetEngine

`func (o *ScannedShip) GetEngine() ScannedShipEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ScannedShip) GetEngineOk() (*ScannedShipEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ScannedShip) SetEngine(v ScannedShipEngine)`

SetEngine sets Engine field to given value.


### GetMounts

`func (o *ScannedShip) GetMounts() []ScannedShipMountsInner`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *ScannedShip) GetMountsOk() (*[]ScannedShipMountsInner, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *ScannedShip) SetMounts(v []ScannedShipMountsInner)`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *ScannedShip) HasMounts() bool`

HasMounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


