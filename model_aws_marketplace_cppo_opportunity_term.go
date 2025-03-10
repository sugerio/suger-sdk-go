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

// checks if the AwsMarketplaceCppoOpportunityTerm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMarketplaceCppoOpportunityTerm{}

// AwsMarketplaceCppoOpportunityTerm struct for AwsMarketplaceCppoOpportunityTerm
type AwsMarketplaceCppoOpportunityTerm struct {
	CurrencyCode *string                                  `json:"CurrencyCode,omitempty"`
	Documents    []AwsMarketplaceCatalogLegalTermDocument `json:"Documents,omitempty"`
	// ISO 8601 duration format. For example, \"P12M\" represents 12 months.
	Duration                  *string                                          `json:"Duration,omitempty"`
	Grants                    []AwsMarketplaceCppoOpportunityUpfrontPriceGrant `json:"Grants,omitempty"`
	Id                        *string                                          `json:"Id,omitempty"`
	MaximumAgreementStartDate *string                                          `json:"MaximumAgreementStartDate,omitempty"`
	PositiveTargeting         *AwsMarketplaceCppoOpportunityPositiveTargeting  `json:"PositiveTargeting,omitempty"`
	// For ResaleFixedUpfrontPricingTerm
	Price     *string                                    `json:"Price,omitempty"`
	RateCards []AwsMarketplaceCatalogPricingTermRateCard `json:"RateCards,omitempty"`
	// For ResalePaymentScheduleTerm
	Schedule []AwsMarketplaceCppoOpportunityPaymentSchedule `json:"Schedule,omitempty"`
	Type     *AwsMarketplaceCppoOpportunityTermType         `json:"Type,omitempty"`
}

// NewAwsMarketplaceCppoOpportunityTerm instantiates a new AwsMarketplaceCppoOpportunityTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMarketplaceCppoOpportunityTerm() *AwsMarketplaceCppoOpportunityTerm {
	this := AwsMarketplaceCppoOpportunityTerm{}
	return &this
}

// NewAwsMarketplaceCppoOpportunityTermWithDefaults instantiates a new AwsMarketplaceCppoOpportunityTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMarketplaceCppoOpportunityTermWithDefaults() *AwsMarketplaceCppoOpportunityTerm {
	this := AwsMarketplaceCppoOpportunityTerm{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetDocuments() []AwsMarketplaceCatalogLegalTermDocument {
	if o == nil || IsNil(o.Documents) {
		var ret []AwsMarketplaceCatalogLegalTermDocument
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetDocumentsOk() ([]AwsMarketplaceCatalogLegalTermDocument, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []AwsMarketplaceCatalogLegalTermDocument and assigns it to the Documents field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetDocuments(v []AwsMarketplaceCatalogLegalTermDocument) {
	o.Documents = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetDuration(v string) {
	o.Duration = &v
}

// GetGrants returns the Grants field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetGrants() []AwsMarketplaceCppoOpportunityUpfrontPriceGrant {
	if o == nil || IsNil(o.Grants) {
		var ret []AwsMarketplaceCppoOpportunityUpfrontPriceGrant
		return ret
	}
	return o.Grants
}

// GetGrantsOk returns a tuple with the Grants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetGrantsOk() ([]AwsMarketplaceCppoOpportunityUpfrontPriceGrant, bool) {
	if o == nil || IsNil(o.Grants) {
		return nil, false
	}
	return o.Grants, true
}

// HasGrants returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasGrants() bool {
	if o != nil && !IsNil(o.Grants) {
		return true
	}

	return false
}

// SetGrants gets a reference to the given []AwsMarketplaceCppoOpportunityUpfrontPriceGrant and assigns it to the Grants field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetGrants(v []AwsMarketplaceCppoOpportunityUpfrontPriceGrant) {
	o.Grants = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetId(v string) {
	o.Id = &v
}

// GetMaximumAgreementStartDate returns the MaximumAgreementStartDate field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetMaximumAgreementStartDate() string {
	if o == nil || IsNil(o.MaximumAgreementStartDate) {
		var ret string
		return ret
	}
	return *o.MaximumAgreementStartDate
}

// GetMaximumAgreementStartDateOk returns a tuple with the MaximumAgreementStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetMaximumAgreementStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.MaximumAgreementStartDate) {
		return nil, false
	}
	return o.MaximumAgreementStartDate, true
}

// HasMaximumAgreementStartDate returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasMaximumAgreementStartDate() bool {
	if o != nil && !IsNil(o.MaximumAgreementStartDate) {
		return true
	}

	return false
}

// SetMaximumAgreementStartDate gets a reference to the given string and assigns it to the MaximumAgreementStartDate field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetMaximumAgreementStartDate(v string) {
	o.MaximumAgreementStartDate = &v
}

