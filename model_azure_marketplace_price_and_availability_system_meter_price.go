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

// checks if the AzureMarketplacePriceAndAvailabilitySystemMeterPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilitySystemMeterPrice{}

// AzureMarketplacePriceAndAvailabilitySystemMeterPrice struct for AzureMarketplacePriceAndAvailabilitySystemMeterPrice
type AzureMarketplacePriceAndAvailabilitySystemMeterPrice struct {
	// default 0
	Price            *float32                `json:"price,omitempty"`
	PriceInputOption *string                 `json:"priceInputOption,omitempty"`
	Prices           []AzureMarketplacePrice `json:"prices,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilitySystemMeterPrice instantiates a new AzureMarketplacePriceAndAvailabilitySystemMeterPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilitySystemMeterPrice() *AzureMarketplacePriceAndAvailabilitySystemMeterPrice {
	this := AzureMarketplacePriceAndAvailabilitySystemMeterPrice{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilitySystemMeterPriceWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilitySystemMeterPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilitySystemMeterPriceWithDefaults() *AzureMarketplacePriceAndAvailabilitySystemMeterPrice {
	this := AzureMarketplacePriceAndAvailabilitySystemMeterPrice{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) SetPrice(v float32) {
	o.Price = &v
}

// GetPriceInputOption returns the PriceInputOption field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) GetPriceInputOption() string {
	if o == nil || IsNil(o.PriceInputOption) {
		var ret string
		return ret
	}
	return *o.PriceInputOption
}

// GetPriceInputOptionOk returns a tuple with the PriceInputOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) GetPriceInputOptionOk() (*string, bool) {
	if o == nil || IsNil(o.PriceInputOption) {
		return nil, false
	}
	return o.PriceInputOption, true
}

// HasPriceInputOption returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) HasPriceInputOption() bool {
	if o != nil && !IsNil(o.PriceInputOption) {
		return true
	}

	return false
}

// SetPriceInputOption gets a reference to the given string and assigns it to the PriceInputOption field.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) SetPriceInputOption(v string) {
	o.PriceInputOption = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) GetPrices() []AzureMarketplacePrice {
	if o == nil || IsNil(o.Prices) {
		var ret []AzureMarketplacePrice
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) GetPricesOk() ([]AzureMarketplacePrice, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []AzureMarketplacePrice and assigns it to the Prices field.
func (o *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) SetPrices(v []AzureMarketplacePrice) {
	o.Prices = v
}

func (o AzureMarketplacePriceAndAvailabilitySystemMeterPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilitySystemMeterPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceInputOption) {
		toSerialize["priceInputOption"] = o.PriceInputOption
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice struct {
	value *AzureMarketplacePriceAndAvailabilitySystemMeterPrice
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice) Get() *AzureMarketplacePriceAndAvailabilitySystemMeterPrice {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice) Set(val *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice(val *AzureMarketplacePriceAndAvailabilitySystemMeterPrice) *NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice {
	return &NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilitySystemMeterPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
