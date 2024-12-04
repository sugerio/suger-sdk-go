# AggregatedMeteringUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount calculated by billable dimension&#39;s price model, this is only used for report billable usage records to marketplace. | [optional] 
**BillableMetricAggregationType** | Pointer to [**BillableMetricAggregationType**](BillableMetricAggregationType.md) |  | [optional] 
**BillableMetricInfo** | Pointer to [**BillableMetricInfo**](BillableMetricInfo.md) |  | [optional] 
**GroupBysExpression** | Pointer to **string** | GroupBysExpression is the string expression of array of group bys. | [optional] 
**Key** | Pointer to **string** | Key is the unique identifier of a billable metric. | [optional] 
**Name** | Pointer to **string** | Name is the name of a billable metric. Optional, it is only for display purpose. | [optional] 
**Quantity** | Pointer to **float32** | Value is the value of a billable metric. | [optional] 
**UniqueCountAggregationResult** | Pointer to [**UniqueCountAggregationResult**](UniqueCountAggregationResult.md) | Unique count metric aggregate result. | [optional] 

## Methods

### NewAggregatedMeteringUsageRecord

`func NewAggregatedMeteringUsageRecord() *AggregatedMeteringUsageRecord`

NewAggregatedMeteringUsageRecord instantiates a new AggregatedMeteringUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatedMeteringUsageRecordWithDefaults

`func NewAggregatedMeteringUsageRecordWithDefaults() *AggregatedMeteringUsageRecord`

NewAggregatedMeteringUsageRecordWithDefaults instantiates a new AggregatedMeteringUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AggregatedMeteringUsageRecord) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AggregatedMeteringUsageRecord) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AggregatedMeteringUsageRecord) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AggregatedMeteringUsageRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillableMetricAggregationType

`func (o *AggregatedMeteringUsageRecord) GetBillableMetricAggregationType() BillableMetricAggregationType`

GetBillableMetricAggregationType returns the BillableMetricAggregationType field if non-nil, zero value otherwise.

### GetBillableMetricAggregationTypeOk

`func (o *AggregatedMeteringUsageRecord) GetBillableMetricAggregationTypeOk() (*BillableMetricAggregationType, bool)`

GetBillableMetricAggregationTypeOk returns a tuple with the BillableMetricAggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricAggregationType

`func (o *AggregatedMeteringUsageRecord) SetBillableMetricAggregationType(v BillableMetricAggregationType)`

SetBillableMetricAggregationType sets BillableMetricAggregationType field to given value.

### HasBillableMetricAggregationType

`func (o *AggregatedMeteringUsageRecord) HasBillableMetricAggregationType() bool`

HasBillableMetricAggregationType returns a boolean if a field has been set.

### GetBillableMetricInfo

`func (o *AggregatedMeteringUsageRecord) GetBillableMetricInfo() BillableMetricInfo`

GetBillableMetricInfo returns the BillableMetricInfo field if non-nil, zero value otherwise.

### GetBillableMetricInfoOk

`func (o *AggregatedMeteringUsageRecord) GetBillableMetricInfoOk() (*BillableMetricInfo, bool)`

GetBillableMetricInfoOk returns a tuple with the BillableMetricInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricInfo

`func (o *AggregatedMeteringUsageRecord) SetBillableMetricInfo(v BillableMetricInfo)`

SetBillableMetricInfo sets BillableMetricInfo field to given value.

### HasBillableMetricInfo

`func (o *AggregatedMeteringUsageRecord) HasBillableMetricInfo() bool`

HasBillableMetricInfo returns a boolean if a field has been set.

### GetGroupBysExpression

`func (o *AggregatedMeteringUsageRecord) GetGroupBysExpression() string`

GetGroupBysExpression returns the GroupBysExpression field if non-nil, zero value otherwise.

### GetGroupBysExpressionOk

`func (o *AggregatedMeteringUsageRecord) GetGroupBysExpressionOk() (*string, bool)`

GetGroupBysExpressionOk returns a tuple with the GroupBysExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBysExpression

`func (o *AggregatedMeteringUsageRecord) SetGroupBysExpression(v string)`

SetGroupBysExpression sets GroupBysExpression field to given value.

### HasGroupBysExpression

`func (o *AggregatedMeteringUsageRecord) HasGroupBysExpression() bool`

HasGroupBysExpression returns a boolean if a field has been set.

### GetKey

`func (o *AggregatedMeteringUsageRecord) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AggregatedMeteringUsageRecord) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AggregatedMeteringUsageRecord) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AggregatedMeteringUsageRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *AggregatedMeteringUsageRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AggregatedMeteringUsageRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AggregatedMeteringUsageRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AggregatedMeteringUsageRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *AggregatedMeteringUsageRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AggregatedMeteringUsageRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AggregatedMeteringUsageRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AggregatedMeteringUsageRecord) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUniqueCountAggregationResult

`func (o *AggregatedMeteringUsageRecord) GetUniqueCountAggregationResult() UniqueCountAggregationResult`

GetUniqueCountAggregationResult returns the UniqueCountAggregationResult field if non-nil, zero value otherwise.

### GetUniqueCountAggregationResultOk

`func (o *AggregatedMeteringUsageRecord) GetUniqueCountAggregationResultOk() (*UniqueCountAggregationResult, bool)`

GetUniqueCountAggregationResultOk returns a tuple with the UniqueCountAggregationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCountAggregationResult

`func (o *AggregatedMeteringUsageRecord) SetUniqueCountAggregationResult(v UniqueCountAggregationResult)`

SetUniqueCountAggregationResult sets UniqueCountAggregationResult field to given value.

### HasUniqueCountAggregationResult

`func (o *AggregatedMeteringUsageRecord) HasUniqueCountAggregationResult() bool`

HasUniqueCountAggregationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


