# OrbItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**ExternalConnections** | Pointer to [**[]OrbExternalConnection**](OrbExternalConnection.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewOrbItem

`func NewOrbItem() *OrbItem`

NewOrbItem instantiates a new OrbItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrbItemWithDefaults

`func NewOrbItemWithDefaults() *OrbItem`

NewOrbItemWithDefaults instantiates a new OrbItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *OrbItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrbItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrbItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrbItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalConnections

`func (o *OrbItem) GetExternalConnections() []OrbExternalConnection`

GetExternalConnections returns the ExternalConnections field if non-nil, zero value otherwise.

### GetExternalConnectionsOk

`func (o *OrbItem) GetExternalConnectionsOk() (*[]OrbExternalConnection, bool)`

GetExternalConnectionsOk returns a tuple with the ExternalConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalConnections

`func (o *OrbItem) SetExternalConnections(v []OrbExternalConnection)`

SetExternalConnections sets ExternalConnections field to given value.

### HasExternalConnections

`func (o *OrbItem) HasExternalConnections() bool`

HasExternalConnections returns a boolean if a field has been set.

### GetId

`func (o *OrbItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrbItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrbItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrbItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OrbItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrbItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrbItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrbItem) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


