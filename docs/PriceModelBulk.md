# PriceModelBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkAmount** | Pointer to **float32** | A currency amount to rate usage by | [optional] 
**BulkSize** | Pointer to **int32** | An integer amount to represent package size. For example, 1000 here would divide usage by 1000 before multiplying by package_amount in rating | [optional] 

## Methods

### NewPriceModelBulk

`func NewPriceModelBulk() *PriceModelBulk`

NewPriceModelBulk instantiates a new PriceModelBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceModelBulkWithDefaults

`func NewPriceModelBulkWithDefaults() *PriceModelBulk`

NewPriceModelBulkWithDefaults instantiates a new PriceModelBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkAmount

`func (o *PriceModelBulk) GetBulkAmount() float32`

GetBulkAmount returns the BulkAmount field if non-nil, zero value otherwise.

### GetBulkAmountOk

`func (o *PriceModelBulk) GetBulkAmountOk() (*float32, bool)`

GetBulkAmountOk returns a tuple with the BulkAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkAmount

`func (o *PriceModelBulk) SetBulkAmount(v float32)`

SetBulkAmount sets BulkAmount field to given value.

### HasBulkAmount

`func (o *PriceModelBulk) HasBulkAmount() bool`

HasBulkAmount returns a boolean if a field has been set.

### GetBulkSize

`func (o *PriceModelBulk) GetBulkSize() int32`

GetBulkSize returns the BulkSize field if non-nil, zero value otherwise.

### GetBulkSizeOk

`func (o *PriceModelBulk) GetBulkSizeOk() (*int32, bool)`

GetBulkSizeOk returns a tuple with the BulkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkSize

`func (o *PriceModelBulk) SetBulkSize(v int32)`

SetBulkSize sets BulkSize field to given value.

### HasBulkSize

`func (o *PriceModelBulk) HasBulkSize() bool`

HasBulkSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


