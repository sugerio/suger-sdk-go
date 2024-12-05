# AwsMarketplaceCppoOpportunityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityEndDate** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Output only. | [optional] 
**NegativeTargeting** | Pointer to [**AwsMarketplaceCppoOpportunityNegativeTargeting**](AwsMarketplaceCppoOpportunityNegativeTargeting.md) | Negative targeting defines the criteria which any customer&#39;s profile should fulfill to be restricted to access the offer. | [optional] 
**OffersMaxQuantity** | Pointer to **int32** |  | [optional] 
**PositiveTargeting** | Pointer to [**AwsMarketplaceCppoOpportunityPositiveTargeting**](AwsMarketplaceCppoOpportunityPositiveTargeting.md) | Positive targeting defines the criteria which any buyer&#39;s profile should fulfill in order to be allowed to access the offer. | [optional] 
**ResellerAccountId** | Pointer to **string** |  | [optional] 
**ResellerLegalName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**AwsMarketplaceCppoOpportunityRuleType**](AwsMarketplaceCppoOpportunityRuleType.md) |  | [optional] 
**Usage** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsMarketplaceCppoOpportunityRule

`func NewAwsMarketplaceCppoOpportunityRule() *AwsMarketplaceCppoOpportunityRule`

NewAwsMarketplaceCppoOpportunityRule instantiates a new AwsMarketplaceCppoOpportunityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMarketplaceCppoOpportunityRuleWithDefaults

`func NewAwsMarketplaceCppoOpportunityRuleWithDefaults() *AwsMarketplaceCppoOpportunityRule`

NewAwsMarketplaceCppoOpportunityRuleWithDefaults instantiates a new AwsMarketplaceCppoOpportunityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityEndDate

`func (o *AwsMarketplaceCppoOpportunityRule) GetAvailabilityEndDate() string`

GetAvailabilityEndDate returns the AvailabilityEndDate field if non-nil, zero value otherwise.

### GetAvailabilityEndDateOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetAvailabilityEndDateOk() (*string, bool)`

GetAvailabilityEndDateOk returns a tuple with the AvailabilityEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityEndDate

`func (o *AwsMarketplaceCppoOpportunityRule) SetAvailabilityEndDate(v string)`

SetAvailabilityEndDate sets AvailabilityEndDate field to given value.

### HasAvailabilityEndDate

`func (o *AwsMarketplaceCppoOpportunityRule) HasAvailabilityEndDate() bool`

HasAvailabilityEndDate returns a boolean if a field has been set.

### GetId

`func (o *AwsMarketplaceCppoOpportunityRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsMarketplaceCppoOpportunityRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsMarketplaceCppoOpportunityRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNegativeTargeting

`func (o *AwsMarketplaceCppoOpportunityRule) GetNegativeTargeting() AwsMarketplaceCppoOpportunityNegativeTargeting`

GetNegativeTargeting returns the NegativeTargeting field if non-nil, zero value otherwise.

### GetNegativeTargetingOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetNegativeTargetingOk() (*AwsMarketplaceCppoOpportunityNegativeTargeting, bool)`

GetNegativeTargetingOk returns a tuple with the NegativeTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeTargeting

`func (o *AwsMarketplaceCppoOpportunityRule) SetNegativeTargeting(v AwsMarketplaceCppoOpportunityNegativeTargeting)`

SetNegativeTargeting sets NegativeTargeting field to given value.

### HasNegativeTargeting

`func (o *AwsMarketplaceCppoOpportunityRule) HasNegativeTargeting() bool`

HasNegativeTargeting returns a boolean if a field has been set.

### GetOffersMaxQuantity

`func (o *AwsMarketplaceCppoOpportunityRule) GetOffersMaxQuantity() int32`

GetOffersMaxQuantity returns the OffersMaxQuantity field if non-nil, zero value otherwise.

### GetOffersMaxQuantityOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetOffersMaxQuantityOk() (*int32, bool)`

GetOffersMaxQuantityOk returns a tuple with the OffersMaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffersMaxQuantity

`func (o *AwsMarketplaceCppoOpportunityRule) SetOffersMaxQuantity(v int32)`

SetOffersMaxQuantity sets OffersMaxQuantity field to given value.

### HasOffersMaxQuantity

`func (o *AwsMarketplaceCppoOpportunityRule) HasOffersMaxQuantity() bool`

HasOffersMaxQuantity returns a boolean if a field has been set.

### GetPositiveTargeting

`func (o *AwsMarketplaceCppoOpportunityRule) GetPositiveTargeting() AwsMarketplaceCppoOpportunityPositiveTargeting`

GetPositiveTargeting returns the PositiveTargeting field if non-nil, zero value otherwise.

### GetPositiveTargetingOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetPositiveTargetingOk() (*AwsMarketplaceCppoOpportunityPositiveTargeting, bool)`

GetPositiveTargetingOk returns a tuple with the PositiveTargeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveTargeting

`func (o *AwsMarketplaceCppoOpportunityRule) SetPositiveTargeting(v AwsMarketplaceCppoOpportunityPositiveTargeting)`

SetPositiveTargeting sets PositiveTargeting field to given value.

### HasPositiveTargeting

`func (o *AwsMarketplaceCppoOpportunityRule) HasPositiveTargeting() bool`

HasPositiveTargeting returns a boolean if a field has been set.

### GetResellerAccountId

`func (o *AwsMarketplaceCppoOpportunityRule) GetResellerAccountId() string`

GetResellerAccountId returns the ResellerAccountId field if non-nil, zero value otherwise.

### GetResellerAccountIdOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetResellerAccountIdOk() (*string, bool)`

GetResellerAccountIdOk returns a tuple with the ResellerAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerAccountId

`func (o *AwsMarketplaceCppoOpportunityRule) SetResellerAccountId(v string)`

SetResellerAccountId sets ResellerAccountId field to given value.

### HasResellerAccountId

`func (o *AwsMarketplaceCppoOpportunityRule) HasResellerAccountId() bool`

HasResellerAccountId returns a boolean if a field has been set.

### GetResellerLegalName

`func (o *AwsMarketplaceCppoOpportunityRule) GetResellerLegalName() string`

GetResellerLegalName returns the ResellerLegalName field if non-nil, zero value otherwise.

### GetResellerLegalNameOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetResellerLegalNameOk() (*string, bool)`

GetResellerLegalNameOk returns a tuple with the ResellerLegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerLegalName

`func (o *AwsMarketplaceCppoOpportunityRule) SetResellerLegalName(v string)`

SetResellerLegalName sets ResellerLegalName field to given value.

### HasResellerLegalName

`func (o *AwsMarketplaceCppoOpportunityRule) HasResellerLegalName() bool`

HasResellerLegalName returns a boolean if a field has been set.

### GetType

`func (o *AwsMarketplaceCppoOpportunityRule) GetType() AwsMarketplaceCppoOpportunityRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetTypeOk() (*AwsMarketplaceCppoOpportunityRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwsMarketplaceCppoOpportunityRule) SetType(v AwsMarketplaceCppoOpportunityRuleType)`

SetType sets Type field to given value.

### HasType

`func (o *AwsMarketplaceCppoOpportunityRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *AwsMarketplaceCppoOpportunityRule) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AwsMarketplaceCppoOpportunityRule) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AwsMarketplaceCppoOpportunityRule) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *AwsMarketplaceCppoOpportunityRule) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


