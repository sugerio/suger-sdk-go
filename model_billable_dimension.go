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

// checks if the BillableDimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillableDimension{}

// BillableDimension struct for BillableDimension
type BillableDimension struct {
	// The ID for the billable metric.
	BillableMetricId *string `json:"billableMetricId,omitempty"`
	// The category of the pricing model. This is used to determine which pricing model to use.
	Category *PriceModelCategory `json:"category,omitempty"`
	// Description of the dimension. This is used in the UI to display the dimension.
	Description *string `json:"description,omitempty"`
	// Discount for the dimension.
	Discount *BillingDiscount `json:"discount,omitempty"`
	// The term length for the commit amount. Applicable to Direct only.
	Length *int32 `json:"length,omitempty"`
	// The minimum commit amount. Applicable to Direct only. Ignored if the value is 0 or less.
	MinimumCommit *float32 `json:"minimumCommit,omitempty"`
	// If the minimum commit is appled with pro-rata. Applicable to Direct only. If true, the minimum commit amount will be prorated based on the usage period (starting time and ending time).
	MinimumCommitProrata *bool `json:"minimumCommitProrata,omitempty"`
	// The minimum commit scope. The default value is \"DIMENSION\" if not set.
	MinimumCommitScope *BillingMinimumCommitScope `json:"minimumCommitScope,omitempty"`
	// Display name of the dimension. This is used in the UI to display the dimension.
	Name *string `json:"name,omitempty"`
	// The configuration for the Basic pricing model. Applicable to Direct only.
	PriceModelBasic *PriceModelBasic `json:"priceModelBasic,omitempty"`
	// The configuration for the Package pricing model. Applicable to Direct only.
	PriceModelBulk *PriceModelBulk `json:"priceModelBulk,omitempty"`
	// The configuration for the Matrix pricing model. Applicable to Direct only.
	PriceModelMatrix *PriceModelMatrix `json:"priceModelMatrix,omitempty"`
	// The configuration for the Percentage pricing model. Applicable to Direct only.
	PriceModelPercentage *PriceModelPercentage `json:"priceModelPercentage,omitempty"`
	// The configuration for the Tiered pricing model. Applicable to Direct only.
	PriceModelTiered *PriceModelTiered `json:"priceModelTiered,omitempty"`
	// The configuration for the Tiered Percentage pricing model. Applicable to Direct only.
	PriceModelTieredPercentage *PriceModelTieredPercentage `json:"priceModelTieredPercentage,omitempty"`
	// The configuration for the Bulk pricing model. Applicable to Direct only.
	PriceModelVolume *PriceModelVolume `json:"priceModelVolume,omitempty"`
	// The term unit for the commit amount. Applicable to Direct only.
	TimeUnit *TimeUnit `json:"timeUnit,omitempty"`
}

// NewBillableDimension instantiates a new BillableDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillableDimension() *BillableDimension {
	this := BillableDimension{}
	return &this
}

// NewBillableDimensionWithDefaults instantiates a new BillableDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillableDimensionWithDefaults() *BillableDimension {
	this := BillableDimension{}
	return &this
}

// GetBillableMetricId returns the BillableMetricId field value if set, zero value otherwise.
func (o *BillableDimension) GetBillableMetricId() string {
	if o == nil || IsNil(o.BillableMetricId) {
		var ret string
		return ret
	}
	return *o.BillableMetricId
}

// GetBillableMetricIdOk returns a tuple with the BillableMetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetBillableMetricIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillableMetricId) {
		return nil, false
	}
	return o.BillableMetricId, true
}

// HasBillableMetricId returns a boolean if a field has been set.
func (o *BillableDimension) HasBillableMetricId() bool {
	if o != nil && !IsNil(o.BillableMetricId) {
		return true
	}

	return false
}

// SetBillableMetricId gets a reference to the given string and assigns it to the BillableMetricId field.
func (o *BillableDimension) SetBillableMetricId(v string) {
	o.BillableMetricId = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *BillableDimension) GetCategory() PriceModelCategory {
	if o == nil || IsNil(o.Category) {
		var ret PriceModelCategory
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetCategoryOk() (*PriceModelCategory, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *BillableDimension) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given PriceModelCategory and assigns it to the Category field.
func (o *BillableDimension) SetCategory(v PriceModelCategory) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BillableDimension) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BillableDimension) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BillableDimension) SetDescription(v string) {
	o.Description = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *BillableDimension) GetDiscount() BillingDiscount {
	if o == nil || IsNil(o.Discount) {
		var ret BillingDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetDiscountOk() (*BillingDiscount, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *BillableDimension) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given BillingDiscount and assigns it to the Discount field.
func (o *BillableDimension) SetDiscount(v BillingDiscount) {
	o.Discount = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *BillableDimension) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *BillableDimension) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *BillableDimension) SetLength(v int32) {
	o.Length = &v
}

