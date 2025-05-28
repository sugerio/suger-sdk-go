# SnowflakeMarketplaceProductPricingPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseFee** | Pointer to **float32** |  | [optional] 
**BillingDurationMonths** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ContractDurationMonths** | Pointer to **string** |  | [optional] 
**ContractType** | Pointer to **string** | PAY_AS_YOU_GO, LIMITED_TIME | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** | For pricing display purpose. eg: {\&quot;price\&quot;:\&quot;10\&quot;,\&quot;button_text\&quot;:\&quot;starter plan\&quot;} | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PricingModel** | Pointer to **string** | pricing model: USAGE_BASED, FLAT_FEE | [optional] 
**SalesMotion** | Pointer to **string** | SELF_SERVE, PROACTIVE | [optional] 
**State** | Pointer to **string** | state of pricing plan | [optional] 
**UpdatedOn** | Pointer to **string** |  | [optional] 
**UsageDetails** | Pointer to [**SnowflakeMarketplacePricingPlanUsageDetails**](SnowflakeMarketplacePricingPlanUsageDetails.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewSnowflakeMarketplaceProductPricingPlan

`func NewSnowflakeMarketplaceProductPricingPlan() *SnowflakeMarketplaceProductPricingPlan`

NewSnowflakeMarketplaceProductPricingPlan instantiates a new SnowflakeMarketplaceProductPricingPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeMarketplaceProductPricingPlanWithDefaults

`func NewSnowflakeMarketplaceProductPricingPlanWithDefaults() *SnowflakeMarketplaceProductPricingPlan`

NewSnowflakeMarketplaceProductPricingPlanWithDefaults instantiates a new SnowflakeMarketplaceProductPricingPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseFee

`func (o *SnowflakeMarketplaceProductPricingPlan) GetBaseFee() float32`

GetBaseFee returns the BaseFee field if non-nil, zero value otherwise.

### GetBaseFeeOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetBaseFeeOk() (*float32, bool)`

GetBaseFeeOk returns a tuple with the BaseFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseFee

`func (o *SnowflakeMarketplaceProductPricingPlan) SetBaseFee(v float32)`

SetBaseFee sets BaseFee field to given value.

### HasBaseFee

`func (o *SnowflakeMarketplaceProductPricingPlan) HasBaseFee() bool`

HasBaseFee returns a boolean if a field has been set.

### GetBillingDurationMonths

`func (o *SnowflakeMarketplaceProductPricingPlan) GetBillingDurationMonths() string`

GetBillingDurationMonths returns the BillingDurationMonths field if non-nil, zero value otherwise.

### GetBillingDurationMonthsOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetBillingDurationMonthsOk() (*string, bool)`

GetBillingDurationMonthsOk returns a tuple with the BillingDurationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingDurationMonths

`func (o *SnowflakeMarketplaceProductPricingPlan) SetBillingDurationMonths(v string)`

SetBillingDurationMonths sets BillingDurationMonths field to given value.

### HasBillingDurationMonths

`func (o *SnowflakeMarketplaceProductPricingPlan) HasBillingDurationMonths() bool`

HasBillingDurationMonths returns a boolean if a field has been set.

### GetComment

`func (o *SnowflakeMarketplaceProductPricingPlan) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SnowflakeMarketplaceProductPricingPlan) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SnowflakeMarketplaceProductPricingPlan) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContractDurationMonths

`func (o *SnowflakeMarketplaceProductPricingPlan) GetContractDurationMonths() string`

GetContractDurationMonths returns the ContractDurationMonths field if non-nil, zero value otherwise.

### GetContractDurationMonthsOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetContractDurationMonthsOk() (*string, bool)`

GetContractDurationMonthsOk returns a tuple with the ContractDurationMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDurationMonths

`func (o *SnowflakeMarketplaceProductPricingPlan) SetContractDurationMonths(v string)`

SetContractDurationMonths sets ContractDurationMonths field to given value.

### HasContractDurationMonths

`func (o *SnowflakeMarketplaceProductPricingPlan) HasContractDurationMonths() bool`

HasContractDurationMonths returns a boolean if a field has been set.

### GetContractType

`func (o *SnowflakeMarketplaceProductPricingPlan) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *SnowflakeMarketplaceProductPricingPlan) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *SnowflakeMarketplaceProductPricingPlan) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### GetCurrency

`func (o *SnowflakeMarketplaceProductPricingPlan) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SnowflakeMarketplaceProductPricingPlan) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SnowflakeMarketplaceProductPricingPlan) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *SnowflakeMarketplaceProductPricingPlan) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SnowflakeMarketplaceProductPricingPlan) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SnowflakeMarketplaceProductPricingPlan) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMetadata

