# BillableDimensionPriceModelDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | Amount is the amount that is calculated based on the FeeExpression | [optional] 
**BillableMetricKey** | Pointer to [**MeteringUsageRecordGroupByKey**](MeteringUsageRecordGroupByKey.md) | BillableMetricKey is the key of the billable metric | [optional] 
**Category** | Pointer to [**PriceModelCategory**](PriceModelCategory.md) | Category of this billable dimension. | [optional] 
**Details** | Pointer to [**[]BillableDimensionFeeDetail**](BillableDimensionFeeDetail.md) | Details is the details of the pricing model that is used to show what the amount is for. | [optional] 
**Discount** | Pointer to [**BillingDiscount**](BillingDiscount.md) | The discount of this billable dimension if applicable. | [optional] 
**DiscountExpression** | Pointer to **string** | DiscountExpression is the expression used to calculate the discount that is used to show how the discount is calculated. | [optional] 
**IsTrial** | Pointer to **bool** | Flag to indicate if this period is a trial period. | [optional] 
**MinimumCommit** | Pointer to **float32** | MinimumCommit is the minimum commit amount that is used to show the minimum commit amount. Will be ignored if the value is 0 or less. | [optional] 
**MinimumCommitScope** | Pointer to [**BillingMinimumCommitScope**](BillingMinimumCommitScope.md) | The minimum commit scope. The default value is \&quot;DIMENSION\&quot; if not set. | [optional] 
**Name** | Pointer to **string** | Name of this billable dimension. | [optional] 
**Quantity** | Pointer to **float32** | Final quantity of the billable dimension in the invoice period, which calculates the fee in price model. It may be the sum value of count/sum/unique_count or latest/max value according to different aggregation type. | [optional] 

## Methods

### NewBillableDimensionPriceModelDetail

`func NewBillableDimensionPriceModelDetail() *BillableDimensionPriceModelDetail`

NewBillableDimensionPriceModelDetail instantiates a new BillableDimensionPriceModelDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableDimensionPriceModelDetailWithDefaults

`func NewBillableDimensionPriceModelDetailWithDefaults() *BillableDimensionPriceModelDetail`

NewBillableDimensionPriceModelDetailWithDefaults instantiates a new BillableDimensionPriceModelDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillableDimensionPriceModelDetail) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillableDimensionPriceModelDetail) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillableDimensionPriceModelDetail) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillableDimensionPriceModelDetail) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBillableMetricKey

`func (o *BillableDimensionPriceModelDetail) GetBillableMetricKey() MeteringUsageRecordGroupByKey`

GetBillableMetricKey returns the BillableMetricKey field if non-nil, zero value otherwise.

### GetBillableMetricKeyOk

`func (o *BillableDimensionPriceModelDetail) GetBillableMetricKeyOk() (*MeteringUsageRecordGroupByKey, bool)`

GetBillableMetricKeyOk returns a tuple with the BillableMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricKey

`func (o *BillableDimensionPriceModelDetail) SetBillableMetricKey(v MeteringUsageRecordGroupByKey)`

SetBillableMetricKey sets BillableMetricKey field to given value.

### HasBillableMetricKey

`func (o *BillableDimensionPriceModelDetail) HasBillableMetricKey() bool`

HasBillableMetricKey returns a boolean if a field has been set.

### GetCategory

`func (o *BillableDimensionPriceModelDetail) GetCategory() PriceModelCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *BillableDimensionPriceModelDetail) GetCategoryOk() (*PriceModelCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *BillableDimensionPriceModelDetail) SetCategory(v PriceModelCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *BillableDimensionPriceModelDetail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDetails

`func (o *BillableDimensionPriceModelDetail) GetDetails() []BillableDimensionFeeDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BillableDimensionPriceModelDetail) GetDetailsOk() (*[]BillableDimensionFeeDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BillableDimensionPriceModelDetail) SetDetails(v []BillableDimensionFeeDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BillableDimensionPriceModelDetail) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDiscount

`func (o *BillableDimensionPriceModelDetail) GetDiscount() BillingDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillableDimensionPriceModelDetail) GetDiscountOk() (*BillingDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillableDimensionPriceModelDetail) SetDiscount(v BillingDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillableDimensionPriceModelDetail) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetDiscountExpression

`func (o *BillableDimensionPriceModelDetail) GetDiscountExpression() string`

GetDiscountExpression returns the DiscountExpression field if non-nil, zero value otherwise.

### GetDiscountExpressionOk

`func (o *BillableDimensionPriceModelDetail) GetDiscountExpressionOk() (*string, bool)`

GetDiscountExpressionOk returns a tuple with the DiscountExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountExpression

`func (o *BillableDimensionPriceModelDetail) SetDiscountExpression(v string)`

SetDiscountExpression sets DiscountExpression field to given value.

### HasDiscountExpression

`func (o *BillableDimensionPriceModelDetail) HasDiscountExpression() bool`

HasDiscountExpression returns a boolean if a field has been set.

### GetIsTrial

`func (o *BillableDimensionPriceModelDetail) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *BillableDimensionPriceModelDetail) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *BillableDimensionPriceModelDetail) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *BillableDimensionPriceModelDetail) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetMinimumCommit

`func (o *BillableDimensionPriceModelDetail) GetMinimumCommit() float32`

GetMinimumCommit returns the MinimumCommit field if non-nil, zero value otherwise.

### GetMinimumCommitOk

`func (o *BillableDimensionPriceModelDetail) GetMinimumCommitOk() (*float32, bool)`

GetMinimumCommitOk returns a tuple with the MinimumCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCommit

`func (o *BillableDimensionPriceModelDetail) SetMinimumCommit(v float32)`

SetMinimumCommit sets MinimumCommit field to given value.

### HasMinimumCommit

`func (o *BillableDimensionPriceModelDetail) HasMinimumCommit() bool`

HasMinimumCommit returns a boolean if a field has been set.

### GetMinimumCommitScope

`func (o *BillableDimensionPriceModelDetail) GetMinimumCommitScope() BillingMinimumCommitScope`

GetMinimumCommitScope returns the MinimumCommitScope field if non-nil, zero value otherwise.

### GetMinimumCommitScopeOk

`func (o *BillableDimensionPriceModelDetail) GetMinimumCommitScopeOk() (*BillingMinimumCommitScope, bool)`

GetMinimumCommitScopeOk returns a tuple with the MinimumCommitScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumCommitScope

`func (o *BillableDimensionPriceModelDetail) SetMinimumCommitScope(v BillingMinimumCommitScope)`

SetMinimumCommitScope sets MinimumCommitScope field to given value.

### HasMinimumCommitScope

`func (o *BillableDimensionPriceModelDetail) HasMinimumCommitScope() bool`

HasMinimumCommitScope returns a boolean if a field has been set.

### GetName

`func (o *BillableDimensionPriceModelDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableDimensionPriceModelDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableDimensionPriceModelDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillableDimensionPriceModelDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *BillableDimensionPriceModelDetail) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillableDimensionPriceModelDetail) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillableDimensionPriceModelDetail) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *BillableDimensionPriceModelDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


