# BillableDimensionFeeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptions** | Pointer to **string** | Description of the pricing model that is used to show what the amount is for. like &#39;Bulk pricing: 0-100 units&#39;, &#39;Tiered pricing: 0-100 units&#39; | [optional] 
**FeeExpressions** | Pointer to **string** | FeeExpression is the expression used to calculate the fee that is used to show how the amount is calculated. like &#39;211 × $0.03&#39; | [optional] 

## Methods

### NewBillableDimensionFeeDetail

`func NewBillableDimensionFeeDetail() *BillableDimensionFeeDetail`

NewBillableDimensionFeeDetail instantiates a new BillableDimensionFeeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableDimensionFeeDetailWithDefaults

`func NewBillableDimensionFeeDetailWithDefaults() *BillableDimensionFeeDetail`

NewBillableDimensionFeeDetailWithDefaults instantiates a new BillableDimensionFeeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptions

`func (o *BillableDimensionFeeDetail) GetDescriptions() string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *BillableDimensionFeeDetail) GetDescriptionsOk() (*string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *BillableDimensionFeeDetail) SetDescriptions(v string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *BillableDimensionFeeDetail) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetFeeExpressions

`func (o *BillableDimensionFeeDetail) GetFeeExpressions() string`

GetFeeExpressions returns the FeeExpressions field if non-nil, zero value otherwise.

### GetFeeExpressionsOk

`func (o *BillableDimensionFeeDetail) GetFeeExpressionsOk() (*string, bool)`

GetFeeExpressionsOk returns a tuple with the FeeExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeExpressions

`func (o *BillableDimensionFeeDetail) SetFeeExpressions(v string)`

SetFeeExpressions sets FeeExpressions field to given value.

### HasFeeExpressions

`func (o *BillableDimensionFeeDetail) HasFeeExpressions() bool`

HasFeeExpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


