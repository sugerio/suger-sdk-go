# ServiceMarketplaceServiceApiGetAIUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | EndTime is the end of the queried time range. | [optional] 
**HasMore** | Pointer to **bool** | HasMore indicates if there are more records available beyond the current page. | [optional] 
**Limit** | Pointer to **int32** | Limit is the limit used for this query. | [optional] 
**Offset** | Pointer to **int32** | Offset is the offset used for this query. | [optional] 
**Records** | Pointer to [**[]ServiceMarketplaceServiceApiAIUsageRecord**](ServiceMarketplaceServiceApiAIUsageRecord.md) | Records contains the usage records for the time range. | [optional] 
**StartTime** | Pointer to **string** | StartTime is the start of the queried time range. | [optional] 
**TotalCount** | Pointer to **int32** | TotalCount is the total number of records matching the query (before pagination). | [optional] 

## Methods

### NewServiceMarketplaceServiceApiGetAIUsageResponse

`func NewServiceMarketplaceServiceApiGetAIUsageResponse() *ServiceMarketplaceServiceApiGetAIUsageResponse`

NewServiceMarketplaceServiceApiGetAIUsageResponse instantiates a new ServiceMarketplaceServiceApiGetAIUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMarketplaceServiceApiGetAIUsageResponseWithDefaults

`func NewServiceMarketplaceServiceApiGetAIUsageResponseWithDefaults() *ServiceMarketplaceServiceApiGetAIUsageResponse`

NewServiceMarketplaceServiceApiGetAIUsageResponseWithDefaults instantiates a new ServiceMarketplaceServiceApiGetAIUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetHasMore

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetLimit

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetRecords

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetRecords() []ServiceMarketplaceServiceApiAIUsageRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetRecordsOk() (*[]ServiceMarketplaceServiceApiAIUsageRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetRecords(v []ServiceMarketplaceServiceApiAIUsageRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetStartTime

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTotalCount

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ServiceMarketplaceServiceApiGetAIUsageResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


