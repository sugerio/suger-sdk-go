# MeteringUsageRecordGroupMetaInfo

## Properties

 Name                                      | Type                                                               | Description                                                                                                                                                                                              | Notes      
-------------------------------------------|--------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------
 **SkipValidation**                        | Pointer to **bool**                                                | If it is true, the validation of the usage record group is skipped.                                                                                                                                      | [optional] 
 **BillableRecords**                       | Pointer to [**[]MeteringUsageRecord**](MeteringUsageRecord.md)     | for usage metering API v2                                                                                                                                                                                | [optional] 
 **LagoAmount**                            | Pointer to **float32**                                             | The lago amount (in dollars) of the customer. This field keeps the largest of the monthly amount. So it can only be updated when the invoice month increases.                                            | [optional] 
 **LagoSubscriptionID**                    | Pointer to **string**                                              | The lago subscription ID of the customer.                                                                                                                                                                | [optional] 
 **LagoUsageStartTime**                    | Pointer to **time.Time**                                           | The lago usage start time of the customer usage.                                                                                                                                                         | [optional] 
 **MetronomeDailyCostAmount**              | Pointer to **float32**                                             | The metronome daily cost amount (in dollars) of the customer.                                                                                                                                            | [optional] 
 **MetronomeInvoiceID**                    | Pointer to **string**                                              | The metronome invoice ID of the customer.                                                                                                                                                                | [optional] 
 **MetronomeMonthlyInvoiceAmount**         | Pointer to **float32**                                             | The metronome monthly invoice amount (in dollars) of the customer. This field keeps the largest amount of the invoice month. So it can only be updated when the invoice month increases.                 | [optional] 
 **MetronomeMonthlyInvoiceAmountAdjusted** | Pointer to **float32**                                             | The metronome monthly invoice amount (in dollars) of the customer, which is adjusted by the seller. This field is populated only when the invoice amount is decreased by the seller via credit granting. | [optional] 
 **OriginRecords**                         | Pointer to **map[string]float32**                                  | The original records reported by the customer before convertion. If no dimension mapping is applied, this field is the same as the records field.                                                        | [optional] 
 **Source**                                | Pointer to [**UsageRecordGroupSource**](UsageRecordGroupSource.md) | The source of the usage record group. Can be from Suger API or other third party services, such as Metronome.                                                                                            | [optional] 
 **StripeInvoiceID**                       | Pointer to **string**                                              |                                                                                                                                                                                                          | [optional] 
 **StripePeriodEndTime**                   | Pointer to **time.Time**                                           | The stripe period end time of the summary or invoice. UTC time in format \&quot;YYYY-MM-DDTHH:MM:SSZ\&quot;.                                                                                             | [optional] 
 **StripePeriodStartTime**                 | Pointer to **time.Time**                                           | The stripe period start time of the summary or invoice. UTC time in format \&quot;YYYY-MM-DDTHH:MM:SSZ\&quot;.                                                                                           | [optional] 
 **StripeSubscriptionItemID**              | Pointer to **string**                                              |                                                                                                                                                                                                          | [optional] 
 **StripeUsageRecordSummaryID**            | Pointer to **string**                                              |                                                                                                                                                                                                          | [optional] 
 **StripeUsageRecordSummaryTotalUsage**    | Pointer to **int32**                                               |                                                                                                                                                                                                          | [optional] 
 **Timestamp**                             | Pointer to **time.Time**                                           | The timestamp (UTC)) of when the usage records were generated. Optional, if not provided, the current report timestamp will be used.                                                                     | [optional] 

## Methods

### NewMeteringUsageRecordGroupMetaInfo

`func NewMeteringUsageRecordGroupMetaInfo() *MeteringUsageRecordGroupMetaInfo`

NewMeteringUsageRecordGroupMetaInfo instantiates a new MeteringUsageRecordGroupMetaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringUsageRecordGroupMetaInfoWithDefaults

`func NewMeteringUsageRecordGroupMetaInfoWithDefaults() *MeteringUsageRecordGroupMetaInfo`

NewMeteringUsageRecordGroupMetaInfoWithDefaults instantiates a new MeteringUsageRecordGroupMetaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipValidation

`func (o *MeteringUsageRecordGroupMetaInfo) GetSkipValidation() bool`

GetSkipValidation returns the SkipValidation field if non-nil, zero value otherwise.

### GetSkipValidationOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetSkipValidationOk() (*bool, bool)`

GetSkipValidationOk returns a tuple with the SkipValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipValidation

`func (o *MeteringUsageRecordGroupMetaInfo) SetSkipValidation(v bool)`

SetSkipValidation sets SkipValidation field to given value.

### HasSkipValidation

`func (o *MeteringUsageRecordGroupMetaInfo) HasSkipValidation() bool`

HasSkipValidation returns a boolean if a field has been set.

### GetBillableRecords

`func (o *MeteringUsageRecordGroupMetaInfo) GetBillableRecords() []MeteringUsageRecord`

GetBillableRecords returns the BillableRecords field if non-nil, zero value otherwise.

### GetBillableRecordsOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetBillableRecordsOk() (*[]MeteringUsageRecord, bool)`

GetBillableRecordsOk returns a tuple with the BillableRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableRecords

`func (o *MeteringUsageRecordGroupMetaInfo) SetBillableRecords(v []MeteringUsageRecord)`

SetBillableRecords sets BillableRecords field to given value.

### HasBillableRecords

`func (o *MeteringUsageRecordGroupMetaInfo) HasBillableRecords() bool`

HasBillableRecords returns a boolean if a field has been set.

### GetLagoAmount

`func (o *MeteringUsageRecordGroupMetaInfo) GetLagoAmount() float32`

GetLagoAmount returns the LagoAmount field if non-nil, zero value otherwise.

### GetLagoAmountOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetLagoAmountOk() (*float32, bool)`

GetLagoAmountOk returns a tuple with the LagoAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoAmount

`func (o *MeteringUsageRecordGroupMetaInfo) SetLagoAmount(v float32)`

SetLagoAmount sets LagoAmount field to given value.

### HasLagoAmount

`func (o *MeteringUsageRecordGroupMetaInfo) HasLagoAmount() bool`

HasLagoAmount returns a boolean if a field has been set.

### GetLagoSubscriptionID

`func (o *MeteringUsageRecordGroupMetaInfo) GetLagoSubscriptionID() string`

GetLagoSubscriptionID returns the LagoSubscriptionID field if non-nil, zero value otherwise.

### GetLagoSubscriptionIDOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetLagoSubscriptionIDOk() (*string, bool)`

GetLagoSubscriptionIDOk returns a tuple with the LagoSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoSubscriptionID

`func (o *MeteringUsageRecordGroupMetaInfo) SetLagoSubscriptionID(v string)`

SetLagoSubscriptionID sets LagoSubscriptionID field to given value.

### HasLagoSubscriptionID

`func (o *MeteringUsageRecordGroupMetaInfo) HasLagoSubscriptionID() bool`

HasLagoSubscriptionID returns a boolean if a field has been set.

### GetLagoUsageStartTime

`func (o *MeteringUsageRecordGroupMetaInfo) GetLagoUsageStartTime() time.Time`

GetLagoUsageStartTime returns the LagoUsageStartTime field if non-nil, zero value otherwise.

### GetLagoUsageStartTimeOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetLagoUsageStartTimeOk() (*time.Time, bool)`

GetLagoUsageStartTimeOk returns a tuple with the LagoUsageStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagoUsageStartTime

`func (o *MeteringUsageRecordGroupMetaInfo) SetLagoUsageStartTime(v time.Time)`

SetLagoUsageStartTime sets LagoUsageStartTime field to given value.

### HasLagoUsageStartTime

`func (o *MeteringUsageRecordGroupMetaInfo) HasLagoUsageStartTime() bool`

HasLagoUsageStartTime returns a boolean if a field has been set.

### GetMetronomeDailyCostAmount

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeDailyCostAmount() float32`

GetMetronomeDailyCostAmount returns the MetronomeDailyCostAmount field if non-nil, zero value otherwise.

### GetMetronomeDailyCostAmountOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeDailyCostAmountOk() (*float32, bool)`

GetMetronomeDailyCostAmountOk returns a tuple with the MetronomeDailyCostAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeDailyCostAmount

`func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeDailyCostAmount(v float32)`

SetMetronomeDailyCostAmount sets MetronomeDailyCostAmount field to given value.

### HasMetronomeDailyCostAmount

`func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeDailyCostAmount() bool`

HasMetronomeDailyCostAmount returns a boolean if a field has been set.

### GetMetronomeInvoiceID

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeInvoiceID() string`

GetMetronomeInvoiceID returns the MetronomeInvoiceID field if non-nil, zero value otherwise.

### GetMetronomeInvoiceIDOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeInvoiceIDOk() (*string, bool)`

