# AzureMarketplacePriceAndAvailabilityPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**BillingTag** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Markets** | Pointer to **[]string** |  | [optional] 
**MeterDefine** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Pricing** | Pointer to [**AzureMarketplacePriceAndAvailabilityPrice**](AzureMarketplacePriceAndAvailabilityPrice.md) |  | [optional] 
**PrivateAudiences** | Pointer to [**[]AzureMarketplacePriceAndAvailabilityAudience**](AzureMarketplacePriceAndAvailabilityAudience.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**SoftwareReservation** | Pointer to [**[]AzureMarketplacePriceAndAvailabilitySoftwareReservation**](AzureMarketplacePriceAndAvailabilitySoftwareReservation.md) |  | [optional] 
**Trial** | Pointer to [**AzureMarketplaceTerm**](AzureMarketplaceTerm.md) |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureMarketplacePriceAndAvailabilityPlan

`func NewAzureMarketplacePriceAndAvailabilityPlan() *AzureMarketplacePriceAndAvailabilityPlan`

NewAzureMarketplacePriceAndAvailabilityPlan instantiates a new AzureMarketplacePriceAndAvailabilityPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceAndAvailabilityPlanWithDefaults

`func NewAzureMarketplacePriceAndAvailabilityPlanWithDefaults() *AzureMarketplacePriceAndAvailabilityPlan`

NewAzureMarketplacePriceAndAvailabilityPlanWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAudience

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetBillingTag

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetBillingTag() string`

GetBillingTag returns the BillingTag field if non-nil, zero value otherwise.

### GetBillingTagOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetBillingTagOk() (*string, bool)`

GetBillingTagOk returns a tuple with the BillingTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTag

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetBillingTag(v string)`

SetBillingTag sets BillingTag field to given value.

### HasBillingTag

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasBillingTag() bool`

HasBillingTag returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMarkets

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetMarkets() []string`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetMarketsOk() (*[]string, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetMarkets(v []string)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.

### GetMeterDefine

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetMeterDefine() string`

GetMeterDefine returns the MeterDefine field if non-nil, zero value otherwise.

### GetMeterDefineOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetMeterDefineOk() (*string, bool)`

GetMeterDefineOk returns a tuple with the MeterDefine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDefine

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetMeterDefine(v string)`

SetMeterDefine sets MeterDefine field to given value.

### HasMeterDefine

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasMeterDefine() bool`

HasMeterDefine returns a boolean if a field has been set.

### GetPlan

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPricing

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetPricing() AzureMarketplacePriceAndAvailabilityPrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetPricingOk() (*AzureMarketplacePriceAndAvailabilityPrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetPricing(v AzureMarketplacePriceAndAvailabilityPrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetPrivateAudiences

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetPrivateAudiences() []AzureMarketplacePriceAndAvailabilityAudience`

GetPrivateAudiences returns the PrivateAudiences field if non-nil, zero value otherwise.

### GetPrivateAudiencesOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetPrivateAudiencesOk() (*[]AzureMarketplacePriceAndAvailabilityAudience, bool)`

GetPrivateAudiencesOk returns a tuple with the PrivateAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAudiences

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetPrivateAudiences(v []AzureMarketplacePriceAndAvailabilityAudience)`

SetPrivateAudiences sets PrivateAudiences field to given value.

### HasPrivateAudiences

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasPrivateAudiences() bool`

HasPrivateAudiences returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSoftwareReservation

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetSoftwareReservation() []AzureMarketplacePriceAndAvailabilitySoftwareReservation`

GetSoftwareReservation returns the SoftwareReservation field if non-nil, zero value otherwise.

### GetSoftwareReservationOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetSoftwareReservationOk() (*[]AzureMarketplacePriceAndAvailabilitySoftwareReservation, bool)`

GetSoftwareReservationOk returns a tuple with the SoftwareReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareReservation

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetSoftwareReservation(v []AzureMarketplacePriceAndAvailabilitySoftwareReservation)`

SetSoftwareReservation sets SoftwareReservation field to given value.

### HasSoftwareReservation

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasSoftwareReservation() bool`

HasSoftwareReservation returns a boolean if a field has been set.

### GetTrial

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetTrial() AzureMarketplaceTerm`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetTrialOk() (*AzureMarketplaceTerm, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetTrial(v AzureMarketplaceTerm)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasValidations() bool`

HasValidations returns a boolean if a field has been set.

### GetVisibility

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AzureMarketplacePriceAndAvailabilityPlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AzureMarketplacePriceAndAvailabilityPlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AzureMarketplacePriceAndAvailabilityPlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


