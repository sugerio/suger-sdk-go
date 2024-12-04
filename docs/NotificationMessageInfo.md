# NotificationMessageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**NotificationEventAction**](NotificationEventAction.md) | The action of this notification message. | [optional] 
**CcRecipients** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | All other fields | [optional] 
**HtmlContent** | Pointer to **string** | The HTML content of the email. | [optional] 
**RccRecipients** | Pointer to **[]string** |  | [optional] 
**StandardFields** | Pointer to **map[string]interface{}** | The standard fields to render the email content. | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**TextContent** | Pointer to **string** | The text content of the email in case the recipient&#39;s email client does not support HTML. | [optional] 

## Methods

### NewNotificationMessageInfo

`func NewNotificationMessageInfo() *NotificationMessageInfo`

NewNotificationMessageInfo instantiates a new NotificationMessageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationMessageInfoWithDefaults

`func NewNotificationMessageInfoWithDefaults() *NotificationMessageInfo`

NewNotificationMessageInfoWithDefaults instantiates a new NotificationMessageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NotificationMessageInfo) GetAction() NotificationEventAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NotificationMessageInfo) GetActionOk() (*NotificationEventAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NotificationMessageInfo) SetAction(v NotificationEventAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *NotificationMessageInfo) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCcRecipients

`func (o *NotificationMessageInfo) GetCcRecipients() []string`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *NotificationMessageInfo) GetCcRecipientsOk() (*[]string, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipients

`func (o *NotificationMessageInfo) SetCcRecipients(v []string)`

SetCcRecipients sets CcRecipients field to given value.

### HasCcRecipients

`func (o *NotificationMessageInfo) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### GetCustomFields

`func (o *NotificationMessageInfo) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *NotificationMessageInfo) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *NotificationMessageInfo) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *NotificationMessageInfo) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetHtmlContent

`func (o *NotificationMessageInfo) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *NotificationMessageInfo) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *NotificationMessageInfo) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *NotificationMessageInfo) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### GetRccRecipients

`func (o *NotificationMessageInfo) GetRccRecipients() []string`

GetRccRecipients returns the RccRecipients field if non-nil, zero value otherwise.

### GetRccRecipientsOk

`func (o *NotificationMessageInfo) GetRccRecipientsOk() (*[]string, bool)`

GetRccRecipientsOk returns a tuple with the RccRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRccRecipients

`func (o *NotificationMessageInfo) SetRccRecipients(v []string)`

SetRccRecipients sets RccRecipients field to given value.

### HasRccRecipients

`func (o *NotificationMessageInfo) HasRccRecipients() bool`

HasRccRecipients returns a boolean if a field has been set.

### GetStandardFields

`func (o *NotificationMessageInfo) GetStandardFields() map[string]interface{}`

GetStandardFields returns the StandardFields field if non-nil, zero value otherwise.

### GetStandardFieldsOk

`func (o *NotificationMessageInfo) GetStandardFieldsOk() (*map[string]interface{}, bool)`

GetStandardFieldsOk returns a tuple with the StandardFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardFields

`func (o *NotificationMessageInfo) SetStandardFields(v map[string]interface{})`

SetStandardFields sets StandardFields field to given value.

### HasStandardFields

`func (o *NotificationMessageInfo) HasStandardFields() bool`

HasStandardFields returns a boolean if a field has been set.

### GetSubject

`func (o *NotificationMessageInfo) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *NotificationMessageInfo) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *NotificationMessageInfo) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *NotificationMessageInfo) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTextContent

`func (o *NotificationMessageInfo) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *NotificationMessageInfo) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *NotificationMessageInfo) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *NotificationMessageInfo) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