GetMetronomeInvoiceIDOk returns a tuple with the MetronomeInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeInvoiceID

`func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeInvoiceID(v string)`

SetMetronomeInvoiceID sets MetronomeInvoiceID field to given value.

### HasMetronomeInvoiceID

`func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeInvoiceID() bool`

HasMetronomeInvoiceID returns a boolean if a field has been set.

### GetMetronomeMonthlyInvoiceAmount

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmount() float32`

GetMetronomeMonthlyInvoiceAmount returns the MetronomeMonthlyInvoiceAmount field if non-nil, zero value otherwise.

### GetMetronomeMonthlyInvoiceAmountOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmountOk() (*float32, bool)`

GetMetronomeMonthlyInvoiceAmountOk returns a tuple with the MetronomeMonthlyInvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeMonthlyInvoiceAmount

`func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeMonthlyInvoiceAmount(v float32)`

SetMetronomeMonthlyInvoiceAmount sets MetronomeMonthlyInvoiceAmount field to given value.

### HasMetronomeMonthlyInvoiceAmount

`func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeMonthlyInvoiceAmount() bool`

HasMetronomeMonthlyInvoiceAmount returns a boolean if a field has been set.

### GetMetronomeMonthlyInvoiceAmountAdjusted

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmountAdjusted() float32`

GetMetronomeMonthlyInvoiceAmountAdjusted returns the MetronomeMonthlyInvoiceAmountAdjusted field if non-nil, zero value otherwise.

### GetMetronomeMonthlyInvoiceAmountAdjustedOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetMetronomeMonthlyInvoiceAmountAdjustedOk() (*float32, bool)`

GetMetronomeMonthlyInvoiceAmountAdjustedOk returns a tuple with the MetronomeMonthlyInvoiceAmountAdjusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetronomeMonthlyInvoiceAmountAdjusted

`func (o *MeteringUsageRecordGroupMetaInfo) SetMetronomeMonthlyInvoiceAmountAdjusted(v float32)`

SetMetronomeMonthlyInvoiceAmountAdjusted sets MetronomeMonthlyInvoiceAmountAdjusted field to given value.

### HasMetronomeMonthlyInvoiceAmountAdjusted

`func (o *MeteringUsageRecordGroupMetaInfo) HasMetronomeMonthlyInvoiceAmountAdjusted() bool`

HasMetronomeMonthlyInvoiceAmountAdjusted returns a boolean if a field has been set.

### GetOriginRecords

`func (o *MeteringUsageRecordGroupMetaInfo) GetOriginRecords() map[string]float32`

GetOriginRecords returns the OriginRecords field if non-nil, zero value otherwise.

### GetOriginRecordsOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetOriginRecordsOk() (*map[string]float32, bool)`

GetOriginRecordsOk returns a tuple with the OriginRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRecords

`func (o *MeteringUsageRecordGroupMetaInfo) SetOriginRecords(v map[string]float32)`

SetOriginRecords sets OriginRecords field to given value.

### HasOriginRecords

`func (o *MeteringUsageRecordGroupMetaInfo) HasOriginRecords() bool`

HasOriginRecords returns a boolean if a field has been set.

### GetSource

`func (o *MeteringUsageRecordGroupMetaInfo) GetSource() UsageRecordGroupSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetSourceOk() (*UsageRecordGroupSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MeteringUsageRecordGroupMetaInfo) SetSource(v UsageRecordGroupSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MeteringUsageRecordGroupMetaInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStripeInvoiceID

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeInvoiceID() string`

GetStripeInvoiceID returns the StripeInvoiceID field if non-nil, zero value otherwise.

### GetStripeInvoiceIDOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeInvoiceIDOk() (*string, bool)`

GetStripeInvoiceIDOk returns a tuple with the StripeInvoiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeInvoiceID

`func (o *MeteringUsageRecordGroupMetaInfo) SetStripeInvoiceID(v string)`

SetStripeInvoiceID sets StripeInvoiceID field to given value.

### HasStripeInvoiceID

`func (o *MeteringUsageRecordGroupMetaInfo) HasStripeInvoiceID() bool`

HasStripeInvoiceID returns a boolean if a field has been set.

### GetStripePeriodEndTime

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripePeriodEndTime() time.Time`

GetStripePeriodEndTime returns the StripePeriodEndTime field if non-nil, zero value otherwise.

### GetStripePeriodEndTimeOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripePeriodEndTimeOk() (*time.Time, bool)`

