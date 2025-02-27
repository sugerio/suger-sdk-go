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

// PartnerService the model 'PartnerService'
type PartnerService string

// List of PartnerService
const (
	PartnerService_DEFAULT     PartnerService = "DEFAULT"
	PartnerService_ACE         PartnerService = "ACE"
	PartnerService_BIGQUERY    PartnerService = "BIGQUERY"
	PartnerService_BILLING     PartnerService = "BILLING"
	PartnerService_CHATBOT     PartnerService = "CHATBOT"
	PartnerService_COSELL      PartnerService = "COSELL"
	PartnerService_CRM         PartnerService = "CRM"
	PartnerService_CPQ         PartnerService = "CPQ"
	PartnerService_DATABASE    PartnerService = "DATABASE"
	PartnerService_DRIVE       PartnerService = "DRIVE"
	PartnerService_EMAIL       PartnerService = "EMAIL"
	PartnerService_MARKETPLACE PartnerService = "MARKETPLACE"
	PartnerService_NETSUITE    PartnerService = "NETSUITE"
	PartnerService_PAYMENT     PartnerService = "PAYMENT"
	PartnerService_QUICKBOOKS  PartnerService = "QUICKBOOKS"
	PartnerService_STORAGE     PartnerService = "STORAGE"
	PartnerService_TEAMS       PartnerService = "TEAMS"
)

// All allowed values of PartnerService enum
var AllowedPartnerServiceEnumValues = []PartnerService{
	"DEFAULT",
	"ACE",
	"BIGQUERY",
	"BILLING",
	"CHATBOT",
	"COSELL",
	"CRM",
	"CPQ",
	"DATABASE",
	"DRIVE",
	"EMAIL",
	"MARKETPLACE",
	"NETSUITE",
	"PAYMENT",
	"QUICKBOOKS",
	"STORAGE",
	"TEAMS",
}

func (v *PartnerService) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartnerService(value)
	for _, existing := range AllowedPartnerServiceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartnerService", value)
}

// NewPartnerServiceFromValue returns a pointer to a valid PartnerService
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartnerServiceFromValue(v string) (*PartnerService, error) {
	ev := PartnerService(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartnerService: valid values are %v", v, AllowedPartnerServiceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartnerService) IsValid() bool {
	for _, existing := range AllowedPartnerServiceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartnerService value
func (v PartnerService) Ptr() *PartnerService {
	return &v
}

type NullablePartnerService struct {
	value *PartnerService
	isSet bool
}

func (v NullablePartnerService) Get() *PartnerService {
	return v.value
}

func (v *NullablePartnerService) Set(val *PartnerService) {
	v.value = val
	v.isSet = true
}

func (v NullablePartnerService) IsSet() bool {
	return v.isSet
}

func (v *NullablePartnerService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartnerService(val *PartnerService) *NullablePartnerService {
	return &NullablePartnerService{value: val, isSet: true}
}

func (v NullablePartnerService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartnerService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
