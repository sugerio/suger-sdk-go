# SalesforceSyncFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to [**SalesforceSyncFilterOperator**](SalesforceSyncFilterOperator.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSalesforceSyncFilter

`func NewSalesforceSyncFilter() *SalesforceSyncFilter`

NewSalesforceSyncFilter instantiates a new SalesforceSyncFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceSyncFilterWithDefaults

`func NewSalesforceSyncFilterWithDefaults() *SalesforceSyncFilter`

NewSalesforceSyncFilterWithDefaults instantiates a new SalesforceSyncFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *SalesforceSyncFilter) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *SalesforceSyncFilter) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *SalesforceSyncFilter) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *SalesforceSyncFilter) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOperator

`func (o *SalesforceSyncFilter) GetOperator() SalesforceSyncFilterOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SalesforceSyncFilter) GetOperatorOk() (*SalesforceSyncFilterOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SalesforceSyncFilter) SetOperator(v SalesforceSyncFilterOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *SalesforceSyncFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *SalesforceSyncFilter) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SalesforceSyncFilter) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SalesforceSyncFilter) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SalesforceSyncFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


