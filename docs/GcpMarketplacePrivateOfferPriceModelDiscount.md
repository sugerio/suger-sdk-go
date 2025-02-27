# GcpMarketplacePrivateOfferPriceModelDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountPercentage** | Pointer to [**GcpAmountUnit**](GcpAmountUnit.md) | such as {\&quot;units\&quot;: \&quot;0\&quot;, \&quot;nanos\&quot;: 0} as no discount, or {\&quot;units\&quot;: \&quot;10\&quot;, \&quot;nanos\&quot;: 0} as 10% off discount | [optional] 
**DiscountedPrice** | Pointer to [**GcpPriceValue**](GcpPriceValue.md) | The discounted price of the private offer. If the discount is 10% off, and the original price is $100, then the discounted price is $90. | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferPriceModelDiscount

`func NewGcpMarketplacePrivateOfferPriceModelDiscount() *GcpMarketplacePrivateOfferPriceModelDiscount`

NewGcpMarketplacePrivateOfferPriceModelDiscount instantiates a new GcpMarketplacePrivateOfferPriceModelDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferPriceModelDiscountWithDefaults

`func NewGcpMarketplacePrivateOfferPriceModelDiscountWithDefaults() *GcpMarketplacePrivateOfferPriceModelDiscount`

NewGcpMarketplacePrivateOfferPriceModelDiscountWithDefaults instantiates a new GcpMarketplacePrivateOfferPriceModelDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountPercentage

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) GetDiscountPercentage() GcpAmountUnit`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) GetDiscountPercentageOk() (*GcpAmountUnit, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) SetDiscountPercentage(v GcpAmountUnit)`

SetDiscountPercentage sets DiscountPercentage field to given value.

### HasDiscountPercentage

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) HasDiscountPercentage() bool`

HasDiscountPercentage returns a boolean if a field has been set.

### GetDiscountedPrice

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) GetDiscountedPrice() GcpPriceValue`

GetDiscountedPrice returns the DiscountedPrice field if non-nil, zero value otherwise.

### GetDiscountedPriceOk

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) GetDiscountedPriceOk() (*GcpPriceValue, bool)`

GetDiscountedPriceOk returns a tuple with the DiscountedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountedPrice

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) SetDiscountedPrice(v GcpPriceValue)`

SetDiscountedPrice sets DiscountedPrice field to given value.

### HasDiscountedPrice

`func (o *GcpMarketplacePrivateOfferPriceModelDiscount) HasDiscountedPrice() bool`

HasDiscountedPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


