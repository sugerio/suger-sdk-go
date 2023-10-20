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

// GcpMarketplacePrivateOfferState the model 'GcpMarketplacePrivateOfferState'
type GcpMarketplacePrivateOfferState string

// List of GcpMarketplacePrivateOfferState
const (
	GcpMarketplacePrivateOfferState_OFFER_PUBLISHED GcpMarketplacePrivateOfferState = "OFFER_PUBLISHED"
	GcpMarketplacePrivateOfferState_OFFER_ACTIVE GcpMarketplacePrivateOfferState = "OFFER_ACTIVE"
	GcpMarketplacePrivateOfferState_OFFER_ACTIVATING GcpMarketplacePrivateOfferState = "OFFER_ACTIVATING"
	GcpMarketplacePrivateOfferState_OFFER_LAPSED GcpMarketplacePrivateOfferState = "OFFER_LAPSED"
	GcpMarketplacePrivateOfferState_OFFER_EXPIRED GcpMarketplacePrivateOfferState = "OFFER_EXPIRED"
	GcpMarketplacePrivateOfferState_OFFER_CANCELLED GcpMarketplacePrivateOfferState = "OFFER_CANCELLED"
	GcpMarketplacePrivateOfferState_OFFER_UNAVAILABLE GcpMarketplacePrivateOfferState = "OFFER_UNAVAILABLE"
	GcpMarketplacePrivateOfferState_OFFER_DRAFT GcpMarketplacePrivateOfferState = "OFFER_DRAFT"
)

// All allowed values of GcpMarketplacePrivateOfferState enum
var AllowedGcpMarketplacePrivateOfferStateEnumValues = []GcpMarketplacePrivateOfferState{
	"OFFER_PUBLISHED",
	"OFFER_ACTIVE",
	"OFFER_ACTIVATING",
	"OFFER_LAPSED",
	"OFFER_EXPIRED",
	"OFFER_CANCELLED",
	"OFFER_UNAVAILABLE",
	"OFFER_DRAFT",
}

func (v *GcpMarketplacePrivateOfferState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GcpMarketplacePrivateOfferState(value)
	for _, existing := range AllowedGcpMarketplacePrivateOfferStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GcpMarketplacePrivateOfferState", value)
}

// NewGcpMarketplacePrivateOfferStateFromValue returns a pointer to a valid GcpMarketplacePrivateOfferState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGcpMarketplacePrivateOfferStateFromValue(v string) (*GcpMarketplacePrivateOfferState, error) {
	ev := GcpMarketplacePrivateOfferState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GcpMarketplacePrivateOfferState: valid values are %v", v, AllowedGcpMarketplacePrivateOfferStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GcpMarketplacePrivateOfferState) IsValid() bool {
	for _, existing := range AllowedGcpMarketplacePrivateOfferStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GcpMarketplacePrivateOfferState value
func (v GcpMarketplacePrivateOfferState) Ptr() *GcpMarketplacePrivateOfferState {
	return &v
}

type NullableGcpMarketplacePrivateOfferState struct {
	value *GcpMarketplacePrivateOfferState
	isSet bool
}

func (v NullableGcpMarketplacePrivateOfferState) Get() *GcpMarketplacePrivateOfferState {
	return v.value
}

func (v *NullableGcpMarketplacePrivateOfferState) Set(val *GcpMarketplacePrivateOfferState) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplacePrivateOfferState) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplacePrivateOfferState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplacePrivateOfferState(val *GcpMarketplacePrivateOfferState) *NullableGcpMarketplacePrivateOfferState {
	return &NullableGcpMarketplacePrivateOfferState{value: val, isSet: true}
}

func (v NullableGcpMarketplacePrivateOfferState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplacePrivateOfferState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

