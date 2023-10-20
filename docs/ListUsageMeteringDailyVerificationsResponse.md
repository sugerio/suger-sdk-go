# ListUsageMeteringDailyVerificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextOffset** | Pointer to **int32** |  | [optional] 
**UsageMeteringDailyVerifications** | Pointer to [**[]UsageMeteringDailyVerification**](UsageMeteringDailyVerification.md) | per day per dimension. | [optional] 

## Methods

### NewListUsageMeteringDailyVerificationsResponse

`func NewListUsageMeteringDailyVerificationsResponse() *ListUsageMeteringDailyVerificationsResponse`

NewListUsageMeteringDailyVerificationsResponse instantiates a new ListUsageMeteringDailyVerificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsageMeteringDailyVerificationsResponseWithDefaults

`func NewListUsageMeteringDailyVerificationsResponseWithDefaults() *ListUsageMeteringDailyVerificationsResponse`

NewListUsageMeteringDailyVerificationsResponseWithDefaults instantiates a new ListUsageMeteringDailyVerificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextOffset

`func (o *ListUsageMeteringDailyVerificationsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListUsageMeteringDailyVerificationsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListUsageMeteringDailyVerificationsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListUsageMeteringDailyVerificationsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetUsageMeteringDailyVerifications

`func (o *ListUsageMeteringDailyVerificationsResponse) GetUsageMeteringDailyVerifications() []UsageMeteringDailyVerification`

GetUsageMeteringDailyVerifications returns the UsageMeteringDailyVerifications field if non-nil, zero value otherwise.

### GetUsageMeteringDailyVerificationsOk

`func (o *ListUsageMeteringDailyVerificationsResponse) GetUsageMeteringDailyVerificationsOk() (*[]UsageMeteringDailyVerification, bool)`

GetUsageMeteringDailyVerificationsOk returns a tuple with the UsageMeteringDailyVerifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageMeteringDailyVerifications

`func (o *ListUsageMeteringDailyVerificationsResponse) SetUsageMeteringDailyVerifications(v []UsageMeteringDailyVerification)`

SetUsageMeteringDailyVerifications sets UsageMeteringDailyVerifications field to given value.

### HasUsageMeteringDailyVerifications

`func (o *ListUsageMeteringDailyVerificationsResponse) HasUsageMeteringDailyVerifications() bool`

HasUsageMeteringDailyVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


