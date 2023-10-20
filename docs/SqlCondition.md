# SqlCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Column** | Pointer to **string** | The column name. | [optional] 
**Operator** | Pointer to [**SqlOperator**](SqlOperator.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value. | [optional] 
**ValueType** | Pointer to [**SqlValueType**](SqlValueType.md) |  | [optional] 

## Methods

### NewSqlCondition

`func NewSqlCondition() *SqlCondition`

NewSqlCondition instantiates a new SqlCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlConditionWithDefaults

`func NewSqlConditionWithDefaults() *SqlCondition`

NewSqlConditionWithDefaults instantiates a new SqlCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumn

`func (o *SqlCondition) GetColumn() string`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *SqlCondition) GetColumnOk() (*string, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *SqlCondition) SetColumn(v string)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *SqlCondition) HasColumn() bool`

HasColumn returns a boolean if a field has been set.

### GetOperator

`func (o *SqlCondition) GetOperator() SqlOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SqlCondition) GetOperatorOk() (*SqlOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SqlCondition) SetOperator(v SqlOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *SqlCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *SqlCondition) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SqlCondition) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SqlCondition) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SqlCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *SqlCondition) GetValueType() SqlValueType`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *SqlCondition) GetValueTypeOk() (*SqlValueType, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *SqlCondition) SetValueType(v SqlValueType)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *SqlCondition) HasValueType() bool`

HasValueType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


