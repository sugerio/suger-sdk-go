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

// checks if the AwsProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsProduct{}

// AwsProduct struct for AwsProduct
type AwsProduct struct {
	Description               *AwsProductDescription               `json:"Description,omitempty"`
	Dimensions                []AwsProductDimension                `json:"Dimensions,omitempty"`
	PromotionalResources      *AwsProductPromotionalResources      `json:"PromotionalResources,omitempty"`
	Repositories              []AwsProductRepository               `json:"Repositories,omitempty"`
	SignatureVerificationKeys []AwsProductSignatureVerificationKey `json:"SignatureVerificationKeys,omitempty"`
	SupportInformation        *AwsProductSupportInformation        `json:"SupportInformation,omitempty"`
	Versions                  []AwsProductVersion                  `json:"Versions,omitempty"`
	// The product Id in AWS Marketplace Data Feed Service.
	DataFeedProductId *string `json:"dataFeedProductId,omitempty"`
	// AWS Product ID
	ProductId *string `json:"productId,omitempty"`
}

// NewAwsProduct instantiates a new AwsProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsProduct() *AwsProduct {
	this := AwsProduct{}
	return &this
}

// NewAwsProductWithDefaults instantiates a new AwsProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsProductWithDefaults() *AwsProduct {
	this := AwsProduct{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AwsProduct) GetDescription() AwsProductDescription {
	if o == nil || IsNil(o.Description) {
		var ret AwsProductDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetDescriptionOk() (*AwsProductDescription, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AwsProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given AwsProductDescription and assigns it to the Description field.
func (o *AwsProduct) SetDescription(v AwsProductDescription) {
	o.Description = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *AwsProduct) GetDimensions() []AwsProductDimension {
	if o == nil || IsNil(o.Dimensions) {
		var ret []AwsProductDimension
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetDimensionsOk() ([]AwsProductDimension, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *AwsProduct) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []AwsProductDimension and assigns it to the Dimensions field.
func (o *AwsProduct) SetDimensions(v []AwsProductDimension) {
	o.Dimensions = v
}

// GetPromotionalResources returns the PromotionalResources field value if set, zero value otherwise.
func (o *AwsProduct) GetPromotionalResources() AwsProductPromotionalResources {
	if o == nil || IsNil(o.PromotionalResources) {
		var ret AwsProductPromotionalResources
		return ret
	}
	return *o.PromotionalResources
}

// GetPromotionalResourcesOk returns a tuple with the PromotionalResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetPromotionalResourcesOk() (*AwsProductPromotionalResources, bool) {
	if o == nil || IsNil(o.PromotionalResources) {
		return nil, false
	}
	return o.PromotionalResources, true
}

// HasPromotionalResources returns a boolean if a field has been set.
func (o *AwsProduct) HasPromotionalResources() bool {
	if o != nil && !IsNil(o.PromotionalResources) {
		return true
	}

	return false
}

// SetPromotionalResources gets a reference to the given AwsProductPromotionalResources and assigns it to the PromotionalResources field.
func (o *AwsProduct) SetPromotionalResources(v AwsProductPromotionalResources) {
	o.PromotionalResources = &v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *AwsProduct) GetRepositories() []AwsProductRepository {
	if o == nil || IsNil(o.Repositories) {
		var ret []AwsProductRepository
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetRepositoriesOk() ([]AwsProductRepository, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *AwsProduct) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []AwsProductRepository and assigns it to the Repositories field.
func (o *AwsProduct) SetRepositories(v []AwsProductRepository) {
	o.Repositories = v
}

// GetSignatureVerificationKeys returns the SignatureVerificationKeys field value if set, zero value otherwise.
func (o *AwsProduct) GetSignatureVerificationKeys() []AwsProductSignatureVerificationKey {
	if o == nil || IsNil(o.SignatureVerificationKeys) {
		var ret []AwsProductSignatureVerificationKey
		return ret
	}
	return o.SignatureVerificationKeys
}

// GetSignatureVerificationKeysOk returns a tuple with the SignatureVerificationKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetSignatureVerificationKeysOk() ([]AwsProductSignatureVerificationKey, bool) {
	if o == nil || IsNil(o.SignatureVerificationKeys) {
		return nil, false
	}
	return o.SignatureVerificationKeys, true
}

// HasSignatureVerificationKeys returns a boolean if a field has been set.
func (o *AwsProduct) HasSignatureVerificationKeys() bool {
	if o != nil && !IsNil(o.SignatureVerificationKeys) {
		return true
	}

	return false
}

// SetSignatureVerificationKeys gets a reference to the given []AwsProductSignatureVerificationKey and assigns it to the SignatureVerificationKeys field.
func (o *AwsProduct) SetSignatureVerificationKeys(v []AwsProductSignatureVerificationKey) {
	o.SignatureVerificationKeys = v
}

// GetSupportInformation returns the SupportInformation field value if set, zero value otherwise.
func (o *AwsProduct) GetSupportInformation() AwsProductSupportInformation {
	if o == nil || IsNil(o.SupportInformation) {
		var ret AwsProductSupportInformation
		return ret
	}
	return *o.SupportInformation
}

// GetSupportInformationOk returns a tuple with the SupportInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetSupportInformationOk() (*AwsProductSupportInformation, bool) {
	if o == nil || IsNil(o.SupportInformation) {
		return nil, false
	}
	return o.SupportInformation, true
}

// HasSupportInformation returns a boolean if a field has been set.
func (o *AwsProduct) HasSupportInformation() bool {
	if o != nil && !IsNil(o.SupportInformation) {
		return true
	}

	return false
}

// SetSupportInformation gets a reference to the given AwsProductSupportInformation and assigns it to the SupportInformation field.
func (o *AwsProduct) SetSupportInformation(v AwsProductSupportInformation) {
	o.SupportInformation = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *AwsProduct) GetVersions() []AwsProductVersion {
	if o == nil || IsNil(o.Versions) {
		var ret []AwsProductVersion
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetVersionsOk() ([]AwsProductVersion, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *AwsProduct) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []AwsProductVersion and assigns it to the Versions field.
func (o *AwsProduct) SetVersions(v []AwsProductVersion) {
	o.Versions = v
}

// GetDataFeedProductId returns the DataFeedProductId field value if set, zero value otherwise.
func (o *AwsProduct) GetDataFeedProductId() string {
	if o == nil || IsNil(o.DataFeedProductId) {
		var ret string
		return ret
	}
	return *o.DataFeedProductId
}

// GetDataFeedProductIdOk returns a tuple with the DataFeedProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetDataFeedProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.DataFeedProductId) {
		return nil, false
	}
	return o.DataFeedProductId, true
}

// HasDataFeedProductId returns a boolean if a field has been set.
func (o *AwsProduct) HasDataFeedProductId() bool {
	if o != nil && !IsNil(o.DataFeedProductId) {
		return true
	}

	return false
}

// SetDataFeedProductId gets a reference to the given string and assigns it to the DataFeedProductId field.
func (o *AwsProduct) SetDataFeedProductId(v string) {
	o.DataFeedProductId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AwsProduct) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsProduct) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AwsProduct) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AwsProduct) SetProductId(v string) {
	o.ProductId = &v
}

func (o AwsProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Dimensions) {
		toSerialize["Dimensions"] = o.Dimensions
	}
	if !IsNil(o.PromotionalResources) {
		toSerialize["PromotionalResources"] = o.PromotionalResources
	}
	if !IsNil(o.Repositories) {
		toSerialize["Repositories"] = o.Repositories
	}
	if !IsNil(o.SignatureVerificationKeys) {
		toSerialize["SignatureVerificationKeys"] = o.SignatureVerificationKeys
	}
	if !IsNil(o.SupportInformation) {
		toSerialize["SupportInformation"] = o.SupportInformation
	}
	if !IsNil(o.Versions) {
		toSerialize["Versions"] = o.Versions
	}
	if !IsNil(o.DataFeedProductId) {
		toSerialize["dataFeedProductId"] = o.DataFeedProductId
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	return toSerialize, nil
}

type NullableAwsProduct struct {
	value *AwsProduct
	isSet bool
}

func (v NullableAwsProduct) Get() *AwsProduct {
	return v.value
}

func (v *NullableAwsProduct) Set(val *AwsProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsProduct(val *AwsProduct) *NullableAwsProduct {
	return &NullableAwsProduct{value: val, isSet: true}
}

func (v NullableAwsProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
