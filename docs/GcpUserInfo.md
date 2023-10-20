# GcpUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to **[]string** | An array of strings representing the user&#39;s roles. Right now, it can be either: ** account_admin, which indicates that the user is a Billing Account Administrator of the billing account that purchased the product, or ** project_editor, which indicates that the user is a Project Editor, but not a Billing Administrator, of the project under that billing account. | [optional] 
**UserIdentity** | Pointer to **string** | The user&#39;s obfuscated GAIA ID, which can be used to initiate Open ID Connect. | [optional] 

## Methods

### NewGcpUserInfo

`func NewGcpUserInfo() *GcpUserInfo`

NewGcpUserInfo instantiates a new GcpUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpUserInfoWithDefaults

`func NewGcpUserInfoWithDefaults() *GcpUserInfo`

NewGcpUserInfoWithDefaults instantiates a new GcpUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GcpUserInfo) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GcpUserInfo) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GcpUserInfo) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GcpUserInfo) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUserIdentity

`func (o *GcpUserInfo) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *GcpUserInfo) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *GcpUserInfo) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *GcpUserInfo) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


