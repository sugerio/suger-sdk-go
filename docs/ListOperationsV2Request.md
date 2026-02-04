# ListOperationsV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**OperationFilter**](OperationFilter.md) | Filter expression for querying workflows. Supports nested logical operators (and/or) and comparison operators (&#x3D;, !&#x3D;, in, not_in). Example: {\&quot;operator\&quot;: \&quot;and\&quot;, \&quot;args\&quot;: [{\&quot;operator\&quot;: \&quot;&#x3D;\&quot;, \&quot;field\&quot;: \&quot;SugerService\&quot;, \&quot;value\&quot;: \&quot;Marketplace\&quot;}, {\&quot;operator\&quot;: \&quot;&#x3D;\&quot;, \&quot;field\&quot;: \&quot;SugerPartner\&quot;, \&quot;value\&quot;: \&quot;AWS\&quot;}]} | [optional] 
**Limit** | Pointer to **int32** | Limit the number of results returned (1-100), default 100 | [optional] 
**OffsetToken** | Pointer to **string** | OffsetToken for pagination, use nextOffsetToken from previous response | [optional] 

## Methods

### NewListOperationsV2Request

`func NewListOperationsV2Request() *ListOperationsV2Request`

NewListOperationsV2Request instantiates a new ListOperationsV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOperationsV2RequestWithDefaults

`func NewListOperationsV2RequestWithDefaults() *ListOperationsV2Request`

NewListOperationsV2RequestWithDefaults instantiates a new ListOperationsV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ListOperationsV2Request) GetFilter() OperationFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListOperationsV2Request) GetFilterOk() (*OperationFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListOperationsV2Request) SetFilter(v OperationFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListOperationsV2Request) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLimit

`func (o *ListOperationsV2Request) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListOperationsV2Request) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListOperationsV2Request) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListOperationsV2Request) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffsetToken

`func (o *ListOperationsV2Request) GetOffsetToken() string`

GetOffsetToken returns the OffsetToken field if non-nil, zero value otherwise.

### GetOffsetTokenOk

`func (o *ListOperationsV2Request) GetOffsetTokenOk() (*string, bool)`

GetOffsetTokenOk returns a tuple with the OffsetToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetToken

`func (o *ListOperationsV2Request) SetOffsetToken(v string)`

SetOffsetToken sets OffsetToken field to given value.

### HasOffsetToken

`func (o *ListOperationsV2Request) HasOffsetToken() bool`

HasOffsetToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


