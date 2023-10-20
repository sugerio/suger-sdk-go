# CommitDimension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsUserCreated** | Pointer to **bool** | Whether this commit dimension is newly created by user, when creating AWS Marketplace Contract private offer. | [optional] 
**Key** | Pointer to **string** | API name of the dimension | [optional] 
**Length** | Pointer to **int32** | The term length for the commit amount, such as 6 months, or 1 year. The length is used together with timeUnit. If the length is zero, use the TermEndTime. | [optional] 
**MaximumUsers** | Pointer to **int32** | The maximum number of users for PER_USER commit | [optional] 
**MinimumUsers** | Pointer to **int32** | The minimum number of users for PER_USER commit | [optional] 
**Name** | Pointer to **string** | Display name of the dimension | [optional] 
**Quantity** | Pointer to **int32** | The quantity of this commit. | [optional] 
**Rate** | Pointer to **float32** | The commit amount. For GCP, it is monthly commitment. | [optional] 
**Term** | Pointer to **string** | The term of the commit amount. It is used for front-end display only. | [optional] 
**TermEndTime** | Pointer to **string** | The end time of the commit term. | [optional] 
**TimeUnit** | Pointer to [**CommitDimensionTimeUnit**](CommitDimensionTimeUnit.md) |  | [optional] 
**Type** | Pointer to [**CommitDimensionType**](CommitDimensionType.md) |  | [optional] 
**Types** | Pointer to **[]string** | These indicate whether the dimension covers metering, entitlement, or support for external metering | [optional] 

## Methods

### NewCommitDimension

`func NewCommitDimension() *CommitDimension`

NewCommitDimension instantiates a new CommitDimension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDimensionWithDefaults

`func NewCommitDimensionWithDefaults() *CommitDimension`

NewCommitDimensionWithDefaults instantiates a new CommitDimension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CommitDimension) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CommitDimension) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CommitDimension) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CommitDimension) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *CommitDimension) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommitDimension) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommitDimension) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommitDimension) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsUserCreated

`func (o *CommitDimension) GetIsUserCreated() bool`

GetIsUserCreated returns the IsUserCreated field if non-nil, zero value otherwise.

### GetIsUserCreatedOk

`func (o *CommitDimension) GetIsUserCreatedOk() (*bool, bool)`

GetIsUserCreatedOk returns a tuple with the IsUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserCreated

`func (o *CommitDimension) SetIsUserCreated(v bool)`

SetIsUserCreated sets IsUserCreated field to given value.

### HasIsUserCreated

`func (o *CommitDimension) HasIsUserCreated() bool`

HasIsUserCreated returns a boolean if a field has been set.

### GetKey

`func (o *CommitDimension) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CommitDimension) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CommitDimension) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CommitDimension) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLength

`func (o *CommitDimension) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *CommitDimension) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *CommitDimension) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *CommitDimension) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMaximumUsers

`func (o *CommitDimension) GetMaximumUsers() int32`

GetMaximumUsers returns the MaximumUsers field if non-nil, zero value otherwise.

### GetMaximumUsersOk

`func (o *CommitDimension) GetMaximumUsersOk() (*int32, bool)`

GetMaximumUsersOk returns a tuple with the MaximumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumUsers

`func (o *CommitDimension) SetMaximumUsers(v int32)`

SetMaximumUsers sets MaximumUsers field to given value.

### HasMaximumUsers

`func (o *CommitDimension) HasMaximumUsers() bool`

HasMaximumUsers returns a boolean if a field has been set.

### GetMinimumUsers

`func (o *CommitDimension) GetMinimumUsers() int32`

GetMinimumUsers returns the MinimumUsers field if non-nil, zero value otherwise.

### GetMinimumUsersOk

`func (o *CommitDimension) GetMinimumUsersOk() (*int32, bool)`

GetMinimumUsersOk returns a tuple with the MinimumUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUsers

`func (o *CommitDimension) SetMinimumUsers(v int32)`

SetMinimumUsers sets MinimumUsers field to given value.

### HasMinimumUsers

`func (o *CommitDimension) HasMinimumUsers() bool`

HasMinimumUsers returns a boolean if a field has been set.

### GetName

`func (o *CommitDimension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommitDimension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommitDimension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommitDimension) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuantity

`func (o *CommitDimension) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CommitDimension) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CommitDimension) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CommitDimension) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRate

`func (o *CommitDimension) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *CommitDimension) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *CommitDimension) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *CommitDimension) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetTerm

`func (o *CommitDimension) GetTerm() string`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *CommitDimension) GetTermOk() (*string, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *CommitDimension) SetTerm(v string)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *CommitDimension) HasTerm() bool`

HasTerm returns a boolean if a field has been set.

### GetTermEndTime

`func (o *CommitDimension) GetTermEndTime() string`

GetTermEndTime returns the TermEndTime field if non-nil, zero value otherwise.

### GetTermEndTimeOk

`func (o *CommitDimension) GetTermEndTimeOk() (*string, bool)`

GetTermEndTimeOk returns a tuple with the TermEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermEndTime

`func (o *CommitDimension) SetTermEndTime(v string)`

SetTermEndTime sets TermEndTime field to given value.

### HasTermEndTime

`func (o *CommitDimension) HasTermEndTime() bool`

HasTermEndTime returns a boolean if a field has been set.

### GetTimeUnit

`func (o *CommitDimension) GetTimeUnit() CommitDimensionTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *CommitDimension) GetTimeUnitOk() (*CommitDimensionTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *CommitDimension) SetTimeUnit(v CommitDimensionTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *CommitDimension) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetType

`func (o *CommitDimension) GetType() CommitDimensionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommitDimension) GetTypeOk() (*CommitDimensionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommitDimension) SetType(v CommitDimensionType)`

SetType sets Type field to given value.

### HasType

`func (o *CommitDimension) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypes

`func (o *CommitDimension) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *CommitDimension) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *CommitDimension) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *CommitDimension) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


