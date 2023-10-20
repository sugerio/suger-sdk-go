# GetApiClientAccessTokenParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the API Client. | 
**OrganizationID** | **string** |  | 
**Secret** | **string** | The secret of the API Client. | 

## Methods

### NewGetApiClientAccessTokenParams

`func NewGetApiClientAccessTokenParams(id string, organizationID string, secret string, ) *GetApiClientAccessTokenParams`

NewGetApiClientAccessTokenParams instantiates a new GetApiClientAccessTokenParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApiClientAccessTokenParamsWithDefaults

`func NewGetApiClientAccessTokenParamsWithDefaults() *GetApiClientAccessTokenParams`

NewGetApiClientAccessTokenParamsWithDefaults instantiates a new GetApiClientAccessTokenParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetApiClientAccessTokenParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetApiClientAccessTokenParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetApiClientAccessTokenParams) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationID

`func (o *GetApiClientAccessTokenParams) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GetApiClientAccessTokenParams) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GetApiClientAccessTokenParams) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.


### GetSecret

`func (o *GetApiClientAccessTokenParams) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetApiClientAccessTokenParams) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetApiClientAccessTokenParams) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


