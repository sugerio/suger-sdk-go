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

// checks if the PriceModelMatrixConfigGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceModelMatrixConfigGroup{}

// PriceModelMatrixConfigGroup struct for PriceModelMatrixConfigGroup
type PriceModelMatrixConfigGroup struct {
	Properties []PriceModelMatrixProperty `json:"properties,omitempty"`
	UnitAmount *float32                   `json:"unitAmount,omitempty"`
}

// NewPriceModelMatrixConfigGroup instantiates a new PriceModelMatrixConfigGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceModelMatrixConfigGroup() *PriceModelMatrixConfigGroup {
	this := PriceModelMatrixConfigGroup{}
	return &this
}

// NewPriceModelMatrixConfigGroupWithDefaults instantiates a new PriceModelMatrixConfigGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceModelMatrixConfigGroupWithDefaults() *PriceModelMatrixConfigGroup {
	this := PriceModelMatrixConfigGroup{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PriceModelMatrixConfigGroup) GetProperties() []PriceModelMatrixProperty {
	if o == nil || IsNil(o.Properties) {
		var ret []PriceModelMatrixProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelMatrixConfigGroup) GetPropertiesOk() ([]PriceModelMatrixProperty, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PriceModelMatrixConfigGroup) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []PriceModelMatrixProperty and assigns it to the Properties field.
func (o *PriceModelMatrixConfigGroup) SetProperties(v []PriceModelMatrixProperty) {
	o.Properties = v
}

// GetUnitAmount returns the UnitAmount field value if set, zero value otherwise.
func (o *PriceModelMatrixConfigGroup) GetUnitAmount() float32 {
	if o == nil || IsNil(o.UnitAmount) {
		var ret float32
		return ret
	}
	return *o.UnitAmount
}

// GetUnitAmountOk returns a tuple with the UnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelMatrixConfigGroup) GetUnitAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitAmount) {
		return nil, false
	}
	return o.UnitAmount, true
}

// HasUnitAmount returns a boolean if a field has been set.
func (o *PriceModelMatrixConfigGroup) HasUnitAmount() bool {
	if o != nil && !IsNil(o.UnitAmount) {
		return true
	}

	return false
}

// SetUnitAmount gets a reference to the given float32 and assigns it to the UnitAmount field.
func (o *PriceModelMatrixConfigGroup) SetUnitAmount(v float32) {
	o.UnitAmount = &v
}

func (o PriceModelMatrixConfigGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceModelMatrixConfigGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.UnitAmount) {
		toSerialize["unitAmount"] = o.UnitAmount
	}
	return toSerialize, nil
}

type NullablePriceModelMatrixConfigGroup struct {
	value *PriceModelMatrixConfigGroup
	isSet bool
}

func (v NullablePriceModelMatrixConfigGroup) Get() *PriceModelMatrixConfigGroup {
	return v.value
}

func (v *NullablePriceModelMatrixConfigGroup) Set(val *PriceModelMatrixConfigGroup) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceModelMatrixConfigGroup) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceModelMatrixConfigGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceModelMatrixConfigGroup(val *PriceModelMatrixConfigGroup) *NullablePriceModelMatrixConfigGroup {
	return &NullablePriceModelMatrixConfigGroup{value: val, isSet: true}
}

func (v NullablePriceModelMatrixConfigGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceModelMatrixConfigGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}