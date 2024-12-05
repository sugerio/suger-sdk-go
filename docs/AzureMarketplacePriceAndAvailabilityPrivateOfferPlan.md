# AzureMarketplacePriceAndAvailabilityPrivateOfferPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) |  | [optional] 
**OfferPricingType** | Pointer to [**AzureMarketplaceOfferPricingType**](AzureMarketplaceOfferPricingType.md) | default \&quot;editExistingOfferPricingOnly\&quot; | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**PlanName** | Pointer to **string** | The azure plan friendly name, from the Azure Marketplace. | [optional] 
**Pricing** | Pointer to [**AzureMarketplacePriceAndAvailabilityPrivateOfferPrice**](AzureMarketplacePriceAndAvailabilityPrivateOfferPrice.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**SoftwareReservation** | Pointer to [**AzureMarketplacePriceAndAvailabilityPrivateOfferPlanSoftwareReservation**](AzureMarketplacePriceAndAvailabilityPrivateOfferPlanSoftwareReservation.md) |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 
**Visibility** | Pointer to **string** | default \&quot;visible\&quot; | [optional] 

## Methods

### NewAzureMarketplacePriceAndAvailabilityPrivateOfferPlan

`func NewAzureMarketplacePriceAndAvailabilityPrivateOfferPlan() *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan`

NewAzureMarketplacePriceAndAvailabilityPrivateOfferPlan instantiates a new AzureMarketplacePriceAndAvailabilityPrivateOfferPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceAndAvailabilityPrivateOfferPlanWithDefaults

`func NewAzureMarketplacePriceAndAvailabilityPrivateOfferPlanWithDefaults() *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan`

NewAzureMarketplacePriceAndAvailabilityPrivateOfferPlanWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityPrivateOfferPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetOfferPricingType

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetOfferPricingType() AzureMarketplaceOfferPricingType`

GetOfferPricingType returns the OfferPricingType field if non-nil, zero value otherwise.

### GetOfferPricingTypeOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetOfferPricingTypeOk() (*AzureMarketplaceOfferPricingType, bool)`

GetOfferPricingTypeOk returns a tuple with the OfferPricingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPricingType

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetOfferPricingType(v AzureMarketplaceOfferPricingType)`

SetOfferPricingType sets OfferPricingType field to given value.

### HasOfferPricingType

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasOfferPricingType() bool`

HasOfferPricingType returns a boolean if a field has been set.

### GetPlan

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlanName

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetPricing

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetPricing() AzureMarketplacePriceAndAvailabilityPrivateOfferPrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetPricingOk() (*AzureMarketplacePriceAndAvailabilityPrivateOfferPrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetPricing(v AzureMarketplacePriceAndAvailabilityPrivateOfferPrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSoftwareReservation

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetSoftwareReservation() AzureMarketplacePriceAndAvailabilityPrivateOfferPlanSoftwareReservation`

GetSoftwareReservation returns the SoftwareReservation field if non-nil, zero value otherwise.

### GetSoftwareReservationOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetSoftwareReservationOk() (*AzureMarketplacePriceAndAvailabilityPrivateOfferPlanSoftwareReservation, bool)`

GetSoftwareReservationOk returns a tuple with the SoftwareReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareReservation

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetSoftwareReservation(v AzureMarketplacePriceAndAvailabilityPrivateOfferPlanSoftwareReservation)`

SetSoftwareReservation sets SoftwareReservation field to given value.

### HasSoftwareReservation

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasSoftwareReservation() bool`

HasSoftwareReservation returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasValidations() bool`

HasValidations returns a boolean if a field has been set.

### GetVisibility

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AzureMarketplacePriceAndAvailabilityPrivateOfferPlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


