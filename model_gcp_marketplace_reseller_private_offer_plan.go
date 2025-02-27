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
	"time"
)

// checks if the GcpMarketplaceResellerPrivateOfferPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceResellerPrivateOfferPlan{}

// GcpMarketplaceResellerPrivateOfferPlan struct for GcpMarketplaceResellerPrivateOfferPlan
type GcpMarketplaceResellerPrivateOfferPlan struct {
	AcceptanceDeadlineTime *time.Time `json:"acceptanceDeadlineTime,omitempty"`
	// The resource name of agreement(entitlement) In format of \"projects/{projectNumber}/agreements/{agreementId}\"
	Agreement                   *string                                                            `json:"agreement,omitempty"`
	AgreementDocuments          *GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments          `json:"agreementDocuments,omitempty"`
	AmendmentContext            map[string]interface{}                                             `json:"amendmentContext,omitempty"`
	DisplayName                 *string                                                            `json:"displayName,omitempty"`
	DurationConfig              *GcpMarketplaceResellerPrivateOfferPlanDurationConfig              `json:"durationConfig,omitempty"`
	Features                    []GcpMarketplaceProductFeatureValue                                `json:"features,omitempty"`
	InstallmentTimelineTemplate *GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate `json:"installmentTimelineTemplate,omitempty"`
	IsvInfo                     *GcpMarketplaceIsvInfo                                             `json:"isvInfo,omitempty"`
	Margin                      *GcpMarketplaceResellerPrivateOfferPlanMargin                      `json:"margin,omitempty"`
	MetaInfo                    *GcpMarketplaceResellerPrivateOfferPlanMetainfo                    `json:"metaInfo,omitempty"`
	// In format of \"projects/{projectNumber}/partnerAccounts/{partnerAccountId}/resellerPrivateOfferPlans/{resellerPrivateOfferPlanId}\"
	Name                  *string                                 `json:"name,omitempty"`
	OfferTemplatePolicies *GcpMarketplaceOfferTemplatePolicies    `json:"offerTemplatePolicies,omitempty"`
	OfferTermTemplate     *GcpMarketplacePrivateOfferTermTemplate `json:"offerTermTemplate,omitempty"`
	PaymentSchedule       *PaymentScheduleType                    `json:"paymentSchedule,omitempty"`
	// Nill if this reseller private offer plan has installmentTimelineTemplate (payment installments).
	PriceModelTemplate  *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate `json:"priceModelTemplate,omitempty"`
	ProductInfo         *GcpMarketplaceProductInfo                                `json:"productInfo,omitempty"`
	ReplacementMetadata map[string]interface{}                                    `json:"replacementMetadata,omitempty"`
	// in format of \"resellOfferTemplates/{resellOfferTemplateId}\"
	ResellOfferTemplate *string                                                 `json:"resellOfferTemplate,omitempty"`
	ResellerInfo        *GcpMarketplaceResellerInfo                             `json:"resellerInfo,omitempty"`
	ReusePolicy         *GcpMarketplaceResellerPrivateOfferPlanReusePolicy      `json:"reusePolicy,omitempty"`
	StartPolicy         *GcpMarketplaceStartPolicy                              `json:"startPolicy,omitempty"`
	State               *GcpMarketplaceResellerPrivateOfferPlanState            `json:"state,omitempty"`
	StateTransitions    []GcpMarketplaceResellerPrivateOfferPlanStateTransition `json:"stateTransitions,omitempty"`
	UpdateTime          *time.Time                                              `json:"updateTime,omitempty"`
}

// NewGcpMarketplaceResellerPrivateOfferPlan instantiates a new GcpMarketplaceResellerPrivateOfferPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceResellerPrivateOfferPlan() *GcpMarketplaceResellerPrivateOfferPlan {
	this := GcpMarketplaceResellerPrivateOfferPlan{}
	return &this
}

// NewGcpMarketplaceResellerPrivateOfferPlanWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceResellerPrivateOfferPlanWithDefaults() *GcpMarketplaceResellerPrivateOfferPlan {
	this := GcpMarketplaceResellerPrivateOfferPlan{}
	return &this
}

// GetAcceptanceDeadlineTime returns the AcceptanceDeadlineTime field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAcceptanceDeadlineTime() time.Time {
	if o == nil || IsNil(o.AcceptanceDeadlineTime) {
		var ret time.Time
		return ret
	}
	return *o.AcceptanceDeadlineTime
}

