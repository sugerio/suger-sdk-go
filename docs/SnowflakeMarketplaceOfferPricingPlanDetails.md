# SnowflakeMarketplaceOfferPricingPlanDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name is the pricing plan name. Used for V2 offers with OVERRIDDEN type. | [optional] 
**Overrides** | Pointer to [**SnowflakeMarketplaceOfferOneTimeOverride**](SnowflakeMarketplaceOfferOneTimeOverride.md) |  | [optional] 
**Type** | Pointer to **string** | Type: DEFAULT for V1 private offer, INLINE for one time offer, OVERRIDDEN for V2 private offer with pricing plan | [optional] 

## Methods

### NewSnowflakeMarketplaceOfferPricingPlanDetails

`func NewSnowflakeMarketplaceOfferPricingPlanDetails() *SnowflakeMarketplaceOfferPricingPlanDetails`

NewSnowflakeMarketplaceOfferPricingPlanDetails instantiates a new SnowflakeMarketplaceOfferPricingPlanDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeMarketplaceOfferPricingPlanDetailsWithDefaults

`func NewSnowflakeMarketplaceOfferPricingPlanDetailsWithDefaults() *SnowflakeMarketplaceOfferPricingPlanDetails`

NewSnowflakeMarketplaceOfferPricingPlanDetailsWithDefaults instantiates a new SnowflakeMarketplaceOfferPricingPlanDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrides

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) GetOverrides() SnowflakeMarketplaceOfferOneTimeOverride`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) GetOverridesOk() (*SnowflakeMarketplaceOfferOneTimeOverride, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) SetOverrides(v SnowflakeMarketplaceOfferOneTimeOverride)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetType

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SnowflakeMarketplaceOfferPricingPlanDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


