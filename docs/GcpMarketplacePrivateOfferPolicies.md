# GcpMarketplacePrivateOfferPolicies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancellationPolicy** | Pointer to **string** | The cancellation policy of the offer. | [optional] 
**DefaultRenewalPolicy** | Pointer to **string** | The default renewal policy of the offer. | [optional] 
**DowngradePolicy** | Pointer to **string** | The downgrade policy of the offer. | [optional] 
**MaxRenewalTimes** | Pointer to **string** | The max renewal times of the offer. | [optional] 
**OfferDealType** | Pointer to [**GcpMarketplaceOfferDealType**](GcpMarketplaceOfferDealType.md) | The offer deal type of the offer. | [optional] 
**PurchaseApproval** | Pointer to **string** | The purchase approval of the offer. | [optional] 
**PurchaseApprovalOverride** | Pointer to **string** | The purchase approval override of the offer. | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferPolicies

`func NewGcpMarketplacePrivateOfferPolicies() *GcpMarketplacePrivateOfferPolicies`

NewGcpMarketplacePrivateOfferPolicies instantiates a new GcpMarketplacePrivateOfferPolicies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferPoliciesWithDefaults

`func NewGcpMarketplacePrivateOfferPoliciesWithDefaults() *GcpMarketplacePrivateOfferPolicies`

NewGcpMarketplacePrivateOfferPoliciesWithDefaults instantiates a new GcpMarketplacePrivateOfferPolicies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancellationPolicy

`func (o *GcpMarketplacePrivateOfferPolicies) GetCancellationPolicy() string`

GetCancellationPolicy returns the CancellationPolicy field if non-nil, zero value otherwise.

### GetCancellationPolicyOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetCancellationPolicyOk() (*string, bool)`

GetCancellationPolicyOk returns a tuple with the CancellationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationPolicy

`func (o *GcpMarketplacePrivateOfferPolicies) SetCancellationPolicy(v string)`

SetCancellationPolicy sets CancellationPolicy field to given value.

### HasCancellationPolicy

`func (o *GcpMarketplacePrivateOfferPolicies) HasCancellationPolicy() bool`

HasCancellationPolicy returns a boolean if a field has been set.

### GetDefaultRenewalPolicy

`func (o *GcpMarketplacePrivateOfferPolicies) GetDefaultRenewalPolicy() string`

GetDefaultRenewalPolicy returns the DefaultRenewalPolicy field if non-nil, zero value otherwise.

### GetDefaultRenewalPolicyOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetDefaultRenewalPolicyOk() (*string, bool)`

GetDefaultRenewalPolicyOk returns a tuple with the DefaultRenewalPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRenewalPolicy

`func (o *GcpMarketplacePrivateOfferPolicies) SetDefaultRenewalPolicy(v string)`

SetDefaultRenewalPolicy sets DefaultRenewalPolicy field to given value.

### HasDefaultRenewalPolicy

`func (o *GcpMarketplacePrivateOfferPolicies) HasDefaultRenewalPolicy() bool`

HasDefaultRenewalPolicy returns a boolean if a field has been set.

### GetDowngradePolicy

`func (o *GcpMarketplacePrivateOfferPolicies) GetDowngradePolicy() string`

GetDowngradePolicy returns the DowngradePolicy field if non-nil, zero value otherwise.

### GetDowngradePolicyOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetDowngradePolicyOk() (*string, bool)`

GetDowngradePolicyOk returns a tuple with the DowngradePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDowngradePolicy

`func (o *GcpMarketplacePrivateOfferPolicies) SetDowngradePolicy(v string)`

SetDowngradePolicy sets DowngradePolicy field to given value.

### HasDowngradePolicy

`func (o *GcpMarketplacePrivateOfferPolicies) HasDowngradePolicy() bool`

HasDowngradePolicy returns a boolean if a field has been set.

### GetMaxRenewalTimes

`func (o *GcpMarketplacePrivateOfferPolicies) GetMaxRenewalTimes() string`

GetMaxRenewalTimes returns the MaxRenewalTimes field if non-nil, zero value otherwise.

### GetMaxRenewalTimesOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetMaxRenewalTimesOk() (*string, bool)`

GetMaxRenewalTimesOk returns a tuple with the MaxRenewalTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRenewalTimes

`func (o *GcpMarketplacePrivateOfferPolicies) SetMaxRenewalTimes(v string)`

SetMaxRenewalTimes sets MaxRenewalTimes field to given value.

### HasMaxRenewalTimes

`func (o *GcpMarketplacePrivateOfferPolicies) HasMaxRenewalTimes() bool`

HasMaxRenewalTimes returns a boolean if a field has been set.

### GetOfferDealType

`func (o *GcpMarketplacePrivateOfferPolicies) GetOfferDealType() GcpMarketplaceOfferDealType`

GetOfferDealType returns the OfferDealType field if non-nil, zero value otherwise.

### GetOfferDealTypeOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetOfferDealTypeOk() (*GcpMarketplaceOfferDealType, bool)`

GetOfferDealTypeOk returns a tuple with the OfferDealType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferDealType

`func (o *GcpMarketplacePrivateOfferPolicies) SetOfferDealType(v GcpMarketplaceOfferDealType)`

SetOfferDealType sets OfferDealType field to given value.

### HasOfferDealType

`func (o *GcpMarketplacePrivateOfferPolicies) HasOfferDealType() bool`

HasOfferDealType returns a boolean if a field has been set.

### GetPurchaseApproval

`func (o *GcpMarketplacePrivateOfferPolicies) GetPurchaseApproval() string`

GetPurchaseApproval returns the PurchaseApproval field if non-nil, zero value otherwise.

### GetPurchaseApprovalOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetPurchaseApprovalOk() (*string, bool)`

GetPurchaseApprovalOk returns a tuple with the PurchaseApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseApproval

`func (o *GcpMarketplacePrivateOfferPolicies) SetPurchaseApproval(v string)`

SetPurchaseApproval sets PurchaseApproval field to given value.

### HasPurchaseApproval

`func (o *GcpMarketplacePrivateOfferPolicies) HasPurchaseApproval() bool`

HasPurchaseApproval returns a boolean if a field has been set.

### GetPurchaseApprovalOverride

`func (o *GcpMarketplacePrivateOfferPolicies) GetPurchaseApprovalOverride() string`

GetPurchaseApprovalOverride returns the PurchaseApprovalOverride field if non-nil, zero value otherwise.

### GetPurchaseApprovalOverrideOk

`func (o *GcpMarketplacePrivateOfferPolicies) GetPurchaseApprovalOverrideOk() (*string, bool)`

GetPurchaseApprovalOverrideOk returns a tuple with the PurchaseApprovalOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseApprovalOverride

`func (o *GcpMarketplacePrivateOfferPolicies) SetPurchaseApprovalOverride(v string)`

SetPurchaseApprovalOverride sets PurchaseApprovalOverride field to given value.

### HasPurchaseApprovalOverride

`func (o *GcpMarketplacePrivateOfferPolicies) HasPurchaseApprovalOverride() bool`

HasPurchaseApprovalOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


