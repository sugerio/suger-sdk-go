# AwsProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**AwsProductDescription**](AwsProductDescription.md) |  | [optional] 
**Dimensions** | Pointer to [**[]AwsProductDimension**](AwsProductDimension.md) |  | [optional] 
**PromotionalResources** | Pointer to [**AwsProductPromotionalResources**](AwsProductPromotionalResources.md) |  | [optional] 
**Repositories** | Pointer to [**[]AwsProductRepository**](AwsProductRepository.md) |  | [optional] 
**SignatureVerificationKeys** | Pointer to [**[]AwsProductSignatureVerificationKey**](AwsProductSignatureVerificationKey.md) |  | [optional] 
**SupportInformation** | Pointer to [**AwsProductSupportInformation**](AwsProductSupportInformation.md) |  | [optional] 
**Versions** | Pointer to [**[]AwsProductVersion**](AwsProductVersion.md) |  | [optional] 
**DataFeedProductId** | Pointer to **string** | The product Id in AWS Marketplace Data Feed Service. | [optional] 
**ProductId** | Pointer to **string** | AWS Product ID | [optional] 

## Methods

### NewAwsProduct

`func NewAwsProduct() *AwsProduct`

NewAwsProduct instantiates a new AwsProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProductWithDefaults

`func NewAwsProductWithDefaults() *AwsProduct`

NewAwsProductWithDefaults instantiates a new AwsProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AwsProduct) GetDescription() AwsProductDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsProduct) GetDescriptionOk() (*AwsProductDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsProduct) SetDescription(v AwsProductDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDimensions

`func (o *AwsProduct) GetDimensions() []AwsProductDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AwsProduct) GetDimensionsOk() (*[]AwsProductDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AwsProduct) SetDimensions(v []AwsProductDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *AwsProduct) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetPromotionalResources

`func (o *AwsProduct) GetPromotionalResources() AwsProductPromotionalResources`

GetPromotionalResources returns the PromotionalResources field if non-nil, zero value otherwise.

### GetPromotionalResourcesOk

`func (o *AwsProduct) GetPromotionalResourcesOk() (*AwsProductPromotionalResources, bool)`

GetPromotionalResourcesOk returns a tuple with the PromotionalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionalResources

`func (o *AwsProduct) SetPromotionalResources(v AwsProductPromotionalResources)`

SetPromotionalResources sets PromotionalResources field to given value.

### HasPromotionalResources

`func (o *AwsProduct) HasPromotionalResources() bool`

HasPromotionalResources returns a boolean if a field has been set.

### GetRepositories

`func (o *AwsProduct) GetRepositories() []AwsProductRepository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AwsProduct) GetRepositoriesOk() (*[]AwsProductRepository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AwsProduct) SetRepositories(v []AwsProductRepository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *AwsProduct) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetSignatureVerificationKeys

`func (o *AwsProduct) GetSignatureVerificationKeys() []AwsProductSignatureVerificationKey`

GetSignatureVerificationKeys returns the SignatureVerificationKeys field if non-nil, zero value otherwise.

### GetSignatureVerificationKeysOk

`func (o *AwsProduct) GetSignatureVerificationKeysOk() (*[]AwsProductSignatureVerificationKey, bool)`

GetSignatureVerificationKeysOk returns a tuple with the SignatureVerificationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerificationKeys

`func (o *AwsProduct) SetSignatureVerificationKeys(v []AwsProductSignatureVerificationKey)`

SetSignatureVerificationKeys sets SignatureVerificationKeys field to given value.

### HasSignatureVerificationKeys

`func (o *AwsProduct) HasSignatureVerificationKeys() bool`

HasSignatureVerificationKeys returns a boolean if a field has been set.

### GetSupportInformation

`func (o *AwsProduct) GetSupportInformation() AwsProductSupportInformation`

GetSupportInformation returns the SupportInformation field if non-nil, zero value otherwise.

### GetSupportInformationOk

`func (o *AwsProduct) GetSupportInformationOk() (*AwsProductSupportInformation, bool)`

GetSupportInformationOk returns a tuple with the SupportInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportInformation

`func (o *AwsProduct) SetSupportInformation(v AwsProductSupportInformation)`

SetSupportInformation sets SupportInformation field to given value.

### HasSupportInformation

`func (o *AwsProduct) HasSupportInformation() bool`

HasSupportInformation returns a boolean if a field has been set.

### GetVersions

`func (o *AwsProduct) GetVersions() []AwsProductVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AwsProduct) GetVersionsOk() (*[]AwsProductVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AwsProduct) SetVersions(v []AwsProductVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AwsProduct) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetDataFeedProductId

`func (o *AwsProduct) GetDataFeedProductId() string`

GetDataFeedProductId returns the DataFeedProductId field if non-nil, zero value otherwise.

### GetDataFeedProductIdOk

`func (o *AwsProduct) GetDataFeedProductIdOk() (*string, bool)`

GetDataFeedProductIdOk returns a tuple with the DataFeedProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFeedProductId

`func (o *AwsProduct) SetDataFeedProductId(v string)`

SetDataFeedProductId sets DataFeedProductId field to given value.

### HasDataFeedProductId

`func (o *AwsProduct) HasDataFeedProductId() bool`

HasDataFeedProductId returns a boolean if a field has been set.

### GetProductId

`func (o *AwsProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AwsProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AwsProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AwsProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


