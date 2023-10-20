# AzureProductFeatureAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMeters** | Pointer to [**[]AzureProductVariantCustomMeter**](AzureProductVariantCustomMeter.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**MarketStates** | Pointer to [**[]AzureMarketState**](AzureMarketState.md) |  | [optional] 
**Markets** | Pointer to [**[]AzureMarket**](AzureMarket.md) |  | [optional] 
**PriceSchedules** | Pointer to [**[]AzureProductVariantPriceSchedule**](AzureProductVariantPriceSchedule.md) |  | [optional] 
**Properties** | Pointer to [**[]AzureTypeValue**](AzureTypeValue.md) |  | [optional] 
**ResourceType** | Pointer to **string** | ResourceType &#x3D; FeatureAvailability | [optional] 
**SubscriptionAudiences** | Pointer to [**[]AzureAudience**](AzureAudience.md) |  | [optional] 
**TenantAudiences** | Pointer to [**[]AzureAudience**](AzureAudience.md) |  | [optional] 
**Trial** | Pointer to [**AzureProductVariantTrial**](AzureProductVariantTrial.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureProductFeatureAvailability

`func NewAzureProductFeatureAvailability() *AzureProductFeatureAvailability`

NewAzureProductFeatureAvailability instantiates a new AzureProductFeatureAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductFeatureAvailabilityWithDefaults

`func NewAzureProductFeatureAvailabilityWithDefaults() *AzureProductFeatureAvailability`

NewAzureProductFeatureAvailabilityWithDefaults instantiates a new AzureProductFeatureAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMeters

`func (o *AzureProductFeatureAvailability) GetCustomMeters() []AzureProductVariantCustomMeter`

GetCustomMeters returns the CustomMeters field if non-nil, zero value otherwise.

### GetCustomMetersOk

`func (o *AzureProductFeatureAvailability) GetCustomMetersOk() (*[]AzureProductVariantCustomMeter, bool)`

GetCustomMetersOk returns a tuple with the CustomMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMeters

`func (o *AzureProductFeatureAvailability) SetCustomMeters(v []AzureProductVariantCustomMeter)`

SetCustomMeters sets CustomMeters field to given value.

### HasCustomMeters

`func (o *AzureProductFeatureAvailability) HasCustomMeters() bool`

HasCustomMeters returns a boolean if a field has been set.

### GetId

`func (o *AzureProductFeatureAvailability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureProductFeatureAvailability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureProductFeatureAvailability) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureProductFeatureAvailability) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsHidden

`func (o *AzureProductFeatureAvailability) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AzureProductFeatureAvailability) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AzureProductFeatureAvailability) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AzureProductFeatureAvailability) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetMarketStates

`func (o *AzureProductFeatureAvailability) GetMarketStates() []AzureMarketState`

GetMarketStates returns the MarketStates field if non-nil, zero value otherwise.

### GetMarketStatesOk

`func (o *AzureProductFeatureAvailability) GetMarketStatesOk() (*[]AzureMarketState, bool)`

GetMarketStatesOk returns a tuple with the MarketStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketStates

`func (o *AzureProductFeatureAvailability) SetMarketStates(v []AzureMarketState)`

SetMarketStates sets MarketStates field to given value.

### HasMarketStates

`func (o *AzureProductFeatureAvailability) HasMarketStates() bool`

HasMarketStates returns a boolean if a field has been set.

### GetMarkets

`func (o *AzureProductFeatureAvailability) GetMarkets() []AzureMarket`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *AzureProductFeatureAvailability) GetMarketsOk() (*[]AzureMarket, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *AzureProductFeatureAvailability) SetMarkets(v []AzureMarket)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *AzureProductFeatureAvailability) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.

### GetPriceSchedules

`func (o *AzureProductFeatureAvailability) GetPriceSchedules() []AzureProductVariantPriceSchedule`

GetPriceSchedules returns the PriceSchedules field if non-nil, zero value otherwise.

### GetPriceSchedulesOk

`func (o *AzureProductFeatureAvailability) GetPriceSchedulesOk() (*[]AzureProductVariantPriceSchedule, bool)`

GetPriceSchedulesOk returns a tuple with the PriceSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSchedules

`func (o *AzureProductFeatureAvailability) SetPriceSchedules(v []AzureProductVariantPriceSchedule)`

SetPriceSchedules sets PriceSchedules field to given value.

### HasPriceSchedules

`func (o *AzureProductFeatureAvailability) HasPriceSchedules() bool`

HasPriceSchedules returns a boolean if a field has been set.

### GetProperties

`func (o *AzureProductFeatureAvailability) GetProperties() []AzureTypeValue`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AzureProductFeatureAvailability) GetPropertiesOk() (*[]AzureTypeValue, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AzureProductFeatureAvailability) SetProperties(v []AzureTypeValue)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AzureProductFeatureAvailability) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetResourceType

`func (o *AzureProductFeatureAvailability) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AzureProductFeatureAvailability) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AzureProductFeatureAvailability) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AzureProductFeatureAvailability) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSubscriptionAudiences

`func (o *AzureProductFeatureAvailability) GetSubscriptionAudiences() []AzureAudience`

GetSubscriptionAudiences returns the SubscriptionAudiences field if non-nil, zero value otherwise.

### GetSubscriptionAudiencesOk

`func (o *AzureProductFeatureAvailability) GetSubscriptionAudiencesOk() (*[]AzureAudience, bool)`

GetSubscriptionAudiencesOk returns a tuple with the SubscriptionAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionAudiences

`func (o *AzureProductFeatureAvailability) SetSubscriptionAudiences(v []AzureAudience)`

SetSubscriptionAudiences sets SubscriptionAudiences field to given value.

### HasSubscriptionAudiences

`func (o *AzureProductFeatureAvailability) HasSubscriptionAudiences() bool`

HasSubscriptionAudiences returns a boolean if a field has been set.

### GetTenantAudiences

`func (o *AzureProductFeatureAvailability) GetTenantAudiences() []AzureAudience`

GetTenantAudiences returns the TenantAudiences field if non-nil, zero value otherwise.

### GetTenantAudiencesOk

`func (o *AzureProductFeatureAvailability) GetTenantAudiencesOk() (*[]AzureAudience, bool)`

GetTenantAudiencesOk returns a tuple with the TenantAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantAudiences

`func (o *AzureProductFeatureAvailability) SetTenantAudiences(v []AzureAudience)`

SetTenantAudiences sets TenantAudiences field to given value.

### HasTenantAudiences

`func (o *AzureProductFeatureAvailability) HasTenantAudiences() bool`

HasTenantAudiences returns a boolean if a field has been set.

### GetTrial

`func (o *AzureProductFeatureAvailability) GetTrial() AzureProductVariantTrial`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *AzureProductFeatureAvailability) GetTrialOk() (*AzureProductVariantTrial, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *AzureProductFeatureAvailability) SetTrial(v AzureProductVariantTrial)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *AzureProductFeatureAvailability) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetVisibility

`func (o *AzureProductFeatureAvailability) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AzureProductFeatureAvailability) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AzureProductFeatureAvailability) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AzureProductFeatureAvailability) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


