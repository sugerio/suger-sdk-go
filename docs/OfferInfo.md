# OfferInfo

## Properties

 Name                              | Type                                                                                                                           | Description                                                                                                                                                                                                                                                                                                                         | Notes      
-----------------------------------|--------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **AdditionalEulaUrls**            | Pointer to **[]string**                                                                                                        | The URL of the additional EULA files. Only applicable when EulaType &#x3D; CUSTOM. The additional EULA files will be attached to the EULA file in the EulaUrl, and form a single EULA file.                                                                                                                                         | [optional] 
 **AdditionalResellerEulaUrls**    | Pointer to **[]string**                                                                                                        | The URL of the additional reseller EULA files. Only applicable when ResellerEulaType &#x3D; CUSTOM.                                                                                                                                                                                                                                 | [optional] 
 **AttachEulaType**                | Pointer to [**EulaType**](EulaType.md)                                                                                         | Attach the standard EULA file to the CUSTOM EULA file. Only applicable when EulaType &#x3D; CUSTOM                                                                                                                                                                                                                                  | [optional] 
 **AutoRenew**                     | Pointer to **bool**                                                                                                            | Is this offer Auto Renew enabled.                                                                                                                                                                                                                                                                                                   | [optional] 
 **AwsAgreementDuration**          | Pointer to **string**                                                                                                          | Aws private subscription offer Usage duration. ISO8601 format. P300D means the contract Usage start date At acceptance, and with duration 300 days.                                                                                                                                                                                 | [optional] 
 **AwsChannelPartner**             | Pointer to [**AwsChannelPartner**](AwsChannelPartner.md)                                                                       | The AWS channel partner (reseller), only applicable for AWS Marketplace CPPO_OUT or CPPO offers.                                                                                                                                                                                                                                    | [optional] 
 **AwsCppoEventDetail**            | Pointer to [**AwsMarketplaceEventBridgeEventDetail**](AwsMarketplaceEventBridgeEventDetail.md)                                 | AWS EventBridge Event, only applicable for AWS Marketplace CPPO offers.                                                                                                                                                                                                                                                             | [optional] 
 **AwsCppoOpportunity**            | Pointer to [**AwsMarketplaceCppoOpportunity**](AwsMarketplaceCppoOpportunity.md)                                               | AWS CPPO Opportunity, only applicable for AWS Marketplace CPPO_OUT or CPPO_IN offers.                                                                                                                                                                                                                                               | [optional] 
 **AwsMarkupPercentage**           | Pointer to **float32**                                                                                                         | AWS private reseller offer using markup percentage. 10.0 represent 10% partner margin.                                                                                                                                                                                                                                              | [optional] 
 **AwsResaleAuthorizationId**      | Pointer to **string**                                                                                                          | AWS ResaleAuthorizationId(CPPO_IN offer id) for CPPO offers of the reseller.                                                                                                                                                                                                                                                        | [optional] 
 **AzureOriginalPlan**             | Pointer to [**AzureMarketplacePriceAndAvailabilityPrivateOfferPlan**](AzureMarketplacePriceAndAvailabilityPrivateOfferPlan.md) | The origin pricing of Azure plan. Only applicable for Azure Marketplace plans.                                                                                                                                                                                                                                                      | [optional] 
 **AzurePrivateOffer**             | Pointer to [**AzureMarketplacePrivateOffer**](AzureMarketplacePrivateOffer.md)                                                 | The private offer for Azure Marketplace. Only applicable for Azure Marketplace private offers.                                                                                                                                                                                                                                      | [optional] 
 **AzureProductVariant**           | Pointer to [**AzureProductVariant**](AzureProductVariant.md)                                                                   | For Azure marketplace only.                                                                                                                                                                                                                                                                                                         | [optional] 
 **BillableDimensions**            | Pointer to [**[]BillableDimension**](BillableDimension.md)                                                                     | Usage based metering dimensions based on Billable Metrics, managed by Suger only.                                                                                                                                                                                                                                                   | [optional] 
 **BillingCycle**                  | Pointer to [**BillingCycle**](BillingCycle.md)                                                                                 | Billing Cycle for the offer.                                                                                                                                                                                                                                                                                                        | [optional] 
 **BillingIntervalInMonths**       | Pointer to **int32**                                                                                                           | Billing interval in months for the offer.                                                                                                                                                                                                                                                                                           | [optional] 
 **BuyerAwsAccountIds**            | Pointer to **[]string**                                                                                                        | The buyers&#39; AWS Account IDs of this offer.                                                                                                                                                                                                                                                                                      | [optional] 
 **BuyerAzureTenants**             | Pointer to [**[]AzureAudience**](AzureAudience.md)                                                                             | The buyers&#39; Azure tenants of this offer.                                                                                                                                                                                                                                                                                        | [optional] 
 **CommitAmount**                  | Pointer to **float32**                                                                                                         | The amount that the buyer has committed to pay, before discount if applicable. It can be monthly commitment or total commitment. For frontend display or analysis purposes, not used for billing.                                                                                                                                   | [optional] 
 **CommitBillingIntervalInMonths** | Pointer to **int32**                                                                                                           | Deprecated: Use BillingIntervalInMonths instead.                                                                                                                                                                                                                                                                                    | [optional] 
 **Commits**                       | Pointer to [**[]CommitDimension**](CommitDimension.md)                                                                         | Recurring flat fee for the offer, managed by cloud marketplaces or Suger.                                                                                                                                                                                                                                                           | [optional] 
 **Currency**                      | Pointer to **string**                                                                                                          | The currency code of the offer. ISO 4217 format.                                                                                                                                                                                                                                                                                    | [optional] 
 **Dimensions**                    | Pointer to [**[]MeteringDimension**](MeteringDimension.md)                                                                     | Usage based metering dimensions defined on cloud marketplaces, managed by Cloud marketplaces only.                                                                                                                                                                                                                                  | [optional] 
 **DiscountPercentage**            | Pointer to **float32**                                                                                                         | The discount percentage off the original price. For example, 20 means 20% off. 0 means no discount. It can be discount off the commitment amount or discount off the usage price.                                                                                                                                                   | [optional] 
 **EulaType**                      | Pointer to [**EulaType**](EulaType.md)                                                                                         |                                                                                                                                                                                                                                                                                                                                     | [optional] 
 **EulaUrl**                       | Pointer to **string**                                                                                                          |                                                                                                                                                                                                                                                                                                                                     | [optional] 
 **GcpCustomerInfo**               | Pointer to [**GcpMarketplacePrivateOfferCustomerInfo**](GcpMarketplacePrivateOfferCustomerInfo.md)                             | Only required when creating GCP Marketplace private offer.                                                                                                                                                                                                                                                                          | [optional] 
 **GcpDuration**                   | Pointer to **int32**                                                                                                           | The duration of the offer in months. Only required when creating GCP Marketplace private offer.                                                                                                                                                                                                                                     | [optional] 
 **GcpMetrics**                    | Pointer to [**[]GcpMarketplaceProductMeteringMetric**](GcpMarketplaceProductMeteringMetric.md)                                 | Only applicable for GCP Marketplace Offers (the default or private offer)                                                                                                                                                                                                                                                           | [optional] 
 **GcpPaymentSchedule**            | Pointer to [**PaymentScheduleType**](PaymentScheduleType.md)                                                                   | Only required when creating GCP Marketplace private offer, to specify the payment schedule for the private offer. TODO: It will be deprecated in the future and replaced by PaymentSchedule.                                                                                                                                        | [optional] 
 **GcpPlans**                      | Pointer to [**[]GcpMarketplaceProductPurchaseOptionSpec**](GcpMarketplaceProductPurchaseOptionSpec.md)                         | Only applicable for GCP Marketplace                                                                                                                                                                                                                                                                                                 | [optional] 
 **GcpPrivateOffer**               | Pointer to [**GcpMarketplacePrivateOffer**](GcpMarketplacePrivateOffer.md)                                                     | The private offer for GCP Marketplace. Only applicable for GCP Marketplace private offers.                                                                                                                                                                                                                                          | [optional] 
 **GcpProviderInfo**               | Pointer to [**GcpMarketplacePrivateOfferProviderInfo**](GcpMarketplacePrivateOfferProviderInfo.md)                             | Only required when creating GCP Marketplace private offer.                                                                                                                                                                                                                                                                          | [optional] 
 **GcpProviderInternalNote**       | Pointer to **string**                                                                                                          | Optional when creating GCP Marketplace private offer. The internal note for the seller/ISV. It is only visible to the seller/ISV.                                                                                                                                                                                                   | [optional] 
 **GcpProviderPublicNote**         | Pointer to **string**                                                                                                          | Optional when creating GCP Marketplace private offer. By default, it is the same as offer name. The public note for the buyer. It is visible to the buyer.                                                                                                                                                                          | [optional] 
 **GcpResellerPrivateOfferPlan**   | Pointer to [**GcpMarketplaceResellerPrivateOfferPlan**](GcpMarketplaceResellerPrivateOfferPlan.md)                             |                                                                                                                                                                                                                                                                                                                                     | [optional] 
 **GcpUsagePlanPriceModel**        | Pointer to [**GcpMarketplaceUsagePlanPriceModel**](GcpMarketplaceUsagePlanPriceModel.md)                                       | Only applicable for GCP Marketplace with Usage plan. Not appliable for Subscription plan.                                                                                                                                                                                                                                           | [optional] 
 **GracePeriodInDays**             | Pointer to **int32**                                                                                                           | The grace period in days for the offer. This is the number of days during which invoices remain in draft status, for reviewing. This filed can be overridden at the entitlement level.                                                                                                                                              | [optional] 
 **IsMeteringOverageCommit**       | Pointer to **bool**                                                                                                            | Whether the usage metering will only be charged for the amount that exceeds the committed amount. e.g. the buyer has committed $100, and the usage is $120, - if true, the buyer will be charged for the usage at $20, and the commit at $100. - if false, the buyer will be charged for the usage at $120, and the commit at $100. | [optional] 
 **NetTermsInDays**                | Pointer to **int32**                                                                                                           | The net terms in days for the offer. This is the number of days the buyer has to pay the invoice. This filed can be overridden at the entitlement level.                                                                                                                                                                            | [optional] 
 **PaymentInstallments**           | Pointer to [**[]PaymentInstallment**](PaymentInstallment.md)                                                                   | For flexible payment schedule, managed by cloud marketplaces or Suger.                                                                                                                                                                                                                                                              | [optional] 
 **PaymentSchedule**               | Pointer to [**PaymentScheduleType**](PaymentScheduleType.md)                                                                   | The payment schedule for the offer. PREPAY means the buyer pays before the service is provided. POSTPAY means the buyer pays after the service is provided.                                                                                                                                                                         | [optional] 
 **PrivateOfferUrl**               | Pointer to **string**                                                                                                          | The URL of the private offer sent to buyers to accept. Only applicable for private offer.                                                                                                                                                                                                                                           | [optional] 
 **ProratedBilling**               | Pointer to **bool**                                                                                                            | Prorated billing for the offer. If true, the billing is prorated based on the start date and end date. If false, the billing is not prorated. This filed can be overridden at the entitlement level.                                                                                                                                | [optional] 
 **RefundCancellationPolicy**      | Pointer to **string**                                                                                                          |                                                                                                                                                                                                                                                                                                                                     | [optional] 
 **ResellerAttachEulaType**        | Pointer to [**EulaType**](EulaType.md)                                                                                         | Attach the standard EULA file to the CUSTOM EULA file. Only applicable when EulaType &#x3D; CUSTOM                                                                                                                                                                                                                                  | [optional] 
 **ResellerEulaType**              | Pointer to [**EulaType**](EulaType.md)                                                                                         | The type of the reseller EULA. Only applicable for CPPO offers.                                                                                                                                                                                                                                                                     | [optional] 
 **ResellerEulaUrl**               | Pointer to **string**                                                                                                          |                                                                                                                                                                                                                                                                                                                                     | [optional] 
 **SellerNotes**                   | Pointer to **string**                                                                                                          |                                                                                                                                                                                                                                                                                                                                     | [optional] 
 **StartTime**                     | Pointer to **time.Time**                                                                                                       | Optional when creating AWS or GCP Marketplace private offer on the contract product. The future start time of the offer if it is not started on the acceptance.                                                                                                                                                                     | [optional] 
 **TaxIds**                        | Pointer to **[]string**                                                                                                        | Tax ids for the offer, used to calculate the tax amount for the offer. This field can be overridden at the entitlement level.                                                                                                                                                                                                       | [optional] 
 **TrialConfig**                   | Pointer to [**TrialConfig**](TrialConfig.md)                                                                                   | The offer for Direct. Only applicable for Direct offers. It is used in Stripe, Adyen, and other direct payment providers. The trial configuration for the offer.                                                                                                                                                                    | [optional] 
 **UsageBillingIntervalInMonths**  | Pointer to **int32**                                                                                                           | Deprecated: Use BillingIntervalInMonths instead.                                                                                                                                                                                                                                                                                    | [optional] 
 **Visibility**                    | Pointer to **string**                                                                                                          | The default visibility of offer is PRIVATE.                                                                                                                                                                                                                                                                                         | [optional] 

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

