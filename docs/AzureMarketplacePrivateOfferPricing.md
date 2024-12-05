# AzureMarketplacePrivateOfferPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasePlan** | Pointer to **string** | required for SaasNewCustomizedPlans | [optional] 
**DiscountPercentage** | Pointer to **float32** | between 0.01 to 100 | [optional] 
**DiscountType** | Pointer to [**PrivateOfferDiscountType**](PrivateOfferDiscountType.md) |  | [optional] 
**MarkupPercentage** | Pointer to **float32** | between 0.00000001 to 100 | [optional] 
**NewPlanDetails** | Pointer to [**AzureMarketplacePrivateOfferPricingNewPlanDetails**](AzureMarketplacePrivateOfferPricingNewPlanDetails.md) | required for SaasNewCustomizedPlans | [optional] 
**OriginalPlan** | Pointer to [**AzureMarketplacePriceAndAvailabilityPrivateOfferPlan**](AzureMarketplacePriceAndAvailabilityPrivateOfferPlan.md) | the pricing plan of the original plan. | [optional] 
**Plan** | Pointer to **string** | The base/original/default plan of the private offer, in format of \&quot;plan/product-durable-id/plan-durable-id\&quot; | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**PlanType** | Pointer to **string** | The type of the plan, FLAT_RATE or PER_USER. | [optional] 
**PriceDetails** | Pointer to **map[string]interface{}** |  | [optional] 
**PrivateOfferPlan** | Pointer to [**AzureMarketplacePriceAndAvailabilityPrivateOfferPlan**](AzureMarketplacePriceAndAvailabilityPrivateOfferPlan.md) | the pricing plan of the private offer | [optional] 
**Product** | Pointer to **string** | in format of \&quot;product/product-durable-id\&quot; | [optional] 
**ProductName** | Pointer to **string** |  | [optional] 
**SugerOfferId** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureMarketplacePrivateOfferPricing

`func NewAzureMarketplacePrivateOfferPricing() *AzureMarketplacePrivateOfferPricing`

NewAzureMarketplacePrivateOfferPricing instantiates a new AzureMarketplacePrivateOfferPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePrivateOfferPricingWithDefaults

`func NewAzureMarketplacePrivateOfferPricingWithDefaults() *AzureMarketplacePrivateOfferPricing`

NewAzureMarketplacePrivateOfferPricingWithDefaults instantiates a new AzureMarketplacePrivateOfferPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasePlan

`func (o *AzureMarketplacePrivateOfferPricing) GetBasePlan() string`

GetBasePlan returns the BasePlan field if non-nil, zero value otherwise.

### GetBasePlanOk

`func (o *AzureMarketplacePrivateOfferPricing) GetBasePlanOk() (*string, bool)`

GetBasePlanOk returns a tuple with the BasePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePlan

`func (o *AzureMarketplacePrivateOfferPricing) SetBasePlan(v string)`

SetBasePlan sets BasePlan field to given value.

### HasBasePlan

`func (o *AzureMarketplacePrivateOfferPricing) HasBasePlan() bool`

HasBasePlan returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *AzureMarketplacePrivateOfferPricing) GetDiscountPercentage() float32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *AzureMarketplacePrivateOfferPricing) GetDiscountPercentageOk() (*float32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *AzureMarketplacePrivateOfferPricing) SetDiscountPercentage(v float32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *AzureMarketplacePrivateOfferPricing) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountType

