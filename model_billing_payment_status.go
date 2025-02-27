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

// BillingPaymentStatus the model 'BillingPaymentStatus'
type BillingPaymentStatus string

// List of BillingPaymentStatus
const (
	BillingPaymentStatus_PENDING    BillingPaymentStatus = "PENDING"
	BillingPaymentStatus_PROCESSING BillingPaymentStatus = "PROCESSING"
	BillingPaymentStatus_SUCCESS    BillingPaymentStatus = "SUCCESS"
	BillingPaymentStatus_FAILED     BillingPaymentStatus = "FAILED"
)

// All allowed values of BillingPaymentStatus enum
var AllowedBillingPaymentStatusEnumValues = []BillingPaymentStatus{
	"PENDING",
	"PROCESSING",
	"SUCCESS",
	"FAILED",
}

func (v *BillingPaymentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingPaymentStatus(value)
	for _, existing := range AllowedBillingPaymentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingPaymentStatus", value)
}

// NewBillingPaymentStatusFromValue returns a pointer to a valid BillingPaymentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingPaymentStatusFromValue(v string) (*BillingPaymentStatus, error) {
	ev := BillingPaymentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingPaymentStatus: valid values are %v", v, AllowedBillingPaymentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingPaymentStatus) IsValid() bool {
	for _, existing := range AllowedBillingPaymentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingPaymentStatus value
func (v BillingPaymentStatus) Ptr() *BillingPaymentStatus {
	return &v
}

type NullableBillingPaymentStatus struct {
	value *BillingPaymentStatus
	isSet bool
}

func (v NullableBillingPaymentStatus) Get() *BillingPaymentStatus {
	return v.value
}

func (v *NullableBillingPaymentStatus) Set(val *BillingPaymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPaymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPaymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPaymentStatus(val *BillingPaymentStatus) *NullableBillingPaymentStatus {
	return &NullableBillingPaymentStatus{value: val, isSet: true}
}

func (v NullableBillingPaymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPaymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
