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
	"time"
)

// checks if the AzureMarketplacePrivateOffer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePrivateOffer{}

// AzureMarketplacePrivateOffer struct for AzureMarketplacePrivateOffer
type AzureMarketplacePrivateOffer struct {
	Schema *string `json:"$schema,omitempty"`
	// in format YYYY-MM-DD
	AcceptBy *time.Time `json:"acceptBy,omitempty"`
	AcceptanceLinks []AzureMarketplacePrivateOfferAcceptanceLink `json:"acceptanceLinks,omitempty"`
	Beneficiaries []AzureMarketplacePrivateOfferBeneficiary `json:"beneficiaries,omitempty"`
	ETag *string `json:"eTag,omitempty"`
	// in format YYYY-MM-DD
	End *time.Time `json:"end,omitempty"`
	// in format of \"private-offer/private-offer-durable-id\"
	Id *string `json:"id,omitempty"`
	// in format YYYY-MM-DD
	LastModified *time.Time `json:"lastModified,omitempty"`
	Name *string `json:"name,omitempty"`
	// array of email addresses of the users to be notified of any changes in the private offer status.
	NotificationContacts []string `json:"notificationContacts,omitempty"`
	OfferPricingType *AzureMarketplaceOfferPricingType `json:"offerPricingType,omitempty"`
	Partners []AzureMarketplacePrivateOfferPartner `json:"partners,omitempty"`
	PreparedBy *string `json:"preparedBy,omitempty"`
	// Up to 10 pricing entries are allowed.
	Pricing []AzureMarketplacePrivateOfferPricing `json:"pricing,omitempty"`
	PrivateOfferType *AzureMarketplacePrivateOfferType `json:"privateOfferType,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	// in format YYYY-MM-DD, if VariableStartDate = true, this field should be empty.
	Start *time.Time `json:"start,omitempty"`
	State *AzureMarketplacePrivateOfferState `json:"state,omitempty"`
	SubState *AzureMarketplacePrivateOfferSubState `json:"subState,omitempty"`
	// Only applicable to private offers with privateOfferType = customerPromotion || cspPromotion
	TermsAndConditionsDocSasUrl *string `json:"termsAndConditionsDocSasUrl,omitempty"`
	// Only applicable to private offers with privateOfferType = multipartyPromotionOriginator || multipartyPromotionChannelPartner
	TermsAndConditionsDocs []AzureMarketplacePrivateOfferTermsDoc `json:"termsAndConditionsDocs,omitempty"`
	UpgradedFrom *AzureMarketplacePrivateOfferPromotionReference `json:"upgradedFrom,omitempty"`
	Validations []AzureMarketplaceValidation `json:"validations,omitempty"`
	VariableStartDate *bool `json:"variableStartDate,omitempty"`
}

// NewAzureMarketplacePrivateOffer instantiates a new AzureMarketplacePrivateOffer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePrivateOffer() *AzureMarketplacePrivateOffer {
	this := AzureMarketplacePrivateOffer{}
	return &this
}

// NewAzureMarketplacePrivateOfferWithDefaults instantiates a new AzureMarketplacePrivateOffer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePrivateOfferWithDefaults() *AzureMarketplacePrivateOffer {
	this := AzureMarketplacePrivateOffer{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *AzureMarketplacePrivateOffer) SetSchema(v string) {
	o.Schema = &v
}

// GetAcceptBy returns the AcceptBy field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetAcceptBy() time.Time {
	if o == nil || IsNil(o.AcceptBy) {
		var ret time.Time
		return ret
	}
	return *o.AcceptBy
}

// GetAcceptByOk returns a tuple with the AcceptBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetAcceptByOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AcceptBy) {
		return nil, false
	}
	return o.AcceptBy, true
}

// HasAcceptBy returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasAcceptBy() bool {
	if o != nil && !IsNil(o.AcceptBy) {
		return true
	}

	return false
}

// SetAcceptBy gets a reference to the given time.Time and assigns it to the AcceptBy field.
func (o *AzureMarketplacePrivateOffer) SetAcceptBy(v time.Time) {
	o.AcceptBy = &v
}

// GetAcceptanceLinks returns the AcceptanceLinks field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetAcceptanceLinks() []AzureMarketplacePrivateOfferAcceptanceLink {
	if o == nil || IsNil(o.AcceptanceLinks) {
		var ret []AzureMarketplacePrivateOfferAcceptanceLink
		return ret
	}
	return o.AcceptanceLinks
}

// GetAcceptanceLinksOk returns a tuple with the AcceptanceLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetAcceptanceLinksOk() ([]AzureMarketplacePrivateOfferAcceptanceLink, bool) {
	if o == nil || IsNil(o.AcceptanceLinks) {
		return nil, false
	}
	return o.AcceptanceLinks, true
}

// HasAcceptanceLinks returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasAcceptanceLinks() bool {
	if o != nil && !IsNil(o.AcceptanceLinks) {
		return true
	}

	return false
}

// SetAcceptanceLinks gets a reference to the given []AzureMarketplacePrivateOfferAcceptanceLink and assigns it to the AcceptanceLinks field.
func (o *AzureMarketplacePrivateOffer) SetAcceptanceLinks(v []AzureMarketplacePrivateOfferAcceptanceLink) {
	o.AcceptanceLinks = v
}

// GetBeneficiaries returns the Beneficiaries field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetBeneficiaries() []AzureMarketplacePrivateOfferBeneficiary {
	if o == nil || IsNil(o.Beneficiaries) {
		var ret []AzureMarketplacePrivateOfferBeneficiary
		return ret
	}
	return o.Beneficiaries
}

// GetBeneficiariesOk returns a tuple with the Beneficiaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetBeneficiariesOk() ([]AzureMarketplacePrivateOfferBeneficiary, bool) {
	if o == nil || IsNil(o.Beneficiaries) {
		return nil, false
	}
	return o.Beneficiaries, true
}

// HasBeneficiaries returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasBeneficiaries() bool {
	if o != nil && !IsNil(o.Beneficiaries) {
		return true
	}

	return false
}

// SetBeneficiaries gets a reference to the given []AzureMarketplacePrivateOfferBeneficiary and assigns it to the Beneficiaries field.
func (o *AzureMarketplacePrivateOffer) SetBeneficiaries(v []AzureMarketplacePrivateOfferBeneficiary) {
	o.Beneficiaries = v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *AzureMarketplacePrivateOffer) SetETag(v string) {
	o.ETag = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *AzureMarketplacePrivateOffer) SetEnd(v time.Time) {
	o.End = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplacePrivateOffer) SetId(v string) {
	o.Id = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *AzureMarketplacePrivateOffer) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AzureMarketplacePrivateOffer) SetName(v string) {
	o.Name = &v
}

// GetNotificationContacts returns the NotificationContacts field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetNotificationContacts() []string {
	if o == nil || IsNil(o.NotificationContacts) {
		var ret []string
		return ret
	}
	return o.NotificationContacts
}

// GetNotificationContactsOk returns a tuple with the NotificationContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetNotificationContactsOk() ([]string, bool) {
	if o == nil || IsNil(o.NotificationContacts) {
		return nil, false
	}
	return o.NotificationContacts, true
}

// HasNotificationContacts returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasNotificationContacts() bool {
	if o != nil && !IsNil(o.NotificationContacts) {
		return true
	}

	return false
}

// SetNotificationContacts gets a reference to the given []string and assigns it to the NotificationContacts field.
func (o *AzureMarketplacePrivateOffer) SetNotificationContacts(v []string) {
	o.NotificationContacts = v
}

// GetOfferPricingType returns the OfferPricingType field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetOfferPricingType() AzureMarketplaceOfferPricingType {
	if o == nil || IsNil(o.OfferPricingType) {
		var ret AzureMarketplaceOfferPricingType
		return ret
	}
	return *o.OfferPricingType
}

// GetOfferPricingTypeOk returns a tuple with the OfferPricingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetOfferPricingTypeOk() (*AzureMarketplaceOfferPricingType, bool) {
	if o == nil || IsNil(o.OfferPricingType) {
		return nil, false
	}
	return o.OfferPricingType, true
}

// HasOfferPricingType returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasOfferPricingType() bool {
	if o != nil && !IsNil(o.OfferPricingType) {
		return true
	}

	return false
}

// SetOfferPricingType gets a reference to the given AzureMarketplaceOfferPricingType and assigns it to the OfferPricingType field.
func (o *AzureMarketplacePrivateOffer) SetOfferPricingType(v AzureMarketplaceOfferPricingType) {
	o.OfferPricingType = &v
}

// GetPartners returns the Partners field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetPartners() []AzureMarketplacePrivateOfferPartner {
	if o == nil || IsNil(o.Partners) {
		var ret []AzureMarketplacePrivateOfferPartner
		return ret
	}
	return o.Partners
}

// GetPartnersOk returns a tuple with the Partners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetPartnersOk() ([]AzureMarketplacePrivateOfferPartner, bool) {
	if o == nil || IsNil(o.Partners) {
		return nil, false
	}
	return o.Partners, true
}

// HasPartners returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasPartners() bool {
	if o != nil && !IsNil(o.Partners) {
		return true
	}

	return false
}

// SetPartners gets a reference to the given []AzureMarketplacePrivateOfferPartner and assigns it to the Partners field.
func (o *AzureMarketplacePrivateOffer) SetPartners(v []AzureMarketplacePrivateOfferPartner) {
	o.Partners = v
}

// GetPreparedBy returns the PreparedBy field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetPreparedBy() string {
	if o == nil || IsNil(o.PreparedBy) {
		var ret string
		return ret
	}
	return *o.PreparedBy
}

// GetPreparedByOk returns a tuple with the PreparedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetPreparedByOk() (*string, bool) {
	if o == nil || IsNil(o.PreparedBy) {
		return nil, false
	}
	return o.PreparedBy, true
}

// HasPreparedBy returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasPreparedBy() bool {
	if o != nil && !IsNil(o.PreparedBy) {
		return true
	}

	return false
}

// SetPreparedBy gets a reference to the given string and assigns it to the PreparedBy field.
func (o *AzureMarketplacePrivateOffer) SetPreparedBy(v string) {
	o.PreparedBy = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetPricing() []AzureMarketplacePrivateOfferPricing {
	if o == nil || IsNil(o.Pricing) {
		var ret []AzureMarketplacePrivateOfferPricing
		return ret
	}
	return o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetPricingOk() ([]AzureMarketplacePrivateOfferPricing, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasPricing() bool {
	if o != nil && !IsNil(o.Pricing) {
		return true
	}

	return false
}

// SetPricing gets a reference to the given []AzureMarketplacePrivateOfferPricing and assigns it to the Pricing field.
func (o *AzureMarketplacePrivateOffer) SetPricing(v []AzureMarketplacePrivateOfferPricing) {
	o.Pricing = v
}

// GetPrivateOfferType returns the PrivateOfferType field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetPrivateOfferType() AzureMarketplacePrivateOfferType {
	if o == nil || IsNil(o.PrivateOfferType) {
		var ret AzureMarketplacePrivateOfferType
		return ret
	}
	return *o.PrivateOfferType
}

// GetPrivateOfferTypeOk returns a tuple with the PrivateOfferType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetPrivateOfferTypeOk() (*AzureMarketplacePrivateOfferType, bool) {
	if o == nil || IsNil(o.PrivateOfferType) {
		return nil, false
	}
	return o.PrivateOfferType, true
}

// HasPrivateOfferType returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasPrivateOfferType() bool {
	if o != nil && !IsNil(o.PrivateOfferType) {
		return true
	}

	return false
}

// SetPrivateOfferType gets a reference to the given AzureMarketplacePrivateOfferType and assigns it to the PrivateOfferType field.
func (o *AzureMarketplacePrivateOffer) SetPrivateOfferType(v AzureMarketplacePrivateOfferType) {
	o.PrivateOfferType = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *AzureMarketplacePrivateOffer) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *AzureMarketplacePrivateOffer) SetStart(v time.Time) {
	o.Start = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetState() AzureMarketplacePrivateOfferState {
	if o == nil || IsNil(o.State) {
		var ret AzureMarketplacePrivateOfferState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetStateOk() (*AzureMarketplacePrivateOfferState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given AzureMarketplacePrivateOfferState and assigns it to the State field.
func (o *AzureMarketplacePrivateOffer) SetState(v AzureMarketplacePrivateOfferState) {
	o.State = &v
}

// GetSubState returns the SubState field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetSubState() AzureMarketplacePrivateOfferSubState {
	if o == nil || IsNil(o.SubState) {
		var ret AzureMarketplacePrivateOfferSubState
		return ret
	}
	return *o.SubState
}

// GetSubStateOk returns a tuple with the SubState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetSubStateOk() (*AzureMarketplacePrivateOfferSubState, bool) {
	if o == nil || IsNil(o.SubState) {
		return nil, false
	}
	return o.SubState, true
}

// HasSubState returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasSubState() bool {
	if o != nil && !IsNil(o.SubState) {
		return true
	}

	return false
}

// SetSubState gets a reference to the given AzureMarketplacePrivateOfferSubState and assigns it to the SubState field.
func (o *AzureMarketplacePrivateOffer) SetSubState(v AzureMarketplacePrivateOfferSubState) {
	o.SubState = &v
}

// GetTermsAndConditionsDocSasUrl returns the TermsAndConditionsDocSasUrl field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocSasUrl() string {
	if o == nil || IsNil(o.TermsAndConditionsDocSasUrl) {
		var ret string
		return ret
	}
	return *o.TermsAndConditionsDocSasUrl
}

// GetTermsAndConditionsDocSasUrlOk returns a tuple with the TermsAndConditionsDocSasUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocSasUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TermsAndConditionsDocSasUrl) {
		return nil, false
	}
	return o.TermsAndConditionsDocSasUrl, true
}

// HasTermsAndConditionsDocSasUrl returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasTermsAndConditionsDocSasUrl() bool {
	if o != nil && !IsNil(o.TermsAndConditionsDocSasUrl) {
		return true
	}

	return false
}

// SetTermsAndConditionsDocSasUrl gets a reference to the given string and assigns it to the TermsAndConditionsDocSasUrl field.
func (o *AzureMarketplacePrivateOffer) SetTermsAndConditionsDocSasUrl(v string) {
	o.TermsAndConditionsDocSasUrl = &v
}

// GetTermsAndConditionsDocs returns the TermsAndConditionsDocs field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocs() []AzureMarketplacePrivateOfferTermsDoc {
	if o == nil || IsNil(o.TermsAndConditionsDocs) {
		var ret []AzureMarketplacePrivateOfferTermsDoc
		return ret
	}
	return o.TermsAndConditionsDocs
}

// GetTermsAndConditionsDocsOk returns a tuple with the TermsAndConditionsDocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetTermsAndConditionsDocsOk() ([]AzureMarketplacePrivateOfferTermsDoc, bool) {
	if o == nil || IsNil(o.TermsAndConditionsDocs) {
		return nil, false
	}
	return o.TermsAndConditionsDocs, true
}

// HasTermsAndConditionsDocs returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasTermsAndConditionsDocs() bool {
	if o != nil && !IsNil(o.TermsAndConditionsDocs) {
		return true
	}

	return false
}

// SetTermsAndConditionsDocs gets a reference to the given []AzureMarketplacePrivateOfferTermsDoc and assigns it to the TermsAndConditionsDocs field.
func (o *AzureMarketplacePrivateOffer) SetTermsAndConditionsDocs(v []AzureMarketplacePrivateOfferTermsDoc) {
	o.TermsAndConditionsDocs = v
}

// GetUpgradedFrom returns the UpgradedFrom field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetUpgradedFrom() AzureMarketplacePrivateOfferPromotionReference {
	if o == nil || IsNil(o.UpgradedFrom) {
		var ret AzureMarketplacePrivateOfferPromotionReference
		return ret
	}
	return *o.UpgradedFrom
}

// GetUpgradedFromOk returns a tuple with the UpgradedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetUpgradedFromOk() (*AzureMarketplacePrivateOfferPromotionReference, bool) {
	if o == nil || IsNil(o.UpgradedFrom) {
		return nil, false
	}
	return o.UpgradedFrom, true
}

// HasUpgradedFrom returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasUpgradedFrom() bool {
	if o != nil && !IsNil(o.UpgradedFrom) {
		return true
	}

	return false
}

// SetUpgradedFrom gets a reference to the given AzureMarketplacePrivateOfferPromotionReference and assigns it to the UpgradedFrom field.
func (o *AzureMarketplacePrivateOffer) SetUpgradedFrom(v AzureMarketplacePrivateOfferPromotionReference) {
	o.UpgradedFrom = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetValidations() []AzureMarketplaceValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []AzureMarketplaceValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetValidationsOk() ([]AzureMarketplaceValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []AzureMarketplaceValidation and assigns it to the Validations field.
func (o *AzureMarketplacePrivateOffer) SetValidations(v []AzureMarketplaceValidation) {
	o.Validations = v
}

// GetVariableStartDate returns the VariableStartDate field value if set, zero value otherwise.
func (o *AzureMarketplacePrivateOffer) GetVariableStartDate() bool {
	if o == nil || IsNil(o.VariableStartDate) {
		var ret bool
		return ret
	}
	return *o.VariableStartDate
}

// GetVariableStartDateOk returns a tuple with the VariableStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePrivateOffer) GetVariableStartDateOk() (*bool, bool) {
	if o == nil || IsNil(o.VariableStartDate) {
		return nil, false
	}
	return o.VariableStartDate, true
}

// HasVariableStartDate returns a boolean if a field has been set.
func (o *AzureMarketplacePrivateOffer) HasVariableStartDate() bool {
	if o != nil && !IsNil(o.VariableStartDate) {
		return true
	}

	return false
}

// SetVariableStartDate gets a reference to the given bool and assigns it to the VariableStartDate field.
func (o *AzureMarketplacePrivateOffer) SetVariableStartDate(v bool) {
	o.VariableStartDate = &v
}

func (o AzureMarketplacePrivateOffer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePrivateOffer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.AcceptBy) {
		toSerialize["acceptBy"] = o.AcceptBy
	}
	if !IsNil(o.AcceptanceLinks) {
		toSerialize["acceptanceLinks"] = o.AcceptanceLinks
	}
	if !IsNil(o.Beneficiaries) {
		toSerialize["beneficiaries"] = o.Beneficiaries
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NotificationContacts) {
		toSerialize["notificationContacts"] = o.NotificationContacts
	}
	if !IsNil(o.OfferPricingType) {
		toSerialize["offerPricingType"] = o.OfferPricingType
	}
	if !IsNil(o.Partners) {
		toSerialize["partners"] = o.Partners
	}
	if !IsNil(o.PreparedBy) {
		toSerialize["preparedBy"] = o.PreparedBy
	}
	if !IsNil(o.Pricing) {
		toSerialize["pricing"] = o.Pricing
	}
	if !IsNil(o.PrivateOfferType) {
		toSerialize["privateOfferType"] = o.PrivateOfferType
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.SubState) {
		toSerialize["subState"] = o.SubState
	}
	if !IsNil(o.TermsAndConditionsDocSasUrl) {
		toSerialize["termsAndConditionsDocSasUrl"] = o.TermsAndConditionsDocSasUrl
	}
	if !IsNil(o.TermsAndConditionsDocs) {
		toSerialize["termsAndConditionsDocs"] = o.TermsAndConditionsDocs
	}
	if !IsNil(o.UpgradedFrom) {
		toSerialize["upgradedFrom"] = o.UpgradedFrom
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	if !IsNil(o.VariableStartDate) {
		toSerialize["variableStartDate"] = o.VariableStartDate
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePrivateOffer struct {
	value *AzureMarketplacePrivateOffer
	isSet bool
}

func (v NullableAzureMarketplacePrivateOffer) Get() *AzureMarketplacePrivateOffer {
	return v.value
}

func (v *NullableAzureMarketplacePrivateOffer) Set(val *AzureMarketplacePrivateOffer) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePrivateOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePrivateOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePrivateOffer(val *AzureMarketplacePrivateOffer) *NullableAzureMarketplacePrivateOffer {
	return &NullableAzureMarketplacePrivateOffer{value: val, isSet: true}
}

func (v NullableAzureMarketplacePrivateOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePrivateOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


