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

// checks if the AzureMarketplaceReseller type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceReseller{}

// AzureMarketplaceReseller struct for AzureMarketplaceReseller
type AzureMarketplaceReseller struct {
	Schema               *string                           `json:"$schema,omitempty"`
	Audiences            []AzureMarketplacePreviewAudience `json:"audiences,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	Product              *string                           `json:"product,omitempty"`
	ResellerChannelState *string                           `json:"resellerChannelState,omitempty"`
	ResourceName         *string                           `json:"resourceName,omitempty"`
	Validations          []AzureMarketplaceValidation      `json:"validations,omitempty"`
}

// NewAzureMarketplaceReseller instantiates a new AzureMarketplaceReseller object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceReseller() *AzureMarketplaceReseller {
	this := AzureMarketplaceReseller{}
	return &this
}

// NewAzureMarketplaceResellerWithDefaults instantiates a new AzureMarketplaceReseller object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceResellerWithDefaults() *AzureMarketplaceReseller {
	this := AzureMarketplaceReseller{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *AzureMarketplaceReseller) SetSchema(v string) {
	o.Schema = &v
}

// GetAudiences returns the Audiences field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetAudiences() []AzureMarketplacePreviewAudience {
	if o == nil || IsNil(o.Audiences) {
		var ret []AzureMarketplacePreviewAudience
		return ret
	}
	return o.Audiences
}

// GetAudiencesOk returns a tuple with the Audiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetAudiencesOk() ([]AzureMarketplacePreviewAudience, bool) {
	if o == nil || IsNil(o.Audiences) {
		return nil, false
	}
	return o.Audiences, true
}

// HasAudiences returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasAudiences() bool {
	if o != nil && !IsNil(o.Audiences) {
		return true
	}

	return false
}

// SetAudiences gets a reference to the given []AzureMarketplacePreviewAudience and assigns it to the Audiences field.
func (o *AzureMarketplaceReseller) SetAudiences(v []AzureMarketplacePreviewAudience) {
	o.Audiences = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplaceReseller) SetId(v string) {
	o.Id = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AzureMarketplaceReseller) SetProduct(v string) {
	o.Product = &v
}

// GetResellerChannelState returns the ResellerChannelState field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetResellerChannelState() string {
	if o == nil || IsNil(o.ResellerChannelState) {
		var ret string
		return ret
	}
	return *o.ResellerChannelState
}

// GetResellerChannelStateOk returns a tuple with the ResellerChannelState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetResellerChannelStateOk() (*string, bool) {
	if o == nil || IsNil(o.ResellerChannelState) {
		return nil, false
	}
	return o.ResellerChannelState, true
}

// HasResellerChannelState returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasResellerChannelState() bool {
	if o != nil && !IsNil(o.ResellerChannelState) {
		return true
	}

	return false
}

// SetResellerChannelState gets a reference to the given string and assigns it to the ResellerChannelState field.
func (o *AzureMarketplaceReseller) SetResellerChannelState(v string) {
	o.ResellerChannelState = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *AzureMarketplaceReseller) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *AzureMarketplaceReseller) GetValidations() []AzureMarketplaceValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []AzureMarketplaceValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceReseller) GetValidationsOk() ([]AzureMarketplaceValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *AzureMarketplaceReseller) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []AzureMarketplaceValidation and assigns it to the Validations field.
func (o *AzureMarketplaceReseller) SetValidations(v []AzureMarketplaceValidation) {
	o.Validations = v
}

func (o AzureMarketplaceReseller) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceReseller) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Audiences) {
		toSerialize["audiences"] = o.Audiences
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.ResellerChannelState) {
		toSerialize["resellerChannelState"] = o.ResellerChannelState
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceReseller struct {
	value *AzureMarketplaceReseller
	isSet bool
}

func (v NullableAzureMarketplaceReseller) Get() *AzureMarketplaceReseller {
	return v.value
}

func (v *NullableAzureMarketplaceReseller) Set(val *AzureMarketplaceReseller) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceReseller) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceReseller) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceReseller(val *AzureMarketplaceReseller) *NullableAzureMarketplaceReseller {
	return &NullableAzureMarketplaceReseller{value: val, isSet: true}
}

func (v NullableAzureMarketplaceReseller) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceReseller) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}