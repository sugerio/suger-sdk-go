# RevenueRecordInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRevenueRecords** | Pointer to [**[]GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent**](GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent.md) | For raw revenue records in AWS Marketplace | [optional] 
**AzureRevenueRecords** | Pointer to [**[]GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue**](GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue.md) | For raw revenue records in Azure Marketplace | [optional] 
**DisbursementNotificationSent** | Pointer to **bool** | Whether the disbursement notification has been sent to the seller/ISV. | [optional] 
**GcpRevenueRecords** | Pointer to [**[]GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage**](GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage.md) | For raw revenue records in GCP Marketplace | [optional] 
**Resource** | Pointer to **string** | Resource name for the revenue record. Applicable only to GCP Marketplace. | [optional] 

## Methods

### NewRevenueRecordInfo

`func NewRevenueRecordInfo() *RevenueRecordInfo`

NewRevenueRecordInfo instantiates a new RevenueRecordInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueRecordInfoWithDefaults

`func NewRevenueRecordInfoWithDefaults() *RevenueRecordInfo`

NewRevenueRecordInfoWithDefaults instantiates a new RevenueRecordInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRevenueRecords

`func (o *RevenueRecordInfo) GetAwsRevenueRecords() []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent`

GetAwsRevenueRecords returns the AwsRevenueRecords field if non-nil, zero value otherwise.

### GetAwsRevenueRecordsOk

`func (o *RevenueRecordInfo) GetAwsRevenueRecordsOk() (*[]GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent, bool)`

GetAwsRevenueRecordsOk returns a tuple with the AwsRevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRevenueRecords

`func (o *RevenueRecordInfo) SetAwsRevenueRecords(v []GithubComSugerioMarketplaceServiceRdsDbLibBillingAwsBillingEvent)`

SetAwsRevenueRecords sets AwsRevenueRecords field to given value.

### HasAwsRevenueRecords

`func (o *RevenueRecordInfo) HasAwsRevenueRecords() bool`

HasAwsRevenueRecords returns a boolean if a field has been set.

### GetAzureRevenueRecords

`func (o *RevenueRecordInfo) GetAzureRevenueRecords() []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue`

GetAzureRevenueRecords returns the AzureRevenueRecords field if non-nil, zero value otherwise.

### GetAzureRevenueRecordsOk

`func (o *RevenueRecordInfo) GetAzureRevenueRecordsOk() (*[]GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue, bool)`

GetAzureRevenueRecordsOk returns a tuple with the AzureRevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRevenueRecords

`func (o *RevenueRecordInfo) SetAzureRevenueRecords(v []GithubComSugerioMarketplaceServiceRdsDbLibBillingAzureCmaRevenue)`

SetAzureRevenueRecords sets AzureRevenueRecords field to given value.

### HasAzureRevenueRecords

`func (o *RevenueRecordInfo) HasAzureRevenueRecords() bool`

HasAzureRevenueRecords returns a boolean if a field has been set.

### GetDisbursementNotificationSent

`func (o *RevenueRecordInfo) GetDisbursementNotificationSent() bool`

GetDisbursementNotificationSent returns the DisbursementNotificationSent field if non-nil, zero value otherwise.

### GetDisbursementNotificationSentOk

`func (o *RevenueRecordInfo) GetDisbursementNotificationSentOk() (*bool, bool)`

GetDisbursementNotificationSentOk returns a tuple with the DisbursementNotificationSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursementNotificationSent

`func (o *RevenueRecordInfo) SetDisbursementNotificationSent(v bool)`

SetDisbursementNotificationSent sets DisbursementNotificationSent field to given value.

### HasDisbursementNotificationSent

`func (o *RevenueRecordInfo) HasDisbursementNotificationSent() bool`

HasDisbursementNotificationSent returns a boolean if a field has been set.

### GetGcpRevenueRecords

`func (o *RevenueRecordInfo) GetGcpRevenueRecords() []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage`

GetGcpRevenueRecords returns the GcpRevenueRecords field if non-nil, zero value otherwise.

### GetGcpRevenueRecordsOk

`func (o *RevenueRecordInfo) GetGcpRevenueRecordsOk() (*[]GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage, bool)`

GetGcpRevenueRecordsOk returns a tuple with the GcpRevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRevenueRecords

`func (o *RevenueRecordInfo) SetGcpRevenueRecords(v []GithubComSugerioMarketplaceServiceRdsDbLibBillingGcpChargeUsage)`

SetGcpRevenueRecords sets GcpRevenueRecords field to given value.

### HasGcpRevenueRecords

`func (o *RevenueRecordInfo) HasGcpRevenueRecords() bool`

HasGcpRevenueRecords returns a boolean if a field has been set.

### GetResource

`func (o *RevenueRecordInfo) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RevenueRecordInfo) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RevenueRecordInfo) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RevenueRecordInfo) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


