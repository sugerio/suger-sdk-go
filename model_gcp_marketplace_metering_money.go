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
)

// checks if the GcpMarketplaceMeteringMoney type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceMeteringMoney{}

// GcpMarketplaceMeteringMoney struct for GcpMarketplaceMeteringMoney
type GcpMarketplaceMeteringMoney struct {
	// CurrencyCode: The three-letter currency code defined in ISO 4217.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Nanos: Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If `units` is positive, `nanos` must be positive or zero. If `units` is zero, `nanos` can be positive, zero, or negative. If `units` is negative, `nanos` must be negative or zero. For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	Nanos *int32 `json:"nanos,omitempty"`
	// Units: The whole units of the amount. For example if `currencyCode` is \"USD\", then 1 unit is one US dollar.
	Units *string `json:"units,omitempty"`
}

// NewGcpMarketplaceMeteringMoney instantiates a new GcpMarketplaceMeteringMoney object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceMeteringMoney() *GcpMarketplaceMeteringMoney {
	this := GcpMarketplaceMeteringMoney{}
	return &this
}

// NewGcpMarketplaceMeteringMoneyWithDefaults instantiates a new GcpMarketplaceMeteringMoney object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceMeteringMoneyWithDefaults() *GcpMarketplaceMeteringMoney {
	this := GcpMarketplaceMeteringMoney{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *GcpMarketplaceMeteringMoney) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceMeteringMoney) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GcpMarketplaceMeteringMoney) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *GcpMarketplaceMeteringMoney) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *GcpMarketplaceMeteringMoney) GetNanos() int32 {
	if o == nil || IsNil(o.Nanos) {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceMeteringMoney) GetNanosOk() (*int32, bool) {
	if o == nil || IsNil(o.Nanos) {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *GcpMarketplaceMeteringMoney) HasNanos() bool {
	if o != nil && !IsNil(o.Nanos) {
		return true
	}

	return false
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *GcpMarketplaceMeteringMoney) SetNanos(v int32) {
	o.Nanos = &v
}

// GetUnits returns the Units field value if set, zero value otherwise.
func (o *GcpMarketplaceMeteringMoney) GetUnits() string {
	if o == nil || IsNil(o.Units) {
		var ret string
		return ret
	}
	return *o.Units
}

// GetUnitsOk returns a tuple with the Units field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceMeteringMoney) GetUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return o.Units, true
}

// HasUnits returns a boolean if a field has been set.
func (o *GcpMarketplaceMeteringMoney) HasUnits() bool {
	if o != nil && !IsNil(o.Units) {
		return true
	}

	return false
}

// SetUnits gets a reference to the given string and assigns it to the Units field.
func (o *GcpMarketplaceMeteringMoney) SetUnits(v string) {
	o.Units = &v
}

func (o GcpMarketplaceMeteringMoney) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceMeteringMoney) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Nanos) {
		toSerialize["nanos"] = o.Nanos
	}
	if !IsNil(o.Units) {
		toSerialize["units"] = o.Units
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceMeteringMoney struct {
	value *GcpMarketplaceMeteringMoney
	isSet bool
}

func (v NullableGcpMarketplaceMeteringMoney) Get() *GcpMarketplaceMeteringMoney {
	return v.value
}

func (v *NullableGcpMarketplaceMeteringMoney) Set(val *GcpMarketplaceMeteringMoney) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceMeteringMoney) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceMeteringMoney) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceMeteringMoney(val *GcpMarketplaceMeteringMoney) *NullableGcpMarketplaceMeteringMoney {
	return &NullableGcpMarketplaceMeteringMoney{value: val, isSet: true}
}

func (v NullableGcpMarketplaceMeteringMoney) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceMeteringMoney) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
