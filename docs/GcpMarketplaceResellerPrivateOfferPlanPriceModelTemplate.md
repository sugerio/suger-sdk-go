# GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseOffer** | Pointer to **string** | in format of \&quot;projects/{projectNumber}/services/service-name.endpoints.gcp-project-id.cloud.goog/standardOffers/base-offer-id\&quot; | [optional] 
**Commitment** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateCommitment**](GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateCommitment.md) |  | [optional] 
**Consumption** | Pointer to **string** |  | [optional] 
**FixedPrice** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateFixedPrice**](GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateFixedPrice.md) |  | [optional] 
**Overage** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateOverage**](GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateOverage.md) |  | [optional] 
**Payg** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplatePayg**](GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplatePayg.md) |  | [optional] 
**Subscription** | Pointer to **string** |  | [optional] 

## Methods

### NewGcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate

`func NewGcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate() *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate`

NewGcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate instantiates a new GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateWithDefaults

`func NewGcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateWithDefaults() *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate`

NewGcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseOffer

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetBaseOffer() string`

GetBaseOffer returns the BaseOffer field if non-nil, zero value otherwise.

### GetBaseOfferOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetBaseOfferOk() (*string, bool)`

GetBaseOfferOk returns a tuple with the BaseOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOffer

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetBaseOffer(v string)`

SetBaseOffer sets BaseOffer field to given value.

### HasBaseOffer

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasBaseOffer() bool`

HasBaseOffer returns a boolean if a field has been set.

### GetCommitment

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetCommitment() GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateCommitment`

GetCommitment returns the Commitment field if non-nil, zero value otherwise.

### GetCommitmentOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetCommitmentOk() (*GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateCommitment, bool)`

GetCommitmentOk returns a tuple with the Commitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitment

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetCommitment(v GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateCommitment)`

SetCommitment sets Commitment field to given value.

### HasCommitment

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasCommitment() bool`

HasCommitment returns a boolean if a field has been set.

### GetConsumption

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetConsumption() string`

GetConsumption returns the Consumption field if non-nil, zero value otherwise.

### GetConsumptionOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetConsumptionOk() (*string, bool)`

GetConsumptionOk returns a tuple with the Consumption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumption

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetConsumption(v string)`

SetConsumption sets Consumption field to given value.

### HasConsumption

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasConsumption() bool`

HasConsumption returns a boolean if a field has been set.

### GetFixedPrice

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetFixedPrice() GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateFixedPrice`

GetFixedPrice returns the FixedPrice field if non-nil, zero value otherwise.

### GetFixedPriceOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetFixedPriceOk() (*GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateFixedPrice, bool)`

GetFixedPriceOk returns a tuple with the FixedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPrice

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetFixedPrice(v GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateFixedPrice)`

SetFixedPrice sets FixedPrice field to given value.

### HasFixedPrice

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasFixedPrice() bool`

HasFixedPrice returns a boolean if a field has been set.

### GetOverage

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetOverage() GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateOverage`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetOverageOk() (*GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateOverage, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetOverage(v GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplateOverage)`

SetOverage sets Overage field to given value.

### HasOverage

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasOverage() bool`

HasOverage returns a boolean if a field has been set.

### GetPayg

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetPayg() GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplatePayg`

GetPayg returns the Payg field if non-nil, zero value otherwise.

### GetPaygOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetPaygOk() (*GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplatePayg, bool)`

GetPaygOk returns a tuple with the Payg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayg

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetPayg(v GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplatePayg)`

SetPayg sets Payg field to given value.

### HasPayg

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasPayg() bool`

HasPayg returns a boolean if a field has been set.

### GetSubscription

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


