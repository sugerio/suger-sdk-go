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

// checks if the AwsProductRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsProductRepository{}

// AwsProductRepository struct for AwsProductRepository
type AwsProductRepository struct {
	Type *string `json:"Type,omitempty"`
	Url  *string `json:"Url,omitempty"`
}

// NewAwsProductRepository instantiates a new AwsProductRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsProductRepository() *AwsProductRepository {
	this := AwsProductRepository{}
	return &this
}

// NewAwsProductRepositoryWithDefaults instantiates a new AwsProductRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsProductRepositoryWithDefaults() *AwsProductRepository {
	this := AwsProductRepository{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AwsProductRepository) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductRepository) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AwsProductRepository) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AwsProductRepository) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AwsProductRepository) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProductRepository) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AwsProductRepository) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AwsProductRepository) SetUrl(v string) {
	o.Url = &v
}

func (o AwsProductRepository) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsProductRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAwsProductRepository struct {
	value *AwsProductRepository
	isSet bool
}

func (v NullableAwsProductRepository) Get() *AwsProductRepository {
	return v.value
}

func (v *NullableAwsProductRepository) Set(val *AwsProductRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsProductRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsProductRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsProductRepository(val *AwsProductRepository) *NullableAwsProductRepository {
	return &NullableAwsProductRepository{value: val, isSet: true}
}

func (v NullableAwsProductRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsProductRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
