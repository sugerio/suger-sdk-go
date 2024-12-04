# GcpMarketplaceMeteringMoney

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | CurrencyCode: The three-letter currency code defined in ISO 4217. | [optional] 
**Nanos** | Pointer to **int32** | Nanos: Number of nano (10^-9) units of the amount. The value must be between -999,999,999 and +999,999,999 inclusive. If &#x60;units&#x60; is positive, &#x60;nanos&#x60; must be positive or zero. If &#x60;units&#x60; is zero, &#x60;nanos&#x60; can be positive, zero, or negative. If &#x60;units&#x60; is negative, &#x60;nanos&#x60; must be negative or zero. For example $-1.75 is represented as &#x60;units&#x60;&#x3D;-1 and &#x60;nanos&#x60;&#x3D;-750,000,000. | [optional] 
**Units** | Pointer to **string** | Units: The whole units of the amount. For example if &#x60;currencyCode&#x60; is \&quot;USD\&quot;, then 1 unit is one US dollar. | [optional] 

## Methods

### NewGcpMarketplaceMeteringMoney

`func NewGcpMarketplaceMeteringMoney() *GcpMarketplaceMeteringMoney`

NewGcpMarketplaceMeteringMoney instantiates a new GcpMarketplaceMeteringMoney object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpMarketplaceMeteringMoneyWithDefaults

`func NewGcpMarketplaceMeteringMoneyWithDefaults() *GcpMarketplaceMeteringMoney`

NewGcpMarketplaceMeteringMoneyWithDefaults instantiates a new GcpMarketplaceMeteringMoney object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GcpMarketplaceMeteringMoney) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GcpMarketplaceMeteringMoney) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GcpMarketplaceMeteringMoney) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GcpMarketplaceMeteringMoney) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetNanos

`func (o *GcpMarketplaceMeteringMoney) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *GcpMarketplaceMeteringMoney) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *GcpMarketplaceMeteringMoney) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *GcpMarketplaceMeteringMoney) HasNanos() bool`

HasNanos returns a boolean if a field has been set.

### GetUnits

`func (o *GcpMarketplaceMeteringMoney) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GcpMarketplaceMeteringMoney) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GcpMarketplaceMeteringMoney) SetUnits(v string)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GcpMarketplaceMeteringMoney) HasUnits() bool`

HasUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


