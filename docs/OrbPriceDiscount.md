# OrbPriceDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountDiscount** | Pointer to **string** | Only available if discount_type is amount. | [optional] 
**AppliesToPriceIds** | Pointer to **[]string** |  | [optional] 
**DiscountType** | Pointer to [**OrbPriceDiscountType**](OrbPriceDiscountType.md) |  | [optional] 
**PercentageDiscount** | Pointer to **float32** | Only available if discount_type is percentage.This is a number between 0 and 1. | [optional] 
**TrialAmountDiscount** | Pointer to **string** | Only available if discount_type is trial | [optional] 
**UsageDiscount** | Pointer to **string** | Only available if discount_type is usage. Number of usage units that this discount is for | [optional] 

## Methods

### NewOrbPriceDiscount

`func NewOrbPriceDiscount() *OrbPriceDiscount`

NewOrbPriceDiscount instantiates a new OrbPriceDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPriceDiscountWithDefaults

`func NewOrbPriceDiscountWithDefaults() *OrbPriceDiscount`

NewOrbPriceDiscountWithDefaults instantiates a new OrbPriceDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountDiscount

`func (o *OrbPriceDiscount) GetAmountDiscount() string`

GetAmountDiscount returns the AmountDiscount field if non-nil, zero value otherwise.

### GetAmountDiscountOk

`func (o *OrbPriceDiscount) GetAmountDiscountOk() (*string, bool)`

GetAmountDiscountOk returns a tuple with the AmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDiscount

`func (o *OrbPriceDiscount) SetAmountDiscount(v string)`

SetAmountDiscount sets AmountDiscount field to given value.

### HasAmountDiscount

`func (o *OrbPriceDiscount) HasAmountDiscount() bool`

HasAmountDiscount returns a boolean if a field has been set.

### GetAppliesToPriceIds

`func (o *OrbPriceDiscount) GetAppliesToPriceIds() []string`

GetAppliesToPriceIds returns the AppliesToPriceIds field if non-nil, zero value otherwise.

### GetAppliesToPriceIdsOk

`func (o *OrbPriceDiscount) GetAppliesToPriceIdsOk() (*[]string, bool)`

GetAppliesToPriceIdsOk returns a tuple with the AppliesToPriceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToPriceIds

`func (o *OrbPriceDiscount) SetAppliesToPriceIds(v []string)`

SetAppliesToPriceIds sets AppliesToPriceIds field to given value.

### HasAppliesToPriceIds

`func (o *OrbPriceDiscount) HasAppliesToPriceIds() bool`

HasAppliesToPriceIds returns a boolean if a field has been set.

### GetDiscountType

`func (o *OrbPriceDiscount) GetDiscountType() OrbPriceDiscountType`

GetDiscountType returns the DiscountType field if non-nil, zero value otherwise.

### GetDiscountTypeOk

`func (o *OrbPriceDiscount) GetDiscountTypeOk() (*OrbPriceDiscountType, bool)`

GetDiscountTypeOk returns a tuple with the DiscountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountType

`func (o *OrbPriceDiscount) SetDiscountType(v OrbPriceDiscountType)`

SetDiscountType sets DiscountType field to given value.

### HasDiscountType

`func (o *OrbPriceDiscount) HasDiscountType() bool`

HasDiscountType returns a boolean if a field has been set.

### GetPercentageDiscount

`func (o *OrbPriceDiscount) GetPercentageDiscount() float32`

GetPercentageDiscount returns the PercentageDiscount field if non-nil, zero value otherwise.

### GetPercentageDiscountOk

`func (o *OrbPriceDiscount) GetPercentageDiscountOk() (*float32, bool)`

GetPercentageDiscountOk returns a tuple with the PercentageDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageDiscount

`func (o *OrbPriceDiscount) SetPercentageDiscount(v float32)`

SetPercentageDiscount sets PercentageDiscount field to given value.

### HasPercentageDiscount

`func (o *OrbPriceDiscount) HasPercentageDiscount() bool`

HasPercentageDiscount returns a boolean if a field has been set.

### GetTrialAmountDiscount

`func (o *OrbPriceDiscount) GetTrialAmountDiscount() string`

GetTrialAmountDiscount returns the TrialAmountDiscount field if non-nil, zero value otherwise.

### GetTrialAmountDiscountOk

`func (o *OrbPriceDiscount) GetTrialAmountDiscountOk() (*string, bool)`

GetTrialAmountDiscountOk returns a tuple with the TrialAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialAmountDiscount

`func (o *OrbPriceDiscount) SetTrialAmountDiscount(v string)`

SetTrialAmountDiscount sets TrialAmountDiscount field to given value.

### HasTrialAmountDiscount

`func (o *OrbPriceDiscount) HasTrialAmountDiscount() bool`

HasTrialAmountDiscount returns a boolean if a field has been set.

### GetUsageDiscount

`func (o *OrbPriceDiscount) GetUsageDiscount() string`

GetUsageDiscount returns the UsageDiscount field if non-nil, zero value otherwise.

### GetUsageDiscountOk

`func (o *OrbPriceDiscount) GetUsageDiscountOk() (*string, bool)`

GetUsageDiscountOk returns a tuple with the UsageDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDiscount

`func (o *OrbPriceDiscount) SetUsageDiscount(v string)`

SetUsageDiscount sets UsageDiscount field to given value.

### HasUsageDiscount

`func (o *OrbPriceDiscount) HasUsageDiscount() bool`

HasUsageDiscount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


