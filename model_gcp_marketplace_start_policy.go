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

// GcpMarketplaceStartPolicy the model 'GcpMarketplaceStartPolicy'
type GcpMarketplaceStartPolicy string

// List of GcpMarketplaceStartPolicy
const (
	GcpMarketplaceStartPolicy_UNKNOWN       GcpMarketplaceStartPolicy = ""
	GcpMarketplaceStartPolicy_ON_ACCEPTANCE GcpMarketplaceStartPolicy = "START_POLICY_ON_ACCEPTANCE"
	GcpMarketplaceStartPolicy_SCHEDULED     GcpMarketplaceStartPolicy = "START_POLICY_SCHEDULED"
)

// All allowed values of GcpMarketplaceStartPolicy enum
var AllowedGcpMarketplaceStartPolicyEnumValues = []GcpMarketplaceStartPolicy{
	"",
	"START_POLICY_ON_ACCEPTANCE",
	"START_POLICY_SCHEDULED",
}

func (v *GcpMarketplaceStartPolicy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GcpMarketplaceStartPolicy(value)
	for _, existing := range AllowedGcpMarketplaceStartPolicyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GcpMarketplaceStartPolicy", value)
}

// NewGcpMarketplaceStartPolicyFromValue returns a pointer to a valid GcpMarketplaceStartPolicy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGcpMarketplaceStartPolicyFromValue(v string) (*GcpMarketplaceStartPolicy, error) {
	ev := GcpMarketplaceStartPolicy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GcpMarketplaceStartPolicy: valid values are %v", v, AllowedGcpMarketplaceStartPolicyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GcpMarketplaceStartPolicy) IsValid() bool {
	for _, existing := range AllowedGcpMarketplaceStartPolicyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GcpMarketplaceStartPolicy value
func (v GcpMarketplaceStartPolicy) Ptr() *GcpMarketplaceStartPolicy {
	return &v
}

type NullableGcpMarketplaceStartPolicy struct {
	value *GcpMarketplaceStartPolicy
	isSet bool
}

func (v NullableGcpMarketplaceStartPolicy) Get() *GcpMarketplaceStartPolicy {
	return v.value
}

func (v *NullableGcpMarketplaceStartPolicy) Set(val *GcpMarketplaceStartPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceStartPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceStartPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceStartPolicy(val *GcpMarketplaceStartPolicy) *NullableGcpMarketplaceStartPolicy {
	return &NullableGcpMarketplaceStartPolicy{value: val, isSet: true}
}

func (v NullableGcpMarketplaceStartPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceStartPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}