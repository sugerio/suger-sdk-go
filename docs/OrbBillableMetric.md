# OrbBillableMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**OrbBillableMetricStatus**](OrbBillableMetricStatus.md) |  | [optional] 

## Methods

### NewOrbBillableMetric

`func NewOrbBillableMetric() *OrbBillableMetric`

NewOrbBillableMetric instantiates a new OrbBillableMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbBillableMetricWithDefaults

`func NewOrbBillableMetricWithDefaults() *OrbBillableMetric`

NewOrbBillableMetricWithDefaults instantiates a new OrbBillableMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OrbBillableMetric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrbBillableMetric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrbBillableMetric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrbBillableMetric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *OrbBillableMetric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrbBillableMetric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrbBillableMetric) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrbBillableMetric) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *OrbBillableMetric) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OrbBillableMetric) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OrbBillableMetric) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OrbBillableMetric) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *OrbBillableMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrbBillableMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrbBillableMetric) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrbBillableMetric) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OrbBillableMetric) GetStatus() OrbBillableMetricStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrbBillableMetric) GetStatusOk() (*OrbBillableMetricStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrbBillableMetric) SetStatus(v OrbBillableMetricStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrbBillableMetric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


