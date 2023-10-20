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

// checks if the GcpMarketplaceProductSupportSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductSupportSpec{}

// GcpMarketplaceProductSupportSpec struct for GcpMarketplaceProductSupportSpec
type GcpMarketplaceProductSupportSpec struct {
	Description *string `json:"description,omitempty"`
	Email *string `json:"email,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewGcpMarketplaceProductSupportSpec instantiates a new GcpMarketplaceProductSupportSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductSupportSpec() *GcpMarketplaceProductSupportSpec {
	this := GcpMarketplaceProductSupportSpec{}
	return &this
}

// NewGcpMarketplaceProductSupportSpecWithDefaults instantiates a new GcpMarketplaceProductSupportSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductSupportSpecWithDefaults() *GcpMarketplaceProductSupportSpec {
	this := GcpMarketplaceProductSupportSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GcpMarketplaceProductSupportSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductSupportSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GcpMarketplaceProductSupportSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GcpMarketplaceProductSupportSpec) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GcpMarketplaceProductSupportSpec) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductSupportSpec) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GcpMarketplaceProductSupportSpec) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GcpMarketplaceProductSupportSpec) SetEmail(v string) {
	o.Email = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *GcpMarketplaceProductSupportSpec) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductSupportSpec) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *GcpMarketplaceProductSupportSpec) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *GcpMarketplaceProductSupportSpec) SetUri(v string) {
	o.Uri = &v
}

func (o GcpMarketplaceProductSupportSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductSupportSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductSupportSpec struct {
	value *GcpMarketplaceProductSupportSpec
	isSet bool
}

func (v NullableGcpMarketplaceProductSupportSpec) Get() *GcpMarketplaceProductSupportSpec {
	return v.value
}

func (v *NullableGcpMarketplaceProductSupportSpec) Set(val *GcpMarketplaceProductSupportSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductSupportSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductSupportSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductSupportSpec(val *GcpMarketplaceProductSupportSpec) *NullableGcpMarketplaceProductSupportSpec {
	return &NullableGcpMarketplaceProductSupportSpec{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductSupportSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductSupportSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


