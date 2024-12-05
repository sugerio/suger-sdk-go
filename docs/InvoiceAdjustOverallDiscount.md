# InvoiceAdjustOverallDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discount** | Pointer to [**BillingDiscount**](BillingDiscount.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceAdjustOverallDiscount

`func NewInvoiceAdjustOverallDiscount() *InvoiceAdjustOverallDiscount`

NewInvoiceAdjustOverallDiscount instantiates a new InvoiceAdjustOverallDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAdjustOverallDiscountWithDefaults

`func NewInvoiceAdjustOverallDiscountWithDefaults() *InvoiceAdjustOverallDiscount`

NewInvoiceAdjustOverallDiscountWithDefaults instantiates a new InvoiceAdjustOverallDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscount

`func (o *InvoiceAdjustOverallDiscount) GetDiscount() BillingDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InvoiceAdjustOverallDiscount) GetDiscountOk() (*BillingDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InvoiceAdjustOverallDiscount) SetDiscount(v BillingDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *InvoiceAdjustOverallDiscount) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetReason

`func (o *InvoiceAdjustOverallDiscount) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceAdjustOverallDiscount) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceAdjustOverallDiscount) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoiceAdjustOverallDiscount) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


