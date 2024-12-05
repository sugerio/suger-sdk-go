# ProductInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlibabaProduct** | Pointer to [**AlibabaMarketplaceProduct**](AlibabaMarketplaceProduct.md) |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**AwsAmiProduct** | Pointer to [**AwsProduct**](AwsProduct.md) |  | [optional] 
**AwsContainerProduct** | Pointer to [**AwsProduct**](AwsProduct.md) |  | [optional] 
**AwsProfessionalServicesProduct** | Pointer to [**AwsProduct**](AwsProduct.md) |  | [optional] 
**AwsSaasProduct** | Pointer to [**AwsProduct**](AwsProduct.md) |  | [optional] 
**AwsSnsSubscriptions** | Pointer to [**[]AwsSnsSubscription**](AwsSnsSubscription.md) |  | [optional] 
**AzureProduct** | Pointer to [**AzureProduct**](AzureProduct.md) |  | [optional] 
**AzureProductResource** | Pointer to [**AzureMarketplaceProductResource**](AzureMarketplaceProductResource.md) |  | [optional] 
**Commits** | Pointer to [**[]CommitDimension**](CommitDimension.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Dimensions** | Pointer to [**[]MeteringDimension**](MeteringDimension.md) |  | [optional] 
**EulaType** | Pointer to [**EulaType**](EulaType.md) | The public offer&#39;s EULA type. | [optional] 
**EulaUrl** | Pointer to **string** | The public offer&#39;s EULA URL. | [optional] 
**GcpProduct** | Pointer to [**GcpMarketplaceProduct**](GcpMarketplaceProduct.md) |  | [optional] 
**RefundCancellationPolicy** | Pointer to **string** |  | [optional] 
**SellerNotes** | Pointer to **string** |  | [optional] 
**StripeProduct** | Pointer to [**StripeProduct**](StripeProduct.md) |  | [optional] 

## Methods

### NewProductInfo

`func NewProductInfo() *ProductInfo`

NewProductInfo instantiates a new ProductInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInfoWithDefaults

`func NewProductInfoWithDefaults() *ProductInfo`

NewProductInfoWithDefaults instantiates a new ProductInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlibabaProduct

`func (o *ProductInfo) GetAlibabaProduct() AlibabaMarketplaceProduct`

GetAlibabaProduct returns the AlibabaProduct field if non-nil, zero value otherwise.

### GetAlibabaProductOk

`func (o *ProductInfo) GetAlibabaProductOk() (*AlibabaMarketplaceProduct, bool)`

GetAlibabaProductOk returns a tuple with the AlibabaProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlibabaProduct

`func (o *ProductInfo) SetAlibabaProduct(v AlibabaMarketplaceProduct)`

SetAlibabaProduct sets AlibabaProduct field to given value.

### HasAlibabaProduct

`func (o *ProductInfo) HasAlibabaProduct() bool`

HasAlibabaProduct returns a boolean if a field has been set.

### GetAttributes

`func (o *ProductInfo) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ProductInfo) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ProductInfo) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ProductInfo) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAwsAmiProduct

`func (o *ProductInfo) GetAwsAmiProduct() AwsProduct`

GetAwsAmiProduct returns the AwsAmiProduct field if non-nil, zero value otherwise.

### GetAwsAmiProductOk

`func (o *ProductInfo) GetAwsAmiProductOk() (*AwsProduct, bool)`

GetAwsAmiProductOk returns a tuple with the AwsAmiProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAmiProduct

`func (o *ProductInfo) SetAwsAmiProduct(v AwsProduct)`

SetAwsAmiProduct sets AwsAmiProduct field to given value.

### HasAwsAmiProduct

`func (o *ProductInfo) HasAwsAmiProduct() bool`

HasAwsAmiProduct returns a boolean if a field has been set.

### GetAwsContainerProduct

`func (o *ProductInfo) GetAwsContainerProduct() AwsProduct`

GetAwsContainerProduct returns the AwsContainerProduct field if non-nil, zero value otherwise.

### GetAwsContainerProductOk

`func (o *ProductInfo) GetAwsContainerProductOk() (*AwsProduct, bool)`

GetAwsContainerProductOk returns a tuple with the AwsContainerProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsContainerProduct

`func (o *ProductInfo) SetAwsContainerProduct(v AwsProduct)`

SetAwsContainerProduct sets AwsContainerProduct field to given value.

### HasAwsContainerProduct

`func (o *ProductInfo) HasAwsContainerProduct() bool`

HasAwsContainerProduct returns a boolean if a field has been set.

### GetAwsProfessionalServicesProduct

`func (o *ProductInfo) GetAwsProfessionalServicesProduct() AwsProduct`

GetAwsProfessionalServicesProduct returns the AwsProfessionalServicesProduct field if non-nil, zero value otherwise.

### GetAwsProfessionalServicesProductOk

`func (o *ProductInfo) GetAwsProfessionalServicesProductOk() (*AwsProduct, bool)`

GetAwsProfessionalServicesProductOk returns a tuple with the AwsProfessionalServicesProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsProfessionalServicesProduct

`func (o *ProductInfo) SetAwsProfessionalServicesProduct(v AwsProduct)`

SetAwsProfessionalServicesProduct sets AwsProfessionalServicesProduct field to given value.

### HasAwsProfessionalServicesProduct

`func (o *ProductInfo) HasAwsProfessionalServicesProduct() bool`

HasAwsProfessionalServicesProduct returns a boolean if a field has been set.

### GetAwsSaasProduct

`func (o *ProductInfo) GetAwsSaasProduct() AwsProduct`

GetAwsSaasProduct returns the AwsSaasProduct field if non-nil, zero value otherwise.

### GetAwsSaasProductOk

`func (o *ProductInfo) GetAwsSaasProductOk() (*AwsProduct, bool)`

GetAwsSaasProductOk returns a tuple with the AwsSaasProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSaasProduct

`func (o *ProductInfo) SetAwsSaasProduct(v AwsProduct)`

SetAwsSaasProduct sets AwsSaasProduct field to given value.

### HasAwsSaasProduct

`func (o *ProductInfo) HasAwsSaasProduct() bool`

HasAwsSaasProduct returns a boolean if a field has been set.

### GetAwsSnsSubscriptions

`func (o *ProductInfo) GetAwsSnsSubscriptions() []AwsSnsSubscription`

GetAwsSnsSubscriptions returns the AwsSnsSubscriptions field if non-nil, zero value otherwise.

### GetAwsSnsSubscriptionsOk

`func (o *ProductInfo) GetAwsSnsSubscriptionsOk() (*[]AwsSnsSubscription, bool)`

GetAwsSnsSubscriptionsOk returns a tuple with the AwsSnsSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSnsSubscriptions

`func (o *ProductInfo) SetAwsSnsSubscriptions(v []AwsSnsSubscription)`

SetAwsSnsSubscriptions sets AwsSnsSubscriptions field to given value.

### HasAwsSnsSubscriptions

`func (o *ProductInfo) HasAwsSnsSubscriptions() bool`

HasAwsSnsSubscriptions returns a boolean if a field has been set.

### GetAzureProduct

`func (o *ProductInfo) GetAzureProduct() AzureProduct`

GetAzureProduct returns the AzureProduct field if non-nil, zero value otherwise.

### GetAzureProductOk

`func (o *ProductInfo) GetAzureProductOk() (*AzureProduct, bool)`

GetAzureProductOk returns a tuple with the AzureProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureProduct

`func (o *ProductInfo) SetAzureProduct(v AzureProduct)`

SetAzureProduct sets AzureProduct field to given value.

### HasAzureProduct

`func (o *ProductInfo) HasAzureProduct() bool`

HasAzureProduct returns a boolean if a field has been set.

### GetAzureProductResource

`func (o *ProductInfo) GetAzureProductResource() AzureMarketplaceProductResource`

GetAzureProductResource returns the AzureProductResource field if non-nil, zero value otherwise.

### GetAzureProductResourceOk

`func (o *ProductInfo) GetAzureProductResourceOk() (*AzureMarketplaceProductResource, bool)`

GetAzureProductResourceOk returns a tuple with the AzureProductResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureProductResource

`func (o *ProductInfo) SetAzureProductResource(v AzureMarketplaceProductResource)`

SetAzureProductResource sets AzureProductResource field to given value.

### HasAzureProductResource

`func (o *ProductInfo) HasAzureProductResource() bool`

HasAzureProductResource returns a boolean if a field has been set.

### GetCommits

`func (o *ProductInfo) GetCommits() []CommitDimension`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *ProductInfo) GetCommitsOk() (*[]CommitDimension, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *ProductInfo) SetCommits(v []CommitDimension)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *ProductInfo) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetCurrency

`func (o *ProductInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDimensions

`func (o *ProductInfo) GetDimensions() []MeteringDimension`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *ProductInfo) GetDimensionsOk() (*[]MeteringDimension, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *ProductInfo) SetDimensions(v []MeteringDimension)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *ProductInfo) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetEulaType

