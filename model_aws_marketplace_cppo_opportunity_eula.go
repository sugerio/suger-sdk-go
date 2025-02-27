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

// checks if the AwsMarketplaceCppoOpportunityEula type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceCppoOpportunityEula{}

// AwsMarketplaceCppoOpportunityEula struct for AwsMarketplaceCppoOpportunityEula
type AwsMarketplaceCppoOpportunityEula struct {
	// The S3 signed URL of the EULA file.
	AccessUrl *string `json:"accessUrl,omitempty"`
	Key       *string `json:"key,omitempty"`
	ObjectUrl *string `json:"objectUrl,omitempty"`
}

// NewAwsMarketplaceCppoOpportunityEula instantiates a new AwsMarketplaceCppoOpportunityEula object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceCppoOpportunityEula() *AwsMarketplaceCppoOpportunityEula {
	this := AwsMarketplaceCppoOpportunityEula{}
	return &this
}

// NewAwsMarketplaceCppoOpportunityEulaWithDefaults instantiates a new AwsMarketplaceCppoOpportunityEula object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceCppoOpportunityEulaWithDefaults() *AwsMarketplaceCppoOpportunityEula {
	this := AwsMarketplaceCppoOpportunityEula{}
	return &this
}

// GetAccessUrl returns the AccessUrl field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityEula) GetAccessUrl() string {
	if o == nil || IsNil(o.AccessUrl) {
		var ret string
		return ret
	}
	return *o.AccessUrl
}

// GetAccessUrlOk returns a tuple with the AccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityEula) GetAccessUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AccessUrl) {
		return nil, false
	}
	return o.AccessUrl, true
}

// HasAccessUrl returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityEula) HasAccessUrl() bool {
	if o != nil && !IsNil(o.AccessUrl) {
		return true
	}

	return false
}

// SetAccessUrl gets a reference to the given string and assigns it to the AccessUrl field.
func (o *AwsMarketplaceCppoOpportunityEula) SetAccessUrl(v string) {
	o.AccessUrl = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityEula) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityEula) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityEula) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AwsMarketplaceCppoOpportunityEula) SetKey(v string) {
	o.Key = &v
}

// GetObjectUrl returns the ObjectUrl field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityEula) GetObjectUrl() string {
	if o == nil || IsNil(o.ObjectUrl) {
		var ret string
		return ret
	}
	return *o.ObjectUrl
}

// GetObjectUrlOk returns a tuple with the ObjectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityEula) GetObjectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectUrl) {
		return nil, false
	}
	return o.ObjectUrl, true
}

// HasObjectUrl returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityEula) HasObjectUrl() bool {
	if o != nil && !IsNil(o.ObjectUrl) {
		return true
	}

	return false
}

// SetObjectUrl gets a reference to the given string and assigns it to the ObjectUrl field.
func (o *AwsMarketplaceCppoOpportunityEula) SetObjectUrl(v string) {
	o.ObjectUrl = &v
}

func (o AwsMarketplaceCppoOpportunityEula) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceCppoOpportunityEula) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessUrl) {
		toSerialize["accessUrl"] = o.AccessUrl
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.ObjectUrl) {
		toSerialize["objectUrl"] = o.ObjectUrl
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceCppoOpportunityEula struct {
	value *AwsMarketplaceCppoOpportunityEula
	isSet bool
}

func (v NullableAwsMarketplaceCppoOpportunityEula) Get() *AwsMarketplaceCppoOpportunityEula {
	return v.value
}

func (v *NullableAwsMarketplaceCppoOpportunityEula) Set(val *AwsMarketplaceCppoOpportunityEula) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceCppoOpportunityEula) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceCppoOpportunityEula) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceCppoOpportunityEula(val *AwsMarketplaceCppoOpportunityEula) *NullableAwsMarketplaceCppoOpportunityEula {
	return &NullableAwsMarketplaceCppoOpportunityEula{value: val, isSet: true}
}

func (v NullableAwsMarketplaceCppoOpportunityEula) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceCppoOpportunityEula) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
