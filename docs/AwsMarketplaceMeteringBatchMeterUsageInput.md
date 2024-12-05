# AwsMarketplaceMeteringBatchMeterUsageInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductCode** | Pointer to **string** | Product code is used to uniquely identify a product in AWS Marketplace. The product code should be the same as the one used during the publishing of a new product. | [optional] 
**UsageRecords** | Pointer to [**[]AwsMarketplaceMeteringUsageRecord**](AwsMarketplaceMeteringUsageRecord.md) | The set of UsageRecords to submit. BatchMeterUsage accepts up to 25 UsageRecords at a time. | [optional] 

## Methods

### NewAwsMarketplaceMeteringBatchMeterUsageInput

`func NewAwsMarketplaceMeteringBatchMeterUsageInput() *AwsMarketplaceMeteringBatchMeterUsageInput`

NewAwsMarketplaceMeteringBatchMeterUsageInput instantiates a new AwsMarketplaceMeteringBatchMeterUsageInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceMeteringBatchMeterUsageInputWithDefaults

`func NewAwsMarketplaceMeteringBatchMeterUsageInputWithDefaults() *AwsMarketplaceMeteringBatchMeterUsageInput`

NewAwsMarketplaceMeteringBatchMeterUsageInputWithDefaults instantiates a new AwsMarketplaceMeteringBatchMeterUsageInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductCode

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetUsageRecords

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) GetUsageRecords() []AwsMarketplaceMeteringUsageRecord`

GetUsageRecords returns the UsageRecords field if non-nil, zero value otherwise.

### GetUsageRecordsOk

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) GetUsageRecordsOk() (*[]AwsMarketplaceMeteringUsageRecord, bool)`

GetUsageRecordsOk returns a tuple with the UsageRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageRecords

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) SetUsageRecords(v []AwsMarketplaceMeteringUsageRecord)`

SetUsageRecords sets UsageRecords field to given value.

### HasUsageRecords

`func (o *AwsMarketplaceMeteringBatchMeterUsageInput) HasUsageRecords() bool`

HasUsageRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


