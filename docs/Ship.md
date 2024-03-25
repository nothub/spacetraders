# Ship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The globally unique identifier of the ship in the following format: &#x60;[AGENT_SYMBOL]-[HEX_ID]&#x60; | 
**Registration** | [**ShipRegistration**](ShipRegistration.md) |  | 
**Nav** | [**ShipNav**](ShipNav.md) |  | 
**Crew** | [**ShipCrew**](ShipCrew.md) |  | 
**Frame** | [**ShipFrame**](ShipFrame.md) |  | 
**Reactor** | [**ShipReactor**](ShipReactor.md) |  | 
**Engine** | [**ShipEngine**](ShipEngine.md) |  | 
**Cooldown** | [**Cooldown**](Cooldown.md) |  | 
**Modules** | [**[]ShipModule**](ShipModule.md) | Modules installed in this ship. | 
**Mounts** | [**[]ShipMount**](ShipMount.md) | Mounts installed in this ship. | 
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Fuel** | [**ShipFuel**](ShipFuel.md) |  | 

## Methods

### NewShip

`func NewShip(symbol string, registration ShipRegistration, nav ShipNav, crew ShipCrew, frame ShipFrame, reactor ShipReactor, engine ShipEngine, cooldown Cooldown, modules []ShipModule, mounts []ShipMount, cargo ShipCargo, fuel ShipFuel, ) *Ship`

NewShip instantiates a new Ship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipWithDefaults

`func NewShipWithDefaults() *Ship`

NewShipWithDefaults instantiates a new Ship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *Ship) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Ship) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Ship) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetRegistration

`func (o *Ship) GetRegistration() ShipRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *Ship) GetRegistrationOk() (*ShipRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *Ship) SetRegistration(v ShipRegistration)`

SetRegistration sets Registration field to given value.


### GetNav

`func (o *Ship) GetNav() ShipNav`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *Ship) GetNavOk() (*ShipNav, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *Ship) SetNav(v ShipNav)`

SetNav sets Nav field to given value.


### GetCrew

`func (o *Ship) GetCrew() ShipCrew`

GetCrew returns the Crew field if non-nil, zero value otherwise.

### GetCrewOk

`func (o *Ship) GetCrewOk() (*ShipCrew, bool)`

GetCrewOk returns a tuple with the Crew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrew

`func (o *Ship) SetCrew(v ShipCrew)`

SetCrew sets Crew field to given value.


### GetFrame

`func (o *Ship) GetFrame() ShipFrame`

GetFrame returns the Frame field if non-nil, zero value otherwise.

### GetFrameOk

`func (o *Ship) GetFrameOk() (*ShipFrame, bool)`

GetFrameOk returns a tuple with the Frame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrame

`func (o *Ship) SetFrame(v ShipFrame)`

SetFrame sets Frame field to given value.


### GetReactor

`func (o *Ship) GetReactor() ShipReactor`

GetReactor returns the Reactor field if non-nil, zero value otherwise.

### GetReactorOk

`func (o *Ship) GetReactorOk() (*ShipReactor, bool)`

GetReactorOk returns a tuple with the Reactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactor

`func (o *Ship) SetReactor(v ShipReactor)`

SetReactor sets Reactor field to given value.


### GetEngine

`func (o *Ship) GetEngine() ShipEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *Ship) GetEngineOk() (*ShipEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *Ship) SetEngine(v ShipEngine)`

SetEngine sets Engine field to given value.


### GetCooldown

`func (o *Ship) GetCooldown() Cooldown`

GetCooldown returns the Cooldown field if non-nil, zero value otherwise.

### GetCooldownOk

`func (o *Ship) GetCooldownOk() (*Cooldown, bool)`

GetCooldownOk returns a tuple with the Cooldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldown

`func (o *Ship) SetCooldown(v Cooldown)`

SetCooldown sets Cooldown field to given value.


### GetModules

`func (o *Ship) GetModules() []ShipModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *Ship) GetModulesOk() (*[]ShipModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *Ship) SetModules(v []ShipModule)`

SetModules sets Modules field to given value.


### GetMounts

`func (o *Ship) GetMounts() []ShipMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *Ship) GetMountsOk() (*[]ShipMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *Ship) SetMounts(v []ShipMount)`

SetMounts sets Mounts field to given value.


### GetCargo

`func (o *Ship) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *Ship) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *Ship) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetFuel

`func (o *Ship) GetFuel() ShipFuel`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *Ship) GetFuelOk() (*ShipFuel, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *Ship) SetFuel(v ShipFuel)`

SetFuel sets Fuel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


