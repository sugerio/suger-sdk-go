# GcpPriceTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAmount** | Pointer to **float32** | such as 0 | [optional] 
**Price** | Pointer to [**GcpPriceValue**](GcpPriceValue.md) |  | [optional] 
**StartingUsageAmount** | Pointer to **string** | such as \&quot;0\&quot; | [optional] 

## Methods

### NewGcpPriceTier

`func NewGcpPriceTier() *GcpPriceTier`

NewGcpPriceTier instantiates a new GcpPriceTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpPriceTierWithDefaults

`func NewGcpPriceTierWithDefaults() *GcpPriceTier`

NewGcpPriceTierWithDefaults instantiates a new GcpPriceTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAmount

`func (o *GcpPriceTier) GetFromAmount() float32`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *GcpPriceTier) GetFromAmountOk() (*float32, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *GcpPriceTier) SetFromAmount(v float32)`

SetFromAmount sets FromAmount field to given value.

### HasFromAmount

`func (o *GcpPriceTier) HasFromAmount() bool`

HasFromAmount returns a boolean if a field has been set.

### GetPrice

`func (o *GcpPriceTier) GetPrice() GcpPriceValue`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GcpPriceTier) GetPriceOk() (*GcpPriceValue, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GcpPriceTier) SetPrice(v GcpPriceValue)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GcpPriceTier) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStartingUsageAmount

`func (o *GcpPriceTier) GetStartingUsageAmount() string`

GetStartingUsageAmount returns the StartingUsageAmount field if non-nil, zero value otherwise.

### GetStartingUsageAmountOk

`func (o *GcpPriceTier) GetStartingUsageAmountOk() (*string, bool)`

GetStartingUsageAmountOk returns a tuple with the StartingUsageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingUsageAmount

`func (o *GcpPriceTier) SetStartingUsageAmount(v string)`

SetStartingUsageAmount sets StartingUsageAmount field to given value.

### HasStartingUsageAmount

`func (o *GcpPriceTier) HasStartingUsageAmount() bool`

HasStartingUsageAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


