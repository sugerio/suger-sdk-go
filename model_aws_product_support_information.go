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

// checks if the AwsProductSupportInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsProductSupportInformation{}

// AwsProductSupportInformation struct for AwsProductSupportInformation
type AwsProductSupportInformation struct {
	Description *string `json:"Description,omitempty"`
}

// NewAwsProductSupportInformation instantiates a new AwsProductSupportInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsProductSupportInformation() *AwsProductSupportInformation {
	this := AwsProductSupportInformation{}
	return &this
}

// NewAwsProductSupportInformationWithDefaults instantiates a new AwsProductSupportInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsProductSupportInformationWithDefaults() *AwsProductSupportInformation {
	this := AwsProductSupportInformation{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsProductSupportInformation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductSupportInformation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsProductSupportInformation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AwsProductSupportInformation) SetDescription(v string) {
	o.Description = &v
}

func (o AwsProductSupportInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsProductSupportInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAwsProductSupportInformation struct {
	value *AwsProductSupportInformation
	isSet bool
}

func (v NullableAwsProductSupportInformation) Get() *AwsProductSupportInformation {
	return v.value
}

func (v *NullableAwsProductSupportInformation) Set(val *AwsProductSupportInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsProductSupportInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsProductSupportInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsProductSupportInformation(val *AwsProductSupportInformation) *NullableAwsProductSupportInformation {
	return &NullableAwsProductSupportInformation{value: val, isSet: true}
}

func (v NullableAwsProductSupportInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsProductSupportInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
