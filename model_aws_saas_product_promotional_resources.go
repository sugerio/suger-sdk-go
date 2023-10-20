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

// checks if the AwsSaasProductPromotionalResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsSaasProductPromotionalResources{}

// AwsSaasProductPromotionalResources struct for AwsSaasProductPromotionalResources
type AwsSaasProductPromotionalResources struct {
	AdditionalResources []AwsSaasProductAdditionalResource `json:"AdditionalResources,omitempty"`
	LogoUrl *string `json:"LogoUrl,omitempty"`
	// Currently, AWS only support 1 url in the array.
	VideoUrls []string `json:"VideoUrls,omitempty"`
}

// NewAwsSaasProductPromotionalResources instantiates a new AwsSaasProductPromotionalResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSaasProductPromotionalResources() *AwsSaasProductPromotionalResources {
	this := AwsSaasProductPromotionalResources{}
	return &this
}

// NewAwsSaasProductPromotionalResourcesWithDefaults instantiates a new AwsSaasProductPromotionalResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSaasProductPromotionalResourcesWithDefaults() *AwsSaasProductPromotionalResources {
	this := AwsSaasProductPromotionalResources{}
	return &this
}

// GetAdditionalResources returns the AdditionalResources field value if set, zero value otherwise.
func (o *AwsSaasProductPromotionalResources) GetAdditionalResources() []AwsSaasProductAdditionalResource {
	if o == nil || IsNil(o.AdditionalResources) {
		var ret []AwsSaasProductAdditionalResource
		return ret
	}
	return o.AdditionalResources
}

// GetAdditionalResourcesOk returns a tuple with the AdditionalResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSaasProductPromotionalResources) GetAdditionalResourcesOk() ([]AwsSaasProductAdditionalResource, bool) {
	if o == nil || IsNil(o.AdditionalResources) {
		return nil, false
	}
	return o.AdditionalResources, true
}

// HasAdditionalResources returns a boolean if a field has been set.
func (o *AwsSaasProductPromotionalResources) HasAdditionalResources() bool {
	if o != nil && !IsNil(o.AdditionalResources) {
		return true
	}

	return false
}

// SetAdditionalResources gets a reference to the given []AwsSaasProductAdditionalResource and assigns it to the AdditionalResources field.
func (o *AwsSaasProductPromotionalResources) SetAdditionalResources(v []AwsSaasProductAdditionalResource) {
	o.AdditionalResources = v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AwsSaasProductPromotionalResources) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSaasProductPromotionalResources) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AwsSaasProductPromotionalResources) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AwsSaasProductPromotionalResources) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetVideoUrls returns the VideoUrls field value if set, zero value otherwise.
func (o *AwsSaasProductPromotionalResources) GetVideoUrls() []string {
	if o == nil || IsNil(o.VideoUrls) {
		var ret []string
		return ret
	}
	return o.VideoUrls
}

// GetVideoUrlsOk returns a tuple with the VideoUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSaasProductPromotionalResources) GetVideoUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoUrls) {
		return nil, false
	}
	return o.VideoUrls, true
}

// HasVideoUrls returns a boolean if a field has been set.
func (o *AwsSaasProductPromotionalResources) HasVideoUrls() bool {
	if o != nil && !IsNil(o.VideoUrls) {
		return true
	}

	return false
}

// SetVideoUrls gets a reference to the given []string and assigns it to the VideoUrls field.
func (o *AwsSaasProductPromotionalResources) SetVideoUrls(v []string) {
	o.VideoUrls = v
}

func (o AwsSaasProductPromotionalResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsSaasProductPromotionalResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalResources) {
		toSerialize["AdditionalResources"] = o.AdditionalResources
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["LogoUrl"] = o.LogoUrl
	}
	if !IsNil(o.VideoUrls) {
		toSerialize["VideoUrls"] = o.VideoUrls
	}
	return toSerialize, nil
}

type NullableAwsSaasProductPromotionalResources struct {
	value *AwsSaasProductPromotionalResources
	isSet bool
}

func (v NullableAwsSaasProductPromotionalResources) Get() *AwsSaasProductPromotionalResources {
	return v.value
}

func (v *NullableAwsSaasProductPromotionalResources) Set(val *AwsSaasProductPromotionalResources) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSaasProductPromotionalResources) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSaasProductPromotionalResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSaasProductPromotionalResources(val *AwsSaasProductPromotionalResources) *NullableAwsSaasProductPromotionalResources {
	return &NullableAwsSaasProductPromotionalResources{value: val, isSet: true}
}

func (v NullableAwsSaasProductPromotionalResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSaasProductPromotionalResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


