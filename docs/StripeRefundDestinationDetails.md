# StripeRefundDestinationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**StripeRefundDestinationDetailsCard**](StripeRefundDestinationDetailsCard.md) |  | [optional] 
**Type** | Pointer to **string** | The type of transaction-specific details of the payment method used in the refund (e.g., &#x60;card&#x60;). An additional hash is included on &#x60;destination_details&#x60; with a name matching this value. It contains information specific to the refund transaction. | [optional] 
**UsBankTransfer** | Pointer to [**StripeRefundDestinationDetailsUSBankTransfer**](StripeRefundDestinationDetailsUSBankTransfer.md) |  | [optional] 

## Methods

### NewStripeRefundDestinationDetails

`func NewStripeRefundDestinationDetails() *StripeRefundDestinationDetails`

NewStripeRefundDestinationDetails instantiates a new StripeRefundDestinationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeRefundDestinationDetailsWithDefaults

`func NewStripeRefundDestinationDetailsWithDefaults() *StripeRefundDestinationDetails`

NewStripeRefundDestinationDetailsWithDefaults instantiates a new StripeRefundDestinationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *StripeRefundDestinationDetails) GetCard() StripeRefundDestinationDetailsCard`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *StripeRefundDestinationDetails) GetCardOk() (*StripeRefundDestinationDetailsCard, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *StripeRefundDestinationDetails) SetCard(v StripeRefundDestinationDetailsCard)`

SetCard sets Card field to given value.

### HasCard

`func (o *StripeRefundDestinationDetails) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetType

`func (o *StripeRefundDestinationDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeRefundDestinationDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeRefundDestinationDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeRefundDestinationDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsBankTransfer

`func (o *StripeRefundDestinationDetails) GetUsBankTransfer() StripeRefundDestinationDetailsUSBankTransfer`

GetUsBankTransfer returns the UsBankTransfer field if non-nil, zero value otherwise.

### GetUsBankTransferOk

`func (o *StripeRefundDestinationDetails) GetUsBankTransferOk() (*StripeRefundDestinationDetailsUSBankTransfer, bool)`

GetUsBankTransferOk returns a tuple with the UsBankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsBankTransfer

`func (o *StripeRefundDestinationDetails) SetUsBankTransfer(v StripeRefundDestinationDetailsUSBankTransfer)`

SetUsBankTransfer sets UsBankTransfer field to given value.

### HasUsBankTransfer

`func (o *StripeRefundDestinationDetails) HasUsBankTransfer() bool`

HasUsBankTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


