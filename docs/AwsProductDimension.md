# AwsProductDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** | The term length for the commit amount, such as 6 months, or 1 year. The length is used together with timeUnit. Length and TimeUnit are only used for commit dimension. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Rate** | Pointer to **float32** | Below three fields are only used for pass data when create or update product&#39;s public offer pricing. Rate is only used for update public offer, becasue rate will be set as 0.01 when create new product. | [optional] 
**TimeUnit** | Pointer to [**TimeUnit**](TimeUnit.md) |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsProductDimension

`func NewAwsProductDimension() *AwsProductDimension`

NewAwsProductDimension instantiates a new AwsProductDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProductDimensionWithDefaults

`func NewAwsProductDimensionWithDefaults() *AwsProductDimension`

NewAwsProductDimensionWithDefaults instantiates a new AwsProductDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AwsProductDimension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AwsProductDimension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AwsProductDimension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AwsProductDimension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *AwsProductDimension) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AwsProductDimension) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AwsProductDimension) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AwsProductDimension) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLength

`func (o *AwsProductDimension) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *AwsProductDimension) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *AwsProductDimension) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *AwsProductDimension) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetName

`func (o *AwsProductDimension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsProductDimension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsProductDimension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AwsProductDimension) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRate

`func (o *AwsProductDimension) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AwsProductDimension) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AwsProductDimension) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AwsProductDimension) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTimeUnit

`func (o *AwsProductDimension) GetTimeUnit() TimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *AwsProductDimension) GetTimeUnitOk() (*TimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *AwsProductDimension) SetTimeUnit(v TimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *AwsProductDimension) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetTypes

`func (o *AwsProductDimension) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *AwsProductDimension) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *AwsProductDimension) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *AwsProductDimension) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetUnit

`func (o *AwsProductDimension) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AwsProductDimension) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AwsProductDimension) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AwsProductDimension) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


