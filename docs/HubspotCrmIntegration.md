# HubspotCrmIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyFields** | Pointer to **[]string** |  | [optional] 
**ContactFields** | Pointer to **[]string** |  | [optional] 
**Credential** | Pointer to [**HubspotCrmCredential**](HubspotCrmCredential.md) |  | [optional] 
**DealFields** | Pointer to **[]string** |  | [optional] 
**Paused** | Pointer to **bool** | Paused means the integration is not syncing. | [optional] 
**PortalId** | Pointer to **int32** | Hubspot Account Id | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**SyncFilters** | Pointer to [**[]HubspotSyncFilter**](HubspotSyncFilter.md) | Can have at most 3 filters which will all be AND-ed. | [optional] 

## Methods

### NewHubspotCrmIntegration

`func NewHubspotCrmIntegration() *HubspotCrmIntegration`

NewHubspotCrmIntegration instantiates a new HubspotCrmIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubspotCrmIntegrationWithDefaults

`func NewHubspotCrmIntegrationWithDefaults() *HubspotCrmIntegration`

NewHubspotCrmIntegrationWithDefaults instantiates a new HubspotCrmIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanyFields

`func (o *HubspotCrmIntegration) GetCompanyFields() []string`

GetCompanyFields returns the CompanyFields field if non-nil, zero value otherwise.

### GetCompanyFieldsOk

`func (o *HubspotCrmIntegration) GetCompanyFieldsOk() (*[]string, bool)`

GetCompanyFieldsOk returns a tuple with the CompanyFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyFields

`func (o *HubspotCrmIntegration) SetCompanyFields(v []string)`

SetCompanyFields sets CompanyFields field to given value.

### HasCompanyFields

`func (o *HubspotCrmIntegration) HasCompanyFields() bool`

HasCompanyFields returns a boolean if a field has been set.

### GetContactFields

`func (o *HubspotCrmIntegration) GetContactFields() []string`

GetContactFields returns the ContactFields field if non-nil, zero value otherwise.

### GetContactFieldsOk

`func (o *HubspotCrmIntegration) GetContactFieldsOk() (*[]string, bool)`

GetContactFieldsOk returns a tuple with the ContactFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFields

`func (o *HubspotCrmIntegration) SetContactFields(v []string)`

SetContactFields sets ContactFields field to given value.

### HasContactFields

`func (o *HubspotCrmIntegration) HasContactFields() bool`

HasContactFields returns a boolean if a field has been set.

### GetCredential

`func (o *HubspotCrmIntegration) GetCredential() HubspotCrmCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *HubspotCrmIntegration) GetCredentialOk() (*HubspotCrmCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *HubspotCrmIntegration) SetCredential(v HubspotCrmCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *HubspotCrmIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDealFields

`func (o *HubspotCrmIntegration) GetDealFields() []string`

GetDealFields returns the DealFields field if non-nil, zero value otherwise.

### GetDealFieldsOk

`func (o *HubspotCrmIntegration) GetDealFieldsOk() (*[]string, bool)`

GetDealFieldsOk returns a tuple with the DealFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealFields

`func (o *HubspotCrmIntegration) SetDealFields(v []string)`

SetDealFields sets DealFields field to given value.

### HasDealFields

`func (o *HubspotCrmIntegration) HasDealFields() bool`

HasDealFields returns a boolean if a field has been set.

### GetPaused

`func (o *HubspotCrmIntegration) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *HubspotCrmIntegration) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *HubspotCrmIntegration) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *HubspotCrmIntegration) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetPortalId

`func (o *HubspotCrmIntegration) GetPortalId() int32`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *HubspotCrmIntegration) GetPortalIdOk() (*int32, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *HubspotCrmIntegration) SetPortalId(v int32)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *HubspotCrmIntegration) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### GetSecretKey

`func (o *HubspotCrmIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *HubspotCrmIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *HubspotCrmIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *HubspotCrmIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetSyncFilters

`func (o *HubspotCrmIntegration) GetSyncFilters() []HubspotSyncFilter`

GetSyncFilters returns the SyncFilters field if non-nil, zero value otherwise.

### GetSyncFiltersOk

`func (o *HubspotCrmIntegration) GetSyncFiltersOk() (*[]HubspotSyncFilter, bool)`

GetSyncFiltersOk returns a tuple with the SyncFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncFilters

`func (o *HubspotCrmIntegration) SetSyncFilters(v []HubspotSyncFilter)`

SetSyncFilters sets SyncFilters field to given value.

### HasSyncFilters

`func (o *HubspotCrmIntegration) HasSyncFilters() bool`

HasSyncFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


