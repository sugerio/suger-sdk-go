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

// checks if the WorkloadMetaInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkloadMetaInfo{}

// WorkloadMetaInfo struct for WorkloadMetaInfo
type WorkloadMetaInfo struct {
	// The linked ACE ApnCrmUniqueIdentifier of the private offer if available.
	AceApnCrmUniqueIdentifier *string `json:"aceApnCrmUniqueIdentifier,omitempty"`
	// The AWS SaaS product dimensions. Applicable for AWS SaaS products only. This is used to save price info when creating AWS SaaS product.
	AwsSaasProductDimensions []AwsProductDimension `json:"awsSaasProductDimensions,omitempty"`
	// Applicable for AWS Marketplace only, when the IsAgreementBasedOffer is true.
	BaseAgreementId *string `json:"baseAgreementId,omitempty"`
	// The Suger buyer IDs of the private offer if available.
	BuyerIds []string `json:"buyerIds,omitempty"`
	// The contacts of the offer to notify if any updates.
	Contacts []Contact `json:"contacts,omitempty"`
	// The Suger CPPO_IN offer ID.
	CppoInOfferId *string `json:"cppoInOfferId,omitempty"`
	// The Suger CPPO offer ID.
	CppoOfferId *string `json:"cppoOfferId,omitempty"`
	// The Suger CPPO_OUT offer ID.
	CppoOutOfferId *string `json:"cppoOutOfferId,omitempty"`
	// The custom meta info of the offer can be updated by seller via API or console.
	CustomMetaInfo *map[string]string `json:"customMetaInfo,omitempty"`
	// If enabled, Suger will test metering the usage for this entitlement hourly.
	EnableTestUsageMetering *bool `json:"enableTestUsageMetering,omitempty"`
	// The cancellation schedule for the entitlement. It is nill if no cancellation schedule.
	EntitlementCancellationSchedule *CancellationSchedule `json:"entitlementCancellationSchedule,omitempty"`
	// The error messages when the offer is invalid or offer related tasks failed. Populated by Suger service.
	ErrorMessages []string `json:"errorMessages,omitempty"`
	// Hubsport deal ID of the private offer if available.
	HubspotDealId *string `json:"hubspotDealId,omitempty"`
	// The Internal note of the private offer. It is only visible to the seller/ISV, not visible to the buyer. Up to 1000 characters.
	InternalNote *string `json:"internalNote,omitempty"`
	// Applicable for AWS Marketplace only, If this offer is agreement based offer.
	IsAgreementBasedOffer *bool `json:"isAgreementBasedOffer,omitempty"`
	// Whether the gross revenue is fully synced for the entitlement.
	IsGrossRevenueFullSync *bool `json:"isGrossRevenueFullSync,omitempty"`
	// Applicable for AWS Marketplace only. If this offer is renewal offer of existing agreement. The existing agreement can be within or outside AWS Marketplace. AWS may audit and verify your offer is a renewal. If AWS is unable to verify your offer, then AWS may revoke the offer and entitlements from your customer.
	IsRenewalOffer *bool `json:"isRenewalOffer,omitempty"`
	// If this offer is a GCP replacement offer. Applicable for GCP Marketplace replacement offer only.
	IsReplacementOffer *bool `json:"isReplacementOffer,omitempty"`
	// The user who last modified the product/offer/buyer/contact.
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// The notifications of the offer if any updates. In most cases, it is to notify contacts/buyers when the offer is pending acceptance.
	Notifications []NotificationEvent `json:"notifications,omitempty"`
	// The date when the offer is accepted by the buyer. Only available when the private offer has been accepted.
	OfferAcceptDate *time.Time `json:"offerAcceptDate,omitempty"`
	// Applicable for AWS Marketplace only, required when the IsRenewalOffer is true.
	RenewalOfferType *AwsRenewalOfferType `json:"renewalOfferType,omitempty"`
	// The end time of the replaced offer. Applicable for GCP Marketplace replacement offer only.
	ReplacedOfferEndTime *time.Time `json:"replacedOfferEndTime,omitempty"`
	// The resource name of the GCP Marketplace offer that this offer is replacing. In format of \"projects/{gcpProjectNumber}/services/{productServiceName}/privateOffers/{privateOfferId}\" Applicable for GCP Marketplace replacement offer only.
	ReplacedOfferResourceName *string `json:"replacedOfferResourceName,omitempty"`
	// The Salesforce entitlement URL
	SalesforceEntitlementURL *string `json:"salesforceEntitlementURL,omitempty"`
	// The Salesforce opportunity ID of the private offer if available.
	SalesforceOpportunityId *string `json:"salesforceOpportunityId,omitempty"`
	// The test usage metering end time. It is used for test usage metering only. Required if EnableTestUsageMetering is true.
	TestUsageMeteringEndTime *time.Time `json:"testUsageMeteringEndTime,omitempty"`
	// The message to notify when the offer is updated.
	UpdateMessage *string `json:"updateMessage,omitempty"`
}

