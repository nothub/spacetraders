# ConnectedSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** |  | 
**SectorSymbol** | **string** |  | 
**Type** | [**SystemType**](SystemType.md) |  | 
**FactionSymbol** | Pointer to **string** | The symbol of the faction that owns the connected jump gate in the system. | [optional] 
**X** | **int32** |  | 
**Y** | **int32** |  | 
**Distance** | **int32** |  | 

## Methods

### NewConnectedSystem

`func NewConnectedSystem(symbol string, sectorSymbol string, type_ SystemType, x int32, y int32, distance int32, ) *ConnectedSystem`

NewConnectedSystem instantiates a new ConnectedSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedSystemWithDefaults

`func NewConnectedSystemWithDefaults() *ConnectedSystem`

NewConnectedSystemWithDefaults instantiates a new ConnectedSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ConnectedSystem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ConnectedSystem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ConnectedSystem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSectorSymbol

`func (o *ConnectedSystem) GetSectorSymbol() string`

GetSectorSymbol returns the SectorSymbol field if non-nil, zero value otherwise.

### GetSectorSymbolOk

`func (o *ConnectedSystem) GetSectorSymbolOk() (*string, bool)`

GetSectorSymbolOk returns a tuple with the SectorSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSymbol

`func (o *ConnectedSystem) SetSectorSymbol(v string)`

SetSectorSymbol sets SectorSymbol field to given value.


### GetType

`func (o *ConnectedSystem) GetType() SystemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectedSystem) GetTypeOk() (*SystemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectedSystem) SetType(v SystemType)`

SetType sets Type field to given value.


### GetFactionSymbol

`func (o *ConnectedSystem) GetFactionSymbol() string`

GetFactionSymbol returns the FactionSymbol field if non-nil, zero value otherwise.

### GetFactionSymbolOk

`func (o *ConnectedSystem) GetFactionSymbolOk() (*string, bool)`

GetFactionSymbolOk returns a tuple with the FactionSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactionSymbol

`func (o *ConnectedSystem) SetFactionSymbol(v string)`

SetFactionSymbol sets FactionSymbol field to given value.

### HasFactionSymbol

`func (o *ConnectedSystem) HasFactionSymbol() bool`

HasFactionSymbol returns a boolean if a field has been set.

### GetX

`func (o *ConnectedSystem) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ConnectedSystem) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ConnectedSystem) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *ConnectedSystem) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ConnectedSystem) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ConnectedSystem) SetY(v int32)`

SetY sets Y field to given value.


### GetDistance

`func (o *ConnectedSystem) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *ConnectedSystem) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *ConnectedSystem) SetDistance(v int32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


