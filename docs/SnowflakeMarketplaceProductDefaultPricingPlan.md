# SnowflakeMarketplaceProductDefaultPricingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowEarlyAccess** | Pointer to **bool** | If true, consumers can access the listing before making a payment. You&#39;ll need to remind them to pay. | [optional] 
**BaseFee** | Pointer to **float32** |  | [optional] 
**BillingDuration** | Pointer to **int32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FreeUnitKind** | Pointer to **string** |  | [optional] 
**FreeUnits** | Pointer to **float32** |  | [optional] 
**InstallmentSchedule** | Pointer to [**SnowflakeMarketplacePlanInstallmentSchedule**](SnowflakeMarketplacePlanInstallmentSchedule.md) |  | [optional] 
**IsAutoRenewable** | Pointer to **bool** | IsAutoRenewable is true if the product is auto renewable. | [optional] 
**MaxFee** | Pointer to **float32** |  | [optional] 
**PaymentType** | Pointer to **string** |  | [optional] 
**TrialUsageLimit** | Pointer to **float32** |  | [optional] 
**TrialUsageUnit** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsageUnitKind** | Pointer to **string** |  | [optional] 
**UsageUnitPrice** | Pointer to **float32** |  | [optional] 

## Methods

### NewSnowflakeMarketplaceProductDefaultPricingPlan

`func NewSnowflakeMarketplaceProductDefaultPricingPlan() *SnowflakeMarketplaceProductDefaultPricingPlan`

NewSnowflakeMarketplaceProductDefaultPricingPlan instantiates a new SnowflakeMarketplaceProductDefaultPricingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeMarketplaceProductDefaultPricingPlanWithDefaults

`func NewSnowflakeMarketplaceProductDefaultPricingPlanWithDefaults() *SnowflakeMarketplaceProductDefaultPricingPlan`

NewSnowflakeMarketplaceProductDefaultPricingPlanWithDefaults instantiates a new SnowflakeMarketplaceProductDefaultPricingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowEarlyAccess

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetAllowEarlyAccess() bool`

GetAllowEarlyAccess returns the AllowEarlyAccess field if non-nil, zero value otherwise.

### GetAllowEarlyAccessOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetAllowEarlyAccessOk() (*bool, bool)`

GetAllowEarlyAccessOk returns a tuple with the AllowEarlyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEarlyAccess

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetAllowEarlyAccess(v bool)`

SetAllowEarlyAccess sets AllowEarlyAccess field to given value.

### HasAllowEarlyAccess

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasAllowEarlyAccess() bool`

HasAllowEarlyAccess returns a boolean if a field has been set.

