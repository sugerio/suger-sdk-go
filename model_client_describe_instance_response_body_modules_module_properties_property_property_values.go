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

// checks if the ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues{}

// ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues struct for ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues
type ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues struct {
	PropertyValue []ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue `json:"PropertyValue,omitempty"`
}

// NewClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues instantiates a new ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues() *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	this := ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues{}
	return &this
}

// NewClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesWithDefaults instantiates a new ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesWithDefaults() *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	this := ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues{}
	return &this
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) GetPropertyValue() []ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	if o == nil || IsNil(o.PropertyValue) {
		var ret []ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue
		return ret
	}
	return o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) GetPropertyValueOk() ([]ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue, bool) {
	if o == nil || IsNil(o.PropertyValue) {
		return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) HasPropertyValue() bool {
	if o != nil && !IsNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given []ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue and assigns it to the PropertyValue field.
func (o *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) SetPropertyValue(v []ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValuesPropertyValue) {
	o.PropertyValue = v
}

func (o ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyValue) {
		toSerialize["PropertyValue"] = o.PropertyValue
	}
	return toSerialize, nil
}

type NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues struct {
	value *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues
	isSet bool
}

func (v NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) Get() *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	return v.value
}

func (v *NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) Set(val *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) {
	v.value = val
	v.isSet = true
}

func (v NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) IsSet() bool {
	return v.isSet
}

func (v *NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues(val *ClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) *NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues {
	return &NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues{value: val, isSet: true}
}

func (v NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientDescribeInstanceResponseBodyModulesModulePropertiesPropertyPropertyValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
