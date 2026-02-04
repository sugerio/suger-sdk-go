# RevenueRecordInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRevenueRecords** | Pointer to [**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAwsBillingEvent.md) | For raw revenue records in AWS Marketplace | [optional] 
**AzureRevenueRecords** | Pointer to [**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingAzureCmaRevenue.md) | For raw revenue records in Azure Marketplace | [optional] 
**BankTraceId** | Pointer to **string** | The bank trace ID of the revenue record if applicable | [optional] 
**BillingModel** | Pointer to [**RevenueBillingModel**](RevenueBillingModel.md) | The billing model of the revenue record if applicable The value is one of the following: - SubscriptionBased: The revenue record is from a subscription or recurring commitment. - UsageBased: The revenue record is from a usage-based metering. | [optional] 
**Channel** | Pointer to [**RevenueChannel**](RevenueChannel.md) | The channel of revenue record. | [optional] 
**ChargebeeInvoiceId** | Pointer to **string** | The Chargebee invoice ID for this revenue record if synced to Chargebee | [optional] 
**ChargebeePaymentTransactionId** | Pointer to **string** | The Chargebee payment transaction ID for this revenue record if a payment was recorded in Chargebee | [optional] 
**ChargebeeRefundTransactionId** | Pointer to **string** | The Chargebee refund transaction ID for this revenue record if a refund was recorded in Chargebee | [optional] 
**CreditAmount** | Pointer to **float32** | The credit amount used in the revenue record. | [optional] 
**DisbursementBillingEventId** | Pointer to **string** | The disbursement ID of the revenue record if applicable | [optional] 
**DisbursementNotificationSent** | Pointer to **bool** | Whether the disbursement notification has been sent to the seller/ISV. | [optional] 
**EarningId** | Pointer to **string** | The earning ID of the revenue record if applicable | [optional] 
**GcpRevenueRecords** | Pointer to [**[]GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage**](GithubComSugerioMarketplaceServicePkgLegacyRdsDbLibBillingGcpChargeUsage.md) | For raw revenue records in GCP Marketplace | [optional] 
**IdSource** | Pointer to **string** | Source of the revenue record ID. | [optional] 
**InvoiceId** | Pointer to **string** | The invoice ID of the revenue record if applicable | [optional] 
**PaymentId** | Pointer to **string** | The payment  ID of the revenue record if applicable | [optional] 
**ResellerId** | Pointer to **string** | The reseller ID of the revenue record if applicable | [optional] 
**ResellerName** | Pointer to **string** | The reseller name of the revenue record if application | [optional] 
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

### GetBankTraceId

`func (o *RevenueRecordInfo) GetBankTraceId() string`

GetBankTraceId returns the BankTraceId field if non-nil, zero value otherwise.

### GetBankTraceIdOk

`func (o *RevenueRecordInfo) GetBankTraceIdOk() (*string, bool)`

GetBankTraceIdOk returns a tuple with the BankTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTraceId

`func (o *RevenueRecordInfo) SetBankTraceId(v string)`

SetBankTraceId sets BankTraceId field to given value.

### HasBankTraceId

`func (o *RevenueRecordInfo) HasBankTraceId() bool`

HasBankTraceId returns a boolean if a field has been set.

### GetBillingModel

`func (o *RevenueRecordInfo) GetBillingModel() RevenueBillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *RevenueRecordInfo) GetBillingModelOk() (*RevenueBillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *RevenueRecordInfo) SetBillingModel(v RevenueBillingModel)`

SetBillingModel sets BillingModel field to given value.

### HasBillingModel

`func (o *RevenueRecordInfo) HasBillingModel() bool`

HasBillingModel returns a boolean if a field has been set.

### GetChannel

