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

// checks if the AzureMarketplacePriceAndAvailabilityPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePriceAndAvailabilityPrice{}

// AzureMarketplacePriceAndAvailabilityPrice struct for AzureMarketplacePriceAndAvailabilityPrice
type AzureMarketplacePriceAndAvailabilityPrice struct {
	CorePricing        *AzureMarketplacePriceAndAvailabilityCorePrice        `json:"corePricing,omitempty"`
	CustomMeters       *AzureMarketplacePriceAndAvailabilityCustomMeterPrice `json:"customMeters,omitempty"`
	LicenseModel       *string                                               `json:"licenseModel,omitempty"`
	RecurrentPrice     *AzureMarketplacePriceAndAvailabilityRecurrentPrice   `json:"recurrentPrice,omitempty"`
	SystemMeterPricing *AzureMarketplacePriceAndAvailabilitySystemMeterPrice `json:"systemMeterPricing,omitempty"`
}

// NewAzureMarketplacePriceAndAvailabilityPrice instantiates a new AzureMarketplacePriceAndAvailabilityPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePriceAndAvailabilityPrice() *AzureMarketplacePriceAndAvailabilityPrice {
	this := AzureMarketplacePriceAndAvailabilityPrice{}
	return &this
}

// NewAzureMarketplacePriceAndAvailabilityPriceWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePriceAndAvailabilityPriceWithDefaults() *AzureMarketplacePriceAndAvailabilityPrice {
	this := AzureMarketplacePriceAndAvailabilityPrice{}
	return &this
}

// GetCorePricing returns the CorePricing field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCorePricing() AzureMarketplacePriceAndAvailabilityCorePrice {
	if o == nil || IsNil(o.CorePricing) {
		var ret AzureMarketplacePriceAndAvailabilityCorePrice
		return ret
	}
	return *o.CorePricing
}

// GetCorePricingOk returns a tuple with the CorePricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCorePricingOk() (*AzureMarketplacePriceAndAvailabilityCorePrice, bool) {
	if o == nil || IsNil(o.CorePricing) {
		return nil, false
	}
	return o.CorePricing, true
}

// HasCorePricing returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) HasCorePricing() bool {
	if o != nil && !IsNil(o.CorePricing) {
		return true
	}

	return false
}

// SetCorePricing gets a reference to the given AzureMarketplacePriceAndAvailabilityCorePrice and assigns it to the CorePricing field.
func (o *AzureMarketplacePriceAndAvailabilityPrice) SetCorePricing(v AzureMarketplacePriceAndAvailabilityCorePrice) {
	o.CorePricing = &v
}

// GetCustomMeters returns the CustomMeters field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCustomMeters() AzureMarketplacePriceAndAvailabilityCustomMeterPrice {
	if o == nil || IsNil(o.CustomMeters) {
		var ret AzureMarketplacePriceAndAvailabilityCustomMeterPrice
		return ret
	}
	return *o.CustomMeters
}

// GetCustomMetersOk returns a tuple with the CustomMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCustomMetersOk() (*AzureMarketplacePriceAndAvailabilityCustomMeterPrice, bool) {
	if o == nil || IsNil(o.CustomMeters) {
		return nil, false
	}
	return o.CustomMeters, true
}

// HasCustomMeters returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) HasCustomMeters() bool {
	if o != nil && !IsNil(o.CustomMeters) {
		return true
	}

	return false
}

// SetCustomMeters gets a reference to the given AzureMarketplacePriceAndAvailabilityCustomMeterPrice and assigns it to the CustomMeters field.
func (o *AzureMarketplacePriceAndAvailabilityPrice) SetCustomMeters(v AzureMarketplacePriceAndAvailabilityCustomMeterPrice) {
	o.CustomMeters = &v
}

// GetLicenseModel returns the LicenseModel field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetLicenseModel() string {
	if o == nil || IsNil(o.LicenseModel) {
		var ret string
		return ret
	}
	return *o.LicenseModel
}

// GetLicenseModelOk returns a tuple with the LicenseModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetLicenseModelOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseModel) {
		return nil, false
	}
	return o.LicenseModel, true
}

// HasLicenseModel returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) HasLicenseModel() bool {
	if o != nil && !IsNil(o.LicenseModel) {
		return true
	}

	return false
}

