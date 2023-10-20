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

// checks if the AzurePriceAndAvailabilityPrivateOfferCustomMeters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzurePriceAndAvailabilityPrivateOfferCustomMeters{}

// AzurePriceAndAvailabilityPrivateOfferCustomMeters struct for AzurePriceAndAvailabilityPrivateOfferCustomMeters
type AzurePriceAndAvailabilityPrivateOfferCustomMeters struct {
	// One of PriceAndAvailabilityCustomMeter_USD or PriceAndAvailabilityCustomMeter_PerMarket
	Meters map[string]interface{} `json:"meters,omitempty"`
	// default \"usd\"
	PriceInputOption *string `json:"priceInputOption,omitempty"`
}

// NewAzurePriceAndAvailabilityPrivateOfferCustomMeters instantiates a new AzurePriceAndAvailabilityPrivateOfferCustomMeters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzurePriceAndAvailabilityPrivateOfferCustomMeters() *AzurePriceAndAvailabilityPrivateOfferCustomMeters {
	this := AzurePriceAndAvailabilityPrivateOfferCustomMeters{}
	return &this
}

// NewAzurePriceAndAvailabilityPrivateOfferCustomMetersWithDefaults instantiates a new AzurePriceAndAvailabilityPrivateOfferCustomMeters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzurePriceAndAvailabilityPrivateOfferCustomMetersWithDefaults() *AzurePriceAndAvailabilityPrivateOfferCustomMeters {
	this := AzurePriceAndAvailabilityPrivateOfferCustomMeters{}
	return &this
}

// GetMeters returns the Meters field value if set, zero value otherwise.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) GetMeters() map[string]interface{} {
	if o == nil || IsNil(o.Meters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Meters
}

// GetMetersOk returns a tuple with the Meters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) GetMetersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Meters) {
		return map[string]interface{}{}, false
	}
	return o.Meters, true
}

// HasMeters returns a boolean if a field has been set.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) HasMeters() bool {
	if o != nil && !IsNil(o.Meters) {
		return true
	}

	return false
}

// SetMeters gets a reference to the given map[string]interface{} and assigns it to the Meters field.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) SetMeters(v map[string]interface{}) {
	o.Meters = v
}

// GetPriceInputOption returns the PriceInputOption field value if set, zero value otherwise.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) GetPriceInputOption() string {
	if o == nil || IsNil(o.PriceInputOption) {
		var ret string
		return ret
	}
	return *o.PriceInputOption
}

// GetPriceInputOptionOk returns a tuple with the PriceInputOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) GetPriceInputOptionOk() (*string, bool) {
	if o == nil || IsNil(o.PriceInputOption) {
		return nil, false
	}
	return o.PriceInputOption, true
}

// HasPriceInputOption returns a boolean if a field has been set.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) HasPriceInputOption() bool {
	if o != nil && !IsNil(o.PriceInputOption) {
		return true
	}

	return false
}

// SetPriceInputOption gets a reference to the given string and assigns it to the PriceInputOption field.
func (o *AzurePriceAndAvailabilityPrivateOfferCustomMeters) SetPriceInputOption(v string) {
	o.PriceInputOption = &v
}

func (o AzurePriceAndAvailabilityPrivateOfferCustomMeters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzurePriceAndAvailabilityPrivateOfferCustomMeters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meters) {
		toSerialize["meters"] = o.Meters
	}
	if !IsNil(o.PriceInputOption) {
		toSerialize["priceInputOption"] = o.PriceInputOption
	}
	return toSerialize, nil
}

type NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters struct {
	value *AzurePriceAndAvailabilityPrivateOfferCustomMeters
	isSet bool
}

func (v NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters) Get() *AzurePriceAndAvailabilityPrivateOfferCustomMeters {
	return v.value
}

func (v *NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters) Set(val *AzurePriceAndAvailabilityPrivateOfferCustomMeters) {
	v.value = val
	v.isSet = true
}

func (v NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters) IsSet() bool {
	return v.isSet
}

func (v *NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzurePriceAndAvailabilityPrivateOfferCustomMeters(val *AzurePriceAndAvailabilityPrivateOfferCustomMeters) *NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters {
	return &NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters{value: val, isSet: true}
}

func (v NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzurePriceAndAvailabilityPrivateOfferCustomMeters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


