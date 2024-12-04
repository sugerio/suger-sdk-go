# BillingAddon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** | Description of the addon | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Info** | Pointer to [**BillingAddonInfo**](BillingAddonInfo.md) |  | [optional] 
**LastUpdateTime** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** | Name of the addon, e.g. \&quot;Support Plan\&quot; | [optional] 
**OrganizationID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BillingAddonStatus**](BillingAddonStatus.md) |  | [optional] 

## Methods

### NewBillingAddon

`func NewBillingAddon() *BillingAddon`

NewBillingAddon instantiates a new BillingAddon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddonWithDefaults

`func NewBillingAddonWithDefaults() *BillingAddon`

NewBillingAddonWithDefaults instantiates a new BillingAddon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *BillingAddon) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BillingAddon) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BillingAddon) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *BillingAddon) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDescription

`func (o *BillingAddon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillingAddon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillingAddon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillingAddon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *BillingAddon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingAddon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingAddon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingAddon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInfo

`func (o *BillingAddon) GetInfo() BillingAddonInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingAddon) GetInfoOk() (*BillingAddonInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingAddon) SetInfo(v BillingAddonInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingAddon) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *BillingAddon) GetLastUpdateTime() time.Time`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *BillingAddon) GetLastUpdateTimeOk() (*time.Time, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *BillingAddon) SetLastUpdateTime(v time.Time)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *BillingAddon) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *BillingAddon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingAddon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingAddon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingAddon) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationID

`func (o *BillingAddon) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *BillingAddon) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *BillingAddon) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *BillingAddon) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetStatus

`func (o *BillingAddon) GetStatus() BillingAddonStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingAddon) GetStatusOk() (*BillingAddonStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingAddon) SetStatus(v BillingAddonStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingAddon) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


