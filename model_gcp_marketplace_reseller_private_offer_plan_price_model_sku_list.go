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

// checks if the GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList{}

// GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList struct for GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList
type GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList struct {
	Skus []string `json:"skus,omitempty"`
}

// NewGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList instantiates a new GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList() *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList {
	this := GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList{}
	return &this
}

// NewGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuListWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuListWithDefaults() *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList {
	this := GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList{}
	return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) GetSkus() []string {
	if o == nil || IsNil(o.Skus) {
		var ret []string
		return ret
	}
	return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) GetSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given []string and assigns it to the Skus field.
func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) SetSkus(v []string) {
	o.Skus = v
}

func (o GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList struct {
	value *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList
	isSet bool
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) Get() *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList {
	return v.value
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) Set(val *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList(val *GcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) *NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList {
	return &NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList{value: val, isSet: true}
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlanPriceModelSkuList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
