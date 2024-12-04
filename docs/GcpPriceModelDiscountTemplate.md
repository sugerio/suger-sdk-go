# GcpPriceModelDiscountTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountEconomics** | Pointer to **string** |  | [optional] 
**DiscountPercentage** | Pointer to [**GcpAmountConstraint**](GcpAmountConstraint.md) |  | [optional] 
**DiscountedPrice** | Pointer to [**GcpPriceValue**](GcpPriceValue.md) |  | [optional] 
**HideDiscountPercentage** | Pointer to **bool** |  | [optional] 

## Methods

### NewGcpPriceModelDiscountTemplate

`func NewGcpPriceModelDiscountTemplate() *GcpPriceModelDiscountTemplate`

NewGcpPriceModelDiscountTemplate instantiates a new GcpPriceModelDiscountTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpPriceModelDiscountTemplateWithDefaults

`func NewGcpPriceModelDiscountTemplateWithDefaults() *GcpPriceModelDiscountTemplate`

NewGcpPriceModelDiscountTemplateWithDefaults instantiates a new GcpPriceModelDiscountTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountEconomics

`func (o *GcpPriceModelDiscountTemplate) GetDiscountEconomics() string`

GetDiscountEconomics returns the DiscountEconomics field if non-nil, zero value otherwise.

### GetDiscountEconomicsOk

`func (o *GcpPriceModelDiscountTemplate) GetDiscountEconomicsOk() (*string, bool)`

GetDiscountEconomicsOk returns a tuple with the DiscountEconomics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEconomics

`func (o *GcpPriceModelDiscountTemplate) SetDiscountEconomics(v string)`

SetDiscountEconomics sets DiscountEconomics field to given value.

### HasDiscountEconomics

`func (o *GcpPriceModelDiscountTemplate) HasDiscountEconomics() bool`

HasDiscountEconomics returns a boolean if a field has been set.

### GetDiscountPercentage

`func (o *GcpPriceModelDiscountTemplate) GetDiscountPercentage() GcpAmountConstraint`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *GcpPriceModelDiscountTemplate) GetDiscountPercentageOk() (*GcpAmountConstraint, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *GcpPriceModelDiscountTemplate) SetDiscountPercentage(v GcpAmountConstraint)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *GcpPriceModelDiscountTemplate) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *GcpPriceModelDiscountTemplate) GetDiscountedPrice() GcpPriceValue`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *GcpPriceModelDiscountTemplate) GetDiscountedPriceOk() (*GcpPriceValue, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *GcpPriceModelDiscountTemplate) SetDiscountedPrice(v GcpPriceValue)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *GcpPriceModelDiscountTemplate) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.

### GetHideDiscountPercentage

`func (o *GcpPriceModelDiscountTemplate) GetHideDiscountPercentage() bool`

GetHideDiscountPercentage returns the HideDiscountPercentage field if non-nil, zero value otherwise.

### GetHideDiscountPercentageOk

`func (o *GcpPriceModelDiscountTemplate) GetHideDiscountPercentageOk() (*bool, bool)`

GetHideDiscountPercentageOk returns a tuple with the HideDiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDiscountPercentage

`func (o *GcpPriceModelDiscountTemplate) SetHideDiscountPercentage(v bool)`

SetHideDiscountPercentage sets HideDiscountPercentage field to given value.

### HasHideDiscountPercentage

`func (o *GcpPriceModelDiscountTemplate) HasHideDiscountPercentage() bool`

HasHideDiscountPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


