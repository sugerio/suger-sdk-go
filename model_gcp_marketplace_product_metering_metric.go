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

// checks if the GcpMarketplaceProductMeteringMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductMeteringMetric{}

// GcpMarketplaceProductMeteringMetric struct for GcpMarketplaceProductMeteringMetric
type GcpMarketplaceProductMeteringMetric struct {
	// Description: A detailed description of the metric, which can be used in documentation.
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	// such as \"min\"
	DisplayUnit *string `json:"displayUnit,omitempty"`
	// such as \"minute\"
	DisplayUnitDescription *string `json:"displayUnitDescription,omitempty"`
	// The usage metering metric/dimension key, all in lower case with underscore. It is in format of \"{plan_id}_{usage_dimension_key}\". For example, \"basic_plan_storage\".
	Id *string `json:"id,omitempty"`
	// such as \"DELTA\"
	MetricKind *string `json:"metricKind,omitempty"`
	// Name: The resource name of the metric descriptor, in format of \"{productServiceName}/{plan_id}_{usage_dimension_key}\"
	Name *string `json:"name,omitempty"`
	// Price info of this usage metering metric. Only applicable for the default offer (plan) and private offer.
	PriceTiers []GcpPriceTier `json:"priceTiers,omitempty"`
	// such as \"min\"
	ReportingUnit *string `json:"reportingUnit,omitempty"`
	// The SKU ID of this usage metering metric. Applicable only in Private Offer.
	SkuId *string `json:"skuId,omitempty"`
	// such as \"min\"
	Unit *string `json:"unit,omitempty"`
	// such as \"INT64\"
	ValueType *ValueType `json:"valueType,omitempty"`
}

// NewGcpMarketplaceProductMeteringMetric instantiates a new GcpMarketplaceProductMeteringMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductMeteringMetric() *GcpMarketplaceProductMeteringMetric {
	this := GcpMarketplaceProductMeteringMetric{}
	return &this
}

// NewGcpMarketplaceProductMeteringMetricWithDefaults instantiates a new GcpMarketplaceProductMeteringMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductMeteringMetricWithDefaults() *GcpMarketplaceProductMeteringMetric {
	this := GcpMarketplaceProductMeteringMetric{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GcpMarketplaceProductMeteringMetric) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GcpMarketplaceProductMeteringMetric) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayUnit returns the DisplayUnit field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnit() string {
	if o == nil || IsNil(o.DisplayUnit) {
		var ret string
		return ret
	}
	return *o.DisplayUnit
}

// GetDisplayUnitOk returns a tuple with the DisplayUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnitOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUnit) {
		return nil, false
	}
	return o.DisplayUnit, true
}

// HasDisplayUnit returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasDisplayUnit() bool {
	if o != nil && !IsNil(o.DisplayUnit) {
		return true
	}

	return false
}

// SetDisplayUnit gets a reference to the given string and assigns it to the DisplayUnit field.
func (o *GcpMarketplaceProductMeteringMetric) SetDisplayUnit(v string) {
	o.DisplayUnit = &v
}

// GetDisplayUnitDescription returns the DisplayUnitDescription field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnitDescription() string {
	if o == nil || IsNil(o.DisplayUnitDescription) {
		var ret string
		return ret
	}
	return *o.DisplayUnitDescription
}

// GetDisplayUnitDescriptionOk returns a tuple with the DisplayUnitDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnitDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUnitDescription) {
		return nil, false
	}
	return o.DisplayUnitDescription, true
}

// HasDisplayUnitDescription returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasDisplayUnitDescription() bool {
	if o != nil && !IsNil(o.DisplayUnitDescription) {
		return true
	}

	return false
}

// SetDisplayUnitDescription gets a reference to the given string and assigns it to the DisplayUnitDescription field.
func (o *GcpMarketplaceProductMeteringMetric) SetDisplayUnitDescription(v string) {
	o.DisplayUnitDescription = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GcpMarketplaceProductMeteringMetric) SetId(v string) {
	o.Id = &v
}

// GetMetricKind returns the MetricKind field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetMetricKind() string {
	if o == nil || IsNil(o.MetricKind) {
		var ret string
		return ret
	}
	return *o.MetricKind
}

// GetMetricKindOk returns a tuple with the MetricKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetMetricKindOk() (*string, bool) {
	if o == nil || IsNil(o.MetricKind) {
		return nil, false
	}
	return o.MetricKind, true
}

// HasMetricKind returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasMetricKind() bool {
	if o != nil && !IsNil(o.MetricKind) {
		return true
	}

	return false
}

