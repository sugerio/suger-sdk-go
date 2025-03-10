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

// checks if the AzureMarketplacePriceAndAvailabilityRecurrentPriceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilityRecurrentPriceItem{}

// AzureMarketplacePriceAndAvailabilityRecurrentPriceItem struct for AzureMarketplacePriceAndAvailabilityRecurrentPriceItem
type AzureMarketplacePriceAndAvailabilityRecurrentPriceItem struct {
	BillingTerm          *AzureMarketplaceTerm   `json:"billingTerm,omitempty"`
	PaymentOption        *AzureMarketplaceTerm   `json:"paymentOption,omitempty"`
	PricePerPaymentInUsd *float32                `json:"pricePerPaymentInUsd,omitempty"`
	Prices               []AzureMarketplacePrice `json:"prices,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItem instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPriceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItem() *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem {
	this := AzureMarketplacePriceAndAvailabilityRecurrentPriceItem{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItemWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPriceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilityRecurrentPriceItemWithDefaults() *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem {
	this := AzureMarketplacePriceAndAvailabilityRecurrentPriceItem{}
	return &this
}

// GetBillingTerm returns the BillingTerm field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetBillingTerm() AzureMarketplaceTerm {
	if o == nil || IsNil(o.BillingTerm) {
		var ret AzureMarketplaceTerm
		return ret
	}
	return *o.BillingTerm
}

// GetBillingTermOk returns a tuple with the BillingTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetBillingTermOk() (*AzureMarketplaceTerm, bool) {
	if o == nil || IsNil(o.BillingTerm) {
		return nil, false
	}
	return o.BillingTerm, true
}

// HasBillingTerm returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasBillingTerm() bool {
	if o != nil && !IsNil(o.BillingTerm) {
		return true
	}

	return false
}

// SetBillingTerm gets a reference to the given AzureMarketplaceTerm and assigns it to the BillingTerm field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetBillingTerm(v AzureMarketplaceTerm) {
	o.BillingTerm = &v
}

// GetPaymentOption returns the PaymentOption field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPaymentOption() AzureMarketplaceTerm {
	if o == nil || IsNil(o.PaymentOption) {
		var ret AzureMarketplaceTerm
		return ret
	}
	return *o.PaymentOption
}

// GetPaymentOptionOk returns a tuple with the PaymentOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPaymentOptionOk() (*AzureMarketplaceTerm, bool) {
	if o == nil || IsNil(o.PaymentOption) {
		return nil, false
	}
	return o.PaymentOption, true
}

// HasPaymentOption returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasPaymentOption() bool {
	if o != nil && !IsNil(o.PaymentOption) {
		return true
	}

	return false
}

// SetPaymentOption gets a reference to the given AzureMarketplaceTerm and assigns it to the PaymentOption field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetPaymentOption(v AzureMarketplaceTerm) {
	o.PaymentOption = &v
}

// GetPricePerPaymentInUsd returns the PricePerPaymentInUsd field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPricePerPaymentInUsd() float32 {
	if o == nil || IsNil(o.PricePerPaymentInUsd) {
		var ret float32
		return ret
	}
	return *o.PricePerPaymentInUsd
}

// GetPricePerPaymentInUsdOk returns a tuple with the PricePerPaymentInUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPricePerPaymentInUsdOk() (*float32, bool) {
	if o == nil || IsNil(o.PricePerPaymentInUsd) {
		return nil, false
	}
	return o.PricePerPaymentInUsd, true
}

// HasPricePerPaymentInUsd returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasPricePerPaymentInUsd() bool {
	if o != nil && !IsNil(o.PricePerPaymentInUsd) {
		return true
	}

	return false
}

// SetPricePerPaymentInUsd gets a reference to the given float32 and assigns it to the PricePerPaymentInUsd field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetPricePerPaymentInUsd(v float32) {
	o.PricePerPaymentInUsd = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPrices() []AzureMarketplacePrice {
	if o == nil || IsNil(o.Prices) {
		var ret []AzureMarketplacePrice
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) GetPricesOk() ([]AzureMarketplacePrice, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []AzureMarketplacePrice and assigns it to the Prices field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) SetPrices(v []AzureMarketplacePrice) {
	o.Prices = v
}

func (o AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingTerm) {
		toSerialize["billingTerm"] = o.BillingTerm
	}
	if !IsNil(o.PaymentOption) {
		toSerialize["paymentOption"] = o.PaymentOption
	}
	if !IsNil(o.PricePerPaymentInUsd) {
		toSerialize["pricePerPaymentInUsd"] = o.PricePerPaymentInUsd
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem struct {
	value *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem) Get() *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem) Set(val *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem(val *AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) *NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem {
	return &NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilityRecurrentPriceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
