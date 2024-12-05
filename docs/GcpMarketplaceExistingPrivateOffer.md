# GcpMarketplaceExistingPrivateOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agreement** | Pointer to **string** |  | [optional] 
**CustomEula** | Pointer to [**GcpMarketplaceDocument**](GcpMarketplaceDocument.md) |  | [optional] 
**InstallmentTimeline** | Pointer to [**GcpMarketplacePrivateOfferInstallmentTimeline**](GcpMarketplacePrivateOfferInstallmentTimeline.md) |  | [optional] 
**Name** | Pointer to **string** | GCP private offer resource name. | [optional] 
**OfferTerm** | Pointer to [**GcpMarketplacePrivateOfferTerm**](GcpMarketplacePrivateOfferTerm.md) |  | [optional] 
**PaymentSchedule** | Pointer to [**PaymentScheduleType**](PaymentScheduleType.md) |  | [optional] 
**PriceModel** | Pointer to [**GcpMarketplacePrivateOfferPriceModel**](GcpMarketplacePrivateOfferPriceModel.md) | Nill if the offer has payment installments. | [optional] 
**ServiceLevel** | Pointer to **string** | The Plan of the offer. | [optional] 

## Methods

### NewGcpMarketplaceExistingPrivateOffer

`func NewGcpMarketplaceExistingPrivateOffer() *GcpMarketplaceExistingPrivateOffer`

NewGcpMarketplaceExistingPrivateOffer instantiates a new GcpMarketplaceExistingPrivateOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceExistingPrivateOfferWithDefaults

`func NewGcpMarketplaceExistingPrivateOfferWithDefaults() *GcpMarketplaceExistingPrivateOffer`

NewGcpMarketplaceExistingPrivateOfferWithDefaults instantiates a new GcpMarketplaceExistingPrivateOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreement

`func (o *GcpMarketplaceExistingPrivateOffer) GetAgreement() string`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetAgreementOk() (*string, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *GcpMarketplaceExistingPrivateOffer) SetAgreement(v string)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *GcpMarketplaceExistingPrivateOffer) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetCustomEula

`func (o *GcpMarketplaceExistingPrivateOffer) GetCustomEula() GcpMarketplaceDocument`

GetCustomEula returns the CustomEula field if non-nil, zero value otherwise.

### GetCustomEulaOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetCustomEulaOk() (*GcpMarketplaceDocument, bool)`

GetCustomEulaOk returns a tuple with the CustomEula field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEula

`func (o *GcpMarketplaceExistingPrivateOffer) SetCustomEula(v GcpMarketplaceDocument)`

SetCustomEula sets CustomEula field to given value.

### HasCustomEula

`func (o *GcpMarketplaceExistingPrivateOffer) HasCustomEula() bool`

HasCustomEula returns a boolean if a field has been set.

### GetInstallmentTimeline

`func (o *GcpMarketplaceExistingPrivateOffer) GetInstallmentTimeline() GcpMarketplacePrivateOfferInstallmentTimeline`

GetInstallmentTimeline returns the InstallmentTimeline field if non-nil, zero value otherwise.

### GetInstallmentTimelineOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetInstallmentTimelineOk() (*GcpMarketplacePrivateOfferInstallmentTimeline, bool)`

GetInstallmentTimelineOk returns a tuple with the InstallmentTimeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallmentTimeline

`func (o *GcpMarketplaceExistingPrivateOffer) SetInstallmentTimeline(v GcpMarketplacePrivateOfferInstallmentTimeline)`

SetInstallmentTimeline sets InstallmentTimeline field to given value.

### HasInstallmentTimeline

`func (o *GcpMarketplaceExistingPrivateOffer) HasInstallmentTimeline() bool`

HasInstallmentTimeline returns a boolean if a field has been set.

### GetName

`func (o *GcpMarketplaceExistingPrivateOffer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GcpMarketplaceExistingPrivateOffer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GcpMarketplaceExistingPrivateOffer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferTerm

`func (o *GcpMarketplaceExistingPrivateOffer) GetOfferTerm() GcpMarketplacePrivateOfferTerm`

GetOfferTerm returns the OfferTerm field if non-nil, zero value otherwise.

### GetOfferTermOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetOfferTermOk() (*GcpMarketplacePrivateOfferTerm, bool)`

GetOfferTermOk returns a tuple with the OfferTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferTerm

`func (o *GcpMarketplaceExistingPrivateOffer) SetOfferTerm(v GcpMarketplacePrivateOfferTerm)`

SetOfferTerm sets OfferTerm field to given value.

### HasOfferTerm

`func (o *GcpMarketplaceExistingPrivateOffer) HasOfferTerm() bool`

HasOfferTerm returns a boolean if a field has been set.

### GetPaymentSchedule

`func (o *GcpMarketplaceExistingPrivateOffer) GetPaymentSchedule() PaymentScheduleType`

GetPaymentSchedule returns the PaymentSchedule field if non-nil, zero value otherwise.

### GetPaymentScheduleOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetPaymentScheduleOk() (*PaymentScheduleType, bool)`

GetPaymentScheduleOk returns a tuple with the PaymentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSchedule

`func (o *GcpMarketplaceExistingPrivateOffer) SetPaymentSchedule(v PaymentScheduleType)`

SetPaymentSchedule sets PaymentSchedule field to given value.

### HasPaymentSchedule

`func (o *GcpMarketplaceExistingPrivateOffer) HasPaymentSchedule() bool`

HasPaymentSchedule returns a boolean if a field has been set.

### GetPriceModel

`func (o *GcpMarketplaceExistingPrivateOffer) GetPriceModel() GcpMarketplacePrivateOfferPriceModel`

GetPriceModel returns the PriceModel field if non-nil, zero value otherwise.

### GetPriceModelOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetPriceModelOk() (*GcpMarketplacePrivateOfferPriceModel, bool)`

GetPriceModelOk returns a tuple with the PriceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceModel

`func (o *GcpMarketplaceExistingPrivateOffer) SetPriceModel(v GcpMarketplacePrivateOfferPriceModel)`

SetPriceModel sets PriceModel field to given value.

### HasPriceModel

`func (o *GcpMarketplaceExistingPrivateOffer) HasPriceModel() bool`

HasPriceModel returns a boolean if a field has been set.

### GetServiceLevel

`func (o *GcpMarketplaceExistingPrivateOffer) GetServiceLevel() string`

GetServiceLevel returns the ServiceLevel field if non-nil, zero value otherwise.

### GetServiceLevelOk

`func (o *GcpMarketplaceExistingPrivateOffer) GetServiceLevelOk() (*string, bool)`

GetServiceLevelOk returns a tuple with the ServiceLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevel

`func (o *GcpMarketplaceExistingPrivateOffer) SetServiceLevel(v string)`

SetServiceLevel sets ServiceLevel field to given value.

### HasServiceLevel

`func (o *GcpMarketplaceExistingPrivateOffer) HasServiceLevel() bool`

HasServiceLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