// GetMinimumCommit returns the MinimumCommit field value if set, zero value otherwise.
func (o *BillableDimension) GetMinimumCommit() float32 {
	if o == nil || IsNil(o.MinimumCommit) {
		var ret float32
		return ret
	}
	return *o.MinimumCommit
}

// GetMinimumCommitOk returns a tuple with the MinimumCommit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetMinimumCommitOk() (*float32, bool) {
	if o == nil || IsNil(o.MinimumCommit) {
		return nil, false
	}
	return o.MinimumCommit, true
}

// HasMinimumCommit returns a boolean if a field has been set.
func (o *BillableDimension) HasMinimumCommit() bool {
	if o != nil && !IsNil(o.MinimumCommit) {
		return true
	}

	return false
}

// SetMinimumCommit gets a reference to the given float32 and assigns it to the MinimumCommit field.
func (o *BillableDimension) SetMinimumCommit(v float32) {
	o.MinimumCommit = &v
}

// GetMinimumCommitProrata returns the MinimumCommitProrata field value if set, zero value otherwise.
func (o *BillableDimension) GetMinimumCommitProrata() bool {
	if o == nil || IsNil(o.MinimumCommitProrata) {
		var ret bool
		return ret
	}
	return *o.MinimumCommitProrata
}

// GetMinimumCommitProrataOk returns a tuple with the MinimumCommitProrata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetMinimumCommitProrataOk() (*bool, bool) {
	if o == nil || IsNil(o.MinimumCommitProrata) {
		return nil, false
	}
	return o.MinimumCommitProrata, true
}

// HasMinimumCommitProrata returns a boolean if a field has been set.
func (o *BillableDimension) HasMinimumCommitProrata() bool {
	if o != nil && !IsNil(o.MinimumCommitProrata) {
		return true
	}

	return false
}

// SetMinimumCommitProrata gets a reference to the given bool and assigns it to the MinimumCommitProrata field.
func (o *BillableDimension) SetMinimumCommitProrata(v bool) {
	o.MinimumCommitProrata = &v
}

// GetMinimumCommitScope returns the MinimumCommitScope field value if set, zero value otherwise.
func (o *BillableDimension) GetMinimumCommitScope() BillingMinimumCommitScope {
	if o == nil || IsNil(o.MinimumCommitScope) {
		var ret BillingMinimumCommitScope
		return ret
	}
	return *o.MinimumCommitScope
}

// GetMinimumCommitScopeOk returns a tuple with the MinimumCommitScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetMinimumCommitScopeOk() (*BillingMinimumCommitScope, bool) {
	if o == nil || IsNil(o.MinimumCommitScope) {
		return nil, false
	}
	return o.MinimumCommitScope, true
}

// HasMinimumCommitScope returns a boolean if a field has been set.
func (o *BillableDimension) HasMinimumCommitScope() bool {
	if o != nil && !IsNil(o.MinimumCommitScope) {
		return true
	}

	return false
}

// SetMinimumCommitScope gets a reference to the given BillingMinimumCommitScope and assigns it to the MinimumCommitScope field.
func (o *BillableDimension) SetMinimumCommitScope(v BillingMinimumCommitScope) {
	o.MinimumCommitScope = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillableDimension) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillableDimension) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillableDimension) SetName(v string) {
	o.Name = &v
}

// GetPriceModelBasic returns the PriceModelBasic field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelBasic() PriceModelBasic {
	if o == nil || IsNil(o.PriceModelBasic) {
		var ret PriceModelBasic
		return ret
	}
	return *o.PriceModelBasic
}

// GetPriceModelBasicOk returns a tuple with the PriceModelBasic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelBasicOk() (*PriceModelBasic, bool) {
	if o == nil || IsNil(o.PriceModelBasic) {
		return nil, false
	}
	return o.PriceModelBasic, true
}

// HasPriceModelBasic returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelBasic() bool {
	if o != nil && !IsNil(o.PriceModelBasic) {
		return true
	}

	return false
}

