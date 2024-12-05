# AzureCommercialMarketplaceSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**AccessUrl** | Pointer to **string** | in patern of \&quot;^(http|https)://\&quot; | [optional] 
**CallToAction** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | In format of \&quot;commercial-marketplace-setup/setup-durable-id\&quot; | [optional] 
**Product** | Pointer to **string** | Product resource name, in format of \&quot;product/product-durable-id\&quot; | [optional] 
**RequireLicenseForInstall** | Pointer to **bool** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**SellThroughMicrosoft** | Pointer to **bool** |  | [optional] 
**UseMicrosoftLicenseManagementService** | Pointer to **bool** | If true, only per_user pricing model is allowed. | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureCommercialMarketplaceSetup

`func NewAzureCommercialMarketplaceSetup() *AzureCommercialMarketplaceSetup`

NewAzureCommercialMarketplaceSetup instantiates a new AzureCommercialMarketplaceSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCommercialMarketplaceSetupWithDefaults

`func NewAzureCommercialMarketplaceSetupWithDefaults() *AzureCommercialMarketplaceSetup`

NewAzureCommercialMarketplaceSetupWithDefaults instantiates a new AzureCommercialMarketplaceSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureCommercialMarketplaceSetup) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureCommercialMarketplaceSetup) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureCommercialMarketplaceSetup) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureCommercialMarketplaceSetup) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAccessUrl

`func (o *AzureCommercialMarketplaceSetup) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *AzureCommercialMarketplaceSetup) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *AzureCommercialMarketplaceSetup) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.

### HasAccessUrl

`func (o *AzureCommercialMarketplaceSetup) HasAccessUrl() bool`

HasAccessUrl returns a boolean if a field has been set.

### GetCallToAction

`func (o *AzureCommercialMarketplaceSetup) GetCallToAction() string`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *AzureCommercialMarketplaceSetup) GetCallToActionOk() (*string, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *AzureCommercialMarketplaceSetup) SetCallToAction(v string)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *AzureCommercialMarketplaceSetup) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetId

`func (o *AzureCommercialMarketplaceSetup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCommercialMarketplaceSetup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCommercialMarketplaceSetup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureCommercialMarketplaceSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProduct

`func (o *AzureCommercialMarketplaceSetup) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureCommercialMarketplaceSetup) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureCommercialMarketplaceSetup) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureCommercialMarketplaceSetup) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRequireLicenseForInstall

`func (o *AzureCommercialMarketplaceSetup) GetRequireLicenseForInstall() bool`

GetRequireLicenseForInstall returns the RequireLicenseForInstall field if non-nil, zero value otherwise.

### GetRequireLicenseForInstallOk

`func (o *AzureCommercialMarketplaceSetup) GetRequireLicenseForInstallOk() (*bool, bool)`

GetRequireLicenseForInstallOk returns a tuple with the RequireLicenseForInstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireLicenseForInstall

`func (o *AzureCommercialMarketplaceSetup) SetRequireLicenseForInstall(v bool)`

SetRequireLicenseForInstall sets RequireLicenseForInstall field to given value.

### HasRequireLicenseForInstall

`func (o *AzureCommercialMarketplaceSetup) HasRequireLicenseForInstall() bool`

HasRequireLicenseForInstall returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureCommercialMarketplaceSetup) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureCommercialMarketplaceSetup) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureCommercialMarketplaceSetup) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureCommercialMarketplaceSetup) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSellThroughMicrosoft

`func (o *AzureCommercialMarketplaceSetup) GetSellThroughMicrosoft() bool`

GetSellThroughMicrosoft returns the SellThroughMicrosoft field if non-nil, zero value otherwise.

### GetSellThroughMicrosoftOk

`func (o *AzureCommercialMarketplaceSetup) GetSellThroughMicrosoftOk() (*bool, bool)`

GetSellThroughMicrosoftOk returns a tuple with the SellThroughMicrosoft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellThroughMicrosoft

`func (o *AzureCommercialMarketplaceSetup) SetSellThroughMicrosoft(v bool)`

SetSellThroughMicrosoft sets SellThroughMicrosoft field to given value.

### HasSellThroughMicrosoft

`func (o *AzureCommercialMarketplaceSetup) HasSellThroughMicrosoft() bool`

HasSellThroughMicrosoft returns a boolean if a field has been set.

### GetUseMicrosoftLicenseManagementService

`func (o *AzureCommercialMarketplaceSetup) GetUseMicrosoftLicenseManagementService() bool`

GetUseMicrosoftLicenseManagementService returns the UseMicrosoftLicenseManagementService field if non-nil, zero value otherwise.

### GetUseMicrosoftLicenseManagementServiceOk

`func (o *AzureCommercialMarketplaceSetup) GetUseMicrosoftLicenseManagementServiceOk() (*bool, bool)`

GetUseMicrosoftLicenseManagementServiceOk returns a tuple with the UseMicrosoftLicenseManagementService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMicrosoftLicenseManagementService

`func (o *AzureCommercialMarketplaceSetup) SetUseMicrosoftLicenseManagementService(v bool)`

SetUseMicrosoftLicenseManagementService sets UseMicrosoftLicenseManagementService field to given value.

### HasUseMicrosoftLicenseManagementService

`func (o *AzureCommercialMarketplaceSetup) HasUseMicrosoftLicenseManagementService() bool`

HasUseMicrosoftLicenseManagementService returns a boolean if a field has been set.

### GetValidations

`func (o *AzureCommercialMarketplaceSetup) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureCommercialMarketplaceSetup) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureCommercialMarketplaceSetup) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureCommercialMarketplaceSetup) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


