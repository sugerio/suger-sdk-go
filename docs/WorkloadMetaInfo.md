# WorkloadMetaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseAgreementId** | Pointer to **string** | Applicable for AWS Marketplace only, when the IsAgreementBasedOffer is true. | [optional] 
**BuyerIds** | Pointer to **[]string** | The Suger buyer IDs of the private offer if available. | [optional] 
**Contacts** | Pointer to [**[]Contact**](Contact.md) | The contacts of the offer to notify if any updates. | [optional] 
**CustomMetaInfo** | Pointer to **map[string]string** | The custom meta info of the offer can be updated by seller via API or console. | [optional] 
**ErrorMessages** | Pointer to **[]string** | The error messages when the offer is invalid or offer related tasks failed. Populated by Suger service. | [optional] 
**InternalNote** | Pointer to **string** | The Internal note of the private offer. It is only visible to the seller/ISV, not visible to the buyer. Up to 1000 characters. | [optional] 
**IsAgreementBasedOffer** | Pointer to **bool** | Applicable for AWS Marketplace only, If this offer is agreement based offer. | [optional] 
**IsRenewalOffer** | Pointer to **bool** | Applicable for AWS Marketplace only. If this offer is renewal offer of existing agreement. The existing agreement can be within or outside AWS Marketplace. AWS may audit and verify your offer is a renewal. If AWS is unable to verify your offer, then AWS may revoke the offer and entitlements from your customer. | [optional] 
**Notifications** | Pointer to [**[]NotificationEvent**](NotificationEvent.md) | The notifications of the offer if any updates. In most cases, it is to notify contacts/buyers when the offer is pending acceptance. | [optional] 
**RenewalOfferType** | Pointer to [**AwsRenewalOfferType**](AwsRenewalOfferType.md) |  | [optional] 

## Methods

### NewWorkloadMetaInfo

`func NewWorkloadMetaInfo() *WorkloadMetaInfo`

NewWorkloadMetaInfo instantiates a new WorkloadMetaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadMetaInfoWithDefaults

`func NewWorkloadMetaInfoWithDefaults() *WorkloadMetaInfo`

NewWorkloadMetaInfoWithDefaults instantiates a new WorkloadMetaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseAgreementId

`func (o *WorkloadMetaInfo) GetBaseAgreementId() string`

GetBaseAgreementId returns the BaseAgreementId field if non-nil, zero value otherwise.

### GetBaseAgreementIdOk

`func (o *WorkloadMetaInfo) GetBaseAgreementIdOk() (*string, bool)`

GetBaseAgreementIdOk returns a tuple with the BaseAgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAgreementId

`func (o *WorkloadMetaInfo) SetBaseAgreementId(v string)`

SetBaseAgreementId sets BaseAgreementId field to given value.

### HasBaseAgreementId

`func (o *WorkloadMetaInfo) HasBaseAgreementId() bool`

HasBaseAgreementId returns a boolean if a field has been set.

### GetBuyerIds

`func (o *WorkloadMetaInfo) GetBuyerIds() []string`

GetBuyerIds returns the BuyerIds field if non-nil, zero value otherwise.

### GetBuyerIdsOk

`func (o *WorkloadMetaInfo) GetBuyerIdsOk() (*[]string, bool)`

GetBuyerIdsOk returns a tuple with the BuyerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerIds

`func (o *WorkloadMetaInfo) SetBuyerIds(v []string)`

SetBuyerIds sets BuyerIds field to given value.

### HasBuyerIds

`func (o *WorkloadMetaInfo) HasBuyerIds() bool`

HasBuyerIds returns a boolean if a field has been set.

### GetContacts

`func (o *WorkloadMetaInfo) GetContacts() []Contact`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *WorkloadMetaInfo) GetContactsOk() (*[]Contact, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *WorkloadMetaInfo) SetContacts(v []Contact)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *WorkloadMetaInfo) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetCustomMetaInfo

`func (o *WorkloadMetaInfo) GetCustomMetaInfo() map[string]string`

