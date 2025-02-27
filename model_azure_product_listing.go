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

// checks if the AzureProductListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureProductListing{}

// AzureProductListing struct for AzureProductListing
type AzureProductListing struct {
	AccessInformation *string `json:"accessInformation,omitempty"`
	// Not original fields. They are populated by other API calls
	Assets                     []AzureProductListingAsset `json:"assets,omitempty"`
	CompatibleProducts         []string                   `json:"compatibleProducts,omitempty"`
	Description                *string                    `json:"description,omitempty"`
	GettingStartedInstructions *string                    `json:"gettingStartedInstructions,omitempty"`
	Id                         *string                    `json:"id,omitempty"`
	Keywords                   []string                   `json:"keywords,omitempty"`
	LanguageCode               *string                    `json:"languageCode,omitempty"`
	ListingContacts            []AzureListingContact      `json:"listingContacts,omitempty"`
	ListingUris                []AzureListingUri          `json:"listingUris,omitempty"`
	ProductDisplayName         *string                    `json:"productDisplayName,omitempty"`
	PublisherName              *string                    `json:"publisherName,omitempty"`
	ResourceType               *string                    `json:"resourceType,omitempty"`
	ShortDescription           *string                    `json:"shortDescription,omitempty"`
	Summary                    *string                    `json:"summary,omitempty"`
	Title                      *string                    `json:"title,omitempty"`
}

// NewAzureProductListing instantiates a new AzureProductListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureProductListing() *AzureProductListing {
	this := AzureProductListing{}
	return &this
}

// NewAzureProductListingWithDefaults instantiates a new AzureProductListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureProductListingWithDefaults() *AzureProductListing {
	this := AzureProductListing{}
	return &this
}

// GetAccessInformation returns the AccessInformation field value if set, zero value otherwise.
func (o *AzureProductListing) GetAccessInformation() string {
	if o == nil || IsNil(o.AccessInformation) {
		var ret string
		return ret
	}
	return *o.AccessInformation
}

// GetAccessInformationOk returns a tuple with the AccessInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetAccessInformationOk() (*string, bool) {
	if o == nil || IsNil(o.AccessInformation) {
		return nil, false
	}
	return o.AccessInformation, true
}

// HasAccessInformation returns a boolean if a field has been set.
func (o *AzureProductListing) HasAccessInformation() bool {
	if o != nil && !IsNil(o.AccessInformation) {
		return true
	}

	return false
}

// SetAccessInformation gets a reference to the given string and assigns it to the AccessInformation field.
func (o *AzureProductListing) SetAccessInformation(v string) {
	o.AccessInformation = &v
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *AzureProductListing) GetAssets() []AzureProductListingAsset {
	if o == nil || IsNil(o.Assets) {
		var ret []AzureProductListingAsset
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetAssetsOk() ([]AzureProductListingAsset, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *AzureProductListing) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AzureProductListingAsset and assigns it to the Assets field.
func (o *AzureProductListing) SetAssets(v []AzureProductListingAsset) {
	o.Assets = v
}

// GetCompatibleProducts returns the CompatibleProducts field value if set, zero value otherwise.
func (o *AzureProductListing) GetCompatibleProducts() []string {
	if o == nil || IsNil(o.CompatibleProducts) {
		var ret []string
		return ret
	}
	return o.CompatibleProducts
}

// GetCompatibleProductsOk returns a tuple with the CompatibleProducts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetCompatibleProductsOk() ([]string, bool) {
	if o == nil || IsNil(o.CompatibleProducts) {
		return nil, false
	}
	return o.CompatibleProducts, true
}

// HasCompatibleProducts returns a boolean if a field has been set.
func (o *AzureProductListing) HasCompatibleProducts() bool {
	if o != nil && !IsNil(o.CompatibleProducts) {
		return true
	}

	return false
}

// SetCompatibleProducts gets a reference to the given []string and assigns it to the CompatibleProducts field.
func (o *AzureProductListing) SetCompatibleProducts(v []string) {
	o.CompatibleProducts = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AzureProductListing) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AzureProductListing) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AzureProductListing) SetDescription(v string) {
	o.Description = &v
}

// GetGettingStartedInstructions returns the GettingStartedInstructions field value if set, zero value otherwise.
func (o *AzureProductListing) GetGettingStartedInstructions() string {
	if o == nil || IsNil(o.GettingStartedInstructions) {
		var ret string
		return ret
	}
	return *o.GettingStartedInstructions
}

// GetGettingStartedInstructionsOk returns a tuple with the GettingStartedInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetGettingStartedInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.GettingStartedInstructions) {
		return nil, false
	}
	return o.GettingStartedInstructions, true
}

// HasGettingStartedInstructions returns a boolean if a field has been set.
func (o *AzureProductListing) HasGettingStartedInstructions() bool {
	if o != nil && !IsNil(o.GettingStartedInstructions) {
		return true
	}

	return false
}

// SetGettingStartedInstructions gets a reference to the given string and assigns it to the GettingStartedInstructions field.
func (o *AzureProductListing) SetGettingStartedInstructions(v string) {
	o.GettingStartedInstructions = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureProductListing) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureProductListing) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureProductListing) SetId(v string) {
	o.Id = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *AzureProductListing) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *AzureProductListing) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *AzureProductListing) SetKeywords(v []string) {
	o.Keywords = v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *AzureProductListing) GetLanguageCode() string {
	if o == nil || IsNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetLanguageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCode) {
		return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *AzureProductListing) HasLanguageCode() bool {
	if o != nil && !IsNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *AzureProductListing) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetListingContacts returns the ListingContacts field value if set, zero value otherwise.
