# EntitlementTermInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DimensionQuantityDecimalParts** | Pointer to **map[string]float32** | The decimal part of the dimension quantity, in format of &lt;dimensionKey, decimalPart&gt; It is used to save the decimal part of the dimension quantity for AWS Marketplace only. Because AWS Marketplace only accepts integer for dimension quantity. If the dimension quantity is a decimal number, we need to save the decimal part for future use. | [optional] 
**IsCommitDivided** | Pointer to **bool** | Whether the commit is divided into multiple sub entitlement terms. If true, it has subEntitlementTermIds. | [optional] 
**ParentEntitlementTermId** | Pointer to **string** | The partner&#39;s entitlement term ID. It stands for the partner&#39;s entitlement term. Applicable to the sub entitlement term only. | [optional] 
**SubEntitlementTermIds** | Pointer to **[]string** | All sub entitlement terms id of this entitlement term if it is commit divided. | [optional] 
**Type** | Pointer to [**EntitlementTermType**](EntitlementTermType.md) |  | [optional] 

## Methods

### NewEntitlementTermInfo

`func NewEntitlementTermInfo() *EntitlementTermInfo`

NewEntitlementTermInfo instantiates a new EntitlementTermInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementTermInfoWithDefaults

`func NewEntitlementTermInfoWithDefaults() *EntitlementTermInfo`

NewEntitlementTermInfoWithDefaults instantiates a new EntitlementTermInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensionQuantityDecimalParts

`func (o *EntitlementTermInfo) GetDimensionQuantityDecimalParts() map[string]float32`

GetDimensionQuantityDecimalParts returns the DimensionQuantityDecimalParts field if non-nil, zero value otherwise.

### GetDimensionQuantityDecimalPartsOk

`func (o *EntitlementTermInfo) GetDimensionQuantityDecimalPartsOk() (*map[string]float32, bool)`

GetDimensionQuantityDecimalPartsOk returns a tuple with the DimensionQuantityDecimalParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionQuantityDecimalParts

`func (o *EntitlementTermInfo) SetDimensionQuantityDecimalParts(v map[string]float32)`

SetDimensionQuantityDecimalParts sets DimensionQuantityDecimalParts field to given value.

### HasDimensionQuantityDecimalParts

`func (o *EntitlementTermInfo) HasDimensionQuantityDecimalParts() bool`

HasDimensionQuantityDecimalParts returns a boolean if a field has been set.

### GetIsCommitDivided

`func (o *EntitlementTermInfo) GetIsCommitDivided() bool`

GetIsCommitDivided returns the IsCommitDivided field if non-nil, zero value otherwise.

### GetIsCommitDividedOk

`func (o *EntitlementTermInfo) GetIsCommitDividedOk() (*bool, bool)`

GetIsCommitDividedOk returns a tuple with the IsCommitDivided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommitDivided

`func (o *EntitlementTermInfo) SetIsCommitDivided(v bool)`

SetIsCommitDivided sets IsCommitDivided field to given value.

### HasIsCommitDivided

`func (o *EntitlementTermInfo) HasIsCommitDivided() bool`

HasIsCommitDivided returns a boolean if a field has been set.

### GetParentEntitlementTermId

`func (o *EntitlementTermInfo) GetParentEntitlementTermId() string`

GetParentEntitlementTermId returns the ParentEntitlementTermId field if non-nil, zero value otherwise.

### GetParentEntitlementTermIdOk

`func (o *EntitlementTermInfo) GetParentEntitlementTermIdOk() (*string, bool)`

GetParentEntitlementTermIdOk returns a tuple with the ParentEntitlementTermId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEntitlementTermId

`func (o *EntitlementTermInfo) SetParentEntitlementTermId(v string)`

SetParentEntitlementTermId sets ParentEntitlementTermId field to given value.

### HasParentEntitlementTermId

`func (o *EntitlementTermInfo) HasParentEntitlementTermId() bool`

HasParentEntitlementTermId returns a boolean if a field has been set.

### GetSubEntitlementTermIds

`func (o *EntitlementTermInfo) GetSubEntitlementTermIds() []string`

GetSubEntitlementTermIds returns the SubEntitlementTermIds field if non-nil, zero value otherwise.

### GetSubEntitlementTermIdsOk

`func (o *EntitlementTermInfo) GetSubEntitlementTermIdsOk() (*[]string, bool)`

GetSubEntitlementTermIdsOk returns a tuple with the SubEntitlementTermIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubEntitlementTermIds

`func (o *EntitlementTermInfo) SetSubEntitlementTermIds(v []string)`

SetSubEntitlementTermIds sets SubEntitlementTermIds field to given value.

### HasSubEntitlementTermIds

`func (o *EntitlementTermInfo) HasSubEntitlementTermIds() bool`

HasSubEntitlementTermIds returns a boolean if a field has been set.

### GetType

`func (o *EntitlementTermInfo) GetType() EntitlementTermType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntitlementTermInfo) GetTypeOk() (*EntitlementTermType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntitlementTermInfo) SetType(v EntitlementTermType)`

SetType sets Type field to given value.

### HasType

`func (o *EntitlementTermInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


