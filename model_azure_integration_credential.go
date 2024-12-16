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

// checks if the AzureIntegrationCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureIntegrationCredential{}

// AzureIntegrationCredential struct for AzureIntegrationCredential
type AzureIntegrationCredential struct {
	AccessToken  *string `json:"accessToken,omitempty"`
	ClientID     *string `json:"clientID,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The time when the access token expires.
	ExpiresOn *string `json:"expiresOn,omitempty"`
	// The refresh token used to refresh the access token.
	RefreshToken *string `json:"refreshToken,omitempty"`
	TenantID     *string `json:"tenantID,omitempty"`
	TokenScope   *string `json:"tokenScope,omitempty"`
	TokenType    *string `json:"tokenType,omitempty"`
}

// NewAzureIntegrationCredential instantiates a new AzureIntegrationCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureIntegrationCredential() *AzureIntegrationCredential {
	this := AzureIntegrationCredential{}
	return &this
}

// NewAzureIntegrationCredentialWithDefaults instantiates a new AzureIntegrationCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureIntegrationCredentialWithDefaults() *AzureIntegrationCredential {
	this := AzureIntegrationCredential{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AzureIntegrationCredential) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetClientID returns the ClientID field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetClientID() string {
	if o == nil || IsNil(o.ClientID) {
		var ret string
		return ret
	}
	return *o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetClientIDOk() (*string, bool) {
	if o == nil || IsNil(o.ClientID) {
		return nil, false
	}
	return o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasClientID() bool {
	if o != nil && !IsNil(o.ClientID) {
		return true
	}

	return false
}

// SetClientID gets a reference to the given string and assigns it to the ClientID field.
func (o *AzureIntegrationCredential) SetClientID(v string) {
	o.ClientID = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AzureIntegrationCredential) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetExpiresOn() string {
	if o == nil || IsNil(o.ExpiresOn) {
		var ret string
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetExpiresOnOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresOn) {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasExpiresOn() bool {
	if o != nil && !IsNil(o.ExpiresOn) {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given string and assigns it to the ExpiresOn field.
func (o *AzureIntegrationCredential) SetExpiresOn(v string) {
	o.ExpiresOn = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AzureIntegrationCredential) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetTenantID returns the TenantID field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetTenantID() string {
	if o == nil || IsNil(o.TenantID) {
		var ret string
		return ret
	}
	return *o.TenantID
}

// GetTenantIDOk returns a tuple with the TenantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetTenantIDOk() (*string, bool) {
	if o == nil || IsNil(o.TenantID) {
		return nil, false
	}
	return o.TenantID, true
}

// HasTenantID returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasTenantID() bool {
	if o != nil && !IsNil(o.TenantID) {
		return true
	}

	return false
}

// SetTenantID gets a reference to the given string and assigns it to the TenantID field.
func (o *AzureIntegrationCredential) SetTenantID(v string) {
	o.TenantID = &v
}

// GetTokenScope returns the TokenScope field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetTokenScope() string {
	if o == nil || IsNil(o.TokenScope) {
		var ret string
		return ret
	}
	return *o.TokenScope
}

// GetTokenScopeOk returns a tuple with the TokenScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetTokenScopeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenScope) {
		return nil, false
	}
	return o.TokenScope, true
}

// HasTokenScope returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasTokenScope() bool {
	if o != nil && !IsNil(o.TokenScope) {
		return true
	}

	return false
}

// SetTokenScope gets a reference to the given string and assigns it to the TokenScope field.
func (o *AzureIntegrationCredential) SetTokenScope(v string) {
	o.TokenScope = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *AzureIntegrationCredential) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureIntegrationCredential) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *AzureIntegrationCredential) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *AzureIntegrationCredential) SetTokenType(v string) {
	o.TokenType = &v
}

func (o AzureIntegrationCredential) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureIntegrationCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["accessToken"] = o.AccessToken
	}
	if !IsNil(o.ClientID) {
		toSerialize["clientID"] = o.ClientID
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.ExpiresOn) {
		toSerialize["expiresOn"] = o.ExpiresOn
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refreshToken"] = o.RefreshToken
	}
	if !IsNil(o.TenantID) {
		toSerialize["tenantID"] = o.TenantID
	}
	if !IsNil(o.TokenScope) {
		toSerialize["tokenScope"] = o.TokenScope
	}
	if !IsNil(o.TokenType) {
		toSerialize["tokenType"] = o.TokenType
	}
	return toSerialize, nil
}

type NullableAzureIntegrationCredential struct {
	value *AzureIntegrationCredential
	isSet bool
}

func (v NullableAzureIntegrationCredential) Get() *AzureIntegrationCredential {
	return v.value
}

func (v *NullableAzureIntegrationCredential) Set(val *AzureIntegrationCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureIntegrationCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureIntegrationCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureIntegrationCredential(val *AzureIntegrationCredential) *NullableAzureIntegrationCredential {
	return &NullableAzureIntegrationCredential{value: val, isSet: true}
}

func (v NullableAzureIntegrationCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureIntegrationCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
