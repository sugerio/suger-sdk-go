# WorkloadMetaInfo

## Properties

 Name                                | Type                                                           | Description                                                                                                                                                                                                                                                                                                             | Notes      
-------------------------------------|----------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **AceApnCrmUniqueIdentifier**       | Pointer to **string**                                          | The linked ACE ApnCrmUniqueIdentifier of the private offer if available.                                                                                                                                                                                                                                                | [optional] 
 **AwsSaasProductDimensions**        | Pointer to [**[]AwsProductDimension**](AwsProductDimension.md) | The AWS SaaS product dimensions. Applicable for AWS SaaS products only. This is used to save price info when creating AWS SaaS product.                                                                                                                                                                                 | [optional] 
 **BaseAgreementId**                 | Pointer to **string**                                          | Applicable for AWS Marketplace only, when the IsAgreementBasedOffer is true.                                                                                                                                                                                                                                            | [optional] 
 **BuyerIds**                        | Pointer to **[]string**                                        | The Suger buyer IDs of the private offer if available.                                                                                                                                                                                                                                                                  | [optional] 
 **Contacts**                        | Pointer to [**[]Contact**](Contact.md)                         | The contacts of the offer to notify if any updates.                                                                                                                                                                                                                                                                     | [optional] 
 **CppoInOfferId**                   | Pointer to **string**                                          | The Suger CPPO_IN offer ID.                                                                                                                                                                                                                                                                                             | [optional] 
 **CppoOfferId**                     | Pointer to **string**                                          | The Suger CPPO offer ID. Reseller to end buyer                                                                                                                                                                                                                                                                          | [optional] 
 **CppoOutOfferId**                  | Pointer to **string**                                          | The Suger CPPO_OUT offer ID. ISV to reseller                                                                                                                                                                                                                                                                            | [optional] 
 **CustomMetaInfo**                  | Pointer to **map[string]string**                               | The custom meta info of the offer can be updated by seller via API or console.                                                                                                                                                                                                                                          | [optional] 
 **EnableTestUsageMetering**         | Pointer to **bool**                                            | If enabled, Suger will test metering the usage for this entitlement hourly.                                                                                                                                                                                                                                             | [optional] 
 **EntitlementCancellationSchedule** | Pointer to [**CancellationSchedule**](CancellationSchedule.md) | The cancellation schedule for the entitlement. It is nill if no cancellation schedule.                                                                                                                                                                                                                                  | [optional] 
 **ErrorMessages**                   | Pointer to **[]string**                                        | The error messages when the offer is invalid or offer related tasks failed. Populated by Suger service.                                                                                                                                                                                                                 | [optional] 
 **HubspotDealId**                   | Pointer to **string**                                          | Hubsport deal ID of the private offer if available.                                                                                                                                                                                                                                                                     | [optional] 
 **InternalNote**                    | Pointer to **string**                                          | The Internal note of the private offer. It is only visible to the seller/ISV, not visible to the buyer. Up to 1000 characters.                                                                                                                                                                                          | [optional] 
 **IsAgreementBasedOffer**           | Pointer to **bool**                                            | Applicable for AWS Marketplace only, If this offer is agreement based offer.                                                                                                                                                                                                                                            | [optional] 
 **IsGrossRevenueFullSync**          | Pointer to **bool**                                            | Whether the gross revenue is fully synced for the entitlement.                                                                                                                                                                                                                                                          | [optional] 
 **IsRenewalOffer**                  | Pointer to **bool**                                            | Applicable for AWS Marketplace only. If this offer is renewal offer of existing agreement. The existing agreement can be within or outside AWS Marketplace. AWS may audit and verify your offer is a renewal. If AWS is unable to verify your offer, then AWS may revoke the offer and entitlements from your customer. | [optional] 
 **IsReplacementOffer**              | Pointer to **bool**                                            | If this offer is a GCP replacement offer. Applicable for GCP Marketplace replacement offer only.                                                                                                                                                                                                                        | [optional] 
 **LastModifiedBy**                  | Pointer to [**LastModifiedBy**](LastModifiedBy.md)             | The user who last modified the product/offer/buyer/contact.                                                                                                                                                                                                                                                             | [optional] 
 **Notifications**                   | Pointer to [**[]NotificationEvent**](NotificationEvent.md)     | The notifications of the offer if any updates. In most cases, it is to notify contacts/buyers when the offer is pending acceptance.                                                                                                                                                                                     | [optional] 
 **OfferAcceptDate**                 | Pointer to **time.Time**                                       | The date when the offer is accepted by the buyer. Only available when the private offer has been accepted.                                                                                                                                                                                                              | [optional] 
 **RenewalOfferType**                | Pointer to [**AwsRenewalOfferType**](AwsRenewalOfferType.md)   | Applicable for AWS Marketplace only, required when the IsRenewalOffer is true.                                                                                                                                                                                                                                          | [optional] 
 **ReplacedOfferEndTime**            | Pointer to **time.Time**                                       | The end time of the replaced offer. Applicable for GCP Marketplace replacement offer only.                                                                                                                                                                                                                              | [optional] 
 **ReplacedOfferResourceName**       | Pointer to **string**                                          | The resource name of the GCP Marketplace offer that this offer is replacing. In format of \&quot;projects/{gcpProjectNumber}/services/{productServiceName}/privateOffers/{privateOfferId}\&quot; Applicable for GCP Marketplace replacement offer only.                                                                 | [optional] 
 **SalesforceEntitlementURL**        | Pointer to **string**                                          | The Salesforce entitlement URL                                                                                                                                                                                                                                                                                          | [optional] 
 **SalesforceOpportunityId**         | Pointer to **string**                                          | The Salesforce opportunity ID of the private offer if available.                                                                                                                                                                                                                                                        | [optional] 
 **TestUsageMeteringEndTime**        | Pointer to **time.Time**                                       | The test usage metering end time. It is used for test usage metering only. Required if EnableTestUsageMetering is true.                                                                                                                                                                                                 | [optional] 
 **UpdateMessage**                   | Pointer to **string**                                          | The message to notify when the offer is updated.                                                                                                                                                                                                                                                                        | [optional] 

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

