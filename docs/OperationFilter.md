# OperationFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to [**[]OperationFilter**](OperationFilter.md) | Args for logical operators (and, or) | [optional] 
**Field** | Pointer to [**TemporalWorkflowAttr**](TemporalWorkflowAttr.md) | Field is temporal workflow filter field | [optional] 
**Operator** | Pointer to **string** | Operator: \&quot;&#x3D;\&quot;, \&quot;!&#x3D;\&quot;, \&quot;in\&quot;, \&quot;not_in\&quot;, \&quot;and\&quot;, \&quot;or\&quot; | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value for single value operators (&#x3D;, !&#x3D;, etc) | [optional] 
**Values** | Pointer to **[]map[string]interface{}** | Values for array operators (in, not_in) | [optional] 

## Methods

### NewOperationFilter

`func NewOperationFilter() *OperationFilter`

NewOperationFilter instantiates a new OperationFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationFilterWithDefaults

`func NewOperationFilterWithDefaults() *OperationFilter`

NewOperationFilterWithDefaults instantiates a new OperationFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *OperationFilter) GetArgs() []OperationFilter`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *OperationFilter) GetArgsOk() (*[]OperationFilter, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *OperationFilter) SetArgs(v []OperationFilter)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *OperationFilter) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetField

`func (o *OperationFilter) GetField() TemporalWorkflowAttr`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *OperationFilter) GetFieldOk() (*TemporalWorkflowAttr, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *OperationFilter) SetField(v TemporalWorkflowAttr)`

SetField sets Field field to given value.

### HasField

`func (o *OperationFilter) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOperator

`func (o *OperationFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *OperationFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *OperationFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *OperationFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *OperationFilter) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OperationFilter) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OperationFilter) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *OperationFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *OperationFilter) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *OperationFilter) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *OperationFilter) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *OperationFilter) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


