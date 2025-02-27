# UpdateBillableMetricParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BillableMetricStatus**](BillableMetricStatus.md) |  | [optional] 

## Methods

### NewUpdateBillableMetricParams

`func NewUpdateBillableMetricParams() *UpdateBillableMetricParams`

NewUpdateBillableMetricParams instantiates a new UpdateBillableMetricParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBillableMetricParamsWithDefaults

`func NewUpdateBillableMetricParamsWithDefaults() *UpdateBillableMetricParams`

NewUpdateBillableMetricParamsWithDefaults instantiates a new UpdateBillableMetricParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateBillableMetricParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateBillableMetricParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateBillableMetricParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateBillableMetricParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateBillableMetricParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBillableMetricParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBillableMetricParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBillableMetricParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateBillableMetricParams) GetStatus() BillableMetricStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateBillableMetricParams) GetStatusOk() (*BillableMetricStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateBillableMetricParams) SetStatus(v BillableMetricStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateBillableMetricParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


