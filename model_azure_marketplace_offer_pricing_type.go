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

// AzureMarketplaceOfferPricingType the model 'AzureMarketplaceOfferPricingType'
type AzureMarketplaceOfferPricingType string

// List of AzureMarketplaceOfferPricingType
const (
	AzureMarketplaceOfferPricingType_Unknown AzureMarketplaceOfferPricingType = ""
	AzureMarketplaceOfferPricingType_EditExistingOfferPricingOnly AzureMarketplaceOfferPricingType = "editExistingOfferPricingOnly"
	AzureMarketplaceOfferPricingType_SaasNewCustomizedPlans AzureMarketplaceOfferPricingType = "saasNewCustomizedPlans"
	AzureMarketplaceOfferPricingType_VmSoftwareReservations AzureMarketplaceOfferPricingType = "vmSoftwareReservations"
)

// All allowed values of AzureMarketplaceOfferPricingType enum
var AllowedAzureMarketplaceOfferPricingTypeEnumValues = []AzureMarketplaceOfferPricingType{
	"",
	"editExistingOfferPricingOnly",
	"saasNewCustomizedPlans",
	"vmSoftwareReservations",
}

func (v *AzureMarketplaceOfferPricingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AzureMarketplaceOfferPricingType(value)
	for _, existing := range AllowedAzureMarketplaceOfferPricingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AzureMarketplaceOfferPricingType", value)
}

// NewAzureMarketplaceOfferPricingTypeFromValue returns a pointer to a valid AzureMarketplaceOfferPricingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureMarketplaceOfferPricingTypeFromValue(v string) (*AzureMarketplaceOfferPricingType, error) {
	ev := AzureMarketplaceOfferPricingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureMarketplaceOfferPricingType: valid values are %v", v, AllowedAzureMarketplaceOfferPricingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureMarketplaceOfferPricingType) IsValid() bool {
	for _, existing := range AllowedAzureMarketplaceOfferPricingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureMarketplaceOfferPricingType value
func (v AzureMarketplaceOfferPricingType) Ptr() *AzureMarketplaceOfferPricingType {
	return &v
}

type NullableAzureMarketplaceOfferPricingType struct {
	value *AzureMarketplaceOfferPricingType
	isSet bool
}

func (v NullableAzureMarketplaceOfferPricingType) Get() *AzureMarketplaceOfferPricingType {
	return v.value
}

func (v *NullableAzureMarketplaceOfferPricingType) Set(val *AzureMarketplaceOfferPricingType) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceOfferPricingType) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceOfferPricingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceOfferPricingType(val *AzureMarketplaceOfferPricingType) *NullableAzureMarketplaceOfferPricingType {
	return &NullableAzureMarketplaceOfferPricingType{value: val, isSet: true}
}

func (v NullableAzureMarketplaceOfferPricingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceOfferPricingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

