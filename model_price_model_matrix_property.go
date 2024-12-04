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

// checks if the PriceModelMatrixProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceModelMatrixProperty{}

// PriceModelMatrixProperty struct for PriceModelMatrixProperty
type PriceModelMatrixProperty struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewPriceModelMatrixProperty instantiates a new PriceModelMatrixProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceModelMatrixProperty() *PriceModelMatrixProperty {
	this := PriceModelMatrixProperty{}
	return &this
}

// NewPriceModelMatrixPropertyWithDefaults instantiates a new PriceModelMatrixProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceModelMatrixPropertyWithDefaults() *PriceModelMatrixProperty {
	this := PriceModelMatrixProperty{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PriceModelMatrixProperty) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelMatrixProperty) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PriceModelMatrixProperty) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PriceModelMatrixProperty) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PriceModelMatrixProperty) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelMatrixProperty) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PriceModelMatrixProperty) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PriceModelMatrixProperty) SetValue(v string) {
	o.Value = &v
}

func (o PriceModelMatrixProperty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceModelMatrixProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePriceModelMatrixProperty struct {
	value *PriceModelMatrixProperty
	isSet bool
}

func (v NullablePriceModelMatrixProperty) Get() *PriceModelMatrixProperty {
	return v.value
}

func (v *NullablePriceModelMatrixProperty) Set(val *PriceModelMatrixProperty) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceModelMatrixProperty) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceModelMatrixProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceModelMatrixProperty(val *PriceModelMatrixProperty) *NullablePriceModelMatrixProperty {
	return &NullablePriceModelMatrixProperty{value: val, isSet: true}
}

func (v NullablePriceModelMatrixProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceModelMatrixProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}