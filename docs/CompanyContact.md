# CompanyContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** |  | [optional] 
**CompanyDomain** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** | Company information | [optional] 
**Country** | Pointer to **string** | Location information | [optional] 
**CreationTime** | Pointer to **string** | Timestamps | [optional] 
**Email** | Pointer to **string** | Contact information | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Core identification fields | [optional] 
**Info** | Pointer to [**CompanyContactInfo**](CompanyContactInfo.md) | Additional structured information | [optional] 
**JobFunction** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** | Professional information | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LastUpdateTime** | Pointer to **string** |  | [optional] 
**LinkedIn** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**S3KeyPicture** | Pointer to **string** | Media and social | [optional] 
**Status** | Pointer to [**EnrichmentDataStatus**](EnrichmentDataStatus.md) | active/stale/archived | [optional] 
**Twitter** | Pointer to **string** |  | [optional] 

## Methods

### NewCompanyContact

`func NewCompanyContact() *CompanyContact`

NewCompanyContact instantiates a new CompanyContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyContactWithDefaults

`func NewCompanyContactWithDefaults() *CompanyContact`

NewCompanyContactWithDefaults instantiates a new CompanyContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *CompanyContact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CompanyContact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CompanyContact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CompanyContact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyDomain

`func (o *CompanyContact) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *CompanyContact) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *CompanyContact) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.

### HasCompanyDomain

`func (o *CompanyContact) HasCompanyDomain() bool`

HasCompanyDomain returns a boolean if a field has been set.

### GetCompanyName

`func (o *CompanyContact) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CompanyContact) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CompanyContact) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CompanyContact) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *CompanyContact) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CompanyContact) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CompanyContact) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CompanyContact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreationTime

`func (o *CompanyContact) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *CompanyContact) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *CompanyContact) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *CompanyContact) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetEmail

`func (o *CompanyContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CompanyContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CompanyContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CompanyContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *CompanyContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CompanyContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CompanyContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CompanyContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetId

`func (o *CompanyContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyContact) GetInfo() CompanyContactInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyContact) GetInfoOk() (*CompanyContactInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyContact) SetInfo(v CompanyContactInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyContact) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetJobFunction

`func (o *CompanyContact) GetJobFunction() string`

GetJobFunction returns the JobFunction field if non-nil, zero value otherwise.

### GetJobFunctionOk

`func (o *CompanyContact) GetJobFunctionOk() (*string, bool)`

GetJobFunctionOk returns a tuple with the JobFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFunction

`func (o *CompanyContact) SetJobFunction(v string)`

SetJobFunction sets JobFunction field to given value.

### HasJobFunction

`func (o *CompanyContact) HasJobFunction() bool`

HasJobFunction returns a boolean if a field has been set.

### GetJobTitle

`func (o *CompanyContact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CompanyContact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *CompanyContact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *CompanyContact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastName

`func (o *CompanyContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CompanyContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CompanyContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CompanyContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *CompanyContact) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *CompanyContact) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *CompanyContact) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *CompanyContact) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetLinkedIn

`func (o *CompanyContact) GetLinkedIn() string`

GetLinkedIn returns the LinkedIn field if non-nil, zero value otherwise.

### GetLinkedInOk

`func (o *CompanyContact) GetLinkedInOk() (*string, bool)`

GetLinkedInOk returns a tuple with the LinkedIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedIn

`func (o *CompanyContact) SetLinkedIn(v string)`

SetLinkedIn sets LinkedIn field to given value.

### HasLinkedIn

`func (o *CompanyContact) HasLinkedIn() bool`

HasLinkedIn returns a boolean if a field has been set.

### GetName

`func (o *CompanyContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *CompanyContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetS3KeyPicture

`func (o *CompanyContact) GetS3KeyPicture() string`

GetS3KeyPicture returns the S3KeyPicture field if non-nil, zero value otherwise.

### GetS3KeyPictureOk

`func (o *CompanyContact) GetS3KeyPictureOk() (*string, bool)`

GetS3KeyPictureOk returns a tuple with the S3KeyPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3KeyPicture

`func (o *CompanyContact) SetS3KeyPicture(v string)`

SetS3KeyPicture sets S3KeyPicture field to given value.

### HasS3KeyPicture

`func (o *CompanyContact) HasS3KeyPicture() bool`

HasS3KeyPicture returns a boolean if a field has been set.

### GetStatus

`func (o *CompanyContact) GetStatus() EnrichmentDataStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyContact) GetStatusOk() (*EnrichmentDataStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyContact) SetStatus(v EnrichmentDataStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyContact) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTwitter

`func (o *CompanyContact) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *CompanyContact) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *CompanyContact) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *CompanyContact) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


