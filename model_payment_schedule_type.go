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

// PaymentScheduleType the model 'PaymentScheduleType'
type PaymentScheduleType string

// List of PaymentScheduleType
const (
	PaymentScheduleType_UNKNOWN PaymentScheduleType = ""
	PaymentScheduleType_PREPAY  PaymentScheduleType = "PREPAY"
	PaymentScheduleType_POSTPAY PaymentScheduleType = "POSTPAY"
)

// All allowed values of PaymentScheduleType enum
var AllowedPaymentScheduleTypeEnumValues = []PaymentScheduleType{
	"",
	"PREPAY",
	"POSTPAY",
}

func (v *PaymentScheduleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentScheduleType(value)
	for _, existing := range AllowedPaymentScheduleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentScheduleType", value)
}

// NewPaymentScheduleTypeFromValue returns a pointer to a valid PaymentScheduleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentScheduleTypeFromValue(v string) (*PaymentScheduleType, error) {
	ev := PaymentScheduleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentScheduleType: valid values are %v", v, AllowedPaymentScheduleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentScheduleType) IsValid() bool {
	for _, existing := range AllowedPaymentScheduleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentScheduleType value
func (v PaymentScheduleType) Ptr() *PaymentScheduleType {
	return &v
}

type NullablePaymentScheduleType struct {
	value *PaymentScheduleType
	isSet bool
}

func (v NullablePaymentScheduleType) Get() *PaymentScheduleType {
	return v.value
}

func (v *NullablePaymentScheduleType) Set(val *PaymentScheduleType) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentScheduleType) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentScheduleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentScheduleType(val *PaymentScheduleType) *NullablePaymentScheduleType {
	return &NullablePaymentScheduleType{value: val, isSet: true}
}

func (v NullablePaymentScheduleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentScheduleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}