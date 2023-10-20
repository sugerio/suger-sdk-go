# GcpMarketplaceProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**DerivedDiscoveryState** | Pointer to [**GcpMarketplaceProductDerivedDiscoveryState**](GcpMarketplaceProductDerivedDiscoveryState.md) |  | [optional] 
**Id** | Pointer to **string** | Nullable, GCP Marketplace Product UUID | [optional] 
**LastPublishTime** | Pointer to **time.Time** |  | [optional] 
**ListingSpec** | Pointer to [**GcpMarketplaceProductListingSpec**](GcpMarketplaceProductListingSpec.md) |  | [optional] 
**Marketplace** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | In format of \&quot;projects/{project-number}/listings/{product-name}.endpoints.{provider-id}.cloud.goog\&quot; | [optional] 
**RevisionCreateTime** | Pointer to **time.Time** |  | [optional] 
**RevisionId** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** | In format of \&quot;services/{product-name}.endpoints.{provider-id}.cloud.goog\&quot; | [optional] 
**ServiceConfig** | Pointer to [**GcpMarketplaceProductServiceConfig**](GcpMarketplaceProductServiceConfig.md) |  | [optional] 
**ValidationSummary** | Pointer to **map[string]interface{}** | TODO: add type | [optional] 

## Methods

### NewGcpMarketplaceProduct

`func NewGcpMarketplaceProduct() *GcpMarketplaceProduct`

NewGcpMarketplaceProduct instantiates a new GcpMarketplaceProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductWithDefaults

`func NewGcpMarketplaceProductWithDefaults() *GcpMarketplaceProduct`

NewGcpMarketplaceProductWithDefaults instantiates a new GcpMarketplaceProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *GcpMarketplaceProduct) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GcpMarketplaceProduct) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GcpMarketplaceProduct) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GcpMarketplaceProduct) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDerivedDiscoveryState

`func (o *GcpMarketplaceProduct) GetDerivedDiscoveryState() GcpMarketplaceProductDerivedDiscoveryState`

GetDerivedDiscoveryState returns the DerivedDiscoveryState field if non-nil, zero value otherwise.

### GetDerivedDiscoveryStateOk

`func (o *GcpMarketplaceProduct) GetDerivedDiscoveryStateOk() (*GcpMarketplaceProductDerivedDiscoveryState, bool)`

GetDerivedDiscoveryStateOk returns a tuple with the DerivedDiscoveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedDiscoveryState

`func (o *GcpMarketplaceProduct) SetDerivedDiscoveryState(v GcpMarketplaceProductDerivedDiscoveryState)`

SetDerivedDiscoveryState sets DerivedDiscoveryState field to given value.

### HasDerivedDiscoveryState

`func (o *GcpMarketplaceProduct) HasDerivedDiscoveryState() bool`

HasDerivedDiscoveryState returns a boolean if a field has been set.

### GetId

`func (o *GcpMarketplaceProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpMarketplaceProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpMarketplaceProduct) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GcpMarketplaceProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPublishTime

`func (o *GcpMarketplaceProduct) GetLastPublishTime() time.Time`

GetLastPublishTime returns the LastPublishTime field if non-nil, zero value otherwise.

### GetLastPublishTimeOk

`func (o *GcpMarketplaceProduct) GetLastPublishTimeOk() (*time.Time, bool)`

GetLastPublishTimeOk returns a tuple with the LastPublishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublishTime

`func (o *GcpMarketplaceProduct) SetLastPublishTime(v time.Time)`

SetLastPublishTime sets LastPublishTime field to given value.

### HasLastPublishTime

`func (o *GcpMarketplaceProduct) HasLastPublishTime() bool`

HasLastPublishTime returns a boolean if a field has been set.

### GetListingSpec

`func (o *GcpMarketplaceProduct) GetListingSpec() GcpMarketplaceProductListingSpec`

GetListingSpec returns the ListingSpec field if non-nil, zero value otherwise.

### GetListingSpecOk

`func (o *GcpMarketplaceProduct) GetListingSpecOk() (*GcpMarketplaceProductListingSpec, bool)`

GetListingSpecOk returns a tuple with the ListingSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingSpec

`func (o *GcpMarketplaceProduct) SetListingSpec(v GcpMarketplaceProductListingSpec)`

SetListingSpec sets ListingSpec field to given value.

### HasListingSpec

`func (o *GcpMarketplaceProduct) HasListingSpec() bool`

HasListingSpec returns a boolean if a field has been set.

### GetMarketplace

`func (o *GcpMarketplaceProduct) GetMarketplace() string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *GcpMarketplaceProduct) GetMarketplaceOk() (*string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *GcpMarketplaceProduct) SetMarketplace(v string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *GcpMarketplaceProduct) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceProduct) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceProduct) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevisionCreateTime

`func (o *GcpMarketplaceProduct) GetRevisionCreateTime() time.Time`

GetRevisionCreateTime returns the RevisionCreateTime field if non-nil, zero value otherwise.

### GetRevisionCreateTimeOk

`func (o *GcpMarketplaceProduct) GetRevisionCreateTimeOk() (*time.Time, bool)`

GetRevisionCreateTimeOk returns a tuple with the RevisionCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionCreateTime

`func (o *GcpMarketplaceProduct) SetRevisionCreateTime(v time.Time)`

SetRevisionCreateTime sets RevisionCreateTime field to given value.

### HasRevisionCreateTime

`func (o *GcpMarketplaceProduct) HasRevisionCreateTime() bool`

HasRevisionCreateTime returns a boolean if a field has been set.

### GetRevisionId

`func (o *GcpMarketplaceProduct) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GcpMarketplaceProduct) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GcpMarketplaceProduct) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GcpMarketplaceProduct) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetService

`func (o *GcpMarketplaceProduct) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GcpMarketplaceProduct) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GcpMarketplaceProduct) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GcpMarketplaceProduct) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceConfig

`func (o *GcpMarketplaceProduct) GetServiceConfig() GcpMarketplaceProductServiceConfig`

GetServiceConfig returns the ServiceConfig field if non-nil, zero value otherwise.

### GetServiceConfigOk

`func (o *GcpMarketplaceProduct) GetServiceConfigOk() (*GcpMarketplaceProductServiceConfig, bool)`

GetServiceConfigOk returns a tuple with the ServiceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfig

`func (o *GcpMarketplaceProduct) SetServiceConfig(v GcpMarketplaceProductServiceConfig)`

SetServiceConfig sets ServiceConfig field to given value.

### HasServiceConfig

`func (o *GcpMarketplaceProduct) HasServiceConfig() bool`

HasServiceConfig returns a boolean if a field has been set.

### GetValidationSummary

`func (o *GcpMarketplaceProduct) GetValidationSummary() map[string]interface{}`

GetValidationSummary returns the ValidationSummary field if non-nil, zero value otherwise.

### GetValidationSummaryOk

`func (o *GcpMarketplaceProduct) GetValidationSummaryOk() (*map[string]interface{}, bool)`

GetValidationSummaryOk returns a tuple with the ValidationSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationSummary

`func (o *GcpMarketplaceProduct) SetValidationSummary(v map[string]interface{})`

SetValidationSummary sets ValidationSummary field to given value.

### HasValidationSummary

`func (o *GcpMarketplaceProduct) HasValidationSummary() bool`

HasValidationSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


