# BillingDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountType** | Pointer to [**BillingDiscountType**](BillingDiscountType.md) |  | [optional] 
**Value** | Pointer to **float32** |  | [optional] 

## Methods

### NewBillingDiscount

`func NewBillingDiscount() *BillingDiscount`

NewBillingDiscount instantiates a new BillingDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingDiscountWithDefaults

`func NewBillingDiscountWithDefaults() *BillingDiscount`

NewBillingDiscountWithDefaults instantiates a new BillingDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountType

`func (o *BillingDiscount) GetDiscountType() BillingDiscountType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *BillingDiscount) GetDiscountTypeOk() (*BillingDiscountType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *BillingDiscount) SetDiscountType(v BillingDiscountType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *BillingDiscount) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetValue

`func (o *BillingDiscount) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BillingDiscount) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BillingDiscount) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *BillingDiscount) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


