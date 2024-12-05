# GcpMarketplaceResellerPrivateOfferPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptanceDeadlineTime** | Pointer to **time.Time** |  | [optional] 
**Agreement** | Pointer to **string** | The resource name of agreement(entitlement) In format of \&quot;projects/{projectNumber}/agreements/{agreementId}\&quot; | [optional] 
**AgreementDocuments** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments**](GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments.md) |  | [optional] 
**AmendmentContext** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**DurationConfig** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanDurationConfig**](GcpMarketplaceResellerPrivateOfferPlanDurationConfig.md) |  | [optional] 
**Features** | Pointer to [**[]GcpMarketplaceProductFeatureValue**](GcpMarketplaceProductFeatureValue.md) |  | [optional] 
**InstallmentTimelineTemplate** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate**](GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate.md) |  | [optional] 
**IsvInfo** | Pointer to [**GcpMarketplaceIsvInfo**](GcpMarketplaceIsvInfo.md) |  | [optional] 
**Margin** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanMargin**](GcpMarketplaceResellerPrivateOfferPlanMargin.md) |  | [optional] 
**MetaInfo** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanMetainfo**](GcpMarketplaceResellerPrivateOfferPlanMetainfo.md) |  | [optional] 
**Name** | Pointer to **string** | In format of \&quot;projects/{projectNumber}/partnerAccounts/{partnerAccountId}/resellerPrivateOfferPlans/{resellerPrivateOfferPlanId}\&quot; | [optional] 
**OfferTemplatePolicies** | Pointer to [**GcpMarketplaceOfferTemplatePolicies**](GcpMarketplaceOfferTemplatePolicies.md) |  | [optional] 
**OfferTermTemplate** | Pointer to [**GcpMarketplacePrivateOfferTermTemplate**](GcpMarketplacePrivateOfferTermTemplate.md) |  | [optional] 
**PaymentSchedule** | Pointer to [**PaymentScheduleType**](PaymentScheduleType.md) |  | [optional] 
**PriceModelTemplate** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate**](GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate.md) | Nill if this reseller private offer plan has installmentTimelineTemplate (payment installments). | [optional] 
**ProductInfo** | Pointer to [**GcpMarketplaceProductInfo**](GcpMarketplaceProductInfo.md) |  | [optional] 
**ReplacementMetadata** | Pointer to **map[string]interface{}** |  | [optional] 
**ResellOfferTemplate** | Pointer to **string** | in format of \&quot;resellOfferTemplates/{resellOfferTemplateId}\&quot; | [optional] 
**ResellerInfo** | Pointer to [**GcpMarketplaceResellerInfo**](GcpMarketplaceResellerInfo.md) |  | [optional] 
**ReusePolicy** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanReusePolicy**](GcpMarketplaceResellerPrivateOfferPlanReusePolicy.md) |  | [optional] 
**StartPolicy** | Pointer to [**GcpMarketplaceStartPolicy**](GcpMarketplaceStartPolicy.md) |  | [optional] 
**State** | Pointer to [**GcpMarketplaceResellerPrivateOfferPlanState**](GcpMarketplaceResellerPrivateOfferPlanState.md) |  | [optional] 
**StateTransitions** | Pointer to [**[]GcpMarketplaceResellerPrivateOfferPlanStateTransition**](GcpMarketplaceResellerPrivateOfferPlanStateTransition.md) |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGcpMarketplaceResellerPrivateOfferPlan

`func NewGcpMarketplaceResellerPrivateOfferPlan() *GcpMarketplaceResellerPrivateOfferPlan`

NewGcpMarketplaceResellerPrivateOfferPlan instantiates a new GcpMarketplaceResellerPrivateOfferPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceResellerPrivateOfferPlanWithDefaults

`func NewGcpMarketplaceResellerPrivateOfferPlanWithDefaults() *GcpMarketplaceResellerPrivateOfferPlan`

NewGcpMarketplaceResellerPrivateOfferPlanWithDefaults instantiates a new GcpMarketplaceResellerPrivateOfferPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptanceDeadlineTime

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAcceptanceDeadlineTime() time.Time`

GetAcceptanceDeadlineTime returns the AcceptanceDeadlineTime field if non-nil, zero value otherwise.

### GetAcceptanceDeadlineTimeOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAcceptanceDeadlineTimeOk() (*time.Time, bool)`

