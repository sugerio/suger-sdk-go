# GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeteringRecordId** | Pointer to **string** | The MeteringRecordId is a unique identifier for this metering event. | [optional] 
**Status** | Pointer to [**GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultStatus**](GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultStatus.md) | The UsageRecordResult Status indicates the status of an individual UsageRecord processed by BatchMeterUsage .    - Success- The UsageRecord was accepted and honored by BatchMeterUsage .    - CustomerNotSubscribed- The CustomerIdentifier specified is not able to use   your product. The UsageRecord was not honored. There are three causes for this   result:    - The customer identifier is invalid.    - The customer identifier provided in the metering record does not have an   active agreement or subscription with this product. Future UsageRecords for   this customer will fail until the customer subscribes to your product.    - The customer&#39;s Amazon Web Services account was suspended.    - DuplicateRecord- Indicates that the UsageRecord was invalid and not honored.   A previously metered UsageRecord had the same customer, dimension, and time,   but a different quantity. | [optional] 
**UsageRecord** | Pointer to [**GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord**](GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord.md) | The UsageRecord that was part of the BatchMeterUsage request. | [optional] 

## Methods

### NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult() *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult`

NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultWithDefaults

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultWithDefaults() *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult`

NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultWithDefaults instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeteringRecordId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) GetMeteringRecordId() string`

GetMeteringRecordId returns the MeteringRecordId field if non-nil, zero value otherwise.

### GetMeteringRecordIdOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) GetMeteringRecordIdOk() (*string, bool)`

GetMeteringRecordIdOk returns a tuple with the MeteringRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteringRecordId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) SetMeteringRecordId(v string)`

SetMeteringRecordId sets MeteringRecordId field to given value.

### HasMeteringRecordId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) HasMeteringRecordId() bool`

HasMeteringRecordId returns a boolean if a field has been set.

### GetStatus

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) GetStatus() GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) GetStatusOk() (*GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) SetStatus(v GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResultStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsageRecord

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) GetUsageRecord() GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord`

GetUsageRecord returns the UsageRecord field if non-nil, zero value otherwise.

### GetUsageRecordOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) GetUsageRecordOk() (*GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord, bool)`

GetUsageRecordOk returns a tuple with the UsageRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecord

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) SetUsageRecord(v GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord)`

SetUsageRecord sets UsageRecord field to given value.

### HasUsageRecord

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordResult) HasUsageRecord() bool`

HasUsageRecord returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


