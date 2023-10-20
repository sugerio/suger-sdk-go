# AzurePriceAndAvailabilityPrivateOfferPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** | The azure plan friendly name, from the Azure Marketplace. | [optional] 
**Pricing** | Pointer to [**AzurePriceAndAvailabilityPrivateOfferPrice**](AzurePriceAndAvailabilityPrivateOfferPrice.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzurePriceAndAvailabilityPrivateOfferPlan

`func NewAzurePriceAndAvailabilityPrivateOfferPlan() *AzurePriceAndAvailabilityPrivateOfferPlan`

NewAzurePriceAndAvailabilityPrivateOfferPlan instantiates a new AzurePriceAndAvailabilityPrivateOfferPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePriceAndAvailabilityPrivateOfferPlanWithDefaults

`func NewAzurePriceAndAvailabilityPrivateOfferPlanWithDefaults() *AzurePriceAndAvailabilityPrivateOfferPlan`

NewAzurePriceAndAvailabilityPrivateOfferPlanWithDefaults instantiates a new AzurePriceAndAvailabilityPrivateOfferPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetId

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlan

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanName

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPricing

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetPricing() AzurePriceAndAvailabilityPrivateOfferPrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetPricingOk() (*AzurePriceAndAvailabilityPrivateOfferPrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetPricing(v AzurePriceAndAvailabilityPrivateOfferPrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProduct

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetValidations

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzurePriceAndAvailabilityPrivateOfferPlan) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


