# BillableMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationType** | Pointer to [**BillableMetricAggregationType**](BillableMetricAggregationType.md) |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**BillableMetricInfo**](BillableMetricInfo.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BillableMetricStatus**](BillableMetricStatus.md) |  | [optional] 

## Methods

### NewBillableMetric

`func NewBillableMetric() *BillableMetric`

NewBillableMetric instantiates a new BillableMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableMetricWithDefaults

`func NewBillableMetricWithDefaults() *BillableMetric`

NewBillableMetricWithDefaults instantiates a new BillableMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationType

`func (o *BillableMetric) GetAggregationType() BillableMetricAggregationType`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *BillableMetric) GetAggregationTypeOk() (*BillableMetricAggregationType, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *BillableMetric) SetAggregationType(v BillableMetricAggregationType)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *BillableMetric) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetCreationTime

`func (o *BillableMetric) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BillableMetric) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BillableMetric) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BillableMetric) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *BillableMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillableMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillableMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillableMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BillableMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillableMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillableMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillableMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *BillableMetric) GetInfo() BillableMetricInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillableMetric) GetInfoOk() (*BillableMetricInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillableMetric) SetInfo(v BillableMetricInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillableMetric) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *BillableMetric) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *BillableMetric) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *BillableMetric) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *BillableMetric) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *BillableMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillableMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *BillableMetric) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BillableMetric) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BillableMetric) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *BillableMetric) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStatus

`func (o *BillableMetric) GetStatus() BillableMetricStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillableMetric) GetStatusOk() (*BillableMetricStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillableMetric) SetStatus(v BillableMetricStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillableMetric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


