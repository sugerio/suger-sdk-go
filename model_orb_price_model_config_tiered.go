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

// checks if the OrbPriceModelConfigTIERED type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbPriceModelConfigTIERED{}

// OrbPriceModelConfigTIERED struct for OrbPriceModelConfigTIERED
type OrbPriceModelConfigTIERED struct {
	Tiers []OrbPriceTier `json:"tiers,omitempty"`
}

// NewOrbPriceModelConfigTIERED instantiates a new OrbPriceModelConfigTIERED object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbPriceModelConfigTIERED() *OrbPriceModelConfigTIERED {
	this := OrbPriceModelConfigTIERED{}
	return &this
}

// NewOrbPriceModelConfigTIEREDWithDefaults instantiates a new OrbPriceModelConfigTIERED object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbPriceModelConfigTIEREDWithDefaults() *OrbPriceModelConfigTIERED {
	this := OrbPriceModelConfigTIERED{}
	return &this
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *OrbPriceModelConfigTIERED) GetTiers() []OrbPriceTier {
	if o == nil || IsNil(o.Tiers) {
		var ret []OrbPriceTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigTIERED) GetTiersOk() ([]OrbPriceTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *OrbPriceModelConfigTIERED) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []OrbPriceTier and assigns it to the Tiers field.
func (o *OrbPriceModelConfigTIERED) SetTiers(v []OrbPriceTier) {
	o.Tiers = v
}

func (o OrbPriceModelConfigTIERED) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbPriceModelConfigTIERED) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}
	return toSerialize, nil
}

type NullableOrbPriceModelConfigTIERED struct {
	value *OrbPriceModelConfigTIERED
	isSet bool
}

func (v NullableOrbPriceModelConfigTIERED) Get() *OrbPriceModelConfigTIERED {
	return v.value
}

func (v *NullableOrbPriceModelConfigTIERED) Set(val *OrbPriceModelConfigTIERED) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPriceModelConfigTIERED) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPriceModelConfigTIERED) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPriceModelConfigTIERED(val *OrbPriceModelConfigTIERED) *NullableOrbPriceModelConfigTIERED {
	return &NullableOrbPriceModelConfigTIERED{value: val, isSet: true}
}

func (v NullableOrbPriceModelConfigTIERED) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPriceModelConfigTIERED) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
