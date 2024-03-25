# ScannedSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Symbol of the system. | 
**SectorSymbol** | **string** | Symbol of the system&#39;s sector. | 
**Type** | [**SystemType**](SystemType.md) |  | 
**X** | **int32** | Position in the universe in the x axis. | 
**Y** | **int32** | Position in the universe in the y axis. | 
**Distance** | **int32** | The system&#39;s distance from the scanning ship. | 

## Methods

### NewScannedSystem

`func NewScannedSystem(symbol string, sectorSymbol string, type_ SystemType, x int32, y int32, distance int32, ) *ScannedSystem`

NewScannedSystem instantiates a new ScannedSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScannedSystemWithDefaults

`func NewScannedSystemWithDefaults() *ScannedSystem`

NewScannedSystemWithDefaults instantiates a new ScannedSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ScannedSystem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ScannedSystem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ScannedSystem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetSectorSymbol

`func (o *ScannedSystem) GetSectorSymbol() string`

GetSectorSymbol returns the SectorSymbol field if non-nil, zero value otherwise.

### GetSectorSymbolOk

`func (o *ScannedSystem) GetSectorSymbolOk() (*string, bool)`

GetSectorSymbolOk returns a tuple with the SectorSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorSymbol

`func (o *ScannedSystem) SetSectorSymbol(v string)`

SetSectorSymbol sets SectorSymbol field to given value.


### GetType

`func (o *ScannedSystem) GetType() SystemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScannedSystem) GetTypeOk() (*SystemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScannedSystem) SetType(v SystemType)`

SetType sets Type field to given value.


### GetX

`func (o *ScannedSystem) GetX() int32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ScannedSystem) GetXOk() (*int32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ScannedSystem) SetX(v int32)`

SetX sets X field to given value.


### GetY

`func (o *ScannedSystem) GetY() int32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ScannedSystem) GetYOk() (*int32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ScannedSystem) SetY(v int32)`

SetY sets Y field to given value.


### GetDistance

`func (o *ScannedSystem) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *ScannedSystem) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *ScannedSystem) SetDistance(v int32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


