# Chart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WaypointSymbol** | Pointer to **string** | The symbol of the waypoint. | [optional] 
**SubmittedBy** | Pointer to **string** | The agent that submitted the chart for this waypoint. | [optional] 
**SubmittedOn** | Pointer to **time.Time** | The time the chart for this waypoint was submitted. | [optional] 

## Methods

### NewChart

`func NewChart() *Chart`

NewChart instantiates a new Chart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartWithDefaults

`func NewChartWithDefaults() *Chart`

NewChartWithDefaults instantiates a new Chart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWaypointSymbol

`func (o *Chart) GetWaypointSymbol() string`

GetWaypointSymbol returns the WaypointSymbol field if non-nil, zero value otherwise.

### GetWaypointSymbolOk

`func (o *Chart) GetWaypointSymbolOk() (*string, bool)`

GetWaypointSymbolOk returns a tuple with the WaypointSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaypointSymbol

`func (o *Chart) SetWaypointSymbol(v string)`

SetWaypointSymbol sets WaypointSymbol field to given value.

### HasWaypointSymbol

`func (o *Chart) HasWaypointSymbol() bool`

HasWaypointSymbol returns a boolean if a field has been set.

### GetSubmittedBy

`func (o *Chart) GetSubmittedBy() string`

GetSubmittedBy returns the SubmittedBy field if non-nil, zero value otherwise.

### GetSubmittedByOk

`func (o *Chart) GetSubmittedByOk() (*string, bool)`

GetSubmittedByOk returns a tuple with the SubmittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedBy

`func (o *Chart) SetSubmittedBy(v string)`

SetSubmittedBy sets SubmittedBy field to given value.

### HasSubmittedBy

`func (o *Chart) HasSubmittedBy() bool`

HasSubmittedBy returns a boolean if a field has been set.

### GetSubmittedOn

`func (o *Chart) GetSubmittedOn() time.Time`

GetSubmittedOn returns the SubmittedOn field if non-nil, zero value otherwise.

### GetSubmittedOnOk

`func (o *Chart) GetSubmittedOnOk() (*time.Time, bool)`

GetSubmittedOnOk returns a tuple with the SubmittedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedOn

`func (o *Chart) SetSubmittedOn(v time.Time)`

SetSubmittedOn sets SubmittedOn field to given value.

### HasSubmittedOn

`func (o *Chart) HasSubmittedOn() bool`

HasSubmittedOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


