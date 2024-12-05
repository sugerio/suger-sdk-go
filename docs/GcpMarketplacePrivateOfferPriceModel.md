# GcpMarketplacePrivateOfferPriceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseOffer** | Pointer to **string** | in format of \&quot;projects/{projectNumber}/services/service-name.endpoints.gcp-project-id.cloud.goog/standardOffers/base-offer-id\&quot; | [optional] 
**Commitment** | Pointer to [**GcpMarketplacePrivateOfferPriceModelCommitment**](GcpMarketplacePrivateOfferPriceModelCommitment.md) |  | [optional] 
**FixedPrice** | Pointer to [**GcpMarketplacePrivateOfferPriceModelFixed**](GcpMarketplacePrivateOfferPriceModelFixed.md) |  | [optional] 
**OneTimeCredit** | Pointer to [**GcpPriceValue**](GcpPriceValue.md) | The one time credit in amount of money | [optional] 
**Overage** | Pointer to [**GcpMarketplacePrivateOfferPriceModelOverage**](GcpMarketplacePrivateOfferPriceModelOverage.md) |  | [optional] 
**Payg** | Pointer to [**GcpMarketplacePrivateOfferPriceModelPayg**](GcpMarketplacePrivateOfferPriceModelPayg.md) | Pay as you go. | [optional] 
**PreviousCreditBalancePolicy** | Pointer to **string** | such as \&quot;PREVIOUS_CREDIT_BALANCE_POLICY_UNSPECIFIED\&quot; | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferPriceModel

`func NewGcpMarketplacePrivateOfferPriceModel() *GcpMarketplacePrivateOfferPriceModel`

NewGcpMarketplacePrivateOfferPriceModel instantiates a new GcpMarketplacePrivateOfferPriceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferPriceModelWithDefaults

`func NewGcpMarketplacePrivateOfferPriceModelWithDefaults() *GcpMarketplacePrivateOfferPriceModel`

NewGcpMarketplacePrivateOfferPriceModelWithDefaults instantiates a new GcpMarketplacePrivateOfferPriceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseOffer

`func (o *GcpMarketplacePrivateOfferPriceModel) GetBaseOffer() string`

GetBaseOffer returns the BaseOffer field if non-nil, zero value otherwise.

### GetBaseOfferOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetBaseOfferOk() (*string, bool)`

GetBaseOfferOk returns a tuple with the BaseOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseOffer

`func (o *GcpMarketplacePrivateOfferPriceModel) SetBaseOffer(v string)`

SetBaseOffer sets BaseOffer field to given value.

### HasBaseOffer

`func (o *GcpMarketplacePrivateOfferPriceModel) HasBaseOffer() bool`

HasBaseOffer returns a boolean if a field has been set.

### GetCommitment

`func (o *GcpMarketplacePrivateOfferPriceModel) GetCommitment() GcpMarketplacePrivateOfferPriceModelCommitment`

GetCommitment returns the Commitment field if non-nil, zero value otherwise.

### GetCommitmentOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetCommitmentOk() (*GcpMarketplacePrivateOfferPriceModelCommitment, bool)`

GetCommitmentOk returns a tuple with the Commitment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitment

`func (o *GcpMarketplacePrivateOfferPriceModel) SetCommitment(v GcpMarketplacePrivateOfferPriceModelCommitment)`

SetCommitment sets Commitment field to given value.

### HasCommitment

`func (o *GcpMarketplacePrivateOfferPriceModel) HasCommitment() bool`

HasCommitment returns a boolean if a field has been set.

### GetFixedPrice

`func (o *GcpMarketplacePrivateOfferPriceModel) GetFixedPrice() GcpMarketplacePrivateOfferPriceModelFixed`

GetFixedPrice returns the FixedPrice field if non-nil, zero value otherwise.

### GetFixedPriceOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetFixedPriceOk() (*GcpMarketplacePrivateOfferPriceModelFixed, bool)`

GetFixedPriceOk returns a tuple with the FixedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedPrice

`func (o *GcpMarketplacePrivateOfferPriceModel) SetFixedPrice(v GcpMarketplacePrivateOfferPriceModelFixed)`

SetFixedPrice sets FixedPrice field to given value.

### HasFixedPrice

`func (o *GcpMarketplacePrivateOfferPriceModel) HasFixedPrice() bool`

HasFixedPrice returns a boolean if a field has been set.

### GetOneTimeCredit

`func (o *GcpMarketplacePrivateOfferPriceModel) GetOneTimeCredit() GcpPriceValue`

GetOneTimeCredit returns the OneTimeCredit field if non-nil, zero value otherwise.

### GetOneTimeCreditOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetOneTimeCreditOk() (*GcpPriceValue, bool)`

GetOneTimeCreditOk returns a tuple with the OneTimeCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeCredit

`func (o *GcpMarketplacePrivateOfferPriceModel) SetOneTimeCredit(v GcpPriceValue)`

SetOneTimeCredit sets OneTimeCredit field to given value.

### HasOneTimeCredit

`func (o *GcpMarketplacePrivateOfferPriceModel) HasOneTimeCredit() bool`

HasOneTimeCredit returns a boolean if a field has been set.

### GetOverage

`func (o *GcpMarketplacePrivateOfferPriceModel) GetOverage() GcpMarketplacePrivateOfferPriceModelOverage`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetOverageOk() (*GcpMarketplacePrivateOfferPriceModelOverage, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *GcpMarketplacePrivateOfferPriceModel) SetOverage(v GcpMarketplacePrivateOfferPriceModelOverage)`

SetOverage sets Overage field to given value.

### HasOverage

`func (o *GcpMarketplacePrivateOfferPriceModel) HasOverage() bool`

HasOverage returns a boolean if a field has been set.

### GetPayg

`func (o *GcpMarketplacePrivateOfferPriceModel) GetPayg() GcpMarketplacePrivateOfferPriceModelPayg`

GetPayg returns the Payg field if non-nil, zero value otherwise.

### GetPaygOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetPaygOk() (*GcpMarketplacePrivateOfferPriceModelPayg, bool)`

GetPaygOk returns a tuple with the Payg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayg

`func (o *GcpMarketplacePrivateOfferPriceModel) SetPayg(v GcpMarketplacePrivateOfferPriceModelPayg)`

SetPayg sets Payg field to given value.

### HasPayg

`func (o *GcpMarketplacePrivateOfferPriceModel) HasPayg() bool`

HasPayg returns a boolean if a field has been set.

### GetPreviousCreditBalancePolicy

`func (o *GcpMarketplacePrivateOfferPriceModel) GetPreviousCreditBalancePolicy() string`

GetPreviousCreditBalancePolicy returns the PreviousCreditBalancePolicy field if non-nil, zero value otherwise.

### GetPreviousCreditBalancePolicyOk

`func (o *GcpMarketplacePrivateOfferPriceModel) GetPreviousCreditBalancePolicyOk() (*string, bool)`

GetPreviousCreditBalancePolicyOk returns a tuple with the PreviousCreditBalancePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCreditBalancePolicy

`func (o *GcpMarketplacePrivateOfferPriceModel) SetPreviousCreditBalancePolicy(v string)`

SetPreviousCreditBalancePolicy sets PreviousCreditBalancePolicy field to given value.

### HasPreviousCreditBalancePolicy

`func (o *GcpMarketplacePrivateOfferPriceModel) HasPreviousCreditBalancePolicy() bool`

HasPreviousCreditBalancePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