// NewWorkloadMetaInfo instantiates a new WorkloadMetaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkloadMetaInfo() *WorkloadMetaInfo {
	this := WorkloadMetaInfo{}
	return &this
}

// NewWorkloadMetaInfoWithDefaults instantiates a new WorkloadMetaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkloadMetaInfoWithDefaults() *WorkloadMetaInfo {
	this := WorkloadMetaInfo{}
	return &this
}

// GetAceApnCrmUniqueIdentifier returns the AceApnCrmUniqueIdentifier field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetAceApnCrmUniqueIdentifier() string {
	if o == nil || IsNil(o.AceApnCrmUniqueIdentifier) {
		var ret string
		return ret
	}
	return *o.AceApnCrmUniqueIdentifier
}

// GetAceApnCrmUniqueIdentifierOk returns a tuple with the AceApnCrmUniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetAceApnCrmUniqueIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.AceApnCrmUniqueIdentifier) {
		return nil, false
	}
	return o.AceApnCrmUniqueIdentifier, true
}

// HasAceApnCrmUniqueIdentifier returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasAceApnCrmUniqueIdentifier() bool {
	if o != nil && !IsNil(o.AceApnCrmUniqueIdentifier) {
		return true
	}

	return false
}

// SetAceApnCrmUniqueIdentifier gets a reference to the given string and assigns it to the AceApnCrmUniqueIdentifier field.
func (o *WorkloadMetaInfo) SetAceApnCrmUniqueIdentifier(v string) {
	o.AceApnCrmUniqueIdentifier = &v
}

// GetAwsSaasProductDimensions returns the AwsSaasProductDimensions field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetAwsSaasProductDimensions() []AwsProductDimension {
	if o == nil || IsNil(o.AwsSaasProductDimensions) {
		var ret []AwsProductDimension
		return ret
	}
	return o.AwsSaasProductDimensions
}

// GetAwsSaasProductDimensionsOk returns a tuple with the AwsSaasProductDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetAwsSaasProductDimensionsOk() ([]AwsProductDimension, bool) {
	if o == nil || IsNil(o.AwsSaasProductDimensions) {
		return nil, false
	}
	return o.AwsSaasProductDimensions, true
}

// HasAwsSaasProductDimensions returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasAwsSaasProductDimensions() bool {
	if o != nil && !IsNil(o.AwsSaasProductDimensions) {
		return true
	}

	return false
}

// SetAwsSaasProductDimensions gets a reference to the given []AwsProductDimension and assigns it to the AwsSaasProductDimensions field.
func (o *WorkloadMetaInfo) SetAwsSaasProductDimensions(v []AwsProductDimension) {
	o.AwsSaasProductDimensions = v
}

// GetBaseAgreementId returns the BaseAgreementId field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetBaseAgreementId() string {
	if o == nil || IsNil(o.BaseAgreementId) {
		var ret string
		return ret
	}
	return *o.BaseAgreementId
}

// GetBaseAgreementIdOk returns a tuple with the BaseAgreementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetBaseAgreementIdOk() (*string, bool) {
	if o == nil || IsNil(o.BaseAgreementId) {
		return nil, false
	}
	return o.BaseAgreementId, true
}

// HasBaseAgreementId returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasBaseAgreementId() bool {
	if o != nil && !IsNil(o.BaseAgreementId) {
		return true
	}

	return false
}

// SetBaseAgreementId gets a reference to the given string and assigns it to the BaseAgreementId field.
func (o *WorkloadMetaInfo) SetBaseAgreementId(v string) {
	o.BaseAgreementId = &v
}

