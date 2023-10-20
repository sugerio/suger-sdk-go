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

// MicrosoftPartnerReferralStatus the model 'MicrosoftPartnerReferralStatus'
type MicrosoftPartnerReferralStatus string

// List of MicrosoftPartnerReferralStatus
const (
	MicrosoftPartnerReferralStatus_Unknown MicrosoftPartnerReferralStatus = ""
	MicrosoftPartnerReferralStatus_New MicrosoftPartnerReferralStatus = "New"
	MicrosoftPartnerReferralStatus_Active MicrosoftPartnerReferralStatus = "Active"
	MicrosoftPartnerReferralStatus_Closed MicrosoftPartnerReferralStatus = "Closed"
	MicrosoftPartnerReferralStatus_None MicrosoftPartnerReferralStatus = "None"
)

// All allowed values of MicrosoftPartnerReferralStatus enum
var AllowedMicrosoftPartnerReferralStatusEnumValues = []MicrosoftPartnerReferralStatus{
	"",
	"New",
	"Active",
	"Closed",
	"None",
}

func (v *MicrosoftPartnerReferralStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftPartnerReferralStatus(value)
	for _, existing := range AllowedMicrosoftPartnerReferralStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftPartnerReferralStatus", value)
}

// NewMicrosoftPartnerReferralStatusFromValue returns a pointer to a valid MicrosoftPartnerReferralStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftPartnerReferralStatusFromValue(v string) (*MicrosoftPartnerReferralStatus, error) {
	ev := MicrosoftPartnerReferralStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftPartnerReferralStatus: valid values are %v", v, AllowedMicrosoftPartnerReferralStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftPartnerReferralStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftPartnerReferralStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftPartnerReferralStatus value
func (v MicrosoftPartnerReferralStatus) Ptr() *MicrosoftPartnerReferralStatus {
	return &v
}

type NullableMicrosoftPartnerReferralStatus struct {
	value *MicrosoftPartnerReferralStatus
	isSet bool
}

func (v NullableMicrosoftPartnerReferralStatus) Get() *MicrosoftPartnerReferralStatus {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralStatus) Set(val *MicrosoftPartnerReferralStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralStatus(val *MicrosoftPartnerReferralStatus) *NullableMicrosoftPartnerReferralStatus {
	return &NullableMicrosoftPartnerReferralStatus{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

