# GcpMarketplaceProductServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Billing** | Pointer to [**GcpMarketplaceProductServiceConfigBilling**](GcpMarketplaceProductServiceConfigBilling.md) |  | [optional] 
**Metrics** | Pointer to [**[]GcpMarketplaceProductMeteringMetric**](GcpMarketplaceProductMeteringMetric.md) |  | [optional] 
**Name** | Pointer to **string** | in format of \&quot;product-name.endpoints.gcp-project-id.cloud.goog\&quot; | [optional] 
**ProducerProjectId** | Pointer to **string** | The GCP project ID of the producer. | [optional] 
**Title** | Pointer to **string** | The title of the product listing. | [optional] 

## Methods

### NewGcpMarketplaceProductServiceConfig

`func NewGcpMarketplaceProductServiceConfig() *GcpMarketplaceProductServiceConfig`

NewGcpMarketplaceProductServiceConfig instantiates a new GcpMarketplaceProductServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductServiceConfigWithDefaults

`func NewGcpMarketplaceProductServiceConfigWithDefaults() *GcpMarketplaceProductServiceConfig`

NewGcpMarketplaceProductServiceConfigWithDefaults instantiates a new GcpMarketplaceProductServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilling

`func (o *GcpMarketplaceProductServiceConfig) GetBilling() GcpMarketplaceProductServiceConfigBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *GcpMarketplaceProductServiceConfig) GetBillingOk() (*GcpMarketplaceProductServiceConfigBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *GcpMarketplaceProductServiceConfig) SetBilling(v GcpMarketplaceProductServiceConfigBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *GcpMarketplaceProductServiceConfig) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetMetrics

`func (o *GcpMarketplaceProductServiceConfig) GetMetrics() []GcpMarketplaceProductMeteringMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GcpMarketplaceProductServiceConfig) GetMetricsOk() (*[]GcpMarketplaceProductMeteringMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GcpMarketplaceProductServiceConfig) SetMetrics(v []GcpMarketplaceProductMeteringMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GcpMarketplaceProductServiceConfig) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceProductServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceProductServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceProductServiceConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceProductServiceConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProducerProjectId

`func (o *GcpMarketplaceProductServiceConfig) GetProducerProjectId() string`

GetProducerProjectId returns the ProducerProjectId field if non-nil, zero value otherwise.

### GetProducerProjectIdOk

`func (o *GcpMarketplaceProductServiceConfig) GetProducerProjectIdOk() (*string, bool)`

GetProducerProjectIdOk returns a tuple with the ProducerProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerProjectId

`func (o *GcpMarketplaceProductServiceConfig) SetProducerProjectId(v string)`

SetProducerProjectId sets ProducerProjectId field to given value.

### HasProducerProjectId

`func (o *GcpMarketplaceProductServiceConfig) HasProducerProjectId() bool`

HasProducerProjectId returns a boolean if a field has been set.

### GetTitle

`func (o *GcpMarketplaceProductServiceConfig) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GcpMarketplaceProductServiceConfig) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GcpMarketplaceProductServiceConfig) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GcpMarketplaceProductServiceConfig) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