// GetBuyerIds returns the BuyerIds field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetBuyerIds() []string {
	if o == nil || IsNil(o.BuyerIds) {
		var ret []string
		return ret
	}
	return o.BuyerIds
}

// GetBuyerIdsOk returns a tuple with the BuyerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetBuyerIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.BuyerIds) {
		return nil, false
	}
	return o.BuyerIds, true
}

// HasBuyerIds returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasBuyerIds() bool {
	if o != nil && !IsNil(o.BuyerIds) {
		return true
	}

	return false
}

// SetBuyerIds gets a reference to the given []string and assigns it to the BuyerIds field.
func (o *WorkloadMetaInfo) SetBuyerIds(v []string) {
	o.BuyerIds = v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetContacts() []Contact {
	if o == nil || IsNil(o.Contacts) {
		var ret []Contact
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetContactsOk() ([]Contact, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasContacts() bool {
	if o != nil && !IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []Contact and assigns it to the Contacts field.
func (o *WorkloadMetaInfo) SetContacts(v []Contact) {
	o.Contacts = v
}

// GetCppoInOfferId returns the CppoInOfferId field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetCppoInOfferId() string {
	if o == nil || IsNil(o.CppoInOfferId) {
		var ret string
		return ret
	}
	return *o.CppoInOfferId
}

// GetCppoInOfferIdOk returns a tuple with the CppoInOfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetCppoInOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.CppoInOfferId) {
		return nil, false
	}
	return o.CppoInOfferId, true
}

// HasCppoInOfferId returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasCppoInOfferId() bool {
	if o != nil && !IsNil(o.CppoInOfferId) {
		return true
	}

	return false
}

// SetCppoInOfferId gets a reference to the given string and assigns it to the CppoInOfferId field.
func (o *WorkloadMetaInfo) SetCppoInOfferId(v string) {
	o.CppoInOfferId = &v
}

// GetCppoOfferId returns the CppoOfferId field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetCppoOfferId() string {
	if o == nil || IsNil(o.CppoOfferId) {
		var ret string
		return ret
	}
	return *o.CppoOfferId
}

// GetCppoOfferIdOk returns a tuple with the CppoOfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetCppoOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.CppoOfferId) {
		return nil, false
	}
	return o.CppoOfferId, true
}

// HasCppoOfferId returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasCppoOfferId() bool {
	if o != nil && !IsNil(o.CppoOfferId) {
		return true
	}

	return false
}

// SetCppoOfferId gets a reference to the given string and assigns it to the CppoOfferId field.
func (o *WorkloadMetaInfo) SetCppoOfferId(v string) {
	o.CppoOfferId = &v
}

// GetCppoOutOfferId returns the CppoOutOfferId field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetCppoOutOfferId() string {
	if o == nil || IsNil(o.CppoOutOfferId) {
		var ret string
		return ret
	}
	return *o.CppoOutOfferId
}

// GetCppoOutOfferIdOk returns a tuple with the CppoOutOfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetCppoOutOfferIdOk() (*string, bool) {
	if o == nil || IsNil(o.CppoOutOfferId) {
		return nil, false
	}
	return o.CppoOutOfferId, true
}

// HasCppoOutOfferId returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasCppoOutOfferId() bool {
	if o != nil && !IsNil(o.CppoOutOfferId) {
		return true
	}

	return false
}

// SetCppoOutOfferId gets a reference to the given string and assigns it to the CppoOutOfferId field.
func (o *WorkloadMetaInfo) SetCppoOutOfferId(v string) {
	o.CppoOutOfferId = &v
}

// GetCustomMetaInfo returns the CustomMetaInfo field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetCustomMetaInfo() map[string]string {
	if o == nil || IsNil(o.CustomMetaInfo) {
		var ret map[string]string
		return ret
	}
	return *o.CustomMetaInfo
}

// GetCustomMetaInfoOk returns a tuple with the CustomMetaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetCustomMetaInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomMetaInfo) {
		return nil, false
	}
	return o.CustomMetaInfo, true
}

// HasCustomMetaInfo returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasCustomMetaInfo() bool {
	if o != nil && !IsNil(o.CustomMetaInfo) {
		return true
	}

	return false
}