// SetPriceModelBasic gets a reference to the given PriceModelBasic and assigns it to the PriceModelBasic field.
func (o *BillableDimension) SetPriceModelBasic(v PriceModelBasic) {
	o.PriceModelBasic = &v
}

// GetPriceModelBulk returns the PriceModelBulk field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelBulk() PriceModelBulk {
	if o == nil || IsNil(o.PriceModelBulk) {
		var ret PriceModelBulk
		return ret
	}
	return *o.PriceModelBulk
}

// GetPriceModelBulkOk returns a tuple with the PriceModelBulk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelBulkOk() (*PriceModelBulk, bool) {
	if o == nil || IsNil(o.PriceModelBulk) {
		return nil, false
	}
	return o.PriceModelBulk, true
}

// HasPriceModelBulk returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelBulk() bool {
	if o != nil && !IsNil(o.PriceModelBulk) {
		return true
	}

	return false
}

// SetPriceModelBulk gets a reference to the given PriceModelBulk and assigns it to the PriceModelBulk field.
func (o *BillableDimension) SetPriceModelBulk(v PriceModelBulk) {
	o.PriceModelBulk = &v
}

// GetPriceModelMatrix returns the PriceModelMatrix field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelMatrix() PriceModelMatrix {
	if o == nil || IsNil(o.PriceModelMatrix) {
		var ret PriceModelMatrix
		return ret
	}
	return *o.PriceModelMatrix
}

// GetPriceModelMatrixOk returns a tuple with the PriceModelMatrix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelMatrixOk() (*PriceModelMatrix, bool) {
	if o == nil || IsNil(o.PriceModelMatrix) {
		return nil, false
	}
	return o.PriceModelMatrix, true
}

// HasPriceModelMatrix returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelMatrix() bool {
	if o != nil && !IsNil(o.PriceModelMatrix) {
		return true
	}

	return false
}

// SetPriceModelMatrix gets a reference to the given PriceModelMatrix and assigns it to the PriceModelMatrix field.
func (o *BillableDimension) SetPriceModelMatrix(v PriceModelMatrix) {
	o.PriceModelMatrix = &v
}

// GetPriceModelPercentage returns the PriceModelPercentage field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelPercentage() PriceModelPercentage {
	if o == nil || IsNil(o.PriceModelPercentage) {
		var ret PriceModelPercentage
		return ret
	}
	return *o.PriceModelPercentage
}

// GetPriceModelPercentageOk returns a tuple with the PriceModelPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelPercentageOk() (*PriceModelPercentage, bool) {
	if o == nil || IsNil(o.PriceModelPercentage) {
		return nil, false
	}
	return o.PriceModelPercentage, true
}

// HasPriceModelPercentage returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelPercentage() bool {
	if o != nil && !IsNil(o.PriceModelPercentage) {
		return true
	}

	return false
}

// SetPriceModelPercentage gets a reference to the given PriceModelPercentage and assigns it to the PriceModelPercentage field.
func (o *BillableDimension) SetPriceModelPercentage(v PriceModelPercentage) {
	o.PriceModelPercentage = &v
}

// GetPriceModelTiered returns the PriceModelTiered field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelTiered() PriceModelTiered {
	if o == nil || IsNil(o.PriceModelTiered) {
		var ret PriceModelTiered
		return ret
	}
	return *o.PriceModelTiered
}

// GetPriceModelTieredOk returns a tuple with the PriceModelTiered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelTieredOk() (*PriceModelTiered, bool) {
	if o == nil || IsNil(o.PriceModelTiered) {
		return nil, false
	}
	return o.PriceModelTiered, true
}

// HasPriceModelTiered returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelTiered() bool {
	if o != nil && !IsNil(o.PriceModelTiered) {
		return true
	}

	return false
}

// SetPriceModelTiered gets a reference to the given PriceModelTiered and assigns it to the PriceModelTiered field.
func (o *BillableDimension) SetPriceModelTiered(v PriceModelTiered) {
	o.PriceModelTiered = &v
}

// GetPriceModelTieredPercentage returns the PriceModelTieredPercentage field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelTieredPercentage() PriceModelTieredPercentage {
	if o == nil || IsNil(o.PriceModelTieredPercentage) {
		var ret PriceModelTieredPercentage
		return ret
	}
	return *o.PriceModelTieredPercentage
}

// GetPriceModelTieredPercentageOk returns a tuple with the PriceModelTieredPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelTieredPercentageOk() (*PriceModelTieredPercentage, bool) {
	if o == nil || IsNil(o.PriceModelTieredPercentage) {
		return nil, false
	}
	return o.PriceModelTieredPercentage, true
}

