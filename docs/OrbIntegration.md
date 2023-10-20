# OrbIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | The Bearer key to access the orb API. | [optional] 
**BillableMetricFullList** | Pointer to [**[]OrbBillableMetric**](OrbBillableMetric.md) |  | [optional] 
**BillingMode** | Pointer to [**OrbBillingMode**](OrbBillingMode.md) |  | [optional] 
**EnableAutoReportUsage** | Pointer to **bool** | Whether to enable the auto usage report for the orb integration. If enabled, cron job runs every hour to fetch usage events from Orb to Suger as UsageRecordGroups. | [optional] 
**PlanFullList** | Pointer to [**[]OrbPlan**](OrbPlan.md) |  | [optional] 
**SecretKey** | Pointer to **string** | The secret key used to store the ApiKey in AWS Secret manager. For internal usage only. | [optional] 

## Methods

### NewOrbIntegration

`func NewOrbIntegration() *OrbIntegration`

NewOrbIntegration instantiates a new OrbIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbIntegrationWithDefaults

`func NewOrbIntegrationWithDefaults() *OrbIntegration`

NewOrbIntegrationWithDefaults instantiates a new OrbIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OrbIntegration) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OrbIntegration) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OrbIntegration) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OrbIntegration) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetBillableMetricFullList

`func (o *OrbIntegration) GetBillableMetricFullList() []OrbBillableMetric`

GetBillableMetricFullList returns the BillableMetricFullList field if non-nil, zero value otherwise.

### GetBillableMetricFullListOk

`func (o *OrbIntegration) GetBillableMetricFullListOk() (*[]OrbBillableMetric, bool)`

GetBillableMetricFullListOk returns a tuple with the BillableMetricFullList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableMetricFullList

`func (o *OrbIntegration) SetBillableMetricFullList(v []OrbBillableMetric)`

SetBillableMetricFullList sets BillableMetricFullList field to given value.

### HasBillableMetricFullList

`func (o *OrbIntegration) HasBillableMetricFullList() bool`

HasBillableMetricFullList returns a boolean if a field has been set.

### GetBillingMode

`func (o *OrbIntegration) GetBillingMode() OrbBillingMode`

GetBillingMode returns the BillingMode field if non-nil, zero value otherwise.

### GetBillingModeOk

`func (o *OrbIntegration) GetBillingModeOk() (*OrbBillingMode, bool)`

GetBillingModeOk returns a tuple with the BillingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMode

`func (o *OrbIntegration) SetBillingMode(v OrbBillingMode)`

SetBillingMode sets BillingMode field to given value.

### HasBillingMode

`func (o *OrbIntegration) HasBillingMode() bool`

HasBillingMode returns a boolean if a field has been set.

### GetEnableAutoReportUsage

`func (o *OrbIntegration) GetEnableAutoReportUsage() bool`

GetEnableAutoReportUsage returns the EnableAutoReportUsage field if non-nil, zero value otherwise.

### GetEnableAutoReportUsageOk

`func (o *OrbIntegration) GetEnableAutoReportUsageOk() (*bool, bool)`

GetEnableAutoReportUsageOk returns a tuple with the EnableAutoReportUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoReportUsage

`func (o *OrbIntegration) SetEnableAutoReportUsage(v bool)`

SetEnableAutoReportUsage sets EnableAutoReportUsage field to given value.

### HasEnableAutoReportUsage

`func (o *OrbIntegration) HasEnableAutoReportUsage() bool`

HasEnableAutoReportUsage returns a boolean if a field has been set.

### GetPlanFullList

`func (o *OrbIntegration) GetPlanFullList() []OrbPlan`

GetPlanFullList returns the PlanFullList field if non-nil, zero value otherwise.

### GetPlanFullListOk

`func (o *OrbIntegration) GetPlanFullListOk() (*[]OrbPlan, bool)`

GetPlanFullListOk returns a tuple with the PlanFullList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanFullList

`func (o *OrbIntegration) SetPlanFullList(v []OrbPlan)`

SetPlanFullList sets PlanFullList field to given value.

### HasPlanFullList

`func (o *OrbIntegration) HasPlanFullList() bool`

HasPlanFullList returns a boolean if a field has been set.

### GetSecretKey

`func (o *OrbIntegration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OrbIntegration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OrbIntegration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *OrbIntegration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


