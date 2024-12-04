# BillableMetricFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to [**BillableMetricFilterOperation**](BillableMetricFilterOperation.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value of the filter. The type of the value depends on the valueType. | [optional] 
**ValueType** | Pointer to [**BillableMetricFilterValueType**](BillableMetricFilterValueType.md) |  | [optional] 

## Methods

### NewBillableMetricFilter

`func NewBillableMetricFilter() *BillableMetricFilter`

NewBillableMetricFilter instantiates a new BillableMetricFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricFilterWithDefaults

`func NewBillableMetricFilterWithDefaults() *BillableMetricFilter`

NewBillableMetricFilterWithDefaults instantiates a new BillableMetricFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BillableMetricFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableMetricFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableMetricFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillableMetricFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *BillableMetricFilter) GetOperation() BillableMetricFilterOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BillableMetricFilter) GetOperationOk() (*BillableMetricFilterOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BillableMetricFilter) SetOperation(v BillableMetricFilterOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *BillableMetricFilter) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *BillableMetricFilter) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BillableMetricFilter) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BillableMetricFilter) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BillableMetricFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *BillableMetricFilter) GetValueType() BillableMetricFilterValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *BillableMetricFilter) GetValueTypeOk() (*BillableMetricFilterValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *BillableMetricFilter) SetValueType(v BillableMetricFilterValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *BillableMetricFilter) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


