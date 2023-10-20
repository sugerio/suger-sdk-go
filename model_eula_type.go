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
	"fmt"
)

// EulaType the model 'EulaType'
type EulaType string

// List of EulaType
const (
	EulaType_UNKNOWN EulaType = ""
	EulaType_SCMP EulaType = "SCMP"
	EulaType_ECMP EulaType = "ECMP"
	EulaType_CUSTOM EulaType = "CUSTOM"
	EulaType_ISV EulaType = "ISV"
	EulaType_CURRENT EulaType = "CURRENT"
)

// All allowed values of EulaType enum
var AllowedEulaTypeEnumValues = []EulaType{
	"",
	"SCMP",
	"ECMP",
	"CUSTOM",
	"ISV",
	"CURRENT",
}

func (v *EulaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EulaType(value)
	for _, existing := range AllowedEulaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EulaType", value)
}

// NewEulaTypeFromValue returns a pointer to a valid EulaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEulaTypeFromValue(v string) (*EulaType, error) {
	ev := EulaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EulaType: valid values are %v", v, AllowedEulaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EulaType) IsValid() bool {
	for _, existing := range AllowedEulaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EulaType value
func (v EulaType) Ptr() *EulaType {
	return &v
}

type NullableEulaType struct {
	value *EulaType
	isSet bool
}

func (v NullableEulaType) Get() *EulaType {
	return v.value
}

func (v *NullableEulaType) Set(val *EulaType) {
	v.value = val
	v.isSet = true
}

func (v NullableEulaType) IsSet() bool {
	return v.isSet
}

func (v *NullableEulaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEulaType(val *EulaType) *NullableEulaType {
	return &NullableEulaType{value: val, isSet: true}
}

func (v NullableEulaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEulaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

