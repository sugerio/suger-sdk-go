# AwsMarketplaceCppoOpportunityPositiveTargeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerAccounts** | Pointer to [**[]AwsMarketplaceBuyerAccount**](AwsMarketplaceBuyerAccount.md) | List of AWS account IDs that are allowed to subscribe to the offer. | [optional] 
**CountryCodes** | Pointer to **[]string** | List as option for allowing targeting based on country. If the intention isn’t to target the offer to a country, this field should be omitted. If it’s present, the list must contain at least one country code. Each element in this list should be a valid 2-letter country code, using this format: ISO 3166-1 alpha-2. | [optional] 

## Methods

### NewAwsMarketplaceCppoOpportunityPositiveTargeting

`func NewAwsMarketplaceCppoOpportunityPositiveTargeting() *AwsMarketplaceCppoOpportunityPositiveTargeting`

NewAwsMarketplaceCppoOpportunityPositiveTargeting instantiates a new AwsMarketplaceCppoOpportunityPositiveTargeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCppoOpportunityPositiveTargetingWithDefaults

`func NewAwsMarketplaceCppoOpportunityPositiveTargetingWithDefaults() *AwsMarketplaceCppoOpportunityPositiveTargeting`

NewAwsMarketplaceCppoOpportunityPositiveTargetingWithDefaults instantiates a new AwsMarketplaceCppoOpportunityPositiveTargeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerAccounts

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) GetBuyerAccounts() []AwsMarketplaceBuyerAccount`

GetBuyerAccounts returns the BuyerAccounts field if non-nil, zero value otherwise.

### GetBuyerAccountsOk

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) GetBuyerAccountsOk() (*[]AwsMarketplaceBuyerAccount, bool)`

GetBuyerAccountsOk returns a tuple with the BuyerAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerAccounts

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) SetBuyerAccounts(v []AwsMarketplaceBuyerAccount)`

SetBuyerAccounts sets BuyerAccounts field to given value.

### HasBuyerAccounts

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) HasBuyerAccounts() bool`

HasBuyerAccounts returns a boolean if a field has been set.

### GetCountryCodes

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) GetCountryCodes() []string`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) GetCountryCodesOk() (*[]string, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) SetCountryCodes(v []string)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *AwsMarketplaceCppoOpportunityPositiveTargeting) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