GetAcceptanceDeadlineTimeOk returns a tuple with the AcceptanceDeadlineTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptanceDeadlineTime

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAcceptanceDeadlineTime(v time.Time)`

SetAcceptanceDeadlineTime sets AcceptanceDeadlineTime field to given value.

### HasAcceptanceDeadlineTime

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAcceptanceDeadlineTime() bool`

HasAcceptanceDeadlineTime returns a boolean if a field has been set.

### GetAgreement

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAgreementDocuments

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreementDocuments() GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments`

GetAgreementDocuments returns the AgreementDocuments field if non-nil, zero value otherwise.

### GetAgreementDocumentsOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAgreementDocumentsOk() (*GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments, bool)`

GetAgreementDocumentsOk returns a tuple with the AgreementDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementDocuments

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAgreementDocuments(v GcpMarketplaceResellerPrivateOfferPlanAgreementDocuments)`

SetAgreementDocuments sets AgreementDocuments field to given value.

### HasAgreementDocuments

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAgreementDocuments() bool`

HasAgreementDocuments returns a boolean if a field has been set.

### GetAmendmentContext

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAmendmentContext() map[string]interface{}`

GetAmendmentContext returns the AmendmentContext field if non-nil, zero value otherwise.

### GetAmendmentContextOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetAmendmentContextOk() (*map[string]interface{}, bool)`

GetAmendmentContextOk returns a tuple with the AmendmentContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendmentContext

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetAmendmentContext(v map[string]interface{})`

SetAmendmentContext sets AmendmentContext field to given value.

### HasAmendmentContext

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasAmendmentContext() bool`

HasAmendmentContext returns a boolean if a field has been set.

### GetDisplayName

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDurationConfig

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDurationConfig() GcpMarketplaceResellerPrivateOfferPlanDurationConfig`

GetDurationConfig returns the DurationConfig field if non-nil, zero value otherwise.

### GetDurationConfigOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetDurationConfigOk() (*GcpMarketplaceResellerPrivateOfferPlanDurationConfig, bool)`

GetDurationConfigOk returns a tuple with the DurationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationConfig

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetDurationConfig(v GcpMarketplaceResellerPrivateOfferPlanDurationConfig)`

SetDurationConfig sets DurationConfig field to given value.

### HasDurationConfig

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasDurationConfig() bool`

HasDurationConfig returns a boolean if a field has been set.

### GetFeatures

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetFeatures() []GcpMarketplaceProductFeatureValue`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetFeaturesOk() (*[]GcpMarketplaceProductFeatureValue, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetFeatures(v []GcpMarketplaceProductFeatureValue)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetInstallmentTimelineTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetInstallmentTimelineTemplate() GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate`

GetInstallmentTimelineTemplate returns the InstallmentTimelineTemplate field if non-nil, zero value otherwise.

### GetInstallmentTimelineTemplateOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetInstallmentTimelineTemplateOk() (*GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate, bool)`

GetInstallmentTimelineTemplateOk returns a tuple with the InstallmentTimelineTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentTimelineTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetInstallmentTimelineTemplate(v GcpMarketplaceResellerPrivateOfferPlanInstallmentTimelineTemplate)`

SetInstallmentTimelineTemplate sets InstallmentTimelineTemplate field to given value.

### HasInstallmentTimelineTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasInstallmentTimelineTemplate() bool`

HasInstallmentTimelineTemplate returns a boolean if a field has been set.

### GetIsvInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetIsvInfo() GcpMarketplaceIsvInfo`

GetIsvInfo returns the IsvInfo field if non-nil, zero value otherwise.

### GetIsvInfoOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetIsvInfoOk() (*GcpMarketplaceIsvInfo, bool)`

GetIsvInfoOk returns a tuple with the IsvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetIsvInfo(v GcpMarketplaceIsvInfo)`

SetIsvInfo sets IsvInfo field to given value.

### HasIsvInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasIsvInfo() bool`

HasIsvInfo returns a boolean if a field has been set.

