# AwsPaymentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AWS Account ID (payer) | [optional] 
**Amount** | Pointer to **float32** | Payment amount (positive for payments, negative for refunds) | [optional] 
**BillingEntity** | Pointer to **string** | BillingEntity holds the value of the \&quot;billing_entity\&quot; field. | [optional] 
**BillingPeriodEnd** | Pointer to **string** | End of billing period this payment covers | [optional] 
**BillingPeriodStart** | Pointer to **string** | Start of billing period this payment covers | [optional] 
**ChargeType** | Pointer to **string** | Anniversary (AWS) or Subscription (Marketplace) | [optional] 
**CurrencyCode** | Pointer to **string** | CurrencyCode holds the value of the \&quot;currency_code\&quot; field. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoiceId** | Pointer to **string** | Links to buyer_aws_invoice.invoice_id | [optional] 
**OrganizationID** | Pointer to **string** | OrganizationID holds the value of the \&quot;organization_id\&quot; field. | [optional] 
**PaymentMethodDescription** | Pointer to **string** | User-friendly description (e.g., &#39;Bank account ending in 901&#39;) | [optional] 
**PaymentMethodType** | Pointer to **string** | BankAccount or CreditCard | [optional] 
**SyncTimestamp** | Pointer to **string** | SyncTimestamp holds the value of the \&quot;sync_timestamp\&quot; field. | [optional] 
**SystemIdentifierArn** | Pointer to **string** | ARN from console API | [optional] 
**TransactionDate** | Pointer to **string** | When payment was processed | [optional] 
**TransactionType** | Pointer to **string** | TransactionType holds the value of the \&quot;transaction_type\&quot; field. | [optional] 

## Methods

### NewAwsPaymentTransaction

`func NewAwsPaymentTransaction() *AwsPaymentTransaction`

NewAwsPaymentTransaction instantiates a new AwsPaymentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsPaymentTransactionWithDefaults

`func NewAwsPaymentTransactionWithDefaults() *AwsPaymentTransaction`

NewAwsPaymentTransactionWithDefaults instantiates a new AwsPaymentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AwsPaymentTransaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsPaymentTransaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsPaymentTransaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AwsPaymentTransaction) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *AwsPaymentTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AwsPaymentTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AwsPaymentTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AwsPaymentTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillingEntity

`func (o *AwsPaymentTransaction) GetBillingEntity() string`

GetBillingEntity returns the BillingEntity field if non-nil, zero value otherwise.

### GetBillingEntityOk

`func (o *AwsPaymentTransaction) GetBillingEntityOk() (*string, bool)`

GetBillingEntityOk returns a tuple with the BillingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEntity

`func (o *AwsPaymentTransaction) SetBillingEntity(v string)`

SetBillingEntity sets BillingEntity field to given value.

### HasBillingEntity

`func (o *AwsPaymentTransaction) HasBillingEntity() bool`

HasBillingEntity returns a boolean if a field has been set.

### GetBillingPeriodEnd

`func (o *AwsPaymentTransaction) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *AwsPaymentTransaction) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *AwsPaymentTransaction) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.

### HasBillingPeriodEnd

`func (o *AwsPaymentTransaction) HasBillingPeriodEnd() bool`

HasBillingPeriodEnd returns a boolean if a field has been set.

### GetBillingPeriodStart

`func (o *AwsPaymentTransaction) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *AwsPaymentTransaction) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *AwsPaymentTransaction) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.

### HasBillingPeriodStart

`func (o *AwsPaymentTransaction) HasBillingPeriodStart() bool`

HasBillingPeriodStart returns a boolean if a field has been set.

### GetChargeType

`func (o *AwsPaymentTransaction) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *AwsPaymentTransaction) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *AwsPaymentTransaction) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *AwsPaymentTransaction) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AwsPaymentTransaction) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AwsPaymentTransaction) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AwsPaymentTransaction) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AwsPaymentTransaction) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetId

`func (o *AwsPaymentTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsPaymentTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsPaymentTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsPaymentTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *AwsPaymentTransaction) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *AwsPaymentTransaction) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *AwsPaymentTransaction) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *AwsPaymentTransaction) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetOrganizationID

`func (o *AwsPaymentTransaction) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *AwsPaymentTransaction) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *AwsPaymentTransaction) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *AwsPaymentTransaction) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPaymentMethodDescription

`func (o *AwsPaymentTransaction) GetPaymentMethodDescription() string`

GetPaymentMethodDescription returns the PaymentMethodDescription field if non-nil, zero value otherwise.

### GetPaymentMethodDescriptionOk

`func (o *AwsPaymentTransaction) GetPaymentMethodDescriptionOk() (*string, bool)`

GetPaymentMethodDescriptionOk returns a tuple with the PaymentMethodDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodDescription

`func (o *AwsPaymentTransaction) SetPaymentMethodDescription(v string)`

SetPaymentMethodDescription sets PaymentMethodDescription field to given value.

### HasPaymentMethodDescription

`func (o *AwsPaymentTransaction) HasPaymentMethodDescription() bool`

HasPaymentMethodDescription returns a boolean if a field has been set.

### GetPaymentMethodType

`func (o *AwsPaymentTransaction) GetPaymentMethodType() string`

GetPaymentMethodType returns the PaymentMethodType field if non-nil, zero value otherwise.

### GetPaymentMethodTypeOk

`func (o *AwsPaymentTransaction) GetPaymentMethodTypeOk() (*string, bool)`

GetPaymentMethodTypeOk returns a tuple with the PaymentMethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodType

`func (o *AwsPaymentTransaction) SetPaymentMethodType(v string)`

SetPaymentMethodType sets PaymentMethodType field to given value.

### HasPaymentMethodType

`func (o *AwsPaymentTransaction) HasPaymentMethodType() bool`

HasPaymentMethodType returns a boolean if a field has been set.

### GetSyncTimestamp

`func (o *AwsPaymentTransaction) GetSyncTimestamp() string`

GetSyncTimestamp returns the SyncTimestamp field if non-nil, zero value otherwise.

### GetSyncTimestampOk

`func (o *AwsPaymentTransaction) GetSyncTimestampOk() (*string, bool)`

GetSyncTimestampOk returns a tuple with the SyncTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTimestamp

`func (o *AwsPaymentTransaction) SetSyncTimestamp(v string)`

SetSyncTimestamp sets SyncTimestamp field to given value.

### HasSyncTimestamp

`func (o *AwsPaymentTransaction) HasSyncTimestamp() bool`

HasSyncTimestamp returns a boolean if a field has been set.

### GetSystemIdentifierArn

`func (o *AwsPaymentTransaction) GetSystemIdentifierArn() string`

GetSystemIdentifierArn returns the SystemIdentifierArn field if non-nil, zero value otherwise.

### GetSystemIdentifierArnOk

`func (o *AwsPaymentTransaction) GetSystemIdentifierArnOk() (*string, bool)`

GetSystemIdentifierArnOk returns a tuple with the SystemIdentifierArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentifierArn

`func (o *AwsPaymentTransaction) SetSystemIdentifierArn(v string)`

SetSystemIdentifierArn sets SystemIdentifierArn field to given value.

### HasSystemIdentifierArn

`func (o *AwsPaymentTransaction) HasSystemIdentifierArn() bool`

HasSystemIdentifierArn returns a boolean if a field has been set.

### GetTransactionDate

`func (o *AwsPaymentTransaction) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *AwsPaymentTransaction) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *AwsPaymentTransaction) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *AwsPaymentTransaction) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionType

`func (o *AwsPaymentTransaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *AwsPaymentTransaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *AwsPaymentTransaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *AwsPaymentTransaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


