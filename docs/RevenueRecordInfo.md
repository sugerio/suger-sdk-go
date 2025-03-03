# RevenueRecordInfo

## Properties

 Name                             | Type                                                                                                                                                                       | Description                                                               | Notes      
----------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|------------
 **AwsRevenueRecords**            | Pointer to [**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent.md) | For raw revenue records in AWS Marketplace                                | [optional] 
 **AzureRevenueRecords**          | Pointer to [**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue.md) | For raw revenue records in Azure Marketplace                              | [optional] 
 **CreditAmount**                 | Pointer to **float32**                                                                                                                                                     | The credit amount used in the revenue record.                             | [optional] 
 **DisbursementNotificationSent** | Pointer to **bool**                                                                                                                                                        | Whether the disbursement notification has been sent to the seller/ISV.    | [optional] 
 **GcpRevenueRecords**            | Pointer to [**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage.md)   | For raw revenue records in GCP Marketplace                                | [optional] 
 **IdSource**                     | Pointer to **string**                                                                                                                                                      | Source of the revenue record ID.                                          | [optional] 
 **Resource**                     | Pointer to **string**                                                                                                                                                      | Resource name for the revenue record. Applicable only to GCP Marketplace. | [optional] 

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

`func (o *RevenueRecordInfo) GetAwsRevenueRecords() []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent`

GetAwsRevenueRecords returns the AwsRevenueRecords field if non-nil, zero value otherwise.

### GetAwsRevenueRecordsOk

`func (o *RevenueRecordInfo) GetAwsRevenueRecordsOk() (*[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent, bool)`

GetAwsRevenueRecordsOk returns a tuple with the AwsRevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRevenueRecords

`func (o *RevenueRecordInfo) SetAwsRevenueRecords(v []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent)`

SetAwsRevenueRecords sets AwsRevenueRecords field to given value.

### HasAwsRevenueRecords

`func (o *RevenueRecordInfo) HasAwsRevenueRecords() bool`

HasAwsRevenueRecords returns a boolean if a field has been set.

### GetAzureRevenueRecords

`func (o *RevenueRecordInfo) GetAzureRevenueRecords() []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue`

GetAzureRevenueRecords returns the AzureRevenueRecords field if non-nil, zero value otherwise.

### GetAzureRevenueRecordsOk

`func (o *RevenueRecordInfo) GetAzureRevenueRecordsOk() (*[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue, bool)`

GetAzureRevenueRecordsOk returns a tuple with the AzureRevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRevenueRecords

`func (o *RevenueRecordInfo) SetAzureRevenueRecords(v []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue)`

SetAzureRevenueRecords sets AzureRevenueRecords field to given value.

### HasAzureRevenueRecords

`func (o *RevenueRecordInfo) HasAzureRevenueRecords() bool`

HasAzureRevenueRecords returns a boolean if a field has been set.

### GetCreditAmount

`func (o *RevenueRecordInfo) GetCreditAmount() float32`

GetCreditAmount returns the CreditAmount field if non-nil, zero value otherwise.

### GetCreditAmountOk

`func (o *RevenueRecordInfo) GetCreditAmountOk() (*float32, bool)`

GetCreditAmountOk returns a tuple with the CreditAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmount

`func (o *RevenueRecordInfo) SetCreditAmount(v float32)`

SetCreditAmount sets CreditAmount field to given value.

### HasCreditAmount

`func (o *RevenueRecordInfo) HasCreditAmount() bool`

HasCreditAmount returns a boolean if a field has been set.

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

`func (o *RevenueRecordInfo) GetGcpRevenueRecords() []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage`

GetGcpRevenueRecords returns the GcpRevenueRecords field if non-nil, zero value otherwise.

### GetGcpRevenueRecordsOk

`func (o *RevenueRecordInfo) GetGcpRevenueRecordsOk() (*[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage, bool)`

GetGcpRevenueRecordsOk returns a tuple with the GcpRevenueRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRevenueRecords

`func (o *RevenueRecordInfo) SetGcpRevenueRecords(v []GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage)`

SetGcpRevenueRecords sets GcpRevenueRecords field to given value.

### HasGcpRevenueRecords

`func (o *RevenueRecordInfo) HasGcpRevenueRecords() bool`

HasGcpRevenueRecords returns a boolean if a field has been set.

### GetIdSource

`func (o *RevenueRecordInfo) GetIdSource() string`

GetIdSource returns the IdSource field if non-nil, zero value otherwise.

### GetIdSourceOk

`func (o *RevenueRecordInfo) GetIdSourceOk() (*string, bool)`

GetIdSourceOk returns a tuple with the IdSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdSource

`func (o *RevenueRecordInfo) SetIdSource(v string)`

SetIdSource sets IdSource field to given value.

### HasIdSource

`func (o *RevenueRecordInfo) HasIdSource() bool`

HasIdSource returns a boolean if a field has been set.

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


