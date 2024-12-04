# StripePaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BacsDebit** | Pointer to [**StripePaymentMethodBACSDebit**](StripePaymentMethodBACSDebit.md) |  | [optional] 
**Card** | Pointer to [**StripePaymentMethodCard**](StripePaymentMethodCard.md) |  | [optional] 
**Created** | Pointer to **int32** | Time at which the object was created. Measured in seconds since the Unix epoch. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the payment method on stripe. | [optional] 
**Livemode** | Pointer to **bool** | Has the value &#x60;true&#x60; if the object exists in live mode or the value &#x60;false&#x60; if the object exists in test mode. | [optional] 
**Object** | Pointer to **string** | String representing the object’s type. Always has the value &#x60;payment_method&#x60;. | [optional] 
**SepaDebit** | Pointer to [**StripePaymentMethodSEPADebit**](StripePaymentMethodSEPADebit.md) |  | [optional] 
**UsBankAccount** | Pointer to [**StripePaymentMethodUSBankAccount**](StripePaymentMethodUSBankAccount.md) |  | [optional] 

## Methods

### NewStripePaymentMethod

`func NewStripePaymentMethod() *StripePaymentMethod`

NewStripePaymentMethod instantiates a new StripePaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentMethodWithDefaults

`func NewStripePaymentMethodWithDefaults() *StripePaymentMethod`

NewStripePaymentMethodWithDefaults instantiates a new StripePaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBacsDebit

`func (o *StripePaymentMethod) GetBacsDebit() StripePaymentMethodBACSDebit`

GetBacsDebit returns the BacsDebit field if non-nil, zero value otherwise.

### GetBacsDebitOk

`func (o *StripePaymentMethod) GetBacsDebitOk() (*StripePaymentMethodBACSDebit, bool)`

GetBacsDebitOk returns a tuple with the BacsDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacsDebit

`func (o *StripePaymentMethod) SetBacsDebit(v StripePaymentMethodBACSDebit)`

SetBacsDebit sets BacsDebit field to given value.

### HasBacsDebit

`func (o *StripePaymentMethod) HasBacsDebit() bool`

HasBacsDebit returns a boolean if a field has been set.

### GetCard

`func (o *StripePaymentMethod) GetCard() StripePaymentMethodCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *StripePaymentMethod) GetCardOk() (*StripePaymentMethodCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *StripePaymentMethod) SetCard(v StripePaymentMethodCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *StripePaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetCreated

`func (o *StripePaymentMethod) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StripePaymentMethod) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StripePaymentMethod) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StripePaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *StripePaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripePaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripePaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripePaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLivemode

`func (o *StripePaymentMethod) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *StripePaymentMethod) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *StripePaymentMethod) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *StripePaymentMethod) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetObject

`func (o *StripePaymentMethod) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *StripePaymentMethod) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *StripePaymentMethod) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *StripePaymentMethod) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSepaDebit

`func (o *StripePaymentMethod) GetSepaDebit() StripePaymentMethodSEPADebit`

GetSepaDebit returns the SepaDebit field if non-nil, zero value otherwise.

### GetSepaDebitOk

`func (o *StripePaymentMethod) GetSepaDebitOk() (*StripePaymentMethodSEPADebit, bool)`

GetSepaDebitOk returns a tuple with the SepaDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaDebit

`func (o *StripePaymentMethod) SetSepaDebit(v StripePaymentMethodSEPADebit)`

SetSepaDebit sets SepaDebit field to given value.

### HasSepaDebit

`func (o *StripePaymentMethod) HasSepaDebit() bool`

HasSepaDebit returns a boolean if a field has been set.

### GetUsBankAccount

`func (o *StripePaymentMethod) GetUsBankAccount() StripePaymentMethodUSBankAccount`

GetUsBankAccount returns the UsBankAccount field if non-nil, zero value otherwise.

### GetUsBankAccountOk

`func (o *StripePaymentMethod) GetUsBankAccountOk() (*StripePaymentMethodUSBankAccount, bool)`

GetUsBankAccountOk returns a tuple with the UsBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsBankAccount

`func (o *StripePaymentMethod) SetUsBankAccount(v StripePaymentMethodUSBankAccount)`

SetUsBankAccount sets UsBankAccount field to given value.

### HasUsBankAccount

`func (o *StripePaymentMethod) HasUsBankAccount() bool`

HasUsBankAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


