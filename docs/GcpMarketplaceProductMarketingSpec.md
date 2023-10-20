# GcpMarketplaceProductMarketingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DisplayNames** | Pointer to **[]string** |  | [optional] 
**DocumentationSpecs** | Pointer to [**[]GcpMarketplaceProductDocumentationSpec**](GcpMarketplaceProductDocumentationSpec.md) |  | [optional] 
**EulaUrl** | Pointer to **string** |  | [optional] 
**ExternalLicenseSpecs** | Pointer to [**[]GcpMarketplaceProductLicenseSpec**](GcpMarketplaceProductLicenseSpec.md) |  | [optional] 
**ExternalMarketingUrl** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** | in format of \&quot;base64://...\&quot; | [optional] 
**SearchCategories** | Pointer to **[]string** |  | [optional] 
**SearchDescription** | Pointer to **string** |  | [optional] 
**SearchKeywords** | Pointer to **[]string** |  | [optional] 
**SignupUri** | Pointer to **string** |  | [optional] 
**SupportSpec** | Pointer to [**GcpMarketplaceProductSupportSpec**](GcpMarketplaceProductSupportSpec.md) |  | [optional] 
**TagLine** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewGcpMarketplaceProductMarketingSpec

`func NewGcpMarketplaceProductMarketingSpec() *GcpMarketplaceProductMarketingSpec`

NewGcpMarketplaceProductMarketingSpec instantiates a new GcpMarketplaceProductMarketingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductMarketingSpecWithDefaults

`func NewGcpMarketplaceProductMarketingSpecWithDefaults() *GcpMarketplaceProductMarketingSpec`

NewGcpMarketplaceProductMarketingSpecWithDefaults instantiates a new GcpMarketplaceProductMarketingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GcpMarketplaceProductMarketingSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GcpMarketplaceProductMarketingSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GcpMarketplaceProductMarketingSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GcpMarketplaceProductMarketingSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayNames

`func (o *GcpMarketplaceProductMarketingSpec) GetDisplayNames() []string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *GcpMarketplaceProductMarketingSpec) GetDisplayNamesOk() (*[]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *GcpMarketplaceProductMarketingSpec) SetDisplayNames(v []string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *GcpMarketplaceProductMarketingSpec) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### GetDocumentationSpecs

`func (o *GcpMarketplaceProductMarketingSpec) GetDocumentationSpecs() []GcpMarketplaceProductDocumentationSpec`

GetDocumentationSpecs returns the DocumentationSpecs field if non-nil, zero value otherwise.

### GetDocumentationSpecsOk

`func (o *GcpMarketplaceProductMarketingSpec) GetDocumentationSpecsOk() (*[]GcpMarketplaceProductDocumentationSpec, bool)`

GetDocumentationSpecsOk returns a tuple with the DocumentationSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationSpecs

`func (o *GcpMarketplaceProductMarketingSpec) SetDocumentationSpecs(v []GcpMarketplaceProductDocumentationSpec)`

SetDocumentationSpecs sets DocumentationSpecs field to given value.

### HasDocumentationSpecs

`func (o *GcpMarketplaceProductMarketingSpec) HasDocumentationSpecs() bool`

HasDocumentationSpecs returns a boolean if a field has been set.

### GetEulaUrl

`func (o *GcpMarketplaceProductMarketingSpec) GetEulaUrl() string`

GetEulaUrl returns the EulaUrl field if non-nil, zero value otherwise.

### GetEulaUrlOk

`func (o *GcpMarketplaceProductMarketingSpec) GetEulaUrlOk() (*string, bool)`

GetEulaUrlOk returns a tuple with the EulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaUrl

`func (o *GcpMarketplaceProductMarketingSpec) SetEulaUrl(v string)`

SetEulaUrl sets EulaUrl field to given value.

### HasEulaUrl

`func (o *GcpMarketplaceProductMarketingSpec) HasEulaUrl() bool`

HasEulaUrl returns a boolean if a field has been set.

### GetExternalLicenseSpecs

`func (o *GcpMarketplaceProductMarketingSpec) GetExternalLicenseSpecs() []GcpMarketplaceProductLicenseSpec`

GetExternalLicenseSpecs returns the ExternalLicenseSpecs field if non-nil, zero value otherwise.

### GetExternalLicenseSpecsOk

`func (o *GcpMarketplaceProductMarketingSpec) GetExternalLicenseSpecsOk() (*[]GcpMarketplaceProductLicenseSpec, bool)`

GetExternalLicenseSpecsOk returns a tuple with the ExternalLicenseSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLicenseSpecs

`func (o *GcpMarketplaceProductMarketingSpec) SetExternalLicenseSpecs(v []GcpMarketplaceProductLicenseSpec)`

SetExternalLicenseSpecs sets ExternalLicenseSpecs field to given value.

### HasExternalLicenseSpecs

`func (o *GcpMarketplaceProductMarketingSpec) HasExternalLicenseSpecs() bool`

HasExternalLicenseSpecs returns a boolean if a field has been set.

### GetExternalMarketingUrl

`func (o *GcpMarketplaceProductMarketingSpec) GetExternalMarketingUrl() string`

GetExternalMarketingUrl returns the ExternalMarketingUrl field if non-nil, zero value otherwise.

### GetExternalMarketingUrlOk

`func (o *GcpMarketplaceProductMarketingSpec) GetExternalMarketingUrlOk() (*string, bool)`

GetExternalMarketingUrlOk returns a tuple with the ExternalMarketingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMarketingUrl

`func (o *GcpMarketplaceProductMarketingSpec) SetExternalMarketingUrl(v string)`

SetExternalMarketingUrl sets ExternalMarketingUrl field to given value.

### HasExternalMarketingUrl

`func (o *GcpMarketplaceProductMarketingSpec) HasExternalMarketingUrl() bool`

HasExternalMarketingUrl returns a boolean if a field has been set.

### GetIcon

`func (o *GcpMarketplaceProductMarketingSpec) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GcpMarketplaceProductMarketingSpec) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GcpMarketplaceProductMarketingSpec) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GcpMarketplaceProductMarketingSpec) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetSearchCategories

