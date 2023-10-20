# OrbPriceMinimum

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesToPriceIds** | Pointer to **[]string** | List of price_ids that this minimum amount applies to. For plan/plan phase minimums, this can be a subset of prices. | [optional] 
**MinimumAmount** | Pointer to **string** |  | [optional] 

## Methods

### NewOrbPriceMinimum

`func NewOrbPriceMinimum() *OrbPriceMinimum`

NewOrbPriceMinimum instantiates a new OrbPriceMinimum object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbPriceMinimumWithDefaults

`func NewOrbPriceMinimumWithDefaults() *OrbPriceMinimum`

NewOrbPriceMinimumWithDefaults instantiates a new OrbPriceMinimum object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesToPriceIds

`func (o *OrbPriceMinimum) GetAppliesToPriceIds() []string`

GetAppliesToPriceIds returns the AppliesToPriceIds field if non-nil, zero value otherwise.

### GetAppliesToPriceIdsOk

`func (o *OrbPriceMinimum) GetAppliesToPriceIdsOk() (*[]string, bool)`

GetAppliesToPriceIdsOk returns a tuple with the AppliesToPriceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesToPriceIds

`func (o *OrbPriceMinimum) SetAppliesToPriceIds(v []string)`

SetAppliesToPriceIds sets AppliesToPriceIds field to given value.

### HasAppliesToPriceIds

`func (o *OrbPriceMinimum) HasAppliesToPriceIds() bool`

HasAppliesToPriceIds returns a boolean if a field has been set.

### GetMinimumAmount

`func (o *OrbPriceMinimum) GetMinimumAmount() string`

GetMinimumAmount returns the MinimumAmount field if non-nil, zero value otherwise.

### GetMinimumAmountOk

`func (o *OrbPriceMinimum) GetMinimumAmountOk() (*string, bool)`

GetMinimumAmountOk returns a tuple with the MinimumAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumAmount

`func (o *OrbPriceMinimum) SetMinimumAmount(v string)`

SetMinimumAmount sets MinimumAmount field to given value.

### HasMinimumAmount

`func (o *OrbPriceMinimum) HasMinimumAmount() bool`

HasMinimumAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