// SetMetricKind gets a reference to the given string and assigns it to the MetricKind field.
func (o *GcpMarketplaceProductMeteringMetric) SetMetricKind(v string) {
	o.MetricKind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GcpMarketplaceProductMeteringMetric) SetName(v string) {
	o.Name = &v
}

// GetPriceTiers returns the PriceTiers field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetPriceTiers() []GcpPriceTier {
	if o == nil || IsNil(o.PriceTiers) {
		var ret []GcpPriceTier
		return ret
	}
	return o.PriceTiers
}

// GetPriceTiersOk returns a tuple with the PriceTiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetPriceTiersOk() ([]GcpPriceTier, bool) {
	if o == nil || IsNil(o.PriceTiers) {
		return nil, false
	}
	return o.PriceTiers, true
}

// HasPriceTiers returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasPriceTiers() bool {
	if o != nil && !IsNil(o.PriceTiers) {
		return true
	}

	return false
}

// SetPriceTiers gets a reference to the given []GcpPriceTier and assigns it to the PriceTiers field.
func (o *GcpMarketplaceProductMeteringMetric) SetPriceTiers(v []GcpPriceTier) {
	o.PriceTiers = v
}

// GetReportingUnit returns the ReportingUnit field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetReportingUnit() string {
	if o == nil || IsNil(o.ReportingUnit) {
		var ret string
		return ret
	}
	return *o.ReportingUnit
}

// GetReportingUnitOk returns a tuple with the ReportingUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetReportingUnitOk() (*string, bool) {
	if o == nil || IsNil(o.ReportingUnit) {
		return nil, false
	}
	return o.ReportingUnit, true
}

// HasReportingUnit returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasReportingUnit() bool {
	if o != nil && !IsNil(o.ReportingUnit) {
		return true
	}

	return false
}

// SetReportingUnit gets a reference to the given string and assigns it to the ReportingUnit field.
func (o *GcpMarketplaceProductMeteringMetric) SetReportingUnit(v string) {
	o.ReportingUnit = &v
}

// GetSkuId returns the SkuId field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetSkuId() string {
	if o == nil || IsNil(o.SkuId) {
		var ret string
		return ret
	}
	return *o.SkuId
}

// GetSkuIdOk returns a tuple with the SkuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetSkuIdOk() (*string, bool) {
	if o == nil || IsNil(o.SkuId) {
		return nil, false
	}
	return o.SkuId, true
}

// HasSkuId returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasSkuId() bool {
	if o != nil && !IsNil(o.SkuId) {
		return true
	}

	return false
}

// SetSkuId gets a reference to the given string and assigns it to the SkuId field.
func (o *GcpMarketplaceProductMeteringMetric) SetSkuId(v string) {
	o.SkuId = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GcpMarketplaceProductMeteringMetric) SetUnit(v string) {
	o.Unit = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMeteringMetric) GetValueType() ValueType {
	if o == nil || IsNil(o.ValueType) {
		var ret ValueType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMeteringMetric) GetValueTypeOk() (*ValueType, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMeteringMetric) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given ValueType and assigns it to the ValueType field.
func (o *GcpMarketplaceProductMeteringMetric) SetValueType(v ValueType) {
	o.ValueType = &v
}

func (o GcpMarketplaceProductMeteringMetric) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductMeteringMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DisplayUnit) {
		toSerialize["displayUnit"] = o.DisplayUnit
	}
	if !IsNil(o.DisplayUnitDescription) {
		toSerialize["displayUnitDescription"] = o.DisplayUnitDescription
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MetricKind) {
		toSerialize["metricKind"] = o.MetricKind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PriceTiers) {
		toSerialize["priceTiers"] = o.PriceTiers
	}
	if !IsNil(o.ReportingUnit) {
		toSerialize["reportingUnit"] = o.ReportingUnit
	}
	if !IsNil(o.SkuId) {
		toSerialize["skuId"] = o.SkuId
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductMeteringMetric struct {
	value *GcpMarketplaceProductMeteringMetric
	isSet bool
}

func (v NullableGcpMarketplaceProductMeteringMetric) Get() *GcpMarketplaceProductMeteringMetric {
	return v.value
}

func (v *NullableGcpMarketplaceProductMeteringMetric) Set(val *GcpMarketplaceProductMeteringMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductMeteringMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductMeteringMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductMeteringMetric(val *GcpMarketplaceProductMeteringMetric) *NullableGcpMarketplaceProductMeteringMetric {
	return &NullableGcpMarketplaceProductMeteringMetric{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductMeteringMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductMeteringMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