### GetAdditionalEulaUrls

`func (o *OfferInfo) GetAdditionalEulaUrls() []string`

GetAdditionalEulaUrls returns the AdditionalEulaUrls field if non-nil, zero value otherwise.

### GetAdditionalEulaUrlsOk

`func (o *OfferInfo) GetAdditionalEulaUrlsOk() (*[]string, bool)`

GetAdditionalEulaUrlsOk returns a tuple with the AdditionalEulaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEulaUrls

`func (o *OfferInfo) SetAdditionalEulaUrls(v []string)`

SetAdditionalEulaUrls sets AdditionalEulaUrls field to given value.

### HasAdditionalEulaUrls

`func (o *OfferInfo) HasAdditionalEulaUrls() bool`

HasAdditionalEulaUrls returns a boolean if a field has been set.

### GetAdditionalResellerEulaUrls

`func (o *OfferInfo) GetAdditionalResellerEulaUrls() []string`

GetAdditionalResellerEulaUrls returns the AdditionalResellerEulaUrls field if non-nil, zero value otherwise.

### GetAdditionalResellerEulaUrlsOk

`func (o *OfferInfo) GetAdditionalResellerEulaUrlsOk() (*[]string, bool)`

GetAdditionalResellerEulaUrlsOk returns a tuple with the AdditionalResellerEulaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResellerEulaUrls

