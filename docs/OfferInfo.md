# OfferInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachEulaType** | Pointer to [**EulaType**](EulaType.md) |  | [optional] 
**AutoRenew** | Pointer to **bool** | Is this offer Auto Renew enabled. | [optional] 
**AwsCppoEventDetail** | Pointer to [**AwsMarketplaceEventBridgeEventDetail**](AwsMarketplaceEventBridgeEventDetail.md) |  | [optional] 
**AwsCppoOpportunity** | Pointer to [**AwsMarketplaceCppoOpportunity**](AwsMarketplaceCppoOpportunity.md) |  | [optional] 
**AzureOriginalPlan** | Pointer to [**AzurePriceAndAvailabilityPrivateOfferPlan**](AzurePriceAndAvailabilityPrivateOfferPlan.md) |  | [optional] 
**AzurePrivateOffer** | Pointer to [**AzureMarketplacePrivateOffer**](AzureMarketplacePrivateOffer.md) |  | [optional] 
**AzureProductVariant** | Pointer to [**AzureProductVariant**](AzureProductVariant.md) |  | [optional] 
**BuyerAwsAccountIds** | Pointer to **[]string** | The buyers&#39; AWS Account IDs of this offer. | [optional] 
**BuyerAzureTenants** | Pointer to [**[]AzureAudience**](AzureAudience.md) | The buyers&#39; Azure tenants of this offer. | [optional] 
**CommitAmount** | Pointer to **float32** | The amount that the buyer has committed to pay, before discount if applicable. It can be monthly commitment or total commitment. | [optional] 
**Commits** | Pointer to [**[]CommitDimension**](CommitDimension.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Dimensions** | Pointer to [**[]MeteringDimension**](MeteringDimension.md) |  | [optional] 
**DiscountPercentage** | Pointer to **float32** | The discount percentage off the original price. For example, 20 means 20% off. 0 means no discount. It can be discount off the commitment amount or discount off the usage price. | [optional] 
**EulaType** | Pointer to [**EulaType**](EulaType.md) |  | [optional] 
**EulaUrl** | Pointer to **string** |  | [optional] 
**GcpCustomerInfo** | Pointer to [**GcpMarketplacePrivateOfferCustomerInfo**](GcpMarketplacePrivateOfferCustomerInfo.md) |  | [optional] 
**GcpDuration** | Pointer to **int32** | The duration of the offer in months. Only required when creating GCP Marketplace private offer. | [optional] 
**GcpMetrics** | Pointer to [**[]GcpMarketplaceProductMeteringMetric**](GcpMarketplaceProductMeteringMetric.md) | Only applicable for GCP Marketplace Offers (the default or private offer) | [optional] 
**GcpPaymentSchedule** | Pointer to [**GcpMarketplacePaymentScheduleType**](GcpMarketplacePaymentScheduleType.md) |  | [optional] 
**GcpPlans** | Pointer to [**[]GcpMarketplaceProductPurchaseOptionSpec**](GcpMarketplaceProductPurchaseOptionSpec.md) | Only applicable for GCP Marketplace | [optional] 
**GcpPrivateOffer** | Pointer to [**GcpMarketplacePrivateOffer**](GcpMarketplacePrivateOffer.md) |  | [optional] 
**GcpProviderInfo** | Pointer to [**GcpMarketplacePrivateOfferProviderInfo**](GcpMarketplacePrivateOfferProviderInfo.md) |  | [optional] 
**GcpProviderInternalNote** | Pointer to **string** | Optional when creating GCP Marketplace private offer. The internal note for the seller/ISV. It is only visible to the seller/ISV. | [optional] 
**GcpProviderPublicNote** | Pointer to **string** | Optional when creating GCP Marketplace private offer. The public note for the buyer. It is visible to the buyer. | [optional] 
**GcpUsagePlanPriceModel** | Pointer to [**GcpMarketplaceUsagePlanPriceModel**](GcpMarketplaceUsagePlanPriceModel.md) |  | [optional] 
**PaymentInstallments** | Pointer to [**[]PaymentInstallment**](PaymentInstallment.md) | For flexible payment schedule. | [optional] 
**PrivateOfferUrl** | Pointer to **string** | The URL of the private offer sent to buyers to accept. Only applicable for private offer. | [optional] 
**RefundCancelationPolicy** | Pointer to **string** |  | [optional] 
**SellerNotes** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** | The default visibility of offer is PRIVATE. | [optional] 

## Methods

### NewOfferInfo

`func NewOfferInfo() *OfferInfo`

NewOfferInfo instantiates a new OfferInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferInfoWithDefaults

`func NewOfferInfoWithDefaults() *OfferInfo`

NewOfferInfoWithDefaults instantiates a new OfferInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachEulaType

`func (o *OfferInfo) GetAttachEulaType() EulaType`

GetAttachEulaType returns the AttachEulaType field if non-nil, zero value otherwise.

### GetAttachEulaTypeOk

`func (o *OfferInfo) GetAttachEulaTypeOk() (*EulaType, bool)`

GetAttachEulaTypeOk returns a tuple with the AttachEulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachEulaType

`func (o *OfferInfo) SetAttachEulaType(v EulaType)`

SetAttachEulaType sets AttachEulaType field to given value.

### HasAttachEulaType

`func (o *OfferInfo) HasAttachEulaType() bool`

HasAttachEulaType returns a boolean if a field has been set.

### GetAutoRenew

`func (o *OfferInfo) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *OfferInfo) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *OfferInfo) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *OfferInfo) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetAwsCppoEventDetail

`func (o *OfferInfo) GetAwsCppoEventDetail() AwsMarketplaceEventBridgeEventDetail`

GetAwsCppoEventDetail returns the AwsCppoEventDetail field if non-nil, zero value otherwise.

### GetAwsCppoEventDetailOk

`func (o *OfferInfo) GetAwsCppoEventDetailOk() (*AwsMarketplaceEventBridgeEventDetail, bool)`

GetAwsCppoEventDetailOk returns a tuple with the AwsCppoEventDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCppoEventDetail

`func (o *OfferInfo) SetAwsCppoEventDetail(v AwsMarketplaceEventBridgeEventDetail)`

SetAwsCppoEventDetail sets AwsCppoEventDetail field to given value.

### HasAwsCppoEventDetail

`func (o *OfferInfo) HasAwsCppoEventDetail() bool`

HasAwsCppoEventDetail returns a boolean if a field has been set.

### GetAwsCppoOpportunity

`func (o *OfferInfo) GetAwsCppoOpportunity() AwsMarketplaceCppoOpportunity`

GetAwsCppoOpportunity returns the AwsCppoOpportunity field if non-nil, zero value otherwise.

### GetAwsCppoOpportunityOk

`func (o *OfferInfo) GetAwsCppoOpportunityOk() (*AwsMarketplaceCppoOpportunity, bool)`

GetAwsCppoOpportunityOk returns a tuple with the AwsCppoOpportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCppoOpportunity

`func (o *OfferInfo) SetAwsCppoOpportunity(v AwsMarketplaceCppoOpportunity)`

SetAwsCppoOpportunity sets AwsCppoOpportunity field to given value.

### HasAwsCppoOpportunity

`func (o *OfferInfo) HasAwsCppoOpportunity() bool`

HasAwsCppoOpportunity returns a boolean if a field has been set.

### GetAzureOriginalPlan

`func (o *OfferInfo) GetAzureOriginalPlan() AzurePriceAndAvailabilityPrivateOfferPlan`

GetAzureOriginalPlan returns the AzureOriginalPlan field if non-nil, zero value otherwise.

### GetAzureOriginalPlanOk

`func (o *OfferInfo) GetAzureOriginalPlanOk() (*AzurePriceAndAvailabilityPrivateOfferPlan, bool)`

GetAzureOriginalPlanOk returns a tuple with the AzureOriginalPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureOriginalPlan

`func (o *OfferInfo) SetAzureOriginalPlan(v AzurePriceAndAvailabilityPrivateOfferPlan)`

SetAzureOriginalPlan sets AzureOriginalPlan field to given value.

### HasAzureOriginalPlan

`func (o *OfferInfo) HasAzureOriginalPlan() bool`

HasAzureOriginalPlan returns a boolean if a field has been set.

### GetAzurePrivateOffer

`func (o *OfferInfo) GetAzurePrivateOffer() AzureMarketplacePrivateOffer`

GetAzurePrivateOffer returns the AzurePrivateOffer field if non-nil, zero value otherwise.

### GetAzurePrivateOfferOk

`func (o *OfferInfo) GetAzurePrivateOfferOk() (*AzureMarketplacePrivateOffer, bool)`

GetAzurePrivateOfferOk returns a tuple with the AzurePrivateOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurePrivateOffer

`func (o *OfferInfo) SetAzurePrivateOffer(v AzureMarketplacePrivateOffer)`

SetAzurePrivateOffer sets AzurePrivateOffer field to given value.

### HasAzurePrivateOffer

`func (o *OfferInfo) HasAzurePrivateOffer() bool`

HasAzurePrivateOffer returns a boolean if a field has been set.

### GetAzureProductVariant

`func (o *OfferInfo) GetAzureProductVariant() AzureProductVariant`

GetAzureProductVariant returns the AzureProductVariant field if non-nil, zero value otherwise.

### GetAzureProductVariantOk

`func (o *OfferInfo) GetAzureProductVariantOk() (*AzureProductVariant, bool)`

GetAzureProductVariantOk returns a tuple with the AzureProductVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureProductVariant

`func (o *OfferInfo) SetAzureProductVariant(v AzureProductVariant)`

SetAzureProductVariant sets AzureProductVariant field to given value.

### HasAzureProductVariant

`func (o *OfferInfo) HasAzureProductVariant() bool`

HasAzureProductVariant returns a boolean if a field has been set.

### GetBuyerAwsAccountIds

`func (o *OfferInfo) GetBuyerAwsAccountIds() []string`

GetBuyerAwsAccountIds returns the BuyerAwsAccountIds field if non-nil, zero value otherwise.

### GetBuyerAwsAccountIdsOk

`func (o *OfferInfo) GetBuyerAwsAccountIdsOk() (*[]string, bool)`

GetBuyerAwsAccountIdsOk returns a tuple with the BuyerAwsAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAwsAccountIds

`func (o *OfferInfo) SetBuyerAwsAccountIds(v []string)`

SetBuyerAwsAccountIds sets BuyerAwsAccountIds field to given value.

### HasBuyerAwsAccountIds

`func (o *OfferInfo) HasBuyerAwsAccountIds() bool`

HasBuyerAwsAccountIds returns a boolean if a field has been set.

### GetBuyerAzureTenants

`func (o *OfferInfo) GetBuyerAzureTenants() []AzureAudience`

GetBuyerAzureTenants returns the BuyerAzureTenants field if non-nil, zero value otherwise.

### GetBuyerAzureTenantsOk

`func (o *OfferInfo) GetBuyerAzureTenantsOk() (*[]AzureAudience, bool)`

GetBuyerAzureTenantsOk returns a tuple with the BuyerAzureTenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAzureTenants

`func (o *OfferInfo) SetBuyerAzureTenants(v []AzureAudience)`

SetBuyerAzureTenants sets BuyerAzureTenants field to given value.

### HasBuyerAzureTenants

`func (o *OfferInfo) HasBuyerAzureTenants() bool`

HasBuyerAzureTenants returns a boolean if a field has been set.

### GetCommitAmount

`func (o *OfferInfo) GetCommitAmount() float32`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *OfferInfo) GetCommitAmountOk() (*float32, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *OfferInfo) SetCommitAmount(v float32)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *OfferInfo) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetCommits

`func (o *OfferInfo) GetCommits() []CommitDimension`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *OfferInfo) GetCommitsOk() (*[]CommitDimension, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *OfferInfo) SetCommits(v []CommitDimension)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *OfferInfo) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetCurrency

`func (o *OfferInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OfferInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OfferInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OfferInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDimensions

`func (o *OfferInfo) GetDimensions() []MeteringDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *OfferInfo) GetDimensionsOk() (*[]MeteringDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *OfferInfo) SetDimensions(v []MeteringDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *OfferInfo) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *OfferInfo) GetDiscountPercentage() float32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *OfferInfo) GetDiscountPercentageOk() (*float32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *OfferInfo) SetDiscountPercentage(v float32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *OfferInfo) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetEulaType

`func (o *OfferInfo) GetEulaType() EulaType`

GetEulaType returns the EulaType field if non-nil, zero value otherwise.

### GetEulaTypeOk

`func (o *OfferInfo) GetEulaTypeOk() (*EulaType, bool)`

GetEulaTypeOk returns a tuple with the EulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaType

`func (o *OfferInfo) SetEulaType(v EulaType)`

SetEulaType sets EulaType field to given value.

### HasEulaType

`func (o *OfferInfo) HasEulaType() bool`

HasEulaType returns a boolean if a field has been set.

### GetEulaUrl

`func (o *OfferInfo) GetEulaUrl() string`

GetEulaUrl returns the EulaUrl field if non-nil, zero value otherwise.

### GetEulaUrlOk

`func (o *OfferInfo) GetEulaUrlOk() (*string, bool)`

GetEulaUrlOk returns a tuple with the EulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaUrl

`func (o *OfferInfo) SetEulaUrl(v string)`

SetEulaUrl sets EulaUrl field to given value.

### HasEulaUrl

`func (o *OfferInfo) HasEulaUrl() bool`

HasEulaUrl returns a boolean if a field has been set.

### GetGcpCustomerInfo

`func (o *OfferInfo) GetGcpCustomerInfo() GcpMarketplacePrivateOfferCustomerInfo`

GetGcpCustomerInfo returns the GcpCustomerInfo field if non-nil, zero value otherwise.

### GetGcpCustomerInfoOk

`func (o *OfferInfo) GetGcpCustomerInfoOk() (*GcpMarketplacePrivateOfferCustomerInfo, bool)`

GetGcpCustomerInfoOk returns a tuple with the GcpCustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCustomerInfo

`func (o *OfferInfo) SetGcpCustomerInfo(v GcpMarketplacePrivateOfferCustomerInfo)`

SetGcpCustomerInfo sets GcpCustomerInfo field to given value.

### HasGcpCustomerInfo

`func (o *OfferInfo) HasGcpCustomerInfo() bool`

HasGcpCustomerInfo returns a boolean if a field has been set.

### GetGcpDuration

`func (o *OfferInfo) GetGcpDuration() int32`

GetGcpDuration returns the GcpDuration field if non-nil, zero value otherwise.

### GetGcpDurationOk

`func (o *OfferInfo) GetGcpDurationOk() (*int32, bool)`

GetGcpDurationOk returns a tuple with the GcpDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpDuration

`func (o *OfferInfo) SetGcpDuration(v int32)`

SetGcpDuration sets GcpDuration field to given value.

### HasGcpDuration

`func (o *OfferInfo) HasGcpDuration() bool`

HasGcpDuration returns a boolean if a field has been set.

### GetGcpMetrics

`func (o *OfferInfo) GetGcpMetrics() []GcpMarketplaceProductMeteringMetric`

GetGcpMetrics returns the GcpMetrics field if non-nil, zero value otherwise.

### GetGcpMetricsOk

`func (o *OfferInfo) GetGcpMetricsOk() (*[]GcpMarketplaceProductMeteringMetric, bool)`

GetGcpMetricsOk returns a tuple with the GcpMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpMetrics

`func (o *OfferInfo) SetGcpMetrics(v []GcpMarketplaceProductMeteringMetric)`

SetGcpMetrics sets GcpMetrics field to given value.

### HasGcpMetrics

`func (o *OfferInfo) HasGcpMetrics() bool`

HasGcpMetrics returns a boolean if a field has been set.

### GetGcpPaymentSchedule

`func (o *OfferInfo) GetGcpPaymentSchedule() GcpMarketplacePaymentScheduleType`

GetGcpPaymentSchedule returns the GcpPaymentSchedule field if non-nil, zero value otherwise.

### GetGcpPaymentScheduleOk

`func (o *OfferInfo) GetGcpPaymentScheduleOk() (*GcpMarketplacePaymentScheduleType, bool)`

GetGcpPaymentScheduleOk returns a tuple with the GcpPaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpPaymentSchedule

`func (o *OfferInfo) SetGcpPaymentSchedule(v GcpMarketplacePaymentScheduleType)`

SetGcpPaymentSchedule sets GcpPaymentSchedule field to given value.

### HasGcpPaymentSchedule

`func (o *OfferInfo) HasGcpPaymentSchedule() bool`

HasGcpPaymentSchedule returns a boolean if a field has been set.

### GetGcpPlans

`func (o *OfferInfo) GetGcpPlans() []GcpMarketplaceProductPurchaseOptionSpec`

GetGcpPlans returns the GcpPlans field if non-nil, zero value otherwise.

### GetGcpPlansOk

`func (o *OfferInfo) GetGcpPlansOk() (*[]GcpMarketplaceProductPurchaseOptionSpec, bool)`

GetGcpPlansOk returns a tuple with the GcpPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpPlans

`func (o *OfferInfo) SetGcpPlans(v []GcpMarketplaceProductPurchaseOptionSpec)`

SetGcpPlans sets GcpPlans field to given value.

### HasGcpPlans

`func (o *OfferInfo) HasGcpPlans() bool`

HasGcpPlans returns a boolean if a field has been set.

### GetGcpPrivateOffer

`func (o *OfferInfo) GetGcpPrivateOffer() GcpMarketplacePrivateOffer`

GetGcpPrivateOffer returns the GcpPrivateOffer field if non-nil, zero value otherwise.

### GetGcpPrivateOfferOk

`func (o *OfferInfo) GetGcpPrivateOfferOk() (*GcpMarketplacePrivateOffer, bool)`

GetGcpPrivateOfferOk returns a tuple with the GcpPrivateOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpPrivateOffer

`func (o *OfferInfo) SetGcpPrivateOffer(v GcpMarketplacePrivateOffer)`

SetGcpPrivateOffer sets GcpPrivateOffer field to given value.

### HasGcpPrivateOffer

`func (o *OfferInfo) HasGcpPrivateOffer() bool`

HasGcpPrivateOffer returns a boolean if a field has been set.

### GetGcpProviderInfo

`func (o *OfferInfo) GetGcpProviderInfo() GcpMarketplacePrivateOfferProviderInfo`

GetGcpProviderInfo returns the GcpProviderInfo field if non-nil, zero value otherwise.

### GetGcpProviderInfoOk

`func (o *OfferInfo) GetGcpProviderInfoOk() (*GcpMarketplacePrivateOfferProviderInfo, bool)`

GetGcpProviderInfoOk returns a tuple with the GcpProviderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProviderInfo

`func (o *OfferInfo) SetGcpProviderInfo(v GcpMarketplacePrivateOfferProviderInfo)`

SetGcpProviderInfo sets GcpProviderInfo field to given value.

### HasGcpProviderInfo

`func (o *OfferInfo) HasGcpProviderInfo() bool`

HasGcpProviderInfo returns a boolean if a field has been set.

### GetGcpProviderInternalNote

`func (o *OfferInfo) GetGcpProviderInternalNote() string`

GetGcpProviderInternalNote returns the GcpProviderInternalNote field if non-nil, zero value otherwise.

### GetGcpProviderInternalNoteOk

`func (o *OfferInfo) GetGcpProviderInternalNoteOk() (*string, bool)`

GetGcpProviderInternalNoteOk returns a tuple with the GcpProviderInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProviderInternalNote

`func (o *OfferInfo) SetGcpProviderInternalNote(v string)`

SetGcpProviderInternalNote sets GcpProviderInternalNote field to given value.

### HasGcpProviderInternalNote

`func (o *OfferInfo) HasGcpProviderInternalNote() bool`

HasGcpProviderInternalNote returns a boolean if a field has been set.

### GetGcpProviderPublicNote

`func (o *OfferInfo) GetGcpProviderPublicNote() string`

GetGcpProviderPublicNote returns the GcpProviderPublicNote field if non-nil, zero value otherwise.

### GetGcpProviderPublicNoteOk

`func (o *OfferInfo) GetGcpProviderPublicNoteOk() (*string, bool)`

GetGcpProviderPublicNoteOk returns a tuple with the GcpProviderPublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProviderPublicNote

`func (o *OfferInfo) SetGcpProviderPublicNote(v string)`

SetGcpProviderPublicNote sets GcpProviderPublicNote field to given value.

### HasGcpProviderPublicNote

`func (o *OfferInfo) HasGcpProviderPublicNote() bool`

HasGcpProviderPublicNote returns a boolean if a field has been set.

### GetGcpUsagePlanPriceModel

`func (o *OfferInfo) GetGcpUsagePlanPriceModel() GcpMarketplaceUsagePlanPriceModel`

GetGcpUsagePlanPriceModel returns the GcpUsagePlanPriceModel field if non-nil, zero value otherwise.

### GetGcpUsagePlanPriceModelOk

`func (o *OfferInfo) GetGcpUsagePlanPriceModelOk() (*GcpMarketplaceUsagePlanPriceModel, bool)`

GetGcpUsagePlanPriceModelOk returns a tuple with the GcpUsagePlanPriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpUsagePlanPriceModel

`func (o *OfferInfo) SetGcpUsagePlanPriceModel(v GcpMarketplaceUsagePlanPriceModel)`

SetGcpUsagePlanPriceModel sets GcpUsagePlanPriceModel field to given value.

### HasGcpUsagePlanPriceModel

`func (o *OfferInfo) HasGcpUsagePlanPriceModel() bool`

HasGcpUsagePlanPriceModel returns a boolean if a field has been set.

### GetPaymentInstallments

`func (o *OfferInfo) GetPaymentInstallments() []PaymentInstallment`

GetPaymentInstallments returns the PaymentInstallments field if non-nil, zero value otherwise.

### GetPaymentInstallmentsOk

`func (o *OfferInfo) GetPaymentInstallmentsOk() (*[]PaymentInstallment, bool)`

GetPaymentInstallmentsOk returns a tuple with the PaymentInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstallments

`func (o *OfferInfo) SetPaymentInstallments(v []PaymentInstallment)`

SetPaymentInstallments sets PaymentInstallments field to given value.

### HasPaymentInstallments

`func (o *OfferInfo) HasPaymentInstallments() bool`

HasPaymentInstallments returns a boolean if a field has been set.

### GetPrivateOfferUrl

`func (o *OfferInfo) GetPrivateOfferUrl() string`

GetPrivateOfferUrl returns the PrivateOfferUrl field if non-nil, zero value otherwise.

### GetPrivateOfferUrlOk

`func (o *OfferInfo) GetPrivateOfferUrlOk() (*string, bool)`

GetPrivateOfferUrlOk returns a tuple with the PrivateOfferUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOfferUrl

`func (o *OfferInfo) SetPrivateOfferUrl(v string)`

SetPrivateOfferUrl sets PrivateOfferUrl field to given value.

### HasPrivateOfferUrl

`func (o *OfferInfo) HasPrivateOfferUrl() bool`

HasPrivateOfferUrl returns a boolean if a field has been set.

### GetRefundCancelationPolicy

`func (o *OfferInfo) GetRefundCancelationPolicy() string`

GetRefundCancelationPolicy returns the RefundCancelationPolicy field if non-nil, zero value otherwise.

### GetRefundCancelationPolicyOk

`func (o *OfferInfo) GetRefundCancelationPolicyOk() (*string, bool)`

GetRefundCancelationPolicyOk returns a tuple with the RefundCancelationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCancelationPolicy

`func (o *OfferInfo) SetRefundCancelationPolicy(v string)`

SetRefundCancelationPolicy sets RefundCancelationPolicy field to given value.

### HasRefundCancelationPolicy

`func (o *OfferInfo) HasRefundCancelationPolicy() bool`

HasRefundCancelationPolicy returns a boolean if a field has been set.

### GetSellerNotes

`func (o *OfferInfo) GetSellerNotes() string`

GetSellerNotes returns the SellerNotes field if non-nil, zero value otherwise.

### GetSellerNotesOk

`func (o *OfferInfo) GetSellerNotesOk() (*string, bool)`

GetSellerNotesOk returns a tuple with the SellerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerNotes

`func (o *OfferInfo) SetSellerNotes(v string)`

SetSellerNotes sets SellerNotes field to given value.

### HasSellerNotes

`func (o *OfferInfo) HasSellerNotes() bool`

HasSellerNotes returns a boolean if a field has been set.

### GetVisibility

`func (o *OfferInfo) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OfferInfo) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OfferInfo) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OfferInfo) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