### GetAceApnCrmUniqueIdentifier

`func (o *WorkloadMetaInfo) GetAceApnCrmUniqueIdentifier() string`

GetAceApnCrmUniqueIdentifier returns the AceApnCrmUniqueIdentifier field if non-nil, zero value otherwise.

### GetAceApnCrmUniqueIdentifierOk

`func (o *WorkloadMetaInfo) GetAceApnCrmUniqueIdentifierOk() (*string, bool)`

GetAceApnCrmUniqueIdentifierOk returns a tuple with the AceApnCrmUniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAceApnCrmUniqueIdentifier

`func (o *WorkloadMetaInfo) SetAceApnCrmUniqueIdentifier(v string)`

SetAceApnCrmUniqueIdentifier sets AceApnCrmUniqueIdentifier field to given value.

### HasAceApnCrmUniqueIdentifier

`func (o *WorkloadMetaInfo) HasAceApnCrmUniqueIdentifier() bool`

HasAceApnCrmUniqueIdentifier returns a boolean if a field has been set.

### GetAwsSaasProductDimensions

`func (o *WorkloadMetaInfo) GetAwsSaasProductDimensions() []AwsProductDimension`

GetAwsSaasProductDimensions returns the AwsSaasProductDimensions field if non-nil, zero value otherwise.

### GetAwsSaasProductDimensionsOk

`func (o *WorkloadMetaInfo) GetAwsSaasProductDimensionsOk() (*[]AwsProductDimension, bool)`

GetAwsSaasProductDimensionsOk returns a tuple with the AwsSaasProductDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSaasProductDimensions

`func (o *WorkloadMetaInfo) SetAwsSaasProductDimensions(v []AwsProductDimension)`

SetAwsSaasProductDimensions sets AwsSaasProductDimensions field to given value.

### HasAwsSaasProductDimensions

`func (o *WorkloadMetaInfo) HasAwsSaasProductDimensions() bool`

HasAwsSaasProductDimensions returns a boolean if a field has been set.

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

### GetCppoInOfferId

`func (o *WorkloadMetaInfo) GetCppoInOfferId() string`

GetCppoInOfferId returns the CppoInOfferId field if non-nil, zero value otherwise.

