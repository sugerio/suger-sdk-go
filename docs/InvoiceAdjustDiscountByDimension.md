# InvoiceAdjustDiscountByDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionKey** | Pointer to **string** |  | [optional] 
**Discount** | Pointer to [**BillingDiscount**](BillingDiscount.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceAdjustDiscountByDimension

`func NewInvoiceAdjustDiscountByDimension() *InvoiceAdjustDiscountByDimension`

NewInvoiceAdjustDiscountByDimension instantiates a new InvoiceAdjustDiscountByDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceAdjustDiscountByDimensionWithDefaults

`func NewInvoiceAdjustDiscountByDimensionWithDefaults() *InvoiceAdjustDiscountByDimension`

NewInvoiceAdjustDiscountByDimensionWithDefaults instantiates a new InvoiceAdjustDiscountByDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionKey

`func (o *InvoiceAdjustDiscountByDimension) GetDimensionKey() string`

GetDimensionKey returns the DimensionKey field if non-nil, zero value otherwise.

### GetDimensionKeyOk

`func (o *InvoiceAdjustDiscountByDimension) GetDimensionKeyOk() (*string, bool)`

GetDimensionKeyOk returns a tuple with the DimensionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionKey

`func (o *InvoiceAdjustDiscountByDimension) SetDimensionKey(v string)`

SetDimensionKey sets DimensionKey field to given value.

### HasDimensionKey

`func (o *InvoiceAdjustDiscountByDimension) HasDimensionKey() bool`

HasDimensionKey returns a boolean if a field has been set.

### GetDiscount

`func (o *InvoiceAdjustDiscountByDimension) GetDiscount() BillingDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InvoiceAdjustDiscountByDimension) GetDiscountOk() (*BillingDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InvoiceAdjustDiscountByDimension) SetDiscount(v BillingDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *InvoiceAdjustDiscountByDimension) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetReason

`func (o *InvoiceAdjustDiscountByDimension) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *InvoiceAdjustDiscountByDimension) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *InvoiceAdjustDiscountByDimension) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *InvoiceAdjustDiscountByDimension) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


