# AzureMarketplacePriceAndAvailabilityPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorePricing** | Pointer to [**AzureMarketplacePriceAndAvailabilityCorePrice**](AzureMarketplacePriceAndAvailabilityCorePrice.md) |  | [optional] 
**CustomMeters** | Pointer to [**AzureMarketplacePriceAndAvailabilityCustomMeterPrice**](AzureMarketplacePriceAndAvailabilityCustomMeterPrice.md) |  | [optional] 
**LicenseModel** | Pointer to **string** |  | [optional] 
**RecurrentPrice** | Pointer to [**AzureMarketplacePriceAndAvailabilityRecurrentPrice**](AzureMarketplacePriceAndAvailabilityRecurrentPrice.md) |  | [optional] 
**SystemMeterPricing** | Pointer to [**AzureMarketplacePriceAndAvailabilitySystemMeterPrice**](AzureMarketplacePriceAndAvailabilitySystemMeterPrice.md) |  | [optional] 

## Methods

### NewAzureMarketplacePriceAndAvailabilityPrice

`func NewAzureMarketplacePriceAndAvailabilityPrice() *AzureMarketplacePriceAndAvailabilityPrice`

NewAzureMarketplacePriceAndAvailabilityPrice instantiates a new AzureMarketplacePriceAndAvailabilityPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMarketplacePriceAndAvailabilityPriceWithDefaults

`func NewAzureMarketplacePriceAndAvailabilityPriceWithDefaults() *AzureMarketplacePriceAndAvailabilityPrice`

NewAzureMarketplacePriceAndAvailabilityPriceWithDefaults instantiates a new AzureMarketplacePriceAndAvailabilityPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorePricing

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCorePricing() AzureMarketplacePriceAndAvailabilityCorePrice`

GetCorePricing returns the CorePricing field if non-nil, zero value otherwise.

### GetCorePricingOk

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCorePricingOk() (*AzureMarketplacePriceAndAvailabilityCorePrice, bool)`

GetCorePricingOk returns a tuple with the CorePricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorePricing

`func (o *AzureMarketplacePriceAndAvailabilityPrice) SetCorePricing(v AzureMarketplacePriceAndAvailabilityCorePrice)`

SetCorePricing sets CorePricing field to given value.

### HasCorePricing

`func (o *AzureMarketplacePriceAndAvailabilityPrice) HasCorePricing() bool`

HasCorePricing returns a boolean if a field has been set.

### GetCustomMeters

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCustomMeters() AzureMarketplacePriceAndAvailabilityCustomMeterPrice`

GetCustomMeters returns the CustomMeters field if non-nil, zero value otherwise.

### GetCustomMetersOk

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetCustomMetersOk() (*AzureMarketplacePriceAndAvailabilityCustomMeterPrice, bool)`

GetCustomMetersOk returns a tuple with the CustomMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMeters

`func (o *AzureMarketplacePriceAndAvailabilityPrice) SetCustomMeters(v AzureMarketplacePriceAndAvailabilityCustomMeterPrice)`

SetCustomMeters sets CustomMeters field to given value.

### HasCustomMeters

`func (o *AzureMarketplacePriceAndAvailabilityPrice) HasCustomMeters() bool`

HasCustomMeters returns a boolean if a field has been set.

### GetLicenseModel

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetLicenseModel() string`

GetLicenseModel returns the LicenseModel field if non-nil, zero value otherwise.

### GetLicenseModelOk

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetLicenseModelOk() (*string, bool)`

GetLicenseModelOk returns a tuple with the LicenseModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseModel

`func (o *AzureMarketplacePriceAndAvailabilityPrice) SetLicenseModel(v string)`

SetLicenseModel sets LicenseModel field to given value.

### HasLicenseModel

`func (o *AzureMarketplacePriceAndAvailabilityPrice) HasLicenseModel() bool`

HasLicenseModel returns a boolean if a field has been set.

### GetRecurrentPrice

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetRecurrentPrice() AzureMarketplacePriceAndAvailabilityRecurrentPrice`

GetRecurrentPrice returns the RecurrentPrice field if non-nil, zero value otherwise.

### GetRecurrentPriceOk

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetRecurrentPriceOk() (*AzureMarketplacePriceAndAvailabilityRecurrentPrice, bool)`

GetRecurrentPriceOk returns a tuple with the RecurrentPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrentPrice

`func (o *AzureMarketplacePriceAndAvailabilityPrice) SetRecurrentPrice(v AzureMarketplacePriceAndAvailabilityRecurrentPrice)`

SetRecurrentPrice sets RecurrentPrice field to given value.

### HasRecurrentPrice

`func (o *AzureMarketplacePriceAndAvailabilityPrice) HasRecurrentPrice() bool`

HasRecurrentPrice returns a boolean if a field has been set.

### GetSystemMeterPricing

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetSystemMeterPricing() AzureMarketplacePriceAndAvailabilitySystemMeterPrice`

GetSystemMeterPricing returns the SystemMeterPricing field if non-nil, zero value otherwise.

### GetSystemMeterPricingOk

`func (o *AzureMarketplacePriceAndAvailabilityPrice) GetSystemMeterPricingOk() (*AzureMarketplacePriceAndAvailabilitySystemMeterPrice, bool)`

GetSystemMeterPricingOk returns a tuple with the SystemMeterPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemMeterPricing

`func (o *AzureMarketplacePriceAndAvailabilityPrice) SetSystemMeterPricing(v AzureMarketplacePriceAndAvailabilitySystemMeterPrice)`

SetSystemMeterPricing sets SystemMeterPricing field to given value.

### HasSystemMeterPricing

`func (o *AzureMarketplacePriceAndAvailabilityPrice) HasSystemMeterPricing() bool`

HasSystemMeterPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


