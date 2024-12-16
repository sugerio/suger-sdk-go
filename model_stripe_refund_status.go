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

// StripeRefundStatus the model 'StripeRefundStatus'
type StripeRefundStatus string

// List of StripeRefundStatus
const (
	StripeRefundStatus_Pending        StripeRefundStatus = "pending"
	StripeRefundStatus_RequiresAction StripeRefundStatus = "requires_action"
	StripeRefundStatus_Succeeded      StripeRefundStatus = "succeeded"
	StripeRefundStatus_Failed         StripeRefundStatus = "failed"
	StripeRefundStatus_Canceled       StripeRefundStatus = "canceled"
)

// All allowed values of StripeRefundStatus enum
var AllowedStripeRefundStatusEnumValues = []StripeRefundStatus{
	"pending",
	"requires_action",
	"succeeded",
	"failed",
	"canceled",
}

func (v *StripeRefundStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StripeRefundStatus(value)
	for _, existing := range AllowedStripeRefundStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StripeRefundStatus", value)
}

// NewStripeRefundStatusFromValue returns a pointer to a valid StripeRefundStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStripeRefundStatusFromValue(v string) (*StripeRefundStatus, error) {
	ev := StripeRefundStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StripeRefundStatus: valid values are %v", v, AllowedStripeRefundStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StripeRefundStatus) IsValid() bool {
	for _, existing := range AllowedStripeRefundStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StripeRefundStatus value
func (v StripeRefundStatus) Ptr() *StripeRefundStatus {
	return &v
}

type NullableStripeRefundStatus struct {
	value *StripeRefundStatus
	isSet bool
}

func (v NullableStripeRefundStatus) Get() *StripeRefundStatus {
	return v.value
}

func (v *NullableStripeRefundStatus) Set(val *StripeRefundStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStripeRefundStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStripeRefundStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStripeRefundStatus(val *StripeRefundStatus) *NullableStripeRefundStatus {
	return &NullableStripeRefundStatus{value: val, isSet: true}
}

func (v NullableStripeRefundStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStripeRefundStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
