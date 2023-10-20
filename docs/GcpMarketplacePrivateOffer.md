# GcpMarketplacePrivateOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveEntitlement** | Pointer to [**GcpMarketplaceEntitlement**](GcpMarketplaceEntitlement.md) |  | [optional] 
**AgencyEnabled** | Pointer to **bool** |  | [optional] 
**Agreement** | Pointer to **string** | The resource name of agreement(entitlement) In format of \&quot;projects/{projectNumber}/agreements/{agreementId}\&quot; | [optional] 
**CancelTime** | Pointer to **time.Time** |  | [optional] 
**CustomEula** | Pointer to [**GcpMarketplaceDocument**](GcpMarketplaceDocument.md) |  | [optional] 
**CustomerInfo** | Pointer to [**GcpMarketplacePrivateOfferCustomerInfo**](GcpMarketplacePrivateOfferCustomerInfo.md) |  | [optional] 
**EulaAgreementDocument** | Pointer to [**GcpMarketplaceDocument**](GcpMarketplaceDocument.md) |  | [optional] 
**ExistingOfferData** | Pointer to [**GcpMarketplaceExistingOfferData**](GcpMarketplaceExistingOfferData.md) |  | [optional] 
**ExpireTime** | Pointer to **time.Time** |  | [optional] 
**Features** | Pointer to [**[]GcpMarketplaceProductFeatureValue**](GcpMarketplaceProductFeatureValue.md) |  | [optional] 
**InstallmentTimeline** | Pointer to [**GcpMarketplacePrivateOfferInstallmentTimeline**](GcpMarketplacePrivateOfferInstallmentTimeline.md) |  | [optional] 
**LifecycleState** | Pointer to **string** | such as \&quot;PUBLISHED\&quot; | [optional] 
**MetricInformation** | Pointer to [**GcpMarketplacePrivateOfferMetricInformation**](GcpMarketplacePrivateOfferMetricInformation.md) |  | [optional] 
**MigrationMetadata** | Pointer to [**GcpMarketplacePrivateOfferMigrationMetadata**](GcpMarketplacePrivateOfferMigrationMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | In format of \&quot;projects/{projectNumber}/services/{serviceName, such as service-name.endpoints.gcp-project-id.cloud.goog}/privateOffers/{privateOfferId}\&quot; | [optional] 
**OfferId** | Pointer to **string** | GCP private offer ID | [optional] 
**OfferSource** | Pointer to **string** | such as \&quot;OFFER\&quot; | [optional] 
**OfferState** | Pointer to [**GcpMarketplacePrivateOfferState**](GcpMarketplacePrivateOfferState.md) |  | [optional] 
**OfferTerm** | Pointer to [**GcpMarketplacePrivateOfferTerm**](GcpMarketplacePrivateOfferTerm.md) |  | [optional] 
**PaymentSchedule** | Pointer to [**GcpMarketplacePaymentScheduleType**](GcpMarketplacePaymentScheduleType.md) |  | [optional] 
**Policies** | Pointer to **map[string]string** |  | [optional] 
**PriceModel** | Pointer to [**GcpMarketplacePrivateOfferPriceModel**](GcpMarketplacePrivateOfferPriceModel.md) |  | [optional] 
**PriceModelType** | Pointer to [**GcpMarketplacePrivateOfferPriceModelType**](GcpMarketplacePrivateOfferPriceModelType.md) |  | [optional] 
**ProviderCancellationInternalNote** | Pointer to **string** |  | [optional] 
**ProviderInfo** | Pointer to [**GcpMarketplacePrivateOfferProviderInfo**](GcpMarketplacePrivateOfferProviderInfo.md) |  | [optional] 
**ProviderInternalNote** | Pointer to **string** |  | [optional] 
**ProviderPublicNote** | Pointer to **string** |  | [optional] 
**PurchaseChannel** | Pointer to [**GcpMarketplacePurchaseChannel**](GcpMarketplacePurchaseChannel.md) |  | [optional] 
**PurchaseTime** | Pointer to **time.Time** |  | [optional] 
**ReplacementMetadata** | Pointer to [**GcpMarketplacePrivateOfferReplacementMetadata**](GcpMarketplacePrivateOfferReplacementMetadata.md) |  | [optional] 
**ServiceLevel** | Pointer to **string** | The Plan of the offer. | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**UseLegacyPartnerEula** | Pointer to **bool** |  | [optional] 
**UserLabels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGcpMarketplacePrivateOffer

`func NewGcpMarketplacePrivateOffer() *GcpMarketplacePrivateOffer`

NewGcpMarketplacePrivateOffer instantiates a new GcpMarketplacePrivateOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferWithDefaults

`func NewGcpMarketplacePrivateOfferWithDefaults() *GcpMarketplacePrivateOffer`

NewGcpMarketplacePrivateOfferWithDefaults instantiates a new GcpMarketplacePrivateOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveEntitlement

`func (o *GcpMarketplacePrivateOffer) GetActiveEntitlement() GcpMarketplaceEntitlement`

GetActiveEntitlement returns the ActiveEntitlement field if non-nil, zero value otherwise.

### GetActiveEntitlementOk

`func (o *GcpMarketplacePrivateOffer) GetActiveEntitlementOk() (*GcpMarketplaceEntitlement, bool)`

GetActiveEntitlementOk returns a tuple with the ActiveEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveEntitlement

`func (o *GcpMarketplacePrivateOffer) SetActiveEntitlement(v GcpMarketplaceEntitlement)`

SetActiveEntitlement sets ActiveEntitlement field to given value.

### HasActiveEntitlement

`func (o *GcpMarketplacePrivateOffer) HasActiveEntitlement() bool`

HasActiveEntitlement returns a boolean if a field has been set.

### GetAgencyEnabled

`func (o *GcpMarketplacePrivateOffer) GetAgencyEnabled() bool`

GetAgencyEnabled returns the AgencyEnabled field if non-nil, zero value otherwise.

### GetAgencyEnabledOk

`func (o *GcpMarketplacePrivateOffer) GetAgencyEnabledOk() (*bool, bool)`

GetAgencyEnabledOk returns a tuple with the AgencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyEnabled

`func (o *GcpMarketplacePrivateOffer) SetAgencyEnabled(v bool)`

SetAgencyEnabled sets AgencyEnabled field to given value.

### HasAgencyEnabled

`func (o *GcpMarketplacePrivateOffer) HasAgencyEnabled() bool`

HasAgencyEnabled returns a boolean if a field has been set.

### GetAgreement

`func (o *GcpMarketplacePrivateOffer) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *GcpMarketplacePrivateOffer) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *GcpMarketplacePrivateOffer) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *GcpMarketplacePrivateOffer) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCancelTime

`func (o *GcpMarketplacePrivateOffer) GetCancelTime() time.Time`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *GcpMarketplacePrivateOffer) GetCancelTimeOk() (*time.Time, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *GcpMarketplacePrivateOffer) SetCancelTime(v time.Time)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *GcpMarketplacePrivateOffer) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetCustomEula

`func (o *GcpMarketplacePrivateOffer) GetCustomEula() GcpMarketplaceDocument`

GetCustomEula returns the CustomEula field if non-nil, zero value otherwise.

### GetCustomEulaOk

`func (o *GcpMarketplacePrivateOffer) GetCustomEulaOk() (*GcpMarketplaceDocument, bool)`

GetCustomEulaOk returns a tuple with the CustomEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEula

`func (o *GcpMarketplacePrivateOffer) SetCustomEula(v GcpMarketplaceDocument)`

SetCustomEula sets CustomEula field to given value.

### HasCustomEula

`func (o *GcpMarketplacePrivateOffer) HasCustomEula() bool`

HasCustomEula returns a boolean if a field has been set.

### GetCustomerInfo

`func (o *GcpMarketplacePrivateOffer) GetCustomerInfo() GcpMarketplacePrivateOfferCustomerInfo`

GetCustomerInfo returns the CustomerInfo field if non-nil, zero value otherwise.

### GetCustomerInfoOk

`func (o *GcpMarketplacePrivateOffer) GetCustomerInfoOk() (*GcpMarketplacePrivateOfferCustomerInfo, bool)`

GetCustomerInfoOk returns a tuple with the CustomerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerInfo

`func (o *GcpMarketplacePrivateOffer) SetCustomerInfo(v GcpMarketplacePrivateOfferCustomerInfo)`

SetCustomerInfo sets CustomerInfo field to given value.

### HasCustomerInfo

`func (o *GcpMarketplacePrivateOffer) HasCustomerInfo() bool`

HasCustomerInfo returns a boolean if a field has been set.

### GetEulaAgreementDocument

`func (o *GcpMarketplacePrivateOffer) GetEulaAgreementDocument() GcpMarketplaceDocument`

GetEulaAgreementDocument returns the EulaAgreementDocument field if non-nil, zero value otherwise.

### GetEulaAgreementDocumentOk

`func (o *GcpMarketplacePrivateOffer) GetEulaAgreementDocumentOk() (*GcpMarketplaceDocument, bool)`

GetEulaAgreementDocumentOk returns a tuple with the EulaAgreementDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaAgreementDocument

`func (o *GcpMarketplacePrivateOffer) SetEulaAgreementDocument(v GcpMarketplaceDocument)`

SetEulaAgreementDocument sets EulaAgreementDocument field to given value.

### HasEulaAgreementDocument

`func (o *GcpMarketplacePrivateOffer) HasEulaAgreementDocument() bool`

HasEulaAgreementDocument returns a boolean if a field has been set.

### GetExistingOfferData

`func (o *GcpMarketplacePrivateOffer) GetExistingOfferData() GcpMarketplaceExistingOfferData`

GetExistingOfferData returns the ExistingOfferData field if non-nil, zero value otherwise.

### GetExistingOfferDataOk

`func (o *GcpMarketplacePrivateOffer) GetExistingOfferDataOk() (*GcpMarketplaceExistingOfferData, bool)`

GetExistingOfferDataOk returns a tuple with the ExistingOfferData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingOfferData

`func (o *GcpMarketplacePrivateOffer) SetExistingOfferData(v GcpMarketplaceExistingOfferData)`

SetExistingOfferData sets ExistingOfferData field to given value.

### HasExistingOfferData

`func (o *GcpMarketplacePrivateOffer) HasExistingOfferData() bool`

HasExistingOfferData returns a boolean if a field has been set.

### GetExpireTime

`func (o *GcpMarketplacePrivateOffer) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *GcpMarketplacePrivateOffer) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *GcpMarketplacePrivateOffer) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *GcpMarketplacePrivateOffer) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetFeatures

