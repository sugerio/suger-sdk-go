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

// checks if the GcpMarketplaceProductMarketingSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GcpMarketplaceProductMarketingSpec{}

// GcpMarketplaceProductMarketingSpec struct for GcpMarketplaceProductMarketingSpec
type GcpMarketplaceProductMarketingSpec struct {
	Description          *string                                  `json:"description,omitempty"`
	DisplayNames         []string                                 `json:"displayNames,omitempty"`
	DocumentationSpecs   []GcpMarketplaceProductDocumentationSpec `json:"documentationSpecs,omitempty"`
	EulaUrl              *string                                  `json:"eulaUrl,omitempty"`
	ExternalLicenseSpecs []GcpMarketplaceProductLicenseSpec       `json:"externalLicenseSpecs,omitempty"`
	ExternalMarketingUrl *string                                  `json:"externalMarketingUrl,omitempty"`
	// in format of \"base64://...\"
	Icon              *string                           `json:"icon,omitempty"`
	SearchCategories  []string                          `json:"searchCategories,omitempty"`
	SearchDescription *string                           `json:"searchDescription,omitempty"`
	SearchKeywords    []string                          `json:"searchKeywords,omitempty"`
	SignupUri         *string                           `json:"signupUri,omitempty"`
	SupportSpec       *GcpMarketplaceProductSupportSpec `json:"supportSpec,omitempty"`
	TagLine           *string                           `json:"tagLine,omitempty"`
	Title             *string                           `json:"title,omitempty"`
}

// NewGcpMarketplaceProductMarketingSpec instantiates a new GcpMarketplaceProductMarketingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGcpMarketplaceProductMarketingSpec() *GcpMarketplaceProductMarketingSpec {
	this := GcpMarketplaceProductMarketingSpec{}
	return &this
}

// NewGcpMarketplaceProductMarketingSpecWithDefaults instantiates a new GcpMarketplaceProductMarketingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGcpMarketplaceProductMarketingSpecWithDefaults() *GcpMarketplaceProductMarketingSpec {
	this := GcpMarketplaceProductMarketingSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GcpMarketplaceProductMarketingSpec) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayNames returns the DisplayNames field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetDisplayNames() []string {
	if o == nil || IsNil(o.DisplayNames) {
		var ret []string
		return ret
	}
	return o.DisplayNames
}

// GetDisplayNamesOk returns a tuple with the DisplayNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetDisplayNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisplayNames) {
		return nil, false
	}
	return o.DisplayNames, true
}

// HasDisplayNames returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasDisplayNames() bool {
	if o != nil && !IsNil(o.DisplayNames) {
		return true
	}

	return false
}

// SetDisplayNames gets a reference to the given []string and assigns it to the DisplayNames field.
func (o *GcpMarketplaceProductMarketingSpec) SetDisplayNames(v []string) {
	o.DisplayNames = v
}

// GetDocumentationSpecs returns the DocumentationSpecs field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetDocumentationSpecs() []GcpMarketplaceProductDocumentationSpec {
	if o == nil || IsNil(o.DocumentationSpecs) {
		var ret []GcpMarketplaceProductDocumentationSpec
		return ret
	}
	return o.DocumentationSpecs
}

// GetDocumentationSpecsOk returns a tuple with the DocumentationSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetDocumentationSpecsOk() ([]GcpMarketplaceProductDocumentationSpec, bool) {
	if o == nil || IsNil(o.DocumentationSpecs) {
		return nil, false
	}
	return o.DocumentationSpecs, true
}

// HasDocumentationSpecs returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasDocumentationSpecs() bool {
	if o != nil && !IsNil(o.DocumentationSpecs) {
		return true
	}

	return false
}

// SetDocumentationSpecs gets a reference to the given []GcpMarketplaceProductDocumentationSpec and assigns it to the DocumentationSpecs field.
func (o *GcpMarketplaceProductMarketingSpec) SetDocumentationSpecs(v []GcpMarketplaceProductDocumentationSpec) {
	o.DocumentationSpecs = v
}

// GetEulaUrl returns the EulaUrl field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetEulaUrl() string {
	if o == nil || IsNil(o.EulaUrl) {
		var ret string
		return ret
	}
	return *o.EulaUrl
}

// GetEulaUrlOk returns a tuple with the EulaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetEulaUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EulaUrl) {
		return nil, false
	}
	return o.EulaUrl, true
}

// HasEulaUrl returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasEulaUrl() bool {
	if o != nil && !IsNil(o.EulaUrl) {
		return true
	}

	return false
}

// SetEulaUrl gets a reference to the given string and assigns it to the EulaUrl field.
func (o *GcpMarketplaceProductMarketingSpec) SetEulaUrl(v string) {
	o.EulaUrl = &v
}

