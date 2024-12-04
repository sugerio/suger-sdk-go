# AdyenBuyer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to **map[string]interface{}** | Settings store key-value pairs such as paymentMethodId,syncWithProvider,providerPaymentMethods. | [optional] 
**ShopperId** | Pointer to **string** | The shopperId on the adyen platform corresponding to the buyer. | [optional] 

## Methods

### NewAdyenBuyer

`func NewAdyenBuyer() *AdyenBuyer`

NewAdyenBuyer instantiates a new AdyenBuyer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenBuyerWithDefaults

`func NewAdyenBuyerWithDefaults() *AdyenBuyer`

NewAdyenBuyerWithDefaults instantiates a new AdyenBuyer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *AdyenBuyer) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AdyenBuyer) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AdyenBuyer) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AdyenBuyer) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetShopperId

`func (o *AdyenBuyer) GetShopperId() string`

GetShopperId returns the ShopperId field if non-nil, zero value otherwise.

### GetShopperIdOk

`func (o *AdyenBuyer) GetShopperIdOk() (*string, bool)`

GetShopperIdOk returns a tuple with the ShopperId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopperId

`func (o *AdyenBuyer) SetShopperId(v string)`

SetShopperId sets ShopperId field to given value.

### HasShopperId

`func (o *AdyenBuyer) HasShopperId() bool`

HasShopperId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