GetStripePeriodEndTimeOk returns a tuple with the StripePeriodEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePeriodEndTime

`func (o *MeteringUsageRecordGroupMetaInfo) SetStripePeriodEndTime(v time.Time)`

SetStripePeriodEndTime sets StripePeriodEndTime field to given value.

### HasStripePeriodEndTime

`func (o *MeteringUsageRecordGroupMetaInfo) HasStripePeriodEndTime() bool`

HasStripePeriodEndTime returns a boolean if a field has been set.

### GetStripePeriodStartTime

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripePeriodStartTime() time.Time`

GetStripePeriodStartTime returns the StripePeriodStartTime field if non-nil, zero value otherwise.

### GetStripePeriodStartTimeOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripePeriodStartTimeOk() (*time.Time, bool)`

GetStripePeriodStartTimeOk returns a tuple with the StripePeriodStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripePeriodStartTime

`func (o *MeteringUsageRecordGroupMetaInfo) SetStripePeriodStartTime(v time.Time)`

SetStripePeriodStartTime sets StripePeriodStartTime field to given value.

### HasStripePeriodStartTime

`func (o *MeteringUsageRecordGroupMetaInfo) HasStripePeriodStartTime() bool`

HasStripePeriodStartTime returns a boolean if a field has been set.

### GetStripeSubscriptionItemID

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeSubscriptionItemID() string`

GetStripeSubscriptionItemID returns the StripeSubscriptionItemID field if non-nil, zero value otherwise.

### GetStripeSubscriptionItemIDOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeSubscriptionItemIDOk() (*string, bool)`

GetStripeSubscriptionItemIDOk returns a tuple with the StripeSubscriptionItemID field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetStripeSubscriptionItemID

`func (o *MeteringUsageRecordGroupMetaInfo) SetStripeSubscriptionItemID(v string)`

SetStripeSubscriptionItemID sets StripeSubscriptionItemID field to given value.

### HasStripeSubscriptionItemID

`func (o *MeteringUsageRecordGroupMetaInfo) HasStripeSubscriptionItemID() bool`

HasStripeSubscriptionItemID returns a boolean if a field has been set.

### GetStripeUsageRecordSummaryID

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeUsageRecordSummaryID() string`

GetStripeUsageRecordSummaryID returns the StripeUsageRecordSummaryID field if non-nil, zero value otherwise.

### GetStripeUsageRecordSummaryIDOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeUsageRecordSummaryIDOk() (*string, bool)`

GetStripeUsageRecordSummaryIDOk returns a tuple with the StripeUsageRecordSummaryID field if it's non-nil, zero value
otherwise
and a boolean to check if the value has been set.

### SetStripeUsageRecordSummaryID

`func (o *MeteringUsageRecordGroupMetaInfo) SetStripeUsageRecordSummaryID(v string)`

SetStripeUsageRecordSummaryID sets StripeUsageRecordSummaryID field to given value.

### HasStripeUsageRecordSummaryID

`func (o *MeteringUsageRecordGroupMetaInfo) HasStripeUsageRecordSummaryID() bool`

HasStripeUsageRecordSummaryID returns a boolean if a field has been set.

### GetStripeUsageRecordSummaryTotalUsage

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeUsageRecordSummaryTotalUsage() int32`

GetStripeUsageRecordSummaryTotalUsage returns the StripeUsageRecordSummaryTotalUsage field if non-nil, zero value
otherwise.

### GetStripeUsageRecordSummaryTotalUsageOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetStripeUsageRecordSummaryTotalUsageOk() (*int32, bool)`

GetStripeUsageRecordSummaryTotalUsageOk returns a tuple with the StripeUsageRecordSummaryTotalUsage field if it's
non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeUsageRecordSummaryTotalUsage

`func (o *MeteringUsageRecordGroupMetaInfo) SetStripeUsageRecordSummaryTotalUsage(v int32)`

SetStripeUsageRecordSummaryTotalUsage sets StripeUsageRecordSummaryTotalUsage field to given value.

### HasStripeUsageRecordSummaryTotalUsage

`func (o *MeteringUsageRecordGroupMetaInfo) HasStripeUsageRecordSummaryTotalUsage() bool`

HasStripeUsageRecordSummaryTotalUsage returns a boolean if a field has been set.

### GetTimestamp

`func (o *MeteringUsageRecordGroupMetaInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MeteringUsageRecordGroupMetaInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MeteringUsageRecordGroupMetaInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MeteringUsageRecordGroupMetaInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


