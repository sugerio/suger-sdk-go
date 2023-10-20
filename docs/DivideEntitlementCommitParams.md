# DivideEntitlementCommitParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitAmount** | Pointer to **float32** | The amount of the commit to be divided. If it is less or equal to 0.0, the total commit of the entitlement will be divided into multiple sub entitlement terms with credit. | [optional] 
**StartDates** | Pointer to [**[]time.Time**](time.Time.md) | The start dates of the sub entitlement terms. The end date of the last sub entitlement term is the end date of the parent entitlement term. The first start date must be the same as the start date of the parent entitlement term. | [optional] 

## Methods

### NewDivideEntitlementCommitParams

`func NewDivideEntitlementCommitParams() *DivideEntitlementCommitParams`

NewDivideEntitlementCommitParams instantiates a new DivideEntitlementCommitParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDivideEntitlementCommitParamsWithDefaults

`func NewDivideEntitlementCommitParamsWithDefaults() *DivideEntitlementCommitParams`

NewDivideEntitlementCommitParamsWithDefaults instantiates a new DivideEntitlementCommitParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitAmount

`func (o *DivideEntitlementCommitParams) GetCommitAmount() float32`

GetCommitAmount returns the CommitAmount field if non-nil, zero value otherwise.

### GetCommitAmountOk

`func (o *DivideEntitlementCommitParams) GetCommitAmountOk() (*float32, bool)`

GetCommitAmountOk returns a tuple with the CommitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitAmount

`func (o *DivideEntitlementCommitParams) SetCommitAmount(v float32)`

SetCommitAmount sets CommitAmount field to given value.

### HasCommitAmount

`func (o *DivideEntitlementCommitParams) HasCommitAmount() bool`

HasCommitAmount returns a boolean if a field has been set.

### GetStartDates

`func (o *DivideEntitlementCommitParams) GetStartDates() []time.Time`

GetStartDates returns the StartDates field if non-nil, zero value otherwise.

### GetStartDatesOk

`func (o *DivideEntitlementCommitParams) GetStartDatesOk() (*[]time.Time, bool)`

GetStartDatesOk returns a tuple with the StartDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDates

`func (o *DivideEntitlementCommitParams) SetStartDates(v []time.Time)`

SetStartDates sets StartDates field to given value.

### HasStartDates

`func (o *DivideEntitlementCommitParams) HasStartDates() bool`

HasStartDates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


