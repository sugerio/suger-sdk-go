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

// checks if the UsageMeteringDimensionMappingValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageMeteringDimensionMappingValue{}

// UsageMeteringDimensionMappingValue struct for UsageMeteringDimensionMappingValue
type UsageMeteringDimensionMappingValue struct {
	// The convertion multiplier when mapping from the source dimension key to the destination dimensionKey by quantity mode. Not required if the mapping mode is AMOUNT.
	ConvertionMultiplier *float32 `json:"convertionMultiplier,omitempty"`
	// The destination dimension key of the usage metering mapping.
	DimensionKey *string `json:"dimensionKey,omitempty"`
	// The conversion mode of UsageMeteringDimensionMapping. The default is QUANTITY if not available.
	MappingMode *UsageMeteringDimensionMappingMode `json:"mappingMode,omitempty"`
}

// NewUsageMeteringDimensionMappingValue instantiates a new UsageMeteringDimensionMappingValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMeteringDimensionMappingValue() *UsageMeteringDimensionMappingValue {
	this := UsageMeteringDimensionMappingValue{}
	return &this
}

// NewUsageMeteringDimensionMappingValueWithDefaults instantiates a new UsageMeteringDimensionMappingValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMeteringDimensionMappingValueWithDefaults() *UsageMeteringDimensionMappingValue {
	this := UsageMeteringDimensionMappingValue{}
	return &this
}

// GetConvertionMultiplier returns the ConvertionMultiplier field value if set, zero value otherwise.
func (o *UsageMeteringDimensionMappingValue) GetConvertionMultiplier() float32 {
	if o == nil || IsNil(o.ConvertionMultiplier) {
		var ret float32
		return ret
	}
	return *o.ConvertionMultiplier
}

// GetConvertionMultiplierOk returns a tuple with the ConvertionMultiplier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDimensionMappingValue) GetConvertionMultiplierOk() (*float32, bool) {
	if o == nil || IsNil(o.ConvertionMultiplier) {
		return nil, false
	}
	return o.ConvertionMultiplier, true
}

// HasConvertionMultiplier returns a boolean if a field has been set.
func (o *UsageMeteringDimensionMappingValue) HasConvertionMultiplier() bool {
	if o != nil && !IsNil(o.ConvertionMultiplier) {
		return true
	}

	return false
}

// SetConvertionMultiplier gets a reference to the given float32 and assigns it to the ConvertionMultiplier field.
func (o *UsageMeteringDimensionMappingValue) SetConvertionMultiplier(v float32) {
	o.ConvertionMultiplier = &v
}

// GetDimensionKey returns the DimensionKey field value if set, zero value otherwise.
func (o *UsageMeteringDimensionMappingValue) GetDimensionKey() string {
	if o == nil || IsNil(o.DimensionKey) {
		var ret string
		return ret
	}
	return *o.DimensionKey
}

// GetDimensionKeyOk returns a tuple with the DimensionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDimensionMappingValue) GetDimensionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionKey) {
		return nil, false
	}
	return o.DimensionKey, true
}

// HasDimensionKey returns a boolean if a field has been set.
func (o *UsageMeteringDimensionMappingValue) HasDimensionKey() bool {
	if o != nil && !IsNil(o.DimensionKey) {
		return true
	}

	return false
}

// SetDimensionKey gets a reference to the given string and assigns it to the DimensionKey field.
func (o *UsageMeteringDimensionMappingValue) SetDimensionKey(v string) {
	o.DimensionKey = &v
}

// GetMappingMode returns the MappingMode field value if set, zero value otherwise.
func (o *UsageMeteringDimensionMappingValue) GetMappingMode() UsageMeteringDimensionMappingMode {
	if o == nil || IsNil(o.MappingMode) {
		var ret UsageMeteringDimensionMappingMode
		return ret
	}
	return *o.MappingMode
}

// GetMappingModeOk returns a tuple with the MappingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMeteringDimensionMappingValue) GetMappingModeOk() (*UsageMeteringDimensionMappingMode, bool) {
	if o == nil || IsNil(o.MappingMode) {
		return nil, false
	}
	return o.MappingMode, true
}

// HasMappingMode returns a boolean if a field has been set.
func (o *UsageMeteringDimensionMappingValue) HasMappingMode() bool {
	if o != nil && !IsNil(o.MappingMode) {
		return true
	}

	return false
}

// SetMappingMode gets a reference to the given UsageMeteringDimensionMappingMode and assigns it to the MappingMode field.
func (o *UsageMeteringDimensionMappingValue) SetMappingMode(v UsageMeteringDimensionMappingMode) {
	o.MappingMode = &v
}

func (o UsageMeteringDimensionMappingValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageMeteringDimensionMappingValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConvertionMultiplier) {
		toSerialize["convertionMultiplier"] = o.ConvertionMultiplier
	}
	if !IsNil(o.DimensionKey) {
		toSerialize["dimensionKey"] = o.DimensionKey
	}
	if !IsNil(o.MappingMode) {
		toSerialize["mappingMode"] = o.MappingMode
	}
	return toSerialize, nil
}

type NullableUsageMeteringDimensionMappingValue struct {
	value *UsageMeteringDimensionMappingValue
	isSet bool
}

func (v NullableUsageMeteringDimensionMappingValue) Get() *UsageMeteringDimensionMappingValue {
	return v.value
}

func (v *NullableUsageMeteringDimensionMappingValue) Set(val *UsageMeteringDimensionMappingValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMeteringDimensionMappingValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMeteringDimensionMappingValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMeteringDimensionMappingValue(val *UsageMeteringDimensionMappingValue) *NullableUsageMeteringDimensionMappingValue {
	return &NullableUsageMeteringDimensionMappingValue{value: val, isSet: true}
}

func (v NullableUsageMeteringDimensionMappingValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMeteringDimensionMappingValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
