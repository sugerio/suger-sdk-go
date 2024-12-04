# AzureMarketplaceProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** | The Product Display Name | [optional] 
**Id** | Pointer to **string** | in format of \&quot;product/product-durable-id\&quot; | [optional] 
**Identity** | Pointer to [**AzureMarketplaceIdentity**](AzureMarketplaceIdentity.md) |  | [optional] 
**LifecycleState** | Pointer to [**AzureMarketplaceResourceLifecycleState**](AzureMarketplaceResourceLifecycleState.md) |  | [optional] 
**ProductGroup** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AzureMarketplaceProductType**](AzureMarketplaceProductType.md) |  | [optional] 
**Validations** | Pointer to [**[]AzureMarketplaceValidation**](AzureMarketplaceValidation.md) |  | [optional] 

## Methods

### NewAzureMarketplaceProduct

`func NewAzureMarketplaceProduct() *AzureMarketplaceProduct`

NewAzureMarketplaceProduct instantiates a new AzureMarketplaceProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplaceProductWithDefaults

`func NewAzureMarketplaceProductWithDefaults() *AzureMarketplaceProduct`

NewAzureMarketplaceProductWithDefaults instantiates a new AzureMarketplaceProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *AzureMarketplaceProduct) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *AzureMarketplaceProduct) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *AzureMarketplaceProduct) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *AzureMarketplaceProduct) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAlias

`func (o *AzureMarketplaceProduct) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *AzureMarketplaceProduct) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *AzureMarketplaceProduct) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *AzureMarketplaceProduct) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetId

`func (o *AzureMarketplaceProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureMarketplaceProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureMarketplaceProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureMarketplaceProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentity

`func (o *AzureMarketplaceProduct) GetIdentity() AzureMarketplaceIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *AzureMarketplaceProduct) GetIdentityOk() (*AzureMarketplaceIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *AzureMarketplaceProduct) SetIdentity(v AzureMarketplaceIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *AzureMarketplaceProduct) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLifecycleState

`func (o *AzureMarketplaceProduct) GetLifecycleState() AzureMarketplaceResourceLifecycleState`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *AzureMarketplaceProduct) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *AzureMarketplaceProduct) SetLifecycleState(v AzureMarketplaceResourceLifecycleState)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *AzureMarketplaceProduct) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetProductGroup

`func (o *AzureMarketplaceProduct) GetProductGroup() string`

GetProductGroup returns the ProductGroup field if non-nil, zero value otherwise.

### GetProductGroupOk

`func (o *AzureMarketplaceProduct) GetProductGroupOk() (*string, bool)`

GetProductGroupOk returns a tuple with the ProductGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroup

`func (o *AzureMarketplaceProduct) SetProductGroup(v string)`

SetProductGroup sets ProductGroup field to given value.

### HasProductGroup

`func (o *AzureMarketplaceProduct) HasProductGroup() bool`

HasProductGroup returns a boolean if a field has been set.

### GetResourceName

`func (o *AzureMarketplaceProduct) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *AzureMarketplaceProduct) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *AzureMarketplaceProduct) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *AzureMarketplaceProduct) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetType

`func (o *AzureMarketplaceProduct) GetType() AzureMarketplaceProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AzureMarketplaceProduct) GetTypeOk() (*AzureMarketplaceProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AzureMarketplaceProduct) SetType(v AzureMarketplaceProductType)`

SetType sets Type field to given value.

### HasType

`func (o *AzureMarketplaceProduct) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidations

`func (o *AzureMarketplaceProduct) GetValidations() []AzureMarketplaceValidation`

GetValidations returns the Validations field if non-nil, zero value otherwise.

### GetValidationsOk

`func (o *AzureMarketplaceProduct) GetValidationsOk() (*[]AzureMarketplaceValidation, bool)`

GetValidationsOk returns a tuple with the Validations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidations

`func (o *AzureMarketplaceProduct) SetValidations(v []AzureMarketplaceValidation)`

SetValidations sets Validations field to given value.

### HasValidations

`func (o *AzureMarketplaceProduct) HasValidations() bool`

HasValidations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