### GetCppoInOfferIdOk

`func (o *WorkloadMetaInfo) GetCppoInOfferIdOk() (*string, bool)`

GetCppoInOfferIdOk returns a tuple with the CppoInOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCppoInOfferId

`func (o *WorkloadMetaInfo) SetCppoInOfferId(v string)`

SetCppoInOfferId sets CppoInOfferId field to given value.

### HasCppoInOfferId

`func (o *WorkloadMetaInfo) HasCppoInOfferId() bool`

HasCppoInOfferId returns a boolean if a field has been set.

### GetCppoOfferId

`func (o *WorkloadMetaInfo) GetCppoOfferId() string`

GetCppoOfferId returns the CppoOfferId field if non-nil, zero value otherwise.

### GetCppoOfferIdOk

`func (o *WorkloadMetaInfo) GetCppoOfferIdOk() (*string, bool)`

GetCppoOfferIdOk returns a tuple with the CppoOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCppoOfferId

`func (o *WorkloadMetaInfo) SetCppoOfferId(v string)`

SetCppoOfferId sets CppoOfferId field to given value.

### HasCppoOfferId

`func (o *WorkloadMetaInfo) HasCppoOfferId() bool`

HasCppoOfferId returns a boolean if a field has been set.

### GetCppoOutOfferId

`func (o *WorkloadMetaInfo) GetCppoOutOfferId() string`

GetCppoOutOfferId returns the CppoOutOfferId field if non-nil, zero value otherwise.

### GetCppoOutOfferIdOk

`func (o *WorkloadMetaInfo) GetCppoOutOfferIdOk() (*string, bool)`

GetCppoOutOfferIdOk returns a tuple with the CppoOutOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCppoOutOfferId

`func (o *WorkloadMetaInfo) SetCppoOutOfferId(v string)`

SetCppoOutOfferId sets CppoOutOfferId field to given value.

### HasCppoOutOfferId

`func (o *WorkloadMetaInfo) HasCppoOutOfferId() bool`

HasCppoOutOfferId returns a boolean if a field has been set.

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

### GetEnableTestUsageMetering

`func (o *WorkloadMetaInfo) GetEnableTestUsageMetering() bool`

GetEnableTestUsageMetering returns the EnableTestUsageMetering field if non-nil, zero value otherwise.

### GetEnableTestUsageMeteringOk

`func (o *WorkloadMetaInfo) GetEnableTestUsageMeteringOk() (*bool, bool)`

GetEnableTestUsageMeteringOk returns a tuple with the EnableTestUsageMetering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTestUsageMetering

`func (o *WorkloadMetaInfo) SetEnableTestUsageMetering(v bool)`

SetEnableTestUsageMetering sets EnableTestUsageMetering field to given value.

### HasEnableTestUsageMetering

`func (o *WorkloadMetaInfo) HasEnableTestUsageMetering() bool`

HasEnableTestUsageMetering returns a boolean if a field has been set.

### GetEntitlementCancellationSchedule

`func (o *WorkloadMetaInfo) GetEntitlementCancellationSchedule() CancellationSchedule`

GetEntitlementCancellationSchedule returns the EntitlementCancellationSchedule field if non-nil, zero value otherwise.

### GetEntitlementCancellationScheduleOk

`func (o *WorkloadMetaInfo) GetEntitlementCancellationScheduleOk() (*CancellationSchedule, bool)`

GetEntitlementCancellationScheduleOk returns a tuple with the EntitlementCancellationSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementCancellationSchedule

`func (o *WorkloadMetaInfo) SetEntitlementCancellationSchedule(v CancellationSchedule)`

SetEntitlementCancellationSchedule sets EntitlementCancellationSchedule field to given value.

### HasEntitlementCancellationSchedule

`func (o *WorkloadMetaInfo) HasEntitlementCancellationSchedule() bool`

HasEntitlementCancellationSchedule returns a boolean if a field has been set.

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

### GetHubspotDealId

`func (o *WorkloadMetaInfo) GetHubspotDealId() string`

GetHubspotDealId returns the HubspotDealId field if non-nil, zero value otherwise.

### GetHubspotDealIdOk

`func (o *WorkloadMetaInfo) GetHubspotDealIdOk() (*string, bool)`

