# StripeProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Whether the product is currently available for purchase. | [optional] 
**Created** | Pointer to **int32** | Time at which the object was created. Measured in seconds since the Unix epoch. | [optional] 
**Description** | Pointer to **string** | The product&#39;s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the product in Stripe. | [optional] 
**Images** | Pointer to **[]string** | A list of up to 8 URLs of images for this product, meant to be displayable to the customer. | [optional] 
**Livemode** | Pointer to **bool** | Has the value &#x60;true&#x60; if the object exists in live mode or the value &#x60;false&#x60; if the object exists in test mode. | [optional] 
**MarketingFeatures** | Pointer to [**[]StripeProductMarketingFeature**](StripeProductMarketingFeature.md) | A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://stripe.com/docs/payments/checkout/pricing-table). | [optional] 
**Metadata** | Pointer to **map[string]string** | Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. | [optional] 
**Name** | Pointer to **string** | The product&#39;s name, meant to be displayable to the customer. | [optional] 
**Object** | Pointer to **string** | String representing the object&#39;s type. Always has the value &#x60;product&#x60;. | [optional] 
**PackageDimensions** | Pointer to [**StripeProductPackageDimensions**](StripeProductPackageDimensions.md) | The dimensions of this product for shipping purposes. | [optional] 
**Shippable** | Pointer to **bool** | Whether this product is shipped (i.e., physical goods). | [optional] 
**StatementDescriptor** | Pointer to **string** | Extra information about a product which will appear on your customer&#39;s credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used. | [optional] 
**TaxCode** | Pointer to **map[string]interface{}** | A [tax code](https://stripe.com/docs/tax/tax-categories) ID. | [optional] 
**UnitLabel** | Pointer to **string** | A label that represents units of this product. When set, this will be included in customers&#39; receipts, invoices, Checkout, and the customer portal. | [optional] 
**Updated** | Pointer to **int32** | Time at which the product was last updated. Measured in seconds since the Unix epoch. | [optional] 
**Url** | Pointer to **string** | A URL of a publicly-accessible webpage for this product. | [optional] 

## Methods

### NewStripeProduct

`func NewStripeProduct() *StripeProduct`

NewStripeProduct instantiates a new StripeProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeProductWithDefaults

`func NewStripeProductWithDefaults() *StripeProduct`

NewStripeProductWithDefaults instantiates a new StripeProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *StripeProduct) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StripeProduct) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StripeProduct) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StripeProduct) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreated

`func (o *StripeProduct) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StripeProduct) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StripeProduct) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StripeProduct) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *StripeProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StripeProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StripeProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StripeProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *StripeProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *StripeProduct) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *StripeProduct) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *StripeProduct) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *StripeProduct) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetLivemode

`func (o *StripeProduct) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *StripeProduct) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *StripeProduct) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *StripeProduct) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetMarketingFeatures

`func (o *StripeProduct) GetMarketingFeatures() []StripeProductMarketingFeature`

GetMarketingFeatures returns the MarketingFeatures field if non-nil, zero value otherwise.

### GetMarketingFeaturesOk

`func (o *StripeProduct) GetMarketingFeaturesOk() (*[]StripeProductMarketingFeature, bool)`

GetMarketingFeaturesOk returns a tuple with the MarketingFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingFeatures

`func (o *StripeProduct) SetMarketingFeatures(v []StripeProductMarketingFeature)`

SetMarketingFeatures sets MarketingFeatures field to given value.

### HasMarketingFeatures

`func (o *StripeProduct) HasMarketingFeatures() bool`

HasMarketingFeatures returns a boolean if a field has been set.

### GetMetadata

`func (o *StripeProduct) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StripeProduct) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StripeProduct) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StripeProduct) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *StripeProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StripeProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StripeProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StripeProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *StripeProduct) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StripeProduct) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StripeProduct) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *StripeProduct) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPackageDimensions

`func (o *StripeProduct) GetPackageDimensions() StripeProductPackageDimensions`

GetPackageDimensions returns the PackageDimensions field if non-nil, zero value otherwise.

### GetPackageDimensionsOk

`func (o *StripeProduct) GetPackageDimensionsOk() (*StripeProductPackageDimensions, bool)`

GetPackageDimensionsOk returns a tuple with the PackageDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDimensions

`func (o *StripeProduct) SetPackageDimensions(v StripeProductPackageDimensions)`

SetPackageDimensions sets PackageDimensions field to given value.

### HasPackageDimensions

`func (o *StripeProduct) HasPackageDimensions() bool`

HasPackageDimensions returns a boolean if a field has been set.

### GetShippable

`func (o *StripeProduct) GetShippable() bool`

GetShippable returns the Shippable field if non-nil, zero value otherwise.

### GetShippableOk

`func (o *StripeProduct) GetShippableOk() (*bool, bool)`

GetShippableOk returns a tuple with the Shippable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippable

`func (o *StripeProduct) SetShippable(v bool)`

SetShippable sets Shippable field to given value.

### HasShippable

`func (o *StripeProduct) HasShippable() bool`

HasShippable returns a boolean if a field has been set.

### GetStatementDescriptor

`func (o *StripeProduct) GetStatementDescriptor() string`

GetStatementDescriptor returns the StatementDescriptor field if non-nil, zero value otherwise.

### GetStatementDescriptorOk

`func (o *StripeProduct) GetStatementDescriptorOk() (*string, bool)`

GetStatementDescriptorOk returns a tuple with the StatementDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDescriptor

`func (o *StripeProduct) SetStatementDescriptor(v string)`

SetStatementDescriptor sets StatementDescriptor field to given value.

### HasStatementDescriptor

`func (o *StripeProduct) HasStatementDescriptor() bool`

HasStatementDescriptor returns a boolean if a field has been set.

### GetTaxCode

`func (o *StripeProduct) GetTaxCode() map[string]interface{}`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *StripeProduct) GetTaxCodeOk() (*map[string]interface{}, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *StripeProduct) SetTaxCode(v map[string]interface{})`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *StripeProduct) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetUnitLabel

`func (o *StripeProduct) GetUnitLabel() string`

GetUnitLabel returns the UnitLabel field if non-nil, zero value otherwise.

### GetUnitLabelOk

`func (o *StripeProduct) GetUnitLabelOk() (*string, bool)`

GetUnitLabelOk returns a tuple with the UnitLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitLabel

`func (o *StripeProduct) SetUnitLabel(v string)`

SetUnitLabel sets UnitLabel field to given value.

### HasUnitLabel

`func (o *StripeProduct) HasUnitLabel() bool`

HasUnitLabel returns a boolean if a field has been set.

### GetUpdated

`func (o *StripeProduct) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *StripeProduct) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *StripeProduct) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *StripeProduct) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUrl

`func (o *StripeProduct) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StripeProduct) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StripeProduct) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StripeProduct) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


