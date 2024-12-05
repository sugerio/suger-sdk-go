# StripeCustomer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**StripeCustomerAddress**](StripeCustomerAddress.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The customer ID on the stripe platform. | [optional] 
**Metadata** | Pointer to **map[string]string** | Set of key-value pairs that you can attach to store additional information about customer. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 

## Methods

### NewStripeCustomer

`func NewStripeCustomer() *StripeCustomer`

NewStripeCustomer instantiates a new StripeCustomer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCustomerWithDefaults

`func NewStripeCustomerWithDefaults() *StripeCustomer`

NewStripeCustomerWithDefaults instantiates a new StripeCustomer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StripeCustomer) GetAddress() StripeCustomerAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StripeCustomer) GetAddressOk() (*StripeCustomerAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StripeCustomer) SetAddress(v StripeCustomerAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StripeCustomer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *StripeCustomer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StripeCustomer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StripeCustomer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StripeCustomer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *StripeCustomer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *StripeCustomer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *StripeCustomer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *StripeCustomer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *StripeCustomer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StripeCustomer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StripeCustomer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StripeCustomer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *StripeCustomer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StripeCustomer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StripeCustomer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StripeCustomer) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *StripeCustomer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StripeCustomer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StripeCustomer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StripeCustomer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *StripeCustomer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *StripeCustomer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *StripeCustomer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *StripeCustomer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


