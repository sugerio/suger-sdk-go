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

// checks if the GcpMarketplaceProductPurchaseOptionSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductPurchaseOptionSpec{}

// GcpMarketplaceProductPurchaseOptionSpec struct for GcpMarketplaceProductPurchaseOptionSpec
type GcpMarketplaceProductPurchaseOptionSpec struct {
	FeatureValues []GcpMarketplaceProductFeatureValue `json:"featureValues,omitempty"`
	// The plan ID, such as \"starter\", without the duration suffix, such as \"P1Y\".
	Name         *string                         `json:"name,omitempty"`
	PriceInfo    *GcpMarketplaceProductPriceInfo `json:"priceInfo,omitempty"`
	PurchaseMode *string                         `json:"purchaseMode,omitempty"`
	// such as \"Starter\"
	Title *string `json:"title,omitempty"`
}

// NewGcpMarketplaceProductPurchaseOptionSpec instantiates a new GcpMarketplaceProductPurchaseOptionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductPurchaseOptionSpec() *GcpMarketplaceProductPurchaseOptionSpec {
	this := GcpMarketplaceProductPurchaseOptionSpec{}
	return &this
}

// NewGcpMarketplaceProductPurchaseOptionSpecWithDefaults instantiates a new GcpMarketplaceProductPurchaseOptionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductPurchaseOptionSpecWithDefaults() *GcpMarketplaceProductPurchaseOptionSpec {
	this := GcpMarketplaceProductPurchaseOptionSpec{}
	return &this
}

// GetFeatureValues returns the FeatureValues field value if set, zero value otherwise.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetFeatureValues() []GcpMarketplaceProductFeatureValue {
	if o == nil || IsNil(o.FeatureValues) {
		var ret []GcpMarketplaceProductFeatureValue
		return ret
	}
	return o.FeatureValues
}

// GetFeatureValuesOk returns a tuple with the FeatureValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetFeatureValuesOk() ([]GcpMarketplaceProductFeatureValue, bool) {
	if o == nil || IsNil(o.FeatureValues) {
		return nil, false
	}
	return o.FeatureValues, true
}

// HasFeatureValues returns a boolean if a field has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) HasFeatureValues() bool {
	if o != nil && !IsNil(o.FeatureValues) {
		return true
	}

	return false
}

// SetFeatureValues gets a reference to the given []GcpMarketplaceProductFeatureValue and assigns it to the FeatureValues field.
func (o *GcpMarketplaceProductPurchaseOptionSpec) SetFeatureValues(v []GcpMarketplaceProductFeatureValue) {
	o.FeatureValues = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GcpMarketplaceProductPurchaseOptionSpec) SetName(v string) {
	o.Name = &v
}

// GetPriceInfo returns the PriceInfo field value if set, zero value otherwise.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetPriceInfo() GcpMarketplaceProductPriceInfo {
	if o == nil || IsNil(o.PriceInfo) {
		var ret GcpMarketplaceProductPriceInfo
		return ret
	}
	return *o.PriceInfo
}

// GetPriceInfoOk returns a tuple with the PriceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetPriceInfoOk() (*GcpMarketplaceProductPriceInfo, bool) {
	if o == nil || IsNil(o.PriceInfo) {
		return nil, false
	}
	return o.PriceInfo, true
}

// HasPriceInfo returns a boolean if a field has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) HasPriceInfo() bool {
	if o != nil && !IsNil(o.PriceInfo) {
		return true
	}

	return false
}

// SetPriceInfo gets a reference to the given GcpMarketplaceProductPriceInfo and assigns it to the PriceInfo field.
func (o *GcpMarketplaceProductPurchaseOptionSpec) SetPriceInfo(v GcpMarketplaceProductPriceInfo) {
	o.PriceInfo = &v
}

// GetPurchaseMode returns the PurchaseMode field value if set, zero value otherwise.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetPurchaseMode() string {
	if o == nil || IsNil(o.PurchaseMode) {
		var ret string
		return ret
	}
	return *o.PurchaseMode
}

// GetPurchaseModeOk returns a tuple with the PurchaseMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetPurchaseModeOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseMode) {
		return nil, false
	}
	return o.PurchaseMode, true
}

// HasPurchaseMode returns a boolean if a field has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) HasPurchaseMode() bool {
	if o != nil && !IsNil(o.PurchaseMode) {
		return true
	}

	return false
}

// SetPurchaseMode gets a reference to the given string and assigns it to the PurchaseMode field.
func (o *GcpMarketplaceProductPurchaseOptionSpec) SetPurchaseMode(v string) {
	o.PurchaseMode = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GcpMarketplaceProductPurchaseOptionSpec) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GcpMarketplaceProductPurchaseOptionSpec) SetTitle(v string) {
	o.Title = &v
}

func (o GcpMarketplaceProductPurchaseOptionSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductPurchaseOptionSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureValues) {
		toSerialize["featureValues"] = o.FeatureValues
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PriceInfo) {
		toSerialize["priceInfo"] = o.PriceInfo
	}
	if !IsNil(o.PurchaseMode) {
		toSerialize["purchaseMode"] = o.PurchaseMode
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductPurchaseOptionSpec struct {
	value *GcpMarketplaceProductPurchaseOptionSpec
	isSet bool
}

func (v NullableGcpMarketplaceProductPurchaseOptionSpec) Get() *GcpMarketplaceProductPurchaseOptionSpec {
	return v.value
}

func (v *NullableGcpMarketplaceProductPurchaseOptionSpec) Set(val *GcpMarketplaceProductPurchaseOptionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductPurchaseOptionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductPurchaseOptionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductPurchaseOptionSpec(val *GcpMarketplaceProductPurchaseOptionSpec) *NullableGcpMarketplaceProductPurchaseOptionSpec {
	return &NullableGcpMarketplaceProductPurchaseOptionSpec{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductPurchaseOptionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductPurchaseOptionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