// GetAcceptanceDeadlineTimeOk returns a tuple with the AcceptanceDeadlineTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAcceptanceDeadlineTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AcceptanceDeadlineTime) {
		return nil, false
	}
	return o.AcceptanceDeadlineTime, true
}

// HasAcceptanceDeadlineTime returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAcceptanceDeadlineTime() bool {
	if o != nil && !IsNil(o.AcceptanceDeadlineTime) {
		return true
	}

	return false
}

// SetAcceptanceDeadlineTime gets a reference to the given time.Time and assigns it to the AcceptanceDeadlineTime field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAcceptanceDeadlineTime(v time.Time) {
	o.AcceptanceDeadlineTime = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreement() string {
	if o == nil || IsNil(o.Agreement) {
		var ret string
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreementOk() (*string, bool) {
	if o == nil || IsNil(o.Agreement) {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAgreement() bool {
	if o != nil && !IsNil(o.Agreement) {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given string and assigns it to the Agreement field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAgreement(v string) {
	o.Agreement = &v
}

// GetAgreementDocuments returns the AgreementDocuments field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreementDocuments() GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments {
	if o == nil || IsNil(o.AgreementDocuments) {
		var ret GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments
		return ret
	}
	return *o.AgreementDocuments
}

// GetAgreementDocumentsOk returns a tuple with the AgreementDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreementDocumentsOk() (*GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments, bool) {
	if o == nil || IsNil(o.AgreementDocuments) {
		return nil, false
	}
	return o.AgreementDocuments, true
}

// HasAgreementDocuments returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAgreementDocuments() bool {
	if o != nil && !IsNil(o.AgreementDocuments) {
		return true
	}

	return false
}

// SetAgreementDocuments gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments and assigns it to the AgreementDocuments field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAgreementDocuments(v GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments) {
	o.AgreementDocuments = &v
}

// GetAmendmentContext returns the AmendmentContext field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAmendmentContext() map[string]interface{} {
	if o == nil || IsNil(o.AmendmentContext) {
		var ret map[string]interface{}
		return ret
	}
	return o.AmendmentContext
}

// GetAmendmentContextOk returns a tuple with the AmendmentContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAmendmentContextOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AmendmentContext) {
		return map[string]interface{}{}, false
	}
	return o.AmendmentContext, true
}

// HasAmendmentContext returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAmendmentContext() bool {
	if o != nil && !IsNil(o.AmendmentContext) {
		return true
	}

	return false
}

// SetAmendmentContext gets a reference to the given map[string]interface{} and assigns it to the AmendmentContext field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAmendmentContext(v map[string]interface{}) {
	o.AmendmentContext = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDurationConfig returns the DurationConfig field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDurationConfig() GcpMarketplaceResellerPrivateOfferPlanDurationConfig {
	if o == nil || IsNil(o.DurationConfig) {
		var ret GcpMarketplaceResellerPrivateOfferPlanDurationConfig
		return ret
	}
	return *o.DurationConfig
}

// GetDurationConfigOk returns a tuple with the DurationConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDurationConfigOk() (*GcpMarketplaceResellerPrivateOfferPlanDurationConfig, bool) {
	if o == nil || IsNil(o.DurationConfig) {
		return nil, false
	}
	return o.DurationConfig, true
}

// HasDurationConfig returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasDurationConfig() bool {
	if o != nil && !IsNil(o.DurationConfig) {
		return true
	}

	return false
}

// SetDurationConfig gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanDurationConfig and assigns it to the DurationConfig field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetDurationConfig(v GcpMarketplaceResellerPrivateOfferPlanDurationConfig) {
	o.DurationConfig = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetFeatures() []GcpMarketplaceProductFeatureValue {
	if o == nil || IsNil(o.Features) {
		var ret []GcpMarketplaceProductFeatureValue
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetFeaturesOk() ([]GcpMarketplaceProductFeatureValue, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []GcpMarketplaceProductFeatureValue and assigns it to the Features field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetFeatures(v []GcpMarketplaceProductFeatureValue) {
	o.Features = v
}

// GetInstallmentTimelineTemplate returns the InstallmentTimelineTemplate field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetInstallmentTimelineTemplate() GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate {
	if o == nil || IsNil(o.InstallmentTimelineTemplate) {
		var ret GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate
		return ret
	}
	return *o.InstallmentTimelineTemplate
}

// GetInstallmentTimelineTemplateOk returns a tuple with the InstallmentTimelineTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetInstallmentTimelineTemplateOk() (*GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate, bool) {
	if o == nil || IsNil(o.InstallmentTimelineTemplate) {
		return nil, false
	}
	return o.InstallmentTimelineTemplate, true
}

// HasInstallmentTimelineTemplate returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasInstallmentTimelineTemplate() bool {
	if o != nil && !IsNil(o.InstallmentTimelineTemplate) {
		return true
	}

	return false
}

