# GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAWSAccountId** | Pointer to **string** | The CustomerAWSAccountID parameter specifies the AWS account ID of the buyer. | [optional] 
**CustomerIdentifier** | Pointer to **string** | The CustomerIdentifier is obtained through the ResolveCustomer operation and represents an individual buyer in your application. | [optional] 
**Dimension** | Pointer to **string** | During the process of registering a product on Amazon Web Services Marketplace, dimensions are specified. These represent different units of value in your application.  This member is required. | [optional] 
**LicenseArn** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** | The quantity of usage consumed by the customer for the given dimension and time. Defaults to 0 if not specified. | [optional] 
**Timestamp** | Pointer to **string** | Timestamp, in UTC, for which the usage is being reported.  Your application can meter usage for up to one hour in the past. Make sure the timestamp value is not before the start of the software usage.  This member is required. | [optional] 
**UsageAllocations** | Pointer to [**[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageAllocation**](GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageAllocation.md) | The set of UsageAllocations to submit. The sum of all UsageAllocation quantities must equal the Quantity of the UsageRecord . | [optional] 

## Methods

### NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord() *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord`

NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordWithDefaults

`func NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordWithDefaults() *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord`

NewGithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecordWithDefaults instantiates a new GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAWSAccountId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetCustomerAWSAccountId() string`

GetCustomerAWSAccountId returns the CustomerAWSAccountId field if non-nil, zero value otherwise.

### GetCustomerAWSAccountIdOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetCustomerAWSAccountIdOk() (*string, bool)`

GetCustomerAWSAccountIdOk returns a tuple with the CustomerAWSAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAWSAccountId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetCustomerAWSAccountId(v string)`

SetCustomerAWSAccountId sets CustomerAWSAccountId field to given value.

### HasCustomerAWSAccountId

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasCustomerAWSAccountId() bool`

HasCustomerAWSAccountId returns a boolean if a field has been set.

### GetCustomerIdentifier

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetCustomerIdentifier() string`

GetCustomerIdentifier returns the CustomerIdentifier field if non-nil, zero value otherwise.

### GetCustomerIdentifierOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetCustomerIdentifierOk() (*string, bool)`

GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdentifier

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetCustomerIdentifier(v string)`

SetCustomerIdentifier sets CustomerIdentifier field to given value.

### HasCustomerIdentifier

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasCustomerIdentifier() bool`

HasCustomerIdentifier returns a boolean if a field has been set.

### GetDimension

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetLicenseArn

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetLicenseArn() string`

GetLicenseArn returns the LicenseArn field if non-nil, zero value otherwise.

### GetLicenseArnOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetLicenseArnOk() (*string, bool)`

GetLicenseArnOk returns a tuple with the LicenseArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseArn

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetLicenseArn(v string)`

SetLicenseArn sets LicenseArn field to given value.

### HasLicenseArn

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasLicenseArn() bool`

HasLicenseArn returns a boolean if a field has been set.

### GetQuantity

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTimestamp

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetUsageAllocations

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetUsageAllocations() []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageAllocation`

GetUsageAllocations returns the UsageAllocations field if non-nil, zero value otherwise.

### GetUsageAllocationsOk

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) GetUsageAllocationsOk() (*[]GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageAllocation, bool)`

GetUsageAllocationsOk returns a tuple with the UsageAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageAllocations

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) SetUsageAllocations(v []GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageAllocation)`

SetUsageAllocations sets UsageAllocations field to given value.

### HasUsageAllocations

`func (o *GithubComAwsAwsSdkGoV2ServiceMarketplacemeteringTypesUsageRecord) HasUsageAllocations() bool`

HasUsageAllocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


