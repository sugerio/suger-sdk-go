# GcpPriceValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | such as \&quot;USD\&quot; | [optional] 
**Nanos** | Pointer to **int32** | for the decimal part, such as 30000000 &#x3D; $0.03 | [optional] 
**Units** | Pointer to **string** | for the integer part, such as \&quot;500000\&quot; &#x3D; $50K | [optional] 

## Methods

### NewGcpPriceValue

`func NewGcpPriceValue() *GcpPriceValue`

NewGcpPriceValue instantiates a new GcpPriceValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpPriceValueWithDefaults

`func NewGcpPriceValueWithDefaults() *GcpPriceValue`

NewGcpPriceValueWithDefaults instantiates a new GcpPriceValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GcpPriceValue) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GcpPriceValue) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GcpPriceValue) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GcpPriceValue) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetNanos

`func (o *GcpPriceValue) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *GcpPriceValue) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *GcpPriceValue) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *GcpPriceValue) HasNanos() bool`

HasNanos returns a boolean if a field has been set.

### GetUnits

`func (o *GcpPriceValue) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GcpPriceValue) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GcpPriceValue) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GcpPriceValue) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


