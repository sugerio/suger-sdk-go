# AwsMarketplaceCatalogPricingTermRateCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**AwsMarketplaceCatalogPricingTermRateCardConstraints**](AwsMarketplaceCatalogPricingTermRateCardConstraints.md) | Defines constraints on how the term can be configured by acceptors. Applicable only to ConfigurableUpfrontPricingTerm. | [optional] 
**RateCard** | Pointer to [**[]AwsMarketplaceCatalogPricingTermRateCardItem**](AwsMarketplaceCatalogPricingTermRateCardItem.md) |  | [optional] 
**Selector** | Pointer to [**AwsMarketplaceCatalogPricingTermRateCardSelector**](AwsMarketplaceCatalogPricingTermRateCardSelector.md) | Selector is used to differentiate between the mutually exclusive rate cards in the same pricing term, to be selected by the buyer. Applicable only to ConfigurableUpfrontPricingTerm. | [optional] 

## Methods

### NewAwsMarketplaceCatalogPricingTermRateCard

`func NewAwsMarketplaceCatalogPricingTermRateCard() *AwsMarketplaceCatalogPricingTermRateCard`

NewAwsMarketplaceCatalogPricingTermRateCard instantiates a new AwsMarketplaceCatalogPricingTermRateCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCatalogPricingTermRateCardWithDefaults

`func NewAwsMarketplaceCatalogPricingTermRateCardWithDefaults() *AwsMarketplaceCatalogPricingTermRateCard`

NewAwsMarketplaceCatalogPricingTermRateCardWithDefaults instantiates a new AwsMarketplaceCatalogPricingTermRateCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *AwsMarketplaceCatalogPricingTermRateCard) GetConstraints() AwsMarketplaceCatalogPricingTermRateCardConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AwsMarketplaceCatalogPricingTermRateCard) GetConstraintsOk() (*AwsMarketplaceCatalogPricingTermRateCardConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AwsMarketplaceCatalogPricingTermRateCard) SetConstraints(v AwsMarketplaceCatalogPricingTermRateCardConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AwsMarketplaceCatalogPricingTermRateCard) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetRateCard

`func (o *AwsMarketplaceCatalogPricingTermRateCard) GetRateCard() []AwsMarketplaceCatalogPricingTermRateCardItem`

GetRateCard returns the RateCard field if non-nil, zero value otherwise.

### GetRateCardOk

`func (o *AwsMarketplaceCatalogPricingTermRateCard) GetRateCardOk() (*[]AwsMarketplaceCatalogPricingTermRateCardItem, bool)`

GetRateCardOk returns a tuple with the RateCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCard

`func (o *AwsMarketplaceCatalogPricingTermRateCard) SetRateCard(v []AwsMarketplaceCatalogPricingTermRateCardItem)`

SetRateCard sets RateCard field to given value.

### HasRateCard

`func (o *AwsMarketplaceCatalogPricingTermRateCard) HasRateCard() bool`

HasRateCard returns a boolean if a field has been set.

### GetSelector

`func (o *AwsMarketplaceCatalogPricingTermRateCard) GetSelector() AwsMarketplaceCatalogPricingTermRateCardSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *AwsMarketplaceCatalogPricingTermRateCard) GetSelectorOk() (*AwsMarketplaceCatalogPricingTermRateCardSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *AwsMarketplaceCatalogPricingTermRateCard) SetSelector(v AwsMarketplaceCatalogPricingTermRateCardSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *AwsMarketplaceCatalogPricingTermRateCard) HasSelector() bool`

HasSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


