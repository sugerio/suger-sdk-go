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

// BillableMetricFilterValueType the model 'BillableMetricFilterValueType'
type BillableMetricFilterValueType string

// List of BillableMetricFilterValueType
const (
	BillableMetricFilterValueType_STRING BillableMetricFilterValueType = "STRING"
	BillableMetricFilterValueType_FLOAT  BillableMetricFilterValueType = "FLOAT"
)

// All allowed values of BillableMetricFilterValueType enum
var AllowedBillableMetricFilterValueTypeEnumValues = []BillableMetricFilterValueType{
	"STRING",
	"FLOAT",
}

func (v *BillableMetricFilterValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillableMetricFilterValueType(value)
	for _, existing := range AllowedBillableMetricFilterValueTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillableMetricFilterValueType", value)
}

// NewBillableMetricFilterValueTypeFromValue returns a pointer to a valid BillableMetricFilterValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillableMetricFilterValueTypeFromValue(v string) (*BillableMetricFilterValueType, error) {
	ev := BillableMetricFilterValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillableMetricFilterValueType: valid values are %v", v, AllowedBillableMetricFilterValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillableMetricFilterValueType) IsValid() bool {
	for _, existing := range AllowedBillableMetricFilterValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillableMetricFilterValueType value
func (v BillableMetricFilterValueType) Ptr() *BillableMetricFilterValueType {
	return &v
}

type NullableBillableMetricFilterValueType struct {
	value *BillableMetricFilterValueType
	isSet bool
}

func (v NullableBillableMetricFilterValueType) Get() *BillableMetricFilterValueType {
	return v.value
}

func (v *NullableBillableMetricFilterValueType) Set(val *BillableMetricFilterValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableMetricFilterValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableMetricFilterValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableMetricFilterValueType(val *BillableMetricFilterValueType) *NullableBillableMetricFilterValueType {
	return &NullableBillableMetricFilterValueType{value: val, isSet: true}
}

func (v NullableBillableMetricFilterValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableMetricFilterValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
