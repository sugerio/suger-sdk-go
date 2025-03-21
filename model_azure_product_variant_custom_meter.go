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

// checks if the AzureProductVariantCustomMeter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureProductVariantCustomMeter{}

// AzureProductVariantCustomMeter struct for AzureProductVariantCustomMeter
type AzureProductVariantCustomMeter struct {
	DisplayName            *string                     `json:"displayName,omitempty"`
	Id                     *string                     `json:"id,omitempty"`
	IncludedBaseQuantities []AzureIncludedBaseQuantity `json:"includedBaseQuantities,omitempty"`
	IsEnabled              *bool                       `json:"isEnabled,omitempty"`
	PriceInUsd             *float32                    `json:"priceInUsd,omitempty"`
	UniqueID               *string                     `json:"uniqueID,omitempty"`
	UnitOfMeasure          *string                     `json:"unitOfMeasure,omitempty"`
}

// NewAzureProductVariantCustomMeter instantiates a new AzureProductVariantCustomMeter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProductVariantCustomMeter() *AzureProductVariantCustomMeter {
	this := AzureProductVariantCustomMeter{}
	return &this
}

// NewAzureProductVariantCustomMeterWithDefaults instantiates a new AzureProductVariantCustomMeter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProductVariantCustomMeterWithDefaults() *AzureProductVariantCustomMeter {
	this := AzureProductVariantCustomMeter{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AzureProductVariantCustomMeter) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureProductVariantCustomMeter) SetId(v string) {
	o.Id = &v
}

// GetIncludedBaseQuantities returns the IncludedBaseQuantities field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetIncludedBaseQuantities() []AzureIncludedBaseQuantity {
	if o == nil || IsNil(o.IncludedBaseQuantities) {
		var ret []AzureIncludedBaseQuantity
		return ret
	}
	return o.IncludedBaseQuantities
}

// GetIncludedBaseQuantitiesOk returns a tuple with the IncludedBaseQuantities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetIncludedBaseQuantitiesOk() ([]AzureIncludedBaseQuantity, bool) {
	if o == nil || IsNil(o.IncludedBaseQuantities) {
		return nil, false
	}
	return o.IncludedBaseQuantities, true
}

// HasIncludedBaseQuantities returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasIncludedBaseQuantities() bool {
	if o != nil && !IsNil(o.IncludedBaseQuantities) {
		return true
	}

	return false
}

// SetIncludedBaseQuantities gets a reference to the given []AzureIncludedBaseQuantity and assigns it to the IncludedBaseQuantities field.
func (o *AzureProductVariantCustomMeter) SetIncludedBaseQuantities(v []AzureIncludedBaseQuantity) {
	o.IncludedBaseQuantities = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *AzureProductVariantCustomMeter) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetPriceInUsd returns the PriceInUsd field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetPriceInUsd() float32 {
	if o == nil || IsNil(o.PriceInUsd) {
		var ret float32
		return ret
	}
	return *o.PriceInUsd
}

// GetPriceInUsdOk returns a tuple with the PriceInUsd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetPriceInUsdOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceInUsd) {
		return nil, false
	}
	return o.PriceInUsd, true
}

// HasPriceInUsd returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasPriceInUsd() bool {
	if o != nil && !IsNil(o.PriceInUsd) {
		return true
	}

	return false
}

// SetPriceInUsd gets a reference to the given float32 and assigns it to the PriceInUsd field.
func (o *AzureProductVariantCustomMeter) SetPriceInUsd(v float32) {
	o.PriceInUsd = &v
}

// GetUniqueID returns the UniqueID field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetUniqueID() string {
	if o == nil || IsNil(o.UniqueID) {
		var ret string
		return ret
	}
	return *o.UniqueID
}

// GetUniqueIDOk returns a tuple with the UniqueID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetUniqueIDOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueID) {
		return nil, false
	}
	return o.UniqueID, true
}

// HasUniqueID returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasUniqueID() bool {
	if o != nil && !IsNil(o.UniqueID) {
		return true
	}

	return false
}

// SetUniqueID gets a reference to the given string and assigns it to the UniqueID field.
func (o *AzureProductVariantCustomMeter) SetUniqueID(v string) {
	o.UniqueID = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *AzureProductVariantCustomMeter) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductVariantCustomMeter) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *AzureProductVariantCustomMeter) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *AzureProductVariantCustomMeter) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

func (o AzureProductVariantCustomMeter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProductVariantCustomMeter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IncludedBaseQuantities) {
		toSerialize["includedBaseQuantities"] = o.IncludedBaseQuantities
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.PriceInUsd) {
		toSerialize["priceInUsd"] = o.PriceInUsd
	}
	if !IsNil(o.UniqueID) {
		toSerialize["uniqueID"] = o.UniqueID
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	return toSerialize, nil
}

type NullableAzureProductVariantCustomMeter struct {
	value *AzureProductVariantCustomMeter
	isSet bool
}

func (v NullableAzureProductVariantCustomMeter) Get() *AzureProductVariantCustomMeter {
	return v.value
}

func (v *NullableAzureProductVariantCustomMeter) Set(val *AzureProductVariantCustomMeter) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProductVariantCustomMeter) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProductVariantCustomMeter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProductVariantCustomMeter(val *AzureProductVariantCustomMeter) *NullableAzureProductVariantCustomMeter {
	return &NullableAzureProductVariantCustomMeter{value: val, isSet: true}
}

func (v NullableAzureProductVariantCustomMeter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProductVariantCustomMeter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
