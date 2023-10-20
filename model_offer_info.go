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

// checks if the OfferInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfferInfo{}

// OfferInfo struct for OfferInfo
type OfferInfo struct {
	AttachEulaType *EulaType `json:"attachEulaType,omitempty"`
	// Is this offer Auto Renew enabled.
	AutoRenew *bool `json:"autoRenew,omitempty"`
	AwsCppoEventDetail *AwsMarketplaceEventBridgeEventDetail `json:"awsCppoEventDetail,omitempty"`
	AwsCppoOpportunity *AwsMarketplaceCppoOpportunity `json:"awsCppoOpportunity,omitempty"`
	AzureOriginalPlan *AzurePriceAndAvailabilityPrivateOfferPlan `json:"azureOriginalPlan,omitempty"`
	AzurePrivateOffer *AzureMarketplacePrivateOffer `json:"azurePrivateOffer,omitempty"`
	AzureProductVariant *AzureProductVariant `json:"azureProductVariant,omitempty"`
	// The buyers' AWS Account IDs of this offer.
	BuyerAwsAccountIds []string `json:"buyerAwsAccountIds,omitempty"`
	// The buyers' Azure tenants of this offer.
	BuyerAzureTenants []AzureAudience `json:"buyerAzureTenants,omitempty"`
	// The amount that the buyer has committed to pay, before discount if applicable. It can be monthly commitment or total commitment.
	CommitAmount *float32 `json:"commitAmount,omitempty"`
	Commits []CommitDimension `json:"commits,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Dimensions []MeteringDimension `json:"dimensions,omitempty"`
	// The discount percentage off the original price. For example, 20 means 20% off. 0 means no discount. It can be discount off the commitment amount or discount off the usage price.
	DiscountPercentage *float32 `json:"discountPercentage,omitempty"`
	EulaType *EulaType `json:"eulaType,omitempty"`
	EulaUrl *string `json:"eulaUrl,omitempty"`
	GcpCustomerInfo *GcpMarketplacePrivateOfferCustomerInfo `json:"gcpCustomerInfo,omitempty"`
	// The duration of the offer in months. Only required when creating GCP Marketplace private offer.
	GcpDuration *int32 `json:"gcpDuration,omitempty"`
	// Only applicable for GCP Marketplace Offers (the default or private offer)
	GcpMetrics []GcpMarketplaceProductMeteringMetric `json:"gcpMetrics,omitempty"`
	GcpPaymentSchedule *GcpMarketplacePaymentScheduleType `json:"gcpPaymentSchedule,omitempty"`
	// Only applicable for GCP Marketplace
	GcpPlans []GcpMarketplaceProductPurchaseOptionSpec `json:"gcpPlans,omitempty"`
	GcpPrivateOffer *GcpMarketplacePrivateOffer `json:"gcpPrivateOffer,omitempty"`
	GcpProviderInfo *GcpMarketplacePrivateOfferProviderInfo `json:"gcpProviderInfo,omitempty"`
	// Optional when creating GCP Marketplace private offer. The internal note for the seller/ISV. It is only visible to the seller/ISV.
	GcpProviderInternalNote *string `json:"gcpProviderInternalNote,omitempty"`
	// Optional when creating GCP Marketplace private offer. The public note for the buyer. It is visible to the buyer.
	GcpProviderPublicNote *string `json:"gcpProviderPublicNote,omitempty"`
	GcpUsagePlanPriceModel *GcpMarketplaceUsagePlanPriceModel `json:"gcpUsagePlanPriceModel,omitempty"`
	// For flexible payment schedule.
	PaymentInstallments []PaymentInstallment `json:"paymentInstallments,omitempty"`
	// The URL of the private offer sent to buyers to accept. Only applicable for private offer.
	PrivateOfferUrl *string `json:"privateOfferUrl,omitempty"`
	RefundCancelationPolicy *string `json:"refundCancelationPolicy,omitempty"`
	SellerNotes *string `json:"sellerNotes,omitempty"`
	// The default visibility of offer is PRIVATE.
	Visibility *string `json:"visibility,omitempty"`
}

// NewOfferInfo instantiates a new OfferInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferInfo() *OfferInfo {
	this := OfferInfo{}
	return &this
}

// NewOfferInfoWithDefaults instantiates a new OfferInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferInfoWithDefaults() *OfferInfo {
	this := OfferInfo{}
	return &this
}

// GetAttachEulaType returns the AttachEulaType field value if set, zero value otherwise.
func (o *OfferInfo) GetAttachEulaType() EulaType {
	if o == nil || IsNil(o.AttachEulaType) {
		var ret EulaType
		return ret
	}
	return *o.AttachEulaType
}

// GetAttachEulaTypeOk returns a tuple with the AttachEulaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAttachEulaTypeOk() (*EulaType, bool) {
	if o == nil || IsNil(o.AttachEulaType) {
		return nil, false
	}
	return o.AttachEulaType, true
}

// HasAttachEulaType returns a boolean if a field has been set.
func (o *OfferInfo) HasAttachEulaType() bool {
	if o != nil && !IsNil(o.AttachEulaType) {
		return true
	}

	return false
}

// SetAttachEulaType gets a reference to the given EulaType and assigns it to the AttachEulaType field.
func (o *OfferInfo) SetAttachEulaType(v EulaType) {
	o.AttachEulaType = &v
}

// GetAutoRenew returns the AutoRenew field value if set, zero value otherwise.
func (o *OfferInfo) GetAutoRenew() bool {
	if o == nil || IsNil(o.AutoRenew) {
		var ret bool
		return ret
	}
	return *o.AutoRenew
}

// GetAutoRenewOk returns a tuple with the AutoRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAutoRenewOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoRenew) {
		return nil, false
	}
	return o.AutoRenew, true
}

// HasAutoRenew returns a boolean if a field has been set.
func (o *OfferInfo) HasAutoRenew() bool {
	if o != nil && !IsNil(o.AutoRenew) {
		return true
	}

	return false
}

// SetAutoRenew gets a reference to the given bool and assigns it to the AutoRenew field.
func (o *OfferInfo) SetAutoRenew(v bool) {
	o.AutoRenew = &v
}

// GetAwsCppoEventDetail returns the AwsCppoEventDetail field value if set, zero value otherwise.
func (o *OfferInfo) GetAwsCppoEventDetail() AwsMarketplaceEventBridgeEventDetail {
	if o == nil || IsNil(o.AwsCppoEventDetail) {
		var ret AwsMarketplaceEventBridgeEventDetail
		return ret
	}
	return *o.AwsCppoEventDetail
}

// GetAwsCppoEventDetailOk returns a tuple with the AwsCppoEventDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAwsCppoEventDetailOk() (*AwsMarketplaceEventBridgeEventDetail, bool) {
	if o == nil || IsNil(o.AwsCppoEventDetail) {
		return nil, false
	}
	return o.AwsCppoEventDetail, true
}

// HasAwsCppoEventDetail returns a boolean if a field has been set.
func (o *OfferInfo) HasAwsCppoEventDetail() bool {
	if o != nil && !IsNil(o.AwsCppoEventDetail) {
		return true
	}

	return false
}

// SetAwsCppoEventDetail gets a reference to the given AwsMarketplaceEventBridgeEventDetail and assigns it to the AwsCppoEventDetail field.
func (o *OfferInfo) SetAwsCppoEventDetail(v AwsMarketplaceEventBridgeEventDetail) {
	o.AwsCppoEventDetail = &v
}

// GetAwsCppoOpportunity returns the AwsCppoOpportunity field value if set, zero value otherwise.
func (o *OfferInfo) GetAwsCppoOpportunity() AwsMarketplaceCppoOpportunity {
	if o == nil || IsNil(o.AwsCppoOpportunity) {
		var ret AwsMarketplaceCppoOpportunity
		return ret
	}
	return *o.AwsCppoOpportunity
}

// GetAwsCppoOpportunityOk returns a tuple with the AwsCppoOpportunity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAwsCppoOpportunityOk() (*AwsMarketplaceCppoOpportunity, bool) {
	if o == nil || IsNil(o.AwsCppoOpportunity) {
		return nil, false
	}
	return o.AwsCppoOpportunity, true
}

// HasAwsCppoOpportunity returns a boolean if a field has been set.
func (o *OfferInfo) HasAwsCppoOpportunity() bool {
	if o != nil && !IsNil(o.AwsCppoOpportunity) {
		return true
	}

	return false
}

// SetAwsCppoOpportunity gets a reference to the given AwsMarketplaceCppoOpportunity and assigns it to the AwsCppoOpportunity field.
func (o *OfferInfo) SetAwsCppoOpportunity(v AwsMarketplaceCppoOpportunity) {
	o.AwsCppoOpportunity = &v
}

// GetAzureOriginalPlan returns the AzureOriginalPlan field value if set, zero value otherwise.
func (o *OfferInfo) GetAzureOriginalPlan() AzurePriceAndAvailabilityPrivateOfferPlan {
	if o == nil || IsNil(o.AzureOriginalPlan) {
		var ret AzurePriceAndAvailabilityPrivateOfferPlan
		return ret
	}
	return *o.AzureOriginalPlan
}

// GetAzureOriginalPlanOk returns a tuple with the AzureOriginalPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAzureOriginalPlanOk() (*AzurePriceAndAvailabilityPrivateOfferPlan, bool) {
	if o == nil || IsNil(o.AzureOriginalPlan) {
		return nil, false
	}
	return o.AzureOriginalPlan, true
}

// HasAzureOriginalPlan returns a boolean if a field has been set.
func (o *OfferInfo) HasAzureOriginalPlan() bool {
	if o != nil && !IsNil(o.AzureOriginalPlan) {
		return true
	}

	return false
}

// SetAzureOriginalPlan gets a reference to the given AzurePriceAndAvailabilityPrivateOfferPlan and assigns it to the AzureOriginalPlan field.
func (o *OfferInfo) SetAzureOriginalPlan(v AzurePriceAndAvailabilityPrivateOfferPlan) {
	o.AzureOriginalPlan = &v
}

// GetAzurePrivateOffer returns the AzurePrivateOffer field value if set, zero value otherwise.
func (o *OfferInfo) GetAzurePrivateOffer() AzureMarketplacePrivateOffer {
	if o == nil || IsNil(o.AzurePrivateOffer) {
		var ret AzureMarketplacePrivateOffer
		return ret
	}
	return *o.AzurePrivateOffer
}

// GetAzurePrivateOfferOk returns a tuple with the AzurePrivateOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAzurePrivateOfferOk() (*AzureMarketplacePrivateOffer, bool) {
	if o == nil || IsNil(o.AzurePrivateOffer) {
		return nil, false
	}
	return o.AzurePrivateOffer, true
}

// HasAzurePrivateOffer returns a boolean if a field has been set.
func (o *OfferInfo) HasAzurePrivateOffer() bool {
	if o != nil && !IsNil(o.AzurePrivateOffer) {
		return true
	}

	return false
}

// SetAzurePrivateOffer gets a reference to the given AzureMarketplacePrivateOffer and assigns it to the AzurePrivateOffer field.
func (o *OfferInfo) SetAzurePrivateOffer(v AzureMarketplacePrivateOffer) {
	o.AzurePrivateOffer = &v
}

// GetAzureProductVariant returns the AzureProductVariant field value if set, zero value otherwise.
func (o *OfferInfo) GetAzureProductVariant() AzureProductVariant {
	if o == nil || IsNil(o.AzureProductVariant) {
		var ret AzureProductVariant
		return ret
	}
	return *o.AzureProductVariant
}

// GetAzureProductVariantOk returns a tuple with the AzureProductVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetAzureProductVariantOk() (*AzureProductVariant, bool) {
	if o == nil || IsNil(o.AzureProductVariant) {
		return nil, false
	}
	return o.AzureProductVariant, true
}

// HasAzureProductVariant returns a boolean if a field has been set.
func (o *OfferInfo) HasAzureProductVariant() bool {
	if o != nil && !IsNil(o.AzureProductVariant) {
		return true
	}

	return false
}

// SetAzureProductVariant gets a reference to the given AzureProductVariant and assigns it to the AzureProductVariant field.
func (o *OfferInfo) SetAzureProductVariant(v AzureProductVariant) {
	o.AzureProductVariant = &v
}

// GetBuyerAwsAccountIds returns the BuyerAwsAccountIds field value if set, zero value otherwise.
func (o *OfferInfo) GetBuyerAwsAccountIds() []string {
	if o == nil || IsNil(o.BuyerAwsAccountIds) {
		var ret []string
		return ret
	}
	return o.BuyerAwsAccountIds
}

// GetBuyerAwsAccountIdsOk returns a tuple with the BuyerAwsAccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetBuyerAwsAccountIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.BuyerAwsAccountIds) {
		return nil, false
	}
	return o.BuyerAwsAccountIds, true
}

// HasBuyerAwsAccountIds returns a boolean if a field has been set.
func (o *OfferInfo) HasBuyerAwsAccountIds() bool {
	if o != nil && !IsNil(o.BuyerAwsAccountIds) {
		return true
	}

	return false
}

// SetBuyerAwsAccountIds gets a reference to the given []string and assigns it to the BuyerAwsAccountIds field.
func (o *OfferInfo) SetBuyerAwsAccountIds(v []string) {
	o.BuyerAwsAccountIds = v
}

// GetBuyerAzureTenants returns the BuyerAzureTenants field value if set, zero value otherwise.
func (o *OfferInfo) GetBuyerAzureTenants() []AzureAudience {
	if o == nil || IsNil(o.BuyerAzureTenants) {
		var ret []AzureAudience
		return ret
	}
	return o.BuyerAzureTenants
}

// GetBuyerAzureTenantsOk returns a tuple with the BuyerAzureTenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetBuyerAzureTenantsOk() ([]AzureAudience, bool) {
	if o == nil || IsNil(o.BuyerAzureTenants) {
		return nil, false
	}
	return o.BuyerAzureTenants, true
}

// HasBuyerAzureTenants returns a boolean if a field has been set.
func (o *OfferInfo) HasBuyerAzureTenants() bool {
	if o != nil && !IsNil(o.BuyerAzureTenants) {
		return true
	}

	return false
}

// SetBuyerAzureTenants gets a reference to the given []AzureAudience and assigns it to the BuyerAzureTenants field.
func (o *OfferInfo) SetBuyerAzureTenants(v []AzureAudience) {
	o.BuyerAzureTenants = v
}

// GetCommitAmount returns the CommitAmount field value if set, zero value otherwise.
func (o *OfferInfo) GetCommitAmount() float32 {
	if o == nil || IsNil(o.CommitAmount) {
		var ret float32
		return ret
	}
	return *o.CommitAmount
}

// GetCommitAmountOk returns a tuple with the CommitAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetCommitAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CommitAmount) {
		return nil, false
	}
	return o.CommitAmount, true
}

// HasCommitAmount returns a boolean if a field has been set.
func (o *OfferInfo) HasCommitAmount() bool {
	if o != nil && !IsNil(o.CommitAmount) {
		return true
	}

	return false
}

// SetCommitAmount gets a reference to the given float32 and assigns it to the CommitAmount field.
func (o *OfferInfo) SetCommitAmount(v float32) {
	o.CommitAmount = &v
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *OfferInfo) GetCommits() []CommitDimension {
	if o == nil || IsNil(o.Commits) {
		var ret []CommitDimension
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetCommitsOk() ([]CommitDimension, bool) {
	if o == nil || IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *OfferInfo) HasCommits() bool {
	if o != nil && !IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []CommitDimension and assigns it to the Commits field.
func (o *OfferInfo) SetCommits(v []CommitDimension) {
	o.Commits = v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *OfferInfo) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *OfferInfo) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *OfferInfo) SetCurrency(v string) {
	o.Currency = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *OfferInfo) GetDimensions() []MeteringDimension {
	if o == nil || IsNil(o.Dimensions) {
		var ret []MeteringDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetDimensionsOk() ([]MeteringDimension, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *OfferInfo) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []MeteringDimension and assigns it to the Dimensions field.
func (o *OfferInfo) SetDimensions(v []MeteringDimension) {
	o.Dimensions = v
}

// GetDiscountPercentage returns the DiscountPercentage field value if set, zero value otherwise.
func (o *OfferInfo) GetDiscountPercentage() float32 {
	if o == nil || IsNil(o.DiscountPercentage) {
		var ret float32
		return ret
	}
	return *o.DiscountPercentage
}

// GetDiscountPercentageOk returns a tuple with the DiscountPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetDiscountPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.DiscountPercentage) {
		return nil, false
	}
	return o.DiscountPercentage, true
}

// HasDiscountPercentage returns a boolean if a field has been set.
func (o *OfferInfo) HasDiscountPercentage() bool {
	if o != nil && !IsNil(o.DiscountPercentage) {
		return true
	}

	return false
}

// SetDiscountPercentage gets a reference to the given float32 and assigns it to the DiscountPercentage field.
func (o *OfferInfo) SetDiscountPercentage(v float32) {
	o.DiscountPercentage = &v
}

// GetEulaType returns the EulaType field value if set, zero value otherwise.
func (o *OfferInfo) GetEulaType() EulaType {
	if o == nil || IsNil(o.EulaType) {
		var ret EulaType
		return ret
	}
	return *o.EulaType
}

// GetEulaTypeOk returns a tuple with the EulaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetEulaTypeOk() (*EulaType, bool) {
	if o == nil || IsNil(o.EulaType) {
		return nil, false
	}
	return o.EulaType, true
}

// HasEulaType returns a boolean if a field has been set.
func (o *OfferInfo) HasEulaType() bool {
	if o != nil && !IsNil(o.EulaType) {
		return true
	}

	return false
}

// SetEulaType gets a reference to the given EulaType and assigns it to the EulaType field.
func (o *OfferInfo) SetEulaType(v EulaType) {
	o.EulaType = &v
}

// GetEulaUrl returns the EulaUrl field value if set, zero value otherwise.
func (o *OfferInfo) GetEulaUrl() string {
	if o == nil || IsNil(o.EulaUrl) {
		var ret string
		return ret
	}
	return *o.EulaUrl
}

// GetEulaUrlOk returns a tuple with the EulaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetEulaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EulaUrl) {
		return nil, false
	}
	return o.EulaUrl, true
}

// HasEulaUrl returns a boolean if a field has been set.
func (o *OfferInfo) HasEulaUrl() bool {
	if o != nil && !IsNil(o.EulaUrl) {
		return true
	}

	return false
}

// SetEulaUrl gets a reference to the given string and assigns it to the EulaUrl field.
func (o *OfferInfo) SetEulaUrl(v string) {
	o.EulaUrl = &v
}

// GetGcpCustomerInfo returns the GcpCustomerInfo field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpCustomerInfo() GcpMarketplacePrivateOfferCustomerInfo {
	if o == nil || IsNil(o.GcpCustomerInfo) {
		var ret GcpMarketplacePrivateOfferCustomerInfo
		return ret
	}
	return *o.GcpCustomerInfo
}

// GetGcpCustomerInfoOk returns a tuple with the GcpCustomerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpCustomerInfoOk() (*GcpMarketplacePrivateOfferCustomerInfo, bool) {
	if o == nil || IsNil(o.GcpCustomerInfo) {
		return nil, false
	}
	return o.GcpCustomerInfo, true
}

// HasGcpCustomerInfo returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpCustomerInfo() bool {
	if o != nil && !IsNil(o.GcpCustomerInfo) {
		return true
	}

	return false
}

// SetGcpCustomerInfo gets a reference to the given GcpMarketplacePrivateOfferCustomerInfo and assigns it to the GcpCustomerInfo field.
func (o *OfferInfo) SetGcpCustomerInfo(v GcpMarketplacePrivateOfferCustomerInfo) {
	o.GcpCustomerInfo = &v
}

// GetGcpDuration returns the GcpDuration field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpDuration() int32 {
	if o == nil || IsNil(o.GcpDuration) {
		var ret int32
		return ret
	}
	return *o.GcpDuration
}

// GetGcpDurationOk returns a tuple with the GcpDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.GcpDuration) {
		return nil, false
	}
	return o.GcpDuration, true
}

// HasGcpDuration returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpDuration() bool {
	if o != nil && !IsNil(o.GcpDuration) {
		return true
	}

	return false
}

// SetGcpDuration gets a reference to the given int32 and assigns it to the GcpDuration field.
func (o *OfferInfo) SetGcpDuration(v int32) {
	o.GcpDuration = &v
}

// GetGcpMetrics returns the GcpMetrics field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpMetrics() []GcpMarketplaceProductMeteringMetric {
	if o == nil || IsNil(o.GcpMetrics) {
		var ret []GcpMarketplaceProductMeteringMetric
		return ret
	}
	return o.GcpMetrics
}

// GetGcpMetricsOk returns a tuple with the GcpMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpMetricsOk() ([]GcpMarketplaceProductMeteringMetric, bool) {
	if o == nil || IsNil(o.GcpMetrics) {
		return nil, false
	}
	return o.GcpMetrics, true
}

// HasGcpMetrics returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpMetrics() bool {
	if o != nil && !IsNil(o.GcpMetrics) {
		return true
	}

	return false
}

// SetGcpMetrics gets a reference to the given []GcpMarketplaceProductMeteringMetric and assigns it to the GcpMetrics field.
func (o *OfferInfo) SetGcpMetrics(v []GcpMarketplaceProductMeteringMetric) {
	o.GcpMetrics = v
}

// GetGcpPaymentSchedule returns the GcpPaymentSchedule field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpPaymentSchedule() GcpMarketplacePaymentScheduleType {
	if o == nil || IsNil(o.GcpPaymentSchedule) {
		var ret GcpMarketplacePaymentScheduleType
		return ret
	}
	return *o.GcpPaymentSchedule
}

// GetGcpPaymentScheduleOk returns a tuple with the GcpPaymentSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpPaymentScheduleOk() (*GcpMarketplacePaymentScheduleType, bool) {
	if o == nil || IsNil(o.GcpPaymentSchedule) {
		return nil, false
	}
	return o.GcpPaymentSchedule, true
}

// HasGcpPaymentSchedule returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpPaymentSchedule() bool {
	if o != nil && !IsNil(o.GcpPaymentSchedule) {
		return true
	}

	return false
}

// SetGcpPaymentSchedule gets a reference to the given GcpMarketplacePaymentScheduleType and assigns it to the GcpPaymentSchedule field.
func (o *OfferInfo) SetGcpPaymentSchedule(v GcpMarketplacePaymentScheduleType) {
	o.GcpPaymentSchedule = &v
}

// GetGcpPlans returns the GcpPlans field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpPlans() []GcpMarketplaceProductPurchaseOptionSpec {
	if o == nil || IsNil(o.GcpPlans) {
		var ret []GcpMarketplaceProductPurchaseOptionSpec
		return ret
	}
	return o.GcpPlans
}

// GetGcpPlansOk returns a tuple with the GcpPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpPlansOk() ([]GcpMarketplaceProductPurchaseOptionSpec, bool) {
	if o == nil || IsNil(o.GcpPlans) {
		return nil, false
	}
	return o.GcpPlans, true
}

// HasGcpPlans returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpPlans() bool {
	if o != nil && !IsNil(o.GcpPlans) {
		return true
	}

	return false
}

// SetGcpPlans gets a reference to the given []GcpMarketplaceProductPurchaseOptionSpec and assigns it to the GcpPlans field.
func (o *OfferInfo) SetGcpPlans(v []GcpMarketplaceProductPurchaseOptionSpec) {
	o.GcpPlans = v
}

// GetGcpPrivateOffer returns the GcpPrivateOffer field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpPrivateOffer() GcpMarketplacePrivateOffer {
	if o == nil || IsNil(o.GcpPrivateOffer) {
		var ret GcpMarketplacePrivateOffer
		return ret
	}
	return *o.GcpPrivateOffer
}

// GetGcpPrivateOfferOk returns a tuple with the GcpPrivateOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpPrivateOfferOk() (*GcpMarketplacePrivateOffer, bool) {
	if o == nil || IsNil(o.GcpPrivateOffer) {
		return nil, false
	}
	return o.GcpPrivateOffer, true
}

// HasGcpPrivateOffer returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpPrivateOffer() bool {
	if o != nil && !IsNil(o.GcpPrivateOffer) {
		return true
	}

	return false
}

// SetGcpPrivateOffer gets a reference to the given GcpMarketplacePrivateOffer and assigns it to the GcpPrivateOffer field.
func (o *OfferInfo) SetGcpPrivateOffer(v GcpMarketplacePrivateOffer) {
	o.GcpPrivateOffer = &v
}

// GetGcpProviderInfo returns the GcpProviderInfo field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpProviderInfo() GcpMarketplacePrivateOfferProviderInfo {
	if o == nil || IsNil(o.GcpProviderInfo) {
		var ret GcpMarketplacePrivateOfferProviderInfo
		return ret
	}
	return *o.GcpProviderInfo
}

// GetGcpProviderInfoOk returns a tuple with the GcpProviderInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpProviderInfoOk() (*GcpMarketplacePrivateOfferProviderInfo, bool) {
	if o == nil || IsNil(o.GcpProviderInfo) {
		return nil, false
	}
	return o.GcpProviderInfo, true
}

// HasGcpProviderInfo returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpProviderInfo() bool {
	if o != nil && !IsNil(o.GcpProviderInfo) {
		return true
	}

	return false
}

// SetGcpProviderInfo gets a reference to the given GcpMarketplacePrivateOfferProviderInfo and assigns it to the GcpProviderInfo field.
func (o *OfferInfo) SetGcpProviderInfo(v GcpMarketplacePrivateOfferProviderInfo) {
	o.GcpProviderInfo = &v
}

// GetGcpProviderInternalNote returns the GcpProviderInternalNote field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpProviderInternalNote() string {
	if o == nil || IsNil(o.GcpProviderInternalNote) {
		var ret string
		return ret
	}
	return *o.GcpProviderInternalNote
}

// GetGcpProviderInternalNoteOk returns a tuple with the GcpProviderInternalNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpProviderInternalNoteOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProviderInternalNote) {
		return nil, false
	}
	return o.GcpProviderInternalNote, true
}

// HasGcpProviderInternalNote returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpProviderInternalNote() bool {
	if o != nil && !IsNil(o.GcpProviderInternalNote) {
		return true
	}

	return false
}

// SetGcpProviderInternalNote gets a reference to the given string and assigns it to the GcpProviderInternalNote field.
func (o *OfferInfo) SetGcpProviderInternalNote(v string) {
	o.GcpProviderInternalNote = &v
}

// GetGcpProviderPublicNote returns the GcpProviderPublicNote field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpProviderPublicNote() string {
	if o == nil || IsNil(o.GcpProviderPublicNote) {
		var ret string
		return ret
	}
	return *o.GcpProviderPublicNote
}

// GetGcpProviderPublicNoteOk returns a tuple with the GcpProviderPublicNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpProviderPublicNoteOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProviderPublicNote) {
		return nil, false
	}
	return o.GcpProviderPublicNote, true
}

// HasGcpProviderPublicNote returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpProviderPublicNote() bool {
	if o != nil && !IsNil(o.GcpProviderPublicNote) {
		return true
	}

	return false
}

// SetGcpProviderPublicNote gets a reference to the given string and assigns it to the GcpProviderPublicNote field.
func (o *OfferInfo) SetGcpProviderPublicNote(v string) {
	o.GcpProviderPublicNote = &v
}

// GetGcpUsagePlanPriceModel returns the GcpUsagePlanPriceModel field value if set, zero value otherwise.
func (o *OfferInfo) GetGcpUsagePlanPriceModel() GcpMarketplaceUsagePlanPriceModel {
	if o == nil || IsNil(o.GcpUsagePlanPriceModel) {
		var ret GcpMarketplaceUsagePlanPriceModel
		return ret
	}
	return *o.GcpUsagePlanPriceModel
}

// GetGcpUsagePlanPriceModelOk returns a tuple with the GcpUsagePlanPriceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetGcpUsagePlanPriceModelOk() (*GcpMarketplaceUsagePlanPriceModel, bool) {
	if o == nil || IsNil(o.GcpUsagePlanPriceModel) {
		return nil, false
	}
	return o.GcpUsagePlanPriceModel, true
}

// HasGcpUsagePlanPriceModel returns a boolean if a field has been set.
func (o *OfferInfo) HasGcpUsagePlanPriceModel() bool {
	if o != nil && !IsNil(o.GcpUsagePlanPriceModel) {
		return true
	}

	return false
}

// SetGcpUsagePlanPriceModel gets a reference to the given GcpMarketplaceUsagePlanPriceModel and assigns it to the GcpUsagePlanPriceModel field.
func (o *OfferInfo) SetGcpUsagePlanPriceModel(v GcpMarketplaceUsagePlanPriceModel) {
	o.GcpUsagePlanPriceModel = &v
}

// GetPaymentInstallments returns the PaymentInstallments field value if set, zero value otherwise.
func (o *OfferInfo) GetPaymentInstallments() []PaymentInstallment {
	if o == nil || IsNil(o.PaymentInstallments) {
		var ret []PaymentInstallment
		return ret
	}
	return o.PaymentInstallments
}

// GetPaymentInstallmentsOk returns a tuple with the PaymentInstallments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetPaymentInstallmentsOk() ([]PaymentInstallment, bool) {
	if o == nil || IsNil(o.PaymentInstallments) {
		return nil, false
	}
	return o.PaymentInstallments, true
}

// HasPaymentInstallments returns a boolean if a field has been set.
func (o *OfferInfo) HasPaymentInstallments() bool {
	if o != nil && !IsNil(o.PaymentInstallments) {
		return true
	}

	return false
}

// SetPaymentInstallments gets a reference to the given []PaymentInstallment and assigns it to the PaymentInstallments field.
func (o *OfferInfo) SetPaymentInstallments(v []PaymentInstallment) {
	o.PaymentInstallments = v
}

// GetPrivateOfferUrl returns the PrivateOfferUrl field value if set, zero value otherwise.
func (o *OfferInfo) GetPrivateOfferUrl() string {
	if o == nil || IsNil(o.PrivateOfferUrl) {
		var ret string
		return ret
	}
	return *o.PrivateOfferUrl
}

// GetPrivateOfferUrlOk returns a tuple with the PrivateOfferUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetPrivateOfferUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateOfferUrl) {
		return nil, false
	}
	return o.PrivateOfferUrl, true
}

// HasPrivateOfferUrl returns a boolean if a field has been set.
func (o *OfferInfo) HasPrivateOfferUrl() bool {
	if o != nil && !IsNil(o.PrivateOfferUrl) {
		return true
	}

	return false
}

// SetPrivateOfferUrl gets a reference to the given string and assigns it to the PrivateOfferUrl field.
func (o *OfferInfo) SetPrivateOfferUrl(v string) {
	o.PrivateOfferUrl = &v
}

// GetRefundCancelationPolicy returns the RefundCancelationPolicy field value if set, zero value otherwise.
func (o *OfferInfo) GetRefundCancelationPolicy() string {
	if o == nil || IsNil(o.RefundCancelationPolicy) {
		var ret string
		return ret
	}
	return *o.RefundCancelationPolicy
}

// GetRefundCancelationPolicyOk returns a tuple with the RefundCancelationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetRefundCancelationPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.RefundCancelationPolicy) {
		return nil, false
	}
	return o.RefundCancelationPolicy, true
}

// HasRefundCancelationPolicy returns a boolean if a field has been set.
func (o *OfferInfo) HasRefundCancelationPolicy() bool {
	if o != nil && !IsNil(o.RefundCancelationPolicy) {
		return true
	}

	return false
}

// SetRefundCancelationPolicy gets a reference to the given string and assigns it to the RefundCancelationPolicy field.
func (o *OfferInfo) SetRefundCancelationPolicy(v string) {
	o.RefundCancelationPolicy = &v
}

// GetSellerNotes returns the SellerNotes field value if set, zero value otherwise.
func (o *OfferInfo) GetSellerNotes() string {
	if o == nil || IsNil(o.SellerNotes) {
		var ret string
		return ret
	}
	return *o.SellerNotes
}

// GetSellerNotesOk returns a tuple with the SellerNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetSellerNotesOk() (*string, bool) {
	if o == nil || IsNil(o.SellerNotes) {
		return nil, false
	}
	return o.SellerNotes, true
}

// HasSellerNotes returns a boolean if a field has been set.
func (o *OfferInfo) HasSellerNotes() bool {
	if o != nil && !IsNil(o.SellerNotes) {
		return true
	}

	return false
}

// SetSellerNotes gets a reference to the given string and assigns it to the SellerNotes field.
func (o *OfferInfo) SetSellerNotes(v string) {
	o.SellerNotes = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *OfferInfo) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferInfo) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *OfferInfo) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *OfferInfo) SetVisibility(v string) {
	o.Visibility = &v
}

func (o OfferInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfferInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachEulaType) {
		toSerialize["attachEulaType"] = o.AttachEulaType
	}
	if !IsNil(o.AutoRenew) {
		toSerialize["autoRenew"] = o.AutoRenew
	}
	if !IsNil(o.AwsCppoEventDetail) {
		toSerialize["awsCppoEventDetail"] = o.AwsCppoEventDetail
	}
	if !IsNil(o.AwsCppoOpportunity) {
		toSerialize["awsCppoOpportunity"] = o.AwsCppoOpportunity
	}
	if !IsNil(o.AzureOriginalPlan) {
		toSerialize["azureOriginalPlan"] = o.AzureOriginalPlan
	}
	if !IsNil(o.AzurePrivateOffer) {
		toSerialize["azurePrivateOffer"] = o.AzurePrivateOffer
	}
	if !IsNil(o.AzureProductVariant) {
		toSerialize["azureProductVariant"] = o.AzureProductVariant
	}
	if !IsNil(o.BuyerAwsAccountIds) {
		toSerialize["buyerAwsAccountIds"] = o.BuyerAwsAccountIds
	}
	if !IsNil(o.BuyerAzureTenants) {
		toSerialize["buyerAzureTenants"] = o.BuyerAzureTenants
	}
	if !IsNil(o.CommitAmount) {
		toSerialize["commitAmount"] = o.CommitAmount
	}
	if !IsNil(o.Commits) {
		toSerialize["commits"] = o.Commits
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.DiscountPercentage) {
		toSerialize["discountPercentage"] = o.DiscountPercentage
	}
	if !IsNil(o.EulaType) {
		toSerialize["eulaType"] = o.EulaType
	}
	if !IsNil(o.EulaUrl) {
		toSerialize["eulaUrl"] = o.EulaUrl
	}
	if !IsNil(o.GcpCustomerInfo) {
		toSerialize["gcpCustomerInfo"] = o.GcpCustomerInfo
	}
	if !IsNil(o.GcpDuration) {
		toSerialize["gcpDuration"] = o.GcpDuration
	}
	if !IsNil(o.GcpMetrics) {
		toSerialize["gcpMetrics"] = o.GcpMetrics
	}
	if !IsNil(o.GcpPaymentSchedule) {
		toSerialize["gcpPaymentSchedule"] = o.GcpPaymentSchedule
	}
	if !IsNil(o.GcpPlans) {
		toSerialize["gcpPlans"] = o.GcpPlans
	}
	if !IsNil(o.GcpPrivateOffer) {
		toSerialize["gcpPrivateOffer"] = o.GcpPrivateOffer
	}
	if !IsNil(o.GcpProviderInfo) {
		toSerialize["gcpProviderInfo"] = o.GcpProviderInfo
	}
	if !IsNil(o.GcpProviderInternalNote) {
		toSerialize["gcpProviderInternalNote"] = o.GcpProviderInternalNote
	}
	if !IsNil(o.GcpProviderPublicNote) {
		toSerialize["gcpProviderPublicNote"] = o.GcpProviderPublicNote
	}
	if !IsNil(o.GcpUsagePlanPriceModel) {
		toSerialize["gcpUsagePlanPriceModel"] = o.GcpUsagePlanPriceModel
	}
	if !IsNil(o.PaymentInstallments) {
		toSerialize["paymentInstallments"] = o.PaymentInstallments
	}
	if !IsNil(o.PrivateOfferUrl) {
		toSerialize["privateOfferUrl"] = o.PrivateOfferUrl
	}
	if !IsNil(o.RefundCancelationPolicy) {
		toSerialize["refundCancelationPolicy"] = o.RefundCancelationPolicy
	}
	if !IsNil(o.SellerNotes) {
		toSerialize["sellerNotes"] = o.SellerNotes
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableOfferInfo struct {
	value *OfferInfo
	isSet bool
}

func (v NullableOfferInfo) Get() *OfferInfo {
	return v.value
}

func (v *NullableOfferInfo) Set(val *OfferInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferInfo(val *OfferInfo) *NullableOfferInfo {
	return &NullableOfferInfo{value: val, isSet: true}
}

func (v NullableOfferInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


