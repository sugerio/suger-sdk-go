/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suger

import (
	"encoding/json"
)

// checks if the SlackIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackIntegration{}

// SlackIntegration struct for SlackIntegration
type SlackIntegration struct {
	AccessToken     *string                            `json:"accessToken,omitempty"`
	AppId           *string                            `json:"appId,omitempty"`
	AuthedUser      *SlackOAuthV2ResponseAuthedUser    `json:"authedUser,omitempty"`
	BotUserId       *string                            `json:"botUserId,omitempty"`
	Enterprise      *SlackOAuthV2ResponseEnterprise    `json:"enterprise,omitempty"`
	ExpiresIn       *int32                             `json:"expiresIn,omitempty"`
	IncomingWebhook *SlackOAuthResponseIncomingWebhook `json:"incomingWebhook,omitempty"`
	RedirectUrl     *string                            `json:"redirectUrl,omitempty"`
	RefreshToken    *string                            `json:"refreshToken,omitempty"`
	// The scope of the access token. multiple scopes are separated by comma.
	Scope     *string                   `json:"scope,omitempty"`
	Team      *SlackOAuthV2ResponseTeam `json:"team,omitempty"`
	TokenType *string                   `json:"tokenType,omitempty"`
}

// NewSlackIntegration instantiates a new SlackIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackIntegration() *SlackIntegration {
	this := SlackIntegration{}
	return &this
}

// NewSlackIntegrationWithDefaults instantiates a new SlackIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackIntegrationWithDefaults() *SlackIntegration {
	this := SlackIntegration{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *SlackIntegration) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *SlackIntegration) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *SlackIntegration) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *SlackIntegration) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *SlackIntegration) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *SlackIntegration) SetAppId(v string) {
	o.AppId = &v
}

// GetAuthedUser returns the AuthedUser field value if set, zero value otherwise.
func (o *SlackIntegration) GetAuthedUser() SlackOAuthV2ResponseAuthedUser {
	if o == nil || IsNil(o.AuthedUser) {
		var ret SlackOAuthV2ResponseAuthedUser
		return ret
	}
	return *o.AuthedUser
}

// GetAuthedUserOk returns a tuple with the AuthedUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetAuthedUserOk() (*SlackOAuthV2ResponseAuthedUser, bool) {
	if o == nil || IsNil(o.AuthedUser) {
		return nil, false
	}
	return o.AuthedUser, true
}

// HasAuthedUser returns a boolean if a field has been set.
func (o *SlackIntegration) HasAuthedUser() bool {
	if o != nil && !IsNil(o.AuthedUser) {
		return true
	}

	return false
}

// SetAuthedUser gets a reference to the given SlackOAuthV2ResponseAuthedUser and assigns it to the AuthedUser field.
func (o *SlackIntegration) SetAuthedUser(v SlackOAuthV2ResponseAuthedUser) {
	o.AuthedUser = &v
}

// GetBotUserId returns the BotUserId field value if set, zero value otherwise.
func (o *SlackIntegration) GetBotUserId() string {
	if o == nil || IsNil(o.BotUserId) {
		var ret string
		return ret
	}
	return *o.BotUserId
}

// GetBotUserIdOk returns a tuple with the BotUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetBotUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.BotUserId) {
		return nil, false
	}
	return o.BotUserId, true
}

// HasBotUserId returns a boolean if a field has been set.
func (o *SlackIntegration) HasBotUserId() bool {
	if o != nil && !IsNil(o.BotUserId) {
		return true
	}

	return false
}

// SetBotUserId gets a reference to the given string and assigns it to the BotUserId field.
func (o *SlackIntegration) SetBotUserId(v string) {
	o.BotUserId = &v
}

// GetEnterprise returns the Enterprise field value if set, zero value otherwise.
func (o *SlackIntegration) GetEnterprise() SlackOAuthV2ResponseEnterprise {
	if o == nil || IsNil(o.Enterprise) {
		var ret SlackOAuthV2ResponseEnterprise
		return ret
	}
	return *o.Enterprise
}

