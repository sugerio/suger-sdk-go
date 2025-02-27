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

// checks if the OrbTrialConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbTrialConfig{}

// OrbTrialConfig struct for OrbTrialConfig
type OrbTrialConfig struct {
	TrialPeriod     *int32  `json:"trial_period,omitempty"`
	TrialPeriodUnit *string `json:"trial_period_unit,omitempty"`
}

// NewOrbTrialConfig instantiates a new OrbTrialConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbTrialConfig() *OrbTrialConfig {
	this := OrbTrialConfig{}
	return &this
}

// NewOrbTrialConfigWithDefaults instantiates a new OrbTrialConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbTrialConfigWithDefaults() *OrbTrialConfig {
	this := OrbTrialConfig{}
	return &this
}

// GetTrialPeriod returns the TrialPeriod field value if set, zero value otherwise.
func (o *OrbTrialConfig) GetTrialPeriod() int32 {
	if o == nil || IsNil(o.TrialPeriod) {
		var ret int32
		return ret
	}
	return *o.TrialPeriod
}

// GetTrialPeriodOk returns a tuple with the TrialPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbTrialConfig) GetTrialPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.TrialPeriod) {
		return nil, false
	}
	return o.TrialPeriod, true
}

// HasTrialPeriod returns a boolean if a field has been set.
func (o *OrbTrialConfig) HasTrialPeriod() bool {
	if o != nil && !IsNil(o.TrialPeriod) {
		return true
	}

	return false
}

// SetTrialPeriod gets a reference to the given int32 and assigns it to the TrialPeriod field.
func (o *OrbTrialConfig) SetTrialPeriod(v int32) {
	o.TrialPeriod = &v
}

// GetTrialPeriodUnit returns the TrialPeriodUnit field value if set, zero value otherwise.
func (o *OrbTrialConfig) GetTrialPeriodUnit() string {
	if o == nil || IsNil(o.TrialPeriodUnit) {
		var ret string
		return ret
	}
	return *o.TrialPeriodUnit
}

// GetTrialPeriodUnitOk returns a tuple with the TrialPeriodUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbTrialConfig) GetTrialPeriodUnitOk() (*string, bool) {
	if o == nil || IsNil(o.TrialPeriodUnit) {
		return nil, false
	}
	return o.TrialPeriodUnit, true
}

// HasTrialPeriodUnit returns a boolean if a field has been set.
func (o *OrbTrialConfig) HasTrialPeriodUnit() bool {
	if o != nil && !IsNil(o.TrialPeriodUnit) {
		return true
	}

	return false
}

// SetTrialPeriodUnit gets a reference to the given string and assigns it to the TrialPeriodUnit field.
func (o *OrbTrialConfig) SetTrialPeriodUnit(v string) {
	o.TrialPeriodUnit = &v
}

func (o OrbTrialConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbTrialConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrialPeriod) {
		toSerialize["trial_period"] = o.TrialPeriod
	}
	if !IsNil(o.TrialPeriodUnit) {
		toSerialize["trial_period_unit"] = o.TrialPeriodUnit
	}
	return toSerialize, nil
}

type NullableOrbTrialConfig struct {
	value *OrbTrialConfig
	isSet bool
}

func (v NullableOrbTrialConfig) Get() *OrbTrialConfig {
	return v.value
}

func (v *NullableOrbTrialConfig) Set(val *OrbTrialConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbTrialConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbTrialConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbTrialConfig(val *OrbTrialConfig) *NullableOrbTrialConfig {
	return &NullableOrbTrialConfig{value: val, isSet: true}
}

func (v NullableOrbTrialConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbTrialConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
