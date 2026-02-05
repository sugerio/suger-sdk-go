# OriginalEulaInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEulaUrls** | Pointer to **[]string** | The URL of the additional EULA files. Only applicable when EulaType &#x3D; CUSTOM. The additional EULA files will be attached to the EULA file in the EulaUrl, and form a single EULA file. | [optional] 
**AdditionalResellerEulaUrls** | Pointer to **[]string** | The URL of the additional reseller EULA files. Only applicable when ResellerEulaType &#x3D; CUSTOM. | [optional] 
**AttachEulaType** | Pointer to [**EulaType**](EulaType.md) | Attach the standard EULA file to the CUSTOM EULA file. Only applicable when EulaType &#x3D; CUSTOM | [optional] 
**EulaMergeOrder** | Pointer to **[]int32** | The merge order of the EULA files. Only applicable when EulaType &#x3D; CUSTOM. Elements are the original index of the EULA files in the index they should be transferred to, where original indexes are: AttachEulaType is index 0, EulaUrl is index 1, additionalEulaUrls is index 2 onwards. | [optional] 
**EulaType** | Pointer to [**EulaType**](EulaType.md) | The type of the EULA. | [optional] 
**EulaUrl** | Pointer to **string** | The URL of the EULA file. | [optional] 
**ResellerAttachEulaType** | Pointer to [**EulaType**](EulaType.md) | Attach the standard EULA file to the CUSTOM EULA file. Only applicable when EulaType &#x3D; CUSTOM | [optional] 
**ResellerEulaType** | Pointer to [**EulaType**](EulaType.md) | The type of the reseller EULA. Only applicable for CPPO offers. | [optional] 
**ResellerEulaUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewOriginalEulaInfo

`func NewOriginalEulaInfo() *OriginalEulaInfo`

NewOriginalEulaInfo instantiates a new OriginalEulaInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginalEulaInfoWithDefaults

`func NewOriginalEulaInfoWithDefaults() *OriginalEulaInfo`

NewOriginalEulaInfoWithDefaults instantiates a new OriginalEulaInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalEulaUrls

`func (o *OriginalEulaInfo) GetAdditionalEulaUrls() []string`

GetAdditionalEulaUrls returns the AdditionalEulaUrls field if non-nil, zero value otherwise.

### GetAdditionalEulaUrlsOk

`func (o *OriginalEulaInfo) GetAdditionalEulaUrlsOk() (*[]string, bool)`

GetAdditionalEulaUrlsOk returns a tuple with the AdditionalEulaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEulaUrls

`func (o *OriginalEulaInfo) SetAdditionalEulaUrls(v []string)`

SetAdditionalEulaUrls sets AdditionalEulaUrls field to given value.

### HasAdditionalEulaUrls

`func (o *OriginalEulaInfo) HasAdditionalEulaUrls() bool`

HasAdditionalEulaUrls returns a boolean if a field has been set.

### GetAdditionalResellerEulaUrls

`func (o *OriginalEulaInfo) GetAdditionalResellerEulaUrls() []string`

GetAdditionalResellerEulaUrls returns the AdditionalResellerEulaUrls field if non-nil, zero value otherwise.

### GetAdditionalResellerEulaUrlsOk

`func (o *OriginalEulaInfo) GetAdditionalResellerEulaUrlsOk() (*[]string, bool)`

GetAdditionalResellerEulaUrlsOk returns a tuple with the AdditionalResellerEulaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalResellerEulaUrls

`func (o *OriginalEulaInfo) SetAdditionalResellerEulaUrls(v []string)`

SetAdditionalResellerEulaUrls sets AdditionalResellerEulaUrls field to given value.

### HasAdditionalResellerEulaUrls

`func (o *OriginalEulaInfo) HasAdditionalResellerEulaUrls() bool`

HasAdditionalResellerEulaUrls returns a boolean if a field has been set.

### GetAttachEulaType

`func (o *OriginalEulaInfo) GetAttachEulaType() EulaType`

GetAttachEulaType returns the AttachEulaType field if non-nil, zero value otherwise.

### GetAttachEulaTypeOk

`func (o *OriginalEulaInfo) GetAttachEulaTypeOk() (*EulaType, bool)`

GetAttachEulaTypeOk returns a tuple with the AttachEulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachEulaType

`func (o *OriginalEulaInfo) SetAttachEulaType(v EulaType)`

SetAttachEulaType sets AttachEulaType field to given value.

### HasAttachEulaType

