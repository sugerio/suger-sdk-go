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

// checks if the MicrosoftPartnerReferralLinkInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MicrosoftPartnerReferralLinkInfo{}

// MicrosoftPartnerReferralLinkInfo struct for MicrosoftPartnerReferralLinkInfo
type MicrosoftPartnerReferralLinkInfo struct {
	Method *string `json:"method,omitempty"`
	Uri    *string `json:"uri,omitempty"`
}

// NewMicrosoftPartnerReferralLinkInfo instantiates a new MicrosoftPartnerReferralLinkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftPartnerReferralLinkInfo() *MicrosoftPartnerReferralLinkInfo {
	this := MicrosoftPartnerReferralLinkInfo{}
	return &this
}

// NewMicrosoftPartnerReferralLinkInfoWithDefaults instantiates a new MicrosoftPartnerReferralLinkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftPartnerReferralLinkInfoWithDefaults() *MicrosoftPartnerReferralLinkInfo {
	this := MicrosoftPartnerReferralLinkInfo{}
	return &this
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralLinkInfo) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralLinkInfo) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralLinkInfo) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *MicrosoftPartnerReferralLinkInfo) SetMethod(v string) {
	o.Method = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *MicrosoftPartnerReferralLinkInfo) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftPartnerReferralLinkInfo) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *MicrosoftPartnerReferralLinkInfo) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *MicrosoftPartnerReferralLinkInfo) SetUri(v string) {
	o.Uri = &v
}

func (o MicrosoftPartnerReferralLinkInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MicrosoftPartnerReferralLinkInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableMicrosoftPartnerReferralLinkInfo struct {
	value *MicrosoftPartnerReferralLinkInfo
	isSet bool
}

func (v NullableMicrosoftPartnerReferralLinkInfo) Get() *MicrosoftPartnerReferralLinkInfo {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralLinkInfo) Set(val *MicrosoftPartnerReferralLinkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralLinkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralLinkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralLinkInfo(val *MicrosoftPartnerReferralLinkInfo) *NullableMicrosoftPartnerReferralLinkInfo {
	return &NullableMicrosoftPartnerReferralLinkInfo{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralLinkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralLinkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