`func (o *GcpMarketplaceProductMarketingSpec) GetSearchCategories() []string`

GetSearchCategories returns the SearchCategories field if non-nil, zero value otherwise.

### GetSearchCategoriesOk

`func (o *GcpMarketplaceProductMarketingSpec) GetSearchCategoriesOk() (*[]string, bool)`

GetSearchCategoriesOk returns a tuple with the SearchCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCategories

`func (o *GcpMarketplaceProductMarketingSpec) SetSearchCategories(v []string)`

SetSearchCategories sets SearchCategories field to given value.

### HasSearchCategories

`func (o *GcpMarketplaceProductMarketingSpec) HasSearchCategories() bool`

HasSearchCategories returns a boolean if a field has been set.

### GetSearchDescription

`func (o *GcpMarketplaceProductMarketingSpec) GetSearchDescription() string`

GetSearchDescription returns the SearchDescription field if non-nil, zero value otherwise.

### GetSearchDescriptionOk

`func (o *GcpMarketplaceProductMarketingSpec) GetSearchDescriptionOk() (*string, bool)`

GetSearchDescriptionOk returns a tuple with the SearchDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDescription

`func (o *GcpMarketplaceProductMarketingSpec) SetSearchDescription(v string)`

SetSearchDescription sets SearchDescription field to given value.

### HasSearchDescription

`func (o *GcpMarketplaceProductMarketingSpec) HasSearchDescription() bool`

HasSearchDescription returns a boolean if a field has been set.

### GetSearchKeywords

`func (o *GcpMarketplaceProductMarketingSpec) GetSearchKeywords() []string`

GetSearchKeywords returns the SearchKeywords field if non-nil, zero value otherwise.

### GetSearchKeywordsOk

`func (o *GcpMarketplaceProductMarketingSpec) GetSearchKeywordsOk() (*[]string, bool)`

GetSearchKeywordsOk returns a tuple with the SearchKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeywords

`func (o *GcpMarketplaceProductMarketingSpec) SetSearchKeywords(v []string)`

SetSearchKeywords sets SearchKeywords field to given value.

### HasSearchKeywords

`func (o *GcpMarketplaceProductMarketingSpec) HasSearchKeywords() bool`

HasSearchKeywords returns a boolean if a field has been set.

### GetSignupUri

`func (o *GcpMarketplaceProductMarketingSpec) GetSignupUri() string`

GetSignupUri returns the SignupUri field if non-nil, zero value otherwise.

### GetSignupUriOk

`func (o *GcpMarketplaceProductMarketingSpec) GetSignupUriOk() (*string, bool)`

GetSignupUriOk returns a tuple with the SignupUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupUri

`func (o *GcpMarketplaceProductMarketingSpec) SetSignupUri(v string)`

SetSignupUri sets SignupUri field to given value.

### HasSignupUri

`func (o *GcpMarketplaceProductMarketingSpec) HasSignupUri() bool`

HasSignupUri returns a boolean if a field has been set.

### GetSupportSpec

`func (o *GcpMarketplaceProductMarketingSpec) GetSupportSpec() GcpMarketplaceProductSupportSpec`

GetSupportSpec returns the SupportSpec field if non-nil, zero value otherwise.

### GetSupportSpecOk

`func (o *GcpMarketplaceProductMarketingSpec) GetSupportSpecOk() (*GcpMarketplaceProductSupportSpec, bool)`

GetSupportSpecOk returns a tuple with the SupportSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSpec

`func (o *GcpMarketplaceProductMarketingSpec) SetSupportSpec(v GcpMarketplaceProductSupportSpec)`

SetSupportSpec sets SupportSpec field to given value.

### HasSupportSpec

`func (o *GcpMarketplaceProductMarketingSpec) HasSupportSpec() bool`

HasSupportSpec returns a boolean if a field has been set.

### GetTagLine

`func (o *GcpMarketplaceProductMarketingSpec) GetTagLine() string`

GetTagLine returns the TagLine field if non-nil, zero value otherwise.

### GetTagLineOk

`func (o *GcpMarketplaceProductMarketingSpec) GetTagLineOk() (*string, bool)`

GetTagLineOk returns a tuple with the TagLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagLine

`func (o *GcpMarketplaceProductMarketingSpec) SetTagLine(v string)`

SetTagLine sets TagLine field to given value.

### HasTagLine

`func (o *GcpMarketplaceProductMarketingSpec) HasTagLine() bool`

HasTagLine returns a boolean if a field has been set.

### GetTitle

`func (o *GcpMarketplaceProductMarketingSpec) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GcpMarketplaceProductMarketingSpec) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GcpMarketplaceProductMarketingSpec) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GcpMarketplaceProductMarketingSpec) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


