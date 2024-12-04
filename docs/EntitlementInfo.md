# EntitlementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**[]BillingAddonRecord**](BillingAddonRecord.md) | The addons for the entitlement. | [optional] 
**AlertDaysBeforeEnd** | Pointer to **int32** | Alert days before the end of the entitlement | [optional] 
**AlibabaEntitlements** | Pointer to [**[]ClientDescribeInstanceResponseBody**](ClientDescribeInstanceResponseBody.md) | Nullable. Alibaba Entitlements from Alibaba Marketplace. | [optional] 
**AlibabaOrders** | Pointer to [**[]ClientDescribeOrderResponseBody**](ClientDescribeOrderResponseBody.md) | Nullable. Alibaba Orders from Alibaba Marketplace. | [optional] 
**AutoRenew** | Pointer to **bool** | Is this Entitlement Auto Renew enabled. | [optional] 
**AwsAgreement** | Pointer to [**AwsMarketplaceAgreementV2**](AwsMarketplaceAgreementV2.md) | Nullable. AWS agreement from AWS Marketplace. | [optional] 
**AwsChannelPartner** | Pointer to [**AwsChannelPartner**](AwsChannelPartner.md) | The AWS channel partner (reseller), only applicable if this entitlement is based on AWS CPPO offer. | [optional] 
**AwsEntitlements** | Pointer to [**[]TypesEntitlement**](TypesEntitlement.md) | Nullable. AWS Entitlements from AWS Marketplace. | [optional] 
**AzureSubscriptions** | Pointer to [**[]AzureMarketplaceSubscription**](AzureMarketplaceSubscription.md) | Nullable. Azure Subscriptions from Azure Marketplace. | [optional] 
**BillableDimensions** | Pointer to [**[]BillableDimension**](BillableDimension.md) | The dimensions for billable metric usage-based metering. It&#39;s for Suger(Stripe, Ayden) metering. | [optional] 
**BillingCycle** | Pointer to [**BillingCycle**](BillingCycle.md) | Billing Cycle | [optional] 
**CollectableAmount** | Pointer to **float32** | The amount that the seller can collect. It excludes the marketplace commision fee. | [optional] 
**CommitAmount** | Pointer to **float32** | The amount that the buyer has committed to pay. It can be the sum of payment installments if applicable. | [optional] 
**Commits** | Pointer to [**[]CommitDimension**](CommitDimension.md) | The dimensions for flatrate commitment (recurring or one-time). | [optional] 
**Currency** | Pointer to **string** | The default Currency is USD. | [optional] 
**Dimensions** | Pointer to [**[]MeteringDimension**](MeteringDimension.md) | The dimensions for usage-based metering. It&#39;s for usage metering in cloud marketplaces. The max size of dimensions is 50. The oversized dimensions won&#39;t be saved in the EntitlementInfo. But the dimensions can be accessed from the connected offer info or product info. | [optional] 
**DimensionsOversized** | Pointer to **bool** | Whether the upper metering dimensions are oversized (exceed the max size 50). | [optional] 
**DisbursedAmount** | Pointer to **float32** | The amount that has been disbursed to the seller account. | [optional] 
**EulaType** | Pointer to [**EulaType**](EulaType.md) |  | [optional] 
**EulaUrl** | Pointer to **string** |  | [optional] 
**GcpEntitlements** | Pointer to [**[]GcpMarketplaceEntitlement**](GcpMarketplaceEntitlement.md) | Nullable. GCP Entitlements from GCP Marketplace. | [optional] 
**GcpPlans** | Pointer to [**[]GcpMarketplaceProductPurchaseOptionSpec**](GcpMarketplaceProductPurchaseOptionSpec.md) | Only applicable for GCP Marketplace Entitlements. | [optional] 
**GracePeriodInDays** | Pointer to **int32** | The grace period for the offer. It is same as the TrialConfig in DirectOfferInfo. But can be overridden at the entitlement level. | [optional] 
**GrossAmount** | Pointer to **float32** | The gross amount that the buyer has committed to pay, including usage metered amount. | [optional] 
**InvoicedAmount** | Pointer to **float32** | The amount that the buyer has got invoiced. | [optional] 
**NetTermsInDays** | Pointer to **int32** | The net terms for the offer. It is same as the TrialConfig in DirectOfferInfo. But can be overridden at the entitlement level. | [optional] 
**PaymentInstallments** | Pointer to [**[]PaymentInstallment**](PaymentInstallment.md) | For flexible payment schedules | [optional] 
**PaymentSchedule** | Pointer to [**PaymentScheduleType**](PaymentScheduleType.md) | The payment schedule for the entitlement. PREPAY means the buyer pays before the service is provided. POSTPAY means the buyer pays after the service is provided. | [optional] 
**RefundCancellationPolicy** | Pointer to **string** |  | [optional] 
**SellerNotes** | Pointer to **string** |  | [optional] 
**SpaUrl** | Pointer to **string** | The URL with JWT as auth method for the entitlement SPA. It can be shared with the buyer to access the SPA without login. | [optional] 
**TrialConfig** | Pointer to [**TrialConfig**](TrialConfig.md) | The trial configuration for the offer. It is same as the TrialConfig in DirectOfferInfo. But can be overridden at the entitlement level. | [optional] 

