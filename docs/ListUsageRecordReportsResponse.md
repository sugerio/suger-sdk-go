# ListUsageRecordReportsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int32** |  | [optional] 
**UsageRecordReports** | Pointer to [**[]MeteringUsageRecordReport**](MeteringUsageRecordReport.md) |  | [optional] 

## Methods

### NewListUsageRecordReportsResponse

`func NewListUsageRecordReportsResponse() *ListUsageRecordReportsResponse`

NewListUsageRecordReportsResponse instantiates a new ListUsageRecordReportsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageRecordReportsResponseWithDefaults

`func NewListUsageRecordReportsResponseWithDefaults() *ListUsageRecordReportsResponse`

NewListUsageRecordReportsResponseWithDefaults instantiates a new ListUsageRecordReportsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListUsageRecordReportsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListUsageRecordReportsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListUsageRecordReportsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListUsageRecordReportsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetUsageRecordReports

`func (o *ListUsageRecordReportsResponse) GetUsageRecordReports() []MeteringUsageRecordReport`

GetUsageRecordReports returns the UsageRecordReports field if non-nil, zero value otherwise.

### GetUsageRecordReportsOk

`func (o *ListUsageRecordReportsResponse) GetUsageRecordReportsOk() (*[]MeteringUsageRecordReport, bool)`

GetUsageRecordReportsOk returns a tuple with the UsageRecordReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecordReports

`func (o *ListUsageRecordReportsResponse) SetUsageRecordReports(v []MeteringUsageRecordReport)`

SetUsageRecordReports sets UsageRecordReports field to given value.

### HasUsageRecordReports

`func (o *ListUsageRecordReportsResponse) HasUsageRecordReports() bool`

HasUsageRecordReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


