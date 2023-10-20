# ListCosellOppReferralsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CosellOppReferrals** | Pointer to [**[]CosellOppReferral**](CosellOppReferral.md) |  | [optional] 
**NextOffset** | Pointer to **int32** | If it is nil, it means there is no more records. | [optional] 
**TotalCount** | Pointer to **int32** | Only available when the request is made with offset&#x3D;0. | [optional] 

## Methods

### NewListCosellOppReferralsResponse

`func NewListCosellOppReferralsResponse() *ListCosellOppReferralsResponse`

NewListCosellOppReferralsResponse instantiates a new ListCosellOppReferralsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCosellOppReferralsResponseWithDefaults

`func NewListCosellOppReferralsResponseWithDefaults() *ListCosellOppReferralsResponse`

NewListCosellOppReferralsResponseWithDefaults instantiates a new ListCosellOppReferralsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosellOppReferrals

`func (o *ListCosellOppReferralsResponse) GetCosellOppReferrals() []CosellOppReferral`

GetCosellOppReferrals returns the CosellOppReferrals field if non-nil, zero value otherwise.

### GetCosellOppReferralsOk

`func (o *ListCosellOppReferralsResponse) GetCosellOppReferralsOk() (*[]CosellOppReferral, bool)`

GetCosellOppReferralsOk returns a tuple with the CosellOppReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosellOppReferrals

`func (o *ListCosellOppReferralsResponse) SetCosellOppReferrals(v []CosellOppReferral)`

SetCosellOppReferrals sets CosellOppReferrals field to given value.

### HasCosellOppReferrals

`func (o *ListCosellOppReferralsResponse) HasCosellOppReferrals() bool`

HasCosellOppReferrals returns a boolean if a field has been set.

### GetNextOffset

`func (o *ListCosellOppReferralsResponse) GetNextOffset() int32`

GetNextOffset returns the NextOffset field if non-nil, zero value otherwise.

### GetNextOffsetOk

`func (o *ListCosellOppReferralsResponse) GetNextOffsetOk() (*int32, bool)`

GetNextOffsetOk returns a tuple with the NextOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOffset

`func (o *ListCosellOppReferralsResponse) SetNextOffset(v int32)`

SetNextOffset sets NextOffset field to given value.

### HasNextOffset

`func (o *ListCosellOppReferralsResponse) HasNextOffset() bool`

HasNextOffset returns a boolean if a field has been set.

### GetTotalCount

`func (o *ListCosellOppReferralsResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListCosellOppReferralsResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListCosellOppReferralsResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ListCosellOppReferralsResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


