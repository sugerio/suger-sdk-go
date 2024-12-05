# BillableMetricInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterGroups** | Pointer to [**[]BillableMetricFilterGroup**](BillableMetricFilterGroup.md) | FilterGroups is a list of filter groups. The filterGroups are connected by AND. | [optional] 
**GroupBys** | Pointer to **[]string** | GroupBys is a list of fields to group by. | [optional] 
**PropertyUniqueOn** | Pointer to **string** | The target property for unique count aggregate. | [optional] 

## Methods

### NewBillableMetricInfo

`func NewBillableMetricInfo() *BillableMetricInfo`

NewBillableMetricInfo instantiates a new BillableMetricInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricInfoWithDefaults

`func NewBillableMetricInfoWithDefaults() *BillableMetricInfo`

NewBillableMetricInfoWithDefaults instantiates a new BillableMetricInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterGroups

`func (o *BillableMetricInfo) GetFilterGroups() []BillableMetricFilterGroup`

GetFilterGroups returns the FilterGroups field if non-nil, zero value otherwise.

### GetFilterGroupsOk

`func (o *BillableMetricInfo) GetFilterGroupsOk() (*[]BillableMetricFilterGroup, bool)`

GetFilterGroupsOk returns a tuple with the FilterGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterGroups

`func (o *BillableMetricInfo) SetFilterGroups(v []BillableMetricFilterGroup)`

SetFilterGroups sets FilterGroups field to given value.

### HasFilterGroups

`func (o *BillableMetricInfo) HasFilterGroups() bool`

HasFilterGroups returns a boolean if a field has been set.

### GetGroupBys

`func (o *BillableMetricInfo) GetGroupBys() []string`

GetGroupBys returns the GroupBys field if non-nil, zero value otherwise.

### GetGroupBysOk

`func (o *BillableMetricInfo) GetGroupBysOk() (*[]string, bool)`

GetGroupBysOk returns a tuple with the GroupBys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBys

`func (o *BillableMetricInfo) SetGroupBys(v []string)`

SetGroupBys sets GroupBys field to given value.

### HasGroupBys

`func (o *BillableMetricInfo) HasGroupBys() bool`

HasGroupBys returns a boolean if a field has been set.

### GetPropertyUniqueOn

`func (o *BillableMetricInfo) GetPropertyUniqueOn() string`

GetPropertyUniqueOn returns the PropertyUniqueOn field if non-nil, zero value otherwise.

### GetPropertyUniqueOnOk

`func (o *BillableMetricInfo) GetPropertyUniqueOnOk() (*string, bool)`

GetPropertyUniqueOnOk returns a tuple with the PropertyUniqueOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyUniqueOn

`func (o *BillableMetricInfo) SetPropertyUniqueOn(v string)`

SetPropertyUniqueOn sets PropertyUniqueOn field to given value.

### HasPropertyUniqueOn

`func (o *BillableMetricInfo) HasPropertyUniqueOn() bool`

HasPropertyUniqueOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


