# StripeBalanceTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | Gross amount of this transaction (in cents (or local equivalent)). A positive value represents funds charged to another party, and a negative value represents funds sent to another party. | [optional] 
**AvailableOn** | Pointer to **int32** | The date that the transaction&#39;s net funds become available in the Stripe balance. | [optional] 
**ChargeId** | Pointer to **string** | ID of the charge which the balance transaction comes from. | [optional] 
**Created** | Pointer to **int32** | Time at which the object was created. Measured in seconds since the Unix epoch. | [optional] 
**Description** | Pointer to **string** | An arbitrary string attached to the object. Often useful for displaying to users. | [optional] 
**ExchangeRate** | Pointer to **float32** | If applicable, this transaction uses an exchange rate. If money converts from currency A to currency B, then the &#x60;amount&#x60; in currency A, multipled by the &#x60;exchange_rate&#x60;, equals the &#x60;amount&#x60; in currency B. For example, if you charge a customer 10.00 EUR, the PaymentIntent&#39;s &#x60;amount&#x60; is &#x60;1000&#x60; and &#x60;currency&#x60; is &#x60;eur&#x60;. If this converts to 12.34 USD in your Stripe account, the BalanceTransaction&#39;s &#x60;amount&#x60; is &#x60;1234&#x60;, its &#x60;currency&#x60; is &#x60;usd&#x60;, and the &#x60;exchange_rate&#x60; is &#x60;1.234&#x60;. | [optional] 
**Fee** | Pointer to **int32** | Fees (in cents (or local equivalent)) paid for this transaction. Represented as a positive integer when assessed. | [optional] 
**FeeDetails** | Pointer to [**[]StripeBalanceTransactionFeeDetail**](StripeBalanceTransactionFeeDetail.md) | Detailed breakdown of fees (in cents (or local equivalent)) paid for this transaction. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**Net** | Pointer to **int32** | Net impact to a Stripe balance (in cents (or local equivalent)). A positive value represents incrementing a Stripe balance, and a negative value decrementing a Stripe balance. You can calculate the net impact of a transaction on a balance by &#x60;amount&#x60; - &#x60;fee&#x60; | [optional] 
**Status** | Pointer to **string** | The transaction&#39;s net funds status in the Stripe balance, which are either &#x60;available&#x60; or &#x60;pending&#x60;. | [optional] 
**Type** | Pointer to **string** | Transaction type: &#x60;adjustment&#x60;, &#x60;advance&#x60;, &#x60;advance_funding&#x60;, &#x60;anticipation_repayment&#x60;, &#x60;application_fee&#x60;, &#x60;application_fee_refund&#x60;, &#x60;charge&#x60;, &#x60;climate_order_purchase&#x60;, &#x60;climate_order_refund&#x60;, &#x60;connect_collection_transfer&#x60;, &#x60;contribution&#x60;, &#x60;issuing_authorization_hold&#x60;, &#x60;issuing_authorization_release&#x60;, &#x60;issuing_dispute&#x60;, &#x60;issuing_transaction&#x60;, &#x60;obligation_outbound&#x60;, &#x60;obligation_reversal_inbound&#x60;, &#x60;payment&#x60;, &#x60;payment_failure_refund&#x60;, &#x60;payment_network_reserve_hold&#x60;, &#x60;payment_network_reserve_release&#x60;, &#x60;payment_refund&#x60;, &#x60;payment_reversal&#x60;, &#x60;payment_unreconciled&#x60;, &#x60;payout&#x60;, &#x60;payout_cancel&#x60;, &#x60;payout_failure&#x60;, &#x60;refund&#x60;, &#x60;refund_failure&#x60;, &#x60;reserve_transaction&#x60;, &#x60;reserved_funds&#x60;, &#x60;stripe_fee&#x60;, &#x60;stripe_fx_fee&#x60;, &#x60;tax_fee&#x60;, &#x60;topup&#x60;, &#x60;topup_reversal&#x60;, &#x60;transfer&#x60;, &#x60;transfer_cancel&#x60;, &#x60;transfer_failure&#x60;, or &#x60;transfer_refund&#x60;. Learn more about [balance transaction types and what they represent](https://stripe.com/docs/reports/balance-transaction-types). To classify transactions for accounting purposes, consider &#x60;reporting_category&#x60; instead. | [optional] 

