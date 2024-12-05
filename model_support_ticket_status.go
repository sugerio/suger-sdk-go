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

// SupportTicketStatus the model 'SupportTicketStatus'
type SupportTicketStatus string

// List of SupportTicketStatus
const (
	SupportTicketStatus_OPEN        SupportTicketStatus = "OPEN"
	SupportTicketStatus_IN_PROGRESS SupportTicketStatus = "IN PROGRESS"
	SupportTicketStatus_IN_REVIEW   SupportTicketStatus = "IN REVIEW"
	SupportTicketStatus_BLOCKED     SupportTicketStatus = "BLOCKED"
	SupportTicketStatus_CLOSED      SupportTicketStatus = "CLOSED"
	SupportTicketStatus_ARCHIVED    SupportTicketStatus = "ARCHIVED"
	SupportTicketStatus_DELETED     SupportTicketStatus = "DELETED"
)

// All allowed values of SupportTicketStatus enum
var AllowedSupportTicketStatusEnumValues = []SupportTicketStatus{
	"OPEN",
	"IN PROGRESS",
	"IN REVIEW",
	"BLOCKED",
	"CLOSED",
	"ARCHIVED",
	"DELETED",
}

func (v *SupportTicketStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SupportTicketStatus(value)
	for _, existing := range AllowedSupportTicketStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SupportTicketStatus", value)
}

// NewSupportTicketStatusFromValue returns a pointer to a valid SupportTicketStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSupportTicketStatusFromValue(v string) (*SupportTicketStatus, error) {
	ev := SupportTicketStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SupportTicketStatus: valid values are %v", v, AllowedSupportTicketStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SupportTicketStatus) IsValid() bool {
	for _, existing := range AllowedSupportTicketStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SupportTicketStatus value
func (v SupportTicketStatus) Ptr() *SupportTicketStatus {
	return &v
}

type NullableSupportTicketStatus struct {
	value *SupportTicketStatus
	isSet bool
}

func (v NullableSupportTicketStatus) Get() *SupportTicketStatus {
	return v.value
}

func (v *NullableSupportTicketStatus) Set(val *SupportTicketStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicketStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicketStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicketStatus(val *SupportTicketStatus) *NullableSupportTicketStatus {
	return &NullableSupportTicketStatus{value: val, isSet: true}
}

func (v NullableSupportTicketStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicketStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