GetCustomMetaInfo returns the CustomMetaInfo field if non-nil, zero value otherwise.

### GetCustomMetaInfoOk

`func (o *WorkloadMetaInfo) GetCustomMetaInfoOk() (*map[string]string, bool)`

GetCustomMetaInfoOk returns a tuple with the CustomMetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMetaInfo

`func (o *WorkloadMetaInfo) SetCustomMetaInfo(v map[string]string)`

SetCustomMetaInfo sets CustomMetaInfo field to given value.

### HasCustomMetaInfo

`func (o *WorkloadMetaInfo) HasCustomMetaInfo() bool`

HasCustomMetaInfo returns a boolean if a field has been set.

### GetErrorMessages

`func (o *WorkloadMetaInfo) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *WorkloadMetaInfo) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *WorkloadMetaInfo) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *WorkloadMetaInfo) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetInternalNote

`func (o *WorkloadMetaInfo) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *WorkloadMetaInfo) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *WorkloadMetaInfo) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.

### HasInternalNote

`func (o *WorkloadMetaInfo) HasInternalNote() bool`

HasInternalNote returns a boolean if a field has been set.

### GetIsAgreementBasedOffer

`func (o *WorkloadMetaInfo) GetIsAgreementBasedOffer() bool`

GetIsAgreementBasedOffer returns the IsAgreementBasedOffer field if non-nil, zero value otherwise.

### GetIsAgreementBasedOfferOk

`func (o *WorkloadMetaInfo) GetIsAgreementBasedOfferOk() (*bool, bool)`

GetIsAgreementBasedOfferOk returns a tuple with the IsAgreementBasedOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAgreementBasedOffer

`func (o *WorkloadMetaInfo) SetIsAgreementBasedOffer(v bool)`

SetIsAgreementBasedOffer sets IsAgreementBasedOffer field to given value.

### HasIsAgreementBasedOffer

`func (o *WorkloadMetaInfo) HasIsAgreementBasedOffer() bool`

HasIsAgreementBasedOffer returns a boolean if a field has been set.

### GetIsRenewalOffer

`func (o *WorkloadMetaInfo) GetIsRenewalOffer() bool`

GetIsRenewalOffer returns the IsRenewalOffer field if non-nil, zero value otherwise.

### GetIsRenewalOfferOk

`func (o *WorkloadMetaInfo) GetIsRenewalOfferOk() (*bool, bool)`

GetIsRenewalOfferOk returns a tuple with the IsRenewalOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRenewalOffer

`func (o *WorkloadMetaInfo) SetIsRenewalOffer(v bool)`

SetIsRenewalOffer sets IsRenewalOffer field to given value.

### HasIsRenewalOffer

`func (o *WorkloadMetaInfo) HasIsRenewalOffer() bool`

HasIsRenewalOffer returns a boolean if a field has been set.

### GetNotifications

`func (o *WorkloadMetaInfo) GetNotifications() []NotificationEvent`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *WorkloadMetaInfo) GetNotificationsOk() (*[]NotificationEvent, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *WorkloadMetaInfo) SetNotifications(v []NotificationEvent)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *WorkloadMetaInfo) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetRenewalOfferType

`func (o *WorkloadMetaInfo) GetRenewalOfferType() AwsRenewalOfferType`

GetRenewalOfferType returns the RenewalOfferType field if non-nil, zero value otherwise.

### GetRenewalOfferTypeOk

`func (o *WorkloadMetaInfo) GetRenewalOfferTypeOk() (*AwsRenewalOfferType, bool)`

GetRenewalOfferTypeOk returns a tuple with the RenewalOfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalOfferType

`func (o *WorkloadMetaInfo) SetRenewalOfferType(v AwsRenewalOfferType)`

SetRenewalOfferType sets RenewalOfferType field to given value.

### HasRenewalOfferType

`func (o *WorkloadMetaInfo) HasRenewalOfferType() bool`

HasRenewalOfferType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


