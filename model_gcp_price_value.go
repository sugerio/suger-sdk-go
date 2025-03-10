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

// checks if the GcpPriceValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpPriceValue{}

// GcpPriceValue struct for GcpPriceValue
type GcpPriceValue struct {
	// such as \"USD\"
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// for the decimal part, such as 30000000 = $0.03
	Nanos *int32 `json:"nanos,omitempty"`
	// for the integer part, such as \"500000\" = $50K
	Units *string `json:"units,omitempty"`
}

// NewGcpPriceValue instantiates a new GcpPriceValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpPriceValue() *GcpPriceValue {
	this := GcpPriceValue{}
	return &this
}

// NewGcpPriceValueWithDefaults instantiates a new GcpPriceValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpPriceValueWithDefaults() *GcpPriceValue {
	this := GcpPriceValue{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GcpPriceValue) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceValue) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GcpPriceValue) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GcpPriceValue) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *GcpPriceValue) GetNanos() int32 {
	if o == nil || IsNil(o.Nanos) {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceValue) GetNanosOk() (*int32, bool) {
	if o == nil || IsNil(o.Nanos) {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *GcpPriceValue) HasNanos() bool {
	if o != nil && !IsNil(o.Nanos) {
		return true
	}

	return false
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *GcpPriceValue) SetNanos(v int32) {
	o.Nanos = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GcpPriceValue) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpPriceValue) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GcpPriceValue) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *GcpPriceValue) SetUnits(v string) {
	o.Units = &v
}

func (o GcpPriceValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpPriceValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Nanos) {
		toSerialize["nanos"] = o.Nanos
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableGcpPriceValue struct {
	value *GcpPriceValue
	isSet bool
}

func (v NullableGcpPriceValue) Get() *GcpPriceValue {
	return v.value
}

func (v *NullableGcpPriceValue) Set(val *GcpPriceValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpPriceValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpPriceValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpPriceValue(val *GcpPriceValue) *NullableGcpPriceValue {
	return &NullableGcpPriceValue{value: val, isSet: true}
}

func (v NullableGcpPriceValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpPriceValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
