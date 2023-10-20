# SlackIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**AuthedUser** | Pointer to [**SlackOAuthV2ResponseAuthedUser**](SlackOAuthV2ResponseAuthedUser.md) |  | [optional] 
**BotUserId** | Pointer to **string** |  | [optional] 
**Enterprise** | Pointer to [**SlackOAuthV2ResponseEnterprise**](SlackOAuthV2ResponseEnterprise.md) |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**IncomingWebhook** | Pointer to [**SlackOAuthResponseIncomingWebhook**](SlackOAuthResponseIncomingWebhook.md) |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | The scope of the access token. multiple scopes are separated by comma. | [optional] 
**Team** | Pointer to [**SlackOAuthV2ResponseTeam**](SlackOAuthV2ResponseTeam.md) |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 

## Methods

### NewSlackIntegration

`func NewSlackIntegration() *SlackIntegration`

NewSlackIntegration instantiates a new SlackIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackIntegrationWithDefaults

`func NewSlackIntegrationWithDefaults() *SlackIntegration`

NewSlackIntegrationWithDefaults instantiates a new SlackIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *SlackIntegration) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SlackIntegration) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SlackIntegration) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *SlackIntegration) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAppId

`func (o *SlackIntegration) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SlackIntegration) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SlackIntegration) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *SlackIntegration) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAuthedUser

`func (o *SlackIntegration) GetAuthedUser() SlackOAuthV2ResponseAuthedUser`

GetAuthedUser returns the AuthedUser field if non-nil, zero value otherwise.

### GetAuthedUserOk

`func (o *SlackIntegration) GetAuthedUserOk() (*SlackOAuthV2ResponseAuthedUser, bool)`

GetAuthedUserOk returns a tuple with the AuthedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthedUser

`func (o *SlackIntegration) SetAuthedUser(v SlackOAuthV2ResponseAuthedUser)`

SetAuthedUser sets AuthedUser field to given value.

### HasAuthedUser

`func (o *SlackIntegration) HasAuthedUser() bool`

HasAuthedUser returns a boolean if a field has been set.

### GetBotUserId

`func (o *SlackIntegration) GetBotUserId() string`

GetBotUserId returns the BotUserId field if non-nil, zero value otherwise.

### GetBotUserIdOk

`func (o *SlackIntegration) GetBotUserIdOk() (*string, bool)`

GetBotUserIdOk returns a tuple with the BotUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUserId

`func (o *SlackIntegration) SetBotUserId(v string)`

SetBotUserId sets BotUserId field to given value.

### HasBotUserId

`func (o *SlackIntegration) HasBotUserId() bool`

HasBotUserId returns a boolean if a field has been set.

### GetEnterprise

`func (o *SlackIntegration) GetEnterprise() SlackOAuthV2ResponseEnterprise`

GetEnterprise returns the Enterprise field if non-nil, zero value otherwise.

### GetEnterpriseOk

`func (o *SlackIntegration) GetEnterpriseOk() (*SlackOAuthV2ResponseEnterprise, bool)`

GetEnterpriseOk returns a tuple with the Enterprise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprise

`func (o *SlackIntegration) SetEnterprise(v SlackOAuthV2ResponseEnterprise)`

SetEnterprise sets Enterprise field to given value.

### HasEnterprise

`func (o *SlackIntegration) HasEnterprise() bool`

HasEnterprise returns a boolean if a field has been set.

### GetExpiresIn

`func (o *SlackIntegration) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *SlackIntegration) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *SlackIntegration) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *SlackIntegration) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIncomingWebhook

`func (o *SlackIntegration) GetIncomingWebhook() SlackOAuthResponseIncomingWebhook`

GetIncomingWebhook returns the IncomingWebhook field if non-nil, zero value otherwise.

### GetIncomingWebhookOk

`func (o *SlackIntegration) GetIncomingWebhookOk() (*SlackOAuthResponseIncomingWebhook, bool)`

GetIncomingWebhookOk returns a tuple with the IncomingWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingWebhook

`func (o *SlackIntegration) SetIncomingWebhook(v SlackOAuthResponseIncomingWebhook)`

SetIncomingWebhook sets IncomingWebhook field to given value.

### HasIncomingWebhook

`func (o *SlackIntegration) HasIncomingWebhook() bool`

HasIncomingWebhook returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SlackIntegration) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SlackIntegration) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SlackIntegration) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SlackIntegration) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetRefreshToken

`func (o *SlackIntegration) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *SlackIntegration) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *SlackIntegration) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *SlackIntegration) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetScope

`func (o *SlackIntegration) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SlackIntegration) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SlackIntegration) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SlackIntegration) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTeam

`func (o *SlackIntegration) GetTeam() SlackOAuthV2ResponseTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *SlackIntegration) GetTeamOk() (*SlackOAuthV2ResponseTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *SlackIntegration) SetTeam(v SlackOAuthV2ResponseTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *SlackIntegration) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTokenType

`func (o *SlackIntegration) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *SlackIntegration) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *SlackIntegration) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *SlackIntegration) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


