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
	"fmt"
)

// BillingInvoiceStatus the model 'BillingInvoiceStatus'
type BillingInvoiceStatus string

// List of BillingInvoiceStatus
const (
	BillingInvoiceStatus_DRAFT     BillingInvoiceStatus = "DRAFT"
	BillingInvoiceStatus_FINALIZED BillingInvoiceStatus = "FINALIZED"
	BillingInvoiceStatus_CANCELED  BillingInvoiceStatus = "CANCELED"
	BillingInvoiceStatus_DELETED   BillingInvoiceStatus = "DELETED"
	BillingInvoiceStatus_UNKNOWN   BillingInvoiceStatus = ""
)

// All allowed values of BillingInvoiceStatus enum
var AllowedBillingInvoiceStatusEnumValues = []BillingInvoiceStatus{
	"DRAFT",
	"FINALIZED",
	"CANCELED",
	"DELETED",
	"",
}

func (v *BillingInvoiceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingInvoiceStatus(value)
	for _, existing := range AllowedBillingInvoiceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingInvoiceStatus", value)
}

// NewBillingInvoiceStatusFromValue returns a pointer to a valid BillingInvoiceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingInvoiceStatusFromValue(v string) (*BillingInvoiceStatus, error) {
	ev := BillingInvoiceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingInvoiceStatus: valid values are %v", v, AllowedBillingInvoiceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingInvoiceStatus) IsValid() bool {
	for _, existing := range AllowedBillingInvoiceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingInvoiceStatus value
func (v BillingInvoiceStatus) Ptr() *BillingInvoiceStatus {
	return &v
}

type NullableBillingInvoiceStatus struct {
	value *BillingInvoiceStatus
	isSet bool
}

func (v NullableBillingInvoiceStatus) Get() *BillingInvoiceStatus {
	return v.value
}

func (v *NullableBillingInvoiceStatus) Set(val *BillingInvoiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInvoiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInvoiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInvoiceStatus(val *BillingInvoiceStatus) *NullableBillingInvoiceStatus {
	return &NullableBillingInvoiceStatus{value: val, isSet: true}
}

func (v NullableBillingInvoiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInvoiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