// SetInstallmentTimelineTemplate gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate and assigns it to the InstallmentTimelineTemplate field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetInstallmentTimelineTemplate(v GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate) {
	o.InstallmentTimelineTemplate = &v
}

// GetIsvInfo returns the IsvInfo field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetIsvInfo() GcpMarketplaceIsvInfo {
	if o == nil || IsNil(o.IsvInfo) {
		var ret GcpMarketplaceIsvInfo
		return ret
	}
	return *o.IsvInfo
}

// GetIsvInfoOk returns a tuple with the IsvInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetIsvInfoOk() (*GcpMarketplaceIsvInfo, bool) {
	if o == nil || IsNil(o.IsvInfo) {
		return nil, false
	}
	return o.IsvInfo, true
}

// HasIsvInfo returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasIsvInfo() bool {
	if o != nil && !IsNil(o.IsvInfo) {
		return true
	}

	return false
}

// SetIsvInfo gets a reference to the given GcpMarketplaceIsvInfo and assigns it to the IsvInfo field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetIsvInfo(v GcpMarketplaceIsvInfo) {
	o.IsvInfo = &v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMargin() GcpMarketplaceResellerPrivateOfferPlanMargin {
	if o == nil || IsNil(o.Margin) {
		var ret GcpMarketplaceResellerPrivateOfferPlanMargin
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMarginOk() (*GcpMarketplaceResellerPrivateOfferPlanMargin, bool) {
	if o == nil || IsNil(o.Margin) {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasMargin() bool {
	if o != nil && !IsNil(o.Margin) {
		return true
	}

	return false
}

// SetMargin gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanMargin and assigns it to the Margin field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetMargin(v GcpMarketplaceResellerPrivateOfferPlanMargin) {
	o.Margin = &v
}

// GetMetaInfo returns the MetaInfo field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMetaInfo() GcpMarketplaceResellerPrivateOfferPlanMetainfo {
	if o == nil || IsNil(o.MetaInfo) {
		var ret GcpMarketplaceResellerPrivateOfferPlanMetainfo
		return ret
	}
	return *o.MetaInfo
}

// GetMetaInfoOk returns a tuple with the MetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMetaInfoOk() (*GcpMarketplaceResellerPrivateOfferPlanMetainfo, bool) {
	if o == nil || IsNil(o.MetaInfo) {
		return nil, false
	}
	return o.MetaInfo, true
}

// HasMetaInfo returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasMetaInfo() bool {
	if o != nil && !IsNil(o.MetaInfo) {
		return true
	}

	return false
}

// SetMetaInfo gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanMetainfo and assigns it to the MetaInfo field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetMetaInfo(v GcpMarketplaceResellerPrivateOfferPlanMetainfo) {
	o.MetaInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetName(v string) {
	o.Name = &v
}

// GetOfferTemplatePolicies returns the OfferTemplatePolicies field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTemplatePolicies() GcpMarketplaceOfferTemplatePolicies {
	if o == nil || IsNil(o.OfferTemplatePolicies) {
		var ret GcpMarketplaceOfferTemplatePolicies
		return ret
	}
	return *o.OfferTemplatePolicies
}

// GetOfferTemplatePoliciesOk returns a tuple with the OfferTemplatePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTemplatePoliciesOk() (*GcpMarketplaceOfferTemplatePolicies, bool) {
	if o == nil || IsNil(o.OfferTemplatePolicies) {
		return nil, false
	}
	return o.OfferTemplatePolicies, true
}

// HasOfferTemplatePolicies returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasOfferTemplatePolicies() bool {
	if o != nil && !IsNil(o.OfferTemplatePolicies) {
		return true
	}

	return false
}

// SetOfferTemplatePolicies gets a reference to the given GcpMarketplaceOfferTemplatePolicies and assigns it to the OfferTemplatePolicies field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetOfferTemplatePolicies(v GcpMarketplaceOfferTemplatePolicies) {
	o.OfferTemplatePolicies = &v
}

