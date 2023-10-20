# GcpMarketplacePrivateOfferTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableScheduledStartTimes** | Pointer to **bool** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**StartPolicy** | Pointer to **string** | such as \&quot;OFFER_START_POLICY_IMMEDIATE\&quot; | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**TermDuration** | Pointer to [**GcpPeriodDuration**](GcpPeriodDuration.md) |  | [optional] 

## Methods

### NewGcpMarketplacePrivateOfferTerm

`func NewGcpMarketplacePrivateOfferTerm() *GcpMarketplacePrivateOfferTerm`

NewGcpMarketplacePrivateOfferTerm instantiates a new GcpMarketplacePrivateOfferTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplacePrivateOfferTermWithDefaults

`func NewGcpMarketplacePrivateOfferTermWithDefaults() *GcpMarketplacePrivateOfferTerm`

NewGcpMarketplacePrivateOfferTermWithDefaults instantiates a new GcpMarketplacePrivateOfferTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableScheduledStartTimes

`func (o *GcpMarketplacePrivateOfferTerm) GetEnableScheduledStartTimes() bool`

GetEnableScheduledStartTimes returns the EnableScheduledStartTimes field if non-nil, zero value otherwise.

### GetEnableScheduledStartTimesOk

`func (o *GcpMarketplacePrivateOfferTerm) GetEnableScheduledStartTimesOk() (*bool, bool)`

GetEnableScheduledStartTimesOk returns a tuple with the EnableScheduledStartTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScheduledStartTimes

`func (o *GcpMarketplacePrivateOfferTerm) SetEnableScheduledStartTimes(v bool)`

SetEnableScheduledStartTimes sets EnableScheduledStartTimes field to given value.

### HasEnableScheduledStartTimes

`func (o *GcpMarketplacePrivateOfferTerm) HasEnableScheduledStartTimes() bool`

HasEnableScheduledStartTimes returns a boolean if a field has been set.

### GetEndTime

`func (o *GcpMarketplacePrivateOfferTerm) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GcpMarketplacePrivateOfferTerm) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GcpMarketplacePrivateOfferTerm) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GcpMarketplacePrivateOfferTerm) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartPolicy

`func (o *GcpMarketplacePrivateOfferTerm) GetStartPolicy() string`

GetStartPolicy returns the StartPolicy field if non-nil, zero value otherwise.

### GetStartPolicyOk

`func (o *GcpMarketplacePrivateOfferTerm) GetStartPolicyOk() (*string, bool)`

GetStartPolicyOk returns a tuple with the StartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPolicy

`func (o *GcpMarketplacePrivateOfferTerm) SetStartPolicy(v string)`

SetStartPolicy sets StartPolicy field to given value.

### HasStartPolicy

`func (o *GcpMarketplacePrivateOfferTerm) HasStartPolicy() bool`

HasStartPolicy returns a boolean if a field has been set.

### GetStartTime

`func (o *GcpMarketplacePrivateOfferTerm) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GcpMarketplacePrivateOfferTerm) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GcpMarketplacePrivateOfferTerm) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GcpMarketplacePrivateOfferTerm) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTermDuration

`func (o *GcpMarketplacePrivateOfferTerm) GetTermDuration() GcpPeriodDuration`

GetTermDuration returns the TermDuration field if non-nil, zero value otherwise.

### GetTermDurationOk

`func (o *GcpMarketplacePrivateOfferTerm) GetTermDurationOk() (*GcpPeriodDuration, bool)`

GetTermDurationOk returns a tuple with the TermDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermDuration

`func (o *GcpMarketplacePrivateOfferTerm) SetTermDuration(v GcpPeriodDuration)`

SetTermDuration sets TermDuration field to given value.

### HasTermDuration

`func (o *GcpMarketplacePrivateOfferTerm) HasTermDuration() bool`

HasTermDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


