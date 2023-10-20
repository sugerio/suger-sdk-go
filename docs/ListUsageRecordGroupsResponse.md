# ListUsageRecordGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int32** |  | [optional] 
**UsageRecordGroups** | Pointer to [**[]MeteringUsageRecordGroup**](MeteringUsageRecordGroup.md) |  | [optional] 

## Methods

### NewListUsageRecordGroupsResponse

`func NewListUsageRecordGroupsResponse() *ListUsageRecordGroupsResponse`

NewListUsageRecordGroupsResponse instantiates a new ListUsageRecordGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageRecordGroupsResponseWithDefaults

`func NewListUsageRecordGroupsResponseWithDefaults() *ListUsageRecordGroupsResponse`

NewListUsageRecordGroupsResponseWithDefaults instantiates a new ListUsageRecordGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListUsageRecordGroupsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListUsageRecordGroupsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListUsageRecordGroupsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListUsageRecordGroupsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetUsageRecordGroups

`func (o *ListUsageRecordGroupsResponse) GetUsageRecordGroups() []MeteringUsageRecordGroup`

GetUsageRecordGroups returns the UsageRecordGroups field if non-nil, zero value otherwise.

### GetUsageRecordGroupsOk

`func (o *ListUsageRecordGroupsResponse) GetUsageRecordGroupsOk() (*[]MeteringUsageRecordGroup, bool)`

GetUsageRecordGroupsOk returns a tuple with the UsageRecordGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecordGroups

`func (o *ListUsageRecordGroupsResponse) SetUsageRecordGroups(v []MeteringUsageRecordGroup)`

SetUsageRecordGroups sets UsageRecordGroups field to given value.

### HasUsageRecordGroups

`func (o *ListUsageRecordGroupsResponse) HasUsageRecordGroups() bool`

HasUsageRecordGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


