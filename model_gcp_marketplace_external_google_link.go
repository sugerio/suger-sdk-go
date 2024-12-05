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

// checks if the GcpMarketplaceExternalGoogleLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceExternalGoogleLink{}

// GcpMarketplaceExternalGoogleLink struct for GcpMarketplaceExternalGoogleLink
type GcpMarketplaceExternalGoogleLink struct {
	Uri *string `json:"uri,omitempty"`
}

// NewGcpMarketplaceExternalGoogleLink instantiates a new GcpMarketplaceExternalGoogleLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceExternalGoogleLink() *GcpMarketplaceExternalGoogleLink {
	this := GcpMarketplaceExternalGoogleLink{}
	return &this
}

// NewGcpMarketplaceExternalGoogleLinkWithDefaults instantiates a new GcpMarketplaceExternalGoogleLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceExternalGoogleLinkWithDefaults() *GcpMarketplaceExternalGoogleLink {
	this := GcpMarketplaceExternalGoogleLink{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *GcpMarketplaceExternalGoogleLink) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceExternalGoogleLink) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *GcpMarketplaceExternalGoogleLink) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *GcpMarketplaceExternalGoogleLink) SetUri(v string) {
	o.Uri = &v
}

func (o GcpMarketplaceExternalGoogleLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceExternalGoogleLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceExternalGoogleLink struct {
	value *GcpMarketplaceExternalGoogleLink
	isSet bool
}

func (v NullableGcpMarketplaceExternalGoogleLink) Get() *GcpMarketplaceExternalGoogleLink {
	return v.value
}

func (v *NullableGcpMarketplaceExternalGoogleLink) Set(val *GcpMarketplaceExternalGoogleLink) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceExternalGoogleLink) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceExternalGoogleLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceExternalGoogleLink(val *GcpMarketplaceExternalGoogleLink) *NullableGcpMarketplaceExternalGoogleLink {
	return &NullableGcpMarketplaceExternalGoogleLink{value: val, isSet: true}
}

func (v NullableGcpMarketplaceExternalGoogleLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceExternalGoogleLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
