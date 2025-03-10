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

// checks if the AzureMarketplacePriceAndAvailabilityPrivateOfferPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilityPrivateOfferPrice{}

// AzureMarketplacePriceAndAvailabilityPrivateOfferPrice struct for AzureMarketplacePriceAndAvailabilityPrivateOfferPrice
type AzureMarketplacePriceAndAvailabilityPrivateOfferPrice struct {
	CustomMeters   *AzureMarketplacePriceAndAvailabilityPrivateOfferCustomMeters `json:"customMeters,omitempty"`
	RecurrentPrice *AzureMarketplacePriceAndAvailabilityRecurrentPrice           `json:"recurrentPrice,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilityPrivateOfferPrice instantiates a new AzureMarketplacePriceAndAvailabilityPrivateOfferPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilityPrivateOfferPrice() *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice {
	this := AzureMarketplacePriceAndAvailabilityPrivateOfferPrice{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilityPrivateOfferPriceWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityPrivateOfferPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilityPrivateOfferPriceWithDefaults() *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice {
	this := AzureMarketplacePriceAndAvailabilityPrivateOfferPrice{}
	return &this
}

// GetCustomMeters returns the CustomMeters field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) GetCustomMeters() AzureMarketplacePriceAndAvailabilityPrivateOfferCustomMeters {
	if o == nil || IsNil(o.CustomMeters) {
		var ret AzureMarketplacePriceAndAvailabilityPrivateOfferCustomMeters
		return ret
	}
	return *o.CustomMeters
}

// GetCustomMetersOk returns a tuple with the CustomMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) GetCustomMetersOk() (*AzureMarketplacePriceAndAvailabilityPrivateOfferCustomMeters, bool) {
	if o == nil || IsNil(o.CustomMeters) {
		return nil, false
	}
	return o.CustomMeters, true
}

// HasCustomMeters returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) HasCustomMeters() bool {
	if o != nil && !IsNil(o.CustomMeters) {
		return true
	}

	return false
}

// SetCustomMeters gets a reference to the given AzureMarketplacePriceAndAvailabilityPrivateOfferCustomMeters and assigns it to the CustomMeters field.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) SetCustomMeters(v AzureMarketplacePriceAndAvailabilityPrivateOfferCustomMeters) {
	o.CustomMeters = &v
}

// GetRecurrentPrice returns the RecurrentPrice field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) GetRecurrentPrice() AzureMarketplacePriceAndAvailabilityRecurrentPrice {
	if o == nil || IsNil(o.RecurrentPrice) {
		var ret AzureMarketplacePriceAndAvailabilityRecurrentPrice
		return ret
	}
	return *o.RecurrentPrice
}

// GetRecurrentPriceOk returns a tuple with the RecurrentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) GetRecurrentPriceOk() (*AzureMarketplacePriceAndAvailabilityRecurrentPrice, bool) {
	if o == nil || IsNil(o.RecurrentPrice) {
		return nil, false
	}
	return o.RecurrentPrice, true
}

// HasRecurrentPrice returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) HasRecurrentPrice() bool {
	if o != nil && !IsNil(o.RecurrentPrice) {
		return true
	}

	return false
}

// SetRecurrentPrice gets a reference to the given AzureMarketplacePriceAndAvailabilityRecurrentPrice and assigns it to the RecurrentPrice field.
func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) SetRecurrentPrice(v AzureMarketplacePriceAndAvailabilityRecurrentPrice) {
	o.RecurrentPrice = &v
}

func (o AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomMeters) {
		toSerialize["customMeters"] = o.CustomMeters
	}
	if !IsNil(o.RecurrentPrice) {
		toSerialize["recurrentPrice"] = o.RecurrentPrice
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice struct {
	value *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice) Get() *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice) Set(val *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice(val *AzureMarketplacePriceAndAvailabilityPrivateOfferPrice) *NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice {
	return &NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilityPrivateOfferPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