### GetBaseFee

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetBaseFee() float32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetBaseFeeOk() (*float32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetBaseFee(v float32)`

SetBaseFee sets BaseFee field to given value.

### HasBaseFee

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasBaseFee() bool`

HasBaseFee returns a boolean if a field has been set.

### GetBillingDuration

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetBillingDuration() int32`

GetBillingDuration returns the BillingDuration field if non-nil, zero value otherwise.

### GetBillingDurationOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetBillingDurationOk() (*int32, bool)`

GetBillingDurationOk returns a tuple with the BillingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDuration

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetBillingDuration(v int32)`

SetBillingDuration sets BillingDuration field to given value.

### HasBillingDuration

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasBillingDuration() bool`

HasBillingDuration returns a boolean if a field has been set.

### GetCurrency

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFreeUnitKind

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetFreeUnitKind() string`

GetFreeUnitKind returns the FreeUnitKind field if non-nil, zero value otherwise.

### GetFreeUnitKindOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetFreeUnitKindOk() (*string, bool)`

GetFreeUnitKindOk returns a tuple with the FreeUnitKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnitKind

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetFreeUnitKind(v string)`

SetFreeUnitKind sets FreeUnitKind field to given value.

### HasFreeUnitKind

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasFreeUnitKind() bool`

HasFreeUnitKind returns a boolean if a field has been set.

### GetFreeUnits

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetFreeUnits() float32`

GetFreeUnits returns the FreeUnits field if non-nil, zero value otherwise.

### GetFreeUnitsOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetFreeUnitsOk() (*float32, bool)`

GetFreeUnitsOk returns a tuple with the FreeUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeUnits

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetFreeUnits(v float32)`

SetFreeUnits sets FreeUnits field to given value.

### HasFreeUnits

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasFreeUnits() bool`

HasFreeUnits returns a boolean if a field has been set.

### GetInstallmentSchedule

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetInstallmentSchedule() SnowflakeMarketplacePlanInstallmentSchedule`

GetInstallmentSchedule returns the InstallmentSchedule field if non-nil, zero value otherwise.

### GetInstallmentScheduleOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetInstallmentScheduleOk() (*SnowflakeMarketplacePlanInstallmentSchedule, bool)`

GetInstallmentScheduleOk returns a tuple with the InstallmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentSchedule

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetInstallmentSchedule(v SnowflakeMarketplacePlanInstallmentSchedule)`

SetInstallmentSchedule sets InstallmentSchedule field to given value.

### HasInstallmentSchedule

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasInstallmentSchedule() bool`

HasInstallmentSchedule returns a boolean if a field has been set.

### GetIsAutoRenewable

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetIsAutoRenewable() bool`

GetIsAutoRenewable returns the IsAutoRenewable field if non-nil, zero value otherwise.

### GetIsAutoRenewableOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetIsAutoRenewableOk() (*bool, bool)`

GetIsAutoRenewableOk returns a tuple with the IsAutoRenewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRenewable

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetIsAutoRenewable(v bool)`

SetIsAutoRenewable sets IsAutoRenewable field to given value.

### HasIsAutoRenewable

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasIsAutoRenewable() bool`

HasIsAutoRenewable returns a boolean if a field has been set.

### GetMaxFee

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetMaxFee() float32`

GetMaxFee returns the MaxFee field if non-nil, zero value otherwise.

### GetMaxFeeOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetMaxFeeOk() (*float32, bool)`

GetMaxFeeOk returns a tuple with the MaxFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFee

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetMaxFee(v float32)`

SetMaxFee sets MaxFee field to given value.

### HasMaxFee

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasMaxFee() bool`

HasMaxFee returns a boolean if a field has been set.

### GetPaymentType

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetPaymentType() string`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetPaymentTypeOk() (*string, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetPaymentType(v string)`

SetPaymentType sets PaymentType field to given value.

### HasPaymentType

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasPaymentType() bool`

HasPaymentType returns a boolean if a field has been set.

### GetTrialUsageLimit

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetTrialUsageLimit() float32`

GetTrialUsageLimit returns the TrialUsageLimit field if non-nil, zero value otherwise.

### GetTrialUsageLimitOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetTrialUsageLimitOk() (*float32, bool)`

GetTrialUsageLimitOk returns a tuple with the TrialUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialUsageLimit

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetTrialUsageLimit(v float32)`

SetTrialUsageLimit sets TrialUsageLimit field to given value.

### HasTrialUsageLimit

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasTrialUsageLimit() bool`

HasTrialUsageLimit returns a boolean if a field has been set.

### GetTrialUsageUnit

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetTrialUsageUnit() string`

GetTrialUsageUnit returns the TrialUsageUnit field if non-nil, zero value otherwise.

### GetTrialUsageUnitOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetTrialUsageUnitOk() (*string, bool)`

GetTrialUsageUnitOk returns a tuple with the TrialUsageUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialUsageUnit

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetTrialUsageUnit(v string)`

SetTrialUsageUnit sets TrialUsageUnit field to given value.

### HasTrialUsageUnit

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasTrialUsageUnit() bool`

HasTrialUsageUnit returns a boolean if a field has been set.

### GetType

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsageUnitKind

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetUsageUnitKind() string`

GetUsageUnitKind returns the UsageUnitKind field if non-nil, zero value otherwise.

### GetUsageUnitKindOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetUsageUnitKindOk() (*string, bool)`

GetUsageUnitKindOk returns a tuple with the UsageUnitKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUnitKind

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetUsageUnitKind(v string)`

SetUsageUnitKind sets UsageUnitKind field to given value.

### HasUsageUnitKind

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasUsageUnitKind() bool`

HasUsageUnitKind returns a boolean if a field has been set.

### GetUsageUnitPrice

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetUsageUnitPrice() float32`

GetUsageUnitPrice returns the UsageUnitPrice field if non-nil, zero value otherwise.

### GetUsageUnitPriceOk

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) GetUsageUnitPriceOk() (*float32, bool)`

GetUsageUnitPriceOk returns a tuple with the UsageUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUnitPrice

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) SetUsageUnitPrice(v float32)`

SetUsageUnitPrice sets UsageUnitPrice field to given value.

### HasUsageUnitPrice

`func (o *SnowflakeMarketplaceProductDefaultPricingPlan) HasUsageUnitPrice() bool`

HasUsageUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