// SetCustomMetaInfo gets a reference to the given map[string]string and assigns it to the CustomMetaInfo field.
func (o *WorkloadMetaInfo) SetCustomMetaInfo(v map[string]string) {
	o.CustomMetaInfo = &v
}

// GetEnableTestUsageMetering returns the EnableTestUsageMetering field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetEnableTestUsageMetering() bool {
	if o == nil || IsNil(o.EnableTestUsageMetering) {
		var ret bool
		return ret
	}
	return *o.EnableTestUsageMetering
}

// GetEnableTestUsageMeteringOk returns a tuple with the EnableTestUsageMetering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetEnableTestUsageMeteringOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTestUsageMetering) {
		return nil, false
	}
	return o.EnableTestUsageMetering, true
}

// HasEnableTestUsageMetering returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasEnableTestUsageMetering() bool {
	if o != nil && !IsNil(o.EnableTestUsageMetering) {
		return true
	}

	return false
}

// SetEnableTestUsageMetering gets a reference to the given bool and assigns it to the EnableTestUsageMetering field.
func (o *WorkloadMetaInfo) SetEnableTestUsageMetering(v bool) {
	o.EnableTestUsageMetering = &v
}

// GetEntitlementCancellationSchedule returns the EntitlementCancellationSchedule field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetEntitlementCancellationSchedule() CancellationSchedule {
	if o == nil || IsNil(o.EntitlementCancellationSchedule) {
		var ret CancellationSchedule
		return ret
	}
	return *o.EntitlementCancellationSchedule
}

// GetEntitlementCancellationScheduleOk returns a tuple with the EntitlementCancellationSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetEntitlementCancellationScheduleOk() (*CancellationSchedule, bool) {
	if o == nil || IsNil(o.EntitlementCancellationSchedule) {
		return nil, false
	}
	return o.EntitlementCancellationSchedule, true
}

// HasEntitlementCancellationSchedule returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasEntitlementCancellationSchedule() bool {
	if o != nil && !IsNil(o.EntitlementCancellationSchedule) {
		return true
	}

	return false
}

// SetEntitlementCancellationSchedule gets a reference to the given CancellationSchedule and assigns it to the EntitlementCancellationSchedule field.
func (o *WorkloadMetaInfo) SetEntitlementCancellationSchedule(v CancellationSchedule) {
	o.EntitlementCancellationSchedule = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetErrorMessages() []string {
	if o == nil || IsNil(o.ErrorMessages) {
		var ret []string
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetErrorMessagesOk() ([]string, bool) {
	if o == nil || IsNil(o.ErrorMessages) {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasErrorMessages() bool {
	if o != nil && !IsNil(o.ErrorMessages) {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []string and assigns it to the ErrorMessages field.
func (o *WorkloadMetaInfo) SetErrorMessages(v []string) {
	o.ErrorMessages = v
}

// GetHubspotDealId returns the HubspotDealId field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetHubspotDealId() string {
	if o == nil || IsNil(o.HubspotDealId) {
		var ret string
		return ret
	}
	return *o.HubspotDealId
}

// GetHubspotDealIdOk returns a tuple with the HubspotDealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetHubspotDealIdOk() (*string, bool) {
	if o == nil || IsNil(o.HubspotDealId) {
		return nil, false
	}
	return o.HubspotDealId, true
}

// HasHubspotDealId returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasHubspotDealId() bool {
	if o != nil && !IsNil(o.HubspotDealId) {
		return true
	}

	return false
}

// SetHubspotDealId gets a reference to the given string and assigns it to the HubspotDealId field.
func (o *WorkloadMetaInfo) SetHubspotDealId(v string) {
	o.HubspotDealId = &v
}

// GetInternalNote returns the InternalNote field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetInternalNote() string {
	if o == nil || IsNil(o.InternalNote) {
		var ret string
		return ret
	}
	return *o.InternalNote
}

// GetInternalNoteOk returns a tuple with the InternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetInternalNoteOk() (*string, bool) {
	if o == nil || IsNil(o.InternalNote) {
		return nil, false
	}
	return o.InternalNote, true
}

// HasInternalNote returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasInternalNote() bool {
	if o != nil && !IsNil(o.InternalNote) {
		return true
	}

	return false
}

// SetInternalNote gets a reference to the given string and assigns it to the InternalNote field.
func (o *WorkloadMetaInfo) SetInternalNote(v string) {
	o.InternalNote = &v
}

// GetIsAgreementBasedOffer returns the IsAgreementBasedOffer field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetIsAgreementBasedOffer() bool {
	if o == nil || IsNil(o.IsAgreementBasedOffer) {
		var ret bool
		return ret
	}
	return *o.IsAgreementBasedOffer
}

