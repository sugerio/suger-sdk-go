# AzureMarketplaceSaasTechnicalConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**AzureAdAppId** | Pointer to **string** | Azure AD Application Id | [optional] 
**AzureAdTenantId** | Pointer to **string** | Azure AD Tenant Id | [optional] 
**ConnectionWebhook** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LandingPageUrl** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** | in format of \&quot;product/product-durable-id\&quot; | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplaceSaasTechnicalConfiguration

`func NewAzureMarketplaceSaasTechnicalConfiguration() *AzureMarketplaceSaasTechnicalConfiguration`

NewAzureMarketplaceSaasTechnicalConfiguration instantiates a new AzureMarketplaceSaasTechnicalConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceSaasTechnicalConfigurationWithDefaults

`func NewAzureMarketplaceSaasTechnicalConfigurationWithDefaults() *AzureMarketplaceSaasTechnicalConfiguration`

NewAzureMarketplaceSaasTechnicalConfigurationWithDefaults instantiates a new AzureMarketplaceSaasTechnicalConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAzureAdAppId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetAzureAdAppId() string`

GetAzureAdAppId returns the AzureAdAppId field if non-nil, zero value otherwise.

### GetAzureAdAppIdOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetAzureAdAppIdOk() (*string, bool)`

GetAzureAdAppIdOk returns a tuple with the AzureAdAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdAppId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetAzureAdAppId(v string)`

SetAzureAdAppId sets AzureAdAppId field to given value.

### HasAzureAdAppId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasAzureAdAppId() bool`

HasAzureAdAppId returns a boolean if a field has been set.

### GetAzureAdTenantId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetAzureAdTenantId() string`

GetAzureAdTenantId returns the AzureAdTenantId field if non-nil, zero value otherwise.

### GetAzureAdTenantIdOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetAzureAdTenantIdOk() (*string, bool)`

GetAzureAdTenantIdOk returns a tuple with the AzureAdTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdTenantId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetAzureAdTenantId(v string)`

SetAzureAdTenantId sets AzureAdTenantId field to given value.

### HasAzureAdTenantId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasAzureAdTenantId() bool`

HasAzureAdTenantId returns a boolean if a field has been set.

### GetConnectionWebhook

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetConnectionWebhook() string`

GetConnectionWebhook returns the ConnectionWebhook field if non-nil, zero value otherwise.

### GetConnectionWebhookOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetConnectionWebhookOk() (*string, bool)`

GetConnectionWebhookOk returns a tuple with the ConnectionWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionWebhook

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetConnectionWebhook(v string)`

SetConnectionWebhook sets ConnectionWebhook field to given value.

### HasConnectionWebhook

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasConnectionWebhook() bool`

HasConnectionWebhook returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLandingPageUrl

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetLandingPageUrl() string`

GetLandingPageUrl returns the LandingPageUrl field if non-nil, zero value otherwise.

### GetLandingPageUrlOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetLandingPageUrlOk() (*string, bool)`

GetLandingPageUrlOk returns a tuple with the LandingPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPageUrl

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetLandingPageUrl(v string)`

SetLandingPageUrl sets LandingPageUrl field to given value.

### HasLandingPageUrl

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasLandingPageUrl() bool`

HasLandingPageUrl returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplaceSaasTechnicalConfiguration) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplaceSaasTechnicalConfiguration) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplaceSaasTechnicalConfiguration) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


