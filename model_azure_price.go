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

// checks if the AzurePrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzurePrice{}

// AzurePrice struct for AzurePrice
type AzurePrice struct {
	// ISO currency code, Three characters
	CurrencyCode *string `json:"currencyCode,omitempty"`
	OpenPrice *float32 `json:"openPrice,omitempty"`
	PriceTierID *string `json:"priceTierID,omitempty"`
}

// NewAzurePrice instantiates a new AzurePrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePrice() *AzurePrice {
	this := AzurePrice{}
	return &this
}

// NewAzurePriceWithDefaults instantiates a new AzurePrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePriceWithDefaults() *AzurePrice {
	this := AzurePrice{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AzurePrice) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrice) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AzurePrice) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AzurePrice) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetOpenPrice returns the OpenPrice field value if set, zero value otherwise.
func (o *AzurePrice) GetOpenPrice() float32 {
	if o == nil || IsNil(o.OpenPrice) {
		var ret float32
		return ret
	}
	return *o.OpenPrice
}

// GetOpenPriceOk returns a tuple with the OpenPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrice) GetOpenPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.OpenPrice) {
		return nil, false
	}
	return o.OpenPrice, true
}

// HasOpenPrice returns a boolean if a field has been set.
func (o *AzurePrice) HasOpenPrice() bool {
	if o != nil && !IsNil(o.OpenPrice) {
		return true
	}

	return false
}

// SetOpenPrice gets a reference to the given float32 and assigns it to the OpenPrice field.
func (o *AzurePrice) SetOpenPrice(v float32) {
	o.OpenPrice = &v
}

// GetPriceTierID returns the PriceTierID field value if set, zero value otherwise.
func (o *AzurePrice) GetPriceTierID() string {
	if o == nil || IsNil(o.PriceTierID) {
		var ret string
		return ret
	}
	return *o.PriceTierID
}

// GetPriceTierIDOk returns a tuple with the PriceTierID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePrice) GetPriceTierIDOk() (*string, bool) {
	if o == nil || IsNil(o.PriceTierID) {
		return nil, false
	}
	return o.PriceTierID, true
}

// HasPriceTierID returns a boolean if a field has been set.
func (o *AzurePrice) HasPriceTierID() bool {
	if o != nil && !IsNil(o.PriceTierID) {
		return true
	}

	return false
}

// SetPriceTierID gets a reference to the given string and assigns it to the PriceTierID field.
func (o *AzurePrice) SetPriceTierID(v string) {
	o.PriceTierID = &v
}

func (o AzurePrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzurePrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.OpenPrice) {
		toSerialize["openPrice"] = o.OpenPrice
	}
	if !IsNil(o.PriceTierID) {
		toSerialize["priceTierID"] = o.PriceTierID
	}
	return toSerialize, nil
}

type NullableAzurePrice struct {
	value *AzurePrice
	isSet bool
}

func (v NullableAzurePrice) Get() *AzurePrice {
	return v.value
}

func (v *NullableAzurePrice) Set(val *AzurePrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePrice(val *AzurePrice) *NullableAzurePrice {
	return &NullableAzurePrice{value: val, isSet: true}
}

func (v NullableAzurePrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


