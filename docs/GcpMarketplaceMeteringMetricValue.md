# GcpMarketplaceMeteringMetricValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoolValue** | Pointer to **bool** | BoolValue: A boolean value. | [optional] 
**DoubleValue** | Pointer to **float32** | DoubleValue: A double precision floating point value. | [optional] 
**EndTime** | Pointer to **string** | EndTime: The end of the time period over which this metric value&#39;s measurement applies. If not specified, google.api.servicecontrol.v1.Operation.end_time will be used. | [optional] 
**Int64Value** | Pointer to **string** | Int64Value: A signed 64-bit integer value. | [optional] 
**Labels** | Pointer to **map[string]string** | Labels: The labels describing the metric value. See comments on google.api.servicecontrol.v1.Operation.labels for the overriding relationship. Note that this map must not contain monitored resource labels. | [optional] 
**MoneyValue** | Pointer to [**GcpMarketplaceMeteringMoney**](GcpMarketplaceMeteringMoney.md) | MoneyValue: A money value. | [optional] 
**StartTime** | Pointer to **string** | StartTime: The start of the time period over which this metric value&#39;s measurement applies. The time period has different semantics for different metric types (cumulative, delta, and gauge). See the metric definition documentation in the service configuration for details. If not specified, google.api.servicecontrol.v1.Operation.start_time will be used. | [optional] 
**StringValue** | Pointer to **string** | StringValue: A text string value. | [optional] 

## Methods

### NewGcpMarketplaceMeteringMetricValue

`func NewGcpMarketplaceMeteringMetricValue() *GcpMarketplaceMeteringMetricValue`

NewGcpMarketplaceMeteringMetricValue instantiates a new GcpMarketplaceMeteringMetricValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceMeteringMetricValueWithDefaults

`func NewGcpMarketplaceMeteringMetricValueWithDefaults() *GcpMarketplaceMeteringMetricValue`

NewGcpMarketplaceMeteringMetricValueWithDefaults instantiates a new GcpMarketplaceMeteringMetricValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolValue

`func (o *GcpMarketplaceMeteringMetricValue) GetBoolValue() bool`

GetBoolValue returns the BoolValue field if non-nil, zero value otherwise.

### GetBoolValueOk

`func (o *GcpMarketplaceMeteringMetricValue) GetBoolValueOk() (*bool, bool)`

GetBoolValueOk returns a tuple with the BoolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValue

`func (o *GcpMarketplaceMeteringMetricValue) SetBoolValue(v bool)`

SetBoolValue sets BoolValue field to given value.

### HasBoolValue

`func (o *GcpMarketplaceMeteringMetricValue) HasBoolValue() bool`

HasBoolValue returns a boolean if a field has been set.

### GetDoubleValue

`func (o *GcpMarketplaceMeteringMetricValue) GetDoubleValue() float32`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *GcpMarketplaceMeteringMetricValue) GetDoubleValueOk() (*float32, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *GcpMarketplaceMeteringMetricValue) SetDoubleValue(v float32)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *GcpMarketplaceMeteringMetricValue) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### GetEndTime

`func (o *GcpMarketplaceMeteringMetricValue) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GcpMarketplaceMeteringMetricValue) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GcpMarketplaceMeteringMetricValue) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GcpMarketplaceMeteringMetricValue) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetInt64Value

`func (o *GcpMarketplaceMeteringMetricValue) GetInt64Value() string`

GetInt64Value returns the Int64Value field if non-nil, zero value otherwise.

### GetInt64ValueOk

`func (o *GcpMarketplaceMeteringMetricValue) GetInt64ValueOk() (*string, bool)`

GetInt64ValueOk returns a tuple with the Int64Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInt64Value

`func (o *GcpMarketplaceMeteringMetricValue) SetInt64Value(v string)`

SetInt64Value sets Int64Value field to given value.

### HasInt64Value

`func (o *GcpMarketplaceMeteringMetricValue) HasInt64Value() bool`

HasInt64Value returns a boolean if a field has been set.

### GetLabels

`func (o *GcpMarketplaceMeteringMetricValue) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GcpMarketplaceMeteringMetricValue) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GcpMarketplaceMeteringMetricValue) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GcpMarketplaceMeteringMetricValue) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMoneyValue

`func (o *GcpMarketplaceMeteringMetricValue) GetMoneyValue() GcpMarketplaceMeteringMoney`

GetMoneyValue returns the MoneyValue field if non-nil, zero value otherwise.

### GetMoneyValueOk

`func (o *GcpMarketplaceMeteringMetricValue) GetMoneyValueOk() (*GcpMarketplaceMeteringMoney, bool)`

GetMoneyValueOk returns a tuple with the MoneyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyValue

`func (o *GcpMarketplaceMeteringMetricValue) SetMoneyValue(v GcpMarketplaceMeteringMoney)`

SetMoneyValue sets MoneyValue field to given value.

### HasMoneyValue

`func (o *GcpMarketplaceMeteringMetricValue) HasMoneyValue() bool`

HasMoneyValue returns a boolean if a field has been set.

### GetStartTime

`func (o *GcpMarketplaceMeteringMetricValue) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GcpMarketplaceMeteringMetricValue) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GcpMarketplaceMeteringMetricValue) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GcpMarketplaceMeteringMetricValue) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStringValue

`func (o *GcpMarketplaceMeteringMetricValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *GcpMarketplaceMeteringMetricValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *GcpMarketplaceMeteringMetricValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *GcpMarketplaceMeteringMetricValue) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


