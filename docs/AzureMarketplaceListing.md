# AzureMarketplaceListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**CloudSolutionProviderContact** | Pointer to [**AzureMarketplaceContact**](AzureMarketplaceContact.md) |  | [optional] 
**CloudSolutionProviderMarketingMaterials** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EngineeringContact** | Pointer to [**AzureMarketplaceContact**](AzureMarketplaceContact.md) |  | [optional] 
**GeneralLinks** | Pointer to [**[]AzureMarketplaceGeneralLink**](AzureMarketplaceGeneralLink.md) |  | [optional] 
**GettingStartedInstructions** | Pointer to **string** |  | [optional] 
**GloabalSupportWebsite** | Pointer to **string** |  | [optional] 
**GovernmentSupportWebsite** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LanguageId** | Pointer to **string** |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) | Default value is \&quot;generallyAvailable\&quot;. | [optional] 
**PrivacyPolicyLink** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** | Product resource name, in format of \&quot;product/product-durable-id\&quot; | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**SearchKeywords** | Pointer to **[]string** |  | [optional] 
**SearchResultSummary** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**SupportContact** | Pointer to [**AzureMarketplaceContact**](AzureMarketplaceContact.md) |  | [optional] 
**Title** | Pointer to **string** | Max string length is 200. | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplaceListing

`func NewAzureMarketplaceListing() *AzureMarketplaceListing`

NewAzureMarketplaceListing instantiates a new AzureMarketplaceListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceListingWithDefaults

`func NewAzureMarketplaceListingWithDefaults() *AzureMarketplaceListing`

NewAzureMarketplaceListingWithDefaults instantiates a new AzureMarketplaceListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplaceListing) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplaceListing) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplaceListing) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplaceListing) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetCloudSolutionProviderContact

`func (o *AzureMarketplaceListing) GetCloudSolutionProviderContact() AzureMarketplaceContact`

GetCloudSolutionProviderContact returns the CloudSolutionProviderContact field if non-nil, zero value otherwise.

### GetCloudSolutionProviderContactOk

`func (o *AzureMarketplaceListing) GetCloudSolutionProviderContactOk() (*AzureMarketplaceContact, bool)`

GetCloudSolutionProviderContactOk returns a tuple with the CloudSolutionProviderContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSolutionProviderContact

`func (o *AzureMarketplaceListing) SetCloudSolutionProviderContact(v AzureMarketplaceContact)`

SetCloudSolutionProviderContact sets CloudSolutionProviderContact field to given value.

### HasCloudSolutionProviderContact

`func (o *AzureMarketplaceListing) HasCloudSolutionProviderContact() bool`

HasCloudSolutionProviderContact returns a boolean if a field has been set.

### GetCloudSolutionProviderMarketingMaterials

`func (o *AzureMarketplaceListing) GetCloudSolutionProviderMarketingMaterials() string`

GetCloudSolutionProviderMarketingMaterials returns the CloudSolutionProviderMarketingMaterials field if non-nil, zero value otherwise.

### GetCloudSolutionProviderMarketingMaterialsOk

`func (o *AzureMarketplaceListing) GetCloudSolutionProviderMarketingMaterialsOk() (*string, bool)`

GetCloudSolutionProviderMarketingMaterialsOk returns a tuple with the CloudSolutionProviderMarketingMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSolutionProviderMarketingMaterials

`func (o *AzureMarketplaceListing) SetCloudSolutionProviderMarketingMaterials(v string)`

SetCloudSolutionProviderMarketingMaterials sets CloudSolutionProviderMarketingMaterials field to given value.

### HasCloudSolutionProviderMarketingMaterials

`func (o *AzureMarketplaceListing) HasCloudSolutionProviderMarketingMaterials() bool`

HasCloudSolutionProviderMarketingMaterials returns a boolean if a field has been set.

### GetDescription

