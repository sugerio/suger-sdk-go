# HubspotCrmCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AcquiredOn** | Pointer to **int32** | UTC timestamp on receiving the auth response | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 

## Methods

### NewHubspotCrmCredential

`func NewHubspotCrmCredential() *HubspotCrmCredential`

NewHubspotCrmCredential instantiates a new HubspotCrmCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubspotCrmCredentialWithDefaults

`func NewHubspotCrmCredentialWithDefaults() *HubspotCrmCredential`

NewHubspotCrmCredentialWithDefaults instantiates a new HubspotCrmCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *HubspotCrmCredential) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *HubspotCrmCredential) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *HubspotCrmCredential) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *HubspotCrmCredential) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAcquiredOn

`func (o *HubspotCrmCredential) GetAcquiredOn() int32`

GetAcquiredOn returns the AcquiredOn field if non-nil, zero value otherwise.

### GetAcquiredOnOk

`func (o *HubspotCrmCredential) GetAcquiredOnOk() (*int32, bool)`

GetAcquiredOnOk returns a tuple with the AcquiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiredOn

`func (o *HubspotCrmCredential) SetAcquiredOn(v int32)`

SetAcquiredOn sets AcquiredOn field to given value.

### HasAcquiredOn

`func (o *HubspotCrmCredential) HasAcquiredOn() bool`

HasAcquiredOn returns a boolean if a field has been set.

### GetExpiresIn

`func (o *HubspotCrmCredential) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *HubspotCrmCredential) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *HubspotCrmCredential) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *HubspotCrmCredential) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *HubspotCrmCredential) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *HubspotCrmCredential) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *HubspotCrmCredential) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *HubspotCrmCredential) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


