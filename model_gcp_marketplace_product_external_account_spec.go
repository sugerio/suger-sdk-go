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

// checks if the GcpMarketplaceProductExternalAccountSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductExternalAccountSpec{}

// GcpMarketplaceProductExternalAccountSpec struct for GcpMarketplaceProductExternalAccountSpec
type GcpMarketplaceProductExternalAccountSpec struct {
	LoginUri *string `json:"loginUri,omitempty"`
	SignupUri *string `json:"signupUri,omitempty"`
	SingleSignOnUri *string `json:"singleSignOnUri,omitempty"`
}

// NewGcpMarketplaceProductExternalAccountSpec instantiates a new GcpMarketplaceProductExternalAccountSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductExternalAccountSpec() *GcpMarketplaceProductExternalAccountSpec {
	this := GcpMarketplaceProductExternalAccountSpec{}
	return &this
}

// NewGcpMarketplaceProductExternalAccountSpecWithDefaults instantiates a new GcpMarketplaceProductExternalAccountSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductExternalAccountSpecWithDefaults() *GcpMarketplaceProductExternalAccountSpec {
	this := GcpMarketplaceProductExternalAccountSpec{}
	return &this
}

// GetLoginUri returns the LoginUri field value if set, zero value otherwise.
func (o *GcpMarketplaceProductExternalAccountSpec) GetLoginUri() string {
	if o == nil || IsNil(o.LoginUri) {
		var ret string
		return ret
	}
	return *o.LoginUri
}

// GetLoginUriOk returns a tuple with the LoginUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductExternalAccountSpec) GetLoginUriOk() (*string, bool) {
	if o == nil || IsNil(o.LoginUri) {
		return nil, false
	}
	return o.LoginUri, true
}

// HasLoginUri returns a boolean if a field has been set.
func (o *GcpMarketplaceProductExternalAccountSpec) HasLoginUri() bool {
	if o != nil && !IsNil(o.LoginUri) {
		return true
	}

	return false
}

// SetLoginUri gets a reference to the given string and assigns it to the LoginUri field.
func (o *GcpMarketplaceProductExternalAccountSpec) SetLoginUri(v string) {
	o.LoginUri = &v
}

// GetSignupUri returns the SignupUri field value if set, zero value otherwise.
func (o *GcpMarketplaceProductExternalAccountSpec) GetSignupUri() string {
	if o == nil || IsNil(o.SignupUri) {
		var ret string
		return ret
	}
	return *o.SignupUri
}

// GetSignupUriOk returns a tuple with the SignupUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductExternalAccountSpec) GetSignupUriOk() (*string, bool) {
	if o == nil || IsNil(o.SignupUri) {
		return nil, false
	}
	return o.SignupUri, true
}

// HasSignupUri returns a boolean if a field has been set.
func (o *GcpMarketplaceProductExternalAccountSpec) HasSignupUri() bool {
	if o != nil && !IsNil(o.SignupUri) {
		return true
	}

	return false
}

// SetSignupUri gets a reference to the given string and assigns it to the SignupUri field.
func (o *GcpMarketplaceProductExternalAccountSpec) SetSignupUri(v string) {
	o.SignupUri = &v
}

// GetSingleSignOnUri returns the SingleSignOnUri field value if set, zero value otherwise.
func (o *GcpMarketplaceProductExternalAccountSpec) GetSingleSignOnUri() string {
	if o == nil || IsNil(o.SingleSignOnUri) {
		var ret string
		return ret
	}
	return *o.SingleSignOnUri
}

// GetSingleSignOnUriOk returns a tuple with the SingleSignOnUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductExternalAccountSpec) GetSingleSignOnUriOk() (*string, bool) {
	if o == nil || IsNil(o.SingleSignOnUri) {
		return nil, false
	}
	return o.SingleSignOnUri, true
}

// HasSingleSignOnUri returns a boolean if a field has been set.
func (o *GcpMarketplaceProductExternalAccountSpec) HasSingleSignOnUri() bool {
	if o != nil && !IsNil(o.SingleSignOnUri) {
		return true
	}

	return false
}

// SetSingleSignOnUri gets a reference to the given string and assigns it to the SingleSignOnUri field.
func (o *GcpMarketplaceProductExternalAccountSpec) SetSingleSignOnUri(v string) {
	o.SingleSignOnUri = &v
}

func (o GcpMarketplaceProductExternalAccountSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductExternalAccountSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoginUri) {
		toSerialize["loginUri"] = o.LoginUri
	}
	if !IsNil(o.SignupUri) {
		toSerialize["signupUri"] = o.SignupUri
	}
	if !IsNil(o.SingleSignOnUri) {
		toSerialize["singleSignOnUri"] = o.SingleSignOnUri
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductExternalAccountSpec struct {
	value *GcpMarketplaceProductExternalAccountSpec
	isSet bool
}

func (v NullableGcpMarketplaceProductExternalAccountSpec) Get() *GcpMarketplaceProductExternalAccountSpec {
	return v.value
}

func (v *NullableGcpMarketplaceProductExternalAccountSpec) Set(val *GcpMarketplaceProductExternalAccountSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductExternalAccountSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductExternalAccountSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductExternalAccountSpec(val *GcpMarketplaceProductExternalAccountSpec) *NullableGcpMarketplaceProductExternalAccountSpec {
	return &NullableGcpMarketplaceProductExternalAccountSpec{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductExternalAccountSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductExternalAccountSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


