# MicrosoftPartnerReferral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedDateTime** | Pointer to **string** |  | [optional] 
**CallToAction** | Pointer to **string** |  | [optional] 
**CampaignId** | Pointer to **string** |  | [optional] 
**ClosedDateTime** | Pointer to **string** |  | [optional] 
**Consent** | Pointer to [**MicrosoftPartnerReferralConsent**](MicrosoftPartnerReferralConsent.md) |  | [optional] 
**CreatedDateTime** | Pointer to **string** |  | [optional] 
**CreatedVia** | Pointer to **string** |  | [optional] 
**CustomerProfile** | Pointer to [**MicrosoftPartnerReferralProfile**](MicrosoftPartnerReferralProfile.md) |  | [optional] 
**DealSensitivity** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**MicrosoftPartnerReferralDetail**](MicrosoftPartnerReferralDetail.md) |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] 
**EngagementId** | Pointer to **string** |  | [optional] 
**Exception** | Pointer to **map[string]interface{}** |  | [optional] 
**ExpirationDateTime** | Pointer to **string** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InviteContext** | Pointer to [**MicrosoftPartnerReferralInviteContext**](MicrosoftPartnerReferralInviteContext.md) |  | [optional] 
**IsAutoRegistrationEnabled** | Pointer to **bool** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**LastModifiedVia** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**MicrosoftPartnerReferralLink**](MicrosoftPartnerReferralLink.md) |  | [optional] 
**MpnId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**Qualification** | Pointer to [**MicrosoftPartnerReferralQualification**](MicrosoftPartnerReferralQualification.md) |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**ReferralProgram** | Pointer to **map[string]interface{}** |  | [optional] 
**ReferralSource** | Pointer to **map[string]interface{}** |  | [optional] 
**RegistrationStatus** | Pointer to **string** |  | [optional] 
**Registrations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SalesStage** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**MicrosoftPartnerReferralStatus**](MicrosoftPartnerReferralStatus.md) |  | [optional] 
**Substatus** | Pointer to [**MicrosoftPartnerReferralSubStatus**](MicrosoftPartnerReferralSubStatus.md) |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Target** | Pointer to [**[]MicrosoftPartnerReferralTarget**](MicrosoftPartnerReferralTarget.md) |  | [optional] 
**Team** | Pointer to [**[]MicrosoftPartnerReferralPerson**](MicrosoftPartnerReferralPerson.md) |  | [optional] 
**TrackingInfo** | Pointer to [**MicrosoftPartnerReferralTrackingInfo**](MicrosoftPartnerReferralTrackingInfo.md) |  | [optional] 
**Type** | Pointer to [**MicrosoftPartnerReferralType**](MicrosoftPartnerReferralType.md) |  | [optional] 
**UpdatedDateTime** | Pointer to **string** |  | [optional] 

## Methods

### NewMicrosoftPartnerReferral

`func NewMicrosoftPartnerReferral() *MicrosoftPartnerReferral`

NewMicrosoftPartnerReferral instantiates a new MicrosoftPartnerReferral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftPartnerReferralWithDefaults

`func NewMicrosoftPartnerReferralWithDefaults() *MicrosoftPartnerReferral`

NewMicrosoftPartnerReferralWithDefaults instantiates a new MicrosoftPartnerReferral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedDateTime

`func (o *MicrosoftPartnerReferral) GetAcceptedDateTime() string`

GetAcceptedDateTime returns the AcceptedDateTime field if non-nil, zero value otherwise.

### GetAcceptedDateTimeOk

`func (o *MicrosoftPartnerReferral) GetAcceptedDateTimeOk() (*string, bool)`

GetAcceptedDateTimeOk returns a tuple with the AcceptedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedDateTime

`func (o *MicrosoftPartnerReferral) SetAcceptedDateTime(v string)`

SetAcceptedDateTime sets AcceptedDateTime field to given value.

### HasAcceptedDateTime

`func (o *MicrosoftPartnerReferral) HasAcceptedDateTime() bool`

HasAcceptedDateTime returns a boolean if a field has been set.

### GetCallToAction

`func (o *MicrosoftPartnerReferral) GetCallToAction() string`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *MicrosoftPartnerReferral) GetCallToActionOk() (*string, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *MicrosoftPartnerReferral) SetCallToAction(v string)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *MicrosoftPartnerReferral) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetCampaignId

