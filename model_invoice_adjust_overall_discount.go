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

// checks if the InvoiceAdjustOverallDiscount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceAdjustOverallDiscount{}

// InvoiceAdjustOverallDiscount struct for InvoiceAdjustOverallDiscount
type InvoiceAdjustOverallDiscount struct {
	Discount *BillingDiscount `json:"discount,omitempty"`
	Reason   *string          `json:"reason,omitempty"`
}

// NewInvoiceAdjustOverallDiscount instantiates a new InvoiceAdjustOverallDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceAdjustOverallDiscount() *InvoiceAdjustOverallDiscount {
	this := InvoiceAdjustOverallDiscount{}
	return &this
}

// NewInvoiceAdjustOverallDiscountWithDefaults instantiates a new InvoiceAdjustOverallDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceAdjustOverallDiscountWithDefaults() *InvoiceAdjustOverallDiscount {
	this := InvoiceAdjustOverallDiscount{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *InvoiceAdjustOverallDiscount) GetDiscount() BillingDiscount {
	if o == nil || IsNil(o.Discount) {
		var ret BillingDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustOverallDiscount) GetDiscountOk() (*BillingDiscount, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *InvoiceAdjustOverallDiscount) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given BillingDiscount and assigns it to the Discount field.
func (o *InvoiceAdjustOverallDiscount) SetDiscount(v BillingDiscount) {
	o.Discount = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InvoiceAdjustOverallDiscount) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceAdjustOverallDiscount) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InvoiceAdjustOverallDiscount) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InvoiceAdjustOverallDiscount) SetReason(v string) {
	o.Reason = &v
}

func (o InvoiceAdjustOverallDiscount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceAdjustOverallDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableInvoiceAdjustOverallDiscount struct {
	value *InvoiceAdjustOverallDiscount
	isSet bool
}

func (v NullableInvoiceAdjustOverallDiscount) Get() *InvoiceAdjustOverallDiscount {
	return v.value
}

func (v *NullableInvoiceAdjustOverallDiscount) Set(val *InvoiceAdjustOverallDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceAdjustOverallDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceAdjustOverallDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceAdjustOverallDiscount(val *InvoiceAdjustOverallDiscount) *NullableInvoiceAdjustOverallDiscount {
	return &NullableInvoiceAdjustOverallDiscount{value: val, isSet: true}
}

func (v NullableInvoiceAdjustOverallDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceAdjustOverallDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}