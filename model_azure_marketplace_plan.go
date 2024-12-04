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

// checks if the AzureMarketplacePlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePlan{}

// AzureMarketplacePlan struct for AzureMarketplacePlan
type AzureMarketplacePlan struct {
	Schema                        *string                                   `json:"$schema,omitempty"`
	Alias                         *string                                   `json:"alias,omitempty"`
	AzureGovernmentCertifications []AzureMarketplaceGovernmentCertification `json:"azureGovernmentCertifications,omitempty"`
	// enums:[azureGlobal,azureGovernment,azureGermany,azureChina]
	AzureRegions        []string                             `json:"azureRegions,omitempty"`
	DeprecationSchedule *AzureMarketplaceDeprecationSchedule `json:"deprecationSchedule,omitempty"`
	// default 2147483647
	DisplayRank *int32 `json:"displayRank,omitempty"`
	// in format of \"plan/product-durable-id/plan-durable-id\"
	Id             *string                                 `json:"id,omitempty"`
	Identity       *AzureMarketplaceIdentity               `json:"identity,omitempty"`
	LifecycleState *AzureMarketplaceResourceLifecycleState `json:"lifecycleState,omitempty"`
	// in format of \"product/product-durable-id\"
	Product      *string `json:"product,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	// Specifies the plan type (AzureApplication-type products only) see: https://go.microsoft.com/fwlink/?linkid=2106322
	Subtype     *string                      `json:"subtype,omitempty"`
	Validations []AzureMarketplaceValidation `json:"validations,omitempty"`
}

// NewAzureMarketplacePlan instantiates a new AzureMarketplacePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePlan() *AzureMarketplacePlan {
	this := AzureMarketplacePlan{}
	return &this
}

// NewAzureMarketplacePlanWithDefaults instantiates a new AzureMarketplacePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePlanWithDefaults() *AzureMarketplacePlan {
	this := AzureMarketplacePlan{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *AzureMarketplacePlan) SetSchema(v string) {
	o.Schema = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *AzureMarketplacePlan) SetAlias(v string) {
	o.Alias = &v
}

// GetAzureGovernmentCertifications returns the AzureGovernmentCertifications field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetAzureGovernmentCertifications() []AzureMarketplaceGovernmentCertification {
	if o == nil || IsNil(o.AzureGovernmentCertifications) {
		var ret []AzureMarketplaceGovernmentCertification
		return ret
	}
	return o.AzureGovernmentCertifications
}

// GetAzureGovernmentCertificationsOk returns a tuple with the AzureGovernmentCertifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetAzureGovernmentCertificationsOk() ([]AzureMarketplaceGovernmentCertification, bool) {
	if o == nil || IsNil(o.AzureGovernmentCertifications) {
		return nil, false
	}
	return o.AzureGovernmentCertifications, true
}

// HasAzureGovernmentCertifications returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasAzureGovernmentCertifications() bool {
	if o != nil && !IsNil(o.AzureGovernmentCertifications) {
		return true
	}

	return false
}

// SetAzureGovernmentCertifications gets a reference to the given []AzureMarketplaceGovernmentCertification and assigns it to the AzureGovernmentCertifications field.
func (o *AzureMarketplacePlan) SetAzureGovernmentCertifications(v []AzureMarketplaceGovernmentCertification) {
	o.AzureGovernmentCertifications = v
}

// GetAzureRegions returns the AzureRegions field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetAzureRegions() []string {
	if o == nil || IsNil(o.AzureRegions) {
		var ret []string
		return ret
	}
	return o.AzureRegions
}

// GetAzureRegionsOk returns a tuple with the AzureRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetAzureRegionsOk() ([]string, bool) {
	if o == nil || IsNil(o.AzureRegions) {
		return nil, false
	}
	return o.AzureRegions, true
}

// HasAzureRegions returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasAzureRegions() bool {
	if o != nil && !IsNil(o.AzureRegions) {
		return true
	}

	return false
}

// SetAzureRegions gets a reference to the given []string and assigns it to the AzureRegions field.
func (o *AzureMarketplacePlan) SetAzureRegions(v []string) {
	o.AzureRegions = v
}

// GetDeprecationSchedule returns the DeprecationSchedule field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetDeprecationSchedule() AzureMarketplaceDeprecationSchedule {
	if o == nil || IsNil(o.DeprecationSchedule) {
		var ret AzureMarketplaceDeprecationSchedule
		return ret
	}
	return *o.DeprecationSchedule
}

// GetDeprecationScheduleOk returns a tuple with the DeprecationSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetDeprecationScheduleOk() (*AzureMarketplaceDeprecationSchedule, bool) {
	if o == nil || IsNil(o.DeprecationSchedule) {
		return nil, false
	}
	return o.DeprecationSchedule, true
}

// HasDeprecationSchedule returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasDeprecationSchedule() bool {
	if o != nil && !IsNil(o.DeprecationSchedule) {
		return true
	}

	return false
}

// SetDeprecationSchedule gets a reference to the given AzureMarketplaceDeprecationSchedule and assigns it to the DeprecationSchedule field.
func (o *AzureMarketplacePlan) SetDeprecationSchedule(v AzureMarketplaceDeprecationSchedule) {
	o.DeprecationSchedule = &v
}

// GetDisplayRank returns the DisplayRank field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetDisplayRank() int32 {
	if o == nil || IsNil(o.DisplayRank) {
		var ret int32
		return ret
	}
	return *o.DisplayRank
}

// GetDisplayRankOk returns a tuple with the DisplayRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetDisplayRankOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayRank) {
		return nil, false
	}
	return o.DisplayRank, true
}

// HasDisplayRank returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasDisplayRank() bool {
	if o != nil && !IsNil(o.DisplayRank) {
		return true
	}

	return false
}

// SetDisplayRank gets a reference to the given int32 and assigns it to the DisplayRank field.
func (o *AzureMarketplacePlan) SetDisplayRank(v int32) {
	o.DisplayRank = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureMarketplacePlan) SetId(v string) {
	o.Id = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetIdentity() AzureMarketplaceIdentity {
	if o == nil || IsNil(o.Identity) {
		var ret AzureMarketplaceIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetIdentityOk() (*AzureMarketplaceIdentity, bool) {
	if o == nil || IsNil(o.Identity) {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasIdentity() bool {
	if o != nil && !IsNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given AzureMarketplaceIdentity and assigns it to the Identity field.
func (o *AzureMarketplacePlan) SetIdentity(v AzureMarketplaceIdentity) {
	o.Identity = &v
}

// GetLifecycleState returns the LifecycleState field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetLifecycleState() AzureMarketplaceResourceLifecycleState {
	if o == nil || IsNil(o.LifecycleState) {
		var ret AzureMarketplaceResourceLifecycleState
		return ret
	}
	return *o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetLifecycleStateOk() (*AzureMarketplaceResourceLifecycleState, bool) {
	if o == nil || IsNil(o.LifecycleState) {
		return nil, false
	}
	return o.LifecycleState, true
}

// HasLifecycleState returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasLifecycleState() bool {
	if o != nil && !IsNil(o.LifecycleState) {
		return true
	}

	return false
}

// SetLifecycleState gets a reference to the given AzureMarketplaceResourceLifecycleState and assigns it to the LifecycleState field.
func (o *AzureMarketplacePlan) SetLifecycleState(v AzureMarketplaceResourceLifecycleState) {
	o.LifecycleState = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *AzureMarketplacePlan) SetProduct(v string) {
	o.Product = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *AzureMarketplacePlan) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *AzureMarketplacePlan) SetSubtype(v string) {
	o.Subtype = &v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *AzureMarketplacePlan) GetValidations() []AzureMarketplaceValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []AzureMarketplaceValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlan) GetValidationsOk() ([]AzureMarketplaceValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *AzureMarketplacePlan) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []AzureMarketplaceValidation and assigns it to the Validations field.
func (o *AzureMarketplacePlan) SetValidations(v []AzureMarketplaceValidation) {
	o.Validations = v
}

func (o AzureMarketplacePlan) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.AzureGovernmentCertifications) {
		toSerialize["azureGovernmentCertifications"] = o.AzureGovernmentCertifications
	}
	if !IsNil(o.AzureRegions) {
		toSerialize["azureRegions"] = o.AzureRegions
	}
	if !IsNil(o.DeprecationSchedule) {
		toSerialize["deprecationSchedule"] = o.DeprecationSchedule
	}
	if !IsNil(o.DisplayRank) {
		toSerialize["displayRank"] = o.DisplayRank
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
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePlan struct {
	value *AzureMarketplacePlan
	isSet bool
}

func (v NullableAzureMarketplacePlan) Get() *AzureMarketplacePlan {
	return v.value
}

func (v *NullableAzureMarketplacePlan) Set(val *AzureMarketplacePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePlan(val *AzureMarketplacePlan) *NullableAzureMarketplacePlan {
	return &NullableAzureMarketplacePlan{value: val, isSet: true}
}

func (v NullableAzureMarketplacePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}