GetHubspotDealIdOk returns a tuple with the HubspotDealId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubspotDealId

`func (o *WorkloadMetaInfo) SetHubspotDealId(v string)`

SetHubspotDealId sets HubspotDealId field to given value.

### HasHubspotDealId

`func (o *WorkloadMetaInfo) HasHubspotDealId() bool`

HasHubspotDealId returns a boolean if a field has been set.

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

### GetIsGrossRevenueFullSync

`func (o *WorkloadMetaInfo) GetIsGrossRevenueFullSync() bool`

GetIsGrossRevenueFullSync returns the IsGrossRevenueFullSync field if non-nil, zero value otherwise.

### GetIsGrossRevenueFullSyncOk

`func (o *WorkloadMetaInfo) GetIsGrossRevenueFullSyncOk() (*bool, bool)`

GetIsGrossRevenueFullSyncOk returns a tuple with the IsGrossRevenueFullSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGrossRevenueFullSync

`func (o *WorkloadMetaInfo) SetIsGrossRevenueFullSync(v bool)`

SetIsGrossRevenueFullSync sets IsGrossRevenueFullSync field to given value.

### HasIsGrossRevenueFullSync

`func (o *WorkloadMetaInfo) HasIsGrossRevenueFullSync() bool`

HasIsGrossRevenueFullSync returns a boolean if a field has been set.

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

### GetIsReplacementOffer

`func (o *WorkloadMetaInfo) GetIsReplacementOffer() bool`

GetIsReplacementOffer returns the IsReplacementOffer field if non-nil, zero value otherwise.

### GetIsReplacementOfferOk

`func (o *WorkloadMetaInfo) GetIsReplacementOfferOk() (*bool, bool)`

GetIsReplacementOfferOk returns a tuple with the IsReplacementOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplacementOffer

`func (o *WorkloadMetaInfo) SetIsReplacementOffer(v bool)`

SetIsReplacementOffer sets IsReplacementOffer field to given value.

### HasIsReplacementOffer

`func (o *WorkloadMetaInfo) HasIsReplacementOffer() bool`

