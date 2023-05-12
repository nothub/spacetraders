# JumpGate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JumpRange** | **float32** | The maximum jump range of the gate. | 
**FactionSymbol** | Pointer to **string** | The symbol of the faction that owns the gate. | [optional] 
**ConnectedSystems** | [**[]ConnectedSystem**](ConnectedSystem.md) | The systems within range of the gate that have a corresponding gate. | 

## Methods

### NewJumpGate

`func NewJumpGate(jumpRange float32, connectedSystems []ConnectedSystem, ) *JumpGate`

NewJumpGate instantiates a new JumpGate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpGateWithDefaults

`func NewJumpGateWithDefaults() *JumpGate`

NewJumpGateWithDefaults instantiates a new JumpGate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJumpRange

`func (o *JumpGate) GetJumpRange() float32`

GetJumpRange returns the JumpRange field if non-nil, zero value otherwise.

### GetJumpRangeOk

`func (o *JumpGate) GetJumpRangeOk() (*float32, bool)`

GetJumpRangeOk returns a tuple with the JumpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJumpRange

`func (o *JumpGate) SetJumpRange(v float32)`

SetJumpRange sets JumpRange field to given value.


### GetFactionSymbol

`func (o *JumpGate) GetFactionSymbol() string`

GetFactionSymbol returns the FactionSymbol field if non-nil, zero value otherwise.

### GetFactionSymbolOk

`func (o *JumpGate) GetFactionSymbolOk() (*string, bool)`

GetFactionSymbolOk returns a tuple with the FactionSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionSymbol

`func (o *JumpGate) SetFactionSymbol(v string)`

SetFactionSymbol sets FactionSymbol field to given value.

### HasFactionSymbol

`func (o *JumpGate) HasFactionSymbol() bool`

HasFactionSymbol returns a boolean if a field has been set.

### GetConnectedSystems

`func (o *JumpGate) GetConnectedSystems() []ConnectedSystem`

GetConnectedSystems returns the ConnectedSystems field if non-nil, zero value otherwise.

### GetConnectedSystemsOk

`func (o *JumpGate) GetConnectedSystemsOk() (*[]ConnectedSystem, bool)`

GetConnectedSystemsOk returns a tuple with the ConnectedSystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSystems

`func (o *JumpGate) SetConnectedSystems(v []ConnectedSystem)`

SetConnectedSystems sets ConnectedSystems field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


