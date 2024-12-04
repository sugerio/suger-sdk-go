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

// checks if the AwsMarketplaceBuyerAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceBuyerAccount{}

// AwsMarketplaceBuyerAccount struct for AwsMarketplaceBuyerAccount
type AwsMarketplaceBuyerAccount struct {
	AwsAccountId *string `json:"AwsAccountId,omitempty"`
}

// NewAwsMarketplaceBuyerAccount instantiates a new AwsMarketplaceBuyerAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceBuyerAccount() *AwsMarketplaceBuyerAccount {
	this := AwsMarketplaceBuyerAccount{}
	return &this
}

// NewAwsMarketplaceBuyerAccountWithDefaults instantiates a new AwsMarketplaceBuyerAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceBuyerAccountWithDefaults() *AwsMarketplaceBuyerAccount {
	this := AwsMarketplaceBuyerAccount{}
	return &this
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise.
func (o *AwsMarketplaceBuyerAccount) GetAwsAccountId() string {
	if o == nil || IsNil(o.AwsAccountId) {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceBuyerAccount) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountId) {
		return nil, false
	}
	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *AwsMarketplaceBuyerAccount) HasAwsAccountId() bool {
	if o != nil && !IsNil(o.AwsAccountId) {
		return true
	}

	return false
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *AwsMarketplaceBuyerAccount) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

func (o AwsMarketplaceBuyerAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceBuyerAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AwsAccountId) {
		toSerialize["AwsAccountId"] = o.AwsAccountId
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceBuyerAccount struct {
	value *AwsMarketplaceBuyerAccount
	isSet bool
}

func (v NullableAwsMarketplaceBuyerAccount) Get() *AwsMarketplaceBuyerAccount {
	return v.value
}

func (v *NullableAwsMarketplaceBuyerAccount) Set(val *AwsMarketplaceBuyerAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceBuyerAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceBuyerAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceBuyerAccount(val *AwsMarketplaceBuyerAccount) *NullableAwsMarketplaceBuyerAccount {
	return &NullableAwsMarketplaceBuyerAccount{value: val, isSet: true}
}

func (v NullableAwsMarketplaceBuyerAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceBuyerAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}