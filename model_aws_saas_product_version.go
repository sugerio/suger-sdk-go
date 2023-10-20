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

// checks if the AwsSaasProductVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsSaasProductVersion{}

// AwsSaasProductVersion struct for AwsSaasProductVersion
type AwsSaasProductVersion struct {
	DeliveryOptions []AwsSaasProductDeliveryOption `json:"DeliveryOptions,omitempty"`
	Id *string `json:"Id,omitempty"`
}

// NewAwsSaasProductVersion instantiates a new AwsSaasProductVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsSaasProductVersion() *AwsSaasProductVersion {
	this := AwsSaasProductVersion{}
	return &this
}

// NewAwsSaasProductVersionWithDefaults instantiates a new AwsSaasProductVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsSaasProductVersionWithDefaults() *AwsSaasProductVersion {
	this := AwsSaasProductVersion{}
	return &this
}

// GetDeliveryOptions returns the DeliveryOptions field value if set, zero value otherwise.
func (o *AwsSaasProductVersion) GetDeliveryOptions() []AwsSaasProductDeliveryOption {
	if o == nil || IsNil(o.DeliveryOptions) {
		var ret []AwsSaasProductDeliveryOption
		return ret
	}
	return o.DeliveryOptions
}

// GetDeliveryOptionsOk returns a tuple with the DeliveryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSaasProductVersion) GetDeliveryOptionsOk() ([]AwsSaasProductDeliveryOption, bool) {
	if o == nil || IsNil(o.DeliveryOptions) {
		return nil, false
	}
	return o.DeliveryOptions, true
}

// HasDeliveryOptions returns a boolean if a field has been set.
func (o *AwsSaasProductVersion) HasDeliveryOptions() bool {
	if o != nil && !IsNil(o.DeliveryOptions) {
		return true
	}

	return false
}

// SetDeliveryOptions gets a reference to the given []AwsSaasProductDeliveryOption and assigns it to the DeliveryOptions field.
func (o *AwsSaasProductVersion) SetDeliveryOptions(v []AwsSaasProductDeliveryOption) {
	o.DeliveryOptions = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AwsSaasProductVersion) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsSaasProductVersion) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AwsSaasProductVersion) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AwsSaasProductVersion) SetId(v string) {
	o.Id = &v
}

func (o AwsSaasProductVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsSaasProductVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryOptions) {
		toSerialize["DeliveryOptions"] = o.DeliveryOptions
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAwsSaasProductVersion struct {
	value *AwsSaasProductVersion
	isSet bool
}

func (v NullableAwsSaasProductVersion) Get() *AwsSaasProductVersion {
	return v.value
}

func (v *NullableAwsSaasProductVersion) Set(val *AwsSaasProductVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsSaasProductVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsSaasProductVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsSaasProductVersion(val *AwsSaasProductVersion) *NullableAwsSaasProductVersion {
	return &NullableAwsSaasProductVersion{value: val, isSet: true}
}

func (v NullableAwsSaasProductVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsSaasProductVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


