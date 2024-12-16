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

// checks if the AlibabaMarketplaceProductShopInfoWangWangs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlibabaMarketplaceProductShopInfoWangWangs{}

// AlibabaMarketplaceProductShopInfoWangWangs struct for AlibabaMarketplaceProductShopInfoWangWangs
type AlibabaMarketplaceProductShopInfoWangWangs struct {
	WangWang []AlibabaMarketplaceProductShopInfoWangWang `json:"WangWang,omitempty"`
}

// NewAlibabaMarketplaceProductShopInfoWangWangs instantiates a new AlibabaMarketplaceProductShopInfoWangWangs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlibabaMarketplaceProductShopInfoWangWangs() *AlibabaMarketplaceProductShopInfoWangWangs {
	this := AlibabaMarketplaceProductShopInfoWangWangs{}
	return &this
}

// NewAlibabaMarketplaceProductShopInfoWangWangsWithDefaults instantiates a new AlibabaMarketplaceProductShopInfoWangWangs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlibabaMarketplaceProductShopInfoWangWangsWithDefaults() *AlibabaMarketplaceProductShopInfoWangWangs {
	this := AlibabaMarketplaceProductShopInfoWangWangs{}
	return &this
}

// GetWangWang returns the WangWang field value if set, zero value otherwise.
func (o *AlibabaMarketplaceProductShopInfoWangWangs) GetWangWang() []AlibabaMarketplaceProductShopInfoWangWang {
	if o == nil || IsNil(o.WangWang) {
		var ret []AlibabaMarketplaceProductShopInfoWangWang
		return ret
	}
	return o.WangWang
}

// GetWangWangOk returns a tuple with the WangWang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlibabaMarketplaceProductShopInfoWangWangs) GetWangWangOk() ([]AlibabaMarketplaceProductShopInfoWangWang, bool) {
	if o == nil || IsNil(o.WangWang) {
		return nil, false
	}
	return o.WangWang, true
}

// HasWangWang returns a boolean if a field has been set.
func (o *AlibabaMarketplaceProductShopInfoWangWangs) HasWangWang() bool {
	if o != nil && !IsNil(o.WangWang) {
		return true
	}

	return false
}

// SetWangWang gets a reference to the given []AlibabaMarketplaceProductShopInfoWangWang and assigns it to the WangWang field.
func (o *AlibabaMarketplaceProductShopInfoWangWangs) SetWangWang(v []AlibabaMarketplaceProductShopInfoWangWang) {
	o.WangWang = v
}

func (o AlibabaMarketplaceProductShopInfoWangWangs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlibabaMarketplaceProductShopInfoWangWangs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WangWang) {
		toSerialize["WangWang"] = o.WangWang
	}
	return toSerialize, nil
}

type NullableAlibabaMarketplaceProductShopInfoWangWangs struct {
	value *AlibabaMarketplaceProductShopInfoWangWangs
	isSet bool
}

func (v NullableAlibabaMarketplaceProductShopInfoWangWangs) Get() *AlibabaMarketplaceProductShopInfoWangWangs {
	return v.value
}

func (v *NullableAlibabaMarketplaceProductShopInfoWangWangs) Set(val *AlibabaMarketplaceProductShopInfoWangWangs) {
	v.value = val
	v.isSet = true
}

func (v NullableAlibabaMarketplaceProductShopInfoWangWangs) IsSet() bool {
	return v.isSet
}

func (v *NullableAlibabaMarketplaceProductShopInfoWangWangs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlibabaMarketplaceProductShopInfoWangWangs(val *AlibabaMarketplaceProductShopInfoWangWangs) *NullableAlibabaMarketplaceProductShopInfoWangWangs {
	return &NullableAlibabaMarketplaceProductShopInfoWangWangs{value: val, isSet: true}
}

func (v NullableAlibabaMarketplaceProductShopInfoWangWangs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlibabaMarketplaceProductShopInfoWangWangs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
