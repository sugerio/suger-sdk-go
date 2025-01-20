# UpdateInvoiceInfoRequest

## Properties

 Name               | Type                                                         | Description                          | Notes      
--------------------|--------------------------------------------------------------|--------------------------------------|------------
 **DiscountAmount** | Pointer to **float32**                                       |                                      | [optional] 
 **DiscountType**   | Pointer to [**BillingDiscountType**](BillingDiscountType.md) | The overall discount of the invoice. | [optional] 
 **DueDate**        | Pointer to **string**                                        |                                      | [optional] 
 **Memo**           | Pointer to **string**                                        |                                      | [optional] 

## Methods

### NewUpdateInvoiceInfoRequest

`func NewUpdateInvoiceInfoRequest() *UpdateInvoiceInfoRequest`

NewUpdateInvoiceInfoRequest instantiates a new UpdateInvoiceInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInvoiceInfoRequestWithDefaults

`func NewUpdateInvoiceInfoRequestWithDefaults() *UpdateInvoiceInfoRequest`

NewUpdateInvoiceInfoRequestWithDefaults instantiates a new UpdateInvoiceInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountAmount

`func (o *UpdateInvoiceInfoRequest) GetDiscountAmount() float32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *UpdateInvoiceInfoRequest) GetDiscountAmountOk() (*float32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *UpdateInvoiceInfoRequest) SetDiscountAmount(v float32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *UpdateInvoiceInfoRequest) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetDiscountType

`func (o *UpdateInvoiceInfoRequest) GetDiscountType() BillingDiscountType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *UpdateInvoiceInfoRequest) GetDiscountTypeOk() (*BillingDiscountType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *UpdateInvoiceInfoRequest) SetDiscountType(v BillingDiscountType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *UpdateInvoiceInfoRequest) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetDueDate

`func (o *UpdateInvoiceInfoRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *UpdateInvoiceInfoRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *UpdateInvoiceInfoRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *UpdateInvoiceInfoRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetMemo

`func (o *UpdateInvoiceInfoRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *UpdateInvoiceInfoRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *UpdateInvoiceInfoRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *UpdateInvoiceInfoRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


