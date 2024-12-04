# AzureMarketplacePlanListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LanguageId** | Pointer to **string** |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplacePlanListing

`func NewAzureMarketplacePlanListing() *AzureMarketplacePlanListing`

NewAzureMarketplacePlanListing instantiates a new AzureMarketplacePlanListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePlanListingWithDefaults

`func NewAzureMarketplacePlanListingWithDefaults() *AzureMarketplacePlanListing`

NewAzureMarketplacePlanListingWithDefaults instantiates a new AzureMarketplacePlanListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplacePlanListing) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplacePlanListing) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplacePlanListing) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplacePlanListing) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetDescription

`func (o *AzureMarketplacePlanListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureMarketplacePlanListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureMarketplacePlanListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureMarketplacePlanListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplacePlanListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplacePlanListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplacePlanListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplacePlanListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AzureMarketplacePlanListing) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AzureMarketplacePlanListing) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AzureMarketplacePlanListing) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AzureMarketplacePlanListing) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguageId

`func (o *AzureMarketplacePlanListing) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *AzureMarketplacePlanListing) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *AzureMarketplacePlanListing) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *AzureMarketplacePlanListing) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplacePlanListing) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplacePlanListing) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplacePlanListing) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplacePlanListing) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetName

`func (o *AzureMarketplacePlanListing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureMarketplacePlanListing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureMarketplacePlanListing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureMarketplacePlanListing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlan

`func (o *AzureMarketplacePlanListing) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AzureMarketplacePlanListing) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AzureMarketplacePlanListing) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AzureMarketplacePlanListing) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplacePlanListing) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplacePlanListing) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplacePlanListing) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplacePlanListing) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplacePlanListing) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplacePlanListing) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplacePlanListing) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplacePlanListing) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSummary

`func (o *AzureMarketplacePlanListing) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AzureMarketplacePlanListing) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AzureMarketplacePlanListing) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AzureMarketplacePlanListing) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplacePlanListing) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplacePlanListing) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplacePlanListing) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplacePlanListing) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


