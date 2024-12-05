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

// checks if the InvoiceAdjustOverallMinimumSpend type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceAdjustOverallMinimumSpend{}

// InvoiceAdjustOverallMinimumSpend struct for InvoiceAdjustOverallMinimumSpend
type InvoiceAdjustOverallMinimumSpend struct {
	MinimumSpend *float32 `json:"minimumSpend,omitempty"`
	Reason       *string  `json:"reason,omitempty"`
}

// NewInvoiceAdjustOverallMinimumSpend instantiates a new InvoiceAdjustOverallMinimumSpend object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceAdjustOverallMinimumSpend() *InvoiceAdjustOverallMinimumSpend {
	this := InvoiceAdjustOverallMinimumSpend{}
	return &this
}

// NewInvoiceAdjustOverallMinimumSpendWithDefaults instantiates a new InvoiceAdjustOverallMinimumSpend object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceAdjustOverallMinimumSpendWithDefaults() *InvoiceAdjustOverallMinimumSpend {
	this := InvoiceAdjustOverallMinimumSpend{}
	return &this
}

// GetMinimumSpend returns the MinimumSpend field value if set, zero value otherwise.
func (o *InvoiceAdjustOverallMinimumSpend) GetMinimumSpend() float32 {
	if o == nil || IsNil(o.MinimumSpend) {
		var ret float32
		return ret
	}
	return *o.MinimumSpend
}

// GetMinimumSpendOk returns a tuple with the MinimumSpend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustOverallMinimumSpend) GetMinimumSpendOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimumSpend) {
		return nil, false
	}
	return o.MinimumSpend, true
}

// HasMinimumSpend returns a boolean if a field has been set.
func (o *InvoiceAdjustOverallMinimumSpend) HasMinimumSpend() bool {
	if o != nil && !IsNil(o.MinimumSpend) {
		return true
	}

	return false
}

// SetMinimumSpend gets a reference to the given float32 and assigns it to the MinimumSpend field.
func (o *InvoiceAdjustOverallMinimumSpend) SetMinimumSpend(v float32) {
	o.MinimumSpend = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvoiceAdjustOverallMinimumSpend) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustOverallMinimumSpend) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvoiceAdjustOverallMinimumSpend) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvoiceAdjustOverallMinimumSpend) SetReason(v string) {
	o.Reason = &v
}

func (o InvoiceAdjustOverallMinimumSpend) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceAdjustOverallMinimumSpend) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinimumSpend) {
		toSerialize["minimumSpend"] = o.MinimumSpend
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableInvoiceAdjustOverallMinimumSpend struct {
	value *InvoiceAdjustOverallMinimumSpend
	isSet bool
}

func (v NullableInvoiceAdjustOverallMinimumSpend) Get() *InvoiceAdjustOverallMinimumSpend {
	return v.value
}

func (v *NullableInvoiceAdjustOverallMinimumSpend) Set(val *InvoiceAdjustOverallMinimumSpend) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceAdjustOverallMinimumSpend) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceAdjustOverallMinimumSpend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceAdjustOverallMinimumSpend(val *InvoiceAdjustOverallMinimumSpend) *NullableInvoiceAdjustOverallMinimumSpend {
	return &NullableInvoiceAdjustOverallMinimumSpend{value: val, isSet: true}
}

func (v NullableInvoiceAdjustOverallMinimumSpend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceAdjustOverallMinimumSpend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
