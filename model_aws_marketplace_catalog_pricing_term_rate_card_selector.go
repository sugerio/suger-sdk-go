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

// checks if the AwsMarketplaceCatalogPricingTermRateCardSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceCatalogPricingTermRateCardSelector{}

// AwsMarketplaceCatalogPricingTermRateCardSelector struct for AwsMarketplaceCatalogPricingTermRateCardSelector
type AwsMarketplaceCatalogPricingTermRateCardSelector struct {
	// At this time, only \"Duration\" is supported.
	Type *string `json:"Type,omitempty"`
	// ISO 8601 duration format. For example, \"P1M\" represents one month.
	Value *string `json:"Value,omitempty"`
}

// NewAwsMarketplaceCatalogPricingTermRateCardSelector instantiates a new AwsMarketplaceCatalogPricingTermRateCardSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceCatalogPricingTermRateCardSelector() *AwsMarketplaceCatalogPricingTermRateCardSelector {
	this := AwsMarketplaceCatalogPricingTermRateCardSelector{}
	return &this
}

// NewAwsMarketplaceCatalogPricingTermRateCardSelectorWithDefaults instantiates a new AwsMarketplaceCatalogPricingTermRateCardSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceCatalogPricingTermRateCardSelectorWithDefaults() *AwsMarketplaceCatalogPricingTermRateCardSelector {
	this := AwsMarketplaceCatalogPricingTermRateCardSelector{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AwsMarketplaceCatalogPricingTermRateCardSelector) SetValue(v string) {
	o.Value = &v
}

func (o AwsMarketplaceCatalogPricingTermRateCardSelector) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceCatalogPricingTermRateCardSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceCatalogPricingTermRateCardSelector struct {
	value *AwsMarketplaceCatalogPricingTermRateCardSelector
	isSet bool
}

func (v NullableAwsMarketplaceCatalogPricingTermRateCardSelector) Get() *AwsMarketplaceCatalogPricingTermRateCardSelector {
	return v.value
}

func (v *NullableAwsMarketplaceCatalogPricingTermRateCardSelector) Set(val *AwsMarketplaceCatalogPricingTermRateCardSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceCatalogPricingTermRateCardSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceCatalogPricingTermRateCardSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceCatalogPricingTermRateCardSelector(val *AwsMarketplaceCatalogPricingTermRateCardSelector) *NullableAwsMarketplaceCatalogPricingTermRateCardSelector {
	return &NullableAwsMarketplaceCatalogPricingTermRateCardSelector{value: val, isSet: true}
}

func (v NullableAwsMarketplaceCatalogPricingTermRateCardSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceCatalogPricingTermRateCardSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
