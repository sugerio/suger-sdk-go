# UsageMeteringConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerUsageMeteringConfigs** | Pointer to [**[]PartnerUsageMeteringConfig**](PartnerUsageMeteringConfig.md) | The usage metering configuration for each Partner, such as AWS, AZURE &amp; GCP. | [optional] 

## Methods

### NewUsageMeteringConfigInfo

`func NewUsageMeteringConfigInfo() *UsageMeteringConfigInfo`

NewUsageMeteringConfigInfo instantiates a new UsageMeteringConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeteringConfigInfoWithDefaults

`func NewUsageMeteringConfigInfoWithDefaults() *UsageMeteringConfigInfo`

NewUsageMeteringConfigInfoWithDefaults instantiates a new UsageMeteringConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerUsageMeteringConfigs

`func (o *UsageMeteringConfigInfo) GetPartnerUsageMeteringConfigs() []PartnerUsageMeteringConfig`

GetPartnerUsageMeteringConfigs returns the PartnerUsageMeteringConfigs field if non-nil, zero value otherwise.

### GetPartnerUsageMeteringConfigsOk

`func (o *UsageMeteringConfigInfo) GetPartnerUsageMeteringConfigsOk() (*[]PartnerUsageMeteringConfig, bool)`

GetPartnerUsageMeteringConfigsOk returns a tuple with the PartnerUsageMeteringConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerUsageMeteringConfigs

`func (o *UsageMeteringConfigInfo) SetPartnerUsageMeteringConfigs(v []PartnerUsageMeteringConfig)`

SetPartnerUsageMeteringConfigs sets PartnerUsageMeteringConfigs field to given value.

### HasPartnerUsageMeteringConfigs

`func (o *UsageMeteringConfigInfo) HasPartnerUsageMeteringConfigs() bool`

HasPartnerUsageMeteringConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


