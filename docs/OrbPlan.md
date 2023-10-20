# OrbPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePlanId** | Pointer to **string** | The parent plan id if the given plan was created by overriding one or more of the parent&#39;s prices. | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DefaultInvoiceMemo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**OrbPriceDiscount**](OrbPriceDiscount.md) |  | [optional] 
**ExternalPlanId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoicingCurrency** | Pointer to **string** |  | [optional] 
**Maximum** | Pointer to [**OrbPriceMaximum**](OrbPriceMaximum.md) |  | [optional] 
**MeteringDimensions** | Pointer to [**[]MeteringDimension**](MeteringDimension.md) | The following fields are populated by Suger. The suger metering dimensions that are mapped to the orb billable metrics. | [optional] 
**Minimum** | Pointer to [**OrbPriceMinimum**](OrbPriceMinimum.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NetTerms** | Pointer to **int32** | Determines the difference between the invoice issue date and the due date. A value of \&quot;0\&quot; here signifies that invoices are due on issue, whereas a value of \&quot;30\&quot; means that the customer has a month to pay the invoice before its overdue. Note that individual subscriptions or invoices may set a different net terms configuration. | [optional] 
**PlanPhases** | Pointer to [**[]OrbPlanPhase**](OrbPlanPhase.md) |  | [optional] 
**Prices** | Pointer to [**[]OrbPrice**](OrbPrice.md) |  | [optional] 
**Product** | Pointer to [**OrbProduct**](OrbProduct.md) |  | [optional] 
**Status** | Pointer to [**OrbPlanStatus**](OrbPlanStatus.md) |  | [optional] 
**TrialConfig** | Pointer to [**OrbTrialConfig**](OrbTrialConfig.md) |  | [optional] 

## Methods

### NewOrbPlan

`func NewOrbPlan() *OrbPlan`

NewOrbPlan instantiates a new OrbPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPlanWithDefaults

`func NewOrbPlanWithDefaults() *OrbPlan`

NewOrbPlanWithDefaults instantiates a new OrbPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePlanId

`func (o *OrbPlan) GetBasePlanId() string`

GetBasePlanId returns the BasePlanId field if non-nil, zero value otherwise.

### GetBasePlanIdOk

`func (o *OrbPlan) GetBasePlanIdOk() (*string, bool)`

GetBasePlanIdOk returns a tuple with the BasePlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePlanId

`func (o *OrbPlan) SetBasePlanId(v string)`

SetBasePlanId sets BasePlanId field to given value.

### HasBasePlanId

`func (o *OrbPlan) HasBasePlanId() bool`

HasBasePlanId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrbPlan) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrbPlan) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrbPlan) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrbPlan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrency

`func (o *OrbPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrbPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrbPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrbPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefaultInvoiceMemo

`func (o *OrbPlan) GetDefaultInvoiceMemo() string`

GetDefaultInvoiceMemo returns the DefaultInvoiceMemo field if non-nil, zero value otherwise.

### GetDefaultInvoiceMemoOk

`func (o *OrbPlan) GetDefaultInvoiceMemoOk() (*string, bool)`

GetDefaultInvoiceMemoOk returns a tuple with the DefaultInvoiceMemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInvoiceMemo

`func (o *OrbPlan) SetDefaultInvoiceMemo(v string)`

SetDefaultInvoiceMemo sets DefaultInvoiceMemo field to given value.

### HasDefaultInvoiceMemo

`func (o *OrbPlan) HasDefaultInvoiceMemo() bool`

HasDefaultInvoiceMemo returns a boolean if a field has been set.

### GetDescription

`func (o *OrbPlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrbPlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrbPlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrbPlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiscount

`func (o *OrbPlan) GetDiscount() OrbPriceDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrbPlan) GetDiscountOk() (*OrbPriceDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrbPlan) SetDiscount(v OrbPriceDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrbPlan) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetExternalPlanId

`func (o *OrbPlan) GetExternalPlanId() string`

GetExternalPlanId returns the ExternalPlanId field if non-nil, zero value otherwise.

### GetExternalPlanIdOk

`func (o *OrbPlan) GetExternalPlanIdOk() (*string, bool)`

GetExternalPlanIdOk returns a tuple with the ExternalPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPlanId

`func (o *OrbPlan) SetExternalPlanId(v string)`

SetExternalPlanId sets ExternalPlanId field to given value.

### HasExternalPlanId

`func (o *OrbPlan) HasExternalPlanId() bool`

HasExternalPlanId returns a boolean if a field has been set.

### GetId

`func (o *OrbPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrbPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrbPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrbPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoicingCurrency

`func (o *OrbPlan) GetInvoicingCurrency() string`

GetInvoicingCurrency returns the InvoicingCurrency field if non-nil, zero value otherwise.

### GetInvoicingCurrencyOk

`func (o *OrbPlan) GetInvoicingCurrencyOk() (*string, bool)`

GetInvoicingCurrencyOk returns a tuple with the InvoicingCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicingCurrency

`func (o *OrbPlan) SetInvoicingCurrency(v string)`

SetInvoicingCurrency sets InvoicingCurrency field to given value.

### HasInvoicingCurrency

`func (o *OrbPlan) HasInvoicingCurrency() bool`

HasInvoicingCurrency returns a boolean if a field has been set.

### GetMaximum

`func (o *OrbPlan) GetMaximum() OrbPriceMaximum`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *OrbPlan) GetMaximumOk() (*OrbPriceMaximum, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *OrbPlan) SetMaximum(v OrbPriceMaximum)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *OrbPlan) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMeteringDimensions

