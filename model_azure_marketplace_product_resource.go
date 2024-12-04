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

// checks if the AzureMarketplaceProductResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureMarketplaceProductResource{}

// AzureMarketplaceProductResource struct for AzureMarketplaceProductResource
type AzureMarketplaceProductResource struct {
	CustomerLeads                   *AzureMarketplaceCustomerLeads                   `json:"customerLeads,omitempty"`
	Listing                         *AzureMarketplaceListing                         `json:"listing,omitempty"`
	ListingAssets                   []AzureMarketplaceListingAsset                   `json:"listingAssets,omitempty"`
	Plans                           []AzureMarketplacePlanResource                   `json:"plans,omitempty"`
	PriceAndAvailabilityCustomMeter *AzureMarketplacePriceAndAvailabilityCustomMeter `json:"priceAndAvailabilityCustomMeter,omitempty"`
	PriceAndAvailabilityOffer       *AzureMarketplacePriceAndAvailabilityOffer       `json:"priceAndAvailabilityOffer,omitempty"`
	Product                         *AzureMarketplaceProduct                         `json:"product,omitempty"`
	Property                        *AzureMarketplaceProperty                        `json:"property,omitempty"`
	Reseller                        *AzureMarketplaceReseller                        `json:"reseller,omitempty"`
	Setup                           *AzureCommercialMarketplaceSetup                 `json:"setup,omitempty"`
	Submission                      *AzureMarketplaceSubmission                      `json:"submission,omitempty"`
	TechnicalConfiguration          *AzureMarketplaceSaasTechnicalConfiguration      `json:"technicalConfiguration,omitempty"`
}

// NewAzureMarketplaceProductResource instantiates a new AzureMarketplaceProductResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureMarketplaceProductResource() *AzureMarketplaceProductResource {
	this := AzureMarketplaceProductResource{}
	return &this
}

// NewAzureMarketplaceProductResourceWithDefaults instantiates a new AzureMarketplaceProductResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureMarketplaceProductResourceWithDefaults() *AzureMarketplaceProductResource {
	this := AzureMarketplaceProductResource{}
	return &this
}

// GetCustomerLeads returns the CustomerLeads field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetCustomerLeads() AzureMarketplaceCustomerLeads {
	if o == nil || IsNil(o.CustomerLeads) {
		var ret AzureMarketplaceCustomerLeads
		return ret
	}
	return *o.CustomerLeads
}

// GetCustomerLeadsOk returns a tuple with the CustomerLeads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetCustomerLeadsOk() (*AzureMarketplaceCustomerLeads, bool) {
	if o == nil || IsNil(o.CustomerLeads) {
		return nil, false
	}
	return o.CustomerLeads, true
}

// HasCustomerLeads returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasCustomerLeads() bool {
	if o != nil && !IsNil(o.CustomerLeads) {
		return true
	}

	return false
}

// SetCustomerLeads gets a reference to the given AzureMarketplaceCustomerLeads and assigns it to the CustomerLeads field.
func (o *AzureMarketplaceProductResource) SetCustomerLeads(v AzureMarketplaceCustomerLeads) {
	o.CustomerLeads = &v
}

// GetListing returns the Listing field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetListing() AzureMarketplaceListing {
	if o == nil || IsNil(o.Listing) {
		var ret AzureMarketplaceListing
		return ret
	}
	return *o.Listing
}

// GetListingOk returns a tuple with the Listing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetListingOk() (*AzureMarketplaceListing, bool) {
	if o == nil || IsNil(o.Listing) {
		return nil, false
	}
	return o.Listing, true
}

// HasListing returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasListing() bool {
	if o != nil && !IsNil(o.Listing) {
		return true
	}

	return false
}

// SetListing gets a reference to the given AzureMarketplaceListing and assigns it to the Listing field.
func (o *AzureMarketplaceProductResource) SetListing(v AzureMarketplaceListing) {
	o.Listing = &v
}

// GetListingAssets returns the ListingAssets field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetListingAssets() []AzureMarketplaceListingAsset {
	if o == nil || IsNil(o.ListingAssets) {
		var ret []AzureMarketplaceListingAsset
		return ret
	}
	return o.ListingAssets
}

// GetListingAssetsOk returns a tuple with the ListingAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetListingAssetsOk() ([]AzureMarketplaceListingAsset, bool) {
	if o == nil || IsNil(o.ListingAssets) {
		return nil, false
	}
	return o.ListingAssets, true
}

// HasListingAssets returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasListingAssets() bool {
	if o != nil && !IsNil(o.ListingAssets) {
		return true
	}

	return false
}

