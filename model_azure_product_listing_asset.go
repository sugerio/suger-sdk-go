/*
Suger API

CRUD operations on a set of resources, including organizations, products, offers, entitlements, usage record groups for meterting, etc.

API version: 1.0
Contact: support@suger.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suger

import (
	"encoding/json"
)

// checks if the AzureProductListingAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureProductListingAsset{}

// AzureProductListingAsset struct for AzureProductListingAsset
type AzureProductListingAsset struct {
	Description            *string `json:"description,omitempty"`
	FileName               *string `json:"fileName,omitempty"`
	FileSasUri             *string `json:"fileSasUri,omitempty"`
	FriendlyName           *string `json:"friendlyName,omitempty"`
	Id                     *string `json:"id,omitempty"`
	Order                  *int32  `json:"order,omitempty"`
	PublisherDefinedSasUri *string `json:"publisherDefinedSasUri,omitempty"`
	ResourceType           *string `json:"resourceType,omitempty"`
	State                  *string `json:"state,omitempty"`
	Type                   *string `json:"type,omitempty"`
}

// NewAzureProductListingAsset instantiates a new AzureProductListingAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProductListingAsset() *AzureProductListingAsset {
	this := AzureProductListingAsset{}
	return &this
}

// NewAzureProductListingAssetWithDefaults instantiates a new AzureProductListingAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProductListingAssetWithDefaults() *AzureProductListingAsset {
	this := AzureProductListingAsset{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureProductListingAsset) SetDescription(v string) {
	o.Description = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *AzureProductListingAsset) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSasUri returns the FileSasUri field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetFileSasUri() string {
	if o == nil || IsNil(o.FileSasUri) {
		var ret string
		return ret
	}
	return *o.FileSasUri
}

// GetFileSasUriOk returns a tuple with the FileSasUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetFileSasUriOk() (*string, bool) {
	if o == nil || IsNil(o.FileSasUri) {
		return nil, false
	}
	return o.FileSasUri, true
}

// HasFileSasUri returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasFileSasUri() bool {
	if o != nil && !IsNil(o.FileSasUri) {
		return true
	}

	return false
}

// SetFileSasUri gets a reference to the given string and assigns it to the FileSasUri field.
func (o *AzureProductListingAsset) SetFileSasUri(v string) {
	o.FileSasUri = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *AzureProductListingAsset) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureProductListingAsset) SetId(v string) {
	o.Id = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetOrder() int32 {
	if o == nil || IsNil(o.Order) {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *AzureProductListingAsset) SetOrder(v int32) {
	o.Order = &v
}

// GetPublisherDefinedSasUri returns the PublisherDefinedSasUri field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetPublisherDefinedSasUri() string {
	if o == nil || IsNil(o.PublisherDefinedSasUri) {
		var ret string
		return ret
	}
	return *o.PublisherDefinedSasUri
}

// GetPublisherDefinedSasUriOk returns a tuple with the PublisherDefinedSasUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetPublisherDefinedSasUriOk() (*string, bool) {
	if o == nil || IsNil(o.PublisherDefinedSasUri) {
		return nil, false
	}
	return o.PublisherDefinedSasUri, true
}

// HasPublisherDefinedSasUri returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasPublisherDefinedSasUri() bool {
	if o != nil && !IsNil(o.PublisherDefinedSasUri) {
		return true
	}

	return false
}

// SetPublisherDefinedSasUri gets a reference to the given string and assigns it to the PublisherDefinedSasUri field.
func (o *AzureProductListingAsset) SetPublisherDefinedSasUri(v string) {
	o.PublisherDefinedSasUri = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AzureProductListingAsset) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AzureProductListingAsset) SetState(v string) {
	o.State = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureProductListingAsset) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListingAsset) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureProductListingAsset) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AzureProductListingAsset) SetType(v string) {
	o.Type = &v
}

func (o AzureProductListingAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProductListingAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FileSasUri) {
		toSerialize["fileSasUri"] = o.FileSasUri
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendlyName"] = o.FriendlyName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.PublisherDefinedSasUri) {
		toSerialize["publisherDefinedSasUri"] = o.PublisherDefinedSasUri
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAzureProductListingAsset struct {
	value *AzureProductListingAsset
	isSet bool
}

func (v NullableAzureProductListingAsset) Get() *AzureProductListingAsset {
	return v.value
}

func (v *NullableAzureProductListingAsset) Set(val *AzureProductListingAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProductListingAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProductListingAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProductListingAsset(val *AzureProductListingAsset) *NullableAzureProductListingAsset {
	return &NullableAzureProductListingAsset{value: val, isSet: true}
}

func (v NullableAzureProductListingAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProductListingAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