### GetMargin

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMargin() GcpMarketplaceResellerPrivateOfferPlanMargin`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMarginOk() (*GcpMarketplaceResellerPrivateOfferPlanMargin, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetMargin(v GcpMarketplaceResellerPrivateOfferPlanMargin)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMetaInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMetaInfo() GcpMarketplaceResellerPrivateOfferPlanMetainfo`

GetMetaInfo returns the MetaInfo field if non-nil, zero value otherwise.

### GetMetaInfoOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetMetaInfoOk() (*GcpMarketplaceResellerPrivateOfferPlanMetainfo, bool)`

GetMetaInfoOk returns a tuple with the MetaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetMetaInfo(v GcpMarketplaceResellerPrivateOfferPlanMetainfo)`

SetMetaInfo sets MetaInfo field to given value.

### HasMetaInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasMetaInfo() bool`

HasMetaInfo returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferTemplatePolicies

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTemplatePolicies() GcpMarketplaceOfferTemplatePolicies`

GetOfferTemplatePolicies returns the OfferTemplatePolicies field if non-nil, zero value otherwise.

### GetOfferTemplatePoliciesOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTemplatePoliciesOk() (*GcpMarketplaceOfferTemplatePolicies, bool)`

GetOfferTemplatePoliciesOk returns a tuple with the OfferTemplatePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferTemplatePolicies

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetOfferTemplatePolicies(v GcpMarketplaceOfferTemplatePolicies)`

SetOfferTemplatePolicies sets OfferTemplatePolicies field to given value.

### HasOfferTemplatePolicies

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasOfferTemplatePolicies() bool`

HasOfferTemplatePolicies returns a boolean if a field has been set.

### GetOfferTermTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTermTemplate() GcpMarketplacePrivateOfferTermTemplate`

GetOfferTermTemplate returns the OfferTermTemplate field if non-nil, zero value otherwise.

### GetOfferTermTemplateOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetOfferTermTemplateOk() (*GcpMarketplacePrivateOfferTermTemplate, bool)`

GetOfferTermTemplateOk returns a tuple with the OfferTermTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferTermTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetOfferTermTemplate(v GcpMarketplacePrivateOfferTermTemplate)`

SetOfferTermTemplate sets OfferTermTemplate field to given value.

### HasOfferTermTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasOfferTermTemplate() bool`

HasOfferTermTemplate returns a boolean if a field has been set.

### GetPaymentSchedule

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPaymentSchedule() PaymentScheduleType`

GetPaymentSchedule returns the PaymentSchedule field if non-nil, zero value otherwise.

### GetPaymentScheduleOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPaymentScheduleOk() (*PaymentScheduleType, bool)`

GetPaymentScheduleOk returns a tuple with the PaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedule

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetPaymentSchedule(v PaymentScheduleType)`

SetPaymentSchedule sets PaymentSchedule field to given value.

### HasPaymentSchedule

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasPaymentSchedule() bool`

HasPaymentSchedule returns a boolean if a field has been set.

### GetPriceModelTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPriceModelTemplate() GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate`

GetPriceModelTemplate returns the PriceModelTemplate field if non-nil, zero value otherwise.

### GetPriceModelTemplateOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetPriceModelTemplateOk() (*GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate, bool)`

GetPriceModelTemplateOk returns a tuple with the PriceModelTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModelTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetPriceModelTemplate(v GcpMarketplaceResellerPrivateOfferPlanPriceModelTemplate)`

SetPriceModelTemplate sets PriceModelTemplate field to given value.

### HasPriceModelTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasPriceModelTemplate() bool`

HasPriceModelTemplate returns a boolean if a field has been set.

### GetProductInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetProductInfo() GcpMarketplaceProductInfo`

GetProductInfo returns the ProductInfo field if non-nil, zero value otherwise.

### GetProductInfoOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetProductInfoOk() (*GcpMarketplaceProductInfo, bool)`

GetProductInfoOk returns a tuple with the ProductInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetProductInfo(v GcpMarketplaceProductInfo)`

SetProductInfo sets ProductInfo field to given value.

### HasProductInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasProductInfo() bool`

HasProductInfo returns a boolean if a field has been set.

