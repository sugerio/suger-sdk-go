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

// checks if the StripeRefundDestinationDetailsUSBankTransfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StripeRefundDestinationDetailsUSBankTransfer{}

// StripeRefundDestinationDetailsUSBankTransfer struct for StripeRefundDestinationDetailsUSBankTransfer
type StripeRefundDestinationDetailsUSBankTransfer struct {
	// The reference assigned to the refund.
	Reference *string `json:"reference,omitempty"`
	// Status of the reference on the refund. This can be `pending`, `available` or `unavailable`.
	ReferenceStatus *string `json:"reference_status,omitempty"`
}

// NewStripeRefundDestinationDetailsUSBankTransfer instantiates a new StripeRefundDestinationDetailsUSBankTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStripeRefundDestinationDetailsUSBankTransfer() *StripeRefundDestinationDetailsUSBankTransfer {
	this := StripeRefundDestinationDetailsUSBankTransfer{}
	return &this
}

// NewStripeRefundDestinationDetailsUSBankTransferWithDefaults instantiates a new StripeRefundDestinationDetailsUSBankTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStripeRefundDestinationDetailsUSBankTransferWithDefaults() *StripeRefundDestinationDetailsUSBankTransfer {
	this := StripeRefundDestinationDetailsUSBankTransfer{}
	return &this
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *StripeRefundDestinationDetailsUSBankTransfer) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *StripeRefundDestinationDetailsUSBankTransfer) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceStatus returns the ReferenceStatus field value if set, zero value otherwise.
func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReferenceStatus() string {
	if o == nil || IsNil(o.ReferenceStatus) {
		var ret string
		return ret
	}
	return *o.ReferenceStatus
}

// GetReferenceStatusOk returns a tuple with the ReferenceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReferenceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceStatus) {
		return nil, false
	}
	return o.ReferenceStatus, true
}

// HasReferenceStatus returns a boolean if a field has been set.
func (o *StripeRefundDestinationDetailsUSBankTransfer) HasReferenceStatus() bool {
	if o != nil && !IsNil(o.ReferenceStatus) {
		return true
	}

	return false
}

// SetReferenceStatus gets a reference to the given string and assigns it to the ReferenceStatus field.
func (o *StripeRefundDestinationDetailsUSBankTransfer) SetReferenceStatus(v string) {
	o.ReferenceStatus = &v
}

func (o StripeRefundDestinationDetailsUSBankTransfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StripeRefundDestinationDetailsUSBankTransfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.ReferenceStatus) {
		toSerialize["reference_status"] = o.ReferenceStatus
	}
	return toSerialize, nil
}

type NullableStripeRefundDestinationDetailsUSBankTransfer struct {
	value *StripeRefundDestinationDetailsUSBankTransfer
	isSet bool
}

func (v NullableStripeRefundDestinationDetailsUSBankTransfer) Get() *StripeRefundDestinationDetailsUSBankTransfer {
	return v.value
}

func (v *NullableStripeRefundDestinationDetailsUSBankTransfer) Set(val *StripeRefundDestinationDetailsUSBankTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeRefundDestinationDetailsUSBankTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeRefundDestinationDetailsUSBankTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeRefundDestinationDetailsUSBankTransfer(val *StripeRefundDestinationDetailsUSBankTransfer) *NullableStripeRefundDestinationDetailsUSBankTransfer {
	return &NullableStripeRefundDestinationDetailsUSBankTransfer{value: val, isSet: true}
}

func (v NullableStripeRefundDestinationDetailsUSBankTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeRefundDestinationDetailsUSBankTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