`func (o *OrbPlan) GetMeteringDimensions() []MeteringDimension`

GetMeteringDimensions returns the MeteringDimensions field if non-nil, zero value otherwise.

### GetMeteringDimensionsOk

`func (o *OrbPlan) GetMeteringDimensionsOk() (*[]MeteringDimension, bool)`

GetMeteringDimensionsOk returns a tuple with the MeteringDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringDimensions

`func (o *OrbPlan) SetMeteringDimensions(v []MeteringDimension)`

SetMeteringDimensions sets MeteringDimensions field to given value.

### HasMeteringDimensions

`func (o *OrbPlan) HasMeteringDimensions() bool`

HasMeteringDimensions returns a boolean if a field has been set.

### GetMinimum

`func (o *OrbPlan) GetMinimum() OrbPriceMinimum`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *OrbPlan) GetMinimumOk() (*OrbPriceMinimum, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *OrbPlan) SetMinimum(v OrbPriceMinimum)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *OrbPlan) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetName

`func (o *OrbPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrbPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrbPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrbPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetTerms

`func (o *OrbPlan) GetNetTerms() int32`

GetNetTerms returns the NetTerms field if non-nil, zero value otherwise.

### GetNetTermsOk

`func (o *OrbPlan) GetNetTermsOk() (*int32, bool)`

GetNetTermsOk returns a tuple with the NetTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetTerms

`func (o *OrbPlan) SetNetTerms(v int32)`

SetNetTerms sets NetTerms field to given value.

### HasNetTerms

`func (o *OrbPlan) HasNetTerms() bool`

HasNetTerms returns a boolean if a field has been set.

### GetPlanPhases

`func (o *OrbPlan) GetPlanPhases() []OrbPlanPhase`

GetPlanPhases returns the PlanPhases field if non-nil, zero value otherwise.

### GetPlanPhasesOk

`func (o *OrbPlan) GetPlanPhasesOk() (*[]OrbPlanPhase, bool)`

GetPlanPhasesOk returns a tuple with the PlanPhases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPhases

`func (o *OrbPlan) SetPlanPhases(v []OrbPlanPhase)`

SetPlanPhases sets PlanPhases field to given value.

### HasPlanPhases

`func (o *OrbPlan) HasPlanPhases() bool`

HasPlanPhases returns a boolean if a field has been set.

### GetPrices

`func (o *OrbPlan) GetPrices() []OrbPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *OrbPlan) GetPricesOk() (*[]OrbPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *OrbPlan) SetPrices(v []OrbPrice)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *OrbPlan) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetProduct

`func (o *OrbPlan) GetProduct() OrbProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *OrbPlan) GetProductOk() (*OrbProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *OrbPlan) SetProduct(v OrbProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *OrbPlan) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetStatus

`func (o *OrbPlan) GetStatus() OrbPlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrbPlan) GetStatusOk() (*OrbPlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrbPlan) SetStatus(v OrbPlanStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrbPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrialConfig

`func (o *OrbPlan) GetTrialConfig() OrbTrialConfig`

GetTrialConfig returns the TrialConfig field if non-nil, zero value otherwise.

### GetTrialConfigOk

`func (o *OrbPlan) GetTrialConfigOk() (*OrbTrialConfig, bool)`

GetTrialConfigOk returns a tuple with the TrialConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialConfig

`func (o *OrbPlan) SetTrialConfig(v OrbTrialConfig)`

SetTrialConfig sets TrialConfig field to given value.

### HasTrialConfig

`func (o *OrbPlan) HasTrialConfig() bool`

HasTrialConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


