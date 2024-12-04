# StripePaymentIntent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the object. | [optional] 
**LastPaymentError** | Pointer to [**StripeError**](StripeError.md) | The payment error encountered in the previous PaymentIntent confirmation. It will be cleared if the PaymentIntent is later updated for any reason. | [optional] 
**Livemode** | Pointer to **bool** | Has the value &#x60;true&#x60; if the object exists in live mode or the value &#x60;false&#x60; if the object exists in test mode. | [optional] 
**Status** | Pointer to [**StripePaymentIntentStatus**](StripePaymentIntentStatus.md) | Status of this PaymentIntent. Read more about each PaymentIntent [status](https://stripe.com/docs/payments/intents#intent-statuses). | [optional] 

## Methods

### NewStripePaymentIntent

`func NewStripePaymentIntent() *StripePaymentIntent`

NewStripePaymentIntent instantiates a new StripePaymentIntent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentIntentWithDefaults

`func NewStripePaymentIntentWithDefaults() *StripePaymentIntent`

NewStripePaymentIntentWithDefaults instantiates a new StripePaymentIntent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StripePaymentIntent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripePaymentIntent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripePaymentIntent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripePaymentIntent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPaymentError

`func (o *StripePaymentIntent) GetLastPaymentError() StripeError`

GetLastPaymentError returns the LastPaymentError field if non-nil, zero value otherwise.

### GetLastPaymentErrorOk

`func (o *StripePaymentIntent) GetLastPaymentErrorOk() (*StripeError, bool)`

GetLastPaymentErrorOk returns a tuple with the LastPaymentError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentError

`func (o *StripePaymentIntent) SetLastPaymentError(v StripeError)`

SetLastPaymentError sets LastPaymentError field to given value.

### HasLastPaymentError

`func (o *StripePaymentIntent) HasLastPaymentError() bool`

HasLastPaymentError returns a boolean if a field has been set.

### GetLivemode

`func (o *StripePaymentIntent) GetLivemode() bool`

GetLivemode returns the Livemode field if non-nil, zero value otherwise.

### GetLivemodeOk

`func (o *StripePaymentIntent) GetLivemodeOk() (*bool, bool)`

GetLivemodeOk returns a tuple with the Livemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivemode

`func (o *StripePaymentIntent) SetLivemode(v bool)`

SetLivemode sets Livemode field to given value.

### HasLivemode

`func (o *StripePaymentIntent) HasLivemode() bool`

HasLivemode returns a boolean if a field has been set.

### GetStatus

`func (o *StripePaymentIntent) GetStatus() StripePaymentIntentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StripePaymentIntent) GetStatusOk() (*StripePaymentIntentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StripePaymentIntent) SetStatus(v StripePaymentIntentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StripePaymentIntent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


