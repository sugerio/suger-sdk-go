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

// UsageMeteringDimensionMappingMode the model 'UsageMeteringDimensionMappingMode'
type UsageMeteringDimensionMappingMode string

// List of UsageMeteringDimensionMappingMode
const (
	UsageMeteringDimensionMappingMode_UNKNOWN  UsageMeteringDimensionMappingMode = ""
	UsageMeteringDimensionMappingMode_QUANTITY UsageMeteringDimensionMappingMode = "QUANTITY"
	UsageMeteringDimensionMappingMode_AMOUNT   UsageMeteringDimensionMappingMode = "AMOUNT"
)

// All allowed values of UsageMeteringDimensionMappingMode enum
var AllowedUsageMeteringDimensionMappingModeEnumValues = []UsageMeteringDimensionMappingMode{
	"",
	"QUANTITY",
	"AMOUNT",
}

func (v *UsageMeteringDimensionMappingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageMeteringDimensionMappingMode(value)
	for _, existing := range AllowedUsageMeteringDimensionMappingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageMeteringDimensionMappingMode", value)
}

// NewUsageMeteringDimensionMappingModeFromValue returns a pointer to a valid UsageMeteringDimensionMappingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageMeteringDimensionMappingModeFromValue(v string) (*UsageMeteringDimensionMappingMode, error) {
	ev := UsageMeteringDimensionMappingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageMeteringDimensionMappingMode: valid values are %v", v, AllowedUsageMeteringDimensionMappingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageMeteringDimensionMappingMode) IsValid() bool {
	for _, existing := range AllowedUsageMeteringDimensionMappingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageMeteringDimensionMappingMode value
func (v UsageMeteringDimensionMappingMode) Ptr() *UsageMeteringDimensionMappingMode {
	return &v
}

type NullableUsageMeteringDimensionMappingMode struct {
	value *UsageMeteringDimensionMappingMode
	isSet bool
}

func (v NullableUsageMeteringDimensionMappingMode) Get() *UsageMeteringDimensionMappingMode {
	return v.value
}

func (v *NullableUsageMeteringDimensionMappingMode) Set(val *UsageMeteringDimensionMappingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMeteringDimensionMappingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMeteringDimensionMappingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMeteringDimensionMappingMode(val *UsageMeteringDimensionMappingMode) *NullableUsageMeteringDimensionMappingMode {
	return &NullableUsageMeteringDimensionMappingMode{value: val, isSet: true}
}

func (v NullableUsageMeteringDimensionMappingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMeteringDimensionMappingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