// GetIsAgreementBasedOfferOk returns a tuple with the IsAgreementBasedOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetIsAgreementBasedOfferOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAgreementBasedOffer) {
		return nil, false
	}
	return o.IsAgreementBasedOffer, true
}

// HasIsAgreementBasedOffer returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasIsAgreementBasedOffer() bool {
	if o != nil && !IsNil(o.IsAgreementBasedOffer) {
		return true
	}

	return false
}

// SetIsAgreementBasedOffer gets a reference to the given bool and assigns it to the IsAgreementBasedOffer field.
func (o *WorkloadMetaInfo) SetIsAgreementBasedOffer(v bool) {
	o.IsAgreementBasedOffer = &v
}

// GetIsGrossRevenueFullSync returns the IsGrossRevenueFullSync field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetIsGrossRevenueFullSync() bool {
	if o == nil || IsNil(o.IsGrossRevenueFullSync) {
		var ret bool
		return ret
	}
	return *o.IsGrossRevenueFullSync
}

// GetIsGrossRevenueFullSyncOk returns a tuple with the IsGrossRevenueFullSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetIsGrossRevenueFullSyncOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGrossRevenueFullSync) {
		return nil, false
	}
	return o.IsGrossRevenueFullSync, true
}

// HasIsGrossRevenueFullSync returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasIsGrossRevenueFullSync() bool {
	if o != nil && !IsNil(o.IsGrossRevenueFullSync) {
		return true
	}

	return false
}

// SetIsGrossRevenueFullSync gets a reference to the given bool and assigns it to the IsGrossRevenueFullSync field.
func (o *WorkloadMetaInfo) SetIsGrossRevenueFullSync(v bool) {
	o.IsGrossRevenueFullSync = &v
}

// GetIsRenewalOffer returns the IsRenewalOffer field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetIsRenewalOffer() bool {
	if o == nil || IsNil(o.IsRenewalOffer) {
		var ret bool
		return ret
	}
	return *o.IsRenewalOffer
}

// GetIsRenewalOfferOk returns a tuple with the IsRenewalOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetIsRenewalOfferOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRenewalOffer) {
		return nil, false
	}
	return o.IsRenewalOffer, true
}

// HasIsRenewalOffer returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasIsRenewalOffer() bool {
	if o != nil && !IsNil(o.IsRenewalOffer) {
		return true
	}

	return false
}

// SetIsRenewalOffer gets a reference to the given bool and assigns it to the IsRenewalOffer field.
func (o *WorkloadMetaInfo) SetIsRenewalOffer(v bool) {
	o.IsRenewalOffer = &v
}

// GetIsReplacementOffer returns the IsReplacementOffer field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetIsReplacementOffer() bool {
	if o == nil || IsNil(o.IsReplacementOffer) {
		var ret bool
		return ret
	}
	return *o.IsReplacementOffer
}

// GetIsReplacementOfferOk returns a tuple with the IsReplacementOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetIsReplacementOfferOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReplacementOffer) {
		return nil, false
	}
	return o.IsReplacementOffer, true
}

// HasIsReplacementOffer returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasIsReplacementOffer() bool {
	if o != nil && !IsNil(o.IsReplacementOffer) {
		return true
	}

	return false
}

