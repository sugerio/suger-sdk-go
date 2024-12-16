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

// checks if the GcpMarketplaceProductTermsSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductTermsSpec{}

// GcpMarketplaceProductTermsSpec struct for GcpMarketplaceProductTermsSpec
type GcpMarketplaceProductTermsSpec struct {
	EulaUri *string `json:"eulaUri,omitempty"`
	// TODO: need to define the type
	InlineEula map[string]interface{} `json:"inlineEula,omitempty"`
	// TODO: need to define the type
	StandardEula map[string]interface{} `json:"standardEula,omitempty"`
}

// NewGcpMarketplaceProductTermsSpec instantiates a new GcpMarketplaceProductTermsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductTermsSpec() *GcpMarketplaceProductTermsSpec {
	this := GcpMarketplaceProductTermsSpec{}
	return &this
}

// NewGcpMarketplaceProductTermsSpecWithDefaults instantiates a new GcpMarketplaceProductTermsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductTermsSpecWithDefaults() *GcpMarketplaceProductTermsSpec {
	this := GcpMarketplaceProductTermsSpec{}
	return &this
}

// GetEulaUri returns the EulaUri field value if set, zero value otherwise.
func (o *GcpMarketplaceProductTermsSpec) GetEulaUri() string {
	if o == nil || IsNil(o.EulaUri) {
		var ret string
		return ret
	}
	return *o.EulaUri
}

// GetEulaUriOk returns a tuple with the EulaUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductTermsSpec) GetEulaUriOk() (*string, bool) {
	if o == nil || IsNil(o.EulaUri) {
		return nil, false
	}
	return o.EulaUri, true
}

// HasEulaUri returns a boolean if a field has been set.
func (o *GcpMarketplaceProductTermsSpec) HasEulaUri() bool {
	if o != nil && !IsNil(o.EulaUri) {
		return true
	}

	return false
}

// SetEulaUri gets a reference to the given string and assigns it to the EulaUri field.
func (o *GcpMarketplaceProductTermsSpec) SetEulaUri(v string) {
	o.EulaUri = &v
}

// GetInlineEula returns the InlineEula field value if set, zero value otherwise.
func (o *GcpMarketplaceProductTermsSpec) GetInlineEula() map[string]interface{} {
	if o == nil || IsNil(o.InlineEula) {
		var ret map[string]interface{}
		return ret
	}
	return o.InlineEula
}

// GetInlineEulaOk returns a tuple with the InlineEula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductTermsSpec) GetInlineEulaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.InlineEula) {
		return map[string]interface{}{}, false
	}
	return o.InlineEula, true
}

// HasInlineEula returns a boolean if a field has been set.
func (o *GcpMarketplaceProductTermsSpec) HasInlineEula() bool {
	if o != nil && !IsNil(o.InlineEula) {
		return true
	}

	return false
}

// SetInlineEula gets a reference to the given map[string]interface{} and assigns it to the InlineEula field.
func (o *GcpMarketplaceProductTermsSpec) SetInlineEula(v map[string]interface{}) {
	o.InlineEula = v
}

// GetStandardEula returns the StandardEula field value if set, zero value otherwise.
func (o *GcpMarketplaceProductTermsSpec) GetStandardEula() map[string]interface{} {
	if o == nil || IsNil(o.StandardEula) {
		var ret map[string]interface{}
		return ret
	}
	return o.StandardEula
}

// GetStandardEulaOk returns a tuple with the StandardEula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductTermsSpec) GetStandardEulaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StandardEula) {
		return map[string]interface{}{}, false
	}
	return o.StandardEula, true
}

// HasStandardEula returns a boolean if a field has been set.
func (o *GcpMarketplaceProductTermsSpec) HasStandardEula() bool {
	if o != nil && !IsNil(o.StandardEula) {
		return true
	}

	return false
}

// SetStandardEula gets a reference to the given map[string]interface{} and assigns it to the StandardEula field.
func (o *GcpMarketplaceProductTermsSpec) SetStandardEula(v map[string]interface{}) {
	o.StandardEula = v
}

func (o GcpMarketplaceProductTermsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductTermsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EulaUri) {
		toSerialize["eulaUri"] = o.EulaUri
	}
	if !IsNil(o.InlineEula) {
		toSerialize["inlineEula"] = o.InlineEula
	}
	if !IsNil(o.StandardEula) {
		toSerialize["standardEula"] = o.StandardEula
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductTermsSpec struct {
	value *GcpMarketplaceProductTermsSpec
	isSet bool
}

func (v NullableGcpMarketplaceProductTermsSpec) Get() *GcpMarketplaceProductTermsSpec {
	return v.value
}

func (v *NullableGcpMarketplaceProductTermsSpec) Set(val *GcpMarketplaceProductTermsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductTermsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductTermsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductTermsSpec(val *GcpMarketplaceProductTermsSpec) *NullableGcpMarketplaceProductTermsSpec {
	return &NullableGcpMarketplaceProductTermsSpec{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductTermsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductTermsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
