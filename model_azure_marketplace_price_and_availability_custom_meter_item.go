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

// checks if the AzureMarketplacePriceAndAvailabilityCustomMeterItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilityCustomMeterItem{}

// AzureMarketplacePriceAndAvailabilityCustomMeterItem struct for AzureMarketplacePriceAndAvailabilityCustomMeterItem
type AzureMarketplacePriceAndAvailabilityCustomMeterItem struct {
	DisplayName *string `json:"displayName,omitempty"`
	// Suger's custom field, for Suger internal use only. Not from Microsoft official schema.
	Price         *float32 `json:"price,omitempty"`
	UnitOfMeasure *string  `json:"unitOfMeasure,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilityCustomMeterItem instantiates a new AzureMarketplacePriceAndAvailabilityCustomMeterItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilityCustomMeterItem() *AzureMarketplacePriceAndAvailabilityCustomMeterItem {
	this := AzureMarketplacePriceAndAvailabilityCustomMeterItem{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilityCustomMeterItemWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityCustomMeterItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilityCustomMeterItemWithDefaults() *AzureMarketplacePriceAndAvailabilityCustomMeterItem {
	this := AzureMarketplacePriceAndAvailabilityCustomMeterItem{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) SetPrice(v float32) {
	o.Price = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *AzureMarketplacePriceAndAvailabilityCustomMeterItem) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

func (o AzureMarketplacePriceAndAvailabilityCustomMeterItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilityCustomMeterItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem struct {
	value *AzureMarketplacePriceAndAvailabilityCustomMeterItem
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem) Get() *AzureMarketplacePriceAndAvailabilityCustomMeterItem {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem) Set(val *AzureMarketplacePriceAndAvailabilityCustomMeterItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilityCustomMeterItem(val *AzureMarketplacePriceAndAvailabilityCustomMeterItem) *NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem {
	return &NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilityCustomMeterItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}