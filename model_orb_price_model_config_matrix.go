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

// checks if the OrbPriceModelConfigMATRIX type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbPriceModelConfigMATRIX{}

// OrbPriceModelConfigMATRIX struct for OrbPriceModelConfigMATRIX
type OrbPriceModelConfigMATRIX struct {
	DefaultUnitAmount *string `json:"default_unit_amount,omitempty"`
	// First and (optional) second dimension grouping keys
	Dimensions   []string              `json:"dimensions,omitempty"`
	MatrixValues []OrbMatrixPriceValue `json:"matrix_values,omitempty"`
}

// NewOrbPriceModelConfigMATRIX instantiates a new OrbPriceModelConfigMATRIX object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbPriceModelConfigMATRIX() *OrbPriceModelConfigMATRIX {
	this := OrbPriceModelConfigMATRIX{}
	return &this
}

// NewOrbPriceModelConfigMATRIXWithDefaults instantiates a new OrbPriceModelConfigMATRIX object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbPriceModelConfigMATRIXWithDefaults() *OrbPriceModelConfigMATRIX {
	this := OrbPriceModelConfigMATRIX{}
	return &this
}

// GetDefaultUnitAmount returns the DefaultUnitAmount field value if set, zero value otherwise.
func (o *OrbPriceModelConfigMATRIX) GetDefaultUnitAmount() string {
	if o == nil || IsNil(o.DefaultUnitAmount) {
		var ret string
		return ret
	}
	return *o.DefaultUnitAmount
}

// GetDefaultUnitAmountOk returns a tuple with the DefaultUnitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigMATRIX) GetDefaultUnitAmountOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultUnitAmount) {
		return nil, false
	}
	return o.DefaultUnitAmount, true
}

// HasDefaultUnitAmount returns a boolean if a field has been set.
func (o *OrbPriceModelConfigMATRIX) HasDefaultUnitAmount() bool {
	if o != nil && !IsNil(o.DefaultUnitAmount) {
		return true
	}

	return false
}

// SetDefaultUnitAmount gets a reference to the given string and assigns it to the DefaultUnitAmount field.
func (o *OrbPriceModelConfigMATRIX) SetDefaultUnitAmount(v string) {
	o.DefaultUnitAmount = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *OrbPriceModelConfigMATRIX) GetDimensions() []string {
	if o == nil || IsNil(o.Dimensions) {
		var ret []string
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigMATRIX) GetDimensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *OrbPriceModelConfigMATRIX) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []string and assigns it to the Dimensions field.
func (o *OrbPriceModelConfigMATRIX) SetDimensions(v []string) {
	o.Dimensions = v
}

// GetMatrixValues returns the MatrixValues field value if set, zero value otherwise.
func (o *OrbPriceModelConfigMATRIX) GetMatrixValues() []OrbMatrixPriceValue {
	if o == nil || IsNil(o.MatrixValues) {
		var ret []OrbMatrixPriceValue
		return ret
	}
	return o.MatrixValues
}

// GetMatrixValuesOk returns a tuple with the MatrixValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPriceModelConfigMATRIX) GetMatrixValuesOk() ([]OrbMatrixPriceValue, bool) {
	if o == nil || IsNil(o.MatrixValues) {
		return nil, false
	}
	return o.MatrixValues, true
}

// HasMatrixValues returns a boolean if a field has been set.
func (o *OrbPriceModelConfigMATRIX) HasMatrixValues() bool {
	if o != nil && !IsNil(o.MatrixValues) {
		return true
	}

	return false
}

// SetMatrixValues gets a reference to the given []OrbMatrixPriceValue and assigns it to the MatrixValues field.
func (o *OrbPriceModelConfigMATRIX) SetMatrixValues(v []OrbMatrixPriceValue) {
	o.MatrixValues = v
}

func (o OrbPriceModelConfigMATRIX) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbPriceModelConfigMATRIX) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultUnitAmount) {
		toSerialize["default_unit_amount"] = o.DefaultUnitAmount
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.MatrixValues) {
		toSerialize["matrix_values"] = o.MatrixValues
	}
	return toSerialize, nil
}

type NullableOrbPriceModelConfigMATRIX struct {
	value *OrbPriceModelConfigMATRIX
	isSet bool
}

func (v NullableOrbPriceModelConfigMATRIX) Get() *OrbPriceModelConfigMATRIX {
	return v.value
}

func (v *NullableOrbPriceModelConfigMATRIX) Set(val *OrbPriceModelConfigMATRIX) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPriceModelConfigMATRIX) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPriceModelConfigMATRIX) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPriceModelConfigMATRIX(val *OrbPriceModelConfigMATRIX) *NullableOrbPriceModelConfigMATRIX {
	return &NullableOrbPriceModelConfigMATRIX{value: val, isSet: true}
}

func (v NullableOrbPriceModelConfigMATRIX) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPriceModelConfigMATRIX) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
