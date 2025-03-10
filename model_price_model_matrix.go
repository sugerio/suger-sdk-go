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

// checks if the PriceModelMatrix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceModelMatrix{}

// PriceModelMatrix struct for PriceModelMatrix
type PriceModelMatrix struct {
	DefaultUnitAmount *float32 `json:"defaultUnitAmount,omitempty"`
	// The matrix of the pricing model
	Matrix []PriceModelMatrixConfigGroup `json:"matrix,omitempty"`
}

// NewPriceModelMatrix instantiates a new PriceModelMatrix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceModelMatrix() *PriceModelMatrix {
	this := PriceModelMatrix{}
	return &this
}

// NewPriceModelMatrixWithDefaults instantiates a new PriceModelMatrix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceModelMatrixWithDefaults() *PriceModelMatrix {
	this := PriceModelMatrix{}
	return &this
}

// GetDefaultUnitAmount returns the DefaultUnitAmount field value if set, zero value otherwise.
func (o *PriceModelMatrix) GetDefaultUnitAmount() float32 {
	if o == nil || IsNil(o.DefaultUnitAmount) {
		var ret float32
		return ret
	}
	return *o.DefaultUnitAmount
}

// GetDefaultUnitAmountOk returns a tuple with the DefaultUnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelMatrix) GetDefaultUnitAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DefaultUnitAmount) {
		return nil, false
	}
	return o.DefaultUnitAmount, true
}

// HasDefaultUnitAmount returns a boolean if a field has been set.
func (o *PriceModelMatrix) HasDefaultUnitAmount() bool {
	if o != nil && !IsNil(o.DefaultUnitAmount) {
		return true
	}

	return false
}

// SetDefaultUnitAmount gets a reference to the given float32 and assigns it to the DefaultUnitAmount field.
func (o *PriceModelMatrix) SetDefaultUnitAmount(v float32) {
	o.DefaultUnitAmount = &v
}

// GetMatrix returns the Matrix field value if set, zero value otherwise.
func (o *PriceModelMatrix) GetMatrix() []PriceModelMatrixConfigGroup {
	if o == nil || IsNil(o.Matrix) {
		var ret []PriceModelMatrixConfigGroup
		return ret
	}
	return o.Matrix
}

// GetMatrixOk returns a tuple with the Matrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceModelMatrix) GetMatrixOk() ([]PriceModelMatrixConfigGroup, bool) {
	if o == nil || IsNil(o.Matrix) {
		return nil, false
	}
	return o.Matrix, true
}

// HasMatrix returns a boolean if a field has been set.
func (o *PriceModelMatrix) HasMatrix() bool {
	if o != nil && !IsNil(o.Matrix) {
		return true
	}

	return false
}

// SetMatrix gets a reference to the given []PriceModelMatrixConfigGroup and assigns it to the Matrix field.
func (o *PriceModelMatrix) SetMatrix(v []PriceModelMatrixConfigGroup) {
	o.Matrix = v
}

func (o PriceModelMatrix) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceModelMatrix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultUnitAmount) {
		toSerialize["defaultUnitAmount"] = o.DefaultUnitAmount
	}
	if !IsNil(o.Matrix) {
		toSerialize["matrix"] = o.Matrix
	}
	return toSerialize, nil
}

type NullablePriceModelMatrix struct {
	value *PriceModelMatrix
	isSet bool
}

func (v NullablePriceModelMatrix) Get() *PriceModelMatrix {
	return v.value
}

func (v *NullablePriceModelMatrix) Set(val *PriceModelMatrix) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceModelMatrix) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceModelMatrix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceModelMatrix(val *PriceModelMatrix) *NullablePriceModelMatrix {
	return &NullablePriceModelMatrix{value: val, isSet: true}
}

func (v NullablePriceModelMatrix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceModelMatrix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
