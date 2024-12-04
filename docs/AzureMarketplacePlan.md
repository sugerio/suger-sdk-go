# AzureMarketplacePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**AzureGovernmentCertifications** | Pointer to [**[]AzureMarketplaceGovernmentCertification**](AzureMarketplaceGovernmentCertification.md) |  | [optional] 
**AzureRegions** | Pointer to **[]string** | enums:[azureGlobal,azureGovernment,azureGermany,azureChina] | [optional] 
**DeprecationSchedule** | Pointer to [**AzureMarketplaceDeprecationSchedule**](AzureMarketplaceDeprecationSchedule.md) |  | [optional] 
**DisplayRank** | Pointer to **int32** | default 2147483647 | [optional] 
**Id** | Pointer to **string** | in format of \&quot;plan/product-durable-id/plan-durable-id\&quot; | [optional] 
**Identity** | Pointer to [**AzureMarketplaceIdentity**](AzureMarketplaceIdentity.md) |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) |  | [optional] 
**Product** | Pointer to **string** | in format of \&quot;product/product-durable-id\&quot; | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to **string** | Specifies the plan type (AzureApplication-type products only) see: https://go.microsoft.com/fwlink/?linkid&#x3D;2106322 | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplacePlan

`func NewAzureMarketplacePlan() *AzureMarketplacePlan`

NewAzureMarketplacePlan instantiates a new AzureMarketplacePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePlanWithDefaults

`func NewAzureMarketplacePlanWithDefaults() *AzureMarketplacePlan`

NewAzureMarketplacePlanWithDefaults instantiates a new AzureMarketplacePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplacePlan) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplacePlan) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplacePlan) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplacePlan) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAlias

`func (o *AzureMarketplacePlan) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AzureMarketplacePlan) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AzureMarketplacePlan) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AzureMarketplacePlan) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAzureGovernmentCertifications

`func (o *AzureMarketplacePlan) GetAzureGovernmentCertifications() []AzureMarketplaceGovernmentCertification`

GetAzureGovernmentCertifications returns the AzureGovernmentCertifications field if non-nil, zero value otherwise.

### GetAzureGovernmentCertificationsOk

`func (o *AzureMarketplacePlan) GetAzureGovernmentCertificationsOk() (*[]AzureMarketplaceGovernmentCertification, bool)`

GetAzureGovernmentCertificationsOk returns a tuple with the AzureGovernmentCertifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureGovernmentCertifications

`func (o *AzureMarketplacePlan) SetAzureGovernmentCertifications(v []AzureMarketplaceGovernmentCertification)`

SetAzureGovernmentCertifications sets AzureGovernmentCertifications field to given value.

### HasAzureGovernmentCertifications

`func (o *AzureMarketplacePlan) HasAzureGovernmentCertifications() bool`

HasAzureGovernmentCertifications returns a boolean if a field has been set.

### GetAzureRegions

`func (o *AzureMarketplacePlan) GetAzureRegions() []string`

GetAzureRegions returns the AzureRegions field if non-nil, zero value otherwise.

### GetAzureRegionsOk

`func (o *AzureMarketplacePlan) GetAzureRegionsOk() (*[]string, bool)`

GetAzureRegionsOk returns a tuple with the AzureRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRegions

`func (o *AzureMarketplacePlan) SetAzureRegions(v []string)`

SetAzureRegions sets AzureRegions field to given value.

### HasAzureRegions

`func (o *AzureMarketplacePlan) HasAzureRegions() bool`

HasAzureRegions returns a boolean if a field has been set.

### GetDeprecationSchedule

`func (o *AzureMarketplacePlan) GetDeprecationSchedule() AzureMarketplaceDeprecationSchedule`

GetDeprecationSchedule returns the DeprecationSchedule field if non-nil, zero value otherwise.

### GetDeprecationScheduleOk

`func (o *AzureMarketplacePlan) GetDeprecationScheduleOk() (*AzureMarketplaceDeprecationSchedule, bool)`

GetDeprecationScheduleOk returns a tuple with the DeprecationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecationSchedule

`func (o *AzureMarketplacePlan) SetDeprecationSchedule(v AzureMarketplaceDeprecationSchedule)`

SetDeprecationSchedule sets DeprecationSchedule field to given value.

### HasDeprecationSchedule

`func (o *AzureMarketplacePlan) HasDeprecationSchedule() bool`

HasDeprecationSchedule returns a boolean if a field has been set.

### GetDisplayRank

`func (o *AzureMarketplacePlan) GetDisplayRank() int32`

GetDisplayRank returns the DisplayRank field if non-nil, zero value otherwise.

### GetDisplayRankOk

`func (o *AzureMarketplacePlan) GetDisplayRankOk() (*int32, bool)`

GetDisplayRankOk returns a tuple with the DisplayRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRank

`func (o *AzureMarketplacePlan) SetDisplayRank(v int32)`

SetDisplayRank sets DisplayRank field to given value.

### HasDisplayRank

`func (o *AzureMarketplacePlan) HasDisplayRank() bool`

HasDisplayRank returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplacePlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplacePlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplacePlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplacePlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentity

`func (o *AzureMarketplacePlan) GetIdentity() AzureMarketplaceIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AzureMarketplacePlan) GetIdentityOk() (*AzureMarketplaceIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AzureMarketplacePlan) SetIdentity(v AzureMarketplaceIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AzureMarketplacePlan) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplacePlan) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplacePlan) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplacePlan) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplacePlan) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplacePlan) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplacePlan) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplacePlan) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplacePlan) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplacePlan) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplacePlan) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplacePlan) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplacePlan) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSubtype

`func (o *AzureMarketplacePlan) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AzureMarketplacePlan) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AzureMarketplacePlan) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *AzureMarketplacePlan) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplacePlan) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplacePlan) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplacePlan) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplacePlan) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


