# OrbPriceMaximum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesToPriceIds** | Pointer to **[]string** | List of price_ids that this maximum amount applies to. For plan/plan phase maximums, this can be a subset of prices. | [optional] 
**MaximumAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewOrbPriceMaximum

`func NewOrbPriceMaximum() *OrbPriceMaximum`

NewOrbPriceMaximum instantiates a new OrbPriceMaximum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPriceMaximumWithDefaults

`func NewOrbPriceMaximumWithDefaults() *OrbPriceMaximum`

NewOrbPriceMaximumWithDefaults instantiates a new OrbPriceMaximum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesToPriceIds

`func (o *OrbPriceMaximum) GetAppliesToPriceIds() []string`

GetAppliesToPriceIds returns the AppliesToPriceIds field if non-nil, zero value otherwise.

### GetAppliesToPriceIdsOk

`func (o *OrbPriceMaximum) GetAppliesToPriceIdsOk() (*[]string, bool)`

GetAppliesToPriceIdsOk returns a tuple with the AppliesToPriceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToPriceIds

`func (o *OrbPriceMaximum) SetAppliesToPriceIds(v []string)`

SetAppliesToPriceIds sets AppliesToPriceIds field to given value.

### HasAppliesToPriceIds

`func (o *OrbPriceMaximum) HasAppliesToPriceIds() bool`

HasAppliesToPriceIds returns a boolean if a field has been set.

### GetMaximumAmount

`func (o *OrbPriceMaximum) GetMaximumAmount() string`

GetMaximumAmount returns the MaximumAmount field if non-nil, zero value otherwise.

### GetMaximumAmountOk

`func (o *OrbPriceMaximum) GetMaximumAmountOk() (*string, bool)`

GetMaximumAmountOk returns a tuple with the MaximumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAmount

`func (o *OrbPriceMaximum) SetMaximumAmount(v string)`

SetMaximumAmount sets MaximumAmount field to given value.

### HasMaximumAmount

`func (o *OrbPriceMaximum) HasMaximumAmount() bool`

HasMaximumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


