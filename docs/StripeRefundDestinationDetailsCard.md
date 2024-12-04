# StripeRefundDestinationDetailsCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reference** | Pointer to **string** | Value of the reference number assigned to the refund. | [optional] 
**ReferenceStatus** | Pointer to **string** | Status of the reference number on the refund. This can be &#x60;pending&#x60;, &#x60;available&#x60; or &#x60;unavailable&#x60;. | [optional] 
**ReferenceType** | Pointer to **string** | Type of the reference number assigned to the refund. | [optional] 
**Type** | Pointer to **string** | The type of refund. This can be &#x60;refund&#x60;, &#x60;reversal&#x60;, or &#x60;pending&#x60;. | [optional] 

## Methods

### NewStripeRefundDestinationDetailsCard

`func NewStripeRefundDestinationDetailsCard() *StripeRefundDestinationDetailsCard`

NewStripeRefundDestinationDetailsCard instantiates a new StripeRefundDestinationDetailsCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeRefundDestinationDetailsCardWithDefaults

`func NewStripeRefundDestinationDetailsCardWithDefaults() *StripeRefundDestinationDetailsCard`

NewStripeRefundDestinationDetailsCardWithDefaults instantiates a new StripeRefundDestinationDetailsCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReference

`func (o *StripeRefundDestinationDetailsCard) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *StripeRefundDestinationDetailsCard) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *StripeRefundDestinationDetailsCard) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *StripeRefundDestinationDetailsCard) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetReferenceStatus

`func (o *StripeRefundDestinationDetailsCard) GetReferenceStatus() string`

GetReferenceStatus returns the ReferenceStatus field if non-nil, zero value otherwise.

### GetReferenceStatusOk

`func (o *StripeRefundDestinationDetailsCard) GetReferenceStatusOk() (*string, bool)`

GetReferenceStatusOk returns a tuple with the ReferenceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStatus

`func (o *StripeRefundDestinationDetailsCard) SetReferenceStatus(v string)`

SetReferenceStatus sets ReferenceStatus field to given value.

### HasReferenceStatus

`func (o *StripeRefundDestinationDetailsCard) HasReferenceStatus() bool`

HasReferenceStatus returns a boolean if a field has been set.

### GetReferenceType

`func (o *StripeRefundDestinationDetailsCard) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *StripeRefundDestinationDetailsCard) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *StripeRefundDestinationDetailsCard) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.

### HasReferenceType

`func (o *StripeRefundDestinationDetailsCard) HasReferenceType() bool`

HasReferenceType returns a boolean if a field has been set.

### GetType

`func (o *StripeRefundDestinationDetailsCard) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StripeRefundDestinationDetailsCard) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StripeRefundDestinationDetailsCard) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StripeRefundDestinationDetailsCard) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


