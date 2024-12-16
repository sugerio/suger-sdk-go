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

// checks if the AwsMarketplaceCppoPriceTermEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceCppoPriceTermEntry{}

// AwsMarketplaceCppoPriceTermEntry struct for AwsMarketplaceCppoPriceTermEntry
type AwsMarketplaceCppoPriceTermEntry struct {
	ConsumptionUnitColumnNames []string `json:"consumptionUnitColumnNames,omitempty"`
	Description                *string  `json:"description,omitempty"`
	// the dimension display name
	DisplayName       *string `json:"displayName,omitempty"`
	IsCustomDimension *bool   `json:"isCustomDimension,omitempty"`
	IsDeleted         *bool   `json:"isDeleted,omitempty"`
	// The dimension Key
	Name *string `json:"name,omitempty"`
	// Key: the unit in ConsumptionUnitColumnName, Value: the unit price
	PricePerConsumptionUnit *map[string]string `json:"pricePerConsumptionUnit,omitempty"`
	PricingDimension        *string            `json:"pricingDimension,omitempty"`
}

// NewAwsMarketplaceCppoPriceTermEntry instantiates a new AwsMarketplaceCppoPriceTermEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceCppoPriceTermEntry() *AwsMarketplaceCppoPriceTermEntry {
	this := AwsMarketplaceCppoPriceTermEntry{}
	return &this
}

// NewAwsMarketplaceCppoPriceTermEntryWithDefaults instantiates a new AwsMarketplaceCppoPriceTermEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceCppoPriceTermEntryWithDefaults() *AwsMarketplaceCppoPriceTermEntry {
	this := AwsMarketplaceCppoPriceTermEntry{}
	return &this
}

// GetConsumptionUnitColumnNames returns the ConsumptionUnitColumnNames field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetConsumptionUnitColumnNames() []string {
	if o == nil || IsNil(o.ConsumptionUnitColumnNames) {
		var ret []string
		return ret
	}
	return o.ConsumptionUnitColumnNames
}

// GetConsumptionUnitColumnNamesOk returns a tuple with the ConsumptionUnitColumnNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetConsumptionUnitColumnNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ConsumptionUnitColumnNames) {
		return nil, false
	}
	return o.ConsumptionUnitColumnNames, true
}

// HasConsumptionUnitColumnNames returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasConsumptionUnitColumnNames() bool {
	if o != nil && !IsNil(o.ConsumptionUnitColumnNames) {
		return true
	}

	return false
}

// SetConsumptionUnitColumnNames gets a reference to the given []string and assigns it to the ConsumptionUnitColumnNames field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetConsumptionUnitColumnNames(v []string) {
	o.ConsumptionUnitColumnNames = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIsCustomDimension returns the IsCustomDimension field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetIsCustomDimension() bool {
	if o == nil || IsNil(o.IsCustomDimension) {
		var ret bool
		return ret
	}
	return *o.IsCustomDimension
}

// GetIsCustomDimensionOk returns a tuple with the IsCustomDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetIsCustomDimensionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustomDimension) {
		return nil, false
	}
	return o.IsCustomDimension, true
}

// HasIsCustomDimension returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasIsCustomDimension() bool {
	if o != nil && !IsNil(o.IsCustomDimension) {
		return true
	}

	return false
}

// SetIsCustomDimension gets a reference to the given bool and assigns it to the IsCustomDimension field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetIsCustomDimension(v bool) {
	o.IsCustomDimension = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetName(v string) {
	o.Name = &v
}

// GetPricePerConsumptionUnit returns the PricePerConsumptionUnit field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetPricePerConsumptionUnit() map[string]string {
	if o == nil || IsNil(o.PricePerConsumptionUnit) {
		var ret map[string]string
		return ret
	}
	return *o.PricePerConsumptionUnit
}

// GetPricePerConsumptionUnitOk returns a tuple with the PricePerConsumptionUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetPricePerConsumptionUnitOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PricePerConsumptionUnit) {
		return nil, false
	}
	return o.PricePerConsumptionUnit, true
}

// HasPricePerConsumptionUnit returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasPricePerConsumptionUnit() bool {
	if o != nil && !IsNil(o.PricePerConsumptionUnit) {
		return true
	}

	return false
}

// SetPricePerConsumptionUnit gets a reference to the given map[string]string and assigns it to the PricePerConsumptionUnit field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetPricePerConsumptionUnit(v map[string]string) {
	o.PricePerConsumptionUnit = &v
}

// GetPricingDimension returns the PricingDimension field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoPriceTermEntry) GetPricingDimension() string {
	if o == nil || IsNil(o.PricingDimension) {
		var ret string
		return ret
	}
	return *o.PricingDimension
}

// GetPricingDimensionOk returns a tuple with the PricingDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) GetPricingDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.PricingDimension) {
		return nil, false
	}
	return o.PricingDimension, true
}

// HasPricingDimension returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoPriceTermEntry) HasPricingDimension() bool {
	if o != nil && !IsNil(o.PricingDimension) {
		return true
	}

	return false
}

// SetPricingDimension gets a reference to the given string and assigns it to the PricingDimension field.
func (o *AwsMarketplaceCppoPriceTermEntry) SetPricingDimension(v string) {
	o.PricingDimension = &v
}

func (o AwsMarketplaceCppoPriceTermEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceCppoPriceTermEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionUnitColumnNames) {
		toSerialize["consumptionUnitColumnNames"] = o.ConsumptionUnitColumnNames
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IsCustomDimension) {
		toSerialize["isCustomDimension"] = o.IsCustomDimension
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["isDeleted"] = o.IsDeleted
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PricePerConsumptionUnit) {
		toSerialize["pricePerConsumptionUnit"] = o.PricePerConsumptionUnit
	}
	if !IsNil(o.PricingDimension) {
		toSerialize["pricingDimension"] = o.PricingDimension
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceCppoPriceTermEntry struct {
	value *AwsMarketplaceCppoPriceTermEntry
	isSet bool
}

func (v NullableAwsMarketplaceCppoPriceTermEntry) Get() *AwsMarketplaceCppoPriceTermEntry {
	return v.value
}

func (v *NullableAwsMarketplaceCppoPriceTermEntry) Set(val *AwsMarketplaceCppoPriceTermEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceCppoPriceTermEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceCppoPriceTermEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceCppoPriceTermEntry(val *AwsMarketplaceCppoPriceTermEntry) *NullableAwsMarketplaceCppoPriceTermEntry {
	return &NullableAwsMarketplaceCppoPriceTermEntry{value: val, isSet: true}
}

func (v NullableAwsMarketplaceCppoPriceTermEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceCppoPriceTermEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