`func (o *OfferInfo) SetAdditionalResellerEulaUrls(v []string)`

SetAdditionalResellerEulaUrls sets AdditionalResellerEulaUrls field to given value.

### HasAdditionalResellerEulaUrls

`func (o *OfferInfo) HasAdditionalResellerEulaUrls() bool`

HasAdditionalResellerEulaUrls returns a boolean if a field has been set.

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

### GetAwsAgreementDuration

`func (o *OfferInfo) GetAwsAgreementDuration() string`

GetAwsAgreementDuration returns the AwsAgreementDuration field if non-nil, zero value otherwise.

### GetAwsAgreementDurationOk

`func (o *OfferInfo) GetAwsAgreementDurationOk() (*string, bool)`

GetAwsAgreementDurationOk returns a tuple with the AwsAgreementDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAgreementDuration

`func (o *OfferInfo) SetAwsAgreementDuration(v string)`

SetAwsAgreementDuration sets AwsAgreementDuration field to given value.

### HasAwsAgreementDuration

`func (o *OfferInfo) HasAwsAgreementDuration() bool`

HasAwsAgreementDuration returns a boolean if a field has been set.

### GetAwsChannelPartner

`func (o *OfferInfo) GetAwsChannelPartner() AwsChannelPartner`

GetAwsChannelPartner returns the AwsChannelPartner field if non-nil, zero value otherwise.

