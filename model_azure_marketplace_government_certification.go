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

// checks if the AzureMarketplaceGovernmentCertification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceGovernmentCertification{}

// AzureMarketplaceGovernmentCertification struct for AzureMarketplaceGovernmentCertification
type AzureMarketplaceGovernmentCertification struct {
	// in patern of \"^(http|https)://\"
	Link *string `json:"link,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewAzureMarketplaceGovernmentCertification instantiates a new AzureMarketplaceGovernmentCertification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceGovernmentCertification() *AzureMarketplaceGovernmentCertification {
	this := AzureMarketplaceGovernmentCertification{}
	return &this
}

// NewAzureMarketplaceGovernmentCertificationWithDefaults instantiates a new AzureMarketplaceGovernmentCertification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceGovernmentCertificationWithDefaults() *AzureMarketplaceGovernmentCertification {
	this := AzureMarketplaceGovernmentCertification{}
	return &this
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *AzureMarketplaceGovernmentCertification) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceGovernmentCertification) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *AzureMarketplaceGovernmentCertification) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *AzureMarketplaceGovernmentCertification) SetLink(v string) {
	o.Link = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AzureMarketplaceGovernmentCertification) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceGovernmentCertification) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AzureMarketplaceGovernmentCertification) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AzureMarketplaceGovernmentCertification) SetName(v string) {
	o.Name = &v
}

func (o AzureMarketplaceGovernmentCertification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceGovernmentCertification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceGovernmentCertification struct {
	value *AzureMarketplaceGovernmentCertification
	isSet bool
}

func (v NullableAzureMarketplaceGovernmentCertification) Get() *AzureMarketplaceGovernmentCertification {
	return v.value
}

func (v *NullableAzureMarketplaceGovernmentCertification) Set(val *AzureMarketplaceGovernmentCertification) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceGovernmentCertification) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceGovernmentCertification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceGovernmentCertification(val *AzureMarketplaceGovernmentCertification) *NullableAzureMarketplaceGovernmentCertification {
	return &NullableAzureMarketplaceGovernmentCertification{value: val, isSet: true}
}

func (v NullableAzureMarketplaceGovernmentCertification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceGovernmentCertification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
