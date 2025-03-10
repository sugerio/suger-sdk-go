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

// checks if the BillingAddonInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingAddonInfo{}

// BillingAddonInfo struct for BillingAddonInfo
type BillingAddonInfo struct {
	Amount *float32 `json:"amount,omitempty"`
}

// NewBillingAddonInfo instantiates a new BillingAddonInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingAddonInfo() *BillingAddonInfo {
	this := BillingAddonInfo{}
	return &this
}

// NewBillingAddonInfoWithDefaults instantiates a new BillingAddonInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingAddonInfoWithDefaults() *BillingAddonInfo {
	this := BillingAddonInfo{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *BillingAddonInfo) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingAddonInfo) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *BillingAddonInfo) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *BillingAddonInfo) SetAmount(v float32) {
	o.Amount = &v
}

func (o BillingAddonInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingAddonInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableBillingAddonInfo struct {
	value *BillingAddonInfo
	isSet bool
}

func (v NullableBillingAddonInfo) Get() *BillingAddonInfo {
	return v.value
}

func (v *NullableBillingAddonInfo) Set(val *BillingAddonInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingAddonInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingAddonInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingAddonInfo(val *BillingAddonInfo) *NullableBillingAddonInfo {
	return &NullableBillingAddonInfo{value: val, isSet: true}
}

func (v NullableBillingAddonInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingAddonInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