`func (o *SnowflakeMarketplaceProductPricingPlan) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SnowflakeMarketplaceProductPricingPlan) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SnowflakeMarketplaceProductPricingPlan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *SnowflakeMarketplaceProductPricingPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnowflakeMarketplaceProductPricingPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnowflakeMarketplaceProductPricingPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPricingModel

`func (o *SnowflakeMarketplaceProductPricingPlan) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *SnowflakeMarketplaceProductPricingPlan) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *SnowflakeMarketplaceProductPricingPlan) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.

### GetSalesMotion

`func (o *SnowflakeMarketplaceProductPricingPlan) GetSalesMotion() string`

GetSalesMotion returns the SalesMotion field if non-nil, zero value otherwise.

### GetSalesMotionOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetSalesMotionOk() (*string, bool)`

GetSalesMotionOk returns a tuple with the SalesMotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesMotion

`func (o *SnowflakeMarketplaceProductPricingPlan) SetSalesMotion(v string)`

SetSalesMotion sets SalesMotion field to given value.

### HasSalesMotion

`func (o *SnowflakeMarketplaceProductPricingPlan) HasSalesMotion() bool`

HasSalesMotion returns a boolean if a field has been set.

### GetState

`func (o *SnowflakeMarketplaceProductPricingPlan) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SnowflakeMarketplaceProductPricingPlan) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *SnowflakeMarketplaceProductPricingPlan) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedOn

`func (o *SnowflakeMarketplaceProductPricingPlan) GetUpdatedOn() string`

GetUpdatedOn returns the UpdatedOn field if non-nil, zero value otherwise.

### GetUpdatedOnOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetUpdatedOnOk() (*string, bool)`

GetUpdatedOnOk returns a tuple with the UpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedOn

`func (o *SnowflakeMarketplaceProductPricingPlan) SetUpdatedOn(v string)`

SetUpdatedOn sets UpdatedOn field to given value.

### HasUpdatedOn

`func (o *SnowflakeMarketplaceProductPricingPlan) HasUpdatedOn() bool`

HasUpdatedOn returns a boolean if a field has been set.

### GetUsageDetails

`func (o *SnowflakeMarketplaceProductPricingPlan) GetUsageDetails() SnowflakeMarketplacePricingPlanUsageDetails`

GetUsageDetails returns the UsageDetails field if non-nil, zero value otherwise.

### GetUsageDetailsOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetUsageDetailsOk() (*SnowflakeMarketplacePricingPlanUsageDetails, bool)`

GetUsageDetailsOk returns a tuple with the UsageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDetails

`func (o *SnowflakeMarketplaceProductPricingPlan) SetUsageDetails(v SnowflakeMarketplacePricingPlanUsageDetails)`

SetUsageDetails sets UsageDetails field to given value.

### HasUsageDetails

`func (o *SnowflakeMarketplaceProductPricingPlan) HasUsageDetails() bool`

HasUsageDetails returns a boolean if a field has been set.

### GetVisibility

`func (o *SnowflakeMarketplaceProductPricingPlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SnowflakeMarketplaceProductPricingPlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SnowflakeMarketplaceProductPricingPlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *SnowflakeMarketplaceProductPricingPlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


