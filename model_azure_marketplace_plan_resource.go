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

// checks if the AzureMarketplacePlanResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplacePlanResource{}

// AzureMarketplacePlanResource struct for AzureMarketplacePlanResource
type AzureMarketplacePlanResource struct {
	Plan                     *AzureMarketplacePlan                     `json:"plan,omitempty"`
	PlanListing              *AzureMarketplacePlanListing              `json:"planListing,omitempty"`
	PriceAndAvailabilityPlan *AzureMarketplacePriceAndAvailabilityPlan `json:"priceAndAvailabilityPlan,omitempty"`
}

// NewAzureMarketplacePlanResource instantiates a new AzureMarketplacePlanResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplacePlanResource() *AzureMarketplacePlanResource {
	this := AzureMarketplacePlanResource{}
	return &this
}

// NewAzureMarketplacePlanResourceWithDefaults instantiates a new AzureMarketplacePlanResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplacePlanResourceWithDefaults() *AzureMarketplacePlanResource {
	this := AzureMarketplacePlanResource{}
	return &this
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *AzureMarketplacePlanResource) GetPlan() AzureMarketplacePlan {
	if o == nil || IsNil(o.Plan) {
		var ret AzureMarketplacePlan
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlanResource) GetPlanOk() (*AzureMarketplacePlan, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *AzureMarketplacePlanResource) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given AzureMarketplacePlan and assigns it to the Plan field.
func (o *AzureMarketplacePlanResource) SetPlan(v AzureMarketplacePlan) {
	o.Plan = &v
}

// GetPlanListing returns the PlanListing field value if set, zero value otherwise.
func (o *AzureMarketplacePlanResource) GetPlanListing() AzureMarketplacePlanListing {
	if o == nil || IsNil(o.PlanListing) {
		var ret AzureMarketplacePlanListing
		return ret
	}
	return *o.PlanListing
}

// GetPlanListingOk returns a tuple with the PlanListing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlanResource) GetPlanListingOk() (*AzureMarketplacePlanListing, bool) {
	if o == nil || IsNil(o.PlanListing) {
		return nil, false
	}
	return o.PlanListing, true
}

// HasPlanListing returns a boolean if a field has been set.
func (o *AzureMarketplacePlanResource) HasPlanListing() bool {
	if o != nil && !IsNil(o.PlanListing) {
		return true
	}

	return false
}

// SetPlanListing gets a reference to the given AzureMarketplacePlanListing and assigns it to the PlanListing field.
func (o *AzureMarketplacePlanResource) SetPlanListing(v AzureMarketplacePlanListing) {
	o.PlanListing = &v
}

// GetPriceAndAvailabilityPlan returns the PriceAndAvailabilityPlan field value if set, zero value otherwise.
func (o *AzureMarketplacePlanResource) GetPriceAndAvailabilityPlan() AzureMarketplacePriceAndAvailabilityPlan {
	if o == nil || IsNil(o.PriceAndAvailabilityPlan) {
		var ret AzureMarketplacePriceAndAvailabilityPlan
		return ret
	}
	return *o.PriceAndAvailabilityPlan
}

// GetPriceAndAvailabilityPlanOk returns a tuple with the PriceAndAvailabilityPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplacePlanResource) GetPriceAndAvailabilityPlanOk() (*AzureMarketplacePriceAndAvailabilityPlan, bool) {
	if o == nil || IsNil(o.PriceAndAvailabilityPlan) {
		return nil, false
	}
	return o.PriceAndAvailabilityPlan, true
}

// HasPriceAndAvailabilityPlan returns a boolean if a field has been set.
func (o *AzureMarketplacePlanResource) HasPriceAndAvailabilityPlan() bool {
	if o != nil && !IsNil(o.PriceAndAvailabilityPlan) {
		return true
	}

	return false
}

// SetPriceAndAvailabilityPlan gets a reference to the given AzureMarketplacePriceAndAvailabilityPlan and assigns it to the PriceAndAvailabilityPlan field.
func (o *AzureMarketplacePlanResource) SetPriceAndAvailabilityPlan(v AzureMarketplacePriceAndAvailabilityPlan) {
	o.PriceAndAvailabilityPlan = &v
}

func (o AzureMarketplacePlanResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplacePlanResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	if !IsNil(o.PlanListing) {
		toSerialize["planListing"] = o.PlanListing
	}
	if !IsNil(o.PriceAndAvailabilityPlan) {
		toSerialize["priceAndAvailabilityPlan"] = o.PriceAndAvailabilityPlan
	}
	return toSerialize, nil
}

type NullableAzureMarketplacePlanResource struct {
	value *AzureMarketplacePlanResource
	isSet bool
}

func (v NullableAzureMarketplacePlanResource) Get() *AzureMarketplacePlanResource {
	return v.value
}

func (v *NullableAzureMarketplacePlanResource) Set(val *AzureMarketplacePlanResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplacePlanResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplacePlanResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplacePlanResource(val *AzureMarketplacePlanResource) *NullableAzureMarketplacePlanResource {
	return &NullableAzureMarketplacePlanResource{value: val, isSet: true}
}

func (v NullableAzureMarketplacePlanResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplacePlanResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
