# AwsSaasProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**AwsSaasProductDescription**](AwsSaasProductDescription.md) |  | [optional] 
**Dimensions** | Pointer to [**[]AwsSaasProductDimension**](AwsSaasProductDimension.md) |  | [optional] 
**PromotionalResources** | Pointer to [**AwsSaasProductPromotionalResources**](AwsSaasProductPromotionalResources.md) |  | [optional] 
**SupportInformation** | Pointer to [**AwsSaasProductSupportInformation**](AwsSaasProductSupportInformation.md) |  | [optional] 
**Versions** | Pointer to [**[]AwsSaasProductVersion**](AwsSaasProductVersion.md) |  | [optional] 
**DataFeedProductId** | Pointer to **string** | The product Id in AWS Marketplace Data Feed Service. | [optional] 
**ProductId** | Pointer to **string** | AWS Product ID | [optional] 

## Methods

### NewAwsSaasProduct

`func NewAwsSaasProduct() *AwsSaasProduct`

NewAwsSaasProduct instantiates a new AwsSaasProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSaasProductWithDefaults

`func NewAwsSaasProductWithDefaults() *AwsSaasProduct`

NewAwsSaasProductWithDefaults instantiates a new AwsSaasProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AwsSaasProduct) GetDescription() AwsSaasProductDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsSaasProduct) GetDescriptionOk() (*AwsSaasProductDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsSaasProduct) SetDescription(v AwsSaasProductDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsSaasProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDimensions

`func (o *AwsSaasProduct) GetDimensions() []AwsSaasProductDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AwsSaasProduct) GetDimensionsOk() (*[]AwsSaasProductDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AwsSaasProduct) SetDimensions(v []AwsSaasProductDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *AwsSaasProduct) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetPromotionalResources

`func (o *AwsSaasProduct) GetPromotionalResources() AwsSaasProductPromotionalResources`

GetPromotionalResources returns the PromotionalResources field if non-nil, zero value otherwise.

### GetPromotionalResourcesOk

`func (o *AwsSaasProduct) GetPromotionalResourcesOk() (*AwsSaasProductPromotionalResources, bool)`

GetPromotionalResourcesOk returns a tuple with the PromotionalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionalResources

`func (o *AwsSaasProduct) SetPromotionalResources(v AwsSaasProductPromotionalResources)`

SetPromotionalResources sets PromotionalResources field to given value.

### HasPromotionalResources

`func (o *AwsSaasProduct) HasPromotionalResources() bool`

HasPromotionalResources returns a boolean if a field has been set.

### GetSupportInformation

`func (o *AwsSaasProduct) GetSupportInformation() AwsSaasProductSupportInformation`

GetSupportInformation returns the SupportInformation field if non-nil, zero value otherwise.

### GetSupportInformationOk

`func (o *AwsSaasProduct) GetSupportInformationOk() (*AwsSaasProductSupportInformation, bool)`

GetSupportInformationOk returns a tuple with the SupportInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportInformation

`func (o *AwsSaasProduct) SetSupportInformation(v AwsSaasProductSupportInformation)`

SetSupportInformation sets SupportInformation field to given value.

### HasSupportInformation

`func (o *AwsSaasProduct) HasSupportInformation() bool`

HasSupportInformation returns a boolean if a field has been set.

### GetVersions

`func (o *AwsSaasProduct) GetVersions() []AwsSaasProductVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AwsSaasProduct) GetVersionsOk() (*[]AwsSaasProductVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AwsSaasProduct) SetVersions(v []AwsSaasProductVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AwsSaasProduct) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetDataFeedProductId

`func (o *AwsSaasProduct) GetDataFeedProductId() string`

GetDataFeedProductId returns the DataFeedProductId field if non-nil, zero value otherwise.

### GetDataFeedProductIdOk

`func (o *AwsSaasProduct) GetDataFeedProductIdOk() (*string, bool)`

GetDataFeedProductIdOk returns a tuple with the DataFeedProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFeedProductId

`func (o *AwsSaasProduct) SetDataFeedProductId(v string)`

SetDataFeedProductId sets DataFeedProductId field to given value.

### HasDataFeedProductId

`func (o *AwsSaasProduct) HasDataFeedProductId() bool`

HasDataFeedProductId returns a boolean if a field has been set.

### GetProductId

`func (o *AwsSaasProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *AwsSaasProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *AwsSaasProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *AwsSaasProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