### GetAwsChannelPartnerOk

`func (o *OfferInfo) GetAwsChannelPartnerOk() (*AwsChannelPartner, bool)`

GetAwsChannelPartnerOk returns a tuple with the AwsChannelPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsChannelPartner

`func (o *OfferInfo) SetAwsChannelPartner(v AwsChannelPartner)`

SetAwsChannelPartner sets AwsChannelPartner field to given value.

### HasAwsChannelPartner

`func (o *OfferInfo) HasAwsChannelPartner() bool`

HasAwsChannelPartner returns a boolean if a field has been set.

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

### GetAwsMarkupPercentage

`func (o *OfferInfo) GetAwsMarkupPercentage() float32`

GetAwsMarkupPercentage returns the AwsMarkupPercentage field if non-nil, zero value otherwise.

### GetAwsMarkupPercentageOk

`func (o *OfferInfo) GetAwsMarkupPercentageOk() (*float32, bool)`

GetAwsMarkupPercentageOk returns a tuple with the AwsMarkupPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsMarkupPercentage

`func (o *OfferInfo) SetAwsMarkupPercentage(v float32)`

SetAwsMarkupPercentage sets AwsMarkupPercentage field to given value.

### HasAwsMarkupPercentage

`func (o *OfferInfo) HasAwsMarkupPercentage() bool`