## Methods

### NewStripeBalanceTransaction

`func NewStripeBalanceTransaction() *StripeBalanceTransaction`

NewStripeBalanceTransaction instantiates a new StripeBalanceTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeBalanceTransactionWithDefaults

`func NewStripeBalanceTransactionWithDefaults() *StripeBalanceTransaction`

NewStripeBalanceTransactionWithDefaults instantiates a new StripeBalanceTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *StripeBalanceTransaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *StripeBalanceTransaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *StripeBalanceTransaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *StripeBalanceTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetAvailableOn

`func (o *StripeBalanceTransaction) GetAvailableOn() int32`

GetAvailableOn returns the AvailableOn field if non-nil, zero value otherwise.

### GetAvailableOnOk

`func (o *StripeBalanceTransaction) GetAvailableOnOk() (*int32, bool)`

GetAvailableOnOk returns a tuple with the AvailableOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableOn

`func (o *StripeBalanceTransaction) SetAvailableOn(v int32)`

SetAvailableOn sets AvailableOn field to given value.

### HasAvailableOn

`func (o *StripeBalanceTransaction) HasAvailableOn() bool`

HasAvailableOn returns a boolean if a field has been set.

### GetChargeId

`func (o *StripeBalanceTransaction) GetChargeId() string`

GetChargeId returns the ChargeId field if non-nil, zero value otherwise.

### GetChargeIdOk

`func (o *StripeBalanceTransaction) GetChargeIdOk() (*string, bool)`

GetChargeIdOk returns a tuple with the ChargeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeId

`func (o *StripeBalanceTransaction) SetChargeId(v string)`

SetChargeId sets ChargeId field to given value.

### HasChargeId

`func (o *StripeBalanceTransaction) HasChargeId() bool`

HasChargeId returns a boolean if a field has been set.

### GetCreated

`func (o *StripeBalanceTransaction) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StripeBalanceTransaction) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StripeBalanceTransaction) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StripeBalanceTransaction) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *StripeBalanceTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StripeBalanceTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StripeBalanceTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StripeBalanceTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExchangeRate

`func (o *StripeBalanceTransaction) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *StripeBalanceTransaction) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *StripeBalanceTransaction) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *StripeBalanceTransaction) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetFee

`func (o *StripeBalanceTransaction) GetFee() int32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *StripeBalanceTransaction) GetFeeOk() (*int32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *StripeBalanceTransaction) SetFee(v int32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *StripeBalanceTransaction) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetFeeDetails

`func (o *StripeBalanceTransaction) GetFeeDetails() []StripeBalanceTransactionFeeDetail`

GetFeeDetails returns the FeeDetails field if non-nil, zero value otherwise.

### GetFeeDetailsOk

`func (o *StripeBalanceTransaction) GetFeeDetailsOk() (*[]StripeBalanceTransactionFeeDetail, bool)`

GetFeeDetailsOk returns a tuple with the FeeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeDetails

`func (o *StripeBalanceTransaction) SetFeeDetails(v []StripeBalanceTransactionFeeDetail)`

SetFeeDetails sets FeeDetails field to given value.

### HasFeeDetails

`func (o *StripeBalanceTransaction) HasFeeDetails() bool`

HasFeeDetails returns a boolean if a field has been set.

### GetId

`func (o *StripeBalanceTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeBalanceTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeBalanceTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeBalanceTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNet

`func (o *StripeBalanceTransaction) GetNet() int32`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *StripeBalanceTransaction) GetNetOk() (*int32, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *StripeBalanceTransaction) SetNet(v int32)`

SetNet sets Net field to given value.

### HasNet

`func (o *StripeBalanceTransaction) HasNet() bool`

HasNet returns a boolean if a field has been set.

### GetStatus

`func (o *StripeBalanceTransaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StripeBalanceTransaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StripeBalanceTransaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StripeBalanceTransaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *StripeBalanceTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeBalanceTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeBalanceTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeBalanceTransaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