// SetIsReplacementOffer gets a reference to the given bool and assigns it to the IsReplacementOffer field.
func (o *WorkloadMetaInfo) SetIsReplacementOffer(v bool) {
	o.IsReplacementOffer = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetLastModifiedBy() LastModifiedBy {
	if o == nil || IsNil(o.LastModifiedBy) {
		var ret LastModifiedBy
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetLastModifiedByOk() (*LastModifiedBy, bool) {
	if o == nil || IsNil(o.LastModifiedBy) {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasLastModifiedBy() bool {
	if o != nil && !IsNil(o.LastModifiedBy) {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given LastModifiedBy and assigns it to the LastModifiedBy field.
func (o *WorkloadMetaInfo) SetLastModifiedBy(v LastModifiedBy) {
	o.LastModifiedBy = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetNotifications() []NotificationEvent {
	if o == nil || IsNil(o.Notifications) {
		var ret []NotificationEvent
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetNotificationsOk() ([]NotificationEvent, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []NotificationEvent and assigns it to the Notifications field.
func (o *WorkloadMetaInfo) SetNotifications(v []NotificationEvent) {
	o.Notifications = v
}

// GetOfferAcceptDate returns the OfferAcceptDate field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetOfferAcceptDate() time.Time {
	if o == nil || IsNil(o.OfferAcceptDate) {
		var ret time.Time
		return ret
	}
	return *o.OfferAcceptDate
}

// GetOfferAcceptDateOk returns a tuple with the OfferAcceptDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetOfferAcceptDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OfferAcceptDate) {
		return nil, false
	}
	return o.OfferAcceptDate, true
}

// HasOfferAcceptDate returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasOfferAcceptDate() bool {
	if o != nil && !IsNil(o.OfferAcceptDate) {
		return true
	}

	return false
}

// SetOfferAcceptDate gets a reference to the given time.Time and assigns it to the OfferAcceptDate field.
func (o *WorkloadMetaInfo) SetOfferAcceptDate(v time.Time) {
	o.OfferAcceptDate = &v
}

// GetRenewalOfferType returns the RenewalOfferType field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetRenewalOfferType() AwsRenewalOfferType {
	if o == nil || IsNil(o.RenewalOfferType) {
		var ret AwsRenewalOfferType
		return ret
	}
	return *o.RenewalOfferType
}

// GetRenewalOfferTypeOk returns a tuple with the RenewalOfferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetRenewalOfferTypeOk() (*AwsRenewalOfferType, bool) {
	if o == nil || IsNil(o.RenewalOfferType) {
		return nil, false
	}
	return o.RenewalOfferType, true
}

// HasRenewalOfferType returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasRenewalOfferType() bool {
	if o != nil && !IsNil(o.RenewalOfferType) {
		return true
	}

	return false
}

// SetRenewalOfferType gets a reference to the given AwsRenewalOfferType and assigns it to the RenewalOfferType field.
func (o *WorkloadMetaInfo) SetRenewalOfferType(v AwsRenewalOfferType) {
	o.RenewalOfferType = &v
}

// GetReplacedOfferEndTime returns the ReplacedOfferEndTime field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetReplacedOfferEndTime() time.Time {
	if o == nil || IsNil(o.ReplacedOfferEndTime) {
		var ret time.Time
		return ret
	}
	return *o.ReplacedOfferEndTime
}

// GetReplacedOfferEndTimeOk returns a tuple with the ReplacedOfferEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetReplacedOfferEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReplacedOfferEndTime) {
		return nil, false
	}
	return o.ReplacedOfferEndTime, true
}

// HasReplacedOfferEndTime returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasReplacedOfferEndTime() bool {
	if o != nil && !IsNil(o.ReplacedOfferEndTime) {
		return true
	}

	return false
}

// SetReplacedOfferEndTime gets a reference to the given time.Time and assigns it to the ReplacedOfferEndTime field.
func (o *WorkloadMetaInfo) SetReplacedOfferEndTime(v time.Time) {
	o.ReplacedOfferEndTime = &v
}

// GetReplacedOfferResourceName returns the ReplacedOfferResourceName field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetReplacedOfferResourceName() string {
	if o == nil || IsNil(o.ReplacedOfferResourceName) {
		var ret string
		return ret
	}
	return *o.ReplacedOfferResourceName
}

// GetReplacedOfferResourceNameOk returns a tuple with the ReplacedOfferResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetReplacedOfferResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReplacedOfferResourceName) {
		return nil, false
	}
	return o.ReplacedOfferResourceName, true
}

// HasReplacedOfferResourceName returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasReplacedOfferResourceName() bool {
	if o != nil && !IsNil(o.ReplacedOfferResourceName) {
		return true
	}

	return false
}