`func (o *GcpMarketplacePrivateOffer) GetFeatures() []GcpMarketplaceProductFeatureValue`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GcpMarketplacePrivateOffer) GetFeaturesOk() (*[]GcpMarketplaceProductFeatureValue, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GcpMarketplacePrivateOffer) SetFeatures(v []GcpMarketplaceProductFeatureValue)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GcpMarketplacePrivateOffer) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetInstallmentTimeline

`func (o *GcpMarketplacePrivateOffer) GetInstallmentTimeline() GcpMarketplacePrivateOfferInstallmentTimeline`

GetInstallmentTimeline returns the InstallmentTimeline field if non-nil, zero value otherwise.

### GetInstallmentTimelineOk

`func (o *GcpMarketplacePrivateOffer) GetInstallmentTimelineOk() (*GcpMarketplacePrivateOfferInstallmentTimeline, bool)`

GetInstallmentTimelineOk returns a tuple with the InstallmentTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentTimeline

`func (o *GcpMarketplacePrivateOffer) SetInstallmentTimeline(v GcpMarketplacePrivateOfferInstallmentTimeline)`

SetInstallmentTimeline sets InstallmentTimeline field to given value.

### HasInstallmentTimeline

`func (o *GcpMarketplacePrivateOffer) HasInstallmentTimeline() bool`

