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

// checks if the AzureMarketplaceProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceProduct{}

// AzureMarketplaceProduct struct for AzureMarketplaceProduct
type AzureMarketplaceProduct struct {
	Schema *string `json:"$schema,omitempty"`
	// The Product Display Name
	Alias *string `json:"alias,omitempty"`
	// in format of \"product/product-durable-id\"
	Id             *string                                 `json:"id,omitempty"`
	Identity       *AzureMarketplaceIdentity               `json:"identity,omitempty"`
	LifecycleState *AzureMarketplaceResourceLifecycleState `json:"lifecycleState,omitempty"`
	ProductGroup   *string                                 `json:"productGroup,omitempty"`
	ResourceName   *string                                 `json:"resourceName,omitempty"`
	Type           *AzureMarketplaceProductType            `json:"type,omitempty"`
	Validations    []AzureMarketplaceValidation            `json:"validations,omitempty"`
}

// NewAzureMarketplaceProduct instantiates a new AzureMarketplaceProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceProduct() *AzureMarketplaceProduct {
	this := AzureMarketplaceProduct{}
	return &this
}

// NewAzureMarketplaceProductWithDefaults instantiates a new AzureMarketplaceProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceProductWithDefaults() *AzureMarketplaceProduct {
	this := AzureMarketplaceProduct{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *AzureMarketplaceProduct) SetSchema(v string) {
	o.Schema = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *AzureMarketplaceProduct) SetAlias(v string) {
	o.Alias = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplaceProduct) SetId(v string) {
	o.Id = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetIdentity() AzureMarketplaceIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret AzureMarketplaceIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetIdentityOk() (*AzureMarketplaceIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given AzureMarketplaceIdentity and assigns it to the Identity field.
func (o *AzureMarketplaceProduct) SetIdentity(v AzureMarketplaceIdentity) {
	o.Identity = &v
}

// GetLifecycleState returns the LifecycleState field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetLifecycleState() AzureMarketplaceResourceLifecycleState {
	if o == nil || IsNil(o.LifecycleState) {
		var ret AzureMarketplaceResourceLifecycleState
		return ret
	}
	return *o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool) {
	if o == nil || IsNil(o.LifecycleState) {
		return nil, false
	}
	return o.LifecycleState, true
}

// HasLifecycleState returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasLifecycleState() bool {
	if o != nil && !IsNil(o.LifecycleState) {
		return true
	}

	return false
}

// SetLifecycleState gets a reference to the given AzureMarketplaceResourceLifecycleState and assigns it to the LifecycleState field.
func (o *AzureMarketplaceProduct) SetLifecycleState(v AzureMarketplaceResourceLifecycleState) {
	o.LifecycleState = &v
}

// GetProductGroup returns the ProductGroup field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetProductGroup() string {
	if o == nil || IsNil(o.ProductGroup) {
		var ret string
		return ret
	}
	return *o.ProductGroup
}

// GetProductGroupOk returns a tuple with the ProductGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetProductGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ProductGroup) {
		return nil, false
	}
	return o.ProductGroup, true
}

// HasProductGroup returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasProductGroup() bool {
	if o != nil && !IsNil(o.ProductGroup) {
		return true
	}

	return false
}

// SetProductGroup gets a reference to the given string and assigns it to the ProductGroup field.
func (o *AzureMarketplaceProduct) SetProductGroup(v string) {
	o.ProductGroup = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *AzureMarketplaceProduct) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetType() AzureMarketplaceProductType {
	if o == nil || IsNil(o.Type) {
		var ret AzureMarketplaceProductType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetTypeOk() (*AzureMarketplaceProductType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AzureMarketplaceProductType and assigns it to the Type field.
func (o *AzureMarketplaceProduct) SetType(v AzureMarketplaceProductType) {
	o.Type = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *AzureMarketplaceProduct) GetValidations() []AzureMarketplaceValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []AzureMarketplaceValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProduct) GetValidationsOk() ([]AzureMarketplaceValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *AzureMarketplaceProduct) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []AzureMarketplaceValidation and assigns it to the Validations field.
func (o *AzureMarketplaceProduct) SetValidations(v []AzureMarketplaceValidation) {
	o.Validations = v
}

func (o AzureMarketplaceProduct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !IsNil(o.LifecycleState) {
		toSerialize["lifecycleState"] = o.LifecycleState
	}
	if !IsNil(o.ProductGroup) {
		toSerialize["productGroup"] = o.ProductGroup
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceProduct struct {
	value *AzureMarketplaceProduct
	isSet bool
}

func (v NullableAzureMarketplaceProduct) Get() *AzureMarketplaceProduct {
	return v.value
}

func (v *NullableAzureMarketplaceProduct) Set(val *AzureMarketplaceProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceProduct(val *AzureMarketplaceProduct) *NullableAzureMarketplaceProduct {
	return &NullableAzureMarketplaceProduct{value: val, isSet: true}
}

func (v NullableAzureMarketplaceProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
