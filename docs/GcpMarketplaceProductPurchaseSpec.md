# GcpMarketplaceProductPurchaseSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | Pointer to [**[]GcpMarketplaceProductFeature**](GcpMarketplaceProductFeature.md) |  | [optional] 
**Metrics** | Pointer to [**[]GcpMarketplaceProductMeteringMetric**](GcpMarketplaceProductMeteringMetric.md) | GCP Marketplace Product Usage Metering Dimension/Metric | [optional] 
**PurchaseOptionSpecs** | Pointer to [**[]GcpMarketplaceProductPurchaseOptionSpec**](GcpMarketplaceProductPurchaseOptionSpec.md) | GCP Marketplace Product Pricing Plans | [optional] 

## Methods

### NewGcpMarketplaceProductPurchaseSpec

`func NewGcpMarketplaceProductPurchaseSpec() *GcpMarketplaceProductPurchaseSpec`

NewGcpMarketplaceProductPurchaseSpec instantiates a new GcpMarketplaceProductPurchaseSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductPurchaseSpecWithDefaults

`func NewGcpMarketplaceProductPurchaseSpecWithDefaults() *GcpMarketplaceProductPurchaseSpec`

NewGcpMarketplaceProductPurchaseSpecWithDefaults instantiates a new GcpMarketplaceProductPurchaseSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *GcpMarketplaceProductPurchaseSpec) GetFeatures() []GcpMarketplaceProductFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GcpMarketplaceProductPurchaseSpec) GetFeaturesOk() (*[]GcpMarketplaceProductFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GcpMarketplaceProductPurchaseSpec) SetFeatures(v []GcpMarketplaceProductFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GcpMarketplaceProductPurchaseSpec) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetMetrics

`func (o *GcpMarketplaceProductPurchaseSpec) GetMetrics() []GcpMarketplaceProductMeteringMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GcpMarketplaceProductPurchaseSpec) GetMetricsOk() (*[]GcpMarketplaceProductMeteringMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GcpMarketplaceProductPurchaseSpec) SetMetrics(v []GcpMarketplaceProductMeteringMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GcpMarketplaceProductPurchaseSpec) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPurchaseOptionSpecs

`func (o *GcpMarketplaceProductPurchaseSpec) GetPurchaseOptionSpecs() []GcpMarketplaceProductPurchaseOptionSpec`

GetPurchaseOptionSpecs returns the PurchaseOptionSpecs field if non-nil, zero value otherwise.

### GetPurchaseOptionSpecsOk

`func (o *GcpMarketplaceProductPurchaseSpec) GetPurchaseOptionSpecsOk() (*[]GcpMarketplaceProductPurchaseOptionSpec, bool)`

GetPurchaseOptionSpecsOk returns a tuple with the PurchaseOptionSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOptionSpecs

`func (o *GcpMarketplaceProductPurchaseSpec) SetPurchaseOptionSpecs(v []GcpMarketplaceProductPurchaseOptionSpec)`

SetPurchaseOptionSpecs sets PurchaseOptionSpecs field to given value.

### HasPurchaseOptionSpecs

`func (o *GcpMarketplaceProductPurchaseSpec) HasPurchaseOptionSpecs() bool`

HasPurchaseOptionSpecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


