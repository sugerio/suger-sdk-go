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

// checks if the AzureMarketplacePriceAndAvailabilityRecurrentPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilityRecurrentPrice{}

// AzureMarketplacePriceAndAvailabilityRecurrentPrice struct for AzureMarketplacePriceAndAvailabilityRecurrentPrice
type AzureMarketplacePriceAndAvailabilityRecurrentPrice struct {
	// default \"usd\"
	PriceInputOption *string                                                  `json:"priceInputOption,omitempty"`
	Prices           []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem `json:"prices,omitempty"`
	// default \"flatRate\"
	RecurrentPriceMode *string                                                      `json:"recurrentPriceMode,omitempty"`
	UserLimits         *AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit `json:"userLimits,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilityRecurrentPrice instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilityRecurrentPrice() *AzureMarketplacePriceAndAvailabilityRecurrentPrice {
	this := AzureMarketplacePriceAndAvailabilityRecurrentPrice{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilityRecurrentPriceWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityRecurrentPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilityRecurrentPriceWithDefaults() *AzureMarketplacePriceAndAvailabilityRecurrentPrice {
	this := AzureMarketplacePriceAndAvailabilityRecurrentPrice{}
	return &this
}

// GetPriceInputOption returns the PriceInputOption field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPriceInputOption() string {
	if o == nil || IsNil(o.PriceInputOption) {
		var ret string
		return ret
	}
	return *o.PriceInputOption
}

// GetPriceInputOptionOk returns a tuple with the PriceInputOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPriceInputOptionOk() (*string, bool) {
	if o == nil || IsNil(o.PriceInputOption) {
		return nil, false
	}
	return o.PriceInputOption, true
}

// HasPriceInputOption returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasPriceInputOption() bool {
	if o != nil && !IsNil(o.PriceInputOption) {
		return true
	}

	return false
}

// SetPriceInputOption gets a reference to the given string and assigns it to the PriceInputOption field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetPriceInputOption(v string) {
	o.PriceInputOption = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPrices() []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem {
	if o == nil || IsNil(o.Prices) {
		var ret []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetPricesOk() ([]AzureMarketplacePriceAndAvailabilityRecurrentPriceItem, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem and assigns it to the Prices field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetPrices(v []AzureMarketplacePriceAndAvailabilityRecurrentPriceItem) {
	o.Prices = v
}

// GetRecurrentPriceMode returns the RecurrentPriceMode field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetRecurrentPriceMode() string {
	if o == nil || IsNil(o.RecurrentPriceMode) {
		var ret string
		return ret
	}
	return *o.RecurrentPriceMode
}

// GetRecurrentPriceModeOk returns a tuple with the RecurrentPriceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetRecurrentPriceModeOk() (*string, bool) {
	if o == nil || IsNil(o.RecurrentPriceMode) {
		return nil, false
	}
	return o.RecurrentPriceMode, true
}

// HasRecurrentPriceMode returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasRecurrentPriceMode() bool {
	if o != nil && !IsNil(o.RecurrentPriceMode) {
		return true
	}

	return false
}

// SetRecurrentPriceMode gets a reference to the given string and assigns it to the RecurrentPriceMode field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetRecurrentPriceMode(v string) {
	o.RecurrentPriceMode = &v
}

// GetUserLimits returns the UserLimits field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetUserLimits() AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit {
	if o == nil || IsNil(o.UserLimits) {
		var ret AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit
		return ret
	}
	return *o.UserLimits
}

// GetUserLimitsOk returns a tuple with the UserLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) GetUserLimitsOk() (*AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit, bool) {
	if o == nil || IsNil(o.UserLimits) {
		return nil, false
	}
	return o.UserLimits, true
}

// HasUserLimits returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) HasUserLimits() bool {
	if o != nil && !IsNil(o.UserLimits) {
		return true
	}

	return false
}

// SetUserLimits gets a reference to the given AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit and assigns it to the UserLimits field.
func (o *AzureMarketplacePriceAndAvailabilityRecurrentPrice) SetUserLimits(v AzureMarketplacePriceAndAvailabilityRecurrentPriceUserLimit) {
	o.UserLimits = &v
}

func (o AzureMarketplacePriceAndAvailabilityRecurrentPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilityRecurrentPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceInputOption) {
		toSerialize["priceInputOption"] = o.PriceInputOption
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.RecurrentPriceMode) {
		toSerialize["recurrentPriceMode"] = o.RecurrentPriceMode
	}
	if !IsNil(o.UserLimits) {
		toSerialize["userLimits"] = o.UserLimits
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice struct {
	value *AzureMarketplacePriceAndAvailabilityRecurrentPrice
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice) Get() *AzureMarketplacePriceAndAvailabilityRecurrentPrice {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice) Set(val *AzureMarketplacePriceAndAvailabilityRecurrentPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilityRecurrentPrice(val *AzureMarketplacePriceAndAvailabilityRecurrentPrice) *NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice {
	return &NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilityRecurrentPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}