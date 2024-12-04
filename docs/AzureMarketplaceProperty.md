# AzureMarketplaceProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**AppVersion** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **map[string][]string** |  | [optional] 
**CloudIndustries** | Pointer to **map[string][]string** |  | [optional] 
**CustomAmendments** | Pointer to [**[]AzureMarketplaceCustomAmendment**](AzureMarketplaceCustomAmendment.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Industries** | Pointer to **map[string][]string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**StandardContractAmendment** | Pointer to **string** |  | [optional] 
**TermsConditions** | Pointer to **string** |  | [optional] 
**TermsOfUse** | Pointer to **string** |  | [optional] 
**TermsOfUseUrl** | Pointer to **string** |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplaceProperty

`func NewAzureMarketplaceProperty() *AzureMarketplaceProperty`

NewAzureMarketplaceProperty instantiates a new AzureMarketplaceProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePropertyWithDefaults

`func NewAzureMarketplacePropertyWithDefaults() *AzureMarketplaceProperty`

NewAzureMarketplacePropertyWithDefaults instantiates a new AzureMarketplaceProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplaceProperty) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplaceProperty) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplaceProperty) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplaceProperty) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAppVersion

`func (o *AzureMarketplaceProperty) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AzureMarketplaceProperty) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AzureMarketplaceProperty) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *AzureMarketplaceProperty) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetCategories

`func (o *AzureMarketplaceProperty) GetCategories() map[string][]string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AzureMarketplaceProperty) GetCategoriesOk() (*map[string][]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AzureMarketplaceProperty) SetCategories(v map[string][]string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AzureMarketplaceProperty) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCloudIndustries

`func (o *AzureMarketplaceProperty) GetCloudIndustries() map[string][]string`

GetCloudIndustries returns the CloudIndustries field if non-nil, zero value otherwise.

### GetCloudIndustriesOk

`func (o *AzureMarketplaceProperty) GetCloudIndustriesOk() (*map[string][]string, bool)`

GetCloudIndustriesOk returns a tuple with the CloudIndustries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudIndustries

`func (o *AzureMarketplaceProperty) SetCloudIndustries(v map[string][]string)`

SetCloudIndustries sets CloudIndustries field to given value.

### HasCloudIndustries

`func (o *AzureMarketplaceProperty) HasCloudIndustries() bool`

HasCloudIndustries returns a boolean if a field has been set.

### GetCustomAmendments

`func (o *AzureMarketplaceProperty) GetCustomAmendments() []AzureMarketplaceCustomAmendment`

GetCustomAmendments returns the CustomAmendments field if non-nil, zero value otherwise.

### GetCustomAmendmentsOk

`func (o *AzureMarketplaceProperty) GetCustomAmendmentsOk() (*[]AzureMarketplaceCustomAmendment, bool)`

GetCustomAmendmentsOk returns a tuple with the CustomAmendments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAmendments

`func (o *AzureMarketplaceProperty) SetCustomAmendments(v []AzureMarketplaceCustomAmendment)`

SetCustomAmendments sets CustomAmendments field to given value.

### HasCustomAmendments

`func (o *AzureMarketplaceProperty) HasCustomAmendments() bool`

HasCustomAmendments returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceProperty) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceProperty) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceProperty) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndustries

`func (o *AzureMarketplaceProperty) GetIndustries() map[string][]string`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *AzureMarketplaceProperty) GetIndustriesOk() (*map[string][]string, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *AzureMarketplaceProperty) SetIndustries(v map[string][]string)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *AzureMarketplaceProperty) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetKind

`func (o *AzureMarketplaceProperty) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AzureMarketplaceProperty) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AzureMarketplaceProperty) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AzureMarketplaceProperty) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplaceProperty) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplaceProperty) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplaceProperty) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplaceProperty) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplaceProperty) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplaceProperty) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplaceProperty) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplaceProperty) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplaceProperty) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplaceProperty) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplaceProperty) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplaceProperty) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetStandardContractAmendment

`func (o *AzureMarketplaceProperty) GetStandardContractAmendment() string`

GetStandardContractAmendment returns the StandardContractAmendment field if non-nil, zero value otherwise.

### GetStandardContractAmendmentOk

`func (o *AzureMarketplaceProperty) GetStandardContractAmendmentOk() (*string, bool)`

GetStandardContractAmendmentOk returns a tuple with the StandardContractAmendment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardContractAmendment

`func (o *AzureMarketplaceProperty) SetStandardContractAmendment(v string)`

SetStandardContractAmendment sets StandardContractAmendment field to given value.

### HasStandardContractAmendment

`func (o *AzureMarketplaceProperty) HasStandardContractAmendment() bool`

HasStandardContractAmendment returns a boolean if a field has been set.

### GetTermsConditions

`func (o *AzureMarketplaceProperty) GetTermsConditions() string`

GetTermsConditions returns the TermsConditions field if non-nil, zero value otherwise.

### GetTermsConditionsOk

`func (o *AzureMarketplaceProperty) GetTermsConditionsOk() (*string, bool)`

GetTermsConditionsOk returns a tuple with the TermsConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsConditions

`func (o *AzureMarketplaceProperty) SetTermsConditions(v string)`

SetTermsConditions sets TermsConditions field to given value.

### HasTermsConditions

`func (o *AzureMarketplaceProperty) HasTermsConditions() bool`

HasTermsConditions returns a boolean if a field has been set.

### GetTermsOfUse

`func (o *AzureMarketplaceProperty) GetTermsOfUse() string`

GetTermsOfUse returns the TermsOfUse field if non-nil, zero value otherwise.

### GetTermsOfUseOk

`func (o *AzureMarketplaceProperty) GetTermsOfUseOk() (*string, bool)`

GetTermsOfUseOk returns a tuple with the TermsOfUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUse

`func (o *AzureMarketplaceProperty) SetTermsOfUse(v string)`

SetTermsOfUse sets TermsOfUse field to given value.

### HasTermsOfUse

`func (o *AzureMarketplaceProperty) HasTermsOfUse() bool`

HasTermsOfUse returns a boolean if a field has been set.

### GetTermsOfUseUrl

`func (o *AzureMarketplaceProperty) GetTermsOfUseUrl() string`

GetTermsOfUseUrl returns the TermsOfUseUrl field if non-nil, zero value otherwise.

### GetTermsOfUseUrlOk

`func (o *AzureMarketplaceProperty) GetTermsOfUseUrlOk() (*string, bool)`

GetTermsOfUseUrlOk returns a tuple with the TermsOfUseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfUseUrl

`func (o *AzureMarketplaceProperty) SetTermsOfUseUrl(v string)`

SetTermsOfUseUrl sets TermsOfUseUrl field to given value.

### HasTermsOfUseUrl

`func (o *AzureMarketplaceProperty) HasTermsOfUseUrl() bool`

HasTermsOfUseUrl returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplaceProperty) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplaceProperty) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplaceProperty) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplaceProperty) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


