# AwsInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | AWS Account ID (12 digits) | [optional] 
**AmountDue** | Pointer to **float32** | Outstanding balance | [optional] 
**BillingEntity** | Pointer to **string** | AWS or AWS_MARKETPLACE - single source of truth for invoice type | [optional] 
**BillingMonth** | Pointer to **int32** | Month number (1-12) | [optional] 
**BillingPeriodEnd** | Pointer to **string** | End of billing period | [optional] 
**BillingPeriodStart** | Pointer to **string** | Start of billing period | [optional] 
**BillingYear** | Pointer to **int32** | Year (YYYY) | [optional] 
**ChargeType** | Pointer to **string** | Anniversary (AWS) or Subscription (Marketplace) | [optional] 
**CurrencyCode** | Pointer to **string** | CurrencyCode holds the value of the \&quot;currency_code\&quot; field. | [optional] 
**DueDate** | Pointer to **string** | Payment due date | [optional] 
**Id** | Pointer to **string** | AWS invoice ID (e.g., &#39;2289273833&#39;) | [optional] 
**InvoiceType** | Pointer to **string** | INVOICE or CREDIT_MEMO | [optional] 
**IssuedDate** | Pointer to **string** | When invoice was issued | [optional] 
**LineItems** | Pointer to [**AwsInvoiceLineItems**](AwsInvoiceLineItems.md) | Extracted line items from PDF processing | [optional] 
**PdfDownloaded** | Pointer to **bool** | PDF download status | [optional] 
**SettleDate** | Pointer to **string** | Invoice settle date | [optional] 
**Status** | Pointer to **string** | Due, PastDue, Scheduled, Processing, Paid | [optional] 
**SyncTimestamp** | Pointer to **string** | SyncTimestamp holds the value of the \&quot;sync_timestamp\&quot; field. | [optional] 
**SystemIdentifierArn** | Pointer to **string** | ARN from console API (e.g., arn:aws:payments:us-east-1:account:invoice:id) | [optional] 
**TaxAmount** | Pointer to **float32** | Tax amount | [optional] 
**TotalAmount** | Pointer to **float32** | Total invoice amount | [optional] 
**TotalAmountBeforeTax** | Pointer to **float32** | Subtotal before tax | [optional] 
**TransactionStatus** | Pointer to **string** | Forgive, Successful, or Unknown | [optional] 

## Methods

### NewAwsInvoice

`func NewAwsInvoice() *AwsInvoice`

NewAwsInvoice instantiates a new AwsInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsInvoiceWithDefaults

`func NewAwsInvoiceWithDefaults() *AwsInvoice`

NewAwsInvoiceWithDefaults instantiates a new AwsInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AwsInvoice) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsInvoice) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AwsInvoice) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AwsInvoice) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmountDue

`func (o *AwsInvoice) GetAmountDue() float32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *AwsInvoice) GetAmountDueOk() (*float32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *AwsInvoice) SetAmountDue(v float32)`

SetAmountDue sets AmountDue field to given value.

### HasAmountDue

`func (o *AwsInvoice) HasAmountDue() bool`

HasAmountDue returns a boolean if a field has been set.

### GetBillingEntity

`func (o *AwsInvoice) GetBillingEntity() string`

GetBillingEntity returns the BillingEntity field if non-nil, zero value otherwise.

### GetBillingEntityOk

`func (o *AwsInvoice) GetBillingEntityOk() (*string, bool)`

GetBillingEntityOk returns a tuple with the BillingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEntity

`func (o *AwsInvoice) SetBillingEntity(v string)`

SetBillingEntity sets BillingEntity field to given value.

### HasBillingEntity

`func (o *AwsInvoice) HasBillingEntity() bool`

HasBillingEntity returns a boolean if a field has been set.

### GetBillingMonth

`func (o *AwsInvoice) GetBillingMonth() int32`

GetBillingMonth returns the BillingMonth field if non-nil, zero value otherwise.

### GetBillingMonthOk

`func (o *AwsInvoice) GetBillingMonthOk() (*int32, bool)`

GetBillingMonthOk returns a tuple with the BillingMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMonth

`func (o *AwsInvoice) SetBillingMonth(v int32)`

SetBillingMonth sets BillingMonth field to given value.

### HasBillingMonth

`func (o *AwsInvoice) HasBillingMonth() bool`

