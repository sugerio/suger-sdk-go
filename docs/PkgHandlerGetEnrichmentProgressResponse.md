# PkgHandlerGetEnrichmentProgressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CycleProgressPercent** | Pointer to **int32** | CycleProgressPercent is the percentage of the cycle completed (0-100). | [optional] 
**CycleStartedAt** | Pointer to **string** | CycleStartedAt is when the current refresh cycle began. | [optional] 
**EstimatedRunsRemaining** | Pointer to **int32** | EstimatedRunsRemaining is the estimated number of runs to complete the cycle. | [optional] 
**LastCycleCompletedAt** | Pointer to **string** | LastCycleCompletedAt is when the previous cycle was completed. | [optional] 
**LastRunAt** | Pointer to **string** | LastRunAt is the timestamp of the most recent enrichment run. | [optional] 
**LastRunRecordsProcessed** | Pointer to **int32** | LastRunRecordsProcessed is the number of records processed in the most recent run. | [optional] 
**NextCycleStartsAt** | Pointer to **string** | NextCycleStartsAt is when the next cycle will start (only set when status is \&quot;waiting\&quot;). | [optional] 
**RecordsProcessedInCycle** | Pointer to **int32** | RecordsProcessedInCycle is the count of records processed in the current cycle. | [optional] 
**Status** | Pointer to **string** | Status indicates the current state of enrichment. Values: \&quot;active\&quot;, \&quot;not_started\&quot;, \&quot;in_progress\&quot;, \&quot;waiting\&quot; - \&quot;active\&quot;: Enrichment is running but refresh cycles are disabled (RefreshAfterDays &#x3D; -1) - \&quot;not_started\&quot;: No enrichment cycle has run yet - \&quot;in_progress\&quot;: Currently processing records in a cycle - \&quot;waiting\&quot;: Cycle completed, waiting for next scheduled run | [optional] 
**TotalMatchingRecords** | Pointer to **int32** | TotalMatchingRecords is the cached count of records matching the query. | [optional] 

## Methods

### NewPkgHandlerGetEnrichmentProgressResponse

`func NewPkgHandlerGetEnrichmentProgressResponse() *PkgHandlerGetEnrichmentProgressResponse`

NewPkgHandlerGetEnrichmentProgressResponse instantiates a new PkgHandlerGetEnrichmentProgressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkgHandlerGetEnrichmentProgressResponseWithDefaults

`func NewPkgHandlerGetEnrichmentProgressResponseWithDefaults() *PkgHandlerGetEnrichmentProgressResponse`

NewPkgHandlerGetEnrichmentProgressResponseWithDefaults instantiates a new PkgHandlerGetEnrichmentProgressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycleProgressPercent

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetCycleProgressPercent() int32`

GetCycleProgressPercent returns the CycleProgressPercent field if non-nil, zero value otherwise.

### GetCycleProgressPercentOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetCycleProgressPercentOk() (*int32, bool)`

GetCycleProgressPercentOk returns a tuple with the CycleProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleProgressPercent

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetCycleProgressPercent(v int32)`

SetCycleProgressPercent sets CycleProgressPercent field to given value.

### HasCycleProgressPercent

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasCycleProgressPercent() bool`

HasCycleProgressPercent returns a boolean if a field has been set.

### GetCycleStartedAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetCycleStartedAt() string`

GetCycleStartedAt returns the CycleStartedAt field if non-nil, zero value otherwise.

### GetCycleStartedAtOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetCycleStartedAtOk() (*string, bool)`

GetCycleStartedAtOk returns a tuple with the CycleStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleStartedAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetCycleStartedAt(v string)`

SetCycleStartedAt sets CycleStartedAt field to given value.

### HasCycleStartedAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasCycleStartedAt() bool`

HasCycleStartedAt returns a boolean if a field has been set.

### GetEstimatedRunsRemaining

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetEstimatedRunsRemaining() int32`

GetEstimatedRunsRemaining returns the EstimatedRunsRemaining field if non-nil, zero value otherwise.

### GetEstimatedRunsRemainingOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetEstimatedRunsRemainingOk() (*int32, bool)`

GetEstimatedRunsRemainingOk returns a tuple with the EstimatedRunsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedRunsRemaining

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetEstimatedRunsRemaining(v int32)`

SetEstimatedRunsRemaining sets EstimatedRunsRemaining field to given value.

### HasEstimatedRunsRemaining

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasEstimatedRunsRemaining() bool`

HasEstimatedRunsRemaining returns a boolean if a field has been set.

### GetLastCycleCompletedAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetLastCycleCompletedAt() string`

GetLastCycleCompletedAt returns the LastCycleCompletedAt field if non-nil, zero value otherwise.

### GetLastCycleCompletedAtOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetLastCycleCompletedAtOk() (*string, bool)`

GetLastCycleCompletedAtOk returns a tuple with the LastCycleCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCycleCompletedAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetLastCycleCompletedAt(v string)`

SetLastCycleCompletedAt sets LastCycleCompletedAt field to given value.

### HasLastCycleCompletedAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasLastCycleCompletedAt() bool`

HasLastCycleCompletedAt returns a boolean if a field has been set.

### GetLastRunAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetLastRunAt() string`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetLastRunAtOk() (*string, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetLastRunAt(v string)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### GetLastRunRecordsProcessed

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetLastRunRecordsProcessed() int32`

GetLastRunRecordsProcessed returns the LastRunRecordsProcessed field if non-nil, zero value otherwise.

### GetLastRunRecordsProcessedOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetLastRunRecordsProcessedOk() (*int32, bool)`

GetLastRunRecordsProcessedOk returns a tuple with the LastRunRecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunRecordsProcessed

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetLastRunRecordsProcessed(v int32)`

SetLastRunRecordsProcessed sets LastRunRecordsProcessed field to given value.

### HasLastRunRecordsProcessed

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasLastRunRecordsProcessed() bool`

HasLastRunRecordsProcessed returns a boolean if a field has been set.

### GetNextCycleStartsAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetNextCycleStartsAt() string`

GetNextCycleStartsAt returns the NextCycleStartsAt field if non-nil, zero value otherwise.

### GetNextCycleStartsAtOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetNextCycleStartsAtOk() (*string, bool)`

GetNextCycleStartsAtOk returns a tuple with the NextCycleStartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCycleStartsAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetNextCycleStartsAt(v string)`

SetNextCycleStartsAt sets NextCycleStartsAt field to given value.

### HasNextCycleStartsAt

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasNextCycleStartsAt() bool`

HasNextCycleStartsAt returns a boolean if a field has been set.

### GetRecordsProcessedInCycle

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetRecordsProcessedInCycle() int32`

GetRecordsProcessedInCycle returns the RecordsProcessedInCycle field if non-nil, zero value otherwise.

### GetRecordsProcessedInCycleOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetRecordsProcessedInCycleOk() (*int32, bool)`

GetRecordsProcessedInCycleOk returns a tuple with the RecordsProcessedInCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsProcessedInCycle

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetRecordsProcessedInCycle(v int32)`

SetRecordsProcessedInCycle sets RecordsProcessedInCycle field to given value.

### HasRecordsProcessedInCycle

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasRecordsProcessedInCycle() bool`

HasRecordsProcessedInCycle returns a boolean if a field has been set.

### GetStatus

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalMatchingRecords

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetTotalMatchingRecords() int32`

GetTotalMatchingRecords returns the TotalMatchingRecords field if non-nil, zero value otherwise.

### GetTotalMatchingRecordsOk

`func (o *PkgHandlerGetEnrichmentProgressResponse) GetTotalMatchingRecordsOk() (*int32, bool)`

GetTotalMatchingRecordsOk returns a tuple with the TotalMatchingRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMatchingRecords

`func (o *PkgHandlerGetEnrichmentProgressResponse) SetTotalMatchingRecords(v int32)`

SetTotalMatchingRecords sets TotalMatchingRecords field to given value.

### HasTotalMatchingRecords

`func (o *PkgHandlerGetEnrichmentProgressResponse) HasTotalMatchingRecords() bool`

HasTotalMatchingRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


