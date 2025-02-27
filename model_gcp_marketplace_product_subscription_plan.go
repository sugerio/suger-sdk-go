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
)

// checks if the GcpMarketplaceProductSubscriptionPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductSubscriptionPlan{}

// GcpMarketplaceProductSubscriptionPlan struct for GcpMarketplaceProductSubscriptionPlan
type GcpMarketplaceProductSubscriptionPlan struct {
	// such as \"ONE_YEAR\", \"TWO_YEAR\", \"THREE_YEAR\"
	Period *string        `json:"period,omitempty"`
	Price  *GcpPriceValue `json:"price,omitempty"`
}

// NewGcpMarketplaceProductSubscriptionPlan instantiates a new GcpMarketplaceProductSubscriptionPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductSubscriptionPlan() *GcpMarketplaceProductSubscriptionPlan {
	this := GcpMarketplaceProductSubscriptionPlan{}
	return &this
}

// NewGcpMarketplaceProductSubscriptionPlanWithDefaults instantiates a new GcpMarketplaceProductSubscriptionPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductSubscriptionPlanWithDefaults() *GcpMarketplaceProductSubscriptionPlan {
	this := GcpMarketplaceProductSubscriptionPlan{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GcpMarketplaceProductSubscriptionPlan) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductSubscriptionPlan) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GcpMarketplaceProductSubscriptionPlan) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GcpMarketplaceProductSubscriptionPlan) SetPeriod(v string) {
	o.Period = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GcpMarketplaceProductSubscriptionPlan) GetPrice() GcpPriceValue {
	if o == nil || IsNil(o.Price) {
		var ret GcpPriceValue
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductSubscriptionPlan) GetPriceOk() (*GcpPriceValue, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *GcpMarketplaceProductSubscriptionPlan) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given GcpPriceValue and assigns it to the Price field.
func (o *GcpMarketplaceProductSubscriptionPlan) SetPrice(v GcpPriceValue) {
	o.Price = &v
}

func (o GcpMarketplaceProductSubscriptionPlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductSubscriptionPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductSubscriptionPlan struct {
	value *GcpMarketplaceProductSubscriptionPlan
	isSet bool
}

func (v NullableGcpMarketplaceProductSubscriptionPlan) Get() *GcpMarketplaceProductSubscriptionPlan {
	return v.value
}

func (v *NullableGcpMarketplaceProductSubscriptionPlan) Set(val *GcpMarketplaceProductSubscriptionPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductSubscriptionPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductSubscriptionPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductSubscriptionPlan(val *GcpMarketplaceProductSubscriptionPlan) *NullableGcpMarketplaceProductSubscriptionPlan {
	return &NullableGcpMarketplaceProductSubscriptionPlan{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductSubscriptionPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductSubscriptionPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
