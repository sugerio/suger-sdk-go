# AwsMarketplaceCatalogLegalTermDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AwsMarketplaceCatalogLegalTermDocumentType**](AwsMarketplaceCatalogLegalTermDocumentType.md) |  | [optional] 
**Url** | Pointer to **string** | A URL to the legal document for buyers to read. Required when Type is one of the following [CustomEula, CustomDsa]. | [optional] 
**Version** | Pointer to **string** | Version of standard contracts provided by AWS Marketplace. Required when Type is one of the following [StandardEula, StandardDsa]. The version of StandardEula is \&quot;2022-07-14\&quot;. The version of StandardDsa is \&quot;2019-12-12\&quot;. | [optional] 

## Methods

### NewAwsMarketplaceCatalogLegalTermDocument

`func NewAwsMarketplaceCatalogLegalTermDocument() *AwsMarketplaceCatalogLegalTermDocument`

NewAwsMarketplaceCatalogLegalTermDocument instantiates a new AwsMarketplaceCatalogLegalTermDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCatalogLegalTermDocumentWithDefaults

`func NewAwsMarketplaceCatalogLegalTermDocumentWithDefaults() *AwsMarketplaceCatalogLegalTermDocument`

NewAwsMarketplaceCatalogLegalTermDocumentWithDefaults instantiates a new AwsMarketplaceCatalogLegalTermDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwsMarketplaceCatalogLegalTermDocument) GetType() AwsMarketplaceCatalogLegalTermDocumentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsMarketplaceCatalogLegalTermDocument) GetTypeOk() (*AwsMarketplaceCatalogLegalTermDocumentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsMarketplaceCatalogLegalTermDocument) SetType(v AwsMarketplaceCatalogLegalTermDocumentType)`

SetType sets Type field to given value.

### HasType

`func (o *AwsMarketplaceCatalogLegalTermDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AwsMarketplaceCatalogLegalTermDocument) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AwsMarketplaceCatalogLegalTermDocument) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AwsMarketplaceCatalogLegalTermDocument) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AwsMarketplaceCatalogLegalTermDocument) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersion

`func (o *AwsMarketplaceCatalogLegalTermDocument) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AwsMarketplaceCatalogLegalTermDocument) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AwsMarketplaceCatalogLegalTermDocument) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AwsMarketplaceCatalogLegalTermDocument) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


