# BillingPaymentInstallmentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**ChargeOn** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingPaymentInstallmentDetail

`func NewBillingPaymentInstallmentDetail() *BillingPaymentInstallmentDetail`

NewBillingPaymentInstallmentDetail instantiates a new BillingPaymentInstallmentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingPaymentInstallmentDetailWithDefaults

`func NewBillingPaymentInstallmentDetailWithDefaults() *BillingPaymentInstallmentDetail`

NewBillingPaymentInstallmentDetailWithDefaults instantiates a new BillingPaymentInstallmentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingPaymentInstallmentDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingPaymentInstallmentDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingPaymentInstallmentDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingPaymentInstallmentDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChargeOn

`func (o *BillingPaymentInstallmentDetail) GetChargeOn() time.Time`

GetChargeOn returns the ChargeOn field if non-nil, zero value otherwise.

### GetChargeOnOk

`func (o *BillingPaymentInstallmentDetail) GetChargeOnOk() (*time.Time, bool)`

GetChargeOnOk returns a tuple with the ChargeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeOn

`func (o *BillingPaymentInstallmentDetail) SetChargeOn(v time.Time)`

SetChargeOn sets ChargeOn field to given value.

### HasChargeOn

`func (o *BillingPaymentInstallmentDetail) HasChargeOn() bool`

HasChargeOn returns a boolean if a field has been set.

### GetDescription

`func (o *BillingPaymentInstallmentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingPaymentInstallmentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingPaymentInstallmentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingPaymentInstallmentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