HasIsReplacementOffer returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *WorkloadMetaInfo) GetLastModifiedBy() LastModifiedBy`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *WorkloadMetaInfo) GetLastModifiedByOk() (*LastModifiedBy, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *WorkloadMetaInfo) SetLastModifiedBy(v LastModifiedBy)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *WorkloadMetaInfo) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

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

### GetOfferAcceptDate

`func (o *WorkloadMetaInfo) GetOfferAcceptDate() time.Time`

GetOfferAcceptDate returns the OfferAcceptDate field if non-nil, zero value otherwise.

### GetOfferAcceptDateOk

`func (o *WorkloadMetaInfo) GetOfferAcceptDateOk() (*time.Time, bool)`

GetOfferAcceptDateOk returns a tuple with the OfferAcceptDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferAcceptDate

`func (o *WorkloadMetaInfo) SetOfferAcceptDate(v time.Time)`

SetOfferAcceptDate sets OfferAcceptDate field to given value.

### HasOfferAcceptDate

`func (o *WorkloadMetaInfo) HasOfferAcceptDate() bool`

HasOfferAcceptDate returns a boolean if a field has been set.

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

### GetReplacedOfferEndTime

`func (o *WorkloadMetaInfo) GetReplacedOfferEndTime() time.Time`

GetReplacedOfferEndTime returns the ReplacedOfferEndTime field if non-nil, zero value otherwise.

### GetReplacedOfferEndTimeOk

`func (o *WorkloadMetaInfo) GetReplacedOfferEndTimeOk() (*time.Time, bool)`

GetReplacedOfferEndTimeOk returns a tuple with the ReplacedOfferEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedOfferEndTime

`func (o *WorkloadMetaInfo) SetReplacedOfferEndTime(v time.Time)`

SetReplacedOfferEndTime sets ReplacedOfferEndTime field to given value.

### HasReplacedOfferEndTime

`func (o *WorkloadMetaInfo) HasReplacedOfferEndTime() bool`

HasReplacedOfferEndTime returns a boolean if a field has been set.

### GetReplacedOfferResourceName

`func (o *WorkloadMetaInfo) GetReplacedOfferResourceName() string`

GetReplacedOfferResourceName returns the ReplacedOfferResourceName field if non-nil, zero value otherwise.

### GetReplacedOfferResourceNameOk

`func (o *WorkloadMetaInfo) GetReplacedOfferResourceNameOk() (*string, bool)`

GetReplacedOfferResourceNameOk returns a tuple with the ReplacedOfferResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedOfferResourceName

`func (o *WorkloadMetaInfo) SetReplacedOfferResourceName(v string)`

SetReplacedOfferResourceName sets ReplacedOfferResourceName field to given value.

### HasReplacedOfferResourceName

`func (o *WorkloadMetaInfo) HasReplacedOfferResourceName() bool`

HasReplacedOfferResourceName returns a boolean if a field has been set.

### GetSalesforceEntitlementURL

`func (o *WorkloadMetaInfo) GetSalesforceEntitlementURL() string`

GetSalesforceEntitlementURL returns the SalesforceEntitlementURL field if non-nil, zero value otherwise.

### GetSalesforceEntitlementURLOk

`func (o *WorkloadMetaInfo) GetSalesforceEntitlementURLOk() (*string, bool)`

GetSalesforceEntitlementURLOk returns a tuple with the SalesforceEntitlementURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceEntitlementURL

`func (o *WorkloadMetaInfo) SetSalesforceEntitlementURL(v string)`

SetSalesforceEntitlementURL sets SalesforceEntitlementURL field to given value.

### HasSalesforceEntitlementURL

`func (o *WorkloadMetaInfo) HasSalesforceEntitlementURL() bool`

HasSalesforceEntitlementURL returns a boolean if a field has been set.

### GetSalesforceOpportunityId

`func (o *WorkloadMetaInfo) GetSalesforceOpportunityId() string`

GetSalesforceOpportunityId returns the SalesforceOpportunityId field if non-nil, zero value otherwise.

### GetSalesforceOpportunityIdOk

`func (o *WorkloadMetaInfo) GetSalesforceOpportunityIdOk() (*string, bool)`

GetSalesforceOpportunityIdOk returns a tuple with the SalesforceOpportunityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesforceOpportunityId

`func (o *WorkloadMetaInfo) SetSalesforceOpportunityId(v string)`

SetSalesforceOpportunityId sets SalesforceOpportunityId field to given value.

### HasSalesforceOpportunityId

`func (o *WorkloadMetaInfo) HasSalesforceOpportunityId() bool`

HasSalesforceOpportunityId returns a boolean if a field has been set.

### GetTestUsageMeteringEndTime

`func (o *WorkloadMetaInfo) GetTestUsageMeteringEndTime() time.Time`

GetTestUsageMeteringEndTime returns the TestUsageMeteringEndTime field if non-nil, zero value otherwise.

### GetTestUsageMeteringEndTimeOk

`func (o *WorkloadMetaInfo) GetTestUsageMeteringEndTimeOk() (*time.Time, bool)`

GetTestUsageMeteringEndTimeOk returns a tuple with the TestUsageMeteringEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestUsageMeteringEndTime

`func (o *WorkloadMetaInfo) SetTestUsageMeteringEndTime(v time.Time)`

SetTestUsageMeteringEndTime sets TestUsageMeteringEndTime field to given value.

### HasTestUsageMeteringEndTime

`func (o *WorkloadMetaInfo) HasTestUsageMeteringEndTime() bool`

HasTestUsageMeteringEndTime returns a boolean if a field has been set.

### GetUpdateMessage

`func (o *WorkloadMetaInfo) GetUpdateMessage() string`

GetUpdateMessage returns the UpdateMessage field if non-nil, zero value otherwise.

### GetUpdateMessageOk

`func (o *WorkloadMetaInfo) GetUpdateMessageOk() (*string, bool)`

GetUpdateMessageOk returns a tuple with the UpdateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMessage

`func (o *WorkloadMetaInfo) SetUpdateMessage(v string)`

SetUpdateMessage sets UpdateMessage field to given value.

### HasUpdateMessage

`func (o *WorkloadMetaInfo) HasUpdateMessage() bool`

HasUpdateMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


