# GcpMarketplaceProductServiceConfigBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to **[]string** | The list of metrics that are available for billing for the product. In format of \&quot;product-name.endpoints.gcp-project-id.cloud.goog/plan_name_metric_name\&quot; | [optional] 

## Methods

### NewGcpMarketplaceProductServiceConfigBilling

`func NewGcpMarketplaceProductServiceConfigBilling() *GcpMarketplaceProductServiceConfigBilling`

NewGcpMarketplaceProductServiceConfigBilling instantiates a new GcpMarketplaceProductServiceConfigBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductServiceConfigBillingWithDefaults

`func NewGcpMarketplaceProductServiceConfigBillingWithDefaults() *GcpMarketplaceProductServiceConfigBilling`

NewGcpMarketplaceProductServiceConfigBillingWithDefaults instantiates a new GcpMarketplaceProductServiceConfigBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *GcpMarketplaceProductServiceConfigBilling) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GcpMarketplaceProductServiceConfigBilling) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GcpMarketplaceProductServiceConfigBilling) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GcpMarketplaceProductServiceConfigBilling) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


