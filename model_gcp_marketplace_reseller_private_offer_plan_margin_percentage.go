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

// checks if the GcpMarketplaceResellerPrivateOfferPlanMarginPercentage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceResellerPrivateOfferPlanMarginPercentage{}

// GcpMarketplaceResellerPrivateOfferPlanMarginPercentage struct for GcpMarketplaceResellerPrivateOfferPlanMarginPercentage
type GcpMarketplaceResellerPrivateOfferPlanMarginPercentage struct {
	// The percentage value of the margin in the range of 0 to 100.
	Value *string `json:"value,omitempty"`
}

// NewGcpMarketplaceResellerPrivateOfferPlanMarginPercentage instantiates a new GcpMarketplaceResellerPrivateOfferPlanMarginPercentage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceResellerPrivateOfferPlanMarginPercentage() *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage {
	this := GcpMarketplaceResellerPrivateOfferPlanMarginPercentage{}
	return &this
}

// NewGcpMarketplaceResellerPrivateOfferPlanMarginPercentageWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlanMarginPercentage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceResellerPrivateOfferPlanMarginPercentageWithDefaults() *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage {
	this := GcpMarketplaceResellerPrivateOfferPlanMarginPercentage{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) SetValue(v string) {
	o.Value = &v
}

func (o GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage struct {
	value *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage
	isSet bool
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage) Get() *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage {
	return v.value
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage) Set(val *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage(val *GcpMarketplaceResellerPrivateOfferPlanMarginPercentage) *NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage {
	return &NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage{value: val, isSet: true}
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanMarginPercentage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
