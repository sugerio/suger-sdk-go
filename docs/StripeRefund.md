# StripeRefund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeId** | Pointer to **string** | ID of the charge that&#39;s refunded. | [optional] 
**EstinationDetails** | Pointer to [**StripeRefundDestinationDetails**](StripeRefundDestinationDetails.md) | Transaction-specific details for the refund. | [optional] 
**FailureBalanceTransactionId** | Pointer to **string** | After the refund fails, this balance transaction describes the adjustment made on your account balance that reverses the initial balance transaction. | [optional] 
**FailureReason** | Pointer to **string** | Provides the reason for the refund failure. Possible values are: &#x60;lost_or_stolen_card&#x60;, &#x60;expired_or_canceled_card&#x60;, &#x60;charge_for_pending_refund_disputed&#x60;, &#x60;insufficient_funds&#x60;, &#x60;declined&#x60;, &#x60;merchant_request&#x60;, or &#x60;unknown&#x60;. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PaymentIntentId** | Pointer to **string** | ID of the PaymentIntent that&#39;s refunded. | [optional] 
**Status** | Pointer to [**StripeRefundStatus**](StripeRefundStatus.md) | Status of the refund. This can be &#x60;pending&#x60;, &#x60;requires_action&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60;, or &#x60;canceled&#x60;. Learn more about [failed refunds](https://stripe.com/docs/refunds#failed-refunds). | [optional] 

## Methods

### NewStripeRefund

`func NewStripeRefund() *StripeRefund`

NewStripeRefund instantiates a new StripeRefund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeRefundWithDefaults

`func NewStripeRefundWithDefaults() *StripeRefund`

NewStripeRefundWithDefaults instantiates a new StripeRefund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChargeId

`func (o *StripeRefund) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *StripeRefund) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *StripeRefund) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *StripeRefund) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetEstinationDetails

`func (o *StripeRefund) GetEstinationDetails() StripeRefundDestinationDetails`

GetEstinationDetails returns the EstinationDetails field if non-nil, zero value otherwise.

### GetEstinationDetailsOk

`func (o *StripeRefund) GetEstinationDetailsOk() (*StripeRefundDestinationDetails, bool)`

GetEstinationDetailsOk returns a tuple with the EstinationDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstinationDetails

`func (o *StripeRefund) SetEstinationDetails(v StripeRefundDestinationDetails)`

SetEstinationDetails sets EstinationDetails field to given value.

### HasEstinationDetails

`func (o *StripeRefund) HasEstinationDetails() bool`

HasEstinationDetails returns a boolean if a field has been set.

### GetFailureBalanceTransactionId

`func (o *StripeRefund) GetFailureBalanceTransactionId() string`

GetFailureBalanceTransactionId returns the FailureBalanceTransactionId field if non-nil, zero value otherwise.

### GetFailureBalanceTransactionIdOk

`func (o *StripeRefund) GetFailureBalanceTransactionIdOk() (*string, bool)`

GetFailureBalanceTransactionIdOk returns a tuple with the FailureBalanceTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureBalanceTransactionId

`func (o *StripeRefund) SetFailureBalanceTransactionId(v string)`

SetFailureBalanceTransactionId sets FailureBalanceTransactionId field to given value.

### HasFailureBalanceTransactionId

`func (o *StripeRefund) HasFailureBalanceTransactionId() bool`

HasFailureBalanceTransactionId returns a boolean if a field has been set.

### GetFailureReason

`func (o *StripeRefund) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *StripeRefund) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *StripeRefund) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *StripeRefund) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetId

`func (o *StripeRefund) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeRefund) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeRefund) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeRefund) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentIntentId

`func (o *StripeRefund) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *StripeRefund) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *StripeRefund) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *StripeRefund) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### GetStatus

`func (o *StripeRefund) GetStatus() StripeRefundStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StripeRefund) GetStatusOk() (*StripeRefundStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StripeRefund) SetStatus(v StripeRefundStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StripeRefund) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


