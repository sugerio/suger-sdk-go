# ClientSignupPageConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcEmails** | Pointer to **[]string** | The cc email contacts of the new client signup notification. | [optional] 
**CompanyLogoUrl** | Pointer to **string** | The company logo url of the seller/ISV. | [optional] 
**CompanyName** | Pointer to **string** | The company name of the seller/ISV to show in the client signup page. | [optional] 
**LandingImageUrl** | Pointer to **string** | The signup landing image aws url of the organization | [optional] 
**NotificationEmailSubject** | Pointer to **string** | The email subject of the new client signup notification. | [optional] 
**PublicNotes** | Pointer to **string** | The public notes to append in new client signup notification email. | [optional] 
**SignupId** | Pointer to **string** | The signup ID for the new client signup page url. It is populated by Suger service when creating the new client signup page. | [optional] 
**VideoLink** | Pointer to **string** | The video link of the product. Optional. | [optional] 

## Methods

### NewClientSignupPageConfigInfo

`func NewClientSignupPageConfigInfo() *ClientSignupPageConfigInfo`

NewClientSignupPageConfigInfo instantiates a new ClientSignupPageConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientSignupPageConfigInfoWithDefaults

`func NewClientSignupPageConfigInfoWithDefaults() *ClientSignupPageConfigInfo`

NewClientSignupPageConfigInfoWithDefaults instantiates a new ClientSignupPageConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCcEmails

`func (o *ClientSignupPageConfigInfo) GetCcEmails() []string`

GetCcEmails returns the CcEmails field if non-nil, zero value otherwise.

### GetCcEmailsOk

`func (o *ClientSignupPageConfigInfo) GetCcEmailsOk() (*[]string, bool)`

GetCcEmailsOk returns a tuple with the CcEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcEmails

`func (o *ClientSignupPageConfigInfo) SetCcEmails(v []string)`

SetCcEmails sets CcEmails field to given value.

### HasCcEmails

`func (o *ClientSignupPageConfigInfo) HasCcEmails() bool`

HasCcEmails returns a boolean if a field has been set.

### GetCompanyLogoUrl

`func (o *ClientSignupPageConfigInfo) GetCompanyLogoUrl() string`

GetCompanyLogoUrl returns the CompanyLogoUrl field if non-nil, zero value otherwise.

### GetCompanyLogoUrlOk

`func (o *ClientSignupPageConfigInfo) GetCompanyLogoUrlOk() (*string, bool)`

GetCompanyLogoUrlOk returns a tuple with the CompanyLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLogoUrl

`func (o *ClientSignupPageConfigInfo) SetCompanyLogoUrl(v string)`

SetCompanyLogoUrl sets CompanyLogoUrl field to given value.

### HasCompanyLogoUrl

`func (o *ClientSignupPageConfigInfo) HasCompanyLogoUrl() bool`

HasCompanyLogoUrl returns a boolean if a field has been set.

### GetCompanyName

`func (o *ClientSignupPageConfigInfo) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *ClientSignupPageConfigInfo) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *ClientSignupPageConfigInfo) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *ClientSignupPageConfigInfo) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetLandingImageUrl

`func (o *ClientSignupPageConfigInfo) GetLandingImageUrl() string`

GetLandingImageUrl returns the LandingImageUrl field if non-nil, zero value otherwise.

### GetLandingImageUrlOk

`func (o *ClientSignupPageConfigInfo) GetLandingImageUrlOk() (*string, bool)`

GetLandingImageUrlOk returns a tuple with the LandingImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingImageUrl

`func (o *ClientSignupPageConfigInfo) SetLandingImageUrl(v string)`

SetLandingImageUrl sets LandingImageUrl field to given value.

### HasLandingImageUrl

`func (o *ClientSignupPageConfigInfo) HasLandingImageUrl() bool`

HasLandingImageUrl returns a boolean if a field has been set.

### GetNotificationEmailSubject

`func (o *ClientSignupPageConfigInfo) GetNotificationEmailSubject() string`

GetNotificationEmailSubject returns the NotificationEmailSubject field if non-nil, zero value otherwise.

### GetNotificationEmailSubjectOk

`func (o *ClientSignupPageConfigInfo) GetNotificationEmailSubjectOk() (*string, bool)`

GetNotificationEmailSubjectOk returns a tuple with the NotificationEmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationEmailSubject

`func (o *ClientSignupPageConfigInfo) SetNotificationEmailSubject(v string)`

SetNotificationEmailSubject sets NotificationEmailSubject field to given value.

### HasNotificationEmailSubject

`func (o *ClientSignupPageConfigInfo) HasNotificationEmailSubject() bool`

HasNotificationEmailSubject returns a boolean if a field has been set.

### GetPublicNotes

`func (o *ClientSignupPageConfigInfo) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *ClientSignupPageConfigInfo) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *ClientSignupPageConfigInfo) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *ClientSignupPageConfigInfo) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetSignupId

`func (o *ClientSignupPageConfigInfo) GetSignupId() string`

GetSignupId returns the SignupId field if non-nil, zero value otherwise.

### GetSignupIdOk

`func (o *ClientSignupPageConfigInfo) GetSignupIdOk() (*string, bool)`

GetSignupIdOk returns a tuple with the SignupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupId

`func (o *ClientSignupPageConfigInfo) SetSignupId(v string)`

SetSignupId sets SignupId field to given value.

### HasSignupId

`func (o *ClientSignupPageConfigInfo) HasSignupId() bool`

HasSignupId returns a boolean if a field has been set.

### GetVideoLink

`func (o *ClientSignupPageConfigInfo) GetVideoLink() string`

GetVideoLink returns the VideoLink field if non-nil, zero value otherwise.

### GetVideoLinkOk

`func (o *ClientSignupPageConfigInfo) GetVideoLinkOk() (*string, bool)`

GetVideoLinkOk returns a tuple with the VideoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLink

`func (o *ClientSignupPageConfigInfo) SetVideoLink(v string)`

SetVideoLink sets VideoLink field to given value.

### HasVideoLink

`func (o *ClientSignupPageConfigInfo) HasVideoLink() bool`

HasVideoLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


