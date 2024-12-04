/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AzureMarketplaceListingAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceListingAsset{}

// AzureMarketplaceListingAsset struct for AzureMarketplaceListingAsset
type AzureMarketplaceListingAsset struct {
	Schema      *string `json:"$schema,omitempty"`
	Description *string `json:"description,omitempty"`
	// minimum: 0
	DisplayOrder *int32  `json:"displayOrder,omitempty"`
	FileName     *string `json:"fileName,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
	Id           *string `json:"id,omitempty"`
	Kind         *string `json:"kind,omitempty"`
	// Max string length is 10.
	LanguageId *string `json:"languageId,omitempty"`
	// Default value is \"generallyAvailable\".
	LifecycleState *AzureMarketplaceResourceLifecycleState `json:"lifecycleState,omitempty"`
	Listing        *string                                 `json:"listing,omitempty"`
	// Product resource name, in format of \"product/product-durable-id\"
	Product      *string                           `json:"product,omitempty"`
	ResourceName *string                           `json:"resourceName,omitempty"`
	Type         *AzureMarketplaceListingAssetType `json:"type,omitempty"`
	// pattern: \"^https?://\"
	Url         *string                      `json:"url,omitempty"`
	Validations []AzureMarketplaceValidation `json:"validations,omitempty"`
}

// NewAzureMarketplaceListingAsset instantiates a new AzureMarketplaceListingAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceListingAsset() *AzureMarketplaceListingAsset {
	this := AzureMarketplaceListingAsset{}
	return &this
}

// NewAzureMarketplaceListingAssetWithDefaults instantiates a new AzureMarketplaceListingAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceListingAssetWithDefaults() *AzureMarketplaceListingAsset {
	this := AzureMarketplaceListingAsset{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *AzureMarketplaceListingAsset) SetSchema(v string) {
	o.Schema = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureMarketplaceListingAsset) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *AzureMarketplaceListingAsset) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AzureMarketplaceListingAsset) SetFileName(v string) {
	o.FileName = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *AzureMarketplaceListingAsset) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplaceListingAsset) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AzureMarketplaceListingAsset) SetKind(v string) {
	o.Kind = &v
}

// GetLanguageId returns the LanguageId field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetLanguageId() string {
	if o == nil || IsNil(o.LanguageId) {
		var ret string
		return ret
	}
	return *o.LanguageId
}

// GetLanguageIdOk returns a tuple with the LanguageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetLanguageIdOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageId) {
		return nil, false
	}
	return o.LanguageId, true
}

// HasLanguageId returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasLanguageId() bool {
	if o != nil && !IsNil(o.LanguageId) {
		return true
	}

	return false
}

// SetLanguageId gets a reference to the given string and assigns it to the LanguageId field.
func (o *AzureMarketplaceListingAsset) SetLanguageId(v string) {
	o.LanguageId = &v
}

// GetLifecycleState returns the LifecycleState field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetLifecycleState() AzureMarketplaceResourceLifecycleState {
	if o == nil || IsNil(o.LifecycleState) {
		var ret AzureMarketplaceResourceLifecycleState
		return ret
	}
	return *o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool) {
	if o == nil || IsNil(o.LifecycleState) {
		return nil, false
	}
	return o.LifecycleState, true
}

// HasLifecycleState returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasLifecycleState() bool {
	if o != nil && !IsNil(o.LifecycleState) {
		return true
	}

	return false
}

// SetLifecycleState gets a reference to the given AzureMarketplaceResourceLifecycleState and assigns it to the LifecycleState field.
func (o *AzureMarketplaceListingAsset) SetLifecycleState(v AzureMarketplaceResourceLifecycleState) {
	o.LifecycleState = &v
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetListing() string {
	if o == nil || IsNil(o.Listing) {
		var ret string
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetListingOk() (*string, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given string and assigns it to the Listing field.
func (o *AzureMarketplaceListingAsset) SetListing(v string) {
	o.Listing = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AzureMarketplaceListingAsset) SetProduct(v string) {
	o.Product = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *AzureMarketplaceListingAsset) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetType() AzureMarketplaceListingAssetType {
	if o == nil || IsNil(o.Type) {
		var ret AzureMarketplaceListingAssetType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetTypeOk() (*AzureMarketplaceListingAssetType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AzureMarketplaceListingAssetType and assigns it to the Type field.
func (o *AzureMarketplaceListingAsset) SetType(v AzureMarketplaceListingAssetType) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AzureMarketplaceListingAsset) SetUrl(v string) {
	o.Url = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *AzureMarketplaceListingAsset) GetValidations() []AzureMarketplaceValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []AzureMarketplaceValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceListingAsset) GetValidationsOk() ([]AzureMarketplaceValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *AzureMarketplaceListingAsset) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []AzureMarketplaceValidation and assigns it to the Validations field.
func (o *AzureMarketplaceListingAsset) SetValidations(v []AzureMarketplaceValidation) {
	o.Validations = v
}

func (o AzureMarketplaceListingAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceListingAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendlyName"] = o.FriendlyName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.LanguageId) {
		toSerialize["languageId"] = o.LanguageId
	}
	if !IsNil(o.LifecycleState) {
		toSerialize["lifecycleState"] = o.LifecycleState
	}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceListingAsset struct {
	value *AzureMarketplaceListingAsset
	isSet bool
}

func (v NullableAzureMarketplaceListingAsset) Get() *AzureMarketplaceListingAsset {
	return v.value
}

func (v *NullableAzureMarketplaceListingAsset) Set(val *AzureMarketplaceListingAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceListingAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceListingAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceListingAsset(val *AzureMarketplaceListingAsset) *NullableAzureMarketplaceListingAsset {
	return &NullableAzureMarketplaceListingAsset{value: val, isSet: true}
}

func (v NullableAzureMarketplaceListingAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceListingAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
