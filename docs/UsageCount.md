# UsageCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditCount** | Pointer to **float32** | The count of this dimension usage records that are handled as credit. | [optional] 
**IncludedCount** | Pointer to **float32** | The count of this dimension usage records that are handled as included in IncludedBaseQuantity | [optional] 
**ReportedCount** | Pointer to **float32** | The count of this dimension usage records that are reported to cloud vendors. | [optional] 

## Methods

### NewUsageCount

`func NewUsageCount() *UsageCount`

NewUsageCount instantiates a new UsageCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageCountWithDefaults

`func NewUsageCountWithDefaults() *UsageCount`

NewUsageCountWithDefaults instantiates a new UsageCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditCount

`func (o *UsageCount) GetCreditCount() float32`

GetCreditCount returns the CreditCount field if non-nil, zero value otherwise.

### GetCreditCountOk

`func (o *UsageCount) GetCreditCountOk() (*float32, bool)`

GetCreditCountOk returns a tuple with the CreditCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCount

`func (o *UsageCount) SetCreditCount(v float32)`

SetCreditCount sets CreditCount field to given value.

### HasCreditCount

`func (o *UsageCount) HasCreditCount() bool`

HasCreditCount returns a boolean if a field has been set.

### GetIncludedCount

`func (o *UsageCount) GetIncludedCount() float32`

GetIncludedCount returns the IncludedCount field if non-nil, zero value otherwise.

### GetIncludedCountOk

`func (o *UsageCount) GetIncludedCountOk() (*float32, bool)`

GetIncludedCountOk returns a tuple with the IncludedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedCount

`func (o *UsageCount) SetIncludedCount(v float32)`

SetIncludedCount sets IncludedCount field to given value.

### HasIncludedCount

`func (o *UsageCount) HasIncludedCount() bool`

HasIncludedCount returns a boolean if a field has been set.

### GetReportedCount

`func (o *UsageCount) GetReportedCount() float32`

GetReportedCount returns the ReportedCount field if non-nil, zero value otherwise.

### GetReportedCountOk

`func (o *UsageCount) GetReportedCountOk() (*float32, bool)`

GetReportedCountOk returns a tuple with the ReportedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedCount

`func (o *UsageCount) SetReportedCount(v float32)`

SetReportedCount sets ReportedCount field to given value.

### HasReportedCount

`func (o *UsageCount) HasReportedCount() bool`

HasReportedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


