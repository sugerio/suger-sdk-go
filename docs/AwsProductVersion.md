# AwsProductVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**DeliveryOptions** | Pointer to [**[]AwsProductDeliveryOption**](AwsProductDeliveryOption.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReleaseNotes** | Pointer to **string** |  | [optional] 
**VersionTitle** | Pointer to **string** |  | [optional] 

## Methods

### NewAwsProductVersion

`func NewAwsProductVersion() *AwsProductVersion`

NewAwsProductVersion instantiates a new AwsProductVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsProductVersionWithDefaults

`func NewAwsProductVersionWithDefaults() *AwsProductVersion`

NewAwsProductVersionWithDefaults instantiates a new AwsProductVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *AwsProductVersion) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *AwsProductVersion) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *AwsProductVersion) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *AwsProductVersion) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetDeliveryOptions

`func (o *AwsProductVersion) GetDeliveryOptions() []AwsProductDeliveryOption`

GetDeliveryOptions returns the DeliveryOptions field if non-nil, zero value otherwise.

### GetDeliveryOptionsOk

`func (o *AwsProductVersion) GetDeliveryOptionsOk() (*[]AwsProductDeliveryOption, bool)`

GetDeliveryOptionsOk returns a tuple with the DeliveryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOptions

`func (o *AwsProductVersion) SetDeliveryOptions(v []AwsProductDeliveryOption)`

SetDeliveryOptions sets DeliveryOptions field to given value.

### HasDeliveryOptions

`func (o *AwsProductVersion) HasDeliveryOptions() bool`

HasDeliveryOptions returns a boolean if a field has been set.

### GetId

`func (o *AwsProductVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwsProductVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwsProductVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwsProductVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReleaseNotes

`func (o *AwsProductVersion) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *AwsProductVersion) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *AwsProductVersion) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *AwsProductVersion) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### GetVersionTitle

`func (o *AwsProductVersion) GetVersionTitle() string`

GetVersionTitle returns the VersionTitle field if non-nil, zero value otherwise.

### GetVersionTitleOk

`func (o *AwsProductVersion) GetVersionTitleOk() (*string, bool)`

GetVersionTitleOk returns a tuple with the VersionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionTitle

`func (o *AwsProductVersion) SetVersionTitle(v string)`

SetVersionTitle sets VersionTitle field to given value.

### HasVersionTitle

`func (o *AwsProductVersion) HasVersionTitle() bool`

HasVersionTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


