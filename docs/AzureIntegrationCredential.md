# AzureIntegrationCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**ClientID** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**ExpiresOn** | Pointer to **string** | The time when the access token expires. | [optional] 
**RefreshToken** | Pointer to **string** | The refresh token used to refresh the access token. | [optional] 
**TenantID** | Pointer to **string** |  | [optional] 
**TokenScope** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 

## Methods

### NewAzureIntegrationCredential

`func NewAzureIntegrationCredential() *AzureIntegrationCredential`

NewAzureIntegrationCredential instantiates a new AzureIntegrationCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureIntegrationCredentialWithDefaults

`func NewAzureIntegrationCredentialWithDefaults() *AzureIntegrationCredential`

NewAzureIntegrationCredentialWithDefaults instantiates a new AzureIntegrationCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AzureIntegrationCredential) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AzureIntegrationCredential) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AzureIntegrationCredential) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AzureIntegrationCredential) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetClientID

`func (o *AzureIntegrationCredential) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AzureIntegrationCredential) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AzureIntegrationCredential) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *AzureIntegrationCredential) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### GetClientSecret

`func (o *AzureIntegrationCredential) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureIntegrationCredential) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureIntegrationCredential) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AzureIntegrationCredential) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetExpiresOn

`func (o *AzureIntegrationCredential) GetExpiresOn() string`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *AzureIntegrationCredential) GetExpiresOnOk() (*string, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *AzureIntegrationCredential) SetExpiresOn(v string)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *AzureIntegrationCredential) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AzureIntegrationCredential) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AzureIntegrationCredential) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AzureIntegrationCredential) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AzureIntegrationCredential) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTenantID

`func (o *AzureIntegrationCredential) GetTenantID() string`

GetTenantID returns the TenantID field if non-nil, zero value otherwise.

### GetTenantIDOk

`func (o *AzureIntegrationCredential) GetTenantIDOk() (*string, bool)`

GetTenantIDOk returns a tuple with the TenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantID

`func (o *AzureIntegrationCredential) SetTenantID(v string)`

SetTenantID sets TenantID field to given value.

### HasTenantID

`func (o *AzureIntegrationCredential) HasTenantID() bool`

HasTenantID returns a boolean if a field has been set.

### GetTokenScope

`func (o *AzureIntegrationCredential) GetTokenScope() string`

GetTokenScope returns the TokenScope field if non-nil, zero value otherwise.

### GetTokenScopeOk

`func (o *AzureIntegrationCredential) GetTokenScopeOk() (*string, bool)`

GetTokenScopeOk returns a tuple with the TokenScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenScope

`func (o *AzureIntegrationCredential) SetTokenScope(v string)`

SetTokenScope sets TokenScope field to given value.

### HasTokenScope

`func (o *AzureIntegrationCredential) HasTokenScope() bool`

HasTokenScope returns a boolean if a field has been set.

### GetTokenType

`func (o *AzureIntegrationCredential) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AzureIntegrationCredential) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AzureIntegrationCredential) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *AzureIntegrationCredential) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


