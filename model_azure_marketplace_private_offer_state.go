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

// AzureMarketplacePrivateOfferState the model 'AzureMarketplacePrivateOfferState'
type AzureMarketplacePrivateOfferState string

// List of AzureMarketplacePrivateOfferState
const (
	AzureMarketplacePrivateOfferState_draft AzureMarketplacePrivateOfferState = "draft"
	AzureMarketplacePrivateOfferState_live AzureMarketplacePrivateOfferState = "live"
	AzureMarketplacePrivateOfferState_deleted AzureMarketplacePrivateOfferState = "deleted"
	AzureMarketplacePrivateOfferState_withdrawn AzureMarketplacePrivateOfferState = "withdrawn"
)

// All allowed values of AzureMarketplacePrivateOfferState enum
var AllowedAzureMarketplacePrivateOfferStateEnumValues = []AzureMarketplacePrivateOfferState{
	"draft",
	"live",
	"deleted",
	"withdrawn",
}

func (v *AzureMarketplacePrivateOfferState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AzureMarketplacePrivateOfferState(value)
	for _, existing := range AllowedAzureMarketplacePrivateOfferStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AzureMarketplacePrivateOfferState", value)
}

// NewAzureMarketplacePrivateOfferStateFromValue returns a pointer to a valid AzureMarketplacePrivateOfferState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureMarketplacePrivateOfferStateFromValue(v string) (*AzureMarketplacePrivateOfferState, error) {
	ev := AzureMarketplacePrivateOfferState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureMarketplacePrivateOfferState: valid values are %v", v, AllowedAzureMarketplacePrivateOfferStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureMarketplacePrivateOfferState) IsValid() bool {
	for _, existing := range AllowedAzureMarketplacePrivateOfferStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureMarketplacePrivateOfferState value
func (v AzureMarketplacePrivateOfferState) Ptr() *AzureMarketplacePrivateOfferState {
	return &v
}

type NullableAzureMarketplacePrivateOfferState struct {
	value *AzureMarketplacePrivateOfferState
	isSet bool
}

func (v NullableAzureMarketplacePrivateOfferState) Get() *AzureMarketplacePrivateOfferState {
	return v.value
}

func (v *NullableAzureMarketplacePrivateOfferState) Set(val *AzureMarketplacePrivateOfferState) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePrivateOfferState) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePrivateOfferState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePrivateOfferState(val *AzureMarketplacePrivateOfferState) *NullableAzureMarketplacePrivateOfferState {
	return &NullableAzureMarketplacePrivateOfferState{value: val, isSet: true}
}

func (v NullableAzureMarketplacePrivateOfferState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePrivateOfferState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