HasBillingMonth returns a boolean if a field has been set.

### GetBillingPeriodEnd

`func (o *AwsInvoice) GetBillingPeriodEnd() string`

GetBillingPeriodEnd returns the BillingPeriodEnd field if non-nil, zero value otherwise.

### GetBillingPeriodEndOk

`func (o *AwsInvoice) GetBillingPeriodEndOk() (*string, bool)`

GetBillingPeriodEndOk returns a tuple with the BillingPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodEnd

`func (o *AwsInvoice) SetBillingPeriodEnd(v string)`

SetBillingPeriodEnd sets BillingPeriodEnd field to given value.

### HasBillingPeriodEnd

`func (o *AwsInvoice) HasBillingPeriodEnd() bool`

HasBillingPeriodEnd returns a boolean if a field has been set.

### GetBillingPeriodStart

`func (o *AwsInvoice) GetBillingPeriodStart() string`

GetBillingPeriodStart returns the BillingPeriodStart field if non-nil, zero value otherwise.

### GetBillingPeriodStartOk

`func (o *AwsInvoice) GetBillingPeriodStartOk() (*string, bool)`

GetBillingPeriodStartOk returns a tuple with the BillingPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriodStart

`func (o *AwsInvoice) SetBillingPeriodStart(v string)`

SetBillingPeriodStart sets BillingPeriodStart field to given value.

### HasBillingPeriodStart

`func (o *AwsInvoice) HasBillingPeriodStart() bool`

HasBillingPeriodStart returns a boolean if a field has been set.

### GetBillingYear

`func (o *AwsInvoice) GetBillingYear() int32`

GetBillingYear returns the BillingYear field if non-nil, zero value otherwise.

### GetBillingYearOk

`func (o *AwsInvoice) GetBillingYearOk() (*int32, bool)`

GetBillingYearOk returns a tuple with the BillingYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingYear

`func (o *AwsInvoice) SetBillingYear(v int32)`

SetBillingYear sets BillingYear field to given value.

### HasBillingYear

`func (o *AwsInvoice) HasBillingYear() bool`

HasBillingYear returns a boolean if a field has been set.

### GetChargeType

`func (o *AwsInvoice) GetChargeType() string`

GetChargeType returns the ChargeType field if non-nil, zero value otherwise.

### GetChargeTypeOk

`func (o *AwsInvoice) GetChargeTypeOk() (*string, bool)`

GetChargeTypeOk returns a tuple with the ChargeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeType

`func (o *AwsInvoice) SetChargeType(v string)`

SetChargeType sets ChargeType field to given value.

### HasChargeType

`func (o *AwsInvoice) HasChargeType() bool`

HasChargeType returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *AwsInvoice) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AwsInvoice) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AwsInvoice) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AwsInvoice) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDueDate

`func (o *AwsInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *AwsInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *AwsInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *AwsInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetId

`func (o *AwsInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoiceType

`func (o *AwsInvoice) GetInvoiceType() string`

GetInvoiceType returns the InvoiceType field if non-nil, zero value otherwise.

### GetInvoiceTypeOk

`func (o *AwsInvoice) GetInvoiceTypeOk() (*string, bool)`

GetInvoiceTypeOk returns a tuple with the InvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceType

`func (o *AwsInvoice) SetInvoiceType(v string)`

SetInvoiceType sets InvoiceType field to given value.

### HasInvoiceType

`func (o *AwsInvoice) HasInvoiceType() bool`

HasInvoiceType returns a boolean if a field has been set.

### GetIssuedDate

`func (o *AwsInvoice) GetIssuedDate() string`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *AwsInvoice) GetIssuedDateOk() (*string, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *AwsInvoice) SetIssuedDate(v string)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *AwsInvoice) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### GetLineItems

`func (o *AwsInvoice) GetLineItems() AwsInvoiceLineItems`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *AwsInvoice) GetLineItemsOk() (*AwsInvoiceLineItems, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *AwsInvoice) SetLineItems(v AwsInvoiceLineItems)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *AwsInvoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetPdfDownloaded

`func (o *AwsInvoice) GetPdfDownloaded() bool`

GetPdfDownloaded returns the PdfDownloaded field if non-nil, zero value otherwise.

### GetPdfDownloadedOk

`func (o *AwsInvoice) GetPdfDownloadedOk() (*bool, bool)`

