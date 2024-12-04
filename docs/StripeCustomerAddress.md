# StripeCustomerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City, district, suburb, town, or village. | [optional] 
**Country** | Pointer to **string** | Two-letter country code (ISO 3166-1 alpha-2) | [optional] 
**Line1** | Pointer to **string** | Address line 1 (e.g., street, PO Box, or company name). | [optional] 
**Line2** | Pointer to **string** | Address line 2 (e.g., apartment, suite, unit, or building). | [optional] 
**PostalCode** | Pointer to **string** | ZIP or postal code. | [optional] 
**State** | Pointer to **string** | State, county, province, or region. | [optional] 

## Methods

### NewStripeCustomerAddress

`func NewStripeCustomerAddress() *StripeCustomerAddress`

NewStripeCustomerAddress instantiates a new StripeCustomerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeCustomerAddressWithDefaults

`func NewStripeCustomerAddressWithDefaults() *StripeCustomerAddress`

NewStripeCustomerAddressWithDefaults instantiates a new StripeCustomerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *StripeCustomerAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *StripeCustomerAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *StripeCustomerAddress) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *StripeCustomerAddress) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *StripeCustomerAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StripeCustomerAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StripeCustomerAddress) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StripeCustomerAddress) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetLine1

`func (o *StripeCustomerAddress) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *StripeCustomerAddress) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *StripeCustomerAddress) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *StripeCustomerAddress) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### GetLine2

`func (o *StripeCustomerAddress) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *StripeCustomerAddress) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *StripeCustomerAddress) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *StripeCustomerAddress) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *StripeCustomerAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *StripeCustomerAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *StripeCustomerAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *StripeCustomerAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *StripeCustomerAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StripeCustomerAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StripeCustomerAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StripeCustomerAddress) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


