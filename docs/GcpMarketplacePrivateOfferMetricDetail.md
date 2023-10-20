# GcpMarketplacePrivateOfferMetricDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | such as \&quot;CPU\&quot; | [optional] 
**ParentCommerceService** | Pointer to **string** | in format of \&quot;projects/{projectNumber}/services/service-name.endpoints.gcp-project-id.cloud.goog\&quot; | [optional] 
**SkuId** | Pointer to **string** | such as \&quot;BC1B-6259-BF57\&quot; | [optional] 
**Tiers** | Pointer to [**[]GcpPriceTier**](GcpPriceTier.md) | Price tiers for this metric. | [optional] 
**UnitDescription** | Pointer to **string** | such as \&quot;minute\&quot; | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferMetricDetail

`func NewGcpMarketplacePrivateOfferMetricDetail() *GcpMarketplacePrivateOfferMetricDetail`

NewGcpMarketplacePrivateOfferMetricDetail instantiates a new GcpMarketplacePrivateOfferMetricDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferMetricDetailWithDefaults

`func NewGcpMarketplacePrivateOfferMetricDetailWithDefaults() *GcpMarketplacePrivateOfferMetricDetail`

NewGcpMarketplacePrivateOfferMetricDetailWithDefaults instantiates a new GcpMarketplacePrivateOfferMetricDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GcpMarketplacePrivateOfferMetricDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GcpMarketplacePrivateOfferMetricDetail) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetParentCommerceService

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetParentCommerceService() string`

GetParentCommerceService returns the ParentCommerceService field if non-nil, zero value otherwise.

### GetParentCommerceServiceOk

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetParentCommerceServiceOk() (*string, bool)`

GetParentCommerceServiceOk returns a tuple with the ParentCommerceService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommerceService

`func (o *GcpMarketplacePrivateOfferMetricDetail) SetParentCommerceService(v string)`

SetParentCommerceService sets ParentCommerceService field to given value.

### HasParentCommerceService

`func (o *GcpMarketplacePrivateOfferMetricDetail) HasParentCommerceService() bool`

HasParentCommerceService returns a boolean if a field has been set.

### GetSkuId

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *GcpMarketplacePrivateOfferMetricDetail) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *GcpMarketplacePrivateOfferMetricDetail) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetTiers

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetTiers() []GcpPriceTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetTiersOk() (*[]GcpPriceTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *GcpMarketplacePrivateOfferMetricDetail) SetTiers(v []GcpPriceTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *GcpMarketplacePrivateOfferMetricDetail) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetUnitDescription

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetUnitDescription() string`

GetUnitDescription returns the UnitDescription field if non-nil, zero value otherwise.

### GetUnitDescriptionOk

`func (o *GcpMarketplacePrivateOfferMetricDetail) GetUnitDescriptionOk() (*string, bool)`

GetUnitDescriptionOk returns a tuple with the UnitDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitDescription

`func (o *GcpMarketplacePrivateOfferMetricDetail) SetUnitDescription(v string)`

SetUnitDescription sets UnitDescription field to given value.

### HasUnitDescription

`func (o *GcpMarketplacePrivateOfferMetricDetail) HasUnitDescription() bool`

HasUnitDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