`func (o *MicrosoftPartnerReferral) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *MicrosoftPartnerReferral) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *MicrosoftPartnerReferral) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *MicrosoftPartnerReferral) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetClosedDateTime

`func (o *MicrosoftPartnerReferral) GetClosedDateTime() string`

GetClosedDateTime returns the ClosedDateTime field if non-nil, zero value otherwise.

### GetClosedDateTimeOk

`func (o *MicrosoftPartnerReferral) GetClosedDateTimeOk() (*string, bool)`

GetClosedDateTimeOk returns a tuple with the ClosedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDateTime

`func (o *MicrosoftPartnerReferral) SetClosedDateTime(v string)`

SetClosedDateTime sets ClosedDateTime field to given value.

### HasClosedDateTime

`func (o *MicrosoftPartnerReferral) HasClosedDateTime() bool`

HasClosedDateTime returns a boolean if a field has been set.

### GetConsent

`func (o *MicrosoftPartnerReferral) GetConsent() MicrosoftPartnerReferralConsent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *MicrosoftPartnerReferral) GetConsentOk() (*MicrosoftPartnerReferralConsent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *MicrosoftPartnerReferral) SetConsent(v MicrosoftPartnerReferralConsent)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *MicrosoftPartnerReferral) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftPartnerReferral) GetCreatedDateTime() string`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftPartnerReferral) GetCreatedDateTimeOk() (*string, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftPartnerReferral) SetCreatedDateTime(v string)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftPartnerReferral) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetCreatedVia

`func (o *MicrosoftPartnerReferral) GetCreatedVia() string`

GetCreatedVia returns the CreatedVia field if non-nil, zero value otherwise.

### GetCreatedViaOk

`func (o *MicrosoftPartnerReferral) GetCreatedViaOk() (*string, bool)`

GetCreatedViaOk returns a tuple with the CreatedVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedVia

`func (o *MicrosoftPartnerReferral) SetCreatedVia(v string)`

SetCreatedVia sets CreatedVia field to given value.

### HasCreatedVia

`func (o *MicrosoftPartnerReferral) HasCreatedVia() bool`

HasCreatedVia returns a boolean if a field has been set.

### GetCustomerProfile

`func (o *MicrosoftPartnerReferral) GetCustomerProfile() MicrosoftPartnerReferralProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *MicrosoftPartnerReferral) GetCustomerProfileOk() (*MicrosoftPartnerReferralProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *MicrosoftPartnerReferral) SetCustomerProfile(v MicrosoftPartnerReferralProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *MicrosoftPartnerReferral) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetDealSensitivity

`func (o *MicrosoftPartnerReferral) GetDealSensitivity() string`

GetDealSensitivity returns the DealSensitivity field if non-nil, zero value otherwise.

### GetDealSensitivityOk

`func (o *MicrosoftPartnerReferral) GetDealSensitivityOk() (*string, bool)`

GetDealSensitivityOk returns a tuple with the DealSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealSensitivity

`func (o *MicrosoftPartnerReferral) SetDealSensitivity(v string)`

SetDealSensitivity sets DealSensitivity field to given value.

### HasDealSensitivity

`func (o *MicrosoftPartnerReferral) HasDealSensitivity() bool`

HasDealSensitivity returns a boolean if a field has been set.

### GetDetails

`func (o *MicrosoftPartnerReferral) GetDetails() MicrosoftPartnerReferralDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftPartnerReferral) GetDetailsOk() (*MicrosoftPartnerReferralDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftPartnerReferral) SetDetails(v MicrosoftPartnerReferralDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftPartnerReferral) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDirection

`func (o *MicrosoftPartnerReferral) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MicrosoftPartnerReferral) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MicrosoftPartnerReferral) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MicrosoftPartnerReferral) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetETag

`func (o *MicrosoftPartnerReferral) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MicrosoftPartnerReferral) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MicrosoftPartnerReferral) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MicrosoftPartnerReferral) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetEngagementId

`func (o *MicrosoftPartnerReferral) GetEngagementId() string`

GetEngagementId returns the EngagementId field if non-nil, zero value otherwise.

### GetEngagementIdOk

`func (o *MicrosoftPartnerReferral) GetEngagementIdOk() (*string, bool)`

GetEngagementIdOk returns a tuple with the EngagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementId

`func (o *MicrosoftPartnerReferral) SetEngagementId(v string)`

SetEngagementId sets EngagementId field to given value.

### HasEngagementId

`func (o *MicrosoftPartnerReferral) HasEngagementId() bool`

HasEngagementId returns a boolean if a field has been set.

### GetException

`func (o *MicrosoftPartnerReferral) GetException() map[string]interface{}`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *MicrosoftPartnerReferral) GetExceptionOk() (*map[string]interface{}, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *MicrosoftPartnerReferral) SetException(v map[string]interface{})`