// GetEnterpriseOk returns a tuple with the Enterprise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetEnterpriseOk() (*SlackOAuthV2ResponseEnterprise, bool) {
	if o == nil || IsNil(o.Enterprise) {
		return nil, false
	}
	return o.Enterprise, true
}

// HasEnterprise returns a boolean if a field has been set.
func (o *SlackIntegration) HasEnterprise() bool {
	if o != nil && !IsNil(o.Enterprise) {
		return true
	}

	return false
}

// SetEnterprise gets a reference to the given SlackOAuthV2ResponseEnterprise and assigns it to the Enterprise field.
func (o *SlackIntegration) SetEnterprise(v SlackOAuthV2ResponseEnterprise) {
	o.Enterprise = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *SlackIntegration) GetExpiresIn() int32 {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetExpiresInOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *SlackIntegration) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *SlackIntegration) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetIncomingWebhook returns the IncomingWebhook field value if set, zero value otherwise.
func (o *SlackIntegration) GetIncomingWebhook() SlackOAuthResponseIncomingWebhook {
	if o == nil || IsNil(o.IncomingWebhook) {
		var ret SlackOAuthResponseIncomingWebhook
		return ret
	}
	return *o.IncomingWebhook
}

// GetIncomingWebhookOk returns a tuple with the IncomingWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetIncomingWebhookOk() (*SlackOAuthResponseIncomingWebhook, bool) {
	if o == nil || IsNil(o.IncomingWebhook) {
		return nil, false
	}
	return o.IncomingWebhook, true
}

// HasIncomingWebhook returns a boolean if a field has been set.
func (o *SlackIntegration) HasIncomingWebhook() bool {
	if o != nil && !IsNil(o.IncomingWebhook) {
		return true
	}

	return false
}

// SetIncomingWebhook gets a reference to the given SlackOAuthResponseIncomingWebhook and assigns it to the IncomingWebhook field.
func (o *SlackIntegration) SetIncomingWebhook(v SlackOAuthResponseIncomingWebhook) {
	o.IncomingWebhook = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *SlackIntegration) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *SlackIntegration) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *SlackIntegration) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *SlackIntegration) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *SlackIntegration) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *SlackIntegration) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SlackIntegration) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SlackIntegration) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SlackIntegration) SetScope(v string) {
	o.Scope = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *SlackIntegration) GetTeam() SlackOAuthV2ResponseTeam {
	if o == nil || IsNil(o.Team) {
		var ret SlackOAuthV2ResponseTeam
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetTeamOk() (*SlackOAuthV2ResponseTeam, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *SlackIntegration) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given SlackOAuthV2ResponseTeam and assigns it to the Team field.
func (o *SlackIntegration) SetTeam(v SlackOAuthV2ResponseTeam) {
	o.Team = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *SlackIntegration) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegration) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *SlackIntegration) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *SlackIntegration) SetTokenType(v string) {
	o.TokenType = &v
}

func (o SlackIntegration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.AuthedUser) {
		toSerialize["authedUser"] = o.AuthedUser
	}
	if !IsNil(o.BotUserId) {
		toSerialize["botUserId"] = o.BotUserId
	}
	if !IsNil(o.Enterprise) {
		toSerialize["enterprise"] = o.Enterprise
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if !IsNil(o.IncomingWebhook) {
		toSerialize["incomingWebhook"] = o.IncomingWebhook
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.TokenType) {
		toSerialize["tokenType"] = o.TokenType
	}
	return toSerialize, nil
}

type NullableSlackIntegration struct {
	value *SlackIntegration
	isSet bool
}

func (v NullableSlackIntegration) Get() *SlackIntegration {
	return v.value
}

func (v *NullableSlackIntegration) Set(val *SlackIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackIntegration(val *SlackIntegration) *NullableSlackIntegration {
	return &NullableSlackIntegration{value: val, isSet: true}
}

func (v NullableSlackIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
