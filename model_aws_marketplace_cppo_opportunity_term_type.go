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

// AwsMarketplaceCppoOpportunityTermType the model 'AwsMarketplaceCppoOpportunityTermType'
type AwsMarketplaceCppoOpportunityTermType string

// List of AwsMarketplaceCppoOpportunityTermType
const (
	AwsMarketplaceCppoOpportunityTermType_BuyerTargetingTerm                   AwsMarketplaceCppoOpportunityTermType = "BuyerTargetingTerm"
	AwsMarketplaceCppoOpportunityTermType_UpdateAvailability                   AwsMarketplaceCppoOpportunityTermType = "UpdateAvailability"
	AwsMarketplaceCppoOpportunityTermType_BuyerValidityTerm                    AwsMarketplaceCppoOpportunityTermType = "BuyerValidityTerm"
	AwsMarketplaceCppoOpportunityTermType_BuyerLegalTerm                       AwsMarketplaceCppoOpportunityTermType = "BuyerLegalTerm"
	AwsMarketplaceCppoOpportunityTermType_ResaleLegalTerm                      AwsMarketplaceCppoOpportunityTermType = "ResaleLegalTerm"
	AwsMarketplaceCppoOpportunityTermType_ResaleUsageBasedPricingTerm          AwsMarketplaceCppoOpportunityTermType = "ResaleUsageBasedPricingTerm"
	AwsMarketplaceCppoOpportunityTermType_ResaleConfigurableUpfrontPricingTerm AwsMarketplaceCppoOpportunityTermType = "ResaleConfigurableUpfrontPricingTerm"
	AwsMarketplaceCppoOpportunityTermType_ResaleFixedUpfrontPricingTerm        AwsMarketplaceCppoOpportunityTermType = "ResaleFixedUpfrontPricingTerm"
	AwsMarketplaceCppoOpportunityTermType_ResalePaymentScheduleTerm            AwsMarketplaceCppoOpportunityTermType = "ResalePaymentScheduleTerm"
)

// All allowed values of AwsMarketplaceCppoOpportunityTermType enum
var AllowedAwsMarketplaceCppoOpportunityTermTypeEnumValues = []AwsMarketplaceCppoOpportunityTermType{
	"BuyerTargetingTerm",
	"UpdateAvailability",
	"BuyerValidityTerm",
	"BuyerLegalTerm",
	"ResaleLegalTerm",
	"ResaleUsageBasedPricingTerm",
	"ResaleConfigurableUpfrontPricingTerm",
	"ResaleFixedUpfrontPricingTerm",
	"ResalePaymentScheduleTerm",
}

func (v *AwsMarketplaceCppoOpportunityTermType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AwsMarketplaceCppoOpportunityTermType(value)
	for _, existing := range AllowedAwsMarketplaceCppoOpportunityTermTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AwsMarketplaceCppoOpportunityTermType", value)
}

// NewAwsMarketplaceCppoOpportunityTermTypeFromValue returns a pointer to a valid AwsMarketplaceCppoOpportunityTermType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAwsMarketplaceCppoOpportunityTermTypeFromValue(v string) (*AwsMarketplaceCppoOpportunityTermType, error) {
	ev := AwsMarketplaceCppoOpportunityTermType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AwsMarketplaceCppoOpportunityTermType: valid values are %v", v, AllowedAwsMarketplaceCppoOpportunityTermTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AwsMarketplaceCppoOpportunityTermType) IsValid() bool {
	for _, existing := range AllowedAwsMarketplaceCppoOpportunityTermTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AwsMarketplaceCppoOpportunityTermType value
func (v AwsMarketplaceCppoOpportunityTermType) Ptr() *AwsMarketplaceCppoOpportunityTermType {
	return &v
}

type NullableAwsMarketplaceCppoOpportunityTermType struct {
	value *AwsMarketplaceCppoOpportunityTermType
	isSet bool
}

func (v NullableAwsMarketplaceCppoOpportunityTermType) Get() *AwsMarketplaceCppoOpportunityTermType {
	return v.value
}

func (v *NullableAwsMarketplaceCppoOpportunityTermType) Set(val *AwsMarketplaceCppoOpportunityTermType) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceCppoOpportunityTermType) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceCppoOpportunityTermType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceCppoOpportunityTermType(val *AwsMarketplaceCppoOpportunityTermType) *NullableAwsMarketplaceCppoOpportunityTermType {
	return &NullableAwsMarketplaceCppoOpportunityTermType{value: val, isSet: true}
}

func (v NullableAwsMarketplaceCppoOpportunityTermType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceCppoOpportunityTermType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
