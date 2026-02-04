# GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultMetadata** | Pointer to **map[string]interface{}** | Metadata pertaining to the operation&#39;s result. | [optional] 
**Results** | Pointer to [**[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult**](GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult.md) | Contains all UsageRecords processed by BatchMeterUsage . These records were either honored by Amazon Web Services Marketplace Metering Service or were invalid. Invalid records should be fixed before being resubmitted. | [optional] 
**UnprocessedRecords** | Pointer to [**[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord**](GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord.md) | Contains all UsageRecords that were not processed by BatchMeterUsage . This is a list of UsageRecords . You can retry the failed request by making another BatchMeterUsage call with this list as input in the BatchMeterUsageRequest . | [optional] 

## Methods

### NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput() *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput`

NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutputWithDefaults

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutputWithDefaults() *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput`

NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutputWithDefaults instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultMetadata

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) GetResultMetadata() map[string]interface{}`

GetResultMetadata returns the ResultMetadata field if non-nil, zero value otherwise.

### GetResultMetadataOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) GetResultMetadataOk() (*map[string]interface{}, bool)`

GetResultMetadataOk returns a tuple with the ResultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMetadata

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) SetResultMetadata(v map[string]interface{})`

SetResultMetadata sets ResultMetadata field to given value.

### HasResultMetadata

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) HasResultMetadata() bool`

HasResultMetadata returns a boolean if a field has been set.

### GetResults

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) GetResults() []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) GetResultsOk() (*[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) SetResults(v []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetUnprocessedRecords

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) GetUnprocessedRecords() []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord`

GetUnprocessedRecords returns the UnprocessedRecords field if non-nil, zero value otherwise.

### GetUnprocessedRecordsOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) GetUnprocessedRecordsOk() (*[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord, bool)`

GetUnprocessedRecordsOk returns a tuple with the UnprocessedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessedRecords

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) SetUnprocessedRecords(v []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord)`

SetUnprocessedRecords sets UnprocessedRecords field to given value.

### HasUnprocessedRecords

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringBatchMeterUsageOutput) HasUnprocessedRecords() bool`

HasUnprocessedRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