`func (o *ProductInfo) GetEulaType() EulaType`

GetEulaType returns the EulaType field if non-nil, zero value otherwise.

### GetEulaTypeOk

`func (o *ProductInfo) GetEulaTypeOk() (*EulaType, bool)`

GetEulaTypeOk returns a tuple with the EulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaType

`func (o *ProductInfo) SetEulaType(v EulaType)`

SetEulaType sets EulaType field to given value.

### HasEulaType

`func (o *ProductInfo) HasEulaType() bool`

HasEulaType returns a boolean if a field has been set.

### GetEulaUrl

`func (o *ProductInfo) GetEulaUrl() string`

GetEulaUrl returns the EulaUrl field if non-nil, zero value otherwise.

### GetEulaUrlOk

`func (o *ProductInfo) GetEulaUrlOk() (*string, bool)`

GetEulaUrlOk returns a tuple with the EulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaUrl

`func (o *ProductInfo) SetEulaUrl(v string)`

SetEulaUrl sets EulaUrl field to given value.

### HasEulaUrl

`func (o *ProductInfo) HasEulaUrl() bool`

HasEulaUrl returns a boolean if a field has been set.

### GetGcpProduct

`func (o *ProductInfo) GetGcpProduct() GcpMarketplaceProduct`

GetGcpProduct returns the GcpProduct field if non-nil, zero value otherwise.

