# AzureProductProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalCategories** | Pointer to **[]string** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**ApplicableProducts** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**CustomAmendments** | Pointer to **[]string** |  | [optional] 
**ExtendedProperties** | Pointer to **[]string** |  | [optional] 
**GlobalAmendmentTerms** | Pointer to **string** |  | [optional] 
**HideKeys** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Industries** | Pointer to **[]string** |  | [optional] 
**LeveledCategories** | Pointer to **map[string]interface{}** |  | [optional] 
**LeveledIndustries** | Pointer to **map[string]interface{}** |  | [optional] 
**MarketingOnlyChange** | Pointer to **bool** |  | [optional] 
**ProductTags** | Pointer to **[]string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**SubmissionVersion** | Pointer to **string** |  | [optional] 
**TermsOfUse** | Pointer to **string** |  | [optional] 
**UseEnterpriseContract** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureProductProperty

`func NewAzureProductProperty() *AzureProductProperty`

NewAzureProductProperty instantiates a new AzureProductProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductPropertyWithDefaults

`func NewAzureProductPropertyWithDefaults() *AzureProductProperty`

NewAzureProductPropertyWithDefaults instantiates a new AzureProductProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalCategories

`func (o *AzureProductProperty) GetAdditionalCategories() []string`

GetAdditionalCategories returns the AdditionalCategories field if non-nil, zero value otherwise.

### GetAdditionalCategoriesOk

`func (o *AzureProductProperty) GetAdditionalCategoriesOk() (*[]string, bool)`

GetAdditionalCategoriesOk returns a tuple with the AdditionalCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCategories

`func (o *AzureProductProperty) SetAdditionalCategories(v []string)`

SetAdditionalCategories sets AdditionalCategories field to given value.

### HasAdditionalCategories

`func (o *AzureProductProperty) HasAdditionalCategories() bool`

HasAdditionalCategories returns a boolean if a field has been set.

### GetAppVersion

`func (o *AzureProductProperty) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AzureProductProperty) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AzureProductProperty) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *AzureProductProperty) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetApplicableProducts

`func (o *AzureProductProperty) GetApplicableProducts() []string`

GetApplicableProducts returns the ApplicableProducts field if non-nil, zero value otherwise.

### GetApplicableProductsOk

`func (o *AzureProductProperty) GetApplicableProductsOk() (*[]string, bool)`

GetApplicableProductsOk returns a tuple with the ApplicableProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableProducts

`func (o *AzureProductProperty) SetApplicableProducts(v []string)`

SetApplicableProducts sets ApplicableProducts field to given value.

### HasApplicableProducts

`func (o *AzureProductProperty) HasApplicableProducts() bool`

HasApplicableProducts returns a boolean if a field has been set.

### GetCategories

`func (o *AzureProductProperty) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AzureProductProperty) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AzureProductProperty) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AzureProductProperty) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCustomAmendments

`func (o *AzureProductProperty) GetCustomAmendments() []string`

GetCustomAmendments returns the CustomAmendments field if non-nil, zero value otherwise.

### GetCustomAmendmentsOk

`func (o *AzureProductProperty) GetCustomAmendmentsOk() (*[]string, bool)`

GetCustomAmendmentsOk returns a tuple with the CustomAmendments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAmendments

`func (o *AzureProductProperty) SetCustomAmendments(v []string)`

SetCustomAmendments sets CustomAmendments field to given value.

### HasCustomAmendments

`func (o *AzureProductProperty) HasCustomAmendments() bool`

HasCustomAmendments returns a boolean if a field has been set.

### GetExtendedProperties

`func (o *AzureProductProperty) GetExtendedProperties() []string`

GetExtendedProperties returns the ExtendedProperties field if non-nil, zero value otherwise.

### GetExtendedPropertiesOk

`func (o *AzureProductProperty) GetExtendedPropertiesOk() (*[]string, bool)`

GetExtendedPropertiesOk returns a tuple with the ExtendedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProperties

`func (o *AzureProductProperty) SetExtendedProperties(v []string)`

SetExtendedProperties sets ExtendedProperties field to given value.

### HasExtendedProperties

`func (o *AzureProductProperty) HasExtendedProperties() bool`

HasExtendedProperties returns a boolean if a field has been set.

### GetGlobalAmendmentTerms

`func (o *AzureProductProperty) GetGlobalAmendmentTerms() string`

GetGlobalAmendmentTerms returns the GlobalAmendmentTerms field if non-nil, zero value otherwise.

### GetGlobalAmendmentTermsOk

`func (o *AzureProductProperty) GetGlobalAmendmentTermsOk() (*string, bool)`

GetGlobalAmendmentTermsOk returns a tuple with the GlobalAmendmentTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAmendmentTerms

`func (o *AzureProductProperty) SetGlobalAmendmentTerms(v string)`

SetGlobalAmendmentTerms sets GlobalAmendmentTerms field to given value.

### HasGlobalAmendmentTerms

`func (o *AzureProductProperty) HasGlobalAmendmentTerms() bool`

HasGlobalAmendmentTerms returns a boolean if a field has been set.

### GetHideKeys

`func (o *AzureProductProperty) GetHideKeys() []string`

GetHideKeys returns the HideKeys field if non-nil, zero value otherwise.

### GetHideKeysOk

`func (o *AzureProductProperty) GetHideKeysOk() (*[]string, bool)`