// GetExternalLicenseSpecs returns the ExternalLicenseSpecs field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetExternalLicenseSpecs() []GcpMarketplaceProductLicenseSpec {
	if o == nil || IsNil(o.ExternalLicenseSpecs) {
		var ret []GcpMarketplaceProductLicenseSpec
		return ret
	}
	return o.ExternalLicenseSpecs
}

// GetExternalLicenseSpecsOk returns a tuple with the ExternalLicenseSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetExternalLicenseSpecsOk() ([]GcpMarketplaceProductLicenseSpec, bool) {
	if o == nil || IsNil(o.ExternalLicenseSpecs) {
		return nil, false
	}
	return o.ExternalLicenseSpecs, true
}

// HasExternalLicenseSpecs returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasExternalLicenseSpecs() bool {
	if o != nil && !IsNil(o.ExternalLicenseSpecs) {
		return true
	}

	return false
}

// SetExternalLicenseSpecs gets a reference to the given []GcpMarketplaceProductLicenseSpec and assigns it to the ExternalLicenseSpecs field.
func (o *GcpMarketplaceProductMarketingSpec) SetExternalLicenseSpecs(v []GcpMarketplaceProductLicenseSpec) {
	o.ExternalLicenseSpecs = v
}

// GetExternalMarketingUrl returns the ExternalMarketingUrl field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetExternalMarketingUrl() string {
	if o == nil || IsNil(o.ExternalMarketingUrl) {
		var ret string
		return ret
	}
	return *o.ExternalMarketingUrl
}

// GetExternalMarketingUrlOk returns a tuple with the ExternalMarketingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetExternalMarketingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalMarketingUrl) {
		return nil, false
	}
	return o.ExternalMarketingUrl, true
}

// HasExternalMarketingUrl returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasExternalMarketingUrl() bool {
	if o != nil && !IsNil(o.ExternalMarketingUrl) {
		return true
	}

	return false
}

// SetExternalMarketingUrl gets a reference to the given string and assigns it to the ExternalMarketingUrl field.
func (o *GcpMarketplaceProductMarketingSpec) SetExternalMarketingUrl(v string) {
	o.ExternalMarketingUrl = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *GcpMarketplaceProductMarketingSpec) SetIcon(v string) {
	o.Icon = &v
}

// GetSearchCategories returns the SearchCategories field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetSearchCategories() []string {
	if o == nil || IsNil(o.SearchCategories) {
		var ret []string
		return ret
	}
	return o.SearchCategories
}

// GetSearchCategoriesOk returns a tuple with the SearchCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetSearchCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchCategories) {
		return nil, false
	}
	return o.SearchCategories, true
}

// HasSearchCategories returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasSearchCategories() bool {
	if o != nil && !IsNil(o.SearchCategories) {
		return true
	}

	return false
}

// SetSearchCategories gets a reference to the given []string and assigns it to the SearchCategories field.
func (o *GcpMarketplaceProductMarketingSpec) SetSearchCategories(v []string) {
	o.SearchCategories = v
}

// GetSearchDescription returns the SearchDescription field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetSearchDescription() string {
	if o == nil || IsNil(o.SearchDescription) {
		var ret string
		return ret
	}
	return *o.SearchDescription
}

// GetSearchDescriptionOk returns a tuple with the SearchDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetSearchDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SearchDescription) {
		return nil, false
	}
	return o.SearchDescription, true
}

// HasSearchDescription returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasSearchDescription() bool {
	if o != nil && !IsNil(o.SearchDescription) {
		return true
	}

	return false
}

// SetSearchDescription gets a reference to the given string and assigns it to the SearchDescription field.
func (o *GcpMarketplaceProductMarketingSpec) SetSearchDescription(v string) {
	o.SearchDescription = &v
}

// GetSearchKeywords returns the SearchKeywords field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetSearchKeywords() []string {
	if o == nil || IsNil(o.SearchKeywords) {
		var ret []string
		return ret
	}
	return o.SearchKeywords
}

// GetSearchKeywordsOk returns a tuple with the SearchKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetSearchKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchKeywords) {
		return nil, false
	}
	return o.SearchKeywords, true
}

// HasSearchKeywords returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasSearchKeywords() bool {
	if o != nil && !IsNil(o.SearchKeywords) {
		return true
	}

	return false
}

// SetSearchKeywords gets a reference to the given []string and assigns it to the SearchKeywords field.
func (o *GcpMarketplaceProductMarketingSpec) SetSearchKeywords(v []string) {
	o.SearchKeywords = v
}

// GetSignupUri returns the SignupUri field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetSignupUri() string {
	if o == nil || IsNil(o.SignupUri) {
		var ret string
		return ret
	}
	return *o.SignupUri
}