// SetReplacedOfferResourceName gets a reference to the given string and assigns it to the ReplacedOfferResourceName field.
func (o *WorkloadMetaInfo) SetReplacedOfferResourceName(v string) {
	o.ReplacedOfferResourceName = &v
}

// GetSalesforceEntitlementURL returns the SalesforceEntitlementURL field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetSalesforceEntitlementURL() string {
	if o == nil || IsNil(o.SalesforceEntitlementURL) {
		var ret string
		return ret
	}
	return *o.SalesforceEntitlementURL
}

// GetSalesforceEntitlementURLOk returns a tuple with the SalesforceEntitlementURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetSalesforceEntitlementURLOk() (*string, bool) {
	if o == nil || IsNil(o.SalesforceEntitlementURL) {
		return nil, false
	}
	return o.SalesforceEntitlementURL, true
}

// HasSalesforceEntitlementURL returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasSalesforceEntitlementURL() bool {
	if o != nil && !IsNil(o.SalesforceEntitlementURL) {
		return true
	}

	return false
}

// SetSalesforceEntitlementURL gets a reference to the given string and assigns it to the SalesforceEntitlementURL field.
func (o *WorkloadMetaInfo) SetSalesforceEntitlementURL(v string) {
	o.SalesforceEntitlementURL = &v
}

// GetSalesforceOpportunityId returns the SalesforceOpportunityId field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetSalesforceOpportunityId() string {
	if o == nil || IsNil(o.SalesforceOpportunityId) {
		var ret string
		return ret
	}
	return *o.SalesforceOpportunityId
}

// GetSalesforceOpportunityIdOk returns a tuple with the SalesforceOpportunityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetSalesforceOpportunityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesforceOpportunityId) {
		return nil, false
	}
	return o.SalesforceOpportunityId, true
}

// HasSalesforceOpportunityId returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasSalesforceOpportunityId() bool {
	if o != nil && !IsNil(o.SalesforceOpportunityId) {
		return true
	}

	return false
}

// SetSalesforceOpportunityId gets a reference to the given string and assigns it to the SalesforceOpportunityId field.
func (o *WorkloadMetaInfo) SetSalesforceOpportunityId(v string) {
	o.SalesforceOpportunityId = &v
}

// GetTestUsageMeteringEndTime returns the TestUsageMeteringEndTime field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetTestUsageMeteringEndTime() time.Time {
	if o == nil || IsNil(o.TestUsageMeteringEndTime) {
		var ret time.Time
		return ret
	}
	return *o.TestUsageMeteringEndTime
}

// GetTestUsageMeteringEndTimeOk returns a tuple with the TestUsageMeteringEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetTestUsageMeteringEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TestUsageMeteringEndTime) {
		return nil, false
	}
	return o.TestUsageMeteringEndTime, true
}

// HasTestUsageMeteringEndTime returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasTestUsageMeteringEndTime() bool {
	if o != nil && !IsNil(o.TestUsageMeteringEndTime) {
		return true
	}

	return false
}

// SetTestUsageMeteringEndTime gets a reference to the given time.Time and assigns it to the TestUsageMeteringEndTime field.
func (o *WorkloadMetaInfo) SetTestUsageMeteringEndTime(v time.Time) {
	o.TestUsageMeteringEndTime = &v
}

// GetUpdateMessage returns the UpdateMessage field value if set, zero value otherwise.
func (o *WorkloadMetaInfo) GetUpdateMessage() string {
	if o == nil || IsNil(o.UpdateMessage) {
		var ret string
		return ret
	}
	return *o.UpdateMessage
}

// GetUpdateMessageOk returns a tuple with the UpdateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkloadMetaInfo) GetUpdateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateMessage) {
		return nil, false
	}
	return o.UpdateMessage, true
}

// HasUpdateMessage returns a boolean if a field has been set.
func (o *WorkloadMetaInfo) HasUpdateMessage() bool {
	if o != nil && !IsNil(o.UpdateMessage) {
		return true
	}

	return false
}

// SetUpdateMessage gets a reference to the given string and assigns it to the UpdateMessage field.
func (o *WorkloadMetaInfo) SetUpdateMessage(v string) {
	o.UpdateMessage = &v
}