// GetPositiveTargeting returns the PositiveTargeting field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetPositiveTargeting() AwsMarketplaceCppoOpportunityPositiveTargeting {
	if o == nil || IsNil(o.PositiveTargeting) {
		var ret AwsMarketplaceCppoOpportunityPositiveTargeting
		return ret
	}
	return *o.PositiveTargeting
}

// GetPositiveTargetingOk returns a tuple with the PositiveTargeting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetPositiveTargetingOk() (*AwsMarketplaceCppoOpportunityPositiveTargeting, bool) {
	if o == nil || IsNil(o.PositiveTargeting) {
		return nil, false
	}
	return o.PositiveTargeting, true
}

// HasPositiveTargeting returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasPositiveTargeting() bool {
	if o != nil && !IsNil(o.PositiveTargeting) {
		return true
	}

	return false
}

// SetPositiveTargeting gets a reference to the given AwsMarketplaceCppoOpportunityPositiveTargeting and assigns it to the PositiveTargeting field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetPositiveTargeting(v AwsMarketplaceCppoOpportunityPositiveTargeting) {
	o.PositiveTargeting = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetPrice(v string) {
	o.Price = &v
}

// GetRateCards returns the RateCards field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetRateCards() []AwsMarketplaceCatalogPricingTermRateCard {
	if o == nil || IsNil(o.RateCards) {
		var ret []AwsMarketplaceCatalogPricingTermRateCard
		return ret
	}
	return o.RateCards
}

// GetRateCardsOk returns a tuple with the RateCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetRateCardsOk() ([]AwsMarketplaceCatalogPricingTermRateCard, bool) {
	if o == nil || IsNil(o.RateCards) {
		return nil, false
	}
	return o.RateCards, true
}

// HasRateCards returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasRateCards() bool {
	if o != nil && !IsNil(o.RateCards) {
		return true
	}

	return false
}

// SetRateCards gets a reference to the given []AwsMarketplaceCatalogPricingTermRateCard and assigns it to the RateCards field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetRateCards(v []AwsMarketplaceCatalogPricingTermRateCard) {
	o.RateCards = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetSchedule() []AwsMarketplaceCppoOpportunityPaymentSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret []AwsMarketplaceCppoOpportunityPaymentSchedule
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetScheduleOk() ([]AwsMarketplaceCppoOpportunityPaymentSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given []AwsMarketplaceCppoOpportunityPaymentSchedule and assigns it to the Schedule field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetSchedule(v []AwsMarketplaceCppoOpportunityPaymentSchedule) {
	o.Schedule = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AwsMarketplaceCppoOpportunityTerm) GetType() AwsMarketplaceCppoOpportunityTermType {
	if o == nil || IsNil(o.Type) {
		var ret AwsMarketplaceCppoOpportunityTermType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) GetTypeOk() (*AwsMarketplaceCppoOpportunityTermType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AwsMarketplaceCppoOpportunityTerm) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AwsMarketplaceCppoOpportunityTermType and assigns it to the Type field.
func (o *AwsMarketplaceCppoOpportunityTerm) SetType(v AwsMarketplaceCppoOpportunityTermType) {
	o.Type = &v
}

func (o AwsMarketplaceCppoOpportunityTerm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMarketplaceCppoOpportunityTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["CurrencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Documents) {
		toSerialize["Documents"] = o.Documents
	}
	if !IsNil(o.Duration) {
		toSerialize["Duration"] = o.Duration
	}
	if !IsNil(o.Grants) {
		toSerialize["Grants"] = o.Grants
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.MaximumAgreementStartDate) {
		toSerialize["MaximumAgreementStartDate"] = o.MaximumAgreementStartDate
	}
	if !IsNil(o.PositiveTargeting) {
		toSerialize["PositiveTargeting"] = o.PositiveTargeting
	}
	if !IsNil(o.Price) {
		toSerialize["Price"] = o.Price
	}
	if !IsNil(o.RateCards) {
		toSerialize["RateCards"] = o.RateCards
	}
	if !IsNil(o.Schedule) {
		toSerialize["Schedule"] = o.Schedule
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAwsMarketplaceCppoOpportunityTerm struct {
	value *AwsMarketplaceCppoOpportunityTerm
	isSet bool
}

func (v NullableAwsMarketplaceCppoOpportunityTerm) Get() *AwsMarketplaceCppoOpportunityTerm {
	return v.value
}

func (v *NullableAwsMarketplaceCppoOpportunityTerm) Set(val *AwsMarketplaceCppoOpportunityTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMarketplaceCppoOpportunityTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMarketplaceCppoOpportunityTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMarketplaceCppoOpportunityTerm(val *AwsMarketplaceCppoOpportunityTerm) *NullableAwsMarketplaceCppoOpportunityTerm {
	return &NullableAwsMarketplaceCppoOpportunityTerm{value: val, isSet: true}
}

func (v NullableAwsMarketplaceCppoOpportunityTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMarketplaceCppoOpportunityTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
