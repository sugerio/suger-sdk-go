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

// OrbPlanStatus the model 'OrbPlanStatus'
type OrbPlanStatus string

// List of OrbPlanStatus
const (
	OrbPlanStatus_ACTIVE OrbPlanStatus = "active"
	OrbPlanStatus_ARCHIVED OrbPlanStatus = "archived"
	OrbPlanStatus_DRAFT OrbPlanStatus = "draft"
)

// All allowed values of OrbPlanStatus enum
var AllowedOrbPlanStatusEnumValues = []OrbPlanStatus{
	"active",
	"archived",
	"draft",
}

func (v *OrbPlanStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrbPlanStatus(value)
	for _, existing := range AllowedOrbPlanStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrbPlanStatus", value)
}

// NewOrbPlanStatusFromValue returns a pointer to a valid OrbPlanStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrbPlanStatusFromValue(v string) (*OrbPlanStatus, error) {
	ev := OrbPlanStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrbPlanStatus: valid values are %v", v, AllowedOrbPlanStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrbPlanStatus) IsValid() bool {
	for _, existing := range AllowedOrbPlanStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrbPlanStatus value
func (v OrbPlanStatus) Ptr() *OrbPlanStatus {
	return &v
}

type NullableOrbPlanStatus struct {
	value *OrbPlanStatus
	isSet bool
}

func (v NullableOrbPlanStatus) Get() *OrbPlanStatus {
	return v.value
}

func (v *NullableOrbPlanStatus) Set(val *OrbPlanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPlanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPlanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPlanStatus(val *OrbPlanStatus) *NullableOrbPlanStatus {
	return &NullableOrbPlanStatus{value: val, isSet: true}
}

func (v NullableOrbPlanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPlanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

