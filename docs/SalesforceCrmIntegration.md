# SalesforceCrmIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**SalesforceCrmCredential**](SalesforceCrmCredential.md) |  | [optional] 
**Filters** | Pointer to [**[]SalesforceSyncFilter**](SalesforceSyncFilter.md) |  | [optional] 
**InstanceUrl** | Pointer to **string** |  | [optional] 
**IsSandbox** | Pointer to **bool** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Subdomain** | Pointer to **string** | User defined when setting up the integration | [optional] 
**SugerAppInstalled** | Pointer to **bool** |  | [optional] 

## Methods

### NewSalesforceCrmIntegration

`func NewSalesforceCrmIntegration() *SalesforceCrmIntegration`

NewSalesforceCrmIntegration instantiates a new SalesforceCrmIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesforceCrmIntegrationWithDefaults

`func NewSalesforceCrmIntegrationWithDefaults() *SalesforceCrmIntegration`

NewSalesforceCrmIntegrationWithDefaults instantiates a new SalesforceCrmIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *SalesforceCrmIntegration) GetCredential() SalesforceCrmCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *SalesforceCrmIntegration) GetCredentialOk() (*SalesforceCrmCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *SalesforceCrmIntegration) SetCredential(v SalesforceCrmCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *SalesforceCrmIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetFilters

`func (o *SalesforceCrmIntegration) GetFilters() []SalesforceSyncFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SalesforceCrmIntegration) GetFiltersOk() (*[]SalesforceSyncFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SalesforceCrmIntegration) SetFilters(v []SalesforceSyncFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SalesforceCrmIntegration) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInstanceUrl

`func (o *SalesforceCrmIntegration) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *SalesforceCrmIntegration) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *SalesforceCrmIntegration) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.

### HasInstanceUrl

`func (o *SalesforceCrmIntegration) HasInstanceUrl() bool`

HasInstanceUrl returns a boolean if a field has been set.

### GetIsSandbox

`func (o *SalesforceCrmIntegration) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *SalesforceCrmIntegration) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *SalesforceCrmIntegration) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *SalesforceCrmIntegration) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.

### GetPaused

`func (o *SalesforceCrmIntegration) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *SalesforceCrmIntegration) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *SalesforceCrmIntegration) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *SalesforceCrmIntegration) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetSecretKey

`func (o *SalesforceCrmIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SalesforceCrmIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SalesforceCrmIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SalesforceCrmIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSubdomain

`func (o *SalesforceCrmIntegration) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *SalesforceCrmIntegration) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *SalesforceCrmIntegration) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *SalesforceCrmIntegration) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetSugerAppInstalled

`func (o *SalesforceCrmIntegration) GetSugerAppInstalled() bool`

GetSugerAppInstalled returns the SugerAppInstalled field if non-nil, zero value otherwise.

### GetSugerAppInstalledOk

`func (o *SalesforceCrmIntegration) GetSugerAppInstalledOk() (*bool, bool)`

GetSugerAppInstalledOk returns a tuple with the SugerAppInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSugerAppInstalled

`func (o *SalesforceCrmIntegration) SetSugerAppInstalled(v bool)`

SetSugerAppInstalled sets SugerAppInstalled field to given value.

### HasSugerAppInstalled

`func (o *SalesforceCrmIntegration) HasSugerAppInstalled() bool`

HasSugerAppInstalled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


