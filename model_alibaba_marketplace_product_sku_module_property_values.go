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

// checks if the AlibabaMarketplaceProductSkuModulePropertyValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlibabaMarketplaceProductSkuModulePropertyValues{}

// AlibabaMarketplaceProductSkuModulePropertyValues struct for AlibabaMarketplaceProductSkuModulePropertyValues
type AlibabaMarketplaceProductSkuModulePropertyValues struct {
	PropertyValue []AlibabaMarketplaceProductSkuModulePropertyValue `json:"PropertyValue,omitempty"`
}

// NewAlibabaMarketplaceProductSkuModulePropertyValues instantiates a new AlibabaMarketplaceProductSkuModulePropertyValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlibabaMarketplaceProductSkuModulePropertyValues() *AlibabaMarketplaceProductSkuModulePropertyValues {
	this := AlibabaMarketplaceProductSkuModulePropertyValues{}
	return &this
}

// NewAlibabaMarketplaceProductSkuModulePropertyValuesWithDefaults instantiates a new AlibabaMarketplaceProductSkuModulePropertyValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlibabaMarketplaceProductSkuModulePropertyValuesWithDefaults() *AlibabaMarketplaceProductSkuModulePropertyValues {
	this := AlibabaMarketplaceProductSkuModulePropertyValues{}
	return &this
}

// GetPropertyValue returns the PropertyValue field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductSkuModulePropertyValues) GetPropertyValue() []AlibabaMarketplaceProductSkuModulePropertyValue {
	if o == nil || IsNil(o.PropertyValue) {
		var ret []AlibabaMarketplaceProductSkuModulePropertyValue
		return ret
	}
	return o.PropertyValue
}

// GetPropertyValueOk returns a tuple with the PropertyValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductSkuModulePropertyValues) GetPropertyValueOk() ([]AlibabaMarketplaceProductSkuModulePropertyValue, bool) {
	if o == nil || IsNil(o.PropertyValue) {
		return nil, false
	}
	return o.PropertyValue, true
}

// HasPropertyValue returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductSkuModulePropertyValues) HasPropertyValue() bool {
	if o != nil && !IsNil(o.PropertyValue) {
		return true
	}

	return false
}

// SetPropertyValue gets a reference to the given []AlibabaMarketplaceProductSkuModulePropertyValue and assigns it to the PropertyValue field.
func (o *AlibabaMarketplaceProductSkuModulePropertyValues) SetPropertyValue(v []AlibabaMarketplaceProductSkuModulePropertyValue) {
	o.PropertyValue = v
}

func (o AlibabaMarketplaceProductSkuModulePropertyValues) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlibabaMarketplaceProductSkuModulePropertyValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyValue) {
		toSerialize["PropertyValue"] = o.PropertyValue
	}
	return toSerialize, nil
}

type NullableAlibabaMarketplaceProductSkuModulePropertyValues struct {
	value *AlibabaMarketplaceProductSkuModulePropertyValues
	isSet bool
}

func (v NullableAlibabaMarketplaceProductSkuModulePropertyValues) Get() *AlibabaMarketplaceProductSkuModulePropertyValues {
	return v.value
}

func (v *NullableAlibabaMarketplaceProductSkuModulePropertyValues) Set(val *AlibabaMarketplaceProductSkuModulePropertyValues) {
	v.value = val
	v.isSet = true
}

func (v NullableAlibabaMarketplaceProductSkuModulePropertyValues) IsSet() bool {
	return v.isSet
}

func (v *NullableAlibabaMarketplaceProductSkuModulePropertyValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlibabaMarketplaceProductSkuModulePropertyValues(val *AlibabaMarketplaceProductSkuModulePropertyValues) *NullableAlibabaMarketplaceProductSkuModulePropertyValues {
	return &NullableAlibabaMarketplaceProductSkuModulePropertyValues{value: val, isSet: true}
}

func (v NullableAlibabaMarketplaceProductSkuModulePropertyValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlibabaMarketplaceProductSkuModulePropertyValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