// HasPriceModelTieredPercentage returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelTieredPercentage() bool {
	if o != nil && !IsNil(o.PriceModelTieredPercentage) {
		return true
	}

	return false
}

// SetPriceModelTieredPercentage gets a reference to the given PriceModelTieredPercentage and assigns it to the PriceModelTieredPercentage field.
func (o *BillableDimension) SetPriceModelTieredPercentage(v PriceModelTieredPercentage) {
	o.PriceModelTieredPercentage = &v
}

// GetPriceModelVolume returns the PriceModelVolume field value if set, zero value otherwise.
func (o *BillableDimension) GetPriceModelVolume() PriceModelVolume {
	if o == nil || IsNil(o.PriceModelVolume) {
		var ret PriceModelVolume
		return ret
	}
	return *o.PriceModelVolume
}

// GetPriceModelVolumeOk returns a tuple with the PriceModelVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetPriceModelVolumeOk() (*PriceModelVolume, bool) {
	if o == nil || IsNil(o.PriceModelVolume) {
		return nil, false
	}
	return o.PriceModelVolume, true
}

// HasPriceModelVolume returns a boolean if a field has been set.
func (o *BillableDimension) HasPriceModelVolume() bool {
	if o != nil && !IsNil(o.PriceModelVolume) {
		return true
	}

	return false
}

// SetPriceModelVolume gets a reference to the given PriceModelVolume and assigns it to the PriceModelVolume field.
func (o *BillableDimension) SetPriceModelVolume(v PriceModelVolume) {
	o.PriceModelVolume = &v
}

// GetTimeUnit returns the TimeUnit field value if set, zero value otherwise.
func (o *BillableDimension) GetTimeUnit() TimeUnit {
	if o == nil || IsNil(o.TimeUnit) {
		var ret TimeUnit
		return ret
	}
	return *o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillableDimension) GetTimeUnitOk() (*TimeUnit, bool) {
	if o == nil || IsNil(o.TimeUnit) {
		return nil, false
	}
	return o.TimeUnit, true
}

// HasTimeUnit returns a boolean if a field has been set.
func (o *BillableDimension) HasTimeUnit() bool {
	if o != nil && !IsNil(o.TimeUnit) {
		return true
	}

	return false
}

// SetTimeUnit gets a reference to the given TimeUnit and assigns it to the TimeUnit field.
func (o *BillableDimension) SetTimeUnit(v TimeUnit) {
	o.TimeUnit = &v
}

func (o BillableDimension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillableDimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillableMetricId) {
		toSerialize["billableMetricId"] = o.BillableMetricId
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.MinimumCommit) {
		toSerialize["minimumCommit"] = o.MinimumCommit
	}
	if !IsNil(o.MinimumCommitProrata) {
		toSerialize["minimumCommitProrata"] = o.MinimumCommitProrata
	}
	if !IsNil(o.MinimumCommitScope) {
		toSerialize["minimumCommitScope"] = o.MinimumCommitScope
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PriceModelBasic) {
		toSerialize["priceModelBasic"] = o.PriceModelBasic
	}
	if !IsNil(o.PriceModelBulk) {
		toSerialize["priceModelBulk"] = o.PriceModelBulk
	}
	if !IsNil(o.PriceModelMatrix) {
		toSerialize["priceModelMatrix"] = o.PriceModelMatrix
	}
	if !IsNil(o.PriceModelPercentage) {
		toSerialize["priceModelPercentage"] = o.PriceModelPercentage
	}
	if !IsNil(o.PriceModelTiered) {
		toSerialize["priceModelTiered"] = o.PriceModelTiered
	}
	if !IsNil(o.PriceModelTieredPercentage) {
		toSerialize["priceModelTieredPercentage"] = o.PriceModelTieredPercentage
	}
	if !IsNil(o.PriceModelVolume) {
		toSerialize["priceModelVolume"] = o.PriceModelVolume
	}
	if !IsNil(o.TimeUnit) {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	return toSerialize, nil
}

type NullableBillableDimension struct {
	value *BillableDimension
	isSet bool
}

func (v NullableBillableDimension) Get() *BillableDimension {
	return v.value
}

func (v *NullableBillableDimension) Set(val *BillableDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableBillableDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableBillableDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillableDimension(val *BillableDimension) *NullableBillableDimension {
	return &NullableBillableDimension{value: val, isSet: true}
}

func (v NullableBillableDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillableDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