// GetOfferTermTemplate returns the OfferTermTemplate field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTermTemplate() GcpMarketplacePrivateOfferTermTemplate {
	if o == nil || IsNil(o.OfferTermTemplate) {
		var ret GcpMarketplacePrivateOfferTermTemplate
		return ret
	}
	return *o.OfferTermTemplate
}

// GetOfferTermTemplateOk returns a tuple with the OfferTermTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTermTemplateOk() (*GcpMarketplacePrivateOfferTermTemplate, bool) {
	if o == nil || IsNil(o.OfferTermTemplate) {
		return nil, false
	}
	return o.OfferTermTemplate, true
}

// HasOfferTermTemplate returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasOfferTermTemplate() bool {
	if o != nil && !IsNil(o.OfferTermTemplate) {
		return true
	}

	return false
}

// SetOfferTermTemplate gets a reference to the given GcpMarketplacePrivateOfferTermTemplate and assigns it to the OfferTermTemplate field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetOfferTermTemplate(v GcpMarketplacePrivateOfferTermTemplate) {
	o.OfferTermTemplate = &v
}

// GetPaymentSchedule returns the PaymentSchedule field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPaymentSchedule() PaymentScheduleType {
	if o == nil || IsNil(o.PaymentSchedule) {
		var ret PaymentScheduleType
		return ret
	}
	return *o.PaymentSchedule
}

// GetPaymentScheduleOk returns a tuple with the PaymentSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPaymentScheduleOk() (*PaymentScheduleType, bool) {
	if o == nil || IsNil(o.PaymentSchedule) {
		return nil, false
	}
	return o.PaymentSchedule, true
}

// HasPaymentSchedule returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasPaymentSchedule() bool {
	if o != nil && !IsNil(o.PaymentSchedule) {
		return true
	}

	return false
}

// SetPaymentSchedule gets a reference to the given PaymentScheduleType and assigns it to the PaymentSchedule field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetPaymentSchedule(v PaymentScheduleType) {
	o.PaymentSchedule = &v
}

// GetPriceModelTemplate returns the PriceModelTemplate field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPriceModelTemplate() GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate {
	if o == nil || IsNil(o.PriceModelTemplate) {
		var ret GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate
		return ret
	}
	return *o.PriceModelTemplate
}

// GetPriceModelTemplateOk returns a tuple with the PriceModelTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPriceModelTemplateOk() (*GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate, bool) {
	if o == nil || IsNil(o.PriceModelTemplate) {
		return nil, false
	}
	return o.PriceModelTemplate, true
}

// HasPriceModelTemplate returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasPriceModelTemplate() bool {
	if o != nil && !IsNil(o.PriceModelTemplate) {
		return true
	}

	return false
}

// SetPriceModelTemplate gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate and assigns it to the PriceModelTemplate field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetPriceModelTemplate(v GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) {
	o.PriceModelTemplate = &v
}

// GetProductInfo returns the ProductInfo field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetProductInfo() GcpMarketplaceProductInfo {
	if o == nil || IsNil(o.ProductInfo) {
		var ret GcpMarketplaceProductInfo
		return ret
	}
	return *o.ProductInfo
}

// GetProductInfoOk returns a tuple with the ProductInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetProductInfoOk() (*GcpMarketplaceProductInfo, bool) {
	if o == nil || IsNil(o.ProductInfo) {
		return nil, false
	}
	return o.ProductInfo, true
}

// HasProductInfo returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasProductInfo() bool {
	if o != nil && !IsNil(o.ProductInfo) {
		return true
	}

	return false
}

// SetProductInfo gets a reference to the given GcpMarketplaceProductInfo and assigns it to the ProductInfo field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetProductInfo(v GcpMarketplaceProductInfo) {
	o.ProductInfo = &v
}

// GetReplacementMetadata returns the ReplacementMetadata field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReplacementMetadata() map[string]interface{} {
	if o == nil || IsNil(o.ReplacementMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.ReplacementMetadata
}

// GetReplacementMetadataOk returns a tuple with the ReplacementMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReplacementMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ReplacementMetadata) {
		return map[string]interface{}{}, false
	}
	return o.ReplacementMetadata, true
}

// HasReplacementMetadata returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasReplacementMetadata() bool {
	if o != nil && !IsNil(o.ReplacementMetadata) {
		return true
	}

	return false
}