SetException sets Exception field to given value.

### HasException

`func (o *MicrosoftPartnerReferral) HasException() bool`

HasException returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *MicrosoftPartnerReferral) GetExpirationDateTime() string`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftPartnerReferral) GetExpirationDateTimeOk() (*string, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftPartnerReferral) SetExpirationDateTime(v string)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftPartnerReferral) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetFavorite

`func (o *MicrosoftPartnerReferral) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *MicrosoftPartnerReferral) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *MicrosoftPartnerReferral) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *MicrosoftPartnerReferral) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetId

`func (o *MicrosoftPartnerReferral) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftPartnerReferral) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftPartnerReferral) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftPartnerReferral) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInviteContext

`func (o *MicrosoftPartnerReferral) GetInviteContext() MicrosoftPartnerReferralInviteContext`

GetInviteContext returns the InviteContext field if non-nil, zero value otherwise.

### GetInviteContextOk

`func (o *MicrosoftPartnerReferral) GetInviteContextOk() (*MicrosoftPartnerReferralInviteContext, bool)`

GetInviteContextOk returns a tuple with the InviteContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteContext

`func (o *MicrosoftPartnerReferral) SetInviteContext(v MicrosoftPartnerReferralInviteContext)`

SetInviteContext sets InviteContext field to given value.

### HasInviteContext

`func (o *MicrosoftPartnerReferral) HasInviteContext() bool`

HasInviteContext returns a boolean if a field has been set.

### GetIsAutoRegistrationEnabled

`func (o *MicrosoftPartnerReferral) GetIsAutoRegistrationEnabled() bool`

GetIsAutoRegistrationEnabled returns the IsAutoRegistrationEnabled field if non-nil, zero value otherwise.

### GetIsAutoRegistrationEnabledOk

`func (o *MicrosoftPartnerReferral) GetIsAutoRegistrationEnabledOk() (*bool, bool)`

GetIsAutoRegistrationEnabledOk returns a tuple with the IsAutoRegistrationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRegistrationEnabled

`func (o *MicrosoftPartnerReferral) SetIsAutoRegistrationEnabled(v bool)`

SetIsAutoRegistrationEnabled sets IsAutoRegistrationEnabled field to given value.

### HasIsAutoRegistrationEnabled

`func (o *MicrosoftPartnerReferral) HasIsAutoRegistrationEnabled() bool`

HasIsAutoRegistrationEnabled returns a boolean if a field has been set.

### GetIsSpam

`func (o *MicrosoftPartnerReferral) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *MicrosoftPartnerReferral) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *MicrosoftPartnerReferral) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *MicrosoftPartnerReferral) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetLastModifiedVia

`func (o *MicrosoftPartnerReferral) GetLastModifiedVia() string`

GetLastModifiedVia returns the LastModifiedVia field if non-nil, zero value otherwise.

### GetLastModifiedViaOk

`func (o *MicrosoftPartnerReferral) GetLastModifiedViaOk() (*string, bool)`

GetLastModifiedViaOk returns a tuple with the LastModifiedVia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedVia

`func (o *MicrosoftPartnerReferral) SetLastModifiedVia(v string)`

SetLastModifiedVia sets LastModifiedVia field to given value.

### HasLastModifiedVia

`func (o *MicrosoftPartnerReferral) HasLastModifiedVia() bool`

HasLastModifiedVia returns a boolean if a field has been set.

### GetLinks

