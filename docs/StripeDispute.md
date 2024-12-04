# StripeDispute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Disputed amount. Usually the amount of the charge, but it can differ (usually because of currency fluctuation or because only part of the order is disputed). | [optional] 
**ChargeId** | Pointer to **string** | ID of the charge that&#39;s disputed. | [optional] 
**Created** | Pointer to **int32** | Time at which the object was created. Measured in seconds since the Unix epoch. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**IsChargeRefundable** | Pointer to **bool** | If true, it&#39;s still possible to refund the disputed payment. After the payment has been fully refunded, no further funds are withdrawn from your Stripe account as a result of this dispute. | [optional] 
**Livemode** | Pointer to **bool** | Has the value &#x60;true&#x60; if the object exists in live mode or the value &#x60;false&#x60; if the object exists in test mode. | [optional] 
**PaymentIntentId** | Pointer to **string** | ID of the PaymentIntent that&#39;s disputed. | [optional] 
**Reason** | Pointer to **string** | Reason given by cardholder for dispute. Possible values are &#x60;bank_cannot_process&#x60;, &#x60;check_returned&#x60;, &#x60;credit_not_processed&#x60;, &#x60;customer_initiated&#x60;, &#x60;debit_not_authorized&#x60;, &#x60;duplicate&#x60;, &#x60;fraudulent&#x60;, &#x60;general&#x60;, &#x60;incorrect_account_details&#x60;, &#x60;insufficient_funds&#x60;, &#x60;product_not_received&#x60;, &#x60;product_unacceptable&#x60;, &#x60;subscription_canceled&#x60;, or &#x60;unrecognized&#x60;. Learn more about [dispute reasons](https://stripe.com/docs/disputes/categories). | [optional] 
**Status** | Pointer to **string** | Current status of dispute. Possible values are &#x60;warning_needs_response&#x60;, &#x60;warning_under_review&#x60;, &#x60;warning_closed&#x60;, &#x60;needs_response&#x60;, &#x60;under_review&#x60;, &#x60;won&#x60;, or &#x60;lost&#x60;. | [optional] 

## Methods

### NewStripeDispute

`func NewStripeDispute() *StripeDispute`

NewStripeDispute instantiates a new StripeDispute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeDisputeWithDefaults

`func NewStripeDisputeWithDefaults() *StripeDispute`

NewStripeDisputeWithDefaults instantiates a new StripeDispute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StripeDispute) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StripeDispute) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StripeDispute) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StripeDispute) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChargeId

`func (o *StripeDispute) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *StripeDispute) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *StripeDispute) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *StripeDispute) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetCreated

`func (o *StripeDispute) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StripeDispute) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StripeDispute) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StripeDispute) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *StripeDispute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeDispute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeDispute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeDispute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsChargeRefundable

`func (o *StripeDispute) GetIsChargeRefundable() bool`

GetIsChargeRefundable returns the IsChargeRefundable field if non-nil, zero value otherwise.

### GetIsChargeRefundableOk

`func (o *StripeDispute) GetIsChargeRefundableOk() (*bool, bool)`

GetIsChargeRefundableOk returns a tuple with the IsChargeRefundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsChargeRefundable

`func (o *StripeDispute) SetIsChargeRefundable(v bool)`

SetIsChargeRefundable sets IsChargeRefundable field to given value.

### HasIsChargeRefundable

`func (o *StripeDispute) HasIsChargeRefundable() bool`

HasIsChargeRefundable returns a boolean if a field has been set.

### GetLivemode

`func (o *StripeDispute) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *StripeDispute) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *StripeDispute) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *StripeDispute) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *StripeDispute) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *StripeDispute) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *StripeDispute) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *StripeDispute) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### GetReason

`func (o *StripeDispute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StripeDispute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StripeDispute) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *StripeDispute) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *StripeDispute) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StripeDispute) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StripeDispute) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StripeDispute) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