GetPdfDownloadedOk returns a tuple with the PdfDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdfDownloaded

`func (o *AwsInvoice) SetPdfDownloaded(v bool)`

SetPdfDownloaded sets PdfDownloaded field to given value.

### HasPdfDownloaded

`func (o *AwsInvoice) HasPdfDownloaded() bool`

HasPdfDownloaded returns a boolean if a field has been set.

### GetSettleDate

`func (o *AwsInvoice) GetSettleDate() string`

GetSettleDate returns the SettleDate field if non-nil, zero value otherwise.

### GetSettleDateOk

`func (o *AwsInvoice) GetSettleDateOk() (*string, bool)`

GetSettleDateOk returns a tuple with the SettleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettleDate

`func (o *AwsInvoice) SetSettleDate(v string)`

SetSettleDate sets SettleDate field to given value.

### HasSettleDate

`func (o *AwsInvoice) HasSettleDate() bool`

HasSettleDate returns a boolean if a field has been set.

### GetStatus

`func (o *AwsInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AwsInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AwsInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AwsInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncTimestamp

`func (o *AwsInvoice) GetSyncTimestamp() string`

GetSyncTimestamp returns the SyncTimestamp field if non-nil, zero value otherwise.

### GetSyncTimestampOk

`func (o *AwsInvoice) GetSyncTimestampOk() (*string, bool)`

GetSyncTimestampOk returns a tuple with the SyncTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTimestamp

`func (o *AwsInvoice) SetSyncTimestamp(v string)`

SetSyncTimestamp sets SyncTimestamp field to given value.

### HasSyncTimestamp

`func (o *AwsInvoice) HasSyncTimestamp() bool`

HasSyncTimestamp returns a boolean if a field has been set.

### GetSystemIdentifierArn

`func (o *AwsInvoice) GetSystemIdentifierArn() string`

GetSystemIdentifierArn returns the SystemIdentifierArn field if non-nil, zero value otherwise.

### GetSystemIdentifierArnOk

`func (o *AwsInvoice) GetSystemIdentifierArnOk() (*string, bool)`

GetSystemIdentifierArnOk returns a tuple with the SystemIdentifierArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemIdentifierArn

`func (o *AwsInvoice) SetSystemIdentifierArn(v string)`

SetSystemIdentifierArn sets SystemIdentifierArn field to given value.

### HasSystemIdentifierArn

`func (o *AwsInvoice) HasSystemIdentifierArn() bool`

HasSystemIdentifierArn returns a boolean if a field has been set.

### GetTaxAmount

`func (o *AwsInvoice) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *AwsInvoice) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *AwsInvoice) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *AwsInvoice) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *AwsInvoice) GetTotalAmount() float32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *AwsInvoice) GetTotalAmountOk() (*float32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *AwsInvoice) SetTotalAmount(v float32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *AwsInvoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetTotalAmountBeforeTax

`func (o *AwsInvoice) GetTotalAmountBeforeTax() float32`

GetTotalAmountBeforeTax returns the TotalAmountBeforeTax field if non-nil, zero value otherwise.

### GetTotalAmountBeforeTaxOk

`func (o *AwsInvoice) GetTotalAmountBeforeTaxOk() (*float32, bool)`

GetTotalAmountBeforeTaxOk returns a tuple with the TotalAmountBeforeTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountBeforeTax

`func (o *AwsInvoice) SetTotalAmountBeforeTax(v float32)`

SetTotalAmountBeforeTax sets TotalAmountBeforeTax field to given value.

### HasTotalAmountBeforeTax

`func (o *AwsInvoice) HasTotalAmountBeforeTax() bool`

HasTotalAmountBeforeTax returns a boolean if a field has been set.

### GetTransactionStatus

`func (o *AwsInvoice) GetTransactionStatus() string`

GetTransactionStatus returns the TransactionStatus field if non-nil, zero value otherwise.

### GetTransactionStatusOk

`func (o *AwsInvoice) GetTransactionStatusOk() (*string, bool)`

GetTransactionStatusOk returns a tuple with the TransactionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionStatus

`func (o *AwsInvoice) SetTransactionStatus(v string)`

SetTransactionStatus sets TransactionStatus field to given value.

### HasTransactionStatus

`func (o *AwsInvoice) HasTransactionStatus() bool`

HasTransactionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


