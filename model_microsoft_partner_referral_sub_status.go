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

// MicrosoftPartnerReferralSubStatus the model 'MicrosoftPartnerReferralSubStatus'
type MicrosoftPartnerReferralSubStatus string

// List of MicrosoftPartnerReferralSubStatus
const (
	MicrosoftPartnerReferralSubStatus_Unknown  MicrosoftPartnerReferralSubStatus = ""
	MicrosoftPartnerReferralSubStatus_None     MicrosoftPartnerReferralSubStatus = "None"
	MicrosoftPartnerReferralSubStatus_Pending  MicrosoftPartnerReferralSubStatus = "Pending"
	MicrosoftPartnerReferralSubStatus_Received MicrosoftPartnerReferralSubStatus = "Received"
	MicrosoftPartnerReferralSubStatus_Won      MicrosoftPartnerReferralSubStatus = "Won"
	MicrosoftPartnerReferralSubStatus_Lost     MicrosoftPartnerReferralSubStatus = "Lost"
	MicrosoftPartnerReferralSubStatus_Declined MicrosoftPartnerReferralSubStatus = "Declined"
	MicrosoftPartnerReferralSubStatus_Expired  MicrosoftPartnerReferralSubStatus = "Expired"
)

// All allowed values of MicrosoftPartnerReferralSubStatus enum
var AllowedMicrosoftPartnerReferralSubStatusEnumValues = []MicrosoftPartnerReferralSubStatus{
	"",
	"None",
	"Pending",
	"Received",
	"Won",
	"Lost",
	"Declined",
	"Expired",
}

func (v *MicrosoftPartnerReferralSubStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MicrosoftPartnerReferralSubStatus(value)
	for _, existing := range AllowedMicrosoftPartnerReferralSubStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MicrosoftPartnerReferralSubStatus", value)
}

// NewMicrosoftPartnerReferralSubStatusFromValue returns a pointer to a valid MicrosoftPartnerReferralSubStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMicrosoftPartnerReferralSubStatusFromValue(v string) (*MicrosoftPartnerReferralSubStatus, error) {
	ev := MicrosoftPartnerReferralSubStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MicrosoftPartnerReferralSubStatus: valid values are %v", v, AllowedMicrosoftPartnerReferralSubStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MicrosoftPartnerReferralSubStatus) IsValid() bool {
	for _, existing := range AllowedMicrosoftPartnerReferralSubStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MicrosoftPartnerReferralSubStatus value
func (v MicrosoftPartnerReferralSubStatus) Ptr() *MicrosoftPartnerReferralSubStatus {
	return &v
}

type NullableMicrosoftPartnerReferralSubStatus struct {
	value *MicrosoftPartnerReferralSubStatus
	isSet bool
}

func (v NullableMicrosoftPartnerReferralSubStatus) Get() *MicrosoftPartnerReferralSubStatus {
	return v.value
}

func (v *NullableMicrosoftPartnerReferralSubStatus) Set(val *MicrosoftPartnerReferralSubStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftPartnerReferralSubStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftPartnerReferralSubStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftPartnerReferralSubStatus(val *MicrosoftPartnerReferralSubStatus) *NullableMicrosoftPartnerReferralSubStatus {
	return &NullableMicrosoftPartnerReferralSubStatus{value: val, isSet: true}
}

func (v NullableMicrosoftPartnerReferralSubStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftPartnerReferralSubStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
