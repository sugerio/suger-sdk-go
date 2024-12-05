# BillingAddonRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**ChargeOn** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingAddonRecord

`func NewBillingAddonRecord() *BillingAddonRecord`

NewBillingAddonRecord instantiates a new BillingAddonRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddonRecordWithDefaults

`func NewBillingAddonRecordWithDefaults() *BillingAddonRecord`

NewBillingAddonRecordWithDefaults instantiates a new BillingAddonRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillingAddonRecord) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillingAddonRecord) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillingAddonRecord) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BillingAddonRecord) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetChargeOn

`func (o *BillingAddonRecord) GetChargeOn() time.Time`

GetChargeOn returns the ChargeOn field if non-nil, zero value otherwise.

### GetChargeOnOk

`func (o *BillingAddonRecord) GetChargeOnOk() (*time.Time, bool)`

GetChargeOnOk returns a tuple with the ChargeOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeOn

`func (o *BillingAddonRecord) SetChargeOn(v time.Time)`

SetChargeOn sets ChargeOn field to given value.

### HasChargeOn

`func (o *BillingAddonRecord) HasChargeOn() bool`

HasChargeOn returns a boolean if a field has been set.

### GetDescription

`func (o *BillingAddonRecord) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingAddonRecord) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingAddonRecord) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingAddonRecord) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BillingAddonRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingAddonRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingAddonRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingAddonRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingAddonRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingAddonRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingAddonRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingAddonRecord) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