`func (o *AzureMarketplacePrivateOfferPricing) GetDiscountType() PrivateOfferDiscountType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *AzureMarketplacePrivateOfferPricing) GetDiscountTypeOk() (*PrivateOfferDiscountType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *AzureMarketplacePrivateOfferPricing) SetDiscountType(v PrivateOfferDiscountType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *AzureMarketplacePrivateOfferPricing) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetMarkupPercentage

`func (o *AzureMarketplacePrivateOfferPricing) GetMarkupPercentage() float32`

GetMarkupPercentage returns the MarkupPercentage field if non-nil, zero value otherwise.

### GetMarkupPercentageOk

`func (o *AzureMarketplacePrivateOfferPricing) GetMarkupPercentageOk() (*float32, bool)`

GetMarkupPercentageOk returns a tuple with the MarkupPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercentage

`func (o *AzureMarketplacePrivateOfferPricing) SetMarkupPercentage(v float32)`

SetMarkupPercentage sets MarkupPercentage field to given value.

### HasMarkupPercentage

`func (o *AzureMarketplacePrivateOfferPricing) HasMarkupPercentage() bool`

HasMarkupPercentage returns a boolean if a field has been set.

### GetNewPlanDetails

`func (o *AzureMarketplacePrivateOfferPricing) GetNewPlanDetails() AzureMarketplacePrivateOfferPricingNewPlanDetails`

GetNewPlanDetails returns the NewPlanDetails field if non-nil, zero value otherwise.

### GetNewPlanDetailsOk

`func (o *AzureMarketplacePrivateOfferPricing) GetNewPlanDetailsOk() (*AzureMarketplacePrivateOfferPricingNewPlanDetails, bool)`

GetNewPlanDetailsOk returns a tuple with the NewPlanDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlanDetails

`func (o *AzureMarketplacePrivateOfferPricing) SetNewPlanDetails(v AzureMarketplacePrivateOfferPricingNewPlanDetails)`

SetNewPlanDetails sets NewPlanDetails field to given value.

### HasNewPlanDetails

`func (o *AzureMarketplacePrivateOfferPricing) HasNewPlanDetails() bool`

HasNewPlanDetails returns a boolean if a field has been set.

### GetOriginalPlan

`func (o *AzureMarketplacePrivateOfferPricing) GetOriginalPlan() AzureMarketplacePriceAndAvailabilityPrivateOfferPlan`

GetOriginalPlan returns the OriginalPlan field if non-nil, zero value otherwise.

### GetOriginalPlanOk

`func (o *AzureMarketplacePrivateOfferPricing) GetOriginalPlanOk() (*AzureMarketplacePriceAndAvailabilityPrivateOfferPlan, bool)`

GetOriginalPlanOk returns a tuple with the OriginalPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPlan

`func (o *AzureMarketplacePrivateOfferPricing) SetOriginalPlan(v AzureMarketplacePriceAndAvailabilityPrivateOfferPlan)`

SetOriginalPlan sets OriginalPlan field to given value.

### HasOriginalPlan

`func (o *AzureMarketplacePrivateOfferPricing) HasOriginalPlan() bool`

HasOriginalPlan returns a boolean if a field has been set.

### GetPlan

`func (o *AzureMarketplacePrivateOfferPricing) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AzureMarketplacePrivateOfferPricing) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AzureMarketplacePrivateOfferPricing) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanId

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AzureMarketplacePrivateOfferPricing) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *AzureMarketplacePrivateOfferPricing) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetPlanName

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *AzureMarketplacePrivateOfferPricing) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *AzureMarketplacePrivateOfferPricing) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPlanType

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *AzureMarketplacePrivateOfferPricing) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *AzureMarketplacePrivateOfferPricing) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *AzureMarketplacePrivateOfferPricing) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPriceDetails

`func (o *AzureMarketplacePrivateOfferPricing) GetPriceDetails() map[string]interface{}`

GetPriceDetails returns the PriceDetails field if non-nil, zero value otherwise.

### GetPriceDetailsOk

`func (o *AzureMarketplacePrivateOfferPricing) GetPriceDetailsOk() (*map[string]interface{}, bool)`

GetPriceDetailsOk returns a tuple with the PriceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDetails

`func (o *AzureMarketplacePrivateOfferPricing) SetPriceDetails(v map[string]interface{})`

SetPriceDetails sets PriceDetails field to given value.

### HasPriceDetails

`func (o *AzureMarketplacePrivateOfferPricing) HasPriceDetails() bool`

HasPriceDetails returns a boolean if a field has been set.

### GetPrivateOfferPlan

`func (o *AzureMarketplacePrivateOfferPricing) GetPrivateOfferPlan() AzureMarketplacePriceAndAvailabilityPrivateOfferPlan`

GetPrivateOfferPlan returns the PrivateOfferPlan field if non-nil, zero value otherwise.

### GetPrivateOfferPlanOk

`func (o *AzureMarketplacePrivateOfferPricing) GetPrivateOfferPlanOk() (*AzureMarketplacePriceAndAvailabilityPrivateOfferPlan, bool)`

GetPrivateOfferPlanOk returns a tuple with the PrivateOfferPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOfferPlan

`func (o *AzureMarketplacePrivateOfferPricing) SetPrivateOfferPlan(v AzureMarketplacePriceAndAvailabilityPrivateOfferPlan)`

SetPrivateOfferPlan sets PrivateOfferPlan field to given value.

### HasPrivateOfferPlan

`func (o *AzureMarketplacePrivateOfferPricing) HasPrivateOfferPlan() bool`

HasPrivateOfferPlan returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplacePrivateOfferPricing) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplacePrivateOfferPricing) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplacePrivateOfferPricing) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplacePrivateOfferPricing) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProductName

`func (o *AzureMarketplacePrivateOfferPricing) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *AzureMarketplacePrivateOfferPricing) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *AzureMarketplacePrivateOfferPricing) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *AzureMarketplacePrivateOfferPricing) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetSugerOfferId

`func (o *AzureMarketplacePrivateOfferPricing) GetSugerOfferId() string`

GetSugerOfferId returns the SugerOfferId field if non-nil, zero value otherwise.

### GetSugerOfferIdOk

`func (o *AzureMarketplacePrivateOfferPricing) GetSugerOfferIdOk() (*string, bool)`

GetSugerOfferIdOk returns a tuple with the SugerOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSugerOfferId

`func (o *AzureMarketplacePrivateOfferPricing) SetSugerOfferId(v string)`

SetSugerOfferId sets SugerOfferId field to given value.

### HasSugerOfferId

`func (o *AzureMarketplacePrivateOfferPricing) HasSugerOfferId() bool`

HasSugerOfferId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


