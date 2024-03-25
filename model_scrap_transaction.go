/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package st

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the ScrapTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScrapTransaction{}

// ScrapTransaction Result of a scrap transaction.
type ScrapTransaction struct {
	// The symbol of the waypoint.
	WaypointSymbol string `json:"waypointSymbol"`
	// The symbol of the ship.
	ShipSymbol string `json:"shipSymbol"`
	// The total price of the transaction.
	TotalPrice int32 `json:"totalPrice"`
	// The timestamp of the transaction.
	Timestamp            time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _ScrapTransaction ScrapTransaction

// NewScrapTransaction instantiates a new ScrapTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScrapTransaction(waypointSymbol string, shipSymbol string, totalPrice int32, timestamp time.Time) *ScrapTransaction {
	this := ScrapTransaction{}
	this.WaypointSymbol = waypointSymbol
	this.ShipSymbol = shipSymbol
	this.TotalPrice = totalPrice
	this.Timestamp = timestamp
	return &this
}

// NewScrapTransactionWithDefaults instantiates a new ScrapTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScrapTransactionWithDefaults() *ScrapTransaction {
	this := ScrapTransaction{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *ScrapTransaction) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *ScrapTransaction) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *ScrapTransaction) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

// GetShipSymbol returns the ShipSymbol field value
func (o *ScrapTransaction) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *ScrapTransaction) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *ScrapTransaction) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *ScrapTransaction) GetTotalPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *ScrapTransaction) GetTotalPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *ScrapTransaction) SetTotalPrice(v int32) {
	o.TotalPrice = v
}

// GetTimestamp returns the Timestamp field value
func (o *ScrapTransaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ScrapTransaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ScrapTransaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ScrapTransaction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScrapTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["totalPrice"] = o.TotalPrice
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScrapTransaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"waypointSymbol",
		"shipSymbol",
		"totalPrice",
		"timestamp",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varScrapTransaction := _ScrapTransaction{}

	err = json.Unmarshal(data, &varScrapTransaction)

	if err != nil {
		return err
	}

	*o = ScrapTransaction(varScrapTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "waypointSymbol")
		delete(additionalProperties, "shipSymbol")
		delete(additionalProperties, "totalPrice")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScrapTransaction struct {
	value *ScrapTransaction
	isSet bool
}

func (v NullableScrapTransaction) Get() *ScrapTransaction {
	return v.value
}

func (v *NullableScrapTransaction) Set(val *ScrapTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableScrapTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableScrapTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScrapTransaction(val *ScrapTransaction) *NullableScrapTransaction {
	return &NullableScrapTransaction{value: val, isSet: true}
}

func (v NullableScrapTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScrapTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