`func (o *MicrosoftPartnerReferral) GetLinks() MicrosoftPartnerReferralLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftPartnerReferral) GetLinksOk() (*MicrosoftPartnerReferralLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MicrosoftPartnerReferral) SetLinks(v MicrosoftPartnerReferralLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MicrosoftPartnerReferral) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMpnId

`func (o *MicrosoftPartnerReferral) GetMpnId() string`

GetMpnId returns the MpnId field if non-nil, zero value otherwise.

### GetMpnIdOk

`func (o *MicrosoftPartnerReferral) GetMpnIdOk() (*string, bool)`

GetMpnIdOk returns a tuple with the MpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMpnId

`func (o *MicrosoftPartnerReferral) SetMpnId(v string)`

SetMpnId sets MpnId field to given value.

### HasMpnId

`func (o *MicrosoftPartnerReferral) HasMpnId() bool`

HasMpnId returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftPartnerReferral) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftPartnerReferral) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftPartnerReferral) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftPartnerReferral) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *MicrosoftPartnerReferral) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *MicrosoftPartnerReferral) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *MicrosoftPartnerReferral) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *MicrosoftPartnerReferral) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *MicrosoftPartnerReferral) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *MicrosoftPartnerReferral) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *MicrosoftPartnerReferral) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *MicrosoftPartnerReferral) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetQualification

`func (o *MicrosoftPartnerReferral) GetQualification() MicrosoftPartnerReferralQualification`

GetQualification returns the Qualification field if non-nil, zero value otherwise.

### GetQualificationOk

`func (o *MicrosoftPartnerReferral) GetQualificationOk() (*MicrosoftPartnerReferralQualification, bool)`

GetQualificationOk returns a tuple with the Qualification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualification

`func (o *MicrosoftPartnerReferral) SetQualification(v MicrosoftPartnerReferralQualification)`

SetQualification sets Qualification field to given value.

### HasQualification

`func (o *MicrosoftPartnerReferral) HasQualification() bool`

HasQualification returns a boolean if a field has been set.

### GetQuality

`func (o *MicrosoftPartnerReferral) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *MicrosoftPartnerReferral) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *MicrosoftPartnerReferral) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *MicrosoftPartnerReferral) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetReferralProgram

`func (o *MicrosoftPartnerReferral) GetReferralProgram() map[string]interface{}`

GetReferralProgram returns the ReferralProgram field if non-nil, zero value otherwise.

### GetReferralProgramOk

`func (o *MicrosoftPartnerReferral) GetReferralProgramOk() (*map[string]interface{}, bool)`

GetReferralProgramOk returns a tuple with the ReferralProgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralProgram

`func (o *MicrosoftPartnerReferral) SetReferralProgram(v map[string]interface{})`

SetReferralProgram sets ReferralProgram field to given value.

### HasReferralProgram

`func (o *MicrosoftPartnerReferral) HasReferralProgram() bool`

HasReferralProgram returns a boolean if a field has been set.

### GetReferralSource

`func (o *MicrosoftPartnerReferral) GetReferralSource() map[string]interface{}`

GetReferralSource returns the ReferralSource field if non-nil, zero value otherwise.

### GetReferralSourceOk

`func (o *MicrosoftPartnerReferral) GetReferralSourceOk() (*map[string]interface{}, bool)`

GetReferralSourceOk returns a tuple with the ReferralSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralSource

`func (o *MicrosoftPartnerReferral) SetReferralSource(v map[string]interface{})`

SetReferralSource sets ReferralSource field to given value.

### HasReferralSource

`func (o *MicrosoftPartnerReferral) HasReferralSource() bool`

HasReferralSource returns a boolean if a field has been set.

### GetRegistrationStatus

`func (o *MicrosoftPartnerReferral) GetRegistrationStatus() string`

GetRegistrationStatus returns the RegistrationStatus field if non-nil, zero value otherwise.

### GetRegistrationStatusOk

`func (o *MicrosoftPartnerReferral) GetRegistrationStatusOk() (*string, bool)`

GetRegistrationStatusOk returns a tuple with the RegistrationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatus

`func (o *MicrosoftPartnerReferral) SetRegistrationStatus(v string)`

SetRegistrationStatus sets RegistrationStatus field to given value.

### HasRegistrationStatus

`func (o *MicrosoftPartnerReferral) HasRegistrationStatus() bool`

HasRegistrationStatus returns a boolean if a field has been set.

### GetRegistrations

`func (o *MicrosoftPartnerReferral) GetRegistrations() []map[string]interface{}`

GetRegistrations returns the Registrations field if non-nil, zero value otherwise.

### GetRegistrationsOk

`func (o *MicrosoftPartnerReferral) GetRegistrationsOk() (*[]map[string]interface{}, bool)`

