# AzureProductVariantPriceSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTimeRange** | Pointer to [**AzureLocalizedTimeRange**](AzureLocalizedTimeRange.md) |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**IsBaseSchedule** | Pointer to **bool** | There is only one base schedule. | [optional] 
**MarketCodes** | Pointer to **[]string** | ISO country code | [optional] 
**Schedules** | Pointer to [**[]AzurePriceSchedule**](AzurePriceSchedule.md) |  | [optional] 

## Methods

### NewAzureProductVariantPriceSchedule

`func NewAzureProductVariantPriceSchedule() *AzureProductVariantPriceSchedule`

NewAzureProductVariantPriceSchedule instantiates a new AzureProductVariantPriceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureProductVariantPriceScheduleWithDefaults

`func NewAzureProductVariantPriceScheduleWithDefaults() *AzureProductVariantPriceSchedule`

NewAzureProductVariantPriceScheduleWithDefaults instantiates a new AzureProductVariantPriceSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTimeRange

`func (o *AzureProductVariantPriceSchedule) GetDateTimeRange() AzureLocalizedTimeRange`

GetDateTimeRange returns the DateTimeRange field if non-nil, zero value otherwise.

### GetDateTimeRangeOk

`func (o *AzureProductVariantPriceSchedule) GetDateTimeRangeOk() (*AzureLocalizedTimeRange, bool)`

GetDateTimeRangeOk returns a tuple with the DateTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeRange

`func (o *AzureProductVariantPriceSchedule) SetDateTimeRange(v AzureLocalizedTimeRange)`

SetDateTimeRange sets DateTimeRange field to given value.

### HasDateTimeRange

`func (o *AzureProductVariantPriceSchedule) HasDateTimeRange() bool`

HasDateTimeRange returns a boolean if a field has been set.

### GetFriendlyName

`func (o *AzureProductVariantPriceSchedule) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AzureProductVariantPriceSchedule) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AzureProductVariantPriceSchedule) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AzureProductVariantPriceSchedule) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetIsBaseSchedule

`func (o *AzureProductVariantPriceSchedule) GetIsBaseSchedule() bool`

GetIsBaseSchedule returns the IsBaseSchedule field if non-nil, zero value otherwise.

### GetIsBaseScheduleOk

`func (o *AzureProductVariantPriceSchedule) GetIsBaseScheduleOk() (*bool, bool)`

GetIsBaseScheduleOk returns a tuple with the IsBaseSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBaseSchedule

`func (o *AzureProductVariantPriceSchedule) SetIsBaseSchedule(v bool)`

SetIsBaseSchedule sets IsBaseSchedule field to given value.

### HasIsBaseSchedule

`func (o *AzureProductVariantPriceSchedule) HasIsBaseSchedule() bool`

HasIsBaseSchedule returns a boolean if a field has been set.

### GetMarketCodes

`func (o *AzureProductVariantPriceSchedule) GetMarketCodes() []string`

GetMarketCodes returns the MarketCodes field if non-nil, zero value otherwise.

### GetMarketCodesOk

`func (o *AzureProductVariantPriceSchedule) GetMarketCodesOk() (*[]string, bool)`

GetMarketCodesOk returns a tuple with the MarketCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCodes

`func (o *AzureProductVariantPriceSchedule) SetMarketCodes(v []string)`

SetMarketCodes sets MarketCodes field to given value.

### HasMarketCodes

`func (o *AzureProductVariantPriceSchedule) HasMarketCodes() bool`

HasMarketCodes returns a boolean if a field has been set.

### GetSchedules

`func (o *AzureProductVariantPriceSchedule) GetSchedules() []AzurePriceSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *AzureProductVariantPriceSchedule) GetSchedulesOk() (*[]AzurePriceSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *AzureProductVariantPriceSchedule) SetSchedules(v []AzurePriceSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *AzureProductVariantPriceSchedule) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


