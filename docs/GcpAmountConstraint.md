# GcpAmountConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAmount** | Pointer to [**GcpAmountUnit**](GcpAmountUnit.md) |  | [optional] 
**MaxAmount** | Pointer to [**GcpAmountUnit**](GcpAmountUnit.md) |  | [optional] 
**MinAmount** | Pointer to [**GcpAmountUnit**](GcpAmountUnit.md) |  | [optional] 

## Methods

### NewGcpAmountConstraint

`func NewGcpAmountConstraint() *GcpAmountConstraint`

NewGcpAmountConstraint instantiates a new GcpAmountConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpAmountConstraintWithDefaults

`func NewGcpAmountConstraintWithDefaults() *GcpAmountConstraint`

NewGcpAmountConstraintWithDefaults instantiates a new GcpAmountConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAmount

`func (o *GcpAmountConstraint) GetDefaultAmount() GcpAmountUnit`

GetDefaultAmount returns the DefaultAmount field if non-nil, zero value otherwise.

### GetDefaultAmountOk

`func (o *GcpAmountConstraint) GetDefaultAmountOk() (*GcpAmountUnit, bool)`

GetDefaultAmountOk returns a tuple with the DefaultAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAmount

`func (o *GcpAmountConstraint) SetDefaultAmount(v GcpAmountUnit)`

SetDefaultAmount sets DefaultAmount field to given value.

### HasDefaultAmount

`func (o *GcpAmountConstraint) HasDefaultAmount() bool`

HasDefaultAmount returns a boolean if a field has been set.

### GetMaxAmount

`func (o *GcpAmountConstraint) GetMaxAmount() GcpAmountUnit`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *GcpAmountConstraint) GetMaxAmountOk() (*GcpAmountUnit, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *GcpAmountConstraint) SetMaxAmount(v GcpAmountUnit)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *GcpAmountConstraint) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetMinAmount

`func (o *GcpAmountConstraint) GetMinAmount() GcpAmountUnit`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *GcpAmountConstraint) GetMinAmountOk() (*GcpAmountUnit, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *GcpAmountConstraint) SetMinAmount(v GcpAmountUnit)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *GcpAmountConstraint) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


