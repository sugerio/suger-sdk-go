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

// checks if the GcpMarketplacePrivateOfferPriceModelOverage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplacePrivateOfferPriceModelOverage{}

// GcpMarketplacePrivateOfferPriceModelOverage struct for GcpMarketplacePrivateOfferPriceModelOverage
type GcpMarketplacePrivateOfferPriceModelOverage struct {
	Discount *GcpMarketplacePrivateOfferPriceModelDiscount `json:"discount,omitempty"`
	// TODO: need to define the type
	SkuDiscounts []map[string]interface{} `json:"skuDiscounts,omitempty"`
}

// NewGcpMarketplacePrivateOfferPriceModelOverage instantiates a new GcpMarketplacePrivateOfferPriceModelOverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplacePrivateOfferPriceModelOverage() *GcpMarketplacePrivateOfferPriceModelOverage {
	this := GcpMarketplacePrivateOfferPriceModelOverage{}
	return &this
}

// NewGcpMarketplacePrivateOfferPriceModelOverageWithDefaults instantiates a new GcpMarketplacePrivateOfferPriceModelOverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplacePrivateOfferPriceModelOverageWithDefaults() *GcpMarketplacePrivateOfferPriceModelOverage {
	this := GcpMarketplacePrivateOfferPriceModelOverage{}
	return &this
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) GetDiscount() GcpMarketplacePrivateOfferPriceModelDiscount {
	if o == nil || IsNil(o.Discount) {
		var ret GcpMarketplacePrivateOfferPriceModelDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) GetDiscountOk() (*GcpMarketplacePrivateOfferPriceModelDiscount, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given GcpMarketplacePrivateOfferPriceModelDiscount and assigns it to the Discount field.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) SetDiscount(v GcpMarketplacePrivateOfferPriceModelDiscount) {
	o.Discount = &v
}

// GetSkuDiscounts returns the SkuDiscounts field value if set, zero value otherwise.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) GetSkuDiscounts() []map[string]interface{} {
	if o == nil || IsNil(o.SkuDiscounts) {
		var ret []map[string]interface{}
		return ret
	}
	return o.SkuDiscounts
}

// GetSkuDiscountsOk returns a tuple with the SkuDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) GetSkuDiscountsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.SkuDiscounts) {
		return nil, false
	}
	return o.SkuDiscounts, true
}

// HasSkuDiscounts returns a boolean if a field has been set.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) HasSkuDiscounts() bool {
	if o != nil && !IsNil(o.SkuDiscounts) {
		return true
	}

	return false
}

// SetSkuDiscounts gets a reference to the given []map[string]interface{} and assigns it to the SkuDiscounts field.
func (o *GcpMarketplacePrivateOfferPriceModelOverage) SetSkuDiscounts(v []map[string]interface{}) {
	o.SkuDiscounts = v
}

func (o GcpMarketplacePrivateOfferPriceModelOverage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplacePrivateOfferPriceModelOverage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.SkuDiscounts) {
		toSerialize["skuDiscounts"] = o.SkuDiscounts
	}
	return toSerialize, nil
}

type NullableGcpMarketplacePrivateOfferPriceModelOverage struct {
	value *GcpMarketplacePrivateOfferPriceModelOverage
	isSet bool
}

func (v NullableGcpMarketplacePrivateOfferPriceModelOverage) Get() *GcpMarketplacePrivateOfferPriceModelOverage {
	return v.value
}

func (v *NullableGcpMarketplacePrivateOfferPriceModelOverage) Set(val *GcpMarketplacePrivateOfferPriceModelOverage) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplacePrivateOfferPriceModelOverage) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplacePrivateOfferPriceModelOverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplacePrivateOfferPriceModelOverage(val *GcpMarketplacePrivateOfferPriceModelOverage) *NullableGcpMarketplacePrivateOfferPriceModelOverage {
	return &NullableGcpMarketplacePrivateOfferPriceModelOverage{value: val, isSet: true}
}

func (v NullableGcpMarketplacePrivateOfferPriceModelOverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplacePrivateOfferPriceModelOverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
