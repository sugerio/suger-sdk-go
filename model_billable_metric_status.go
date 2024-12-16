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

// BillableMetricStatus the model 'BillableMetricStatus'
type BillableMetricStatus string

// List of BillableMetricStatus
const (
	BillableMetricStatus_ACTIVE  BillableMetricStatus = "ACTIVE"
	BillableMetricStatus_DELETED BillableMetricStatus = "DELETED"
)

// All allowed values of BillableMetricStatus enum
var AllowedBillableMetricStatusEnumValues = []BillableMetricStatus{
	"ACTIVE",
	"DELETED",
}

func (v *BillableMetricStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillableMetricStatus(value)
	for _, existing := range AllowedBillableMetricStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillableMetricStatus", value)
}

// NewBillableMetricStatusFromValue returns a pointer to a valid BillableMetricStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillableMetricStatusFromValue(v string) (*BillableMetricStatus, error) {
	ev := BillableMetricStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillableMetricStatus: valid values are %v", v, AllowedBillableMetricStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillableMetricStatus) IsValid() bool {
	for _, existing := range AllowedBillableMetricStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillableMetricStatus value
func (v BillableMetricStatus) Ptr() *BillableMetricStatus {
	return &v
}

type NullableBillableMetricStatus struct {
	value *BillableMetricStatus
	isSet bool
}

func (v NullableBillableMetricStatus) Get() *BillableMetricStatus {
	return v.value
}

func (v *NullableBillableMetricStatus) Set(val *BillableMetricStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricStatus(val *BillableMetricStatus) *NullableBillableMetricStatus {
	return &NullableBillableMetricStatus{value: val, isSet: true}
}

func (v NullableBillableMetricStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
