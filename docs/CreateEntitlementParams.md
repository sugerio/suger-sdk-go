# CreateEntitlementParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerId** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** | The end date of the entitlement. If it is null, the entitlement will ends based on the offer. If it is in the past, the entitlement will be created as CANCELLED status. | [optional] 
**OfferId** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** | The start date of the entitlement. If it is null, the entitlement starts immediately. It can be in the future or in the past. | [optional] 

## Methods

### NewCreateEntitlementParams

`func NewCreateEntitlementParams() *CreateEntitlementParams`

NewCreateEntitlementParams instantiates a new CreateEntitlementParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEntitlementParamsWithDefaults

`func NewCreateEntitlementParamsWithDefaults() *CreateEntitlementParams`

NewCreateEntitlementParamsWithDefaults instantiates a new CreateEntitlementParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerId

`func (o *CreateEntitlementParams) GetBuyerId() string`

GetBuyerId returns the BuyerId field if non-nil, zero value otherwise.

### GetBuyerIdOk

`func (o *CreateEntitlementParams) GetBuyerIdOk() (*string, bool)`

GetBuyerIdOk returns a tuple with the BuyerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerId

`func (o *CreateEntitlementParams) SetBuyerId(v string)`

SetBuyerId sets BuyerId field to given value.

### HasBuyerId

`func (o *CreateEntitlementParams) HasBuyerId() bool`

HasBuyerId returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateEntitlementParams) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateEntitlementParams) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateEntitlementParams) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateEntitlementParams) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOfferId

`func (o *CreateEntitlementParams) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *CreateEntitlementParams) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *CreateEntitlementParams) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *CreateEntitlementParams) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateEntitlementParams) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateEntitlementParams) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateEntitlementParams) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateEntitlementParams) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateEntitlementParams) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateEntitlementParams) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateEntitlementParams) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateEntitlementParams) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


