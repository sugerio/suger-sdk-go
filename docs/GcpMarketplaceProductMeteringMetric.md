# GcpMarketplaceProductMeteringMetric

## Properties

 Name                       | Type                                             | Description                                                                                                                                              | Notes      
----------------------------|--------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **Description**            | Pointer to **string**                            | Description: A detailed description of the metric, which can be used in documentation.                                                                   | [optional] 
 **DisplayName**            | Pointer to **string**                            |                                                                                                                                                          | [optional] 
 **DisplayUnit**            | Pointer to **string**                            | such as \&quot;min\&quot;                                                                                                                                | [optional] 
 **DisplayUnitDescription** | Pointer to **string**                            | such as \&quot;minute\&quot;                                                                                                                             | [optional] 
 **Id**                     | Pointer to **string**                            | The usage metering metric/dimension key It is in format of \&quot;{plan_id}_{usage_dimension_key}\&quot;. For example, \&quot;basic_plan_storage\&quot;. | [optional] 
 **MetricKind**             | Pointer to **string**                            | such as \&quot;DELTA\&quot;                                                                                                                              | [optional] 
 **Name**                   | Pointer to **string**                            | Name: The resource name of the metric descriptor, in format of \&quot;{productServiceName}/{plan_id}_{usage_dimension_key}\&quot;                        | [optional] 
 **PriceTiers**             | Pointer to [**[]GcpPriceTier**](GcpPriceTier.md) | Price info of this usage metering metric. Only applicable for the default offer (plan) and private offer.                                                | [optional] 
 **ReportingUnit**          | Pointer to **string**                            | such as \&quot;min\&quot;                                                                                                                                | [optional] 
 **SkuId**                  | Pointer to **string**                            | The SKU ID of this usage metering metric. Applicable only in Private Offer.                                                                              | [optional] 
 **Unit**                   | Pointer to **string**                            | such as \&quot;min\&quot;                                                                                                                                | [optional] 
 **ValueType**              | Pointer to [**ValueType**](ValueType.md)         | such as \&quot;INT64\&quot;                                                                                                                              | [optional] 

## Methods

### NewGcpMarketplaceProductMeteringMetric

`func NewGcpMarketplaceProductMeteringMetric() *GcpMarketplaceProductMeteringMetric`

NewGcpMarketplaceProductMeteringMetric instantiates a new GcpMarketplaceProductMeteringMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceProductMeteringMetricWithDefaults

`func NewGcpMarketplaceProductMeteringMetricWithDefaults() *GcpMarketplaceProductMeteringMetric`

NewGcpMarketplaceProductMeteringMetricWithDefaults instantiates a new GcpMarketplaceProductMeteringMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GcpMarketplaceProductMeteringMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GcpMarketplaceProductMeteringMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GcpMarketplaceProductMeteringMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GcpMarketplaceProductMeteringMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *GcpMarketplaceProductMeteringMetric) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GcpMarketplaceProductMeteringMetric) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GcpMarketplaceProductMeteringMetric) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GcpMarketplaceProductMeteringMetric) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayUnit

`func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnit() string`

GetDisplayUnit returns the DisplayUnit field if non-nil, zero value otherwise.

### GetDisplayUnitOk

`func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnitOk() (*string, bool)`

GetDisplayUnitOk returns a tuple with the DisplayUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUnit

`func (o *GcpMarketplaceProductMeteringMetric) SetDisplayUnit(v string)`

SetDisplayUnit sets DisplayUnit field to given value.

### HasDisplayUnit

`func (o *GcpMarketplaceProductMeteringMetric) HasDisplayUnit() bool`

HasDisplayUnit returns a boolean if a field has been set.

### GetDisplayUnitDescription

`func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnitDescription() string`

GetDisplayUnitDescription returns the DisplayUnitDescription field if non-nil, zero value otherwise.

### GetDisplayUnitDescriptionOk

`func (o *GcpMarketplaceProductMeteringMetric) GetDisplayUnitDescriptionOk() (*string, bool)`

GetDisplayUnitDescriptionOk returns a tuple with the DisplayUnitDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayUnitDescription

`func (o *GcpMarketplaceProductMeteringMetric) SetDisplayUnitDescription(v string)`

SetDisplayUnitDescription sets DisplayUnitDescription field to given value.

### HasDisplayUnitDescription

`func (o *GcpMarketplaceProductMeteringMetric) HasDisplayUnitDescription() bool`

HasDisplayUnitDescription returns a boolean if a field has been set.

### GetId

`func (o *GcpMarketplaceProductMeteringMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcpMarketplaceProductMeteringMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcpMarketplaceProductMeteringMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GcpMarketplaceProductMeteringMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetricKind

`func (o *GcpMarketplaceProductMeteringMetric) GetMetricKind() string`

GetMetricKind returns the MetricKind field if non-nil, zero value otherwise.

### GetMetricKindOk

`func (o *GcpMarketplaceProductMeteringMetric) GetMetricKindOk() (*string, bool)`

GetMetricKindOk returns a tuple with the MetricKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKind

`func (o *GcpMarketplaceProductMeteringMetric) SetMetricKind(v string)`

SetMetricKind sets MetricKind field to given value.

### HasMetricKind

`func (o *GcpMarketplaceProductMeteringMetric) HasMetricKind() bool`

HasMetricKind returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceProductMeteringMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceProductMeteringMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceProductMeteringMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceProductMeteringMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriceTiers

`func (o *GcpMarketplaceProductMeteringMetric) GetPriceTiers() []GcpPriceTier`

GetPriceTiers returns the PriceTiers field if non-nil, zero value otherwise.

### GetPriceTiersOk

`func (o *GcpMarketplaceProductMeteringMetric) GetPriceTiersOk() (*[]GcpPriceTier, bool)`

GetPriceTiersOk returns a tuple with the PriceTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTiers

`func (o *GcpMarketplaceProductMeteringMetric) SetPriceTiers(v []GcpPriceTier)`

SetPriceTiers sets PriceTiers field to given value.

### HasPriceTiers

`func (o *GcpMarketplaceProductMeteringMetric) HasPriceTiers() bool`

HasPriceTiers returns a boolean if a field has been set.

### GetReportingUnit

`func (o *GcpMarketplaceProductMeteringMetric) GetReportingUnit() string`

GetReportingUnit returns the ReportingUnit field if non-nil, zero value otherwise.

### GetReportingUnitOk

`func (o *GcpMarketplaceProductMeteringMetric) GetReportingUnitOk() (*string, bool)`

GetReportingUnitOk returns a tuple with the ReportingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingUnit

`func (o *GcpMarketplaceProductMeteringMetric) SetReportingUnit(v string)`

SetReportingUnit sets ReportingUnit field to given value.

### HasReportingUnit

`func (o *GcpMarketplaceProductMeteringMetric) HasReportingUnit() bool`

HasReportingUnit returns a boolean if a field has been set.

### GetSkuId

`func (o *GcpMarketplaceProductMeteringMetric) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *GcpMarketplaceProductMeteringMetric) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *GcpMarketplaceProductMeteringMetric) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *GcpMarketplaceProductMeteringMetric) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### GetUnit

`func (o *GcpMarketplaceProductMeteringMetric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *GcpMarketplaceProductMeteringMetric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *GcpMarketplaceProductMeteringMetric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *GcpMarketplaceProductMeteringMetric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValueType

`func (o *GcpMarketplaceProductMeteringMetric) GetValueType() ValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *GcpMarketplaceProductMeteringMetric) GetValueTypeOk() (*ValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *GcpMarketplaceProductMeteringMetric) SetValueType(v ValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *GcpMarketplaceProductMeteringMetric) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


