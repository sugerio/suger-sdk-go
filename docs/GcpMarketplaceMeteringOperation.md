# GcpMarketplaceMeteringOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumerId** | Pointer to **string** | ConsumerId: Identity of the consumer who is using the service. This field should be filled in for the operations initiated by a consumer, but not for service-initiated operations that are not related to a specific consumer. - This can be in one of the following formats: - project:PROJECT_ID, - project&#x60;_&#x60;number:PROJECT_NUMBER, - projects/PROJECT_ID or PROJECT_NUMBER, - folders/FOLDER_NUMBER, - organizations/ORGANIZATION_NUMBER, - api&#x60;_&#x60;key:API_KEY. | [optional] 
**EndTime** | Pointer to **string** | EndTime: End time of the operation. Required when the operation is used in ServiceController.Report, but optional when the operation is used in ServiceController.Check. | [optional] 
**Labels** | Pointer to **map[string]string** | Labels: Labels describing the operation. Only the following labels are allowed: - Labels describing monitored resources as defined in the service configuration. - Default labels of metric values. When specified, labels defined in the metric value override these default. - The following labels defined by Google Cloud Platform: - &#x60;cloud.googleapis.com/location&#x60; describing the location where the operation happened, - &#x60;servicecontrol.googleapis.com/user_agent&#x60; describing the user agent of the API request, - &#x60;servicecontrol.googleapis.com/service_agent&#x60; describing the service used to handle the API request (e.g. ESP), - &#x60;servicecontrol.googleapis.com/platform&#x60; describing the platform where the API is served, such as App Engine, Compute Engine, or Kubernetes Engine. | [optional] 
**MetricValueSets** | Pointer to [**[]GcpMarketplaceMeteringMetricValueSet**](GcpMarketplaceMeteringMetricValueSet.md) | MetricValueSets: Represents information about this operation. Each MetricValueSet corresponds to a metric defined in the service configuration. The data type used in the MetricValueSet must agree with the data type specified in the metric definition. Within a single operation, it is not allowed to have more than one MetricValue instances that have the same metric names and identical label value combinations. If a request has such duplicated MetricValue instances, the entire request is rejected with an invalid argument error. | [optional] 
**OperationId** | Pointer to **string** | OperationId: Identity of the operation. This must be unique within the scope of the service that generated the operation. If the service calls Check() and Report() on the same operation, the two calls should carry the same id. UUID version 4 is recommended, though not required. In scenarios where an operation is computed from existing information and an idempotent id is desirable for deduplication purpose, UUID version 5 is recommended. See RFC 4122 for details. | [optional] 
**OperationName** | Pointer to **string** | OperationName: Fully qualified name of the operation. Reserved for future use. | [optional] 
**StartTime** | Pointer to **string** | StartTime: Required. Start time of the operation. | [optional] 

## Methods

### NewGcpMarketplaceMeteringOperation

`func NewGcpMarketplaceMeteringOperation() *GcpMarketplaceMeteringOperation`

NewGcpMarketplaceMeteringOperation instantiates a new GcpMarketplaceMeteringOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceMeteringOperationWithDefaults

`func NewGcpMarketplaceMeteringOperationWithDefaults() *GcpMarketplaceMeteringOperation`

NewGcpMarketplaceMeteringOperationWithDefaults instantiates a new GcpMarketplaceMeteringOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumerId

`func (o *GcpMarketplaceMeteringOperation) GetConsumerId() string`

GetConsumerId returns the ConsumerId field if non-nil, zero value otherwise.

### GetConsumerIdOk

`func (o *GcpMarketplaceMeteringOperation) GetConsumerIdOk() (*string, bool)`

GetConsumerIdOk returns a tuple with the ConsumerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerId

`func (o *GcpMarketplaceMeteringOperation) SetConsumerId(v string)`

SetConsumerId sets ConsumerId field to given value.

### HasConsumerId

`func (o *GcpMarketplaceMeteringOperation) HasConsumerId() bool`

HasConsumerId returns a boolean if a field has been set.

### GetEndTime

`func (o *GcpMarketplaceMeteringOperation) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GcpMarketplaceMeteringOperation) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GcpMarketplaceMeteringOperation) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GcpMarketplaceMeteringOperation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetLabels

`func (o *GcpMarketplaceMeteringOperation) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GcpMarketplaceMeteringOperation) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GcpMarketplaceMeteringOperation) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GcpMarketplaceMeteringOperation) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMetricValueSets

`func (o *GcpMarketplaceMeteringOperation) GetMetricValueSets() []GcpMarketplaceMeteringMetricValueSet`

GetMetricValueSets returns the MetricValueSets field if non-nil, zero value otherwise.

### GetMetricValueSetsOk

`func (o *GcpMarketplaceMeteringOperation) GetMetricValueSetsOk() (*[]GcpMarketplaceMeteringMetricValueSet, bool)`

GetMetricValueSetsOk returns a tuple with the MetricValueSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValueSets

`func (o *GcpMarketplaceMeteringOperation) SetMetricValueSets(v []GcpMarketplaceMeteringMetricValueSet)`

SetMetricValueSets sets MetricValueSets field to given value.

### HasMetricValueSets

`func (o *GcpMarketplaceMeteringOperation) HasMetricValueSets() bool`

HasMetricValueSets returns a boolean if a field has been set.

### GetOperationId

`func (o *GcpMarketplaceMeteringOperation) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *GcpMarketplaceMeteringOperation) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *GcpMarketplaceMeteringOperation) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *GcpMarketplaceMeteringOperation) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetOperationName

`func (o *GcpMarketplaceMeteringOperation) GetOperationName() string`

GetOperationName returns the OperationName field if non-nil, zero value otherwise.

### GetOperationNameOk

`func (o *GcpMarketplaceMeteringOperation) GetOperationNameOk() (*string, bool)`

GetOperationNameOk returns a tuple with the OperationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationName

`func (o *GcpMarketplaceMeteringOperation) SetOperationName(v string)`

SetOperationName sets OperationName field to given value.

### HasOperationName

`func (o *GcpMarketplaceMeteringOperation) HasOperationName() bool`

HasOperationName returns a boolean if a field has been set.

### GetStartTime

`func (o *GcpMarketplaceMeteringOperation) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GcpMarketplaceMeteringOperation) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GcpMarketplaceMeteringOperation) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GcpMarketplaceMeteringOperation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


