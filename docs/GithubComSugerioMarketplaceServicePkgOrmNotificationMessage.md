# GithubComSugerioMarketplaceServicePkgOrmNotificationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** | CreationTime holds the value of the \&quot;creation_time\&quot; field. | [optional] 
**Id** | Pointer to **string** | ID of the ent. | [optional] 
**Info** | Pointer to [**NotificationMessageInfo**](NotificationMessageInfo.md) | Info holds the value of the \&quot;info\&quot; field. | [optional] 
**OrganizationID** | Pointer to **string** | OrganizationID holds the value of the \&quot;organization_id\&quot; field. | [optional] 
**Recipient** | Pointer to **string** | supports multiple recipients as a comma-separated string | [optional] 
**Type** | Pointer to [**GithubComSugerioMarketplaceServicePkgOrmNotificationmessageType**](GithubComSugerioMarketplaceServicePkgOrmNotificationmessageType.md) | Type holds the value of the \&quot;type\&quot; field. | [optional] 

## Methods

### NewGithubComSugerioMarketplaceServicePkgOrmNotificationMessage

`func NewGithubComSugerioMarketplaceServicePkgOrmNotificationMessage() *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage`

NewGithubComSugerioMarketplaceServicePkgOrmNotificationMessage instantiates a new GithubComSugerioMarketplaceServicePkgOrmNotificationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGithubComSugerioMarketplaceServicePkgOrmNotificationMessageWithDefaults

`func NewGithubComSugerioMarketplaceServicePkgOrmNotificationMessageWithDefaults() *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage`

NewGithubComSugerioMarketplaceServicePkgOrmNotificationMessageWithDefaults instantiates a new GithubComSugerioMarketplaceServicePkgOrmNotificationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetInfo() NotificationMessageInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetInfoOk() (*NotificationMessageInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) SetInfo(v NotificationMessageInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetRecipient

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetType

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetType() GithubComSugerioMarketplaceServicePkgOrmNotificationmessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) GetTypeOk() (*GithubComSugerioMarketplaceServicePkgOrmNotificationmessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) SetType(v GithubComSugerioMarketplaceServicePkgOrmNotificationmessageType)`

SetType sets Type field to given value.

### HasType

`func (o *GithubComSugerioMarketplaceServicePkgOrmNotificationMessage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


