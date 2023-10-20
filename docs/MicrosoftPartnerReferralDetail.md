# MicrosoftPartnerReferralDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClosingDateTime** | Pointer to **string** | in UTC date time format | [optional] 
**Currency** | Pointer to **string** | ISO 4217 currency symbol | [optional] 
**CustomerAction** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomerRequestedContact** | Pointer to **map[string]interface{}** |  | [optional] 
**DealValue** | Pointer to **float32** |  | [optional] 
**IncentiveLevel** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Requirements** | Pointer to [**MicrosoftPartnerReferralRequirements**](MicrosoftPartnerReferralRequirements.md) |  | [optional] 

## Methods

### NewMicrosoftPartnerReferralDetail

`func NewMicrosoftPartnerReferralDetail() *MicrosoftPartnerReferralDetail`

NewMicrosoftPartnerReferralDetail instantiates a new MicrosoftPartnerReferralDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftPartnerReferralDetailWithDefaults

`func NewMicrosoftPartnerReferralDetailWithDefaults() *MicrosoftPartnerReferralDetail`

NewMicrosoftPartnerReferralDetailWithDefaults instantiates a new MicrosoftPartnerReferralDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClosingDateTime

`func (o *MicrosoftPartnerReferralDetail) GetClosingDateTime() string`

GetClosingDateTime returns the ClosingDateTime field if non-nil, zero value otherwise.

### GetClosingDateTimeOk

`func (o *MicrosoftPartnerReferralDetail) GetClosingDateTimeOk() (*string, bool)`

GetClosingDateTimeOk returns a tuple with the ClosingDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingDateTime

`func (o *MicrosoftPartnerReferralDetail) SetClosingDateTime(v string)`

SetClosingDateTime sets ClosingDateTime field to given value.

### HasClosingDateTime

`func (o *MicrosoftPartnerReferralDetail) HasClosingDateTime() bool`

HasClosingDateTime returns a boolean if a field has been set.

### GetCurrency

`func (o *MicrosoftPartnerReferralDetail) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MicrosoftPartnerReferralDetail) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MicrosoftPartnerReferralDetail) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MicrosoftPartnerReferralDetail) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCustomerAction

`func (o *MicrosoftPartnerReferralDetail) GetCustomerAction() map[string]interface{}`

GetCustomerAction returns the CustomerAction field if non-nil, zero value otherwise.

### GetCustomerActionOk

`func (o *MicrosoftPartnerReferralDetail) GetCustomerActionOk() (*map[string]interface{}, bool)`

GetCustomerActionOk returns a tuple with the CustomerAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAction

`func (o *MicrosoftPartnerReferralDetail) SetCustomerAction(v map[string]interface{})`

SetCustomerAction sets CustomerAction field to given value.

### HasCustomerAction

`func (o *MicrosoftPartnerReferralDetail) HasCustomerAction() bool`

HasCustomerAction returns a boolean if a field has been set.

### GetCustomerRequestedContact

`func (o *MicrosoftPartnerReferralDetail) GetCustomerRequestedContact() map[string]interface{}`

GetCustomerRequestedContact returns the CustomerRequestedContact field if non-nil, zero value otherwise.

### GetCustomerRequestedContactOk

`func (o *MicrosoftPartnerReferralDetail) GetCustomerRequestedContactOk() (*map[string]interface{}, bool)`

GetCustomerRequestedContactOk returns a tuple with the CustomerRequestedContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRequestedContact

`func (o *MicrosoftPartnerReferralDetail) SetCustomerRequestedContact(v map[string]interface{})`

SetCustomerRequestedContact sets CustomerRequestedContact field to given value.

### HasCustomerRequestedContact

`func (o *MicrosoftPartnerReferralDetail) HasCustomerRequestedContact() bool`

HasCustomerRequestedContact returns a boolean if a field has been set.

### GetDealValue

`func (o *MicrosoftPartnerReferralDetail) GetDealValue() float32`

GetDealValue returns the DealValue field if non-nil, zero value otherwise.

### GetDealValueOk

`func (o *MicrosoftPartnerReferralDetail) GetDealValueOk() (*float32, bool)`

GetDealValueOk returns a tuple with the DealValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealValue

`func (o *MicrosoftPartnerReferralDetail) SetDealValue(v float32)`

SetDealValue sets DealValue field to given value.

### HasDealValue

`func (o *MicrosoftPartnerReferralDetail) HasDealValue() bool`

HasDealValue returns a boolean if a field has been set.

### GetIncentiveLevel

`func (o *MicrosoftPartnerReferralDetail) GetIncentiveLevel() map[string]interface{}`

GetIncentiveLevel returns the IncentiveLevel field if non-nil, zero value otherwise.

### GetIncentiveLevelOk

`func (o *MicrosoftPartnerReferralDetail) GetIncentiveLevelOk() (*map[string]interface{}, bool)`

GetIncentiveLevelOk returns a tuple with the IncentiveLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncentiveLevel

`func (o *MicrosoftPartnerReferralDetail) SetIncentiveLevel(v map[string]interface{})`

SetIncentiveLevel sets IncentiveLevel field to given value.

### HasIncentiveLevel

`func (o *MicrosoftPartnerReferralDetail) HasIncentiveLevel() bool`

HasIncentiveLevel returns a boolean if a field has been set.

### GetNotes

`func (o *MicrosoftPartnerReferralDetail) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *MicrosoftPartnerReferralDetail) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *MicrosoftPartnerReferralDetail) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *MicrosoftPartnerReferralDetail) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRequirements

`func (o *MicrosoftPartnerReferralDetail) GetRequirements() MicrosoftPartnerReferralRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *MicrosoftPartnerReferralDetail) GetRequirementsOk() (*MicrosoftPartnerReferralRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *MicrosoftPartnerReferralDetail) SetRequirements(v MicrosoftPartnerReferralRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *MicrosoftPartnerReferralDetail) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


