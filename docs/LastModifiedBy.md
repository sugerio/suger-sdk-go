# LastModifiedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email of the creator. | [optional] 
**EntityId** | Pointer to **string** | The ID of the creator. | [optional] 
**EntityType** | Pointer to [**EntityType**](EntityType.md) | The Entity type of the creator, either USER or API_CLIENT. | [optional] 
**Name** | Pointer to **string** | The name of the creator. | [optional] 

## Methods

### NewLastModifiedBy

`func NewLastModifiedBy() *LastModifiedBy`

NewLastModifiedBy instantiates a new LastModifiedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastModifiedByWithDefaults

`func NewLastModifiedByWithDefaults() *LastModifiedBy`

NewLastModifiedByWithDefaults instantiates a new LastModifiedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *LastModifiedBy) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *LastModifiedBy) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *LastModifiedBy) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *LastModifiedBy) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEntityId

`func (o *LastModifiedBy) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *LastModifiedBy) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *LastModifiedBy) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *LastModifiedBy) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *LastModifiedBy) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *LastModifiedBy) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *LastModifiedBy) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *LastModifiedBy) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetName

`func (o *LastModifiedBy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LastModifiedBy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LastModifiedBy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LastModifiedBy) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


