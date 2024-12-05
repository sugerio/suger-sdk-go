# AzureMarketplaceListingAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int32** | minimum: 0 | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**LanguageId** | Pointer to **string** | Max string length is 10. | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) | Default value is \&quot;generallyAvailable\&quot;. | [optional] 
**Listing** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** | Product resource name, in format of \&quot;product/product-durable-id\&quot; | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AzureMarketplaceListingAssetType**](AzureMarketplaceListingAssetType.md) |  | [optional] 
**Url** | Pointer to **string** | pattern: \&quot;^https?://\&quot; | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplaceListingAsset

`func NewAzureMarketplaceListingAsset() *AzureMarketplaceListingAsset`

NewAzureMarketplaceListingAsset instantiates a new AzureMarketplaceListingAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceListingAssetWithDefaults

`func NewAzureMarketplaceListingAssetWithDefaults() *AzureMarketplaceListingAsset`

NewAzureMarketplaceListingAssetWithDefaults instantiates a new AzureMarketplaceListingAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplaceListingAsset) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplaceListingAsset) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplaceListingAsset) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplaceListingAsset) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetDescription

`func (o *AzureMarketplaceListingAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AzureMarketplaceListingAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AzureMarketplaceListingAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AzureMarketplaceListingAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *AzureMarketplaceListingAsset) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *AzureMarketplaceListingAsset) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *AzureMarketplaceListingAsset) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *AzureMarketplaceListingAsset) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetFileName

`func (o *AzureMarketplaceListingAsset) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AzureMarketplaceListingAsset) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AzureMarketplaceListingAsset) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AzureMarketplaceListingAsset) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFriendlyName

`func (o *AzureMarketplaceListingAsset) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AzureMarketplaceListingAsset) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AzureMarketplaceListingAsset) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AzureMarketplaceListingAsset) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceListingAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceListingAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceListingAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceListingAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *AzureMarketplaceListingAsset) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AzureMarketplaceListingAsset) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AzureMarketplaceListingAsset) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AzureMarketplaceListingAsset) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLanguageId

`func (o *AzureMarketplaceListingAsset) GetLanguageId() string`

GetLanguageId returns the LanguageId field if non-nil, zero value otherwise.

### GetLanguageIdOk

`func (o *AzureMarketplaceListingAsset) GetLanguageIdOk() (*string, bool)`

GetLanguageIdOk returns a tuple with the LanguageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageId

`func (o *AzureMarketplaceListingAsset) SetLanguageId(v string)`

SetLanguageId sets LanguageId field to given value.

### HasLanguageId

`func (o *AzureMarketplaceListingAsset) HasLanguageId() bool`

HasLanguageId returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplaceListingAsset) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplaceListingAsset) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplaceListingAsset) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplaceListingAsset) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetListing

`func (o *AzureMarketplaceListingAsset) GetListing() string`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *AzureMarketplaceListingAsset) GetListingOk() (*string, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *AzureMarketplaceListingAsset) SetListing(v string)`

SetListing sets Listing field to given value.

### HasListing

`func (o *AzureMarketplaceListingAsset) HasListing() bool`

HasListing returns a boolean if a field has been set.

### GetProduct

`func (o *AzureMarketplaceListingAsset) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AzureMarketplaceListingAsset) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AzureMarketplaceListingAsset) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AzureMarketplaceListingAsset) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplaceListingAsset) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplaceListingAsset) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplaceListingAsset) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplaceListingAsset) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetType

`func (o *AzureMarketplaceListingAsset) GetType() AzureMarketplaceListingAssetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureMarketplaceListingAsset) GetTypeOk() (*AzureMarketplaceListingAssetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureMarketplaceListingAsset) SetType(v AzureMarketplaceListingAssetType)`

SetType sets Type field to given value.

### HasType

`func (o *AzureMarketplaceListingAsset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AzureMarketplaceListingAsset) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AzureMarketplaceListingAsset) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AzureMarketplaceListingAsset) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AzureMarketplaceListingAsset) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplaceListingAsset) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplaceListingAsset) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplaceListingAsset) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplaceListingAsset) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