func (o *AzureProductListing) GetListingContacts() []AzureListingContact {
	if o == nil || IsNil(o.ListingContacts) {
		var ret []AzureListingContact
		return ret
	}
	return o.ListingContacts
}

// GetListingContactsOk returns a tuple with the ListingContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetListingContactsOk() ([]AzureListingContact, bool) {
	if o == nil || IsNil(o.ListingContacts) {
		return nil, false
	}
	return o.ListingContacts, true
}

// HasListingContacts returns a boolean if a field has been set.
func (o *AzureProductListing) HasListingContacts() bool {
	if o != nil && !IsNil(o.ListingContacts) {
		return true
	}

	return false
}

// SetListingContacts gets a reference to the given []AzureListingContact and assigns it to the ListingContacts field.
func (o *AzureProductListing) SetListingContacts(v []AzureListingContact) {
	o.ListingContacts = v
}

// GetListingUris returns the ListingUris field value if set, zero value otherwise.
func (o *AzureProductListing) GetListingUris() []AzureListingUri {
	if o == nil || IsNil(o.ListingUris) {
		var ret []AzureListingUri
		return ret
	}
	return o.ListingUris
}

// GetListingUrisOk returns a tuple with the ListingUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetListingUrisOk() ([]AzureListingUri, bool) {
	if o == nil || IsNil(o.ListingUris) {
		return nil, false
	}
	return o.ListingUris, true
}

// HasListingUris returns a boolean if a field has been set.
func (o *AzureProductListing) HasListingUris() bool {
	if o != nil && !IsNil(o.ListingUris) {
		return true
	}

	return false
}

// SetListingUris gets a reference to the given []AzureListingUri and assigns it to the ListingUris field.
func (o *AzureProductListing) SetListingUris(v []AzureListingUri) {
	o.ListingUris = v
}

// GetProductDisplayName returns the ProductDisplayName field value if set, zero value otherwise.
func (o *AzureProductListing) GetProductDisplayName() string {
	if o == nil || IsNil(o.ProductDisplayName) {
		var ret string
		return ret
	}
	return *o.ProductDisplayName
}

// GetProductDisplayNameOk returns a tuple with the ProductDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetProductDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDisplayName) {
		return nil, false
	}
	return o.ProductDisplayName, true
}

// HasProductDisplayName returns a boolean if a field has been set.
func (o *AzureProductListing) HasProductDisplayName() bool {
	if o != nil && !IsNil(o.ProductDisplayName) {
		return true
	}

	return false
}

// SetProductDisplayName gets a reference to the given string and assigns it to the ProductDisplayName field.
func (o *AzureProductListing) SetProductDisplayName(v string) {
	o.ProductDisplayName = &v
}

// GetPublisherName returns the PublisherName field value if set, zero value otherwise.
func (o *AzureProductListing) GetPublisherName() string {
	if o == nil || IsNil(o.PublisherName) {
		var ret string
		return ret
	}
	return *o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetPublisherNameOk() (*string, bool) {
	if o == nil || IsNil(o.PublisherName) {
		return nil, false
	}
	return o.PublisherName, true
}

// HasPublisherName returns a boolean if a field has been set.
func (o *AzureProductListing) HasPublisherName() bool {
	if o != nil && !IsNil(o.PublisherName) {
		return true
	}

	return false
}

// SetPublisherName gets a reference to the given string and assigns it to the PublisherName field.
func (o *AzureProductListing) SetPublisherName(v string) {
	o.PublisherName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AzureProductListing) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AzureProductListing) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AzureProductListing) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *AzureProductListing) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *AzureProductListing) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *AzureProductListing) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AzureProductListing) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AzureProductListing) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AzureProductListing) SetSummary(v string) {
	o.Summary = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AzureProductListing) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureProductListing) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AzureProductListing) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AzureProductListing) SetTitle(v string) {
	o.Title = &v
}

func (o AzureProductListing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureProductListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessInformation) {
		toSerialize["accessInformation"] = o.AccessInformation
	}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.CompatibleProducts) {
		toSerialize["compatibleProducts"] = o.CompatibleProducts
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GettingStartedInstructions) {
		toSerialize["gettingStartedInstructions"] = o.GettingStartedInstructions
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	if !IsNil(o.LanguageCode) {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if !IsNil(o.ListingContacts) {
		toSerialize["listingContacts"] = o.ListingContacts
	}
	if !IsNil(o.ListingUris) {
		toSerialize["listingUris"] = o.ListingUris
	}
	if !IsNil(o.ProductDisplayName) {
		toSerialize["productDisplayName"] = o.ProductDisplayName
	}
	if !IsNil(o.PublisherName) {
		toSerialize["publisherName"] = o.PublisherName
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableAzureProductListing struct {
	value *AzureProductListing
	isSet bool
}

func (v NullableAzureProductListing) Get() *AzureProductListing {
	return v.value
}

func (v *NullableAzureProductListing) Set(val *AzureProductListing) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureProductListing) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureProductListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureProductListing(val *AzureProductListing) *NullableAzureProductListing {
	return &NullableAzureProductListing{value: val, isSet: true}
}

func (v NullableAzureProductListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureProductListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
