# UsageMeteringDailyRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Key** | Pointer to **string** | The dimension key of the usage metering. | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 

## Methods

### NewUsageMeteringDailyRecord

`func NewUsageMeteringDailyRecord() *UsageMeteringDailyRecord`

NewUsageMeteringDailyRecord instantiates a new UsageMeteringDailyRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMeteringDailyRecordWithDefaults

`func NewUsageMeteringDailyRecordWithDefaults() *UsageMeteringDailyRecord`

NewUsageMeteringDailyRecordWithDefaults instantiates a new UsageMeteringDailyRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *UsageMeteringDailyRecord) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UsageMeteringDailyRecord) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UsageMeteringDailyRecord) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UsageMeteringDailyRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDate

`func (o *UsageMeteringDailyRecord) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsageMeteringDailyRecord) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsageMeteringDailyRecord) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *UsageMeteringDailyRecord) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetKey

`func (o *UsageMeteringDailyRecord) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UsageMeteringDailyRecord) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UsageMeteringDailyRecord) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UsageMeteringDailyRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPartner

`func (o *UsageMeteringDailyRecord) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *UsageMeteringDailyRecord) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *UsageMeteringDailyRecord) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *UsageMeteringDailyRecord) HasPartner() bool`

HasPartner returns a boolean if a field has been set.

### GetQuantity

`func (o *UsageMeteringDailyRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageMeteringDailyRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageMeteringDailyRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UsageMeteringDailyRecord) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