`func (o *OriginalEulaInfo) HasAttachEulaType() bool`

HasAttachEulaType returns a boolean if a field has been set.

### GetEulaMergeOrder

`func (o *OriginalEulaInfo) GetEulaMergeOrder() []int32`

GetEulaMergeOrder returns the EulaMergeOrder field if non-nil, zero value otherwise.

### GetEulaMergeOrderOk

`func (o *OriginalEulaInfo) GetEulaMergeOrderOk() (*[]int32, bool)`

GetEulaMergeOrderOk returns a tuple with the EulaMergeOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaMergeOrder

`func (o *OriginalEulaInfo) SetEulaMergeOrder(v []int32)`

SetEulaMergeOrder sets EulaMergeOrder field to given value.

### HasEulaMergeOrder

`func (o *OriginalEulaInfo) HasEulaMergeOrder() bool`

HasEulaMergeOrder returns a boolean if a field has been set.

### GetEulaType

`func (o *OriginalEulaInfo) GetEulaType() EulaType`

GetEulaType returns the EulaType field if non-nil, zero value otherwise.

### GetEulaTypeOk

`func (o *OriginalEulaInfo) GetEulaTypeOk() (*EulaType, bool)`

GetEulaTypeOk returns a tuple with the EulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaType

`func (o *OriginalEulaInfo) SetEulaType(v EulaType)`

SetEulaType sets EulaType field to given value.

### HasEulaType

`func (o *OriginalEulaInfo) HasEulaType() bool`

HasEulaType returns a boolean if a field has been set.

### GetEulaUrl

`func (o *OriginalEulaInfo) GetEulaUrl() string`

GetEulaUrl returns the EulaUrl field if non-nil, zero value otherwise.

### GetEulaUrlOk

`func (o *OriginalEulaInfo) GetEulaUrlOk() (*string, bool)`

GetEulaUrlOk returns a tuple with the EulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEulaUrl

`func (o *OriginalEulaInfo) SetEulaUrl(v string)`

SetEulaUrl sets EulaUrl field to given value.

### HasEulaUrl

`func (o *OriginalEulaInfo) HasEulaUrl() bool`

HasEulaUrl returns a boolean if a field has been set.

### GetResellerAttachEulaType

`func (o *OriginalEulaInfo) GetResellerAttachEulaType() EulaType`

GetResellerAttachEulaType returns the ResellerAttachEulaType field if non-nil, zero value otherwise.

### GetResellerAttachEulaTypeOk

`func (o *OriginalEulaInfo) GetResellerAttachEulaTypeOk() (*EulaType, bool)`

GetResellerAttachEulaTypeOk returns a tuple with the ResellerAttachEulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerAttachEulaType

`func (o *OriginalEulaInfo) SetResellerAttachEulaType(v EulaType)`

SetResellerAttachEulaType sets ResellerAttachEulaType field to given value.

### HasResellerAttachEulaType

`func (o *OriginalEulaInfo) HasResellerAttachEulaType() bool`

HasResellerAttachEulaType returns a boolean if a field has been set.

### GetResellerEulaType

`func (o *OriginalEulaInfo) GetResellerEulaType() EulaType`

GetResellerEulaType returns the ResellerEulaType field if non-nil, zero value otherwise.

### GetResellerEulaTypeOk

`func (o *OriginalEulaInfo) GetResellerEulaTypeOk() (*EulaType, bool)`

GetResellerEulaTypeOk returns a tuple with the ResellerEulaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerEulaType

`func (o *OriginalEulaInfo) SetResellerEulaType(v EulaType)`

SetResellerEulaType sets ResellerEulaType field to given value.

### HasResellerEulaType

`func (o *OriginalEulaInfo) HasResellerEulaType() bool`

HasResellerEulaType returns a boolean if a field has been set.

### GetResellerEulaUrl

`func (o *OriginalEulaInfo) GetResellerEulaUrl() string`

GetResellerEulaUrl returns the ResellerEulaUrl field if non-nil, zero value otherwise.

### GetResellerEulaUrlOk

`func (o *OriginalEulaInfo) GetResellerEulaUrlOk() (*string, bool)`

GetResellerEulaUrlOk returns a tuple with the ResellerEulaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResellerEulaUrl

`func (o *OriginalEulaInfo) SetResellerEulaUrl(v string)`

SetResellerEulaUrl sets ResellerEulaUrl field to given value.

### HasResellerEulaUrl

`func (o *OriginalEulaInfo) HasResellerEulaUrl() bool`

HasResellerEulaUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


