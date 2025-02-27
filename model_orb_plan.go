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

// checks if the OrbPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrbPlan{}

// OrbPlan struct for OrbPlan
type OrbPlan struct {
	// The parent plan id if the given plan was created by overriding one or more of the parent's prices.
	BasePlanId         *string           `json:"base_plan_id,omitempty"`
	CreatedAt          *string           `json:"created_at,omitempty"`
	Currency           *string           `json:"currency,omitempty"`
	DefaultInvoiceMemo *string           `json:"default_invoice_memo,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Discount           *OrbPriceDiscount `json:"discount,omitempty"`
	ExternalPlanId     *string           `json:"external_plan_id,omitempty"`
	Id                 *string           `json:"id,omitempty"`
	InvoicingCurrency  *string           `json:"invoicing_currency,omitempty"`
	Maximum            *OrbPriceMaximum  `json:"maximum,omitempty"`
	// The following fields are populated by Suger. The suger metering dimensions that are mapped to the orb billable metrics.
	MeteringDimensions []MeteringDimension `json:"metering_dimensions,omitempty"`
	Minimum            *OrbPriceMinimum    `json:"minimum,omitempty"`
	Name               *string             `json:"name,omitempty"`
	// Determines the difference between the invoice issue date and the due date. A value of \"0\" here signifies that invoices are due on issue, whereas a value of \"30\" means that the customer has a month to pay the invoice before its overdue. Note that individual subscriptions or invoices may set a different net terms configuration.
	NetTerms    *int32          `json:"net_terms,omitempty"`
	PlanPhases  []OrbPlanPhase  `json:"plan_phases,omitempty"`
	Prices      []OrbPrice      `json:"prices,omitempty"`
	Product     *OrbProduct     `json:"product,omitempty"`
	Status      *OrbPlanStatus  `json:"status,omitempty"`
	TrialConfig *OrbTrialConfig `json:"trial_config,omitempty"`
}

// NewOrbPlan instantiates a new OrbPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrbPlan() *OrbPlan {
	this := OrbPlan{}
	return &this
}

// NewOrbPlanWithDefaults instantiates a new OrbPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrbPlanWithDefaults() *OrbPlan {
	this := OrbPlan{}
	return &this
}

// GetBasePlanId returns the BasePlanId field value if set, zero value otherwise.
func (o *OrbPlan) GetBasePlanId() string {
	if o == nil || IsNil(o.BasePlanId) {
		var ret string
		return ret
	}
	return *o.BasePlanId
}

// GetBasePlanIdOk returns a tuple with the BasePlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetBasePlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.BasePlanId) {
		return nil, false
	}
	return o.BasePlanId, true
}

// HasBasePlanId returns a boolean if a field has been set.
func (o *OrbPlan) HasBasePlanId() bool {
	if o != nil && !IsNil(o.BasePlanId) {
		return true
	}

	return false
}

// SetBasePlanId gets a reference to the given string and assigns it to the BasePlanId field.
func (o *OrbPlan) SetBasePlanId(v string) {
	o.BasePlanId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrbPlan) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrbPlan) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OrbPlan) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OrbPlan) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OrbPlan) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *OrbPlan) SetCurrency(v string) {
	o.Currency = &v
}

// GetDefaultInvoiceMemo returns the DefaultInvoiceMemo field value if set, zero value otherwise.
func (o *OrbPlan) GetDefaultInvoiceMemo() string {
	if o == nil || IsNil(o.DefaultInvoiceMemo) {
		var ret string
		return ret
	}
	return *o.DefaultInvoiceMemo
}

// GetDefaultInvoiceMemoOk returns a tuple with the DefaultInvoiceMemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetDefaultInvoiceMemoOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultInvoiceMemo) {
		return nil, false
	}
	return o.DefaultInvoiceMemo, true
}

// HasDefaultInvoiceMemo returns a boolean if a field has been set.
func (o *OrbPlan) HasDefaultInvoiceMemo() bool {
	if o != nil && !IsNil(o.DefaultInvoiceMemo) {
		return true
	}

	return false
}

// SetDefaultInvoiceMemo gets a reference to the given string and assigns it to the DefaultInvoiceMemo field.
func (o *OrbPlan) SetDefaultInvoiceMemo(v string) {
	o.DefaultInvoiceMemo = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrbPlan) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrbPlan) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrbPlan) SetDescription(v string) {
	o.Description = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *OrbPlan) GetDiscount() OrbPriceDiscount {
	if o == nil || IsNil(o.Discount) {
		var ret OrbPriceDiscount
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetDiscountOk() (*OrbPriceDiscount, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *OrbPlan) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given OrbPriceDiscount and assigns it to the Discount field.
func (o *OrbPlan) SetDiscount(v OrbPriceDiscount) {
	o.Discount = &v
}

// GetExternalPlanId returns the ExternalPlanId field value if set, zero value otherwise.
func (o *OrbPlan) GetExternalPlanId() string {
	if o == nil || IsNil(o.ExternalPlanId) {
		var ret string
		return ret
	}
	return *o.ExternalPlanId
}

// GetExternalPlanIdOk returns a tuple with the ExternalPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetExternalPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPlanId) {
		return nil, false
	}
	return o.ExternalPlanId, true
}

// HasExternalPlanId returns a boolean if a field has been set.
func (o *OrbPlan) HasExternalPlanId() bool {
	if o != nil && !IsNil(o.ExternalPlanId) {
		return true
	}

	return false
}

// SetExternalPlanId gets a reference to the given string and assigns it to the ExternalPlanId field.
func (o *OrbPlan) SetExternalPlanId(v string) {
	o.ExternalPlanId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrbPlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrbPlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrbPlan) SetId(v string) {
	o.Id = &v
}

// GetInvoicingCurrency returns the InvoicingCurrency field value if set, zero value otherwise.
func (o *OrbPlan) GetInvoicingCurrency() string {
	if o == nil || IsNil(o.InvoicingCurrency) {
		var ret string
		return ret
	}
	return *o.InvoicingCurrency
}

// GetInvoicingCurrencyOk returns a tuple with the InvoicingCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetInvoicingCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.InvoicingCurrency) {
		return nil, false
	}
	return o.InvoicingCurrency, true
}

// HasInvoicingCurrency returns a boolean if a field has been set.
func (o *OrbPlan) HasInvoicingCurrency() bool {
	if o != nil && !IsNil(o.InvoicingCurrency) {
		return true
	}

	return false
}

// SetInvoicingCurrency gets a reference to the given string and assigns it to the InvoicingCurrency field.
func (o *OrbPlan) SetInvoicingCurrency(v string) {
	o.InvoicingCurrency = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *OrbPlan) GetMaximum() OrbPriceMaximum {
	if o == nil || IsNil(o.Maximum) {
		var ret OrbPriceMaximum
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetMaximumOk() (*OrbPriceMaximum, bool) {
	if o == nil || IsNil(o.Maximum) {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *OrbPlan) HasMaximum() bool {
	if o != nil && !IsNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given OrbPriceMaximum and assigns it to the Maximum field.
func (o *OrbPlan) SetMaximum(v OrbPriceMaximum) {
	o.Maximum = &v
}

// GetMeteringDimensions returns the MeteringDimensions field value if set, zero value otherwise.
func (o *OrbPlan) GetMeteringDimensions() []MeteringDimension {
	if o == nil || IsNil(o.MeteringDimensions) {
		var ret []MeteringDimension
		return ret
	}
	return o.MeteringDimensions
}

// GetMeteringDimensionsOk returns a tuple with the MeteringDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetMeteringDimensionsOk() ([]MeteringDimension, bool) {
	if o == nil || IsNil(o.MeteringDimensions) {
		return nil, false
	}
	return o.MeteringDimensions, true
}

// HasMeteringDimensions returns a boolean if a field has been set.
func (o *OrbPlan) HasMeteringDimensions() bool {
	if o != nil && !IsNil(o.MeteringDimensions) {
		return true
	}

	return false
}

// SetMeteringDimensions gets a reference to the given []MeteringDimension and assigns it to the MeteringDimensions field.
func (o *OrbPlan) SetMeteringDimensions(v []MeteringDimension) {
	o.MeteringDimensions = v
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *OrbPlan) GetMinimum() OrbPriceMinimum {
	if o == nil || IsNil(o.Minimum) {
		var ret OrbPriceMinimum
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetMinimumOk() (*OrbPriceMinimum, bool) {
	if o == nil || IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *OrbPlan) HasMinimum() bool {
	if o != nil && !IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given OrbPriceMinimum and assigns it to the Minimum field.
func (o *OrbPlan) SetMinimum(v OrbPriceMinimum) {
	o.Minimum = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrbPlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrbPlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrbPlan) SetName(v string) {
	o.Name = &v
}

// GetNetTerms returns the NetTerms field value if set, zero value otherwise.
func (o *OrbPlan) GetNetTerms() int32 {
	if o == nil || IsNil(o.NetTerms) {
		var ret int32
		return ret
	}
	return *o.NetTerms
}

// GetNetTermsOk returns a tuple with the NetTerms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetNetTermsOk() (*int32, bool) {
	if o == nil || IsNil(o.NetTerms) {
		return nil, false
	}
	return o.NetTerms, true
}

// HasNetTerms returns a boolean if a field has been set.
func (o *OrbPlan) HasNetTerms() bool {
	if o != nil && !IsNil(o.NetTerms) {
		return true
	}

	return false
}

// SetNetTerms gets a reference to the given int32 and assigns it to the NetTerms field.
func (o *OrbPlan) SetNetTerms(v int32) {
	o.NetTerms = &v
}

// GetPlanPhases returns the PlanPhases field value if set, zero value otherwise.
func (o *OrbPlan) GetPlanPhases() []OrbPlanPhase {
	if o == nil || IsNil(o.PlanPhases) {
		var ret []OrbPlanPhase
		return ret
	}
	return o.PlanPhases
}

// GetPlanPhasesOk returns a tuple with the PlanPhases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetPlanPhasesOk() ([]OrbPlanPhase, bool) {
	if o == nil || IsNil(o.PlanPhases) {
		return nil, false
	}
	return o.PlanPhases, true
}

// HasPlanPhases returns a boolean if a field has been set.
func (o *OrbPlan) HasPlanPhases() bool {
	if o != nil && !IsNil(o.PlanPhases) {
		return true
	}

	return false
}

// SetPlanPhases gets a reference to the given []OrbPlanPhase and assigns it to the PlanPhases field.
func (o *OrbPlan) SetPlanPhases(v []OrbPlanPhase) {
	o.PlanPhases = v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *OrbPlan) GetPrices() []OrbPrice {
	if o == nil || IsNil(o.Prices) {
		var ret []OrbPrice
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetPricesOk() ([]OrbPrice, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *OrbPlan) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []OrbPrice and assigns it to the Prices field.
func (o *OrbPlan) SetPrices(v []OrbPrice) {
	o.Prices = v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *OrbPlan) GetProduct() OrbProduct {
	if o == nil || IsNil(o.Product) {
		var ret OrbProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetProductOk() (*OrbProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *OrbPlan) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given OrbProduct and assigns it to the Product field.
func (o *OrbPlan) SetProduct(v OrbProduct) {
	o.Product = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrbPlan) GetStatus() OrbPlanStatus {
	if o == nil || IsNil(o.Status) {
		var ret OrbPlanStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetStatusOk() (*OrbPlanStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrbPlan) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OrbPlanStatus and assigns it to the Status field.
func (o *OrbPlan) SetStatus(v OrbPlanStatus) {
	o.Status = &v
}

// GetTrialConfig returns the TrialConfig field value if set, zero value otherwise.
func (o *OrbPlan) GetTrialConfig() OrbTrialConfig {
	if o == nil || IsNil(o.TrialConfig) {
		var ret OrbTrialConfig
		return ret
	}
	return *o.TrialConfig
}

// GetTrialConfigOk returns a tuple with the TrialConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrbPlan) GetTrialConfigOk() (*OrbTrialConfig, bool) {
	if o == nil || IsNil(o.TrialConfig) {
		return nil, false
	}
	return o.TrialConfig, true
}

// HasTrialConfig returns a boolean if a field has been set.
func (o *OrbPlan) HasTrialConfig() bool {
	if o != nil && !IsNil(o.TrialConfig) {
		return true
	}

	return false
}

// SetTrialConfig gets a reference to the given OrbTrialConfig and assigns it to the TrialConfig field.
func (o *OrbPlan) SetTrialConfig(v OrbTrialConfig) {
	o.TrialConfig = &v
}

func (o OrbPlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrbPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasePlanId) {
		toSerialize["base_plan_id"] = o.BasePlanId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.DefaultInvoiceMemo) {
		toSerialize["default_invoice_memo"] = o.DefaultInvoiceMemo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.ExternalPlanId) {
		toSerialize["external_plan_id"] = o.ExternalPlanId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoicingCurrency) {
		toSerialize["invoicing_currency"] = o.InvoicingCurrency
	}
	if !IsNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	if !IsNil(o.MeteringDimensions) {
		toSerialize["metering_dimensions"] = o.MeteringDimensions
	}
	if !IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NetTerms) {
		toSerialize["net_terms"] = o.NetTerms
	}
	if !IsNil(o.PlanPhases) {
		toSerialize["plan_phases"] = o.PlanPhases
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TrialConfig) {
		toSerialize["trial_config"] = o.TrialConfig
	}
	return toSerialize, nil
}

type NullableOrbPlan struct {
	value *OrbPlan
	isSet bool
}

func (v NullableOrbPlan) Get() *OrbPlan {
	return v.value
}

func (v *NullableOrbPlan) Set(val *OrbPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableOrbPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableOrbPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrbPlan(val *OrbPlan) *NullableOrbPlan {
	return &NullableOrbPlan{value: val, isSet: true}
}

func (v NullableOrbPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrbPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