GetHideKeysOk returns a tuple with the HideKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideKeys

`func (o *AzureProductProperty) SetHideKeys(v []string)`

SetHideKeys sets HideKeys field to given value.

### HasHideKeys

`func (o *AzureProductProperty) HasHideKeys() bool`

HasHideKeys returns a boolean if a field has been set.

### GetId

`func (o *AzureProductProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureProductProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureProductProperty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureProductProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndustries

`func (o *AzureProductProperty) GetIndustries() []string`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *AzureProductProperty) GetIndustriesOk() (*[]string, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *AzureProductProperty) SetIndustries(v []string)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *AzureProductProperty) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetLeveledCategories

`func (o *AzureProductProperty) GetLeveledCategories() map[string]interface{}`

GetLeveledCategories returns the LeveledCategories field if non-nil, zero value otherwise.

### GetLeveledCategoriesOk

`func (o *AzureProductProperty) GetLeveledCategoriesOk() (*map[string]interface{}, bool)`

GetLeveledCategoriesOk returns a tuple with the LeveledCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveledCategories

`func (o *AzureProductProperty) SetLeveledCategories(v map[string]interface{})`

SetLeveledCategories sets LeveledCategories field to given value.

### HasLeveledCategories

`func (o *AzureProductProperty) HasLeveledCategories() bool`

HasLeveledCategories returns a boolean if a field has been set.

### GetLeveledIndustries

`func (o *AzureProductProperty) GetLeveledIndustries() map[string]interface{}`

GetLeveledIndustries returns the LeveledIndustries field if non-nil, zero value otherwise.

### GetLeveledIndustriesOk

`func (o *AzureProductProperty) GetLeveledIndustriesOk() (*map[string]interface{}, bool)`

GetLeveledIndustriesOk returns a tuple with the LeveledIndustries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveledIndustries

`func (o *AzureProductProperty) SetLeveledIndustries(v map[string]interface{})`

SetLeveledIndustries sets LeveledIndustries field to given value.

### HasLeveledIndustries

`func (o *AzureProductProperty) HasLeveledIndustries() bool`

HasLeveledIndustries returns a boolean if a field has been set.

### GetMarketingOnlyChange

`func (o *AzureProductProperty) GetMarketingOnlyChange() bool`

GetMarketingOnlyChange returns the MarketingOnlyChange field if non-nil, zero value otherwise.

### GetMarketingOnlyChangeOk

`func (o *AzureProductProperty) GetMarketingOnlyChangeOk() (*bool, bool)`

GetMarketingOnlyChangeOk returns a tuple with the MarketingOnlyChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOnlyChange

`func (o *AzureProductProperty) SetMarketingOnlyChange(v bool)`

SetMarketingOnlyChange sets MarketingOnlyChange field to given value.

### HasMarketingOnlyChange

`func (o *AzureProductProperty) HasMarketingOnlyChange() bool`

HasMarketingOnlyChange returns a boolean if a field has been set.

### GetProductTags

`func (o *AzureProductProperty) GetProductTags() []string`

GetProductTags returns the ProductTags field if non-nil, zero value otherwise.

### GetProductTagsOk

`func (o *AzureProductProperty) GetProductTagsOk() (*[]string, bool)`

GetProductTagsOk returns a tuple with the ProductTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTags

`func (o *AzureProductProperty) SetProductTags(v []string)`

SetProductTags sets ProductTags field to given value.

### HasProductTags

`func (o *AzureProductProperty) HasProductTags() bool`

HasProductTags returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureProductProperty) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureProductProperty) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureProductProperty) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureProductProperty) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSubmissionVersion

`func (o *AzureProductProperty) GetSubmissionVersion() string`

GetSubmissionVersion returns the SubmissionVersion field if non-nil, zero value otherwise.

### GetSubmissionVersionOk

`func (o *AzureProductProperty) GetSubmissionVersionOk() (*string, bool)`

GetSubmissionVersionOk returns a tuple with the SubmissionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionVersion

`func (o *AzureProductProperty) SetSubmissionVersion(v string)`

SetSubmissionVersion sets SubmissionVersion field to given value.

### HasSubmissionVersion

`func (o *AzureProductProperty) HasSubmissionVersion() bool`

HasSubmissionVersion returns a boolean if a field has been set.

### GetTermsOfUse

`func (o *AzureProductProperty) GetTermsOfUse() string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *AzureProductProperty) GetTermsOfUseOk() (*string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *AzureProductProperty) SetTermsOfUse(v string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *AzureProductProperty) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### GetUseEnterpriseContract

`func (o *AzureProductProperty) GetUseEnterpriseContract() bool`

GetUseEnterpriseContract returns the UseEnterpriseContract field if non-nil, zero value otherwise.

### GetUseEnterpriseContractOk

`func (o *AzureProductProperty) GetUseEnterpriseContractOk() (*bool, bool)`

GetUseEnterpriseContractOk returns a tuple with the UseEnterpriseContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnterpriseContract

`func (o *AzureProductProperty) SetUseEnterpriseContract(v bool)`

SetUseEnterpriseContract sets UseEnterpriseContract field to given value.

### HasUseEnterpriseContract

`func (o *AzureProductProperty) HasUseEnterpriseContract() bool`

HasUseEnterpriseContract returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


