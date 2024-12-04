# StripeRefundDestinationDetailsUSBankTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **string** | The reference assigned to the refund. | [optional] 
**ReferenceStatus** | Pointer to **string** | Status of the reference on the refund. This can be &#x60;pending&#x60;, &#x60;available&#x60; or &#x60;unavailable&#x60;. | [optional] 

## Methods

### NewStripeRefundDestinationDetailsUSBankTransfer

`func NewStripeRefundDestinationDetailsUSBankTransfer() *StripeRefundDestinationDetailsUSBankTransfer`

NewStripeRefundDestinationDetailsUSBankTransfer instantiates a new StripeRefundDestinationDetailsUSBankTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeRefundDestinationDetailsUSBankTransferWithDefaults

`func NewStripeRefundDestinationDetailsUSBankTransferWithDefaults() *StripeRefundDestinationDetailsUSBankTransfer`

NewStripeRefundDestinationDetailsUSBankTransferWithDefaults instantiates a new StripeRefundDestinationDetailsUSBankTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StripeRefundDestinationDetailsUSBankTransfer) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StripeRefundDestinationDetailsUSBankTransfer) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceStatus

`func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReferenceStatus() string`

GetReferenceStatus returns the ReferenceStatus field if non-nil, zero value otherwise.

### GetReferenceStatusOk

`func (o *StripeRefundDestinationDetailsUSBankTransfer) GetReferenceStatusOk() (*string, bool)`

GetReferenceStatusOk returns a tuple with the ReferenceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStatus

`func (o *StripeRefundDestinationDetailsUSBankTransfer) SetReferenceStatus(v string)`

SetReferenceStatus sets ReferenceStatus field to given value.

### HasReferenceStatus

`func (o *StripeRefundDestinationDetailsUSBankTransfer) HasReferenceStatus() bool`

HasReferenceStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