## Methods

### NewEntitlementInfo

`func NewEntitlementInfo() *EntitlementInfo`

NewEntitlementInfo instantiates a new EntitlementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementInfoWithDefaults

`func NewEntitlementInfoWithDefaults() *EntitlementInfo`

NewEntitlementInfoWithDefaults instantiates a new EntitlementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *EntitlementInfo) GetAddons() []BillingAddonRecord`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *EntitlementInfo) GetAddonsOk() (*[]BillingAddonRecord, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *EntitlementInfo) SetAddons(v []BillingAddonRecord)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *EntitlementInfo) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetAlertDaysBeforeEnd

`func (o *EntitlementInfo) GetAlertDaysBeforeEnd() int32`

GetAlertDaysBeforeEnd returns the AlertDaysBeforeEnd field if non-nil, zero value otherwise.

### GetAlertDaysBeforeEndOk

`func (o *EntitlementInfo) GetAlertDaysBeforeEndOk() (*int32, bool)`

GetAlertDaysBeforeEndOk returns a tuple with the AlertDaysBeforeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDaysBeforeEnd

`func (o *EntitlementInfo) SetAlertDaysBeforeEnd(v int32)`

SetAlertDaysBeforeEnd sets AlertDaysBeforeEnd field to given value.

### HasAlertDaysBeforeEnd

`func (o *EntitlementInfo) HasAlertDaysBeforeEnd() bool`

HasAlertDaysBeforeEnd returns a boolean if a field has been set.

### GetAlibabaEntitlements

`func (o *EntitlementInfo) GetAlibabaEntitlements() []ClientDescribeInstanceResponseBody`

GetAlibabaEntitlements returns the AlibabaEntitlements field if non-nil, zero value otherwise.

### GetAlibabaEntitlementsOk

`func (o *EntitlementInfo) GetAlibabaEntitlementsOk() (*[]ClientDescribeInstanceResponseBody, bool)`

GetAlibabaEntitlementsOk returns a tuple with the AlibabaEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaEntitlements

`func (o *EntitlementInfo) SetAlibabaEntitlements(v []ClientDescribeInstanceResponseBody)`

SetAlibabaEntitlements sets AlibabaEntitlements field to given value.

### HasAlibabaEntitlements

`func (o *EntitlementInfo) HasAlibabaEntitlements() bool`

HasAlibabaEntitlements returns a boolean if a field has been set.

### GetAlibabaOrders

`func (o *EntitlementInfo) GetAlibabaOrders() []ClientDescribeOrderResponseBody`

GetAlibabaOrders returns the AlibabaOrders field if non-nil, zero value otherwise.

### GetAlibabaOrdersOk

`func (o *EntitlementInfo) GetAlibabaOrdersOk() (*[]ClientDescribeOrderResponseBody, bool)`

GetAlibabaOrdersOk returns a tuple with the AlibabaOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaOrders

`func (o *EntitlementInfo) SetAlibabaOrders(v []ClientDescribeOrderResponseBody)`

SetAlibabaOrders sets AlibabaOrders field to given value.

### HasAlibabaOrders

`func (o *EntitlementInfo) HasAlibabaOrders() bool`

HasAlibabaOrders returns a boolean if a field has been set.

### GetAutoRenew

`func (o *EntitlementInfo) GetAutoRenew() bool`

GetAutoRenew returns the AutoRenew field if non-nil, zero value otherwise.

### GetAutoRenewOk

`func (o *EntitlementInfo) GetAutoRenewOk() (*bool, bool)`

GetAutoRenewOk returns a tuple with the AutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenew

`func (o *EntitlementInfo) SetAutoRenew(v bool)`

SetAutoRenew sets AutoRenew field to given value.

### HasAutoRenew

`func (o *EntitlementInfo) HasAutoRenew() bool`

HasAutoRenew returns a boolean if a field has been set.

### GetAwsAgreement

`func (o *EntitlementInfo) GetAwsAgreement() AwsMarketplaceAgreementV2`

GetAwsAgreement returns the AwsAgreement field if non-nil, zero value otherwise.

### GetAwsAgreementOk

`func (o *EntitlementInfo) GetAwsAgreementOk() (*AwsMarketplaceAgreementV2, bool)`

GetAwsAgreementOk returns a tuple with the AwsAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAgreement

`func (o *EntitlementInfo) SetAwsAgreement(v AwsMarketplaceAgreementV2)`

SetAwsAgreement sets AwsAgreement field to given value.

### HasAwsAgreement

`func (o *EntitlementInfo) HasAwsAgreement() bool`

HasAwsAgreement returns a boolean if a field has been set.

### GetAwsChannelPartner

`func (o *EntitlementInfo) GetAwsChannelPartner() AwsChannelPartner`

GetAwsChannelPartner returns the AwsChannelPartner field if non-nil, zero value otherwise.

### GetAwsChannelPartnerOk

`func (o *EntitlementInfo) GetAwsChannelPartnerOk() (*AwsChannelPartner, bool)`

GetAwsChannelPartnerOk returns a tuple with the AwsChannelPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsChannelPartner

`func (o *EntitlementInfo) SetAwsChannelPartner(v AwsChannelPartner)`

SetAwsChannelPartner sets AwsChannelPartner field to given value.

### HasAwsChannelPartner

`func (o *EntitlementInfo) HasAwsChannelPartner() bool`

HasAwsChannelPartner returns a boolean if a field has been set.

### GetAwsEntitlements

`func (o *EntitlementInfo) GetAwsEntitlements() []TypesEntitlement`

GetAwsEntitlements returns the AwsEntitlements field if non-nil, zero value otherwise.

### GetAwsEntitlementsOk

`func (o *EntitlementInfo) GetAwsEntitlementsOk() (*[]TypesEntitlement, bool)`

GetAwsEntitlementsOk returns a tuple with the AwsEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsEntitlements

`func (o *EntitlementInfo) SetAwsEntitlements(v []TypesEntitlement)`

SetAwsEntitlements sets AwsEntitlements field to given value.

### HasAwsEntitlements

`func (o *EntitlementInfo) HasAwsEntitlements() bool`

HasAwsEntitlements returns a boolean if a field has been set.

### GetAzureSubscriptions

`func (o *EntitlementInfo) GetAzureSubscriptions() []AzureMarketplaceSubscription`

GetAzureSubscriptions returns the AzureSubscriptions field if non-nil, zero value otherwise.

### GetAzureSubscriptionsOk

`func (o *EntitlementInfo) GetAzureSubscriptionsOk() (*[]AzureMarketplaceSubscription, bool)`

GetAzureSubscriptionsOk returns a tuple with the AzureSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptions

`func (o *EntitlementInfo) SetAzureSubscriptions(v []AzureMarketplaceSubscription)`

SetAzureSubscriptions sets AzureSubscriptions field to given value.

### HasAzureSubscriptions

`func (o *EntitlementInfo) HasAzureSubscriptions() bool`

HasAzureSubscriptions returns a boolean if a field has been set.

### GetBillableDimensions

`func (o *EntitlementInfo) GetBillableDimensions() []BillableDimension`

GetBillableDimensions returns the BillableDimensions field if non-nil, zero value otherwise.

### GetBillableDimensionsOk

`func (o *EntitlementInfo) GetBillableDimensionsOk() (*[]BillableDimension, bool)`

GetBillableDimensionsOk returns a tuple with the BillableDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableDimensions

`func (o *EntitlementInfo) SetBillableDimensions(v []BillableDimension)`

SetBillableDimensions sets BillableDimensions field to given value.

### HasBillableDimensions

`func (o *EntitlementInfo) HasBillableDimensions() bool`

HasBillableDimensions returns a boolean if a field has been set.

### GetBillingCycle

`func (o *EntitlementInfo) GetBillingCycle() BillingCycle`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *EntitlementInfo) GetBillingCycleOk() (*BillingCycle, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *EntitlementInfo) SetBillingCycle(v BillingCycle)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *EntitlementInfo) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCollectableAmount

`func (o *EntitlementInfo) GetCollectableAmount() float32`

GetCollectableAmount returns the CollectableAmount field if non-nil, zero value otherwise.

### GetCollectableAmountOk

`func (o *EntitlementInfo) GetCollectableAmountOk() (*float32, bool)`

GetCollectableAmountOk returns a tuple with the CollectableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectableAmount

`func (o *EntitlementInfo) SetCollectableAmount(v float32)`

SetCollectableAmount sets CollectableAmount field to given value.

### HasCollectableAmount

`func (o *EntitlementInfo) HasCollectableAmount() bool`

HasCollectableAmount returns a boolean if a field has been set.

### GetCommitAmount

`func (o *EntitlementInfo) GetCommitAmount() float32`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *EntitlementInfo) GetCommitAmountOk() (*float32, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *EntitlementInfo) SetCommitAmount(v float32)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *EntitlementInfo) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetCommits

`func (o *EntitlementInfo) GetCommits() []CommitDimension`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *EntitlementInfo) GetCommitsOk() (*[]CommitDimension, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *EntitlementInfo) SetCommits(v []CommitDimension)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *EntitlementInfo) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetCurrency

`func (o *EntitlementInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EntitlementInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EntitlementInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EntitlementInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDimensions

`func (o *EntitlementInfo) GetDimensions() []MeteringDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *EntitlementInfo) GetDimensionsOk() (*[]MeteringDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *EntitlementInfo) SetDimensions(v []MeteringDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *EntitlementInfo) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetDimensionsOversized

`func (o *EntitlementInfo) GetDimensionsOversized() bool`

GetDimensionsOversized returns the DimensionsOversized field if non-nil, zero value otherwise.

### GetDimensionsOversizedOk

`func (o *EntitlementInfo) GetDimensionsOversizedOk() (*bool, bool)`

GetDimensionsOversizedOk returns a tuple with the DimensionsOversized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionsOversized

`func (o *EntitlementInfo) SetDimensionsOversized(v bool)`

SetDimensionsOversized sets DimensionsOversized field to given value.

### HasDimensionsOversized

`func (o *EntitlementInfo) HasDimensionsOversized() bool`

HasDimensionsOversized returns a boolean if a field has been set.

### GetDisbursedAmount

`func (o *EntitlementInfo) GetDisbursedAmount() float32`

GetDisbursedAmount returns the DisbursedAmount field if non-nil, zero value otherwise.

### GetDisbursedAmountOk

`func (o *EntitlementInfo) GetDisbursedAmountOk() (*float32, bool)`

GetDisbursedAmountOk returns a tuple with the DisbursedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursedAmount

`func (o *EntitlementInfo) SetDisbursedAmount(v float32)`

SetDisbursedAmount sets DisbursedAmount field to given value.

### HasDisbursedAmount

`func (o *EntitlementInfo) HasDisbursedAmount() bool`

HasDisbursedAmount returns a boolean if a field has been set.

### GetEulaType

`func (o *EntitlementInfo) GetEulaType() EulaType`

GetEulaType returns the EulaType field if non-nil, zero value otherwise.

### GetEulaTypeOk

`func (o *EntitlementInfo) GetEulaTypeOk() (*EulaType, bool)`

GetEulaTypeOk returns a tuple with the EulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaType

`func (o *EntitlementInfo) SetEulaType(v EulaType)`

SetEulaType sets EulaType field to given value.

### HasEulaType

`func (o *EntitlementInfo) HasEulaType() bool`

HasEulaType returns a boolean if a field has been set.

### GetEulaUrl

`func (o *EntitlementInfo) GetEulaUrl() string`

GetEulaUrl returns the EulaUrl field if non-nil, zero value otherwise.

### GetEulaUrlOk

`func (o *EntitlementInfo) GetEulaUrlOk() (*string, bool)`

GetEulaUrlOk returns a tuple with the EulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaUrl

`func (o *EntitlementInfo) SetEulaUrl(v string)`

SetEulaUrl sets EulaUrl field to given value.

### HasEulaUrl

`func (o *EntitlementInfo) HasEulaUrl() bool`

HasEulaUrl returns a boolean if a field has been set.

### GetGcpEntitlements

`func (o *EntitlementInfo) GetGcpEntitlements() []GcpMarketplaceEntitlement`

GetGcpEntitlements returns the GcpEntitlements field if non-nil, zero value otherwise.

### GetGcpEntitlementsOk

`func (o *EntitlementInfo) GetGcpEntitlementsOk() (*[]GcpMarketplaceEntitlement, bool)`

GetGcpEntitlementsOk returns a tuple with the GcpEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpEntitlements

`func (o *EntitlementInfo) SetGcpEntitlements(v []GcpMarketplaceEntitlement)`

SetGcpEntitlements sets GcpEntitlements field to given value.

### HasGcpEntitlements

`func (o *EntitlementInfo) HasGcpEntitlements() bool`

HasGcpEntitlements returns a boolean if a field has been set.

### GetGcpPlans

`func (o *EntitlementInfo) GetGcpPlans() []GcpMarketplaceProductPurchaseOptionSpec`

GetGcpPlans returns the GcpPlans field if non-nil, zero value otherwise.

### GetGcpPlansOk

`func (o *EntitlementInfo) GetGcpPlansOk() (*[]GcpMarketplaceProductPurchaseOptionSpec, bool)`

GetGcpPlansOk returns a tuple with the GcpPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpPlans

`func (o *EntitlementInfo) SetGcpPlans(v []GcpMarketplaceProductPurchaseOptionSpec)`

SetGcpPlans sets GcpPlans field to given value.

### HasGcpPlans

`func (o *EntitlementInfo) HasGcpPlans() bool`

HasGcpPlans returns a boolean if a field has been set.

### GetGracePeriodInDays

`func (o *EntitlementInfo) GetGracePeriodInDays() int32`

GetGracePeriodInDays returns the GracePeriodInDays field if non-nil, zero value otherwise.

### GetGracePeriodInDaysOk

`func (o *EntitlementInfo) GetGracePeriodInDaysOk() (*int32, bool)`

GetGracePeriodInDaysOk returns a tuple with the GracePeriodInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodInDays

`func (o *EntitlementInfo) SetGracePeriodInDays(v int32)`

SetGracePeriodInDays sets GracePeriodInDays field to given value.

### HasGracePeriodInDays

`func (o *EntitlementInfo) HasGracePeriodInDays() bool`

HasGracePeriodInDays returns a boolean if a field has been set.

### GetGrossAmount

`func (o *EntitlementInfo) GetGrossAmount() float32`

GetGrossAmount returns the GrossAmount field if non-nil, zero value otherwise.

### GetGrossAmountOk

`func (o *EntitlementInfo) GetGrossAmountOk() (*float32, bool)`

GetGrossAmountOk returns a tuple with the GrossAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossAmount

`func (o *EntitlementInfo) SetGrossAmount(v float32)`

SetGrossAmount sets GrossAmount field to given value.

### HasGrossAmount

`func (o *EntitlementInfo) HasGrossAmount() bool`

HasGrossAmount returns a boolean if a field has been set.

### GetInvoicedAmount

`func (o *EntitlementInfo) GetInvoicedAmount() float32`

GetInvoicedAmount returns the InvoicedAmount field if non-nil, zero value otherwise.

### GetInvoicedAmountOk

`func (o *EntitlementInfo) GetInvoicedAmountOk() (*float32, bool)`

GetInvoicedAmountOk returns a tuple with the InvoicedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedAmount

`func (o *EntitlementInfo) SetInvoicedAmount(v float32)`

SetInvoicedAmount sets InvoicedAmount field to given value.

### HasInvoicedAmount

`func (o *EntitlementInfo) HasInvoicedAmount() bool`

HasInvoicedAmount returns a boolean if a field has been set.

### GetNetTermsInDays

`func (o *EntitlementInfo) GetNetTermsInDays() int32`

GetNetTermsInDays returns the NetTermsInDays field if non-nil, zero value otherwise.

### GetNetTermsInDaysOk

`func (o *EntitlementInfo) GetNetTermsInDaysOk() (*int32, bool)`

GetNetTermsInDaysOk returns a tuple with the NetTermsInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTermsInDays

`func (o *EntitlementInfo) SetNetTermsInDays(v int32)`

SetNetTermsInDays sets NetTermsInDays field to given value.

### HasNetTermsInDays

`func (o *EntitlementInfo) HasNetTermsInDays() bool`

HasNetTermsInDays returns a boolean if a field has been set.

### GetPaymentInstallments

`func (o *EntitlementInfo) GetPaymentInstallments() []PaymentInstallment`

GetPaymentInstallments returns the PaymentInstallments field if non-nil, zero value otherwise.

### GetPaymentInstallmentsOk

`func (o *EntitlementInfo) GetPaymentInstallmentsOk() (*[]PaymentInstallment, bool)`

GetPaymentInstallmentsOk returns a tuple with the PaymentInstallments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInstallments

`func (o *EntitlementInfo) SetPaymentInstallments(v []PaymentInstallment)`

SetPaymentInstallments sets PaymentInstallments field to given value.

### HasPaymentInstallments

`func (o *EntitlementInfo) HasPaymentInstallments() bool`

HasPaymentInstallments returns a boolean if a field has been set.

### GetPaymentSchedule

`func (o *EntitlementInfo) GetPaymentSchedule() PaymentScheduleType`

GetPaymentSchedule returns the PaymentSchedule field if non-nil, zero value otherwise.

### GetPaymentScheduleOk

`func (o *EntitlementInfo) GetPaymentScheduleOk() (*PaymentScheduleType, bool)`

GetPaymentScheduleOk returns a tuple with the PaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedule

`func (o *EntitlementInfo) SetPaymentSchedule(v PaymentScheduleType)`

SetPaymentSchedule sets PaymentSchedule field to given value.

### HasPaymentSchedule

`func (o *EntitlementInfo) HasPaymentSchedule() bool`

HasPaymentSchedule returns a boolean if a field has been set.

### GetRefundCancellationPolicy

`func (o *EntitlementInfo) GetRefundCancellationPolicy() string`

GetRefundCancellationPolicy returns the RefundCancellationPolicy field if non-nil, zero value otherwise.

### GetRefundCancellationPolicyOk

`func (o *EntitlementInfo) GetRefundCancellationPolicyOk() (*string, bool)`

GetRefundCancellationPolicyOk returns a tuple with the RefundCancellationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCancellationPolicy

`func (o *EntitlementInfo) SetRefundCancellationPolicy(v string)`

SetRefundCancellationPolicy sets RefundCancellationPolicy field to given value.

### HasRefundCancellationPolicy

`func (o *EntitlementInfo) HasRefundCancellationPolicy() bool`

HasRefundCancellationPolicy returns a boolean if a field has been set.

### GetSellerNotes

`func (o *EntitlementInfo) GetSellerNotes() string`

GetSellerNotes returns the SellerNotes field if non-nil, zero value otherwise.

### GetSellerNotesOk

`func (o *EntitlementInfo) GetSellerNotesOk() (*string, bool)`

GetSellerNotesOk returns a tuple with the SellerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerNotes

`func (o *EntitlementInfo) SetSellerNotes(v string)`

SetSellerNotes sets SellerNotes field to given value.

### HasSellerNotes

`func (o *EntitlementInfo) HasSellerNotes() bool`

HasSellerNotes returns a boolean if a field has been set.

### GetSpaUrl

`func (o *EntitlementInfo) GetSpaUrl() string`

GetSpaUrl returns the SpaUrl field if non-nil, zero value otherwise.

### GetSpaUrlOk

`func (o *EntitlementInfo) GetSpaUrlOk() (*string, bool)`

GetSpaUrlOk returns a tuple with the SpaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaUrl

`func (o *EntitlementInfo) SetSpaUrl(v string)`

SetSpaUrl sets SpaUrl field to given value.

### HasSpaUrl

`func (o *EntitlementInfo) HasSpaUrl() bool`

HasSpaUrl returns a boolean if a field has been set.

### GetTrialConfig

`func (o *EntitlementInfo) GetTrialConfig() TrialConfig`

GetTrialConfig returns the TrialConfig field if non-nil, zero value otherwise.

### GetTrialConfigOk

`func (o *EntitlementInfo) GetTrialConfigOk() (*TrialConfig, bool)`

GetTrialConfigOk returns a tuple with the TrialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialConfig

`func (o *EntitlementInfo) SetTrialConfig(v TrialConfig)`

SetTrialConfig sets TrialConfig field to given value.

### HasTrialConfig

`func (o *EntitlementInfo) HasTrialConfig() bool`

HasTrialConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


