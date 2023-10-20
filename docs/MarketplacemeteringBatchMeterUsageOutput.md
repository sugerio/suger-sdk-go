# MarketplacemeteringBatchMeterUsageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Results** | Pointer to [**[]TypesUsageRecordResult**](TypesUsageRecordResult.md) | Contains all UsageRecords processed by BatchMeterUsage. These records were either honored by AWS Marketplace Metering Service or were invalid. Invalid records should be fixed before being resubmitted. | [optional] 
**UnprocessedRecords** | Pointer to [**[]TypesUsageRecord**](TypesUsageRecord.md) | Contains all UsageRecords that were not processed by BatchMeterUsage. This is a list of UsageRecords. You can retry the failed request by making another BatchMeterUsage call with this list as input in the BatchMeterUsageRequest. | [optional] 

## Methods

### NewMarketplacemeteringBatchMeterUsageOutput

`func NewMarketplacemeteringBatchMeterUsageOutput() *MarketplacemeteringBatchMeterUsageOutput`

NewMarketplacemeteringBatchMeterUsageOutput instantiates a new MarketplacemeteringBatchMeterUsageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacemeteringBatchMeterUsageOutputWithDefaults

`func NewMarketplacemeteringBatchMeterUsageOutputWithDefaults() *MarketplacemeteringBatchMeterUsageOutput`

NewMarketplacemeteringBatchMeterUsageOutputWithDefaults instantiates a new MarketplacemeteringBatchMeterUsageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultMetadata

`func (o *MarketplacemeteringBatchMeterUsageOutput) GetResultMetadata() map[string]interface{}`

GetResultMetadata returns the ResultMetadata field if non-nil, zero value otherwise.

### GetResultMetadataOk

`func (o *MarketplacemeteringBatchMeterUsageOutput) GetResultMetadataOk() (*map[string]interface{}, bool)`

GetResultMetadataOk returns a tuple with the ResultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMetadata

`func (o *MarketplacemeteringBatchMeterUsageOutput) SetResultMetadata(v map[string]interface{})`

SetResultMetadata sets ResultMetadata field to given value.

### HasResultMetadata

`func (o *MarketplacemeteringBatchMeterUsageOutput) HasResultMetadata() bool`

HasResultMetadata returns a boolean if a field has been set.

### GetResults

`func (o *MarketplacemeteringBatchMeterUsageOutput) GetResults() []TypesUsageRecordResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MarketplacemeteringBatchMeterUsageOutput) GetResultsOk() (*[]TypesUsageRecordResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MarketplacemeteringBatchMeterUsageOutput) SetResults(v []TypesUsageRecordResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *MarketplacemeteringBatchMeterUsageOutput) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetUnprocessedRecords

`func (o *MarketplacemeteringBatchMeterUsageOutput) GetUnprocessedRecords() []TypesUsageRecord`

GetUnprocessedRecords returns the UnprocessedRecords field if non-nil, zero value otherwise.

### GetUnprocessedRecordsOk

`func (o *MarketplacemeteringBatchMeterUsageOutput) GetUnprocessedRecordsOk() (*[]TypesUsageRecord, bool)`

GetUnprocessedRecordsOk returns a tuple with the UnprocessedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprocessedRecords

`func (o *MarketplacemeteringBatchMeterUsageOutput) SetUnprocessedRecords(v []TypesUsageRecord)`

SetUnprocessedRecords sets UnprocessedRecords field to given value.

### HasUnprocessedRecords

`func (o *MarketplacemeteringBatchMeterUsageOutput) HasUnprocessedRecords() bool`

HasUnprocessedRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


