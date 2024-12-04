# UniqueCountAggregationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewItems** | Pointer to **map[string]interface{}** | Unique property values of current hour that are new of today. Leave the value as interface{} to save space. | [optional] 

## Methods

### NewUniqueCountAggregationResult

`func NewUniqueCountAggregationResult() *UniqueCountAggregationResult`

NewUniqueCountAggregationResult instantiates a new UniqueCountAggregationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueCountAggregationResultWithDefaults

`func NewUniqueCountAggregationResultWithDefaults() *UniqueCountAggregationResult`

NewUniqueCountAggregationResultWithDefaults instantiates a new UniqueCountAggregationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewItems

`func (o *UniqueCountAggregationResult) GetNewItems() map[string]interface{}`

GetNewItems returns the NewItems field if non-nil, zero value otherwise.

### GetNewItemsOk

`func (o *UniqueCountAggregationResult) GetNewItemsOk() (*map[string]interface{}, bool)`

GetNewItemsOk returns a tuple with the NewItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewItems

`func (o *UniqueCountAggregationResult) SetNewItems(v map[string]interface{})`

SetNewItems sets NewItems field to given value.

### HasNewItems

`func (o *UniqueCountAggregationResult) HasNewItems() bool`

HasNewItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


