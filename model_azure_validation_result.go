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

// checks if the AzureValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureValidationResult{}

// AzureValidationResult struct for AzureValidationResult
type AzureValidationResult struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	MemberNames []string `json:"memberNames,omitempty"`
}

// NewAzureValidationResult instantiates a new AzureValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureValidationResult() *AzureValidationResult {
	this := AzureValidationResult{}
	return &this
}

// NewAzureValidationResultWithDefaults instantiates a new AzureValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureValidationResultWithDefaults() *AzureValidationResult {
	this := AzureValidationResult{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *AzureValidationResult) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureValidationResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *AzureValidationResult) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *AzureValidationResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetMemberNames returns the MemberNames field value if set, zero value otherwise.
func (o *AzureValidationResult) GetMemberNames() []string {
	if o == nil || IsNil(o.MemberNames) {
		var ret []string
		return ret
	}
	return o.MemberNames
}

// GetMemberNamesOk returns a tuple with the MemberNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureValidationResult) GetMemberNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.MemberNames) {
		return nil, false
	}
	return o.MemberNames, true
}

// HasMemberNames returns a boolean if a field has been set.
func (o *AzureValidationResult) HasMemberNames() bool {
	if o != nil && !IsNil(o.MemberNames) {
		return true
	}

	return false
}

// SetMemberNames gets a reference to the given []string and assigns it to the MemberNames field.
func (o *AzureValidationResult) SetMemberNames(v []string) {
	o.MemberNames = v
}

func (o AzureValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.MemberNames) {
		toSerialize["memberNames"] = o.MemberNames
	}
	return toSerialize, nil
}

type NullableAzureValidationResult struct {
	value *AzureValidationResult
	isSet bool
}

func (v NullableAzureValidationResult) Get() *AzureValidationResult {
	return v.value
}

func (v *NullableAzureValidationResult) Set(val *AzureValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureValidationResult(val *AzureValidationResult) *NullableAzureValidationResult {
	return &NullableAzureValidationResult{value: val, isSet: true}
}

func (v NullableAzureValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


