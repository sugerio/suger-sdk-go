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

// PrivateOfferDiscountType the model 'PrivateOfferDiscountType'
type PrivateOfferDiscountType string

// List of PrivateOfferDiscountType
const (
	PrivateOfferDiscountType_absolute   PrivateOfferDiscountType = "absolute"
	PrivateOfferDiscountType_percentage PrivateOfferDiscountType = "percentage"
)

// All allowed values of PrivateOfferDiscountType enum
var AllowedPrivateOfferDiscountTypeEnumValues = []PrivateOfferDiscountType{
	"absolute",
	"percentage",
}

func (v *PrivateOfferDiscountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivateOfferDiscountType(value)
	for _, existing := range AllowedPrivateOfferDiscountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivateOfferDiscountType", value)
}

// NewPrivateOfferDiscountTypeFromValue returns a pointer to a valid PrivateOfferDiscountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivateOfferDiscountTypeFromValue(v string) (*PrivateOfferDiscountType, error) {
	ev := PrivateOfferDiscountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivateOfferDiscountType: valid values are %v", v, AllowedPrivateOfferDiscountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivateOfferDiscountType) IsValid() bool {
	for _, existing := range AllowedPrivateOfferDiscountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrivateOfferDiscountType value
func (v PrivateOfferDiscountType) Ptr() *PrivateOfferDiscountType {
	return &v
}

type NullablePrivateOfferDiscountType struct {
	value *PrivateOfferDiscountType
	isSet bool
}

func (v NullablePrivateOfferDiscountType) Get() *PrivateOfferDiscountType {
	return v.value
}

func (v *NullablePrivateOfferDiscountType) Set(val *PrivateOfferDiscountType) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateOfferDiscountType) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateOfferDiscountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateOfferDiscountType(val *PrivateOfferDiscountType) *NullablePrivateOfferDiscountType {
	return &NullablePrivateOfferDiscountType{value: val, isSet: true}
}

func (v NullablePrivateOfferDiscountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateOfferDiscountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
