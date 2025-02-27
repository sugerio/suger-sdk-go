/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suger

import (
	"encoding/json"
)

// checks if the OrbPriceModelConfigTIEREDBPS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbPriceModelConfigTIEREDBPS{}

// OrbPriceModelConfigTIEREDBPS struct for OrbPriceModelConfigTIEREDBPS
type OrbPriceModelConfigTIEREDBPS struct {
	Tiers []OrbPriceTier `json:"tiers,omitempty"`
}

// NewOrbPriceModelConfigTIEREDBPS instantiates a new OrbPriceModelConfigTIEREDBPS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbPriceModelConfigTIEREDBPS() *OrbPriceModelConfigTIEREDBPS {
	this := OrbPriceModelConfigTIEREDBPS{}
	return &this
}

// NewOrbPriceModelConfigTIEREDBPSWithDefaults instantiates a new OrbPriceModelConfigTIEREDBPS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbPriceModelConfigTIEREDBPSWithDefaults() *OrbPriceModelConfigTIEREDBPS {
	this := OrbPriceModelConfigTIEREDBPS{}
	return &this
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *OrbPriceModelConfigTIEREDBPS) GetTiers() []OrbPriceTier {
	if o == nil || IsNil(o.Tiers) {
		var ret []OrbPriceTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigTIEREDBPS) GetTiersOk() ([]OrbPriceTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *OrbPriceModelConfigTIEREDBPS) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []OrbPriceTier and assigns it to the Tiers field.
func (o *OrbPriceModelConfigTIEREDBPS) SetTiers(v []OrbPriceTier) {
	o.Tiers = v
}

func (o OrbPriceModelConfigTIEREDBPS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbPriceModelConfigTIEREDBPS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}
	return toSerialize, nil
}

type NullableOrbPriceModelConfigTIEREDBPS struct {
	value *OrbPriceModelConfigTIEREDBPS
	isSet bool
}

func (v NullableOrbPriceModelConfigTIEREDBPS) Get() *OrbPriceModelConfigTIEREDBPS {
	return v.value
}

func (v *NullableOrbPriceModelConfigTIEREDBPS) Set(val *OrbPriceModelConfigTIEREDBPS) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPriceModelConfigTIEREDBPS) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPriceModelConfigTIEREDBPS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPriceModelConfigTIEREDBPS(val *OrbPriceModelConfigTIEREDBPS) *NullableOrbPriceModelConfigTIEREDBPS {
	return &NullableOrbPriceModelConfigTIEREDBPS{value: val, isSet: true}
}

func (v NullableOrbPriceModelConfigTIEREDBPS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPriceModelConfigTIEREDBPS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
