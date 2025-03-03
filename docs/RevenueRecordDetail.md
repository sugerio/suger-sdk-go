# RevenueRecordDetail

## Properties

 Name                         | Type                                                                                                                                                                     | Description           | Notes      
------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------|------------
 **AwsRevenueRecordDetail**   | Pointer to [**GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent.md) | For AWS Marketplace   | [optional] 
 **AzureRevenueRecordDetail** | Pointer to [**GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue.md) | For Azure Marketplace | [optional] 
 **GcpRevenueRecordDetail**   | Pointer to [**GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage.md)   | For GCP Marketplace   | [optional] 

## Methods

### NewRevenueRecordDetail

`func NewRevenueRecordDetail() *RevenueRecordDetail`

NewRevenueRecordDetail instantiates a new RevenueRecordDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevenueRecordDetailWithDefaults

`func NewRevenueRecordDetailWithDefaults() *RevenueRecordDetail`

NewRevenueRecordDetailWithDefaults instantiates a new RevenueRecordDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRevenueRecordDetail

`func (o *RevenueRecordDetail) GetAwsRevenueRecordDetail() GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent`

GetAwsRevenueRecordDetail returns the AwsRevenueRecordDetail field if non-nil, zero value otherwise.

### GetAwsRevenueRecordDetailOk

`func (o *RevenueRecordDetail) GetAwsRevenueRecordDetailOk() (*GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent, bool)`

GetAwsRevenueRecordDetailOk returns a tuple with the AwsRevenueRecordDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRevenueRecordDetail

`func (o *RevenueRecordDetail) SetAwsRevenueRecordDetail(v GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent)`

SetAwsRevenueRecordDetail sets AwsRevenueRecordDetail field to given value.

### HasAwsRevenueRecordDetail

`func (o *RevenueRecordDetail) HasAwsRevenueRecordDetail() bool`

HasAwsRevenueRecordDetail returns a boolean if a field has been set.

### GetAzureRevenueRecordDetail

`func (o *RevenueRecordDetail) GetAzureRevenueRecordDetail() GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue`

GetAzureRevenueRecordDetail returns the AzureRevenueRecordDetail field if non-nil, zero value otherwise.

### GetAzureRevenueRecordDetailOk

`func (o *RevenueRecordDetail) GetAzureRevenueRecordDetailOk() (*GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue, bool)`

GetAzureRevenueRecordDetailOk returns a tuple with the AzureRevenueRecordDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureRevenueRecordDetail

`func (o *RevenueRecordDetail) SetAzureRevenueRecordDetail(v GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue)`

SetAzureRevenueRecordDetail sets AzureRevenueRecordDetail field to given value.

### HasAzureRevenueRecordDetail

`func (o *RevenueRecordDetail) HasAzureRevenueRecordDetail() bool`

HasAzureRevenueRecordDetail returns a boolean if a field has been set.

### GetGcpRevenueRecordDetail

`func (o *RevenueRecordDetail) GetGcpRevenueRecordDetail() GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage`

GetGcpRevenueRecordDetail returns the GcpRevenueRecordDetail field if non-nil, zero value otherwise.

### GetGcpRevenueRecordDetailOk

`func (o *RevenueRecordDetail) GetGcpRevenueRecordDetailOk() (*GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage, bool)`

GetGcpRevenueRecordDetailOk returns a tuple with the GcpRevenueRecordDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpRevenueRecordDetail

`func (o *RevenueRecordDetail) SetGcpRevenueRecordDetail(v GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage)`

SetGcpRevenueRecordDetail sets GcpRevenueRecordDetail field to given value.

### HasGcpRevenueRecordDetail

`func (o *RevenueRecordDetail) HasGcpRevenueRecordDetail() bool`

HasGcpRevenueRecordDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