// SetLicenseModel gets a reference to the given string and assigns it to the LicenseModel field.
func (o *AzureMarketplacePriceAndAvailabilityPrice) SetLicenseModel(v string) {
	o.LicenseModel = &v
}

// GetRecurrentPrice returns the RecurrentPrice field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetRecurrentPrice() AzureMarketplacePriceAndAvailabilityRecurrentPrice {
	if o == nil || IsNil(o.RecurrentPrice) {
		var ret AzureMarketplacePriceAndAvailabilityRecurrentPrice
		return ret
	}
	return *o.RecurrentPrice
}

// GetRecurrentPriceOk returns a tuple with the RecurrentPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetRecurrentPriceOk() (*AzureMarketplacePriceAndAvailabilityRecurrentPrice, bool) {
	if o == nil || IsNil(o.RecurrentPrice) {
		return nil, false
	}
	return o.RecurrentPrice, true
}

// HasRecurrentPrice returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) HasRecurrentPrice() bool {
	if o != nil && !IsNil(o.RecurrentPrice) {
		return true
	}

	return false
}

// SetRecurrentPrice gets a reference to the given AzureMarketplacePriceAndAvailabilityRecurrentPrice and assigns it to the RecurrentPrice field.
func (o *AzureMarketplacePriceAndAvailabilityPrice) SetRecurrentPrice(v AzureMarketplacePriceAndAvailabilityRecurrentPrice) {
	o.RecurrentPrice = &v
}

// GetSystemMeterPricing returns the SystemMeterPricing field value if set, zero value otherwise.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetSystemMeterPricing() AzureMarketplacePriceAndAvailabilitySystemMeterPrice {
	if o == nil || IsNil(o.SystemMeterPricing) {
		var ret AzureMarketplacePriceAndAvailabilitySystemMeterPrice
		return ret
	}
	return *o.SystemMeterPricing
}

// GetSystemMeterPricingOk returns a tuple with the SystemMeterPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) GetSystemMeterPricingOk() (*AzureMarketplacePriceAndAvailabilitySystemMeterPrice, bool) {
	if o == nil || IsNil(o.SystemMeterPricing) {
		return nil, false
	}
	return o.SystemMeterPricing, true
}

// HasSystemMeterPricing returns a boolean if a field has been set.
func (o *AzureMarketplacePriceAndAvailabilityPrice) HasSystemMeterPricing() bool {
	if o != nil && !IsNil(o.SystemMeterPricing) {
		return true
	}

	return false
}

// SetSystemMeterPricing gets a reference to the given AzureMarketplacePriceAndAvailabilitySystemMeterPrice and assigns it to the SystemMeterPricing field.
func (o *AzureMarketplacePriceAndAvailabilityPrice) SetSystemMeterPricing(v AzureMarketplacePriceAndAvailabilitySystemMeterPrice) {
	o.SystemMeterPricing = &v
}

func (o AzureMarketplacePriceAndAvailabilityPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePriceAndAvailabilityPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorePricing) {
		toSerialize["corePricing"] = o.CorePricing
	}
	if !IsNil(o.CustomMeters) {
		toSerialize["customMeters"] = o.CustomMeters
	}
	if !IsNil(o.LicenseModel) {
		toSerialize["licenseModel"] = o.LicenseModel
	}
	if !IsNil(o.RecurrentPrice) {
		toSerialize["recurrentPrice"] = o.RecurrentPrice
	}
	if !IsNil(o.SystemMeterPricing) {
		toSerialize["systemMeterPricing"] = o.SystemMeterPricing
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePriceAndAvailabilityPrice struct {
	value *AzureMarketplacePriceAndAvailabilityPrice
	isSet bool
}

func (v NullableAzureMarketplacePriceAndAvailabilityPrice) Get() *AzureMarketplacePriceAndAvailabilityPrice {
	return v.value
}

func (v *NullableAzureMarketplacePriceAndAvailabilityPrice) Set(val *AzureMarketplacePriceAndAvailabilityPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePriceAndAvailabilityPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePriceAndAvailabilityPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePriceAndAvailabilityPrice(val *AzureMarketplacePriceAndAvailabilityPrice) *NullableAzureMarketplacePriceAndAvailabilityPrice {
	return &NullableAzureMarketplacePriceAndAvailabilityPrice{value: val, isSet: true}
}

func (v NullableAzureMarketplacePriceAndAvailabilityPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePriceAndAvailabilityPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
