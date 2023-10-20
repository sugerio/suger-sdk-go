# AddEntitlementCreditParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditAmountIncrement** | **float32** | The amount to be added to the credit amount. | 
**EntitlementID** | **string** |  | 
**OrganizationID** | **string** |  | 

## Methods

### NewAddEntitlementCreditParams

`func NewAddEntitlementCreditParams(creditAmountIncrement float32, entitlementID string, organizationID string, ) *AddEntitlementCreditParams`

NewAddEntitlementCreditParams instantiates a new AddEntitlementCreditParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEntitlementCreditParamsWithDefaults

`func NewAddEntitlementCreditParamsWithDefaults() *AddEntitlementCreditParams`

NewAddEntitlementCreditParamsWithDefaults instantiates a new AddEntitlementCreditParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditAmountIncrement

`func (o *AddEntitlementCreditParams) GetCreditAmountIncrement() float32`

GetCreditAmountIncrement returns the CreditAmountIncrement field if non-nil, zero value otherwise.

### GetCreditAmountIncrementOk

`func (o *AddEntitlementCreditParams) GetCreditAmountIncrementOk() (*float32, bool)`

GetCreditAmountIncrementOk returns a tuple with the CreditAmountIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditAmountIncrement

`func (o *AddEntitlementCreditParams) SetCreditAmountIncrement(v float32)`

SetCreditAmountIncrement sets CreditAmountIncrement field to given value.


### GetEntitlementID

`func (o *AddEntitlementCreditParams) GetEntitlementID() string`

GetEntitlementID returns the EntitlementID field if non-nil, zero value otherwise.

### GetEntitlementIDOk

`func (o *AddEntitlementCreditParams) GetEntitlementIDOk() (*string, bool)`

GetEntitlementIDOk returns a tuple with the EntitlementID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementID

`func (o *AddEntitlementCreditParams) SetEntitlementID(v string)`

SetEntitlementID sets EntitlementID field to given value.


### GetOrganizationID

`func (o *AddEntitlementCreditParams) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *AddEntitlementCreditParams) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *AddEntitlementCreditParams) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


