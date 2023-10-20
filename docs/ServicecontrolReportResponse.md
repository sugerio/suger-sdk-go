# ServicecontrolReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportErrors** | Pointer to [**[]ServicecontrolReportError**](ServicecontrolReportError.md) | ReportErrors: Partial failures, one for each &#x60;Operation&#x60; in the request that failed processing. There are three possible combinations of the RPC status: 1. The combination of a successful RPC status and an empty &#x60;report_errors&#x60; list indicates a complete success where all &#x60;Operations&#x60; in the request are processed successfully. 2. The combination of a successful RPC status and a non-empty &#x60;report_errors&#x60; list indicates a partial success where some &#x60;Operations&#x60; in the request succeeded. Each &#x60;Operation&#x60; that failed processing has a corresponding item in this list. 3. A failed RPC status indicates a general non-deterministic failure. When this happens, it&#39;s impossible to know which of the &#39;Operations&#39; in the request succeeded or failed. | [optional] 
**ServiceConfigId** | Pointer to **string** | ServiceConfigId: The actual config id used to process the request. | [optional] 
**ServiceRolloutId** | Pointer to **string** | ServiceRolloutId: The current service rollout id used to process the request. | [optional] 

## Methods

### NewServicecontrolReportResponse

`func NewServicecontrolReportResponse() *ServicecontrolReportResponse`

NewServicecontrolReportResponse instantiates a new ServicecontrolReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicecontrolReportResponseWithDefaults

`func NewServicecontrolReportResponseWithDefaults() *ServicecontrolReportResponse`

NewServicecontrolReportResponseWithDefaults instantiates a new ServicecontrolReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportErrors

`func (o *ServicecontrolReportResponse) GetReportErrors() []ServicecontrolReportError`

GetReportErrors returns the ReportErrors field if non-nil, zero value otherwise.

### GetReportErrorsOk

`func (o *ServicecontrolReportResponse) GetReportErrorsOk() (*[]ServicecontrolReportError, bool)`

GetReportErrorsOk returns a tuple with the ReportErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportErrors

`func (o *ServicecontrolReportResponse) SetReportErrors(v []ServicecontrolReportError)`

SetReportErrors sets ReportErrors field to given value.

### HasReportErrors

`func (o *ServicecontrolReportResponse) HasReportErrors() bool`

HasReportErrors returns a boolean if a field has been set.

### GetServiceConfigId

`func (o *ServicecontrolReportResponse) GetServiceConfigId() string`

GetServiceConfigId returns the ServiceConfigId field if non-nil, zero value otherwise.

### GetServiceConfigIdOk

`func (o *ServicecontrolReportResponse) GetServiceConfigIdOk() (*string, bool)`

GetServiceConfigIdOk returns a tuple with the ServiceConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceConfigId

`func (o *ServicecontrolReportResponse) SetServiceConfigId(v string)`

SetServiceConfigId sets ServiceConfigId field to given value.

### HasServiceConfigId

`func (o *ServicecontrolReportResponse) HasServiceConfigId() bool`

HasServiceConfigId returns a boolean if a field has been set.

### GetServiceRolloutId

`func (o *ServicecontrolReportResponse) GetServiceRolloutId() string`

GetServiceRolloutId returns the ServiceRolloutId field if non-nil, zero value otherwise.

### GetServiceRolloutIdOk

`func (o *ServicecontrolReportResponse) GetServiceRolloutIdOk() (*string, bool)`

GetServiceRolloutIdOk returns a tuple with the ServiceRolloutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRolloutId

`func (o *ServicecontrolReportResponse) SetServiceRolloutId(v string)`

SetServiceRolloutId sets ServiceRolloutId field to given value.

### HasServiceRolloutId

`func (o *ServicecontrolReportResponse) HasServiceRolloutId() bool`

HasServiceRolloutId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