// SetListingAssets gets a reference to the given []AzureMarketplaceListingAsset and assigns it to the ListingAssets field.
func (o *AzureMarketplaceProductResource) SetListingAssets(v []AzureMarketplaceListingAsset) {
	o.ListingAssets = v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetPlans() []AzureMarketplacePlanResource {
	if o == nil || IsNil(o.Plans) {
		var ret []AzureMarketplacePlanResource
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetPlansOk() ([]AzureMarketplacePlanResource, bool) {
	if o == nil || IsNil(o.Plans) {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasPlans() bool {
	if o != nil && !IsNil(o.Plans) {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []AzureMarketplacePlanResource and assigns it to the Plans field.
func (o *AzureMarketplaceProductResource) SetPlans(v []AzureMarketplacePlanResource) {
	o.Plans = v
}

// GetPriceAndAvailabilityCustomMeter returns the PriceAndAvailabilityCustomMeter field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityCustomMeter() AzureMarketplacePriceAndAvailabilityCustomMeter {
	if o == nil || IsNil(o.PriceAndAvailabilityCustomMeter) {
		var ret AzureMarketplacePriceAndAvailabilityCustomMeter
		return ret
	}
	return *o.PriceAndAvailabilityCustomMeter
}

// GetPriceAndAvailabilityCustomMeterOk returns a tuple with the PriceAndAvailabilityCustomMeter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityCustomMeterOk() (*AzureMarketplacePriceAndAvailabilityCustomMeter, bool) {
	if o == nil || IsNil(o.PriceAndAvailabilityCustomMeter) {
		return nil, false
	}
	return o.PriceAndAvailabilityCustomMeter, true
}

// HasPriceAndAvailabilityCustomMeter returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasPriceAndAvailabilityCustomMeter() bool {
	if o != nil && !IsNil(o.PriceAndAvailabilityCustomMeter) {
		return true
	}

	return false
}

// SetPriceAndAvailabilityCustomMeter gets a reference to the given AzureMarketplacePriceAndAvailabilityCustomMeter and assigns it to the PriceAndAvailabilityCustomMeter field.
func (o *AzureMarketplaceProductResource) SetPriceAndAvailabilityCustomMeter(v AzureMarketplacePriceAndAvailabilityCustomMeter) {
	o.PriceAndAvailabilityCustomMeter = &v
}

// GetPriceAndAvailabilityOffer returns the PriceAndAvailabilityOffer field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityOffer() AzureMarketplacePriceAndAvailabilityOffer {
	if o == nil || IsNil(o.PriceAndAvailabilityOffer) {
		var ret AzureMarketplacePriceAndAvailabilityOffer
		return ret
	}
	return *o.PriceAndAvailabilityOffer
}

// GetPriceAndAvailabilityOfferOk returns a tuple with the PriceAndAvailabilityOffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetPriceAndAvailabilityOfferOk() (*AzureMarketplacePriceAndAvailabilityOffer, bool) {
	if o == nil || IsNil(o.PriceAndAvailabilityOffer) {
		return nil, false
	}
	return o.PriceAndAvailabilityOffer, true
}

// HasPriceAndAvailabilityOffer returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasPriceAndAvailabilityOffer() bool {
	if o != nil && !IsNil(o.PriceAndAvailabilityOffer) {
		return true
	}

	return false
}

// SetPriceAndAvailabilityOffer gets a reference to the given AzureMarketplacePriceAndAvailabilityOffer and assigns it to the PriceAndAvailabilityOffer field.
func (o *AzureMarketplaceProductResource) SetPriceAndAvailabilityOffer(v AzureMarketplacePriceAndAvailabilityOffer) {
	o.PriceAndAvailabilityOffer = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetProduct() AzureMarketplaceProduct {
	if o == nil || IsNil(o.Product) {
		var ret AzureMarketplaceProduct
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetProductOk() (*AzureMarketplaceProduct, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given AzureMarketplaceProduct and assigns it to the Product field.
func (o *AzureMarketplaceProductResource) SetProduct(v AzureMarketplaceProduct) {
	o.Product = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetProperty() AzureMarketplaceProperty {
	if o == nil || IsNil(o.Property) {
		var ret AzureMarketplaceProperty
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetPropertyOk() (*AzureMarketplaceProperty, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given AzureMarketplaceProperty and assigns it to the Property field.
func (o *AzureMarketplaceProductResource) SetProperty(v AzureMarketplaceProperty) {
	o.Property = &v
}

// GetReseller returns the Reseller field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetReseller() AzureMarketplaceReseller {
	if o == nil || IsNil(o.Reseller) {
		var ret AzureMarketplaceReseller
		return ret
	}
	return *o.Reseller
}

// GetResellerOk returns a tuple with the Reseller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetResellerOk() (*AzureMarketplaceReseller, bool) {
	if o == nil || IsNil(o.Reseller) {
		return nil, false
	}
	return o.Reseller, true
}

// HasReseller returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasReseller() bool {
	if o != nil && !IsNil(o.Reseller) {
		return true
	}

	return false
}

// SetReseller gets a reference to the given AzureMarketplaceReseller and assigns it to the Reseller field.
func (o *AzureMarketplaceProductResource) SetReseller(v AzureMarketplaceReseller) {
	o.Reseller = &v
}

// GetSetup returns the Setup field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetSetup() AzureCommercialMarketplaceSetup {
	if o == nil || IsNil(o.Setup) {
		var ret AzureCommercialMarketplaceSetup
		return ret
	}
	return *o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetSetupOk() (*AzureCommercialMarketplaceSetup, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasSetup() bool {
	if o != nil && !IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given AzureCommercialMarketplaceSetup and assigns it to the Setup field.
func (o *AzureMarketplaceProductResource) SetSetup(v AzureCommercialMarketplaceSetup) {
	o.Setup = &v
}

// GetSubmission returns the Submission field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetSubmission() AzureMarketplaceSubmission {
	if o == nil || IsNil(o.Submission) {
		var ret AzureMarketplaceSubmission
		return ret
	}
	return *o.Submission
}

// GetSubmissionOk returns a tuple with the Submission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetSubmissionOk() (*AzureMarketplaceSubmission, bool) {
	if o == nil || IsNil(o.Submission) {
		return nil, false
	}
	return o.Submission, true
}

// HasSubmission returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasSubmission() bool {
	if o != nil && !IsNil(o.Submission) {
		return true
	}

	return false
}

// SetSubmission gets a reference to the given AzureMarketplaceSubmission and assigns it to the Submission field.
func (o *AzureMarketplaceProductResource) SetSubmission(v AzureMarketplaceSubmission) {
	o.Submission = &v
}

// GetTechnicalConfiguration returns the TechnicalConfiguration field value if set, zero value otherwise.
func (o *AzureMarketplaceProductResource) GetTechnicalConfiguration() AzureMarketplaceSaasTechnicalConfiguration {
	if o == nil || IsNil(o.TechnicalConfiguration) {
		var ret AzureMarketplaceSaasTechnicalConfiguration
		return ret
	}
	return *o.TechnicalConfiguration
}

// GetTechnicalConfigurationOk returns a tuple with the TechnicalConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureMarketplaceProductResource) GetTechnicalConfigurationOk() (*AzureMarketplaceSaasTechnicalConfiguration, bool) {
	if o == nil || IsNil(o.TechnicalConfiguration) {
		return nil, false
	}
	return o.TechnicalConfiguration, true
}

// HasTechnicalConfiguration returns a boolean if a field has been set.
func (o *AzureMarketplaceProductResource) HasTechnicalConfiguration() bool {
	if o != nil && !IsNil(o.TechnicalConfiguration) {
		return true
	}

	return false
}

// SetTechnicalConfiguration gets a reference to the given AzureMarketplaceSaasTechnicalConfiguration and assigns it to the TechnicalConfiguration field.
func (o *AzureMarketplaceProductResource) SetTechnicalConfiguration(v AzureMarketplaceSaasTechnicalConfiguration) {
	o.TechnicalConfiguration = &v
}

func (o AzureMarketplaceProductResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureMarketplaceProductResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerLeads) {
		toSerialize["customerLeads"] = o.CustomerLeads
	}
	if !IsNil(o.Listing) {
		toSerialize["listing"] = o.Listing
	}
	if !IsNil(o.ListingAssets) {
		toSerialize["listingAssets"] = o.ListingAssets
	}
	if !IsNil(o.Plans) {
		toSerialize["plans"] = o.Plans
	}
	if !IsNil(o.PriceAndAvailabilityCustomMeter) {
		toSerialize["priceAndAvailabilityCustomMeter"] = o.PriceAndAvailabilityCustomMeter
	}
	if !IsNil(o.PriceAndAvailabilityOffer) {
		toSerialize["priceAndAvailabilityOffer"] = o.PriceAndAvailabilityOffer
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Reseller) {
		toSerialize["reseller"] = o.Reseller
	}
	if !IsNil(o.Setup) {
		toSerialize["setup"] = o.Setup
	}
	if !IsNil(o.Submission) {
		toSerialize["submission"] = o.Submission
	}
	if !IsNil(o.TechnicalConfiguration) {
		toSerialize["technicalConfiguration"] = o.TechnicalConfiguration
	}
	return toSerialize, nil
}

type NullableAzureMarketplaceProductResource struct {
	value *AzureMarketplaceProductResource
	isSet bool
}

func (v NullableAzureMarketplaceProductResource) Get() *AzureMarketplaceProductResource {
	return v.value
}

func (v *NullableAzureMarketplaceProductResource) Set(val *AzureMarketplaceProductResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureMarketplaceProductResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureMarketplaceProductResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureMarketplaceProductResource(val *AzureMarketplaceProductResource) *NullableAzureMarketplaceProductResource {
	return &NullableAzureMarketplaceProductResource{value: val, isSet: true}
}

func (v NullableAzureMarketplaceProductResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureMarketplaceProductResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}