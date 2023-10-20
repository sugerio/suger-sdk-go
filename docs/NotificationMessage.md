# NotificationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**NotificationMessageInfo**](NotificationMessageInfo.md) |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**NotificationType**](NotificationType.md) |  | [optional] 

## Methods

### NewNotificationMessage

`func NewNotificationMessage() *NotificationMessage`

NewNotificationMessage instantiates a new NotificationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMessageWithDefaults

`func NewNotificationMessageWithDefaults() *NotificationMessage`

NewNotificationMessageWithDefaults instantiates a new NotificationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *NotificationMessage) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NotificationMessage) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NotificationMessage) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *NotificationMessage) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *NotificationMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *NotificationMessage) GetInfo() NotificationMessageInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *NotificationMessage) GetInfoOk() (*NotificationMessageInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *NotificationMessage) SetInfo(v NotificationMessageInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *NotificationMessage) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetOrganizationID

`func (o *NotificationMessage) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *NotificationMessage) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *NotificationMessage) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *NotificationMessage) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetRecipient

`func (o *NotificationMessage) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NotificationMessage) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NotificationMessage) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NotificationMessage) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetType

`func (o *NotificationMessage) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationMessage) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationMessage) SetType(v NotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationMessage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


