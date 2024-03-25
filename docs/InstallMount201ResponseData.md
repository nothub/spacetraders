# InstallMount201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**Agent**](Agent.md) |  | 
**Mounts** | [**[]ShipMount**](ShipMount.md) | List of installed mounts after the installation of the new mount. | 
**Cargo** | [**ShipCargo**](ShipCargo.md) |  | 
**Transaction** | [**ShipModificationTransaction**](ShipModificationTransaction.md) |  | 

## Methods

### NewInstallMount201ResponseData

`func NewInstallMount201ResponseData(agent Agent, mounts []ShipMount, cargo ShipCargo, transaction ShipModificationTransaction, ) *InstallMount201ResponseData`

NewInstallMount201ResponseData instantiates a new InstallMount201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallMount201ResponseDataWithDefaults

`func NewInstallMount201ResponseDataWithDefaults() *InstallMount201ResponseData`

NewInstallMount201ResponseDataWithDefaults instantiates a new InstallMount201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *InstallMount201ResponseData) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *InstallMount201ResponseData) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *InstallMount201ResponseData) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetMounts

`func (o *InstallMount201ResponseData) GetMounts() []ShipMount`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *InstallMount201ResponseData) GetMountsOk() (*[]ShipMount, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *InstallMount201ResponseData) SetMounts(v []ShipMount)`

SetMounts sets Mounts field to given value.


### GetCargo

`func (o *InstallMount201ResponseData) GetCargo() ShipCargo`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *InstallMount201ResponseData) GetCargoOk() (*ShipCargo, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *InstallMount201ResponseData) SetCargo(v ShipCargo)`

SetCargo sets Cargo field to given value.


### GetTransaction

`func (o *InstallMount201ResponseData) GetTransaction() ShipModificationTransaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *InstallMount201ResponseData) GetTransactionOk() (*ShipModificationTransaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *InstallMount201ResponseData) SetTransaction(v ShipModificationTransaction)`

SetTransaction sets Transaction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