// GetSignupUriOk returns a tuple with the SignupUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetSignupUriOk() (*string, bool) {
	if o == nil || IsNil(o.SignupUri) {
		return nil, false
	}
	return o.SignupUri, true
}

// HasSignupUri returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasSignupUri() bool {
	if o != nil && !IsNil(o.SignupUri) {
		return true
	}

	return false
}

// SetSignupUri gets a reference to the given string and assigns it to the SignupUri field.
func (o *GcpMarketplaceProductMarketingSpec) SetSignupUri(v string) {
	o.SignupUri = &v
}

// GetSupportSpec returns the SupportSpec field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetSupportSpec() GcpMarketplaceProductSupportSpec {
	if o == nil || IsNil(o.SupportSpec) {
		var ret GcpMarketplaceProductSupportSpec
		return ret
	}
	return *o.SupportSpec
}

// GetSupportSpecOk returns a tuple with the SupportSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetSupportSpecOk() (*GcpMarketplaceProductSupportSpec, bool) {
	if o == nil || IsNil(o.SupportSpec) {
		return nil, false
	}
	return o.SupportSpec, true
}

// HasSupportSpec returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasSupportSpec() bool {
	if o != nil && !IsNil(o.SupportSpec) {
		return true
	}

	return false
}

// SetSupportSpec gets a reference to the given GcpMarketplaceProductSupportSpec and assigns it to the SupportSpec field.
func (o *GcpMarketplaceProductMarketingSpec) SetSupportSpec(v GcpMarketplaceProductSupportSpec) {
	o.SupportSpec = &v
}

// GetTagLine returns the TagLine field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetTagLine() string {
	if o == nil || IsNil(o.TagLine) {
		var ret string
		return ret
	}
	return *o.TagLine
}

// GetTagLineOk returns a tuple with the TagLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetTagLineOk() (*string, bool) {
	if o == nil || IsNil(o.TagLine) {
		return nil, false
	}
	return o.TagLine, true
}

// HasTagLine returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasTagLine() bool {
	if o != nil && !IsNil(o.TagLine) {
		return true
	}

	return false
}

// SetTagLine gets a reference to the given string and assigns it to the TagLine field.
func (o *GcpMarketplaceProductMarketingSpec) SetTagLine(v string) {
	o.TagLine = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GcpMarketplaceProductMarketingSpec) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GcpMarketplaceProductMarketingSpec) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GcpMarketplaceProductMarketingSpec) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GcpMarketplaceProductMarketingSpec) SetTitle(v string) {
	o.Title = &v
}

func (o GcpMarketplaceProductMarketingSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GcpMarketplaceProductMarketingSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayNames) {
		toSerialize["displayNames"] = o.DisplayNames
	}
	if !IsNil(o.DocumentationSpecs) {
		toSerialize["documentationSpecs"] = o.DocumentationSpecs
	}
	if !IsNil(o.EulaUrl) {
		toSerialize["eulaUrl"] = o.EulaUrl
	}
	if !IsNil(o.ExternalLicenseSpecs) {
		toSerialize["externalLicenseSpecs"] = o.ExternalLicenseSpecs
	}
	if !IsNil(o.ExternalMarketingUrl) {
		toSerialize["externalMarketingUrl"] = o.ExternalMarketingUrl
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.SearchCategories) {
		toSerialize["searchCategories"] = o.SearchCategories
	}
	if !IsNil(o.SearchDescription) {
		toSerialize["searchDescription"] = o.SearchDescription
	}
	if !IsNil(o.SearchKeywords) {
		toSerialize["searchKeywords"] = o.SearchKeywords
	}
	if !IsNil(o.SignupUri) {
		toSerialize["signupUri"] = o.SignupUri
	}
	if !IsNil(o.SupportSpec) {
		toSerialize["supportSpec"] = o.SupportSpec
	}
	if !IsNil(o.TagLine) {
		toSerialize["tagLine"] = o.TagLine
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGcpMarketplaceProductMarketingSpec struct {
	value *GcpMarketplaceProductMarketingSpec
	isSet bool
}

func (v NullableGcpMarketplaceProductMarketingSpec) Get() *GcpMarketplaceProductMarketingSpec {
	return v.value
}

func (v *NullableGcpMarketplaceProductMarketingSpec) Set(val *GcpMarketplaceProductMarketingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGcpMarketplaceProductMarketingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGcpMarketplaceProductMarketingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGcpMarketplaceProductMarketingSpec(val *GcpMarketplaceProductMarketingSpec) *NullableGcpMarketplaceProductMarketingSpec {
	return &NullableGcpMarketplaceProductMarketingSpec{value: val, isSet: true}
}

func (v NullableGcpMarketplaceProductMarketingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGcpMarketplaceProductMarketingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
