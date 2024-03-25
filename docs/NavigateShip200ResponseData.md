# NavigateShip200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fuel** | [**ShipFuel**](ShipFuel.md) |  | 
**Nav** | [**ShipNav**](ShipNav.md) |  | 
**Events** | [**[]ExtractResources201ResponseDataEventsInner**](ExtractResources201ResponseDataEventsInner.md) |  | 

## Methods

### NewNavigateShip200ResponseData

`func NewNavigateShip200ResponseData(fuel ShipFuel, nav ShipNav, events []ExtractResources201ResponseDataEventsInner, ) *NavigateShip200ResponseData`

NewNavigateShip200ResponseData instantiates a new NavigateShip200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigateShip200ResponseDataWithDefaults

`func NewNavigateShip200ResponseDataWithDefaults() *NavigateShip200ResponseData`

NewNavigateShip200ResponseDataWithDefaults instantiates a new NavigateShip200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFuel

`func (o *NavigateShip200ResponseData) GetFuel() ShipFuel`

GetFuel returns the Fuel field if non-nil, zero value otherwise.

### GetFuelOk

`func (o *NavigateShip200ResponseData) GetFuelOk() (*ShipFuel, bool)`

GetFuelOk returns a tuple with the Fuel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuel

`func (o *NavigateShip200ResponseData) SetFuel(v ShipFuel)`

SetFuel sets Fuel field to given value.


### GetNav

`func (o *NavigateShip200ResponseData) GetNav() ShipNav`

GetNav returns the Nav field if non-nil, zero value otherwise.

### GetNavOk

`func (o *NavigateShip200ResponseData) GetNavOk() (*ShipNav, bool)`

GetNavOk returns a tuple with the Nav field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNav

`func (o *NavigateShip200ResponseData) SetNav(v ShipNav)`

SetNav sets Nav field to given value.


### GetEvents

`func (o *NavigateShip200ResponseData) GetEvents() []ExtractResources201ResponseDataEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NavigateShip200ResponseData) GetEventsOk() (*[]ExtractResources201ResponseDataEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NavigateShip200ResponseData) SetEvents(v []ExtractResources201ResponseDataEventsInner)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


