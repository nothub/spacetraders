# JumpGate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The symbol of the waypoint. | 
**Connections** | **[]string** | All the gates that are connected to this waypoint. | 

## Methods

### NewJumpGate

`func NewJumpGate(symbol string, connections []string, ) *JumpGate`

NewJumpGate instantiates a new JumpGate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJumpGateWithDefaults

`func NewJumpGateWithDefaults() *JumpGate`

NewJumpGateWithDefaults instantiates a new JumpGate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *JumpGate) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *JumpGate) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *JumpGate) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetConnections

`func (o *JumpGate) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *JumpGate) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *JumpGate) SetConnections(v []string)`

SetConnections sets Connections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


