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
)

// checks if the ShipRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipRequirements{}

// ShipRequirements The requirements for installation on a ship
type ShipRequirements struct {
	// The amount of power required from the reactor.
	Power *int32 `json:"power,omitempty"`
	// The number of crew required for operation.
	Crew *int32 `json:"crew,omitempty"`
	// The number of module slots required for installation.
	Slots                *int32 `json:"slots,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ShipRequirements ShipRequirements

// NewShipRequirements instantiates a new ShipRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipRequirements() *ShipRequirements {
	this := ShipRequirements{}
	return &this
}

// NewShipRequirementsWithDefaults instantiates a new ShipRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipRequirementsWithDefaults() *ShipRequirements {
	this := ShipRequirements{}
	return &this
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *ShipRequirements) GetPower() int32 {
	if o == nil || IsNil(o.Power) {
		var ret int32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipRequirements) GetPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *ShipRequirements) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given int32 and assigns it to the Power field.
func (o *ShipRequirements) SetPower(v int32) {
	o.Power = &v
}

// GetCrew returns the Crew field value if set, zero value otherwise.
func (o *ShipRequirements) GetCrew() int32 {
	if o == nil || IsNil(o.Crew) {
		var ret int32
		return ret
	}
	return *o.Crew
}

// GetCrewOk returns a tuple with the Crew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipRequirements) GetCrewOk() (*int32, bool) {
	if o == nil || IsNil(o.Crew) {
		return nil, false
	}
	return o.Crew, true
}

// HasCrew returns a boolean if a field has been set.
func (o *ShipRequirements) HasCrew() bool {
	if o != nil && !IsNil(o.Crew) {
		return true
	}

	return false
}

// SetCrew gets a reference to the given int32 and assigns it to the Crew field.
func (o *ShipRequirements) SetCrew(v int32) {
	o.Crew = &v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *ShipRequirements) GetSlots() int32 {
	if o == nil || IsNil(o.Slots) {
		var ret int32
		return ret
	}
	return *o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipRequirements) GetSlotsOk() (*int32, bool) {
	if o == nil || IsNil(o.Slots) {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *ShipRequirements) HasSlots() bool {
	if o != nil && !IsNil(o.Slots) {
		return true
	}

	return false
}

// SetSlots gets a reference to the given int32 and assigns it to the Slots field.
func (o *ShipRequirements) SetSlots(v int32) {
	o.Slots = &v
}

func (o ShipRequirements) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !IsNil(o.Crew) {
		toSerialize["crew"] = o.Crew
	}
	if !IsNil(o.Slots) {
		toSerialize["slots"] = o.Slots
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipRequirements) UnmarshalJSON(data []byte) (err error) {
	varShipRequirements := _ShipRequirements{}

	err = json.Unmarshal(data, &varShipRequirements)

	if err != nil {
		return err
	}

	*o = ShipRequirements(varShipRequirements)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "power")
		delete(additionalProperties, "crew")
		delete(additionalProperties, "slots")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipRequirements struct {
	value *ShipRequirements
	isSet bool
}

func (v NullableShipRequirements) Get() *ShipRequirements {
	return v.value
}

func (v *NullableShipRequirements) Set(val *ShipRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRequirements(val *ShipRequirements) *NullableShipRequirements {
	return &NullableShipRequirements{value: val, isSet: true}
}

func (v NullableShipRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
