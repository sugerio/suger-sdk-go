# GcpMarketplaceMeteringMetricValueSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricName** | Pointer to **string** | MetricName: The metric name defined in the service configuration. | [optional] 
**MetricValues** | Pointer to [**[]GcpMarketplaceMeteringMetricValue**](GcpMarketplaceMeteringMetricValue.md) | MetricValues: The values in this metric. | [optional] 

## Methods

### NewGcpMarketplaceMeteringMetricValueSet

`func NewGcpMarketplaceMeteringMetricValueSet() *GcpMarketplaceMeteringMetricValueSet`

NewGcpMarketplaceMeteringMetricValueSet instantiates a new GcpMarketplaceMeteringMetricValueSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceMeteringMetricValueSetWithDefaults

`func NewGcpMarketplaceMeteringMetricValueSetWithDefaults() *GcpMarketplaceMeteringMetricValueSet`

NewGcpMarketplaceMeteringMetricValueSetWithDefaults instantiates a new GcpMarketplaceMeteringMetricValueSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricName

`func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *GcpMarketplaceMeteringMetricValueSet) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *GcpMarketplaceMeteringMetricValueSet) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricValues

`func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricValues() []GcpMarketplaceMeteringMetricValue`

GetMetricValues returns the MetricValues field if non-nil, zero value otherwise.

### GetMetricValuesOk

`func (o *GcpMarketplaceMeteringMetricValueSet) GetMetricValuesOk() (*[]GcpMarketplaceMeteringMetricValue, bool)`

GetMetricValuesOk returns a tuple with the MetricValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValues

`func (o *GcpMarketplaceMeteringMetricValueSet) SetMetricValues(v []GcpMarketplaceMeteringMetricValue)`

SetMetricValues sets MetricValues field to given value.

### HasMetricValues

`func (o *GcpMarketplaceMeteringMetricValueSet) HasMetricValues() bool`

HasMetricValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


