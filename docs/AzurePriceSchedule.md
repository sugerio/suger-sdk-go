# AzurePriceSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceCadence** | Pointer to [**AzurePriceCadence**](AzurePriceCadence.md) |  | [optional] 
**PricingModel** | Pointer to **string** |  | [optional] 
**PricingUnits** | Pointer to [**[]AzurePricingUnit**](AzurePricingUnit.md) |  | [optional] 
**RetailPrice** | Pointer to [**AzurePrice**](AzurePrice.md) |  | [optional] 

## Methods

### NewAzurePriceSchedule

`func NewAzurePriceSchedule() *AzurePriceSchedule`

NewAzurePriceSchedule instantiates a new AzurePriceSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzurePriceScheduleWithDefaults

`func NewAzurePriceScheduleWithDefaults() *AzurePriceSchedule`

NewAzurePriceScheduleWithDefaults instantiates a new AzurePriceSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceCadence

`func (o *AzurePriceSchedule) GetPriceCadence() AzurePriceCadence`

GetPriceCadence returns the PriceCadence field if non-nil, zero value otherwise.

### GetPriceCadenceOk

`func (o *AzurePriceSchedule) GetPriceCadenceOk() (*AzurePriceCadence, bool)`

GetPriceCadenceOk returns a tuple with the PriceCadence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCadence

`func (o *AzurePriceSchedule) SetPriceCadence(v AzurePriceCadence)`

SetPriceCadence sets PriceCadence field to given value.

### HasPriceCadence

`func (o *AzurePriceSchedule) HasPriceCadence() bool`

HasPriceCadence returns a boolean if a field has been set.

### GetPricingModel

`func (o *AzurePriceSchedule) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *AzurePriceSchedule) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *AzurePriceSchedule) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.

### HasPricingModel

`func (o *AzurePriceSchedule) HasPricingModel() bool`

HasPricingModel returns a boolean if a field has been set.

### GetPricingUnits

`func (o *AzurePriceSchedule) GetPricingUnits() []AzurePricingUnit`

GetPricingUnits returns the PricingUnits field if non-nil, zero value otherwise.

### GetPricingUnitsOk

`func (o *AzurePriceSchedule) GetPricingUnitsOk() (*[]AzurePricingUnit, bool)`

GetPricingUnitsOk returns a tuple with the PricingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingUnits

`func (o *AzurePriceSchedule) SetPricingUnits(v []AzurePricingUnit)`

SetPricingUnits sets PricingUnits field to given value.

### HasPricingUnits

`func (o *AzurePriceSchedule) HasPricingUnits() bool`

HasPricingUnits returns a boolean if a field has been set.

### GetRetailPrice

`func (o *AzurePriceSchedule) GetRetailPrice() AzurePrice`

GetRetailPrice returns the RetailPrice field if non-nil, zero value otherwise.

### GetRetailPriceOk

`func (o *AzurePriceSchedule) GetRetailPriceOk() (*AzurePrice, bool)`

GetRetailPriceOk returns a tuple with the RetailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPrice

`func (o *AzurePriceSchedule) SetRetailPrice(v AzurePrice)`

SetRetailPrice sets RetailPrice field to given value.

### HasRetailPrice

`func (o *AzurePriceSchedule) HasRetailPrice() bool`

HasRetailPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


