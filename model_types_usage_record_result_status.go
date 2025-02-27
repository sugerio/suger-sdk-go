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

// TypesUsageRecordResultStatus the model 'TypesUsageRecordResultStatus'
type TypesUsageRecordResultStatus string

// List of types.UsageRecordResultStatus
const (
	UsageRecordResultStatusSuccess               TypesUsageRecordResultStatus = "Success"
	UsageRecordResultStatusCustomerNotSubscribed TypesUsageRecordResultStatus = "CustomerNotSubscribed"
	UsageRecordResultStatusDuplicateRecord       TypesUsageRecordResultStatus = "DuplicateRecord"
)

// All allowed values of TypesUsageRecordResultStatus enum
var AllowedTypesUsageRecordResultStatusEnumValues = []TypesUsageRecordResultStatus{
	"Success",
	"CustomerNotSubscribed",
	"DuplicateRecord",
}

func (v *TypesUsageRecordResultStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypesUsageRecordResultStatus(value)
	for _, existing := range AllowedTypesUsageRecordResultStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypesUsageRecordResultStatus", value)
}

// NewTypesUsageRecordResultStatusFromValue returns a pointer to a valid TypesUsageRecordResultStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypesUsageRecordResultStatusFromValue(v string) (*TypesUsageRecordResultStatus, error) {
	ev := TypesUsageRecordResultStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TypesUsageRecordResultStatus: valid values are %v", v, AllowedTypesUsageRecordResultStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypesUsageRecordResultStatus) IsValid() bool {
	for _, existing := range AllowedTypesUsageRecordResultStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to types.UsageRecordResultStatus value
func (v TypesUsageRecordResultStatus) Ptr() *TypesUsageRecordResultStatus {
	return &v
}

type NullableTypesUsageRecordResultStatus struct {
	value *TypesUsageRecordResultStatus
	isSet bool
}

func (v NullableTypesUsageRecordResultStatus) Get() *TypesUsageRecordResultStatus {
	return v.value
}

func (v *NullableTypesUsageRecordResultStatus) Set(val *TypesUsageRecordResultStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTypesUsageRecordResultStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTypesUsageRecordResultStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypesUsageRecordResultStatus(val *TypesUsageRecordResultStatus) *NullableTypesUsageRecordResultStatus {
	return &NullableTypesUsageRecordResultStatus{value: val, isSet: true}
}

func (v NullableTypesUsageRecordResultStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypesUsageRecordResultStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
