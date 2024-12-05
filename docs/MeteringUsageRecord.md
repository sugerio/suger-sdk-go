# MeteringUsageRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key is the unique identifier of a billable metric. | [optional] 
**Properties** | Pointer to **map[string]interface{}** | Properties is the filters of a billable metric. It should be equal to the filters of the billable metric. | [optional] 
**Quantity** | Pointer to **float32** | The quantity (or numeric value) of a billable metric. | [optional] 

## Methods

### NewMeteringUsageRecord

`func NewMeteringUsageRecord() *MeteringUsageRecord`

NewMeteringUsageRecord instantiates a new MeteringUsageRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeteringUsageRecordWithDefaults

`func NewMeteringUsageRecordWithDefaults() *MeteringUsageRecord`

NewMeteringUsageRecordWithDefaults instantiates a new MeteringUsageRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MeteringUsageRecord) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MeteringUsageRecord) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MeteringUsageRecord) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MeteringUsageRecord) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetProperties

`func (o *MeteringUsageRecord) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MeteringUsageRecord) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MeteringUsageRecord) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MeteringUsageRecord) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuantity

`func (o *MeteringUsageRecord) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *MeteringUsageRecord) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *MeteringUsageRecord) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *MeteringUsageRecord) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


