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

// GcpMarketplaceOfferStartPolicy the model 'GcpMarketplaceOfferStartPolicy'
type GcpMarketplaceOfferStartPolicy string

// List of GcpMarketplaceOfferStartPolicy
const (
	GcpMarketplaceOfferStartPolicy_UNKNOWN   GcpMarketplaceOfferStartPolicy = ""
	GcpMarketplaceOfferStartPolicy_IMMEDIATE GcpMarketplaceOfferStartPolicy = "OFFER_START_POLICY_IMMEDIATE"
	GcpMarketplaceOfferStartPolicy_SCHEDULED GcpMarketplaceOfferStartPolicy = "OFFER_START_POLICY_SCHEDULED"
)

// All allowed values of GcpMarketplaceOfferStartPolicy enum
var AllowedGcpMarketplaceOfferStartPolicyEnumValues = []GcpMarketplaceOfferStartPolicy{
	"",
	"OFFER_START_POLICY_IMMEDIATE",
	"OFFER_START_POLICY_SCHEDULED",
}

func (v *GcpMarketplaceOfferStartPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GcpMarketplaceOfferStartPolicy(value)
	for _, existing := range AllowedGcpMarketplaceOfferStartPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GcpMarketplaceOfferStartPolicy", value)
}

// NewGcpMarketplaceOfferStartPolicyFromValue returns a pointer to a valid GcpMarketplaceOfferStartPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGcpMarketplaceOfferStartPolicyFromValue(v string) (*GcpMarketplaceOfferStartPolicy, error) {
	ev := GcpMarketplaceOfferStartPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GcpMarketplaceOfferStartPolicy: valid values are %v", v, AllowedGcpMarketplaceOfferStartPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GcpMarketplaceOfferStartPolicy) IsValid() bool {
	for _, existing := range AllowedGcpMarketplaceOfferStartPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GcpMarketplaceOfferStartPolicy value
func (v GcpMarketplaceOfferStartPolicy) Ptr() *GcpMarketplaceOfferStartPolicy {
	return &v
}

type NullableGcpMarketplaceOfferStartPolicy struct {
	value *GcpMarketplaceOfferStartPolicy
	isSet bool
}

func (v NullableGcpMarketplaceOfferStartPolicy) Get() *GcpMarketplaceOfferStartPolicy {
	return v.value
}

func (v *NullableGcpMarketplaceOfferStartPolicy) Set(val *GcpMarketplaceOfferStartPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceOfferStartPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceOfferStartPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceOfferStartPolicy(val *GcpMarketplaceOfferStartPolicy) *NullableGcpMarketplaceOfferStartPolicy {
	return &NullableGcpMarketplaceOfferStartPolicy{value: val, isSet: true}
}

func (v NullableGcpMarketplaceOfferStartPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceOfferStartPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
