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

// AzureMarketplaceListingAssetType the model 'AzureMarketplaceListingAssetType'
type AzureMarketplaceListingAssetType string

// List of AzureMarketplaceListingAssetType
const (
	AzureMarketplaceListingAssetType_AzureLogoSmall      AzureMarketplaceListingAssetType = "azureLogoSmall"
	AzureMarketplaceListingAssetType_AzureLogoMedium     AzureMarketplaceListingAssetType = "azureLogoMedium"
	AzureMarketplaceListingAssetType_AzureLogoLarge      AzureMarketplaceListingAssetType = "azureLogoLarge"
	AzureMarketplaceListingAssetType_AzureLogoWide       AzureMarketplaceListingAssetType = "azureLogoWide"
	AzureMarketplaceListingAssetType_AzureLogoScreenshot AzureMarketplaceListingAssetType = "azureLogoScreenshot"
	AzureMarketplaceListingAssetType_AzureLogoHero       AzureMarketplaceListingAssetType = "azureLogoHero"
	AzureMarketplaceListingAssetType_PdfDocument         AzureMarketplaceListingAssetType = "pdfDocument"
)

// All allowed values of AzureMarketplaceListingAssetType enum
var AllowedAzureMarketplaceListingAssetTypeEnumValues = []AzureMarketplaceListingAssetType{
	"azureLogoSmall",
	"azureLogoMedium",
	"azureLogoLarge",
	"azureLogoWide",
	"azureLogoScreenshot",
	"azureLogoHero",
	"pdfDocument",
}

func (v *AzureMarketplaceListingAssetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AzureMarketplaceListingAssetType(value)
	for _, existing := range AllowedAzureMarketplaceListingAssetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AzureMarketplaceListingAssetType", value)
}

// NewAzureMarketplaceListingAssetTypeFromValue returns a pointer to a valid AzureMarketplaceListingAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAzureMarketplaceListingAssetTypeFromValue(v string) (*AzureMarketplaceListingAssetType, error) {
	ev := AzureMarketplaceListingAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AzureMarketplaceListingAssetType: valid values are %v", v, AllowedAzureMarketplaceListingAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AzureMarketplaceListingAssetType) IsValid() bool {
	for _, existing := range AllowedAzureMarketplaceListingAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AzureMarketplaceListingAssetType value
func (v AzureMarketplaceListingAssetType) Ptr() *AzureMarketplaceListingAssetType {
	return &v
}

type NullableAzureMarketplaceListingAssetType struct {
	value *AzureMarketplaceListingAssetType
	isSet bool
}

func (v NullableAzureMarketplaceListingAssetType) Get() *AzureMarketplaceListingAssetType {
	return v.value
}

func (v *NullableAzureMarketplaceListingAssetType) Set(val *AzureMarketplaceListingAssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceListingAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceListingAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceListingAssetType(val *AzureMarketplaceListingAssetType) *NullableAzureMarketplaceListingAssetType {
	return &NullableAzureMarketplaceListingAssetType{value: val, isSet: true}
}

func (v NullableAzureMarketplaceListingAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceListingAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}