HasAwsMarkupPercentage returns a boolean if a field has been set.

### GetAwsResaleAuthorizationId

`func (o *OfferInfo) GetAwsResaleAuthorizationId() string`

GetAwsResaleAuthorizationId returns the AwsResaleAuthorizationId field if non-nil, zero value otherwise.

### GetAwsResaleAuthorizationIdOk

`func (o *OfferInfo) GetAwsResaleAuthorizationIdOk() (*string, bool)`

GetAwsResaleAuthorizationIdOk returns a tuple with the AwsResaleAuthorizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsResaleAuthorizationId

`func (o *OfferInfo) SetAwsResaleAuthorizationId(v string)`

SetAwsResaleAuthorizationId sets AwsResaleAuthorizationId field to given value.

### HasAwsResaleAuthorizationId

`func (o *OfferInfo) HasAwsResaleAuthorizationId() bool`

HasAwsResaleAuthorizationId returns a boolean if a field has been set.

### GetAzureOriginalPlan

`func (o *OfferInfo) GetAzureOriginalPlan() AzureMarketplacePriceAndAvailabilityPrivateOfferPlan`

GetAzureOriginalPlan returns the AzureOriginalPlan field if non-nil, zero value otherwise.

### GetAzureOriginalPlanOk

`func (o *OfferInfo) GetAzureOriginalPlanOk() (*AzureMarketplacePriceAndAvailabilityPrivateOfferPlan, bool)`

GetAzureOriginalPlanOk returns a tuple with the AzureOriginalPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureOriginalPlan

`func (o *OfferInfo) SetAzureOriginalPlan(v AzureMarketplacePriceAndAvailabilityPrivateOfferPlan)`

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

### GetBillableDimensions

`func (o *OfferInfo) GetBillableDimensions() []BillableDimension`

GetBillableDimensions returns the BillableDimensions field if non-nil, zero value otherwise.

### GetBillableDimensionsOk

`func (o *OfferInfo) GetBillableDimensionsOk() (*[]BillableDimension, bool)`

GetBillableDimensionsOk returns a tuple with the BillableDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableDimensions

`func (o *OfferInfo) SetBillableDimensions(v []BillableDimension)`

SetBillableDimensions sets BillableDimensions field to given value.

### HasBillableDimensions

`func (o *OfferInfo) HasBillableDimensions() bool`

HasBillableDimensions returns a boolean if a field has been set.

### GetBillingCycle

