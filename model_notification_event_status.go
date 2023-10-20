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

// NotificationEventStatus the model 'NotificationEventStatus'
type NotificationEventStatus string

// List of NotificationEventStatus
const (
	NotificationEventStatus_UNKNOWN NotificationEventStatus = ""
	NotificationEventStatus_SCHEDULED NotificationEventStatus = "SCHEDULED"
	NotificationEventStatus_PENDING NotificationEventStatus = "PENDING"
	NotificationEventStatus_DONE NotificationEventStatus = "DONE"
	NotificationEventStatus_FAILED NotificationEventStatus = "FAILED"
)

// All allowed values of NotificationEventStatus enum
var AllowedNotificationEventStatusEnumValues = []NotificationEventStatus{
	"",
	"SCHEDULED",
	"PENDING",
	"DONE",
	"FAILED",
}

func (v *NotificationEventStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationEventStatus(value)
	for _, existing := range AllowedNotificationEventStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationEventStatus", value)
}

// NewNotificationEventStatusFromValue returns a pointer to a valid NotificationEventStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationEventStatusFromValue(v string) (*NotificationEventStatus, error) {
	ev := NotificationEventStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationEventStatus: valid values are %v", v, AllowedNotificationEventStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationEventStatus) IsValid() bool {
	for _, existing := range AllowedNotificationEventStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationEventStatus value
func (v NotificationEventStatus) Ptr() *NotificationEventStatus {
	return &v
}

type NullableNotificationEventStatus struct {
	value *NotificationEventStatus
	isSet bool
}

func (v NullableNotificationEventStatus) Get() *NotificationEventStatus {
	return v.value
}

func (v *NullableNotificationEventStatus) Set(val *NotificationEventStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEventStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEventStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEventStatus(val *NotificationEventStatus) *NullableNotificationEventStatus {
	return &NullableNotificationEventStatus{value: val, isSet: true}
}

func (v NullableNotificationEventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEventStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