HasInstallmentTimeline returns a boolean if a field has been set.

### GetLifecycleState

`func (o *GcpMarketplacePrivateOffer) GetLifecycleState() string`

GetLifecycleState returns the LifecycleState field if non-nil, zero value otherwise.

### GetLifecycleStateOk

`func (o *GcpMarketplacePrivateOffer) GetLifecycleStateOk() (*string, bool)`

GetLifecycleStateOk returns a tuple with the LifecycleState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleState

`func (o *GcpMarketplacePrivateOffer) SetLifecycleState(v string)`

SetLifecycleState sets LifecycleState field to given value.

### HasLifecycleState

`func (o *GcpMarketplacePrivateOffer) HasLifecycleState() bool`

HasLifecycleState returns a boolean if a field has been set.

### GetMetricInformation

`func (o *GcpMarketplacePrivateOffer) GetMetricInformation() GcpMarketplacePrivateOfferMetricInformation`

GetMetricInformation returns the MetricInformation field if non-nil, zero value otherwise.

### GetMetricInformationOk

`func (o *GcpMarketplacePrivateOffer) GetMetricInformationOk() (*GcpMarketplacePrivateOfferMetricInformation, bool)`

GetMetricInformationOk returns a tuple with the MetricInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricInformation

`func (o *GcpMarketplacePrivateOffer) SetMetricInformation(v GcpMarketplacePrivateOfferMetricInformation)`