// SetReplacementMetadata gets a reference to the given map[string]interface{} and assigns it to the ReplacementMetadata field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetReplacementMetadata(v map[string]interface{}) {
	o.ReplacementMetadata = v
}

// GetResellOfferTemplate returns the ResellOfferTemplate field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellOfferTemplate() string {
	if o == nil || IsNil(o.ResellOfferTemplate) {
		var ret string
		return ret
	}
	return *o.ResellOfferTemplate
}

// GetResellOfferTemplateOk returns a tuple with the ResellOfferTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellOfferTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.ResellOfferTemplate) {
		return nil, false
	}
	return o.ResellOfferTemplate, true
}

// HasResellOfferTemplate returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasResellOfferTemplate() bool {
	if o != nil && !IsNil(o.ResellOfferTemplate) {
		return true
	}

	return false
}

// SetResellOfferTemplate gets a reference to the given string and assigns it to the ResellOfferTemplate field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetResellOfferTemplate(v string) {
	o.ResellOfferTemplate = &v
}

// GetResellerInfo returns the ResellerInfo field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellerInfo() GcpMarketplaceResellerInfo {
	if o == nil || IsNil(o.ResellerInfo) {
		var ret GcpMarketplaceResellerInfo
		return ret
	}
	return *o.ResellerInfo
}

// GetResellerInfoOk returns a tuple with the ResellerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellerInfoOk() (*GcpMarketplaceResellerInfo, bool) {
	if o == nil || IsNil(o.ResellerInfo) {
		return nil, false
	}
	return o.ResellerInfo, true
}

// HasResellerInfo returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasResellerInfo() bool {
	if o != nil && !IsNil(o.ResellerInfo) {
		return true
	}

	return false
}

// SetResellerInfo gets a reference to the given GcpMarketplaceResellerInfo and assigns it to the ResellerInfo field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetResellerInfo(v GcpMarketplaceResellerInfo) {
	o.ResellerInfo = &v
}

// GetReusePolicy returns the ReusePolicy field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReusePolicy() GcpMarketplaceResellerPrivateOfferPlanReusePolicy {
	if o == nil || IsNil(o.ReusePolicy) {
		var ret GcpMarketplaceResellerPrivateOfferPlanReusePolicy
		return ret
	}
	return *o.ReusePolicy
}

// GetReusePolicyOk returns a tuple with the ReusePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReusePolicyOk() (*GcpMarketplaceResellerPrivateOfferPlanReusePolicy, bool) {
	if o == nil || IsNil(o.ReusePolicy) {
		return nil, false
	}
	return o.ReusePolicy, true
}

// HasReusePolicy returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasReusePolicy() bool {
	if o != nil && !IsNil(o.ReusePolicy) {
		return true
	}

	return false
}

// SetReusePolicy gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanReusePolicy and assigns it to the ReusePolicy field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetReusePolicy(v GcpMarketplaceResellerPrivateOfferPlanReusePolicy) {
	o.ReusePolicy = &v
}

// GetStartPolicy returns the StartPolicy field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStartPolicy() GcpMarketplaceStartPolicy {
	if o == nil || IsNil(o.StartPolicy) {
		var ret GcpMarketplaceStartPolicy
		return ret
	}
	return *o.StartPolicy
}

// GetStartPolicyOk returns a tuple with the StartPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStartPolicyOk() (*GcpMarketplaceStartPolicy, bool) {
	if o == nil || IsNil(o.StartPolicy) {
		return nil, false
	}
	return o.StartPolicy, true
}

// HasStartPolicy returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasStartPolicy() bool {
	if o != nil && !IsNil(o.StartPolicy) {
		return true
	}

	return false
}

