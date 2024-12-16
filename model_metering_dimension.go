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

// checks if the MeteringDimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeteringDimension{}

// MeteringDimension The dimension to meter usage in entitlement.
type MeteringDimension struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	// how many quantities of this dimension are included in the commit.
	IncludedBaseQuantities []AzureIncludedBaseQuantity `json:"includedBaseQuantities,omitempty"`
	Key                    *string                     `json:"key,omitempty"`
	// Display name of the dimension. For GCP Marketplace, it is the metering metric ID without plan prefix.
	Name *string `json:"name,omitempty"`
	// The plan ID of the metering dimension. Applicable to GCP Marketplace only. No ISO duration suffix.
	PlanId *string `json:"planId,omitempty"`
	// The name of the plan for the metering dimension. Applicable to GCP Marketplace only. It may contains the ISO duration suffix, such as P1Y.
	PlanName *string `json:"planName,omitempty"`
	// The price tiers of the metering dimension. Applicable to GCP Marketplace only.
	PriceTiers []GcpPriceTier `json:"priceTiers,omitempty"`
	// The unit price of this usage metering dimension.
	Rate *float32 `json:"rate,omitempty"`
	// The SKU ID of the metering dimension. Applicable to GCP Marketplace only.
	SkuId *string  `json:"skuId,omitempty"`
	Types []string `json:"types,omitempty"`
	// The current Dimension Usage Count. Available when call GetEntitlement API.
	UsageCount *UsageCount `json:"usageCount,omitempty"`
	// The value type of the metering dimension quantity. Applicable to GCP Marketplace only.
	ValueType *ValueType `json:"valueType,omitempty"`
}

// NewMeteringDimension instantiates a new MeteringDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeteringDimension() *MeteringDimension {
	this := MeteringDimension{}
	return &this
}

// NewMeteringDimensionWithDefaults instantiates a new MeteringDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeteringDimensionWithDefaults() *MeteringDimension {
	this := MeteringDimension{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MeteringDimension) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MeteringDimension) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *MeteringDimension) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MeteringDimension) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MeteringDimension) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MeteringDimension) SetDescription(v string) {
	o.Description = &v
}

// GetIncludedBaseQuantities returns the IncludedBaseQuantities field value if set, zero value otherwise.
func (o *MeteringDimension) GetIncludedBaseQuantities() []AzureIncludedBaseQuantity {
	if o == nil || IsNil(o.IncludedBaseQuantities) {
		var ret []AzureIncludedBaseQuantity
		return ret
	}
	return o.IncludedBaseQuantities
}

// GetIncludedBaseQuantitiesOk returns a tuple with the IncludedBaseQuantities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetIncludedBaseQuantitiesOk() ([]AzureIncludedBaseQuantity, bool) {
	if o == nil || IsNil(o.IncludedBaseQuantities) {
		return nil, false
	}
	return o.IncludedBaseQuantities, true
}

// HasIncludedBaseQuantities returns a boolean if a field has been set.
func (o *MeteringDimension) HasIncludedBaseQuantities() bool {
	if o != nil && !IsNil(o.IncludedBaseQuantities) {
		return true
	}

	return false
}

// SetIncludedBaseQuantities gets a reference to the given []AzureIncludedBaseQuantity and assigns it to the IncludedBaseQuantities field.
func (o *MeteringDimension) SetIncludedBaseQuantities(v []AzureIncludedBaseQuantity) {
	o.IncludedBaseQuantities = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MeteringDimension) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MeteringDimension) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MeteringDimension) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MeteringDimension) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MeteringDimension) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MeteringDimension) SetName(v string) {
	o.Name = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *MeteringDimension) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *MeteringDimension) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *MeteringDimension) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *MeteringDimension) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *MeteringDimension) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *MeteringDimension) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPriceTiers returns the PriceTiers field value if set, zero value otherwise.
func (o *MeteringDimension) GetPriceTiers() []GcpPriceTier {
	if o == nil || IsNil(o.PriceTiers) {
		var ret []GcpPriceTier
		return ret
	}
	return o.PriceTiers
}

// GetPriceTiersOk returns a tuple with the PriceTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetPriceTiersOk() ([]GcpPriceTier, bool) {
	if o == nil || IsNil(o.PriceTiers) {
		return nil, false
	}
	return o.PriceTiers, true
}

// HasPriceTiers returns a boolean if a field has been set.
func (o *MeteringDimension) HasPriceTiers() bool {
	if o != nil && !IsNil(o.PriceTiers) {
		return true
	}

	return false
}

// SetPriceTiers gets a reference to the given []GcpPriceTier and assigns it to the PriceTiers field.
func (o *MeteringDimension) SetPriceTiers(v []GcpPriceTier) {
	o.PriceTiers = v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *MeteringDimension) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *MeteringDimension) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *MeteringDimension) SetRate(v float32) {
	o.Rate = &v
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *MeteringDimension) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *MeteringDimension) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *MeteringDimension) SetSkuId(v string) {
	o.SkuId = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *MeteringDimension) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *MeteringDimension) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *MeteringDimension) SetTypes(v []string) {
	o.Types = v
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *MeteringDimension) GetUsageCount() UsageCount {
	if o == nil || IsNil(o.UsageCount) {
		var ret UsageCount
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetUsageCountOk() (*UsageCount, bool) {
	if o == nil || IsNil(o.UsageCount) {
		return nil, false
	}
	return o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *MeteringDimension) HasUsageCount() bool {
	if o != nil && !IsNil(o.UsageCount) {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given UsageCount and assigns it to the UsageCount field.
func (o *MeteringDimension) SetUsageCount(v UsageCount) {
	o.UsageCount = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *MeteringDimension) GetValueType() ValueType {
	if o == nil || IsNil(o.ValueType) {
		var ret ValueType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeteringDimension) GetValueTypeOk() (*ValueType, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *MeteringDimension) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given ValueType and assigns it to the ValueType field.
func (o *MeteringDimension) SetValueType(v ValueType) {
	o.ValueType = &v
}

func (o MeteringDimension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeteringDimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IncludedBaseQuantities) {
		toSerialize["includedBaseQuantities"] = o.IncludedBaseQuantities
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.PriceTiers) {
		toSerialize["priceTiers"] = o.PriceTiers
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.SkuId) {
		toSerialize["skuId"] = o.SkuId
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.UsageCount) {
		toSerialize["usageCount"] = o.UsageCount
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	return toSerialize, nil
}

type NullableMeteringDimension struct {
	value *MeteringDimension
	isSet bool
}

func (v NullableMeteringDimension) Get() *MeteringDimension {
	return v.value
}

func (v *NullableMeteringDimension) Set(val *MeteringDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableMeteringDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableMeteringDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeteringDimension(val *MeteringDimension) *NullableMeteringDimension {
	return &NullableMeteringDimension{value: val, isSet: true}
}

func (v NullableMeteringDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeteringDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
