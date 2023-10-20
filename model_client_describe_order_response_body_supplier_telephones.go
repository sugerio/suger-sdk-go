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

// checks if the ClientDescribeOrderResponseBodySupplierTelephones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientDescribeOrderResponseBodySupplierTelephones{}

// ClientDescribeOrderResponseBodySupplierTelephones struct for ClientDescribeOrderResponseBodySupplierTelephones
type ClientDescribeOrderResponseBodySupplierTelephones struct {
	Telephone []string `json:"Telephone,omitempty"`
}

// NewClientDescribeOrderResponseBodySupplierTelephones instantiates a new ClientDescribeOrderResponseBodySupplierTelephones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientDescribeOrderResponseBodySupplierTelephones() *ClientDescribeOrderResponseBodySupplierTelephones {
	this := ClientDescribeOrderResponseBodySupplierTelephones{}
	return &this
}

// NewClientDescribeOrderResponseBodySupplierTelephonesWithDefaults instantiates a new ClientDescribeOrderResponseBodySupplierTelephones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientDescribeOrderResponseBodySupplierTelephonesWithDefaults() *ClientDescribeOrderResponseBodySupplierTelephones {
	this := ClientDescribeOrderResponseBodySupplierTelephones{}
	return &this
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *ClientDescribeOrderResponseBodySupplierTelephones) GetTelephone() []string {
	if o == nil || IsNil(o.Telephone) {
		var ret []string
		return ret
	}
	return o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeOrderResponseBodySupplierTelephones) GetTelephoneOk() ([]string, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *ClientDescribeOrderResponseBodySupplierTelephones) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given []string and assigns it to the Telephone field.
func (o *ClientDescribeOrderResponseBodySupplierTelephones) SetTelephone(v []string) {
	o.Telephone = v
}

func (o ClientDescribeOrderResponseBodySupplierTelephones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientDescribeOrderResponseBodySupplierTelephones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Telephone) {
		toSerialize["Telephone"] = o.Telephone
	}
	return toSerialize, nil
}

type NullableClientDescribeOrderResponseBodySupplierTelephones struct {
	value *ClientDescribeOrderResponseBodySupplierTelephones
	isSet bool
}

func (v NullableClientDescribeOrderResponseBodySupplierTelephones) Get() *ClientDescribeOrderResponseBodySupplierTelephones {
	return v.value
}

func (v *NullableClientDescribeOrderResponseBodySupplierTelephones) Set(val *ClientDescribeOrderResponseBodySupplierTelephones) {
	v.value = val
	v.isSet = true
}

func (v NullableClientDescribeOrderResponseBodySupplierTelephones) IsSet() bool {
	return v.isSet
}

func (v *NullableClientDescribeOrderResponseBodySupplierTelephones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientDescribeOrderResponseBodySupplierTelephones(val *ClientDescribeOrderResponseBodySupplierTelephones) *NullableClientDescribeOrderResponseBodySupplierTelephones {
	return &NullableClientDescribeOrderResponseBodySupplierTelephones{value: val, isSet: true}
}

func (v NullableClientDescribeOrderResponseBodySupplierTelephones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientDescribeOrderResponseBodySupplierTelephones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