`func (o *AzureMarketplaceListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureMarketplaceListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureMarketplaceListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureMarketplaceListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEngineeringContact

`func (o *AzureMarketplaceListing) GetEngineeringContact() AzureMarketplaceContact`

GetEngineeringContact returns the EngineeringContact field if non-nil, zero value otherwise.

### GetEngineeringContactOk

`func (o *AzureMarketplaceListing) GetEngineeringContactOk() (*AzureMarketplaceContact, bool)`

GetEngineeringContactOk returns a tuple with the EngineeringContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineeringContact

`func (o *AzureMarketplaceListing) SetEngineeringContact(v AzureMarketplaceContact)`

SetEngineeringContact sets EngineeringContact field to given value.

### HasEngineeringContact

`func (o *AzureMarketplaceListing) HasEngineeringContact() bool`

HasEngineeringContact returns a boolean if a field has been set.

### GetGeneralLinks

`func (o *AzureMarketplaceListing) GetGeneralLinks() []AzureMarketplaceGeneralLink`

GetGeneralLinks returns the GeneralLinks field if non-nil, zero value otherwise.

### GetGeneralLinksOk

`func (o *AzureMarketplaceListing) GetGeneralLinksOk() (*[]AzureMarketplaceGeneralLink, bool)`

GetGeneralLinksOk returns a tuple with the GeneralLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralLinks

`func (o *AzureMarketplaceListing) SetGeneralLinks(v []AzureMarketplaceGeneralLink)`

SetGeneralLinks sets GeneralLinks field to given value.

### HasGeneralLinks

`func (o *AzureMarketplaceListing) HasGeneralLinks() bool`

HasGeneralLinks returns a boolean if a field has been set.

### GetGettingStartedInstructions

`func (o *AzureMarketplaceListing) GetGettingStartedInstructions() string`

GetGettingStartedInstructions returns the GettingStartedInstructions field if non-nil, zero value otherwise.

### GetGettingStartedInstructionsOk

`func (o *AzureMarketplaceListing) GetGettingStartedInstructionsOk() (*string, bool)`

GetGettingStartedInstructionsOk returns a tuple with the GettingStartedInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettingStartedInstructions

`func (o *AzureMarketplaceListing) SetGettingStartedInstructions(v string)`

SetGettingStartedInstructions sets GettingStartedInstructions field to given value.

### HasGettingStartedInstructions

`func (o *AzureMarketplaceListing) HasGettingStartedInstructions() bool`

HasGettingStartedInstructions returns a boolean if a field has been set.

### GetGloabalSupportWebsite

`func (o *AzureMarketplaceListing) GetGloabalSupportWebsite() string`

GetGloabalSupportWebsite returns the GloabalSupportWebsite field if non-nil, zero value otherwise.

### GetGloabalSupportWebsiteOk

`func (o *AzureMarketplaceListing) GetGloabalSupportWebsiteOk() (*string, bool)`

GetGloabalSupportWebsiteOk returns a tuple with the GloabalSupportWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGloabalSupportWebsite

`func (o *AzureMarketplaceListing) SetGloabalSupportWebsite(v string)`

SetGloabalSupportWebsite sets GloabalSupportWebsite field to given value.

### HasGloabalSupportWebsite

`func (o *AzureMarketplaceListing) HasGloabalSupportWebsite() bool`

HasGloabalSupportWebsite returns a boolean if a field has been set.

### GetGovernmentSupportWebsite

`func (o *AzureMarketplaceListing) GetGovernmentSupportWebsite() string`

GetGovernmentSupportWebsite returns the GovernmentSupportWebsite field if non-nil, zero value otherwise.

### GetGovernmentSupportWebsiteOk

`func (o *AzureMarketplaceListing) GetGovernmentSupportWebsiteOk() (*string, bool)`

GetGovernmentSupportWebsiteOk returns a tuple with the GovernmentSupportWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovernmentSupportWebsite

`func (o *AzureMarketplaceListing) SetGovernmentSupportWebsite(v string)`

SetGovernmentSupportWebsite sets GovernmentSupportWebsite field to given value.

### HasGovernmentSupportWebsite

`func (o *AzureMarketplaceListing) HasGovernmentSupportWebsite() bool`

HasGovernmentSupportWebsite returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AzureMarketplaceListing) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AzureMarketplaceListing) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AzureMarketplaceListing) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AzureMarketplaceListing) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguageId

`func (o *AzureMarketplaceListing) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *AzureMarketplaceListing) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *AzureMarketplaceListing) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *AzureMarketplaceListing) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplaceListing) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplaceListing) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplaceListing) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplaceListing) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetPrivacyPolicyLink

`func (o *AzureMarketplaceListing) GetPrivacyPolicyLink() string`

GetPrivacyPolicyLink returns the PrivacyPolicyLink field if non-nil, zero value otherwise.

### GetPrivacyPolicyLinkOk

`func (o *AzureMarketplaceListing) GetPrivacyPolicyLinkOk() (*string, bool)`

GetPrivacyPolicyLinkOk returns a tuple with the PrivacyPolicyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyLink

`func (o *AzureMarketplaceListing) SetPrivacyPolicyLink(v string)`

SetPrivacyPolicyLink sets PrivacyPolicyLink field to given value.

### HasPrivacyPolicyLink

`func (o *AzureMarketplaceListing) HasPrivacyPolicyLink() bool`

HasPrivacyPolicyLink returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplaceListing) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplaceListing) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplaceListing) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplaceListing) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplaceListing) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplaceListing) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplaceListing) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplaceListing) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetSearchKeywords

`func (o *AzureMarketplaceListing) GetSearchKeywords() []string`

GetSearchKeywords returns the SearchKeywords field if non-nil, zero value otherwise.

### GetSearchKeywordsOk

`func (o *AzureMarketplaceListing) GetSearchKeywordsOk() (*[]string, bool)`

GetSearchKeywordsOk returns a tuple with the SearchKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeywords

`func (o *AzureMarketplaceListing) SetSearchKeywords(v []string)`

SetSearchKeywords sets SearchKeywords field to given value.

### HasSearchKeywords

`func (o *AzureMarketplaceListing) HasSearchKeywords() bool`

HasSearchKeywords returns a boolean if a field has been set.

### GetSearchResultSummary

`func (o *AzureMarketplaceListing) GetSearchResultSummary() string`

GetSearchResultSummary returns the SearchResultSummary field if non-nil, zero value otherwise.

### GetSearchResultSummaryOk

`func (o *AzureMarketplaceListing) GetSearchResultSummaryOk() (*string, bool)`

GetSearchResultSummaryOk returns a tuple with the SearchResultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultSummary

`func (o *AzureMarketplaceListing) SetSearchResultSummary(v string)`

SetSearchResultSummary sets SearchResultSummary field to given value.

### HasSearchResultSummary

`func (o *AzureMarketplaceListing) HasSearchResultSummary() bool`

HasSearchResultSummary returns a boolean if a field has been set.

### GetShortDescription

`func (o *AzureMarketplaceListing) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AzureMarketplaceListing) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AzureMarketplaceListing) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *AzureMarketplaceListing) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSupportContact

`func (o *AzureMarketplaceListing) GetSupportContact() AzureMarketplaceContact`

GetSupportContact returns the SupportContact field if non-nil, zero value otherwise.

### GetSupportContactOk

`func (o *AzureMarketplaceListing) GetSupportContactOk() (*AzureMarketplaceContact, bool)`

GetSupportContactOk returns a tuple with the SupportContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportContact

`func (o *AzureMarketplaceListing) SetSupportContact(v AzureMarketplaceContact)`

SetSupportContact sets SupportContact field to given value.

### HasSupportContact

`func (o *AzureMarketplaceListing) HasSupportContact() bool`

HasSupportContact returns a boolean if a field has been set.

### GetTitle

`func (o *AzureMarketplaceListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AzureMarketplaceListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AzureMarketplaceListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AzureMarketplaceListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplaceListing) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplaceListing) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplaceListing) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplaceListing) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