SetMetricInformation sets MetricInformation field to given value.

### HasMetricInformation

`func (o *GcpMarketplacePrivateOffer) HasMetricInformation() bool`

HasMetricInformation returns a boolean if a field has been set.

### GetMigrationMetadata

`func (o *GcpMarketplacePrivateOffer) GetMigrationMetadata() GcpMarketplacePrivateOfferMigrationMetadata`

GetMigrationMetadata returns the MigrationMetadata field if non-nil, zero value otherwise.

### GetMigrationMetadataOk

`func (o *GcpMarketplacePrivateOffer) GetMigrationMetadataOk() (*GcpMarketplacePrivateOfferMigrationMetadata, bool)`

GetMigrationMetadataOk returns a tuple with the MigrationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationMetadata

`func (o *GcpMarketplacePrivateOffer) SetMigrationMetadata(v GcpMarketplacePrivateOfferMigrationMetadata)`

SetMigrationMetadata sets MigrationMetadata field to given value.

### HasMigrationMetadata

`func (o *GcpMarketplacePrivateOffer) HasMigrationMetadata() bool`

HasMigrationMetadata returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplacePrivateOffer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplacePrivateOffer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplacePrivateOffer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplacePrivateOffer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferId

`func (o *GcpMarketplacePrivateOffer) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *GcpMarketplacePrivateOffer) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *GcpMarketplacePrivateOffer) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *GcpMarketplacePrivateOffer) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetOfferSource

`func (o *GcpMarketplacePrivateOffer) GetOfferSource() string`

GetOfferSource returns the OfferSource field if non-nil, zero value otherwise.

### GetOfferSourceOk

`func (o *GcpMarketplacePrivateOffer) GetOfferSourceOk() (*string, bool)`

GetOfferSourceOk returns a tuple with the OfferSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferSource

`func (o *GcpMarketplacePrivateOffer) SetOfferSource(v string)`

SetOfferSource sets OfferSource field to given value.

### HasOfferSource

`func (o *GcpMarketplacePrivateOffer) HasOfferSource() bool`

HasOfferSource returns a boolean if a field has been set.

### GetOfferState

`func (o *GcpMarketplacePrivateOffer) GetOfferState() GcpMarketplacePrivateOfferState`

GetOfferState returns the OfferState field if non-nil, zero value otherwise.

### GetOfferStateOk

`func (o *GcpMarketplacePrivateOffer) GetOfferStateOk() (*GcpMarketplacePrivateOfferState, bool)`

GetOfferStateOk returns a tuple with the OfferState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferState

`func (o *GcpMarketplacePrivateOffer) SetOfferState(v GcpMarketplacePrivateOfferState)`

SetOfferState sets OfferState field to given value.

### HasOfferState

`func (o *GcpMarketplacePrivateOffer) HasOfferState() bool`

HasOfferState returns a boolean if a field has been set.

### GetOfferTerm

`func (o *GcpMarketplacePrivateOffer) GetOfferTerm() GcpMarketplacePrivateOfferTerm`

GetOfferTerm returns the OfferTerm field if non-nil, zero value otherwise.

### GetOfferTermOk

`func (o *GcpMarketplacePrivateOffer) GetOfferTermOk() (*GcpMarketplacePrivateOfferTerm, bool)`

GetOfferTermOk returns a tuple with the OfferTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferTerm

`func (o *GcpMarketplacePrivateOffer) SetOfferTerm(v GcpMarketplacePrivateOfferTerm)`

SetOfferTerm sets OfferTerm field to given value.

### HasOfferTerm

`func (o *GcpMarketplacePrivateOffer) HasOfferTerm() bool`

HasOfferTerm returns a boolean if a field has been set.

### GetPaymentSchedule

`func (o *GcpMarketplacePrivateOffer) GetPaymentSchedule() GcpMarketplacePaymentScheduleType`

GetPaymentSchedule returns the PaymentSchedule field if non-nil, zero value otherwise.

### GetPaymentScheduleOk

`func (o *GcpMarketplacePrivateOffer) GetPaymentScheduleOk() (*GcpMarketplacePaymentScheduleType, bool)`

GetPaymentScheduleOk returns a tuple with the PaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedule

`func (o *GcpMarketplacePrivateOffer) SetPaymentSchedule(v GcpMarketplacePaymentScheduleType)`

SetPaymentSchedule sets PaymentSchedule field to given value.

### HasPaymentSchedule

`func (o *GcpMarketplacePrivateOffer) HasPaymentSchedule() bool`

HasPaymentSchedule returns a boolean if a field has been set.

### GetPolicies

`func (o *GcpMarketplacePrivateOffer) GetPolicies() map[string]string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GcpMarketplacePrivateOffer) GetPoliciesOk() (*map[string]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GcpMarketplacePrivateOffer) SetPolicies(v map[string]string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *GcpMarketplacePrivateOffer) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetPriceModel

`func (o *GcpMarketplacePrivateOffer) GetPriceModel() GcpMarketplacePrivateOfferPriceModel`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *GcpMarketplacePrivateOffer) GetPriceModelOk() (*GcpMarketplacePrivateOfferPriceModel, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *GcpMarketplacePrivateOffer) SetPriceModel(v GcpMarketplacePrivateOfferPriceModel)`

SetPriceModel sets PriceModel field to given value.

### HasPriceModel

`func (o *GcpMarketplacePrivateOffer) HasPriceModel() bool`

HasPriceModel returns a boolean if a field has been set.

### GetPriceModelType

`func (o *GcpMarketplacePrivateOffer) GetPriceModelType() GcpMarketplacePrivateOfferPriceModelType`

GetPriceModelType returns the PriceModelType field if non-nil, zero value otherwise.

### GetPriceModelTypeOk

`func (o *GcpMarketplacePrivateOffer) GetPriceModelTypeOk() (*GcpMarketplacePrivateOfferPriceModelType, bool)`

GetPriceModelTypeOk returns a tuple with the PriceModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelType

`func (o *GcpMarketplacePrivateOffer) SetPriceModelType(v GcpMarketplacePrivateOfferPriceModelType)`

SetPriceModelType sets PriceModelType field to given value.

### HasPriceModelType

`func (o *GcpMarketplacePrivateOffer) HasPriceModelType() bool`

HasPriceModelType returns a boolean if a field has been set.

### GetProviderCancellationInternalNote

`func (o *GcpMarketplacePrivateOffer) GetProviderCancellationInternalNote() string`

GetProviderCancellationInternalNote returns the ProviderCancellationInternalNote field if non-nil, zero value otherwise.

### GetProviderCancellationInternalNoteOk

`func (o *GcpMarketplacePrivateOffer) GetProviderCancellationInternalNoteOk() (*string, bool)`

GetProviderCancellationInternalNoteOk returns a tuple with the ProviderCancellationInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCancellationInternalNote

`func (o *GcpMarketplacePrivateOffer) SetProviderCancellationInternalNote(v string)`

SetProviderCancellationInternalNote sets ProviderCancellationInternalNote field to given value.

### HasProviderCancellationInternalNote

`func (o *GcpMarketplacePrivateOffer) HasProviderCancellationInternalNote() bool`

HasProviderCancellationInternalNote returns a boolean if a field has been set.

### GetProviderInfo

`func (o *GcpMarketplacePrivateOffer) GetProviderInfo() GcpMarketplacePrivateOfferProviderInfo`

GetProviderInfo returns the ProviderInfo field if non-nil, zero value otherwise.

### GetProviderInfoOk

`func (o *GcpMarketplacePrivateOffer) GetProviderInfoOk() (*GcpMarketplacePrivateOfferProviderInfo, bool)`

GetProviderInfoOk returns a tuple with the ProviderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInfo

`func (o *GcpMarketplacePrivateOffer) SetProviderInfo(v GcpMarketplacePrivateOfferProviderInfo)`

SetProviderInfo sets ProviderInfo field to given value.

### HasProviderInfo

`func (o *GcpMarketplacePrivateOffer) HasProviderInfo() bool`

HasProviderInfo returns a boolean if a field has been set.

### GetProviderInternalNote

`func (o *GcpMarketplacePrivateOffer) GetProviderInternalNote() string`

GetProviderInternalNote returns the ProviderInternalNote field if non-nil, zero value otherwise.

### GetProviderInternalNoteOk

`func (o *GcpMarketplacePrivateOffer) GetProviderInternalNoteOk() (*string, bool)`

GetProviderInternalNoteOk returns a tuple with the ProviderInternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderInternalNote

`func (o *GcpMarketplacePrivateOffer) SetProviderInternalNote(v string)`

SetProviderInternalNote sets ProviderInternalNote field to given value.

### HasProviderInternalNote

`func (o *GcpMarketplacePrivateOffer) HasProviderInternalNote() bool`

HasProviderInternalNote returns a boolean if a field has been set.

### GetProviderPublicNote

`func (o *GcpMarketplacePrivateOffer) GetProviderPublicNote() string`

GetProviderPublicNote returns the ProviderPublicNote field if non-nil, zero value otherwise.

### GetProviderPublicNoteOk

`func (o *GcpMarketplacePrivateOffer) GetProviderPublicNoteOk() (*string, bool)`

GetProviderPublicNoteOk returns a tuple with the ProviderPublicNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderPublicNote

`func (o *GcpMarketplacePrivateOffer) SetProviderPublicNote(v string)`

SetProviderPublicNote sets ProviderPublicNote field to given value.

### HasProviderPublicNote

`func (o *GcpMarketplacePrivateOffer) HasProviderPublicNote() bool`

HasProviderPublicNote returns a boolean if a field has been set.

### GetPurchaseChannel

`func (o *GcpMarketplacePrivateOffer) GetPurchaseChannel() GcpMarketplacePurchaseChannel`

GetPurchaseChannel returns the PurchaseChannel field if non-nil, zero value otherwise.

### GetPurchaseChannelOk

`func (o *GcpMarketplacePrivateOffer) GetPurchaseChannelOk() (*GcpMarketplacePurchaseChannel, bool)`

GetPurchaseChannelOk returns a tuple with the PurchaseChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseChannel

`func (o *GcpMarketplacePrivateOffer) SetPurchaseChannel(v GcpMarketplacePurchaseChannel)`

SetPurchaseChannel sets PurchaseChannel field to given value.

### HasPurchaseChannel

`func (o *GcpMarketplacePrivateOffer) HasPurchaseChannel() bool`

HasPurchaseChannel returns a boolean if a field has been set.

### GetPurchaseTime

`func (o *GcpMarketplacePrivateOffer) GetPurchaseTime() time.Time`

GetPurchaseTime returns the PurchaseTime field if non-nil, zero value otherwise.

### GetPurchaseTimeOk

`func (o *GcpMarketplacePrivateOffer) GetPurchaseTimeOk() (*time.Time, bool)`

GetPurchaseTimeOk returns a tuple with the PurchaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTime

`func (o *GcpMarketplacePrivateOffer) SetPurchaseTime(v time.Time)`

SetPurchaseTime sets PurchaseTime field to given value.

### HasPurchaseTime

`func (o *GcpMarketplacePrivateOffer) HasPurchaseTime() bool`

HasPurchaseTime returns a boolean if a field has been set.

### GetReplacementMetadata

`func (o *GcpMarketplacePrivateOffer) GetReplacementMetadata() GcpMarketplacePrivateOfferReplacementMetadata`

GetReplacementMetadata returns the ReplacementMetadata field if non-nil, zero value otherwise.

### GetReplacementMetadataOk

`func (o *GcpMarketplacePrivateOffer) GetReplacementMetadataOk() (*GcpMarketplacePrivateOfferReplacementMetadata, bool)`

GetReplacementMetadataOk returns a tuple with the ReplacementMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementMetadata

`func (o *GcpMarketplacePrivateOffer) SetReplacementMetadata(v GcpMarketplacePrivateOfferReplacementMetadata)`

SetReplacementMetadata sets ReplacementMetadata field to given value.

### HasReplacementMetadata

`func (o *GcpMarketplacePrivateOffer) HasReplacementMetadata() bool`

HasReplacementMetadata returns a boolean if a field has been set.

### GetServiceLevel

`func (o *GcpMarketplacePrivateOffer) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *GcpMarketplacePrivateOffer) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *GcpMarketplacePrivateOffer) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *GcpMarketplacePrivateOffer) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GcpMarketplacePrivateOffer) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GcpMarketplacePrivateOffer) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GcpMarketplacePrivateOffer) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GcpMarketplacePrivateOffer) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUseLegacyPartnerEula

