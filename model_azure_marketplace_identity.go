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

// checks if the AzureMarketplaceIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceIdentity{}

// AzureMarketplaceIdentity struct for AzureMarketplaceIdentity
type AzureMarketplaceIdentity struct {
	ExternalId *string `json:"externalId,omitempty"`
}

// NewAzureMarketplaceIdentity instantiates a new AzureMarketplaceIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceIdentity() *AzureMarketplaceIdentity {
	this := AzureMarketplaceIdentity{}
	return &this
}

// NewAzureMarketplaceIdentityWithDefaults instantiates a new AzureMarketplaceIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceIdentityWithDefaults() *AzureMarketplaceIdentity {
	this := AzureMarketplaceIdentity{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *AzureMarketplaceIdentity) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceIdentity) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *AzureMarketplaceIdentity) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *AzureMarketplaceIdentity) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o AzureMarketplaceIdentity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceIdentity struct {
	value *AzureMarketplaceIdentity
	isSet bool
}

func (v NullableAzureMarketplaceIdentity) Get() *AzureMarketplaceIdentity {
	return v.value
}

func (v *NullableAzureMarketplaceIdentity) Set(val *AzureMarketplaceIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceIdentity(val *AzureMarketplaceIdentity) *NullableAzureMarketplaceIdentity {
	return &NullableAzureMarketplaceIdentity{value: val, isSet: true}
}

func (v NullableAzureMarketplaceIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
