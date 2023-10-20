# AzureProductListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessInformation** | Pointer to **string** |  | [optional] 
**Assets** | Pointer to [**[]AzureProductListingAsset**](AzureProductListingAsset.md) | Not original fields. They are populated by other API calls | [optional] 
**CompatibleProducts** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GettingStartedInstructions** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**LanguageCode** | Pointer to **string** |  | [optional] 
**ListingContacts** | Pointer to [**[]AzureListingContact**](AzureListingContact.md) |  | [optional] 
**ListingUris** | Pointer to [**[]AzureListingUri**](AzureListingUri.md) |  | [optional] 
**ProductDisplayName** | Pointer to **string** |  | [optional] 
**PublisherName** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureProductListing

`func NewAzureProductListing() *AzureProductListing`

NewAzureProductListing instantiates a new AzureProductListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductListingWithDefaults

`func NewAzureProductListingWithDefaults() *AzureProductListing`

NewAzureProductListingWithDefaults instantiates a new AzureProductListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessInformation

`func (o *AzureProductListing) GetAccessInformation() string`

GetAccessInformation returns the AccessInformation field if non-nil, zero value otherwise.

### GetAccessInformationOk

`func (o *AzureProductListing) GetAccessInformationOk() (*string, bool)`

GetAccessInformationOk returns a tuple with the AccessInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessInformation

`func (o *AzureProductListing) SetAccessInformation(v string)`

SetAccessInformation sets AccessInformation field to given value.

### HasAccessInformation

`func (o *AzureProductListing) HasAccessInformation() bool`

HasAccessInformation returns a boolean if a field has been set.

### GetAssets

`func (o *AzureProductListing) GetAssets() []AzureProductListingAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *AzureProductListing) GetAssetsOk() (*[]AzureProductListingAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *AzureProductListing) SetAssets(v []AzureProductListingAsset)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *AzureProductListing) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetCompatibleProducts

`func (o *AzureProductListing) GetCompatibleProducts() []string`

GetCompatibleProducts returns the CompatibleProducts field if non-nil, zero value otherwise.

### GetCompatibleProductsOk

`func (o *AzureProductListing) GetCompatibleProductsOk() (*[]string, bool)`

GetCompatibleProductsOk returns a tuple with the CompatibleProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleProducts

`func (o *AzureProductListing) SetCompatibleProducts(v []string)`

SetCompatibleProducts sets CompatibleProducts field to given value.

### HasCompatibleProducts

`func (o *AzureProductListing) HasCompatibleProducts() bool`

HasCompatibleProducts returns a boolean if a field has been set.

### GetDescription

`func (o *AzureProductListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureProductListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureProductListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureProductListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGettingStartedInstructions

`func (o *AzureProductListing) GetGettingStartedInstructions() string`

GetGettingStartedInstructions returns the GettingStartedInstructions field if non-nil, zero value otherwise.

### GetGettingStartedInstructionsOk

`func (o *AzureProductListing) GetGettingStartedInstructionsOk() (*string, bool)`

GetGettingStartedInstructionsOk returns a tuple with the GettingStartedInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGettingStartedInstructions

`func (o *AzureProductListing) SetGettingStartedInstructions(v string)`

SetGettingStartedInstructions sets GettingStartedInstructions field to given value.

### HasGettingStartedInstructions

`func (o *AzureProductListing) HasGettingStartedInstructions() bool`

HasGettingStartedInstructions returns a boolean if a field has been set.

### GetId

`func (o *AzureProductListing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureProductListing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureProductListing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureProductListing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeywords

`func (o *AzureProductListing) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *AzureProductListing) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *AzureProductListing) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *AzureProductListing) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLanguageCode

`func (o *AzureProductListing) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *AzureProductListing) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *AzureProductListing) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *AzureProductListing) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetListingContacts

`func (o *AzureProductListing) GetListingContacts() []AzureListingContact`

GetListingContacts returns the ListingContacts field if non-nil, zero value otherwise.

### GetListingContactsOk

`func (o *AzureProductListing) GetListingContactsOk() (*[]AzureListingContact, bool)`

GetListingContactsOk returns a tuple with the ListingContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingContacts

`func (o *AzureProductListing) SetListingContacts(v []AzureListingContact)`

SetListingContacts sets ListingContacts field to given value.

### HasListingContacts

`func (o *AzureProductListing) HasListingContacts() bool`

HasListingContacts returns a boolean if a field has been set.

### GetListingUris

`func (o *AzureProductListing) GetListingUris() []AzureListingUri`

GetListingUris returns the ListingUris field if non-nil, zero value otherwise.

### GetListingUrisOk

`func (o *AzureProductListing) GetListingUrisOk() (*[]AzureListingUri, bool)`

GetListingUrisOk returns a tuple with the ListingUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingUris

`func (o *AzureProductListing) SetListingUris(v []AzureListingUri)`

SetListingUris sets ListingUris field to given value.

### HasListingUris

`func (o *AzureProductListing) HasListingUris() bool`

HasListingUris returns a boolean if a field has been set.

### GetProductDisplayName

`func (o *AzureProductListing) GetProductDisplayName() string`

GetProductDisplayName returns the ProductDisplayName field if non-nil, zero value otherwise.

### GetProductDisplayNameOk

`func (o *AzureProductListing) GetProductDisplayNameOk() (*string, bool)`

GetProductDisplayNameOk returns a tuple with the ProductDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDisplayName

`func (o *AzureProductListing) SetProductDisplayName(v string)`

SetProductDisplayName sets ProductDisplayName field to given value.

### HasProductDisplayName

`func (o *AzureProductListing) HasProductDisplayName() bool`

HasProductDisplayName returns a boolean if a field has been set.

### GetPublisherName

`func (o *AzureProductListing) GetPublisherName() string`

GetPublisherName returns the PublisherName field if non-nil, zero value otherwise.

### GetPublisherNameOk

`func (o *AzureProductListing) GetPublisherNameOk() (*string, bool)`

GetPublisherNameOk returns a tuple with the PublisherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherName

`func (o *AzureProductListing) SetPublisherName(v string)`

SetPublisherName sets PublisherName field to given value.

### HasPublisherName

`func (o *AzureProductListing) HasPublisherName() bool`

HasPublisherName returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureProductListing) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureProductListing) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureProductListing) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureProductListing) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetShortDescription

`func (o *AzureProductListing) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AzureProductListing) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AzureProductListing) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *AzureProductListing) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetSummary

`func (o *AzureProductListing) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AzureProductListing) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AzureProductListing) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AzureProductListing) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AzureProductListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AzureProductListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AzureProductListing) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AzureProductListing) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


