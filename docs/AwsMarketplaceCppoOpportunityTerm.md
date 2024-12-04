# AwsMarketplaceCppoOpportunityTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** |  | [optional] 
**Documents** | Pointer to [**[]AwsMarketplaceCatalogLegalTermDocument**](AwsMarketplaceCatalogLegalTermDocument.md) |  | [optional] 
**Duration** | Pointer to **string** | ISO 8601 duration format. For example, \&quot;P12M\&quot; represents 12 months. | [optional] 
**Grants** | Pointer to [**[]AwsMarketplaceCppoOpportunityUpfrontPriceGrant**](AwsMarketplaceCppoOpportunityUpfrontPriceGrant.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaximumAgreementStartDate** | Pointer to **string** |  | [optional] 
**PositiveTargeting** | Pointer to [**AwsMarketplaceCppoOpportunityPositiveTargeting**](AwsMarketplaceCppoOpportunityPositiveTargeting.md) |  | [optional] 
**Price** | Pointer to **string** | For ResaleFixedUpfrontPricingTerm | [optional] 
**RateCards** | Pointer to [**[]AwsMarketplaceCatalogPricingTermRateCard**](AwsMarketplaceCatalogPricingTermRateCard.md) |  | [optional] 
**Schedule** | Pointer to [**[]AwsMarketplaceCppoOpportunityPaymentSchedule**](AwsMarketplaceCppoOpportunityPaymentSchedule.md) | For ResalePaymentScheduleTerm | [optional] 
**Type** | Pointer to [**AwsMarketplaceCppoOpportunityTermType**](AwsMarketplaceCppoOpportunityTermType.md) |  | [optional] 

## Methods

### NewAwsMarketplaceCppoOpportunityTerm

`func NewAwsMarketplaceCppoOpportunityTerm() *AwsMarketplaceCppoOpportunityTerm`

NewAwsMarketplaceCppoOpportunityTerm instantiates a new AwsMarketplaceCppoOpportunityTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCppoOpportunityTermWithDefaults

`func NewAwsMarketplaceCppoOpportunityTermWithDefaults() *AwsMarketplaceCppoOpportunityTerm`

NewAwsMarketplaceCppoOpportunityTermWithDefaults instantiates a new AwsMarketplaceCppoOpportunityTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *AwsMarketplaceCppoOpportunityTerm) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *AwsMarketplaceCppoOpportunityTerm) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *AwsMarketplaceCppoOpportunityTerm) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDocuments

`func (o *AwsMarketplaceCppoOpportunityTerm) GetDocuments() []AwsMarketplaceCatalogLegalTermDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetDocumentsOk() (*[]AwsMarketplaceCatalogLegalTermDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *AwsMarketplaceCppoOpportunityTerm) SetDocuments(v []AwsMarketplaceCatalogLegalTermDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *AwsMarketplaceCppoOpportunityTerm) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetDuration

`func (o *AwsMarketplaceCppoOpportunityTerm) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AwsMarketplaceCppoOpportunityTerm) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AwsMarketplaceCppoOpportunityTerm) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGrants

`func (o *AwsMarketplaceCppoOpportunityTerm) GetGrants() []AwsMarketplaceCppoOpportunityUpfrontPriceGrant`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetGrantsOk() (*[]AwsMarketplaceCppoOpportunityUpfrontPriceGrant, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *AwsMarketplaceCppoOpportunityTerm) SetGrants(v []AwsMarketplaceCppoOpportunityUpfrontPriceGrant)`

SetGrants sets Grants field to given value.

### HasGrants

`func (o *AwsMarketplaceCppoOpportunityTerm) HasGrants() bool`

HasGrants returns a boolean if a field has been set.

### GetId

`func (o *AwsMarketplaceCppoOpportunityTerm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsMarketplaceCppoOpportunityTerm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsMarketplaceCppoOpportunityTerm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaximumAgreementStartDate

`func (o *AwsMarketplaceCppoOpportunityTerm) GetMaximumAgreementStartDate() string`

GetMaximumAgreementStartDate returns the MaximumAgreementStartDate field if non-nil, zero value otherwise.

### GetMaximumAgreementStartDateOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetMaximumAgreementStartDateOk() (*string, bool)`

GetMaximumAgreementStartDateOk returns a tuple with the MaximumAgreementStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAgreementStartDate

`func (o *AwsMarketplaceCppoOpportunityTerm) SetMaximumAgreementStartDate(v string)`

SetMaximumAgreementStartDate sets MaximumAgreementStartDate field to given value.

### HasMaximumAgreementStartDate

`func (o *AwsMarketplaceCppoOpportunityTerm) HasMaximumAgreementStartDate() bool`

HasMaximumAgreementStartDate returns a boolean if a field has been set.

### GetPositiveTargeting

`func (o *AwsMarketplaceCppoOpportunityTerm) GetPositiveTargeting() AwsMarketplaceCppoOpportunityPositiveTargeting`

GetPositiveTargeting returns the PositiveTargeting field if non-nil, zero value otherwise.

### GetPositiveTargetingOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetPositiveTargetingOk() (*AwsMarketplaceCppoOpportunityPositiveTargeting, bool)`

GetPositiveTargetingOk returns a tuple with the PositiveTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveTargeting

`func (o *AwsMarketplaceCppoOpportunityTerm) SetPositiveTargeting(v AwsMarketplaceCppoOpportunityPositiveTargeting)`

SetPositiveTargeting sets PositiveTargeting field to given value.

### HasPositiveTargeting

`func (o *AwsMarketplaceCppoOpportunityTerm) HasPositiveTargeting() bool`

HasPositiveTargeting returns a boolean if a field has been set.

### GetPrice

`func (o *AwsMarketplaceCppoOpportunityTerm) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AwsMarketplaceCppoOpportunityTerm) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AwsMarketplaceCppoOpportunityTerm) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRateCards

`func (o *AwsMarketplaceCppoOpportunityTerm) GetRateCards() []AwsMarketplaceCatalogPricingTermRateCard`

GetRateCards returns the RateCards field if non-nil, zero value otherwise.

### GetRateCardsOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetRateCardsOk() (*[]AwsMarketplaceCatalogPricingTermRateCard, bool)`

GetRateCardsOk returns a tuple with the RateCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateCards

`func (o *AwsMarketplaceCppoOpportunityTerm) SetRateCards(v []AwsMarketplaceCatalogPricingTermRateCard)`

SetRateCards sets RateCards field to given value.

### HasRateCards

`func (o *AwsMarketplaceCppoOpportunityTerm) HasRateCards() bool`

HasRateCards returns a boolean if a field has been set.

### GetSchedule

`func (o *AwsMarketplaceCppoOpportunityTerm) GetSchedule() []AwsMarketplaceCppoOpportunityPaymentSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetScheduleOk() (*[]AwsMarketplaceCppoOpportunityPaymentSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AwsMarketplaceCppoOpportunityTerm) SetSchedule(v []AwsMarketplaceCppoOpportunityPaymentSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AwsMarketplaceCppoOpportunityTerm) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetType

`func (o *AwsMarketplaceCppoOpportunityTerm) GetType() AwsMarketplaceCppoOpportunityTermType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsMarketplaceCppoOpportunityTerm) GetTypeOk() (*AwsMarketplaceCppoOpportunityTermType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsMarketplaceCppoOpportunityTerm) SetType(v AwsMarketplaceCppoOpportunityTermType)`

SetType sets Type field to given value.

### HasType

`func (o *AwsMarketplaceCppoOpportunityTerm) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


