# MicrosoftPartnerReferralProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activities** | Pointer to **map[string]interface{}** |  | [optional] 
**Address** | Pointer to [**MicrosoftPartnerReferralAddress**](MicrosoftPartnerReferralAddress.md) |  | [optional] 
**Ids** | Pointer to [**[]MicrosoftPartnerReferralProfileId**](MicrosoftPartnerReferralProfileId.md) |  | [optional] 
**IsMaccEligible** | Pointer to **bool** |  | [optional] 
**IsMatchingComplete** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Team** | Pointer to [**[]MicrosoftPartnerReferralPerson**](MicrosoftPartnerReferralPerson.md) |  | [optional] 

## Methods

### NewMicrosoftPartnerReferralProfile

`func NewMicrosoftPartnerReferralProfile() *MicrosoftPartnerReferralProfile`

NewMicrosoftPartnerReferralProfile instantiates a new MicrosoftPartnerReferralProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftPartnerReferralProfileWithDefaults

`func NewMicrosoftPartnerReferralProfileWithDefaults() *MicrosoftPartnerReferralProfile`

NewMicrosoftPartnerReferralProfileWithDefaults instantiates a new MicrosoftPartnerReferralProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivities

`func (o *MicrosoftPartnerReferralProfile) GetActivities() map[string]interface{}`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *MicrosoftPartnerReferralProfile) GetActivitiesOk() (*map[string]interface{}, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *MicrosoftPartnerReferralProfile) SetActivities(v map[string]interface{})`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *MicrosoftPartnerReferralProfile) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetAddress

`func (o *MicrosoftPartnerReferralProfile) GetAddress() MicrosoftPartnerReferralAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftPartnerReferralProfile) GetAddressOk() (*MicrosoftPartnerReferralAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftPartnerReferralProfile) SetAddress(v MicrosoftPartnerReferralAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftPartnerReferralProfile) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIds

`func (o *MicrosoftPartnerReferralProfile) GetIds() []MicrosoftPartnerReferralProfileId`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MicrosoftPartnerReferralProfile) GetIdsOk() (*[]MicrosoftPartnerReferralProfileId, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MicrosoftPartnerReferralProfile) SetIds(v []MicrosoftPartnerReferralProfileId)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MicrosoftPartnerReferralProfile) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetIsMaccEligible

`func (o *MicrosoftPartnerReferralProfile) GetIsMaccEligible() bool`

GetIsMaccEligible returns the IsMaccEligible field if non-nil, zero value otherwise.

### GetIsMaccEligibleOk

`func (o *MicrosoftPartnerReferralProfile) GetIsMaccEligibleOk() (*bool, bool)`

GetIsMaccEligibleOk returns a tuple with the IsMaccEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaccEligible

`func (o *MicrosoftPartnerReferralProfile) SetIsMaccEligible(v bool)`

SetIsMaccEligible sets IsMaccEligible field to given value.

### HasIsMaccEligible

`func (o *MicrosoftPartnerReferralProfile) HasIsMaccEligible() bool`

HasIsMaccEligible returns a boolean if a field has been set.

### GetIsMatchingComplete

`func (o *MicrosoftPartnerReferralProfile) GetIsMatchingComplete() bool`

GetIsMatchingComplete returns the IsMatchingComplete field if non-nil, zero value otherwise.

### GetIsMatchingCompleteOk

`func (o *MicrosoftPartnerReferralProfile) GetIsMatchingCompleteOk() (*bool, bool)`

GetIsMatchingCompleteOk returns a tuple with the IsMatchingComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMatchingComplete

`func (o *MicrosoftPartnerReferralProfile) SetIsMatchingComplete(v bool)`

SetIsMatchingComplete sets IsMatchingComplete field to given value.

### HasIsMatchingComplete

`func (o *MicrosoftPartnerReferralProfile) HasIsMatchingComplete() bool`

HasIsMatchingComplete returns a boolean if a field has been set.

### GetName

`func (o *MicrosoftPartnerReferralProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftPartnerReferralProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftPartnerReferralProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftPartnerReferralProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *MicrosoftPartnerReferralProfile) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftPartnerReferralProfile) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftPartnerReferralProfile) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftPartnerReferralProfile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTeam

`func (o *MicrosoftPartnerReferralProfile) GetTeam() []MicrosoftPartnerReferralPerson`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *MicrosoftPartnerReferralProfile) GetTeamOk() (*[]MicrosoftPartnerReferralPerson, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *MicrosoftPartnerReferralProfile) SetTeam(v []MicrosoftPartnerReferralPerson)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *MicrosoftPartnerReferralProfile) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


