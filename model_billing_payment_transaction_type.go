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

// BillingPaymentTransactionType the model 'BillingPaymentTransactionType'
type BillingPaymentTransactionType string

// List of BillingPaymentTransactionType
const (
	BillingPaymentTransactionType_CHARGE BillingPaymentTransactionType = "CHARGE"
	BillingPaymentTransactionType_REFUND BillingPaymentTransactionType = "REFUND"
)

// All allowed values of BillingPaymentTransactionType enum
var AllowedBillingPaymentTransactionTypeEnumValues = []BillingPaymentTransactionType{
	"CHARGE",
	"REFUND",
}

func (v *BillingPaymentTransactionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BillingPaymentTransactionType(value)
	for _, existing := range AllowedBillingPaymentTransactionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BillingPaymentTransactionType", value)
}

// NewBillingPaymentTransactionTypeFromValue returns a pointer to a valid BillingPaymentTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBillingPaymentTransactionTypeFromValue(v string) (*BillingPaymentTransactionType, error) {
	ev := BillingPaymentTransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BillingPaymentTransactionType: valid values are %v", v, AllowedBillingPaymentTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BillingPaymentTransactionType) IsValid() bool {
	for _, existing := range AllowedBillingPaymentTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BillingPaymentTransactionType value
func (v BillingPaymentTransactionType) Ptr() *BillingPaymentTransactionType {
	return &v
}

type NullableBillingPaymentTransactionType struct {
	value *BillingPaymentTransactionType
	isSet bool
}

func (v NullableBillingPaymentTransactionType) Get() *BillingPaymentTransactionType {
	return v.value
}

func (v *NullableBillingPaymentTransactionType) Set(val *BillingPaymentTransactionType) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPaymentTransactionType) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPaymentTransactionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPaymentTransactionType(val *BillingPaymentTransactionType) *NullableBillingPaymentTransactionType {
	return &NullableBillingPaymentTransactionType{value: val, isSet: true}
}

func (v NullableBillingPaymentTransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPaymentTransactionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
