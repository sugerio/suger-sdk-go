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

// checks if the OrbPriceModelConfigBPS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbPriceModelConfigBPS{}

// OrbPriceModelConfigBPS struct for OrbPriceModelConfigBPS
type OrbPriceModelConfigBPS struct {
	Bps            *float32 `json:"bps,omitempty"`
	PerUnitMaximum *string  `json:"per_unit_maximum,omitempty"`
}

// NewOrbPriceModelConfigBPS instantiates a new OrbPriceModelConfigBPS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbPriceModelConfigBPS() *OrbPriceModelConfigBPS {
	this := OrbPriceModelConfigBPS{}
	return &this
}

// NewOrbPriceModelConfigBPSWithDefaults instantiates a new OrbPriceModelConfigBPS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbPriceModelConfigBPSWithDefaults() *OrbPriceModelConfigBPS {
	this := OrbPriceModelConfigBPS{}
	return &this
}

// GetBps returns the Bps field value if set, zero value otherwise.
func (o *OrbPriceModelConfigBPS) GetBps() float32 {
	if o == nil || IsNil(o.Bps) {
		var ret float32
		return ret
	}
	return *o.Bps
}

// GetBpsOk returns a tuple with the Bps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigBPS) GetBpsOk() (*float32, bool) {
	if o == nil || IsNil(o.Bps) {
		return nil, false
	}
	return o.Bps, true
}

// HasBps returns a boolean if a field has been set.
func (o *OrbPriceModelConfigBPS) HasBps() bool {
	if o != nil && !IsNil(o.Bps) {
		return true
	}

	return false
}

// SetBps gets a reference to the given float32 and assigns it to the Bps field.
func (o *OrbPriceModelConfigBPS) SetBps(v float32) {
	o.Bps = &v
}

// GetPerUnitMaximum returns the PerUnitMaximum field value if set, zero value otherwise.
func (o *OrbPriceModelConfigBPS) GetPerUnitMaximum() string {
	if o == nil || IsNil(o.PerUnitMaximum) {
		var ret string
		return ret
	}
	return *o.PerUnitMaximum
}

// GetPerUnitMaximumOk returns a tuple with the PerUnitMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigBPS) GetPerUnitMaximumOk() (*string, bool) {
	if o == nil || IsNil(o.PerUnitMaximum) {
		return nil, false
	}
	return o.PerUnitMaximum, true
}

// HasPerUnitMaximum returns a boolean if a field has been set.
func (o *OrbPriceModelConfigBPS) HasPerUnitMaximum() bool {
	if o != nil && !IsNil(o.PerUnitMaximum) {
		return true
	}

	return false
}

// SetPerUnitMaximum gets a reference to the given string and assigns it to the PerUnitMaximum field.
func (o *OrbPriceModelConfigBPS) SetPerUnitMaximum(v string) {
	o.PerUnitMaximum = &v
}

func (o OrbPriceModelConfigBPS) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbPriceModelConfigBPS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bps) {
		toSerialize["bps"] = o.Bps
	}
	if !IsNil(o.PerUnitMaximum) {
		toSerialize["per_unit_maximum"] = o.PerUnitMaximum
	}
	return toSerialize, nil
}

type NullableOrbPriceModelConfigBPS struct {
	value *OrbPriceModelConfigBPS
	isSet bool
}

func (v NullableOrbPriceModelConfigBPS) Get() *OrbPriceModelConfigBPS {
	return v.value
}

func (v *NullableOrbPriceModelConfigBPS) Set(val *OrbPriceModelConfigBPS) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPriceModelConfigBPS) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPriceModelConfigBPS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPriceModelConfigBPS(val *OrbPriceModelConfigBPS) *NullableOrbPriceModelConfigBPS {
	return &NullableOrbPriceModelConfigBPS{value: val, isSet: true}
}

func (v NullableOrbPriceModelConfigBPS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPriceModelConfigBPS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
