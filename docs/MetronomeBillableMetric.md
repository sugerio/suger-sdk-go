# MetronomeBillableMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupBy** | Pointer to **[]string** | the fields to group by when aggregating usage events. | [optional] 
**Id** | Pointer to **string** | the uuid of the billable metric. | [optional] 
**Name** | Pointer to **string** | the name of the billable metric. | [optional] 

## Methods

### NewMetronomeBillableMetric

`func NewMetronomeBillableMetric() *MetronomeBillableMetric`

NewMetronomeBillableMetric instantiates a new MetronomeBillableMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetronomeBillableMetricWithDefaults

`func NewMetronomeBillableMetricWithDefaults() *MetronomeBillableMetric`

NewMetronomeBillableMetricWithDefaults instantiates a new MetronomeBillableMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupBy

`func (o *MetronomeBillableMetric) GetGroupBy() []string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetronomeBillableMetric) GetGroupByOk() (*[]string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetronomeBillableMetric) SetGroupBy(v []string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetronomeBillableMetric) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetId

`func (o *MetronomeBillableMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetronomeBillableMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetronomeBillableMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MetronomeBillableMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MetronomeBillableMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetronomeBillableMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetronomeBillableMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MetronomeBillableMetric) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