### GetReplacementMetadata

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReplacementMetadata() map[string]interface{}`

GetReplacementMetadata returns the ReplacementMetadata field if non-nil, zero value otherwise.

### GetReplacementMetadataOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReplacementMetadataOk() (*map[string]interface{}, bool)`

GetReplacementMetadataOk returns a tuple with the ReplacementMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementMetadata

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetReplacementMetadata(v map[string]interface{})`

SetReplacementMetadata sets ReplacementMetadata field to given value.

### HasReplacementMetadata

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasReplacementMetadata() bool`

HasReplacementMetadata returns a boolean if a field has been set.

### GetResellOfferTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellOfferTemplate() string`

GetResellOfferTemplate returns the ResellOfferTemplate field if non-nil, zero value otherwise.

### GetResellOfferTemplateOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellOfferTemplateOk() (*string, bool)`

GetResellOfferTemplateOk returns a tuple with the ResellOfferTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellOfferTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetResellOfferTemplate(v string)`

SetResellOfferTemplate sets ResellOfferTemplate field to given value.

### HasResellOfferTemplate

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasResellOfferTemplate() bool`

HasResellOfferTemplate returns a boolean if a field has been set.

### GetResellerInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellerInfo() GcpMarketplaceResellerInfo`

GetResellerInfo returns the ResellerInfo field if non-nil, zero value otherwise.

### GetResellerInfoOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetResellerInfoOk() (*GcpMarketplaceResellerInfo, bool)`

GetResellerInfoOk returns a tuple with the ResellerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetResellerInfo(v GcpMarketplaceResellerInfo)`

SetResellerInfo sets ResellerInfo field to given value.

### HasResellerInfo

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasResellerInfo() bool`

HasResellerInfo returns a boolean if a field has been set.

### GetReusePolicy

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReusePolicy() GcpMarketplaceResellerPrivateOfferPlanReusePolicy`

GetReusePolicy returns the ReusePolicy field if non-nil, zero value otherwise.

### GetReusePolicyOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetReusePolicyOk() (*GcpMarketplaceResellerPrivateOfferPlanReusePolicy, bool)`

GetReusePolicyOk returns a tuple with the ReusePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusePolicy

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetReusePolicy(v GcpMarketplaceResellerPrivateOfferPlanReusePolicy)`

SetReusePolicy sets ReusePolicy field to given value.

### HasReusePolicy

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasReusePolicy() bool`

HasReusePolicy returns a boolean if a field has been set.

### GetStartPolicy

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStartPolicy() GcpMarketplaceStartPolicy`

GetStartPolicy returns the StartPolicy field if non-nil, zero value otherwise.

### GetStartPolicyOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStartPolicyOk() (*GcpMarketplaceStartPolicy, bool)`

GetStartPolicyOk returns a tuple with the StartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPolicy

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetStartPolicy(v GcpMarketplaceStartPolicy)`

SetStartPolicy sets StartPolicy field to given value.

### HasStartPolicy

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasStartPolicy() bool`

HasStartPolicy returns a boolean if a field has been set.

### GetState

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetState() GcpMarketplaceResellerPrivateOfferPlanState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStateOk() (*GcpMarketplaceResellerPrivateOfferPlanState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetState(v GcpMarketplaceResellerPrivateOfferPlanState)`

SetState sets State field to given value.

### HasState

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateTransitions

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStateTransitions() []GcpMarketplaceResellerPrivateOfferPlanStateTransition`

GetStateTransitions returns the StateTransitions field if non-nil, zero value otherwise.

### GetStateTransitionsOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetStateTransitionsOk() (*[]GcpMarketplaceResellerPrivateOfferPlanStateTransition, bool)`

GetStateTransitionsOk returns a tuple with the StateTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTransitions

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetStateTransitions(v []GcpMarketplaceResellerPrivateOfferPlanStateTransition)`

SetStateTransitions sets StateTransitions field to given value.

### HasStateTransitions

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasStateTransitions() bool`

HasStateTransitions returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GcpMarketplaceResellerPrivateOfferPlan) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GcpMarketplaceResellerPrivateOfferPlan) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GcpMarketplaceResellerPrivateOfferPlan) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