### GetGcpProductOk

`func (o *ProductInfo) GetGcpProductOk() (*GcpMarketplaceProduct, bool)`

GetGcpProductOk returns a tuple with the GcpProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProduct

`func (o *ProductInfo) SetGcpProduct(v GcpMarketplaceProduct)`

SetGcpProduct sets GcpProduct field to given value.

### HasGcpProduct

`func (o *ProductInfo) HasGcpProduct() bool`

HasGcpProduct returns a boolean if a field has been set.

### GetRefundCancellationPolicy

`func (o *ProductInfo) GetRefundCancellationPolicy() string`

GetRefundCancellationPolicy returns the RefundCancellationPolicy field if non-nil, zero value otherwise.

### GetRefundCancellationPolicyOk

`func (o *ProductInfo) GetRefundCancellationPolicyOk() (*string, bool)`

GetRefundCancellationPolicyOk returns a tuple with the RefundCancellationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCancellationPolicy

`func (o *ProductInfo) SetRefundCancellationPolicy(v string)`

SetRefundCancellationPolicy sets RefundCancellationPolicy field to given value.

### HasRefundCancellationPolicy

`func (o *ProductInfo) HasRefundCancellationPolicy() bool`

HasRefundCancellationPolicy returns a boolean if a field has been set.

### GetSellerNotes

`func (o *ProductInfo) GetSellerNotes() string`

GetSellerNotes returns the SellerNotes field if non-nil, zero value otherwise.

### GetSellerNotesOk

`func (o *ProductInfo) GetSellerNotesOk() (*string, bool)`

GetSellerNotesOk returns a tuple with the SellerNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerNotes

`func (o *ProductInfo) SetSellerNotes(v string)`

SetSellerNotes sets SellerNotes field to given value.

### HasSellerNotes

`func (o *ProductInfo) HasSellerNotes() bool`

HasSellerNotes returns a boolean if a field has been set.

### GetStripeProduct

`func (o *ProductInfo) GetStripeProduct() StripeProduct`

GetStripeProduct returns the StripeProduct field if non-nil, zero value otherwise.

### GetStripeProductOk

`func (o *ProductInfo) GetStripeProductOk() (*StripeProduct, bool)`

GetStripeProductOk returns a tuple with the StripeProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeProduct

`func (o *ProductInfo) SetStripeProduct(v StripeProduct)`

SetStripeProduct sets StripeProduct field to given value.

### HasStripeProduct

`func (o *ProductInfo) HasStripeProduct() bool`

HasStripeProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


