/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PriceModelTieredPercentageConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceModelTieredPercentageConfig{}

// PriceModelTieredPercentageConfig struct for PriceModelTieredPercentageConfig
type PriceModelTieredPercentageConfig struct {
	// Inclusive tier starting value
	FirstUnit *float32 `json:"firstUnit,omitempty"`
	FlatFee   *float32 `json:"flatFee,omitempty"`
	// Exclusive tier ending value. If null, this is treated as the last tier
	LastUnit       *float32 `json:"lastUnit,omitempty"`
	PercentageRate *float32 `json:"percentageRate,omitempty"`
}

// NewPriceModelTieredPercentageConfig instantiates a new PriceModelTieredPercentageConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceModelTieredPercentageConfig() *PriceModelTieredPercentageConfig {
	this := PriceModelTieredPercentageConfig{}
	return &this
}

// NewPriceModelTieredPercentageConfigWithDefaults instantiates a new PriceModelTieredPercentageConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceModelTieredPercentageConfigWithDefaults() *PriceModelTieredPercentageConfig {
	this := PriceModelTieredPercentageConfig{}
	return &this
}

// GetFirstUnit returns the FirstUnit field value if set, zero value otherwise.
func (o *PriceModelTieredPercentageConfig) GetFirstUnit() float32 {
	if o == nil || IsNil(o.FirstUnit) {
		var ret float32
		return ret
	}
	return *o.FirstUnit
}

// GetFirstUnitOk returns a tuple with the FirstUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelTieredPercentageConfig) GetFirstUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.FirstUnit) {
		return nil, false
	}
	return o.FirstUnit, true
}

// HasFirstUnit returns a boolean if a field has been set.
func (o *PriceModelTieredPercentageConfig) HasFirstUnit() bool {
	if o != nil && !IsNil(o.FirstUnit) {
		return true
	}

	return false
}

// SetFirstUnit gets a reference to the given float32 and assigns it to the FirstUnit field.
func (o *PriceModelTieredPercentageConfig) SetFirstUnit(v float32) {
	o.FirstUnit = &v
}

// GetFlatFee returns the FlatFee field value if set, zero value otherwise.
func (o *PriceModelTieredPercentageConfig) GetFlatFee() float32 {
	if o == nil || IsNil(o.FlatFee) {
		var ret float32
		return ret
	}
	return *o.FlatFee
}

// GetFlatFeeOk returns a tuple with the FlatFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelTieredPercentageConfig) GetFlatFeeOk() (*float32, bool) {
	if o == nil || IsNil(o.FlatFee) {
		return nil, false
	}
	return o.FlatFee, true
}

// HasFlatFee returns a boolean if a field has been set.
func (o *PriceModelTieredPercentageConfig) HasFlatFee() bool {
	if o != nil && !IsNil(o.FlatFee) {
		return true
	}

	return false
}

// SetFlatFee gets a reference to the given float32 and assigns it to the FlatFee field.
func (o *PriceModelTieredPercentageConfig) SetFlatFee(v float32) {
	o.FlatFee = &v
}

// GetLastUnit returns the LastUnit field value if set, zero value otherwise.
func (o *PriceModelTieredPercentageConfig) GetLastUnit() float32 {
	if o == nil || IsNil(o.LastUnit) {
		var ret float32
		return ret
	}
	return *o.LastUnit
}

// GetLastUnitOk returns a tuple with the LastUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelTieredPercentageConfig) GetLastUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.LastUnit) {
		return nil, false
	}
	return o.LastUnit, true
}

// HasLastUnit returns a boolean if a field has been set.
func (o *PriceModelTieredPercentageConfig) HasLastUnit() bool {
	if o != nil && !IsNil(o.LastUnit) {
		return true
	}

	return false
}

// SetLastUnit gets a reference to the given float32 and assigns it to the LastUnit field.
func (o *PriceModelTieredPercentageConfig) SetLastUnit(v float32) {
	o.LastUnit = &v
}

// GetPercentageRate returns the PercentageRate field value if set, zero value otherwise.
func (o *PriceModelTieredPercentageConfig) GetPercentageRate() float32 {
	if o == nil || IsNil(o.PercentageRate) {
		var ret float32
		return ret
	}
	return *o.PercentageRate
}

// GetPercentageRateOk returns a tuple with the PercentageRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelTieredPercentageConfig) GetPercentageRateOk() (*float32, bool) {
	if o == nil || IsNil(o.PercentageRate) {
		return nil, false
	}
	return o.PercentageRate, true
}

// HasPercentageRate returns a boolean if a field has been set.
func (o *PriceModelTieredPercentageConfig) HasPercentageRate() bool {
	if o != nil && !IsNil(o.PercentageRate) {
		return true
	}

	return false
}

// SetPercentageRate gets a reference to the given float32 and assigns it to the PercentageRate field.
func (o *PriceModelTieredPercentageConfig) SetPercentageRate(v float32) {
	o.PercentageRate = &v
}

func (o PriceModelTieredPercentageConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceModelTieredPercentageConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstUnit) {
		toSerialize["firstUnit"] = o.FirstUnit
	}
	if !IsNil(o.FlatFee) {
		toSerialize["flatFee"] = o.FlatFee
	}
	if !IsNil(o.LastUnit) {
		toSerialize["lastUnit"] = o.LastUnit
	}
	if !IsNil(o.PercentageRate) {
		toSerialize["percentageRate"] = o.PercentageRate
	}
	return toSerialize, nil
}

type NullablePriceModelTieredPercentageConfig struct {
	value *PriceModelTieredPercentageConfig
	isSet bool
}

func (v NullablePriceModelTieredPercentageConfig) Get() *PriceModelTieredPercentageConfig {
	return v.value
}

func (v *NullablePriceModelTieredPercentageConfig) Set(val *PriceModelTieredPercentageConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceModelTieredPercentageConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceModelTieredPercentageConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceModelTieredPercentageConfig(val *PriceModelTieredPercentageConfig) *NullablePriceModelTieredPercentageConfig {
	return &NullablePriceModelTieredPercentageConfig{value: val, isSet: true}
}

func (v NullablePriceModelTieredPercentageConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceModelTieredPercentageConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
