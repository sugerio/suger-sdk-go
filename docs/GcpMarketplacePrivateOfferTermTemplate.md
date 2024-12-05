# GcpMarketplacePrivateOfferTermTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableScheduledStartTimes** | Pointer to **bool** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**PaymentRecurrence** | Pointer to **string** |  | [optional] 
**StartPolicy** | Pointer to [**GcpMarketplaceOfferStartPolicy**](GcpMarketplaceOfferStartPolicy.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**TermDuration** | Pointer to [**GcpPeriodDuration**](GcpPeriodDuration.md) |  | [optional] 
**TermDurationConstraint** | Pointer to [**GcpMarketplacePrivateOfferTermDurationConstraint**](GcpMarketplacePrivateOfferTermDurationConstraint.md) |  | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferTermTemplate

`func NewGcpMarketplacePrivateOfferTermTemplate() *GcpMarketplacePrivateOfferTermTemplate`

NewGcpMarketplacePrivateOfferTermTemplate instantiates a new GcpMarketplacePrivateOfferTermTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferTermTemplateWithDefaults

`func NewGcpMarketplacePrivateOfferTermTemplateWithDefaults() *GcpMarketplacePrivateOfferTermTemplate`

NewGcpMarketplacePrivateOfferTermTemplateWithDefaults instantiates a new GcpMarketplacePrivateOfferTermTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableScheduledStartTimes

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetEnableScheduledStartTimes() bool`

GetEnableScheduledStartTimes returns the EnableScheduledStartTimes field if non-nil, zero value otherwise.

### GetEnableScheduledStartTimesOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetEnableScheduledStartTimesOk() (*bool, bool)`

GetEnableScheduledStartTimesOk returns a tuple with the EnableScheduledStartTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScheduledStartTimes

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetEnableScheduledStartTimes(v bool)`

SetEnableScheduledStartTimes sets EnableScheduledStartTimes field to given value.

### HasEnableScheduledStartTimes

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasEnableScheduledStartTimes() bool`

HasEnableScheduledStartTimes returns a boolean if a field has been set.

### GetEndTime

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPaymentRecurrence

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetPaymentRecurrence() string`

GetPaymentRecurrence returns the PaymentRecurrence field if non-nil, zero value otherwise.

### GetPaymentRecurrenceOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetPaymentRecurrenceOk() (*string, bool)`

GetPaymentRecurrenceOk returns a tuple with the PaymentRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRecurrence

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetPaymentRecurrence(v string)`

SetPaymentRecurrence sets PaymentRecurrence field to given value.

### HasPaymentRecurrence

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasPaymentRecurrence() bool`

HasPaymentRecurrence returns a boolean if a field has been set.

### GetStartPolicy

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetStartPolicy() GcpMarketplaceOfferStartPolicy`

GetStartPolicy returns the StartPolicy field if non-nil, zero value otherwise.

### GetStartPolicyOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetStartPolicyOk() (*GcpMarketplaceOfferStartPolicy, bool)`

GetStartPolicyOk returns a tuple with the StartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPolicy

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetStartPolicy(v GcpMarketplaceOfferStartPolicy)`

SetStartPolicy sets StartPolicy field to given value.

### HasStartPolicy

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasStartPolicy() bool`

HasStartPolicy returns a boolean if a field has been set.

### GetStartTime

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTermDuration

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetTermDuration() GcpPeriodDuration`

GetTermDuration returns the TermDuration field if non-nil, zero value otherwise.

### GetTermDurationOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetTermDurationOk() (*GcpPeriodDuration, bool)`

GetTermDurationOk returns a tuple with the TermDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermDuration

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetTermDuration(v GcpPeriodDuration)`

SetTermDuration sets TermDuration field to given value.

### HasTermDuration

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasTermDuration() bool`

HasTermDuration returns a boolean if a field has been set.

### GetTermDurationConstraint

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetTermDurationConstraint() GcpMarketplacePrivateOfferTermDurationConstraint`

GetTermDurationConstraint returns the TermDurationConstraint field if non-nil, zero value otherwise.

### GetTermDurationConstraintOk

`func (o *GcpMarketplacePrivateOfferTermTemplate) GetTermDurationConstraintOk() (*GcpMarketplacePrivateOfferTermDurationConstraint, bool)`

GetTermDurationConstraintOk returns a tuple with the TermDurationConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermDurationConstraint

`func (o *GcpMarketplacePrivateOfferTermTemplate) SetTermDurationConstraint(v GcpMarketplacePrivateOfferTermDurationConstraint)`

SetTermDurationConstraint sets TermDurationConstraint field to given value.

### HasTermDurationConstraint

`func (o *GcpMarketplacePrivateOfferTermTemplate) HasTermDurationConstraint() bool`

HasTermDurationConstraint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


