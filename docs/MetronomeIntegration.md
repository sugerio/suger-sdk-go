# MetronomeIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiToken** | Pointer to **string** | The Bearer token for the metronome API. Required only when the metronome integration is created or updated with new API token. | [optional] 
**BillableMetricFullList** | Pointer to [**[]MetronomeBillableMetric**](MetronomeBillableMetric.md) | The full list of billable metrics fetched from metronome API for the available metronome customers. | [optional] 
**BillableMetricWhitelist** | Pointer to [**[]MetronomeBillableMetric**](MetronomeBillableMetric.md) | The whitelist of billable metrics. Only the metrics in the whitelist will be metered &amp; reported to cloud marketplace. | [optional] 
**EnableAutoReportUsage** | Pointer to **bool** | Whether to enable the auto usage report for the metronome integration. If enabled, cron job runs every hour to fetch usage events from Metronome to Suger as UsageRecordGroups. | [optional] 
**EnableBillableMetricWhitelist** | Pointer to **bool** | Enable whitelist for billable metrics. If enabled, only the metrics in the whitelist will be metered &amp; reported to cloud marketplace. Otherwise all the metrics in the billableMetricFullList will be metered &amp; reported to cloud marketplace. | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to store the ApiToken in AWS Secret manager. For internal usage only. | [optional] 

## Methods

### NewMetronomeIntegration

`func NewMetronomeIntegration() *MetronomeIntegration`

NewMetronomeIntegration instantiates a new MetronomeIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetronomeIntegrationWithDefaults

`func NewMetronomeIntegrationWithDefaults() *MetronomeIntegration`

NewMetronomeIntegrationWithDefaults instantiates a new MetronomeIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiToken

`func (o *MetronomeIntegration) GetApiToken() string`

GetApiToken returns the ApiToken field if non-nil, zero value otherwise.

### GetApiTokenOk

`func (o *MetronomeIntegration) GetApiTokenOk() (*string, bool)`

GetApiTokenOk returns a tuple with the ApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiToken

`func (o *MetronomeIntegration) SetApiToken(v string)`

SetApiToken sets ApiToken field to given value.

### HasApiToken

`func (o *MetronomeIntegration) HasApiToken() bool`

HasApiToken returns a boolean if a field has been set.

### GetBillableMetricFullList

`func (o *MetronomeIntegration) GetBillableMetricFullList() []MetronomeBillableMetric`

GetBillableMetricFullList returns the BillableMetricFullList field if non-nil, zero value otherwise.

### GetBillableMetricFullListOk

`func (o *MetronomeIntegration) GetBillableMetricFullListOk() (*[]MetronomeBillableMetric, bool)`

GetBillableMetricFullListOk returns a tuple with the BillableMetricFullList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricFullList

`func (o *MetronomeIntegration) SetBillableMetricFullList(v []MetronomeBillableMetric)`

SetBillableMetricFullList sets BillableMetricFullList field to given value.

### HasBillableMetricFullList

`func (o *MetronomeIntegration) HasBillableMetricFullList() bool`

HasBillableMetricFullList returns a boolean if a field has been set.

### GetBillableMetricWhitelist

`func (o *MetronomeIntegration) GetBillableMetricWhitelist() []MetronomeBillableMetric`

GetBillableMetricWhitelist returns the BillableMetricWhitelist field if non-nil, zero value otherwise.

### GetBillableMetricWhitelistOk

`func (o *MetronomeIntegration) GetBillableMetricWhitelistOk() (*[]MetronomeBillableMetric, bool)`

GetBillableMetricWhitelistOk returns a tuple with the BillableMetricWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricWhitelist

`func (o *MetronomeIntegration) SetBillableMetricWhitelist(v []MetronomeBillableMetric)`

SetBillableMetricWhitelist sets BillableMetricWhitelist field to given value.

### HasBillableMetricWhitelist

`func (o *MetronomeIntegration) HasBillableMetricWhitelist() bool`

HasBillableMetricWhitelist returns a boolean if a field has been set.

### GetEnableAutoReportUsage

`func (o *MetronomeIntegration) GetEnableAutoReportUsage() bool`

GetEnableAutoReportUsage returns the EnableAutoReportUsage field if non-nil, zero value otherwise.

### GetEnableAutoReportUsageOk

`func (o *MetronomeIntegration) GetEnableAutoReportUsageOk() (*bool, bool)`

GetEnableAutoReportUsageOk returns a tuple with the EnableAutoReportUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoReportUsage

`func (o *MetronomeIntegration) SetEnableAutoReportUsage(v bool)`

SetEnableAutoReportUsage sets EnableAutoReportUsage field to given value.

### HasEnableAutoReportUsage

`func (o *MetronomeIntegration) HasEnableAutoReportUsage() bool`

HasEnableAutoReportUsage returns a boolean if a field has been set.

### GetEnableBillableMetricWhitelist

`func (o *MetronomeIntegration) GetEnableBillableMetricWhitelist() bool`

GetEnableBillableMetricWhitelist returns the EnableBillableMetricWhitelist field if non-nil, zero value otherwise.

### GetEnableBillableMetricWhitelistOk

`func (o *MetronomeIntegration) GetEnableBillableMetricWhitelistOk() (*bool, bool)`

GetEnableBillableMetricWhitelistOk returns a tuple with the EnableBillableMetricWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBillableMetricWhitelist

`func (o *MetronomeIntegration) SetEnableBillableMetricWhitelist(v bool)`

SetEnableBillableMetricWhitelist sets EnableBillableMetricWhitelist field to given value.

### HasEnableBillableMetricWhitelist

`func (o *MetronomeIntegration) HasEnableBillableMetricWhitelist() bool`

HasEnableBillableMetricWhitelist returns a boolean if a field has been set.

### GetSecretKey

`func (o *MetronomeIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *MetronomeIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *MetronomeIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *MetronomeIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


