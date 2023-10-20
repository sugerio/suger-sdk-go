# ProductInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlibabaProduct** | Pointer to [**AlibabaMarketplaceProduct**](AlibabaMarketplaceProduct.md) |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**AwsSaasProduct** | Pointer to [**AwsSaasProduct**](AwsSaasProduct.md) |  | [optional] 
**AwsSnsSubscriptions** | Pointer to [**[]AwsSnsSubscription**](AwsSnsSubscription.md) |  | [optional] 
**AzureProduct** | Pointer to [**AzureProduct**](AzureProduct.md) |  | [optional] 
**Commits** | Pointer to [**[]CommitDimension**](CommitDimension.md) |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Dimensions** | Pointer to [**[]MeteringDimension**](MeteringDimension.md) |  | [optional] 
**EulaUrl** | Pointer to **string** |  | [optional] 
**GcpProduct** | Pointer to [**GcpMarketplaceProduct**](GcpMarketplaceProduct.md) |  | [optional] 
**RefundCancelationPolicy** | Pointer to **string** |  | [optional] 
**SellerNotes** | Pointer to **string** |  | [optional] 

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

### GetAwsSaasProduct

`func (o *ProductInfo) GetAwsSaasProduct() AwsSaasProduct`

GetAwsSaasProduct returns the AwsSaasProduct field if non-nil, zero value otherwise.

### GetAwsSaasProductOk

`func (o *ProductInfo) GetAwsSaasProductOk() (*AwsSaasProduct, bool)`

GetAwsSaasProductOk returns a tuple with the AwsSaasProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSaasProduct

`func (o *ProductInfo) SetAwsSaasProduct(v AwsSaasProduct)`

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

### GetRefundCancelationPolicy

`func (o *ProductInfo) GetRefundCancelationPolicy() string`

GetRefundCancelationPolicy returns the RefundCancelationPolicy field if non-nil, zero value otherwise.

### GetRefundCancelationPolicyOk

`func (o *ProductInfo) GetRefundCancelationPolicyOk() (*string, bool)`

GetRefundCancelationPolicyOk returns a tuple with the RefundCancelationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundCancelationPolicy

`func (o *ProductInfo) SetRefundCancelationPolicy(v string)`

SetRefundCancelationPolicy sets RefundCancelationPolicy field to given value.

### HasRefundCancelationPolicy

`func (o *ProductInfo) HasRefundCancelationPolicy() bool`

HasRefundCancelationPolicy returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


