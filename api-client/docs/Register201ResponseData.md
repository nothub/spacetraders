# Register201ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | [**Agent**](Agent.md) |  | 
**Contract** | [**Contract**](Contract.md) |  | 
**Faction** | [**Faction**](Faction.md) |  | 
**Ship** | [**Ship**](Ship.md) |  | 
**Token** | **string** | A Bearer token for accessing secured API endpoints. | 

## Methods

### NewRegister201ResponseData

`func NewRegister201ResponseData(agent Agent, contract Contract, faction Faction, ship Ship, token string, ) *Register201ResponseData`

NewRegister201ResponseData instantiates a new Register201ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegister201ResponseDataWithDefaults

`func NewRegister201ResponseDataWithDefaults() *Register201ResponseData`

NewRegister201ResponseDataWithDefaults instantiates a new Register201ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *Register201ResponseData) GetAgent() Agent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *Register201ResponseData) GetAgentOk() (*Agent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *Register201ResponseData) SetAgent(v Agent)`

SetAgent sets Agent field to given value.


### GetContract

`func (o *Register201ResponseData) GetContract() Contract`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *Register201ResponseData) GetContractOk() (*Contract, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *Register201ResponseData) SetContract(v Contract)`

SetContract sets Contract field to given value.


### GetFaction

`func (o *Register201ResponseData) GetFaction() Faction`

GetFaction returns the Faction field if non-nil, zero value otherwise.

### GetFactionOk

`func (o *Register201ResponseData) GetFactionOk() (*Faction, bool)`

GetFactionOk returns a tuple with the Faction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaction

`func (o *Register201ResponseData) SetFaction(v Faction)`

SetFaction sets Faction field to given value.


### GetShip

`func (o *Register201ResponseData) GetShip() Ship`

GetShip returns the Ship field if non-nil, zero value otherwise.

### GetShipOk

`func (o *Register201ResponseData) GetShipOk() (*Ship, bool)`

GetShipOk returns a tuple with the Ship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShip

`func (o *Register201ResponseData) SetShip(v Ship)`

SetShip sets Ship field to given value.


### GetToken

`func (o *Register201ResponseData) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Register201ResponseData) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Register201ResponseData) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


