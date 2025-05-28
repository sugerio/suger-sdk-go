# UpdateEntitlementPriceModelParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyTime** | Pointer to **string** |  | [optional] 
**BillableDimensions** | Pointer to [**[]BillableDimension**](BillableDimension.md) |  | [optional] 
**Commits** | Pointer to [**[]CommitDimension**](CommitDimension.md) |  | [optional] 

## Methods

### NewUpdateEntitlementPriceModelParams

`func NewUpdateEntitlementPriceModelParams() *UpdateEntitlementPriceModelParams`

NewUpdateEntitlementPriceModelParams instantiates a new UpdateEntitlementPriceModelParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEntitlementPriceModelParamsWithDefaults

`func NewUpdateEntitlementPriceModelParamsWithDefaults() *UpdateEntitlementPriceModelParams`

NewUpdateEntitlementPriceModelParamsWithDefaults instantiates a new UpdateEntitlementPriceModelParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyTime

`func (o *UpdateEntitlementPriceModelParams) GetApplyTime() string`

GetApplyTime returns the ApplyTime field if non-nil, zero value otherwise.

### GetApplyTimeOk

`func (o *UpdateEntitlementPriceModelParams) GetApplyTimeOk() (*string, bool)`

GetApplyTimeOk returns a tuple with the ApplyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyTime

`func (o *UpdateEntitlementPriceModelParams) SetApplyTime(v string)`

SetApplyTime sets ApplyTime field to given value.

### HasApplyTime

`func (o *UpdateEntitlementPriceModelParams) HasApplyTime() bool`

HasApplyTime returns a boolean if a field has been set.

### GetBillableDimensions

`func (o *UpdateEntitlementPriceModelParams) GetBillableDimensions() []BillableDimension`

GetBillableDimensions returns the BillableDimensions field if non-nil, zero value otherwise.

### GetBillableDimensionsOk

`func (o *UpdateEntitlementPriceModelParams) GetBillableDimensionsOk() (*[]BillableDimension, bool)`

GetBillableDimensionsOk returns a tuple with the BillableDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableDimensions

`func (o *UpdateEntitlementPriceModelParams) SetBillableDimensions(v []BillableDimension)`

SetBillableDimensions sets BillableDimensions field to given value.

### HasBillableDimensions

`func (o *UpdateEntitlementPriceModelParams) HasBillableDimensions() bool`

HasBillableDimensions returns a boolean if a field has been set.

### GetCommits

`func (o *UpdateEntitlementPriceModelParams) GetCommits() []CommitDimension`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *UpdateEntitlementPriceModelParams) GetCommitsOk() (*[]CommitDimension, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *UpdateEntitlementPriceModelParams) SetCommits(v []CommitDimension)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *UpdateEntitlementPriceModelParams) HasCommits() bool`

HasCommits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