// SetStartPolicy gets a reference to the given GcpMarketplaceStartPolicy and assigns it to the StartPolicy field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetStartPolicy(v GcpMarketplaceStartPolicy) {
	o.StartPolicy = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetState() GcpMarketplaceResellerPrivateOfferPlanState {
	if o == nil || IsNil(o.State) {
		var ret GcpMarketplaceResellerPrivateOfferPlanState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStateOk() (*GcpMarketplaceResellerPrivateOfferPlanState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given GcpMarketplaceResellerPrivateOfferPlanState and assigns it to the State field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetState(v GcpMarketplaceResellerPrivateOfferPlanState) {
	o.State = &v
}

// GetStateTransitions returns the StateTransitions field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStateTransitions() []GcpMarketplaceResellerPrivateOfferPlanStateTransition {
	if o == nil || IsNil(o.StateTransitions) {
		var ret []GcpMarketplaceResellerPrivateOfferPlanStateTransition
		return ret
	}
	return o.StateTransitions
}

// GetStateTransitionsOk returns a tuple with the StateTransitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStateTransitionsOk() ([]GcpMarketplaceResellerPrivateOfferPlanStateTransition, bool) {
	if o == nil || IsNil(o.StateTransitions) {
		return nil, false
	}
	return o.StateTransitions, true
}

// HasStateTransitions returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasStateTransitions() bool {
	if o != nil && !IsNil(o.StateTransitions) {
		return true
	}

	return false
}

// SetStateTransitions gets a reference to the given []GcpMarketplaceResellerPrivateOfferPlanStateTransition and assigns it to the StateTransitions field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetStateTransitions(v []GcpMarketplaceResellerPrivateOfferPlanStateTransition) {
	o.StateTransitions = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GcpMarketplaceResellerPrivateOfferPlan) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *GcpMarketplaceResellerPrivateOfferPlan) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

func (o GcpMarketplaceResellerPrivateOfferPlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceResellerPrivateOfferPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptanceDeadlineTime) {
		toSerialize["acceptanceDeadlineTime"] = o.AcceptanceDeadlineTime
	}
	if !IsNil(o.Agreement) {
		toSerialize["agreement"] = o.Agreement
	}
	if !IsNil(o.AgreementDocuments) {
		toSerialize["agreementDocuments"] = o.AgreementDocuments
	}
	if !IsNil(o.AmendmentContext) {
		toSerialize["amendmentContext"] = o.AmendmentContext
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DurationConfig) {
		toSerialize["durationConfig"] = o.DurationConfig
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.InstallmentTimelineTemplate) {
		toSerialize["installmentTimelineTemplate"] = o.InstallmentTimelineTemplate
	}
	if !IsNil(o.IsvInfo) {
		toSerialize["isvInfo"] = o.IsvInfo
	}
	if !IsNil(o.Margin) {
		toSerialize["margin"] = o.Margin
	}
	if !IsNil(o.MetaInfo) {
		toSerialize["metaInfo"] = o.MetaInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OfferTemplatePolicies) {
		toSerialize["offerTemplatePolicies"] = o.OfferTemplatePolicies
	}
	if !IsNil(o.OfferTermTemplate) {
		toSerialize["offerTermTemplate"] = o.OfferTermTemplate
	}
	if !IsNil(o.PaymentSchedule) {
		toSerialize["paymentSchedule"] = o.PaymentSchedule
	}
	if !IsNil(o.PriceModelTemplate) {
		toSerialize["priceModelTemplate"] = o.PriceModelTemplate
	}
	if !IsNil(o.ProductInfo) {
		toSerialize["productInfo"] = o.ProductInfo
	}
	if !IsNil(o.ReplacementMetadata) {
		toSerialize["replacementMetadata"] = o.ReplacementMetadata
	}
	if !IsNil(o.ResellOfferTemplate) {
		toSerialize["resellOfferTemplate"] = o.ResellOfferTemplate
	}
	if !IsNil(o.ResellerInfo) {
		toSerialize["resellerInfo"] = o.ResellerInfo
	}
	if !IsNil(o.ReusePolicy) {
		toSerialize["reusePolicy"] = o.ReusePolicy
	}
	if !IsNil(o.StartPolicy) {
		toSerialize["startPolicy"] = o.StartPolicy
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateTransitions) {
		toSerialize["stateTransitions"] = o.StateTransitions
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["updateTime"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceResellerPrivateOfferPlan struct {
	value *GcpMarketplaceResellerPrivateOfferPlan
	isSet bool
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlan) Get() *GcpMarketplaceResellerPrivateOfferPlan {
	return v.value
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlan) Set(val *GcpMarketplaceResellerPrivateOfferPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceResellerPrivateOfferPlan(val *GcpMarketplaceResellerPrivateOfferPlan) *NullableGcpMarketplaceResellerPrivateOfferPlan {
	return &NullableGcpMarketplaceResellerPrivateOfferPlan{value: val, isSet: true}
}

func (v NullableGcpMarketplaceResellerPrivateOfferPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceResellerPrivateOfferPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