`func (o *OfferInfo) GetBillingCycle() BillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *OfferInfo) GetBillingCycleOk() (*BillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *OfferInfo) SetBillingCycle(v BillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *OfferInfo) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBillingIntervalInMonths

`func (o *OfferInfo) GetBillingIntervalInMonths() int32`

GetBillingIntervalInMonths returns the BillingIntervalInMonths field if non-nil, zero value otherwise.

### GetBillingIntervalInMonthsOk

`func (o *OfferInfo) GetBillingIntervalInMonthsOk() (*int32, bool)`

GetBillingIntervalInMonthsOk returns a tuple with the BillingIntervalInMonths field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetBillingIntervalInMonths

`func (o *OfferInfo) SetBillingIntervalInMonths(v int32)`

SetBillingIntervalInMonths sets BillingIntervalInMonths field to given value.

### HasBillingIntervalInMonths

`func (o *OfferInfo) HasBillingIntervalInMonths() bool`

HasBillingIntervalInMonths returns a boolean if a field has been set.

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

### GetCommitBillingIntervalInMonths

`func (o *OfferInfo) GetCommitBillingIntervalInMonths() int32`

GetCommitBillingIntervalInMonths returns the CommitBillingIntervalInMonths field if non-nil, zero value otherwise.

### GetCommitBillingIntervalInMonthsOk

`func (o *OfferInfo) GetCommitBillingIntervalInMonthsOk() (*int32, bool)`

GetCommitBillingIntervalInMonthsOk returns a tuple with the CommitBillingIntervalInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitBillingIntervalInMonths

`func (o *OfferInfo) SetCommitBillingIntervalInMonths(v int32)`

SetCommitBillingIntervalInMonths sets CommitBillingIntervalInMonths field to given value.

### HasCommitBillingIntervalInMonths

`func (o *OfferInfo) HasCommitBillingIntervalInMonths() bool`

HasCommitBillingIntervalInMonths returns a boolean if a field has been set.

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

`func (o *OfferInfo) GetGcpPaymentSchedule() PaymentScheduleType`

GetGcpPaymentSchedule returns the GcpPaymentSchedule field if non-nil, zero value otherwise.

### GetGcpPaymentScheduleOk

`func (o *OfferInfo) GetGcpPaymentScheduleOk() (*PaymentScheduleType, bool)`

GetGcpPaymentScheduleOk returns a tuple with the GcpPaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpPaymentSchedule

`func (o *OfferInfo) SetGcpPaymentSchedule(v PaymentScheduleType)`

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

### GetGcpResellerPrivateOfferPlan

`func (o *OfferInfo) GetGcpResellerPrivateOfferPlan() GcpMarketplaceResellerPrivateOfferPlan`

GetGcpResellerPrivateOfferPlan returns the GcpResellerPrivateOfferPlan field if non-nil, zero value otherwise.

### GetGcpResellerPrivateOfferPlanOk

`func (o *OfferInfo) GetGcpResellerPrivateOfferPlanOk() (*GcpMarketplaceResellerPrivateOfferPlan, bool)`

GetGcpResellerPrivateOfferPlanOk returns a tuple with the GcpResellerPrivateOfferPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpResellerPrivateOfferPlan

`func (o *OfferInfo) SetGcpResellerPrivateOfferPlan(v GcpMarketplaceResellerPrivateOfferPlan)`

SetGcpResellerPrivateOfferPlan sets GcpResellerPrivateOfferPlan field to given value.

### HasGcpResellerPrivateOfferPlan

`func (o *OfferInfo) HasGcpResellerPrivateOfferPlan() bool`

HasGcpResellerPrivateOfferPlan returns a boolean if a field has been set.

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

### GetGracePeriodInDays

`func (o *OfferInfo) GetGracePeriodInDays() int32`

GetGracePeriodInDays returns the GracePeriodInDays field if non-nil, zero value otherwise.

### GetGracePeriodInDaysOk

`func (o *OfferInfo) GetGracePeriodInDaysOk() (*int32, bool)`

GetGracePeriodInDaysOk returns a tuple with the GracePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodInDays

`func (o *OfferInfo) SetGracePeriodInDays(v int32)`

SetGracePeriodInDays sets GracePeriodInDays field to given value.

### HasGracePeriodInDays

`func (o *OfferInfo) HasGracePeriodInDays() bool`

HasGracePeriodInDays returns a boolean if a field has been set.

### GetIsMeteringOverageCommit

`func (o *OfferInfo) GetIsMeteringOverageCommit() bool`

GetIsMeteringOverageCommit returns the IsMeteringOverageCommit field if non-nil, zero value otherwise.

### GetIsMeteringOverageCommitOk

`func (o *OfferInfo) GetIsMeteringOverageCommitOk() (*bool, bool)`

GetIsMeteringOverageCommitOk returns a tuple with the IsMeteringOverageCommit field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetIsMeteringOverageCommit

`func (o *OfferInfo) SetIsMeteringOverageCommit(v bool)`

SetIsMeteringOverageCommit sets IsMeteringOverageCommit field to given value.

### HasIsMeteringOverageCommit

`func (o *OfferInfo) HasIsMeteringOverageCommit() bool`

HasIsMeteringOverageCommit returns a boolean if a field has been set.

### GetNetTermsInDays

`func (o *OfferInfo) GetNetTermsInDays() int32`

GetNetTermsInDays returns the NetTermsInDays field if non-nil, zero value otherwise.

### GetNetTermsInDaysOk

`func (o *OfferInfo) GetNetTermsInDaysOk() (*int32, bool)`

GetNetTermsInDaysOk returns a tuple with the NetTermsInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTermsInDays

`func (o *OfferInfo) SetNetTermsInDays(v int32)`

SetNetTermsInDays sets NetTermsInDays field to given value.

### HasNetTermsInDays

`func (o *OfferInfo) HasNetTermsInDays() bool`

HasNetTermsInDays returns a boolean if a field has been set.

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

### GetPaymentSchedule

`func (o *OfferInfo) GetPaymentSchedule() PaymentScheduleType`

GetPaymentSchedule returns the PaymentSchedule field if non-nil, zero value otherwise.

### GetPaymentScheduleOk

`func (o *OfferInfo) GetPaymentScheduleOk() (*PaymentScheduleType, bool)`

GetPaymentScheduleOk returns a tuple with the PaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedule

`func (o *OfferInfo) SetPaymentSchedule(v PaymentScheduleType)`

SetPaymentSchedule sets PaymentSchedule field to given value.

### HasPaymentSchedule

`func (o *OfferInfo) HasPaymentSchedule() bool`

HasPaymentSchedule returns a boolean if a field has been set.

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

### GetProratedBilling

`func (o *OfferInfo) GetProratedBilling() bool`

GetProratedBilling returns the ProratedBilling field if non-nil, zero value otherwise.

### GetProratedBillingOk

`func (o *OfferInfo) GetProratedBillingOk() (*bool, bool)`

GetProratedBillingOk returns a tuple with the ProratedBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratedBilling

`func (o *OfferInfo) SetProratedBilling(v bool)`

SetProratedBilling sets ProratedBilling field to given value.

### HasProratedBilling

`func (o *OfferInfo) HasProratedBilling() bool`

HasProratedBilling returns a boolean if a field has been set.

### GetRefundCancellationPolicy

`func (o *OfferInfo) GetRefundCancellationPolicy() string`

GetRefundCancellationPolicy returns the RefundCancellationPolicy field if non-nil, zero value otherwise.

### GetRefundCancellationPolicyOk

`func (o *OfferInfo) GetRefundCancellationPolicyOk() (*string, bool)`

GetRefundCancellationPolicyOk returns a tuple with the RefundCancellationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCancellationPolicy

`func (o *OfferInfo) SetRefundCancellationPolicy(v string)`

SetRefundCancellationPolicy sets RefundCancellationPolicy field to given value.

### HasRefundCancellationPolicy

`func (o *OfferInfo) HasRefundCancellationPolicy() bool`

HasRefundCancellationPolicy returns a boolean if a field has been set.

### GetResellerAttachEulaType

`func (o *OfferInfo) GetResellerAttachEulaType() EulaType`

GetResellerAttachEulaType returns the ResellerAttachEulaType field if non-nil, zero value otherwise.

### GetResellerAttachEulaTypeOk

`func (o *OfferInfo) GetResellerAttachEulaTypeOk() (*EulaType, bool)`

GetResellerAttachEulaTypeOk returns a tuple with the ResellerAttachEulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerAttachEulaType

`func (o *OfferInfo) SetResellerAttachEulaType(v EulaType)`

SetResellerAttachEulaType sets ResellerAttachEulaType field to given value.

### HasResellerAttachEulaType

`func (o *OfferInfo) HasResellerAttachEulaType() bool`

HasResellerAttachEulaType returns a boolean if a field has been set.

### GetResellerEulaType

`func (o *OfferInfo) GetResellerEulaType() EulaType`

GetResellerEulaType returns the ResellerEulaType field if non-nil, zero value otherwise.

### GetResellerEulaTypeOk

`func (o *OfferInfo) GetResellerEulaTypeOk() (*EulaType, bool)`

GetResellerEulaTypeOk returns a tuple with the ResellerEulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerEulaType

`func (o *OfferInfo) SetResellerEulaType(v EulaType)`

SetResellerEulaType sets ResellerEulaType field to given value.

### HasResellerEulaType

`func (o *OfferInfo) HasResellerEulaType() bool`

HasResellerEulaType returns a boolean if a field has been set.

### GetResellerEulaUrl

`func (o *OfferInfo) GetResellerEulaUrl() string`

GetResellerEulaUrl returns the ResellerEulaUrl field if non-nil, zero value otherwise.

### GetResellerEulaUrlOk

`func (o *OfferInfo) GetResellerEulaUrlOk() (*string, bool)`

GetResellerEulaUrlOk returns a tuple with the ResellerEulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerEulaUrl

`func (o *OfferInfo) SetResellerEulaUrl(v string)`

SetResellerEulaUrl sets ResellerEulaUrl field to given value.

### HasResellerEulaUrl

`func (o *OfferInfo) HasResellerEulaUrl() bool`

HasResellerEulaUrl returns a boolean if a field has been set.

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

### GetStartTime

`func (o *OfferInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *OfferInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *OfferInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *OfferInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTaxIds

`func (o *OfferInfo) GetTaxIds() []string`

GetTaxIds returns the TaxIds field if non-nil, zero value otherwise.

### GetTaxIdsOk

`func (o *OfferInfo) GetTaxIdsOk() (*[]string, bool)`

GetTaxIdsOk returns a tuple with the TaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIds

`func (o *OfferInfo) SetTaxIds(v []string)`

SetTaxIds sets TaxIds field to given value.

### HasTaxIds

`func (o *OfferInfo) HasTaxIds() bool`

HasTaxIds returns a boolean if a field has been set.

### GetTrialConfig

`func (o *OfferInfo) GetTrialConfig() TrialConfig`

GetTrialConfig returns the TrialConfig field if non-nil, zero value otherwise.

### GetTrialConfigOk

`func (o *OfferInfo) GetTrialConfigOk() (*TrialConfig, bool)`

GetTrialConfigOk returns a tuple with the TrialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialConfig

`func (o *OfferInfo) SetTrialConfig(v TrialConfig)`

SetTrialConfig sets TrialConfig field to given value.

### HasTrialConfig

`func (o *OfferInfo) HasTrialConfig() bool`

HasTrialConfig returns a boolean if a field has been set.

### GetUsageBillingIntervalInMonths

`func (o *OfferInfo) GetUsageBillingIntervalInMonths() int32`

GetUsageBillingIntervalInMonths returns the UsageBillingIntervalInMonths field if non-nil, zero value otherwise.

### GetUsageBillingIntervalInMonthsOk

`func (o *OfferInfo) GetUsageBillingIntervalInMonthsOk() (*int32, bool)`

GetUsageBillingIntervalInMonthsOk returns a tuple with the UsageBillingIntervalInMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBillingIntervalInMonths

`func (o *OfferInfo) SetUsageBillingIntervalInMonths(v int32)`

SetUsageBillingIntervalInMonths sets UsageBillingIntervalInMonths field to given value.

### HasUsageBillingIntervalInMonths

`func (o *OfferInfo) HasUsageBillingIntervalInMonths() bool`

HasUsageBillingIntervalInMonths returns a boolean if a field has been set.

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