func (o WorkloadMetaInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkloadMetaInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AceApnCrmUniqueIdentifier) {
		toSerialize["aceApnCrmUniqueIdentifier"] = o.AceApnCrmUniqueIdentifier
	}
	if !IsNil(o.AwsSaasProductDimensions) {
		toSerialize["awsSaasProductDimensions"] = o.AwsSaasProductDimensions
	}
	if !IsNil(o.BaseAgreementId) {
		toSerialize["baseAgreementId"] = o.BaseAgreementId
	}
	if !IsNil(o.BuyerIds) {
		toSerialize["buyerIds"] = o.BuyerIds
	}
	if !IsNil(o.Contacts) {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.CppoInOfferId) {
		toSerialize["cppoInOfferId"] = o.CppoInOfferId
	}
	if !IsNil(o.CppoOfferId) {
		toSerialize["cppoOfferId"] = o.CppoOfferId
	}
	if !IsNil(o.CppoOutOfferId) {
		toSerialize["cppoOutOfferId"] = o.CppoOutOfferId
	}
	if !IsNil(o.CustomMetaInfo) {
		toSerialize["customMetaInfo"] = o.CustomMetaInfo
	}
	if !IsNil(o.EnableTestUsageMetering) {
		toSerialize["enableTestUsageMetering"] = o.EnableTestUsageMetering
	}
	if !IsNil(o.EntitlementCancellationSchedule) {
		toSerialize["entitlementCancellationSchedule"] = o.EntitlementCancellationSchedule
	}
	if !IsNil(o.ErrorMessages) {
		toSerialize["errorMessages"] = o.ErrorMessages
	}
	if !IsNil(o.HubspotDealId) {
		toSerialize["hubspotDealId"] = o.HubspotDealId
	}
	if !IsNil(o.InternalNote) {
		toSerialize["internalNote"] = o.InternalNote
	}
	if !IsNil(o.IsAgreementBasedOffer) {
		toSerialize["isAgreementBasedOffer"] = o.IsAgreementBasedOffer
	}
	if !IsNil(o.IsGrossRevenueFullSync) {
		toSerialize["isGrossRevenueFullSync"] = o.IsGrossRevenueFullSync
	}
	if !IsNil(o.IsRenewalOffer) {
		toSerialize["isRenewalOffer"] = o.IsRenewalOffer
	}
	if !IsNil(o.IsReplacementOffer) {
		toSerialize["isReplacementOffer"] = o.IsReplacementOffer
	}
	if !IsNil(o.LastModifiedBy) {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.OfferAcceptDate) {
		toSerialize["offerAcceptDate"] = o.OfferAcceptDate
	}
	if !IsNil(o.RenewalOfferType) {
		toSerialize["renewalOfferType"] = o.RenewalOfferType
	}
	if !IsNil(o.ReplacedOfferEndTime) {
		toSerialize["replacedOfferEndTime"] = o.ReplacedOfferEndTime
	}
	if !IsNil(o.ReplacedOfferResourceName) {
		toSerialize["replacedOfferResourceName"] = o.ReplacedOfferResourceName
	}
	if !IsNil(o.SalesforceEntitlementURL) {
		toSerialize["salesforceEntitlementURL"] = o.SalesforceEntitlementURL
	}
	if !IsNil(o.SalesforceOpportunityId) {
		toSerialize["salesforceOpportunityId"] = o.SalesforceOpportunityId
	}
	if !IsNil(o.TestUsageMeteringEndTime) {
		toSerialize["testUsageMeteringEndTime"] = o.TestUsageMeteringEndTime
	}
	if !IsNil(o.UpdateMessage) {
		toSerialize["updateMessage"] = o.UpdateMessage
	}
	return toSerialize, nil
}

type NullableWorkloadMetaInfo struct {
	value *WorkloadMetaInfo
	isSet bool
}

func (v NullableWorkloadMetaInfo) Get() *WorkloadMetaInfo {
	return v.value
}

func (v *NullableWorkloadMetaInfo) Set(val *WorkloadMetaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkloadMetaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkloadMetaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkloadMetaInfo(val *WorkloadMetaInfo) *NullableWorkloadMetaInfo {
	return &NullableWorkloadMetaInfo{value: val, isSet: true}
}

func (v NullableWorkloadMetaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkloadMetaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
