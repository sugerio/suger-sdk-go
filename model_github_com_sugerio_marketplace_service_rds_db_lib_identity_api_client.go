/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient{}

// GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient struct for GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient
type GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient struct {
	ApiKeyHash *string `json:"apiKeyHash,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
	Id *string `json:"id,omitempty"`
	Info *string `json:"info,omitempty"`
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`
	OrganizationID *string `json:"organizationID,omitempty"`
	Provider *string `json:"provider,omitempty"`
	Role *string `json:"role,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient instantiates a new GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient() *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient {
	this := GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient{}
	return &this
}

// NewGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClientWithDefaults instantiates a new GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClientWithDefaults() *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient {
	this := GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient{}
	return &this
}

// GetApiKeyHash returns the ApiKeyHash field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetApiKeyHash() string {
	if o == nil || IsNil(o.ApiKeyHash) {
		var ret string
		return ret
	}
	return *o.ApiKeyHash
}

// GetApiKeyHashOk returns a tuple with the ApiKeyHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetApiKeyHashOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyHash) {
		return nil, false
	}
	return o.ApiKeyHash, true
}

// HasApiKeyHash returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasApiKeyHash() bool {
	if o != nil && !IsNil(o.ApiKeyHash) {
		return true
	}

	return false
}

// SetApiKeyHash gets a reference to the given string and assigns it to the ApiKeyHash field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetApiKeyHash(v string) {
	o.ApiKeyHash = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime) {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetCreationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetInfo() string {
	if o == nil || IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetInfoOk() (*string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetInfo(v string) {
	o.Info = &v
}

// GetLastUpdateTime returns the LastUpdateTime field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetLastUpdateTime() string {
	if o == nil || IsNil(o.LastUpdateTime) {
		var ret string
		return ret
	}
	return *o.LastUpdateTime
}

// GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetLastUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdateTime) {
		return nil, false
	}
	return o.LastUpdateTime, true
}

// HasLastUpdateTime returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasLastUpdateTime() bool {
	if o != nil && !IsNil(o.LastUpdateTime) {
		return true
	}

	return false
}

// SetLastUpdateTime gets a reference to the given string and assigns it to the LastUpdateTime field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetLastUpdateTime(v string) {
	o.LastUpdateTime = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetProvider(v string) {
	o.Provider = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetRole(v string) {
	o.Role = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetSecret(v string) {
	o.Secret = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) SetType(v string) {
	o.Type = &v
}

func (o GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKeyHash) {
		toSerialize["apiKeyHash"] = o.ApiKeyHash
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.LastUpdateTime) {
		toSerialize["lastUpdateTime"] = o.LastUpdateTime
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient struct {
	value *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient
	isSet bool
}

func (v NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) Get() *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient {
	return v.value
}

func (v *NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) Set(val *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) {
	v.value = val
	v.isSet = true
}

func (v NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) IsSet() bool {
	return v.isSet
}

func (v *NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient(val *GithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) *NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient {
	return &NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient{value: val, isSet: true}
}

func (v NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGithubComSugerioMarketplaceServiceRdsDbLibIdentityApiClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


