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

// checks if the PriceModelVolumeConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceModelVolumeConfig{}

// PriceModelVolumeConfig struct for PriceModelVolumeConfig
type PriceModelVolumeConfig struct {
	FlatFee *float32 `json:"flatFee,omitempty"`
	// Upper bound for this tier
	MaximumUnits *float32 `json:"maximumUnits,omitempty"`
	// Amount per unit
	UnitAmount *float32 `json:"unitAmount,omitempty"`
}

// NewPriceModelVolumeConfig instantiates a new PriceModelVolumeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceModelVolumeConfig() *PriceModelVolumeConfig {
	this := PriceModelVolumeConfig{}
	return &this
}

// NewPriceModelVolumeConfigWithDefaults instantiates a new PriceModelVolumeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceModelVolumeConfigWithDefaults() *PriceModelVolumeConfig {
	this := PriceModelVolumeConfig{}
	return &this
}

// GetFlatFee returns the FlatFee field value if set, zero value otherwise.
func (o *PriceModelVolumeConfig) GetFlatFee() float32 {
	if o == nil || IsNil(o.FlatFee) {
		var ret float32
		return ret
	}
	return *o.FlatFee
}

// GetFlatFeeOk returns a tuple with the FlatFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelVolumeConfig) GetFlatFeeOk() (*float32, bool) {
	if o == nil || IsNil(o.FlatFee) {
		return nil, false
	}
	return o.FlatFee, true
}

// HasFlatFee returns a boolean if a field has been set.
func (o *PriceModelVolumeConfig) HasFlatFee() bool {
	if o != nil && !IsNil(o.FlatFee) {
		return true
	}

	return false
}

// SetFlatFee gets a reference to the given float32 and assigns it to the FlatFee field.
func (o *PriceModelVolumeConfig) SetFlatFee(v float32) {
	o.FlatFee = &v
}

// GetMaximumUnits returns the MaximumUnits field value if set, zero value otherwise.
func (o *PriceModelVolumeConfig) GetMaximumUnits() float32 {
	if o == nil || IsNil(o.MaximumUnits) {
		var ret float32
		return ret
	}
	return *o.MaximumUnits
}

// GetMaximumUnitsOk returns a tuple with the MaximumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelVolumeConfig) GetMaximumUnitsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaximumUnits) {
		return nil, false
	}
	return o.MaximumUnits, true
}

// HasMaximumUnits returns a boolean if a field has been set.
func (o *PriceModelVolumeConfig) HasMaximumUnits() bool {
	if o != nil && !IsNil(o.MaximumUnits) {
		return true
	}

	return false
}

// SetMaximumUnits gets a reference to the given float32 and assigns it to the MaximumUnits field.
func (o *PriceModelVolumeConfig) SetMaximumUnits(v float32) {
	o.MaximumUnits = &v
}

// GetUnitAmount returns the UnitAmount field value if set, zero value otherwise.
func (o *PriceModelVolumeConfig) GetUnitAmount() float32 {
	if o == nil || IsNil(o.UnitAmount) {
		var ret float32
		return ret
	}
	return *o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelVolumeConfig) GetUnitAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitAmount) {
		return nil, false
	}
	return o.UnitAmount, true
}

// HasUnitAmount returns a boolean if a field has been set.
func (o *PriceModelVolumeConfig) HasUnitAmount() bool {
	if o != nil && !IsNil(o.UnitAmount) {
		return true
	}

	return false
}

// SetUnitAmount gets a reference to the given float32 and assigns it to the UnitAmount field.
func (o *PriceModelVolumeConfig) SetUnitAmount(v float32) {
	o.UnitAmount = &v
}

func (o PriceModelVolumeConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceModelVolumeConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlatFee) {
		toSerialize["flatFee"] = o.FlatFee
	}
	if !IsNil(o.MaximumUnits) {
		toSerialize["maximumUnits"] = o.MaximumUnits
	}
	if !IsNil(o.UnitAmount) {
		toSerialize["unitAmount"] = o.UnitAmount
	}
	return toSerialize, nil
}

type NullablePriceModelVolumeConfig struct {
	value *PriceModelVolumeConfig
	isSet bool
}

func (v NullablePriceModelVolumeConfig) Get() *PriceModelVolumeConfig {
	return v.value
}

func (v *NullablePriceModelVolumeConfig) Set(val *PriceModelVolumeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceModelVolumeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceModelVolumeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceModelVolumeConfig(val *PriceModelVolumeConfig) *NullablePriceModelVolumeConfig {
	return &NullablePriceModelVolumeConfig{value: val, isSet: true}
}

func (v NullablePriceModelVolumeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceModelVolumeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}