`func (o *GcpMarketplacePrivateOffer) GetUseLegacyPartnerEula() bool`

GetUseLegacyPartnerEula returns the UseLegacyPartnerEula field if non-nil, zero value otherwise.

### GetUseLegacyPartnerEulaOk

`func (o *GcpMarketplacePrivateOffer) GetUseLegacyPartnerEulaOk() (*bool, bool)`

GetUseLegacyPartnerEulaOk returns a tuple with the UseLegacyPartnerEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyPartnerEula

`func (o *GcpMarketplacePrivateOffer) SetUseLegacyPartnerEula(v bool)`

SetUseLegacyPartnerEula sets UseLegacyPartnerEula field to given value.

### HasUseLegacyPartnerEula

`func (o *GcpMarketplacePrivateOffer) HasUseLegacyPartnerEula() bool`

HasUseLegacyPartnerEula returns a boolean if a field has been set.

### GetUserLabels

`func (o *GcpMarketplacePrivateOffer) GetUserLabels() []string`

GetUserLabels returns the UserLabels field if non-nil, zero value otherwise.

### GetUserLabelsOk

`func (o *GcpMarketplacePrivateOffer) GetUserLabelsOk() (*[]string, bool)`

GetUserLabelsOk returns a tuple with the UserLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabels

`func (o *GcpMarketplacePrivateOffer) SetUserLabels(v []string)`

SetUserLabels sets UserLabels field to given value.

### HasUserLabels

`func (o *GcpMarketplacePrivateOffer) HasUserLabels() bool`

HasUserLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


