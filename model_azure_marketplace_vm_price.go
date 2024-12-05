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

// checks if the AzureMarketplaceVmPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceVmPrice{}

// AzureMarketplaceVmPrice struct for AzureMarketplaceVmPrice
type AzureMarketplaceVmPrice struct {
	PatternProperties *map[string]AzureMarketplaceVmPricePropertyItem `json:"patternProperties,omitempty"`
}

// NewAzureMarketplaceVmPrice instantiates a new AzureMarketplaceVmPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceVmPrice() *AzureMarketplaceVmPrice {
	this := AzureMarketplaceVmPrice{}
	return &this
}

// NewAzureMarketplaceVmPriceWithDefaults instantiates a new AzureMarketplaceVmPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceVmPriceWithDefaults() *AzureMarketplaceVmPrice {
	this := AzureMarketplaceVmPrice{}
	return &this
}

// GetPatternProperties returns the PatternProperties field value if set, zero value otherwise.
func (o *AzureMarketplaceVmPrice) GetPatternProperties() map[string]AzureMarketplaceVmPricePropertyItem {
	if o == nil || IsNil(o.PatternProperties) {
		var ret map[string]AzureMarketplaceVmPricePropertyItem
		return ret
	}
	return *o.PatternProperties
}

// GetPatternPropertiesOk returns a tuple with the PatternProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceVmPrice) GetPatternPropertiesOk() (*map[string]AzureMarketplaceVmPricePropertyItem, bool) {
	if o == nil || IsNil(o.PatternProperties) {
		return nil, false
	}
	return o.PatternProperties, true
}

// HasPatternProperties returns a boolean if a field has been set.
func (o *AzureMarketplaceVmPrice) HasPatternProperties() bool {
	if o != nil && !IsNil(o.PatternProperties) {
		return true
	}

	return false
}

// SetPatternProperties gets a reference to the given map[string]AzureMarketplaceVmPricePropertyItem and assigns it to the PatternProperties field.
func (o *AzureMarketplaceVmPrice) SetPatternProperties(v map[string]AzureMarketplaceVmPricePropertyItem) {
	o.PatternProperties = &v
}

func (o AzureMarketplaceVmPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceVmPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PatternProperties) {
		toSerialize["patternProperties"] = o.PatternProperties
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceVmPrice struct {
	value *AzureMarketplaceVmPrice
	isSet bool
}

func (v NullableAzureMarketplaceVmPrice) Get() *AzureMarketplaceVmPrice {
	return v.value
}

func (v *NullableAzureMarketplaceVmPrice) Set(val *AzureMarketplaceVmPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceVmPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceVmPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceVmPrice(val *AzureMarketplaceVmPrice) *NullableAzureMarketplaceVmPrice {
	return &NullableAzureMarketplaceVmPrice{value: val, isSet: true}
}

func (v NullableAzureMarketplaceVmPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceVmPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
