# AzureLocalizedDateTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTimeInUtc** | Pointer to **string** |  | [optional] 
**LocalizePerMarket** | Pointer to **bool** |  | [optional] 

## Methods

### NewAzureLocalizedDateTime

`func NewAzureLocalizedDateTime() *AzureLocalizedDateTime`

NewAzureLocalizedDateTime instantiates a new AzureLocalizedDateTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureLocalizedDateTimeWithDefaults

`func NewAzureLocalizedDateTimeWithDefaults() *AzureLocalizedDateTime`

NewAzureLocalizedDateTimeWithDefaults instantiates a new AzureLocalizedDateTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTimeInUtc

`func (o *AzureLocalizedDateTime) GetDateTimeInUtc() string`

GetDateTimeInUtc returns the DateTimeInUtc field if non-nil, zero value otherwise.

### GetDateTimeInUtcOk

`func (o *AzureLocalizedDateTime) GetDateTimeInUtcOk() (*string, bool)`

GetDateTimeInUtcOk returns a tuple with the DateTimeInUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeInUtc

`func (o *AzureLocalizedDateTime) SetDateTimeInUtc(v string)`

SetDateTimeInUtc sets DateTimeInUtc field to given value.

### HasDateTimeInUtc

`func (o *AzureLocalizedDateTime) HasDateTimeInUtc() bool`

HasDateTimeInUtc returns a boolean if a field has been set.

### GetLocalizePerMarket

`func (o *AzureLocalizedDateTime) GetLocalizePerMarket() bool`

GetLocalizePerMarket returns the LocalizePerMarket field if non-nil, zero value otherwise.

### GetLocalizePerMarketOk

`func (o *AzureLocalizedDateTime) GetLocalizePerMarketOk() (*bool, bool)`

GetLocalizePerMarketOk returns a tuple with the LocalizePerMarket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizePerMarket

`func (o *AzureLocalizedDateTime) SetLocalizePerMarket(v bool)`

SetLocalizePerMarket sets LocalizePerMarket field to given value.

### HasLocalizePerMarket

`func (o *AzureLocalizedDateTime) HasLocalizePerMarket() bool`

HasLocalizePerMarket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


