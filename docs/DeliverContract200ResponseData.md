# DeliverContract200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | [**Contract**](Contract.md) |  | 
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 

## Methods

### NewDeliverContract200ResponseData

`func NewDeliverContract200ResponseData(contract Contract, cargo ShipCargo, ) *DeliverContract200ResponseData`

NewDeliverContract200ResponseData instantiates a new DeliverContract200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliverContract200ResponseDataWithDefaults

`func NewDeliverContract200ResponseDataWithDefaults() *DeliverContract200ResponseData`

NewDeliverContract200ResponseDataWithDefaults instantiates a new DeliverContract200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *DeliverContract200ResponseData) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *DeliverContract200ResponseData) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *DeliverContract200ResponseData) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetCargo

`func (o *DeliverContract200ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *DeliverContract200ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *DeliverContract200ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


