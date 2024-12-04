# IdentityBuyer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactIds** | Pointer to **[]string** |  | [optional] 
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**BuyerInfo**](BuyerInfo.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Partner** | Pointer to [**Partner**](Partner.md) |  | [optional] 

## Methods

### NewIdentityBuyer

`func NewIdentityBuyer() *IdentityBuyer`

NewIdentityBuyer instantiates a new IdentityBuyer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityBuyerWithDefaults

`func NewIdentityBuyerWithDefaults() *IdentityBuyer`

NewIdentityBuyerWithDefaults instantiates a new IdentityBuyer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactIds

`func (o *IdentityBuyer) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *IdentityBuyer) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *IdentityBuyer) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *IdentityBuyer) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetCreationTime

`func (o *IdentityBuyer) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *IdentityBuyer) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *IdentityBuyer) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *IdentityBuyer) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityBuyer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityBuyer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityBuyer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityBuyer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *IdentityBuyer) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *IdentityBuyer) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *IdentityBuyer) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *IdentityBuyer) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetId

`func (o *IdentityBuyer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityBuyer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityBuyer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityBuyer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *IdentityBuyer) GetInfo() BuyerInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *IdentityBuyer) GetInfoOk() (*BuyerInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *IdentityBuyer) SetInfo(v BuyerInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *IdentityBuyer) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *IdentityBuyer) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *IdentityBuyer) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *IdentityBuyer) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *IdentityBuyer) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *IdentityBuyer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityBuyer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityBuyer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IdentityBuyer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationID

`func (o *IdentityBuyer) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *IdentityBuyer) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *IdentityBuyer) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *IdentityBuyer) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetPartner

`func (o *IdentityBuyer) GetPartner() Partner`

GetPartner returns the Partner field if non-nil, zero value otherwise.

### GetPartnerOk

`func (o *IdentityBuyer) GetPartnerOk() (*Partner, bool)`

GetPartnerOk returns a tuple with the Partner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartner

`func (o *IdentityBuyer) SetPartner(v Partner)`

SetPartner sets Partner field to given value.

### HasPartner

`func (o *IdentityBuyer) HasPartner() bool`

HasPartner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


