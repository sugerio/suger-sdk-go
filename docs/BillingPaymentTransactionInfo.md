# BillingPaymentTransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceDate** | Pointer to **time.Time** | The invoice issue date. | [optional] 
**RefundExists** | Pointer to **bool** | Refund flag marks whether the transaction has any refund records. | [optional] 
**StripeBalanceTransaction** | Pointer to [**StripeBalanceTransaction**](StripeBalanceTransaction.md) | Balance transaction that describes the impact of this charge on your account balance. | [optional] 
**StripeDisputes** | Pointer to [**[]StripeDispute**](StripeDispute.md) | Stripe dispute result, got by Dispute API, there may be multiple disputes. | [optional] 
**StripeError** | Pointer to [**StripeError**](StripeError.md) | Error of stripe API call | [optional] 
**StripePaymentIntent** | Pointer to [**StripePaymentIntent**](StripePaymentIntent.md) | Stripe payment intent result, got by PaymentIntent API | [optional] 
**StripeRefund** | Pointer to [**StripeRefund**](StripeRefund.md) | Stripe refund result, got by Refund API | [optional] 

## Methods

### NewBillingPaymentTransactionInfo

`func NewBillingPaymentTransactionInfo() *BillingPaymentTransactionInfo`

NewBillingPaymentTransactionInfo instantiates a new BillingPaymentTransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPaymentTransactionInfoWithDefaults

`func NewBillingPaymentTransactionInfoWithDefaults() *BillingPaymentTransactionInfo`

NewBillingPaymentTransactionInfoWithDefaults instantiates a new BillingPaymentTransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceDate

`func (o *BillingPaymentTransactionInfo) GetInvoiceDate() time.Time`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *BillingPaymentTransactionInfo) GetInvoiceDateOk() (*time.Time, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *BillingPaymentTransactionInfo) SetInvoiceDate(v time.Time)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *BillingPaymentTransactionInfo) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetRefundExists

`func (o *BillingPaymentTransactionInfo) GetRefundExists() bool`

GetRefundExists returns the RefundExists field if non-nil, zero value otherwise.

### GetRefundExistsOk

`func (o *BillingPaymentTransactionInfo) GetRefundExistsOk() (*bool, bool)`

GetRefundExistsOk returns a tuple with the RefundExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundExists

`func (o *BillingPaymentTransactionInfo) SetRefundExists(v bool)`

SetRefundExists sets RefundExists field to given value.

### HasRefundExists

`func (o *BillingPaymentTransactionInfo) HasRefundExists() bool`

HasRefundExists returns a boolean if a field has been set.

### GetStripeBalanceTransaction

`func (o *BillingPaymentTransactionInfo) GetStripeBalanceTransaction() StripeBalanceTransaction`

GetStripeBalanceTransaction returns the StripeBalanceTransaction field if non-nil, zero value otherwise.

### GetStripeBalanceTransactionOk

`func (o *BillingPaymentTransactionInfo) GetStripeBalanceTransactionOk() (*StripeBalanceTransaction, bool)`

GetStripeBalanceTransactionOk returns a tuple with the StripeBalanceTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeBalanceTransaction

`func (o *BillingPaymentTransactionInfo) SetStripeBalanceTransaction(v StripeBalanceTransaction)`

SetStripeBalanceTransaction sets StripeBalanceTransaction field to given value.

### HasStripeBalanceTransaction

`func (o *BillingPaymentTransactionInfo) HasStripeBalanceTransaction() bool`

HasStripeBalanceTransaction returns a boolean if a field has been set.

### GetStripeDisputes

`func (o *BillingPaymentTransactionInfo) GetStripeDisputes() []StripeDispute`

GetStripeDisputes returns the StripeDisputes field if non-nil, zero value otherwise.

### GetStripeDisputesOk

`func (o *BillingPaymentTransactionInfo) GetStripeDisputesOk() (*[]StripeDispute, bool)`

GetStripeDisputesOk returns a tuple with the StripeDisputes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeDisputes

`func (o *BillingPaymentTransactionInfo) SetStripeDisputes(v []StripeDispute)`

SetStripeDisputes sets StripeDisputes field to given value.

### HasStripeDisputes

`func (o *BillingPaymentTransactionInfo) HasStripeDisputes() bool`

HasStripeDisputes returns a boolean if a field has been set.

### GetStripeError

`func (o *BillingPaymentTransactionInfo) GetStripeError() StripeError`

GetStripeError returns the StripeError field if non-nil, zero value otherwise.

### GetStripeErrorOk

`func (o *BillingPaymentTransactionInfo) GetStripeErrorOk() (*StripeError, bool)`

GetStripeErrorOk returns a tuple with the StripeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeError

`func (o *BillingPaymentTransactionInfo) SetStripeError(v StripeError)`

SetStripeError sets StripeError field to given value.

### HasStripeError

`func (o *BillingPaymentTransactionInfo) HasStripeError() bool`

HasStripeError returns a boolean if a field has been set.

### GetStripePaymentIntent

`func (o *BillingPaymentTransactionInfo) GetStripePaymentIntent() StripePaymentIntent`

GetStripePaymentIntent returns the StripePaymentIntent field if non-nil, zero value otherwise.

### GetStripePaymentIntentOk

`func (o *BillingPaymentTransactionInfo) GetStripePaymentIntentOk() (*StripePaymentIntent, bool)`

GetStripePaymentIntentOk returns a tuple with the StripePaymentIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePaymentIntent

`func (o *BillingPaymentTransactionInfo) SetStripePaymentIntent(v StripePaymentIntent)`

SetStripePaymentIntent sets StripePaymentIntent field to given value.

### HasStripePaymentIntent

`func (o *BillingPaymentTransactionInfo) HasStripePaymentIntent() bool`

HasStripePaymentIntent returns a boolean if a field has been set.

### GetStripeRefund

`func (o *BillingPaymentTransactionInfo) GetStripeRefund() StripeRefund`

GetStripeRefund returns the StripeRefund field if non-nil, zero value otherwise.

### GetStripeRefundOk

`func (o *BillingPaymentTransactionInfo) GetStripeRefundOk() (*StripeRefund, bool)`

GetStripeRefundOk returns a tuple with the StripeRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeRefund

`func (o *BillingPaymentTransactionInfo) SetStripeRefund(v StripeRefund)`

SetStripeRefund sets StripeRefund field to given value.

### HasStripeRefund

`func (o *BillingPaymentTransactionInfo) HasStripeRefund() bool`

HasStripeRefund returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