`func (o *RevenueRecordInfo) GetChannel() RevenueChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RevenueRecordInfo) GetChannelOk() (*RevenueChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RevenueRecordInfo) SetChannel(v RevenueChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *RevenueRecordInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChargebeeInvoiceId

`func (o *RevenueRecordInfo) GetChargebeeInvoiceId() string`

GetChargebeeInvoiceId returns the ChargebeeInvoiceId field if non-nil, zero value otherwise.

### GetChargebeeInvoiceIdOk

`func (o *RevenueRecordInfo) GetChargebeeInvoiceIdOk() (*string, bool)`

GetChargebeeInvoiceIdOk returns a tuple with the ChargebeeInvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargebeeInvoiceId

`func (o *RevenueRecordInfo) SetChargebeeInvoiceId(v string)`

SetChargebeeInvoiceId sets ChargebeeInvoiceId field to given value.

### HasChargebeeInvoiceId

`func (o *RevenueRecordInfo) HasChargebeeInvoiceId() bool`

HasChargebeeInvoiceId returns a boolean if a field has been set.

### GetChargebeePaymentTransactionId

`func (o *RevenueRecordInfo) GetChargebeePaymentTransactionId() string`

GetChargebeePaymentTransactionId returns the ChargebeePaymentTransactionId field if non-nil, zero value otherwise.

### GetChargebeePaymentTransactionIdOk

`func (o *RevenueRecordInfo) GetChargebeePaymentTransactionIdOk() (*string, bool)`

GetChargebeePaymentTransactionIdOk returns a tuple with the ChargebeePaymentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargebeePaymentTransactionId

`func (o *RevenueRecordInfo) SetChargebeePaymentTransactionId(v string)`

SetChargebeePaymentTransactionId sets ChargebeePaymentTransactionId field to given value.

### HasChargebeePaymentTransactionId

`func (o *RevenueRecordInfo) HasChargebeePaymentTransactionId() bool`

HasChargebeePaymentTransactionId returns a boolean if a field has been set.

### GetChargebeeRefundTransactionId

`func (o *RevenueRecordInfo) GetChargebeeRefundTransactionId() string`

GetChargebeeRefundTransactionId returns the ChargebeeRefundTransactionId field if non-nil, zero value otherwise.

### GetChargebeeRefundTransactionIdOk

`func (o *RevenueRecordInfo) GetChargebeeRefundTransactionIdOk() (*string, bool)`

GetChargebeeRefundTransactionIdOk returns a tuple with the ChargebeeRefundTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargebeeRefundTransactionId

`func (o *RevenueRecordInfo) SetChargebeeRefundTransactionId(v string)`

SetChargebeeRefundTransactionId sets ChargebeeRefundTransactionId field to given value.

### HasChargebeeRefundTransactionId

`func (o *RevenueRecordInfo) HasChargebeeRefundTransactionId() bool`

HasChargebeeRefundTransactionId returns a boolean if a field has been set.

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

### GetDisbursementBillingEventId

`func (o *RevenueRecordInfo) GetDisbursementBillingEventId() string`

GetDisbursementBillingEventId returns the DisbursementBillingEventId field if non-nil, zero value otherwise.

### GetDisbursementBillingEventIdOk

`func (o *RevenueRecordInfo) GetDisbursementBillingEventIdOk() (*string, bool)`

GetDisbursementBillingEventIdOk returns a tuple with the DisbursementBillingEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursementBillingEventId

`func (o *RevenueRecordInfo) SetDisbursementBillingEventId(v string)`

SetDisbursementBillingEventId sets DisbursementBillingEventId field to given value.

### HasDisbursementBillingEventId

`func (o *RevenueRecordInfo) HasDisbursementBillingEventId() bool`

HasDisbursementBillingEventId returns a boolean if a field has been set.

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

### GetEarningId

`func (o *RevenueRecordInfo) GetEarningId() string`

GetEarningId returns the EarningId field if non-nil, zero value otherwise.

### GetEarningIdOk

`func (o *RevenueRecordInfo) GetEarningIdOk() (*string, bool)`

GetEarningIdOk returns a tuple with the EarningId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningId

`func (o *RevenueRecordInfo) SetEarningId(v string)`

SetEarningId sets EarningId field to given value.

### HasEarningId

`func (o *RevenueRecordInfo) HasEarningId() bool`

HasEarningId returns a boolean if a field has been set.

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

### GetInvoiceId

`func (o *RevenueRecordInfo) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RevenueRecordInfo) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RevenueRecordInfo) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *RevenueRecordInfo) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPaymentId

`func (o *RevenueRecordInfo) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *RevenueRecordInfo) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *RevenueRecordInfo) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *RevenueRecordInfo) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetResellerId

`func (o *RevenueRecordInfo) GetResellerId() string`

GetResellerId returns the ResellerId field if non-nil, zero value otherwise.

### GetResellerIdOk

`func (o *RevenueRecordInfo) GetResellerIdOk() (*string, bool)`

GetResellerIdOk returns a tuple with the ResellerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerId

`func (o *RevenueRecordInfo) SetResellerId(v string)`

SetResellerId sets ResellerId field to given value.

### HasResellerId

`func (o *RevenueRecordInfo) HasResellerId() bool`

HasResellerId returns a boolean if a field has been set.

### GetResellerName

`func (o *RevenueRecordInfo) GetResellerName() string`

GetResellerName returns the ResellerName field if non-nil, zero value otherwise.

### GetResellerNameOk

`func (o *RevenueRecordInfo) GetResellerNameOk() (*string, bool)`

GetResellerNameOk returns a tuple with the ResellerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerName

`func (o *RevenueRecordInfo) SetResellerName(v string)`

SetResellerName sets ResellerName field to given value.

### HasResellerName

`func (o *RevenueRecordInfo) HasResellerName() bool`

HasResellerName returns a boolean if a field has been set.

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


