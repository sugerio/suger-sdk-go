# PaymentInstallment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The amount the buyer has paid for this installment. If there is a discount off the original price, the amount is the discounted price. | [optional] 
**ChargeOn** | Pointer to **time.Time** | When the buyer will be charged for this installment. If it is null, the buyer will be charged on the effective date of the entitlement. | [optional] 
**ChargeOnStr** | Pointer to **string** | The charge on date in string format. It is used for front-end display only. | [optional] 
**Credit** | Pointer to **float32** | Used in GCP Marketplace private offer as one-time credit. Default as zero if there is no credit. | [optional] 
**DiscountPercentage** | Pointer to **float32** | The discount percentage off the original price. For GCP Marketplace, it can be discount off the commitment amount or discount off the usage price. The value is between 0 to 100. For example, 20 means 20% off. Default as zero if there is no discount. | [optional] 
**OriginalAmount** | Pointer to **float32** | The original amount before discount if there is a discount off the original price. nil if there is no discount. | [optional] 

## Methods

### NewPaymentInstallment

`func NewPaymentInstallment() *PaymentInstallment`

NewPaymentInstallment instantiates a new PaymentInstallment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInstallmentWithDefaults

`func NewPaymentInstallmentWithDefaults() *PaymentInstallment`

NewPaymentInstallmentWithDefaults instantiates a new PaymentInstallment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PaymentInstallment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInstallment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInstallment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentInstallment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChargeOn

`func (o *PaymentInstallment) GetChargeOn() time.Time`

GetChargeOn returns the ChargeOn field if non-nil, zero value otherwise.

### GetChargeOnOk

`func (o *PaymentInstallment) GetChargeOnOk() (*time.Time, bool)`

GetChargeOnOk returns a tuple with the ChargeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeOn

`func (o *PaymentInstallment) SetChargeOn(v time.Time)`

SetChargeOn sets ChargeOn field to given value.

### HasChargeOn

`func (o *PaymentInstallment) HasChargeOn() bool`

HasChargeOn returns a boolean if a field has been set.

### GetChargeOnStr

`func (o *PaymentInstallment) GetChargeOnStr() string`

GetChargeOnStr returns the ChargeOnStr field if non-nil, zero value otherwise.

### GetChargeOnStrOk

`func (o *PaymentInstallment) GetChargeOnStrOk() (*string, bool)`

GetChargeOnStrOk returns a tuple with the ChargeOnStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeOnStr

`func (o *PaymentInstallment) SetChargeOnStr(v string)`

SetChargeOnStr sets ChargeOnStr field to given value.

### HasChargeOnStr

`func (o *PaymentInstallment) HasChargeOnStr() bool`

HasChargeOnStr returns a boolean if a field has been set.

### GetCredit

`func (o *PaymentInstallment) GetCredit() float32`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *PaymentInstallment) GetCreditOk() (*float32, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *PaymentInstallment) SetCredit(v float32)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *PaymentInstallment) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *PaymentInstallment) GetDiscountPercentage() float32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *PaymentInstallment) GetDiscountPercentageOk() (*float32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *PaymentInstallment) SetDiscountPercentage(v float32)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *PaymentInstallment) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *PaymentInstallment) GetOriginalAmount() float32`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *PaymentInstallment) GetOriginalAmountOk() (*float32, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *PaymentInstallment) SetOriginalAmount(v float32)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *PaymentInstallment) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


