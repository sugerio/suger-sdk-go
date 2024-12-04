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

// checks if the AzureMarketplaceResourceTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceResourceTarget{}

// AzureMarketplaceResourceTarget struct for AzureMarketplaceResourceTarget
type AzureMarketplaceResourceTarget struct {
	TargetId   *string `json:"targetId,omitempty"`
	TargetType *string `json:"targetType,omitempty"`
}

// NewAzureMarketplaceResourceTarget instantiates a new AzureMarketplaceResourceTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceResourceTarget() *AzureMarketplaceResourceTarget {
	this := AzureMarketplaceResourceTarget{}
	return &this
}

// NewAzureMarketplaceResourceTargetWithDefaults instantiates a new AzureMarketplaceResourceTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceResourceTargetWithDefaults() *AzureMarketplaceResourceTarget {
	this := AzureMarketplaceResourceTarget{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AzureMarketplaceResourceTarget) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceResourceTarget) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AzureMarketplaceResourceTarget) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AzureMarketplaceResourceTarget) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *AzureMarketplaceResourceTarget) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceResourceTarget) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *AzureMarketplaceResourceTarget) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *AzureMarketplaceResourceTarget) SetTargetType(v string) {
	o.TargetType = &v
}

func (o AzureMarketplaceResourceTarget) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceResourceTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["targetType"] = o.TargetType
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceResourceTarget struct {
	value *AzureMarketplaceResourceTarget
	isSet bool
}

func (v NullableAzureMarketplaceResourceTarget) Get() *AzureMarketplaceResourceTarget {
	return v.value
}

func (v *NullableAzureMarketplaceResourceTarget) Set(val *AzureMarketplaceResourceTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceResourceTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceResourceTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceResourceTarget(val *AzureMarketplaceResourceTarget) *NullableAzureMarketplaceResourceTarget {
	return &NullableAzureMarketplaceResourceTarget{value: val, isSet: true}
}

func (v NullableAzureMarketplaceResourceTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceResourceTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}