GetRegistrationsOk returns a tuple with the Registrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrations

`func (o *MicrosoftPartnerReferral) SetRegistrations(v []map[string]interface{})`

SetRegistrations sets Registrations field to given value.

### HasRegistrations

`func (o *MicrosoftPartnerReferral) HasRegistrations() bool`

HasRegistrations returns a boolean if a field has been set.

### GetSalesStage

`func (o *MicrosoftPartnerReferral) GetSalesStage() map[string]interface{}`

GetSalesStage returns the SalesStage field if non-nil, zero value otherwise.

### GetSalesStageOk

`func (o *MicrosoftPartnerReferral) GetSalesStageOk() (*map[string]interface{}, bool)`

GetSalesStageOk returns a tuple with the SalesStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesStage

`func (o *MicrosoftPartnerReferral) SetSalesStage(v map[string]interface{})`

SetSalesStage sets SalesStage field to given value.

### HasSalesStage

`func (o *MicrosoftPartnerReferral) HasSalesStage() bool`

HasSalesStage returns a boolean if a field has been set.

### GetStatus

`func (o *MicrosoftPartnerReferral) GetStatus() MicrosoftPartnerReferralStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftPartnerReferral) GetStatusOk() (*MicrosoftPartnerReferralStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftPartnerReferral) SetStatus(v MicrosoftPartnerReferralStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftPartnerReferral) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubstatus

`func (o *MicrosoftPartnerReferral) GetSubstatus() MicrosoftPartnerReferralSubStatus`

GetSubstatus returns the Substatus field if non-nil, zero value otherwise.

### GetSubstatusOk

`func (o *MicrosoftPartnerReferral) GetSubstatusOk() (*MicrosoftPartnerReferralSubStatus, bool)`

GetSubstatusOk returns a tuple with the Substatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstatus

`func (o *MicrosoftPartnerReferral) SetSubstatus(v MicrosoftPartnerReferralSubStatus)`

SetSubstatus sets Substatus field to given value.

### HasSubstatus

`func (o *MicrosoftPartnerReferral) HasSubstatus() bool`

HasSubstatus returns a boolean if a field has been set.

### GetTags

`func (o *MicrosoftPartnerReferral) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MicrosoftPartnerReferral) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MicrosoftPartnerReferral) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *MicrosoftPartnerReferral) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTarget

`func (o *MicrosoftPartnerReferral) GetTarget() []MicrosoftPartnerReferralTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftPartnerReferral) GetTargetOk() (*[]MicrosoftPartnerReferralTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftPartnerReferral) SetTarget(v []MicrosoftPartnerReferralTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftPartnerReferral) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTeam

`func (o *MicrosoftPartnerReferral) GetTeam() []MicrosoftPartnerReferralPerson`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MicrosoftPartnerReferral) GetTeamOk() (*[]MicrosoftPartnerReferralPerson, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MicrosoftPartnerReferral) SetTeam(v []MicrosoftPartnerReferralPerson)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MicrosoftPartnerReferral) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetTrackingInfo

`func (o *MicrosoftPartnerReferral) GetTrackingInfo() MicrosoftPartnerReferralTrackingInfo`

GetTrackingInfo returns the TrackingInfo field if non-nil, zero value otherwise.

### GetTrackingInfoOk

`func (o *MicrosoftPartnerReferral) GetTrackingInfoOk() (*MicrosoftPartnerReferralTrackingInfo, bool)`

GetTrackingInfoOk returns a tuple with the TrackingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingInfo

`func (o *MicrosoftPartnerReferral) SetTrackingInfo(v MicrosoftPartnerReferralTrackingInfo)`

SetTrackingInfo sets TrackingInfo field to given value.

### HasTrackingInfo

`func (o *MicrosoftPartnerReferral) HasTrackingInfo() bool`

HasTrackingInfo returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftPartnerReferral) GetType() MicrosoftPartnerReferralType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftPartnerReferral) GetTypeOk() (*MicrosoftPartnerReferralType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftPartnerReferral) SetType(v MicrosoftPartnerReferralType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftPartnerReferral) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *MicrosoftPartnerReferral) GetUpdatedDateTime() string`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *MicrosoftPartnerReferral) GetUpdatedDateTimeOk() (*string, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *MicrosoftPartnerReferral) SetUpdatedDateTime(v string)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *MicrosoftPartnerReferral) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


