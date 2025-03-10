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

// checks if the GcpAmountUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpAmountUnit{}

// GcpAmountUnit struct for GcpAmountUnit
type GcpAmountUnit struct {
	Nanos *int32  `json:"nanos,omitempty"`
	Units *string `json:"units,omitempty"`
}

// NewGcpAmountUnit instantiates a new GcpAmountUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpAmountUnit() *GcpAmountUnit {
	this := GcpAmountUnit{}
	return &this
}

// NewGcpAmountUnitWithDefaults instantiates a new GcpAmountUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpAmountUnitWithDefaults() *GcpAmountUnit {
	this := GcpAmountUnit{}
	return &this
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *GcpAmountUnit) GetNanos() int32 {
	if o == nil || IsNil(o.Nanos) {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpAmountUnit) GetNanosOk() (*int32, bool) {
	if o == nil || IsNil(o.Nanos) {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *GcpAmountUnit) HasNanos() bool {
	if o != nil && !IsNil(o.Nanos) {
		return true
	}

	return false
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *GcpAmountUnit) SetNanos(v int32) {
	o.Nanos = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GcpAmountUnit) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpAmountUnit) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GcpAmountUnit) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *GcpAmountUnit) SetUnits(v string) {
	o.Units = &v
}

func (o GcpAmountUnit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpAmountUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Nanos) {
		toSerialize["nanos"] = o.Nanos
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableGcpAmountUnit struct {
	value *GcpAmountUnit
	isSet bool
}

func (v NullableGcpAmountUnit) Get() *GcpAmountUnit {
	return v.value
}

func (v *NullableGcpAmountUnit) Set(val *GcpAmountUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpAmountUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpAmountUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpAmountUnit(val *GcpAmountUnit) *NullableGcpAmountUnit {
	return &NullableGcpAmountUnit{value: val, isSet: true}
}

func (v NullableGcpAmountUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpAmountUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
