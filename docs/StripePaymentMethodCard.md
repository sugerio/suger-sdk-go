# StripePaymentMethodCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** | Card brand. | [optional] 
**Country** | Pointer to **string** | Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you&#39;ve collected. | [optional] 
**DisplayBrand** | Pointer to **string** | The brand to use when displaying the card, this accounts for customer&#39;s brand choice on dual-branded cards. Can be &#x60;american_express&#x60;, &#x60;cartes_bancaires&#x60;, &#x60;diners_club&#x60;, &#x60;discover&#x60;, &#x60;eftpos_australia&#x60;, &#x60;interac&#x60;, &#x60;jcb&#x60;, &#x60;mastercard&#x60;, &#x60;union_pay&#x60;, &#x60;visa&#x60;, or &#x60;other&#x60; and may contain more values in the future. | [optional] 
**ExpMonth** | Pointer to **int32** | Two-digit number representing the card&#39;s expiration month. | [optional] 
**ExpYear** | Pointer to **int32** | Four-digit number representing the card&#39;s expiration year. | [optional] 
**Fingerprint** | Pointer to **string** | Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number. | [optional] 
**Funding** | Pointer to **string** | Card funding type. Can be &#x60;credit&#x60;, &#x60;debit&#x60;, &#x60;prepaid&#x60;, or &#x60;unknown&#x60;. | [optional] 
**Last4** | Pointer to **string** | The last four digits of the card. | [optional] 

## Methods

### NewStripePaymentMethodCard

`func NewStripePaymentMethodCard() *StripePaymentMethodCard`

NewStripePaymentMethodCard instantiates a new StripePaymentMethodCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripePaymentMethodCardWithDefaults

`func NewStripePaymentMethodCardWithDefaults() *StripePaymentMethodCard`

NewStripePaymentMethodCardWithDefaults instantiates a new StripePaymentMethodCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *StripePaymentMethodCard) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *StripePaymentMethodCard) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *StripePaymentMethodCard) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *StripePaymentMethodCard) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCountry

`func (o *StripePaymentMethodCard) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StripePaymentMethodCard) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StripePaymentMethodCard) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *StripePaymentMethodCard) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDisplayBrand

`func (o *StripePaymentMethodCard) GetDisplayBrand() string`

GetDisplayBrand returns the DisplayBrand field if non-nil, zero value otherwise.

### GetDisplayBrandOk

`func (o *StripePaymentMethodCard) GetDisplayBrandOk() (*string, bool)`

GetDisplayBrandOk returns a tuple with the DisplayBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayBrand

`func (o *StripePaymentMethodCard) SetDisplayBrand(v string)`

SetDisplayBrand sets DisplayBrand field to given value.

### HasDisplayBrand

`func (o *StripePaymentMethodCard) HasDisplayBrand() bool`

HasDisplayBrand returns a boolean if a field has been set.

### GetExpMonth

`func (o *StripePaymentMethodCard) GetExpMonth() int32`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *StripePaymentMethodCard) GetExpMonthOk() (*int32, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *StripePaymentMethodCard) SetExpMonth(v int32)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *StripePaymentMethodCard) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *StripePaymentMethodCard) GetExpYear() int32`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *StripePaymentMethodCard) GetExpYearOk() (*int32, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *StripePaymentMethodCard) SetExpYear(v int32)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *StripePaymentMethodCard) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetFingerprint

`func (o *StripePaymentMethodCard) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *StripePaymentMethodCard) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *StripePaymentMethodCard) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *StripePaymentMethodCard) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFunding

`func (o *StripePaymentMethodCard) GetFunding() string`

GetFunding returns the Funding field if non-nil, zero value otherwise.

### GetFundingOk

`func (o *StripePaymentMethodCard) GetFundingOk() (*string, bool)`

GetFundingOk returns a tuple with the Funding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunding

`func (o *StripePaymentMethodCard) SetFunding(v string)`

SetFunding sets Funding field to given value.

### HasFunding

`func (o *StripePaymentMethodCard) HasFunding() bool`

HasFunding returns a boolean if a field has been set.

### GetLast4

`func (o *StripePaymentMethodCard) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *StripePaymentMethodCard) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *StripePaymentMethodCard) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *StripePaymentMethodCard) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


