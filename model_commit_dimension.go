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

// checks if the CommitDimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitDimension{}

// CommitDimension The commit dimension. There may be one or more commit dimensions defined in single product, offer or entitlement.
type CommitDimension struct {
	Category    *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	// Whether this commit dimension is newly created by user, when creating AWS Marketplace Contract private offer.
	IsUserCreated *bool `json:"isUserCreated,omitempty"`
	// API name of the dimension
	Key *string `json:"key,omitempty"`
	// The term length for the commit amount, such as 6 months, or 1 year. The length is used together with timeUnit. If the length is zero, use the TermEndTime.
	Length *int32 `json:"length,omitempty"`
	// The maximum number of users for PER_USER commit
	MaximumUsers *int32 `json:"maximumUsers,omitempty"`
	// The minimum number of users for PER_USER commit
	MinimumUsers *int32 `json:"minimumUsers,omitempty"`
	// Display name of the dimension
	Name *string `json:"name,omitempty"`
	// The quantity of this commit.
	Quantity *int32 `json:"quantity,omitempty"`
	// The commit amount. For GCP, it is monthly commitment.
	Rate *float32 `json:"rate,omitempty"`
	// The term of the commit amount. It is used for front-end display only.
	Term *string `json:"term,omitempty"`
	// The end time of the commit term.
	TermEndTime *string `json:"termEndTime,omitempty"`
	// The term unit for the commit amount.
	TimeUnit *TimeUnit `json:"timeUnit,omitempty"`
	// The type of the commit dimension. Applicable only to Azure Marketplace.
	Type *CommitDimensionType `json:"type,omitempty"`
	// These indicate whether the dimension covers metering, entitlement, or support for external metering
	Types []string `json:"types,omitempty"`
}

// NewCommitDimension instantiates a new CommitDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitDimension() *CommitDimension {
	this := CommitDimension{}
	return &this
}

// NewCommitDimensionWithDefaults instantiates a new CommitDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDimensionWithDefaults() *CommitDimension {
	this := CommitDimension{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CommitDimension) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CommitDimension) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CommitDimension) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommitDimension) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommitDimension) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommitDimension) SetDescription(v string) {
	o.Description = &v
}

// GetIsUserCreated returns the IsUserCreated field value if set, zero value otherwise.
func (o *CommitDimension) GetIsUserCreated() bool {
	if o == nil || IsNil(o.IsUserCreated) {
		var ret bool
		return ret
	}
	return *o.IsUserCreated
}

// GetIsUserCreatedOk returns a tuple with the IsUserCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetIsUserCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUserCreated) {
		return nil, false
	}
	return o.IsUserCreated, true
}

// HasIsUserCreated returns a boolean if a field has been set.
func (o *CommitDimension) HasIsUserCreated() bool {
	if o != nil && !IsNil(o.IsUserCreated) {
		return true
	}

	return false
}

// SetIsUserCreated gets a reference to the given bool and assigns it to the IsUserCreated field.
func (o *CommitDimension) SetIsUserCreated(v bool) {
	o.IsUserCreated = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CommitDimension) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CommitDimension) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CommitDimension) SetKey(v string) {
	o.Key = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *CommitDimension) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *CommitDimension) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *CommitDimension) SetLength(v int32) {
	o.Length = &v
}

// GetMaximumUsers returns the MaximumUsers field value if set, zero value otherwise.
func (o *CommitDimension) GetMaximumUsers() int32 {
	if o == nil || IsNil(o.MaximumUsers) {
		var ret int32
		return ret
	}
	return *o.MaximumUsers
}

// GetMaximumUsersOk returns a tuple with the MaximumUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetMaximumUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumUsers) {
		return nil, false
	}
	return o.MaximumUsers, true
}

// HasMaximumUsers returns a boolean if a field has been set.
func (o *CommitDimension) HasMaximumUsers() bool {
	if o != nil && !IsNil(o.MaximumUsers) {
		return true
	}

	return false
}

// SetMaximumUsers gets a reference to the given int32 and assigns it to the MaximumUsers field.
func (o *CommitDimension) SetMaximumUsers(v int32) {
	o.MaximumUsers = &v
}

// GetMinimumUsers returns the MinimumUsers field value if set, zero value otherwise.
func (o *CommitDimension) GetMinimumUsers() int32 {
	if o == nil || IsNil(o.MinimumUsers) {
		var ret int32
		return ret
	}
	return *o.MinimumUsers
}

// GetMinimumUsersOk returns a tuple with the MinimumUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetMinimumUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumUsers) {
		return nil, false
	}
	return o.MinimumUsers, true
}

// HasMinimumUsers returns a boolean if a field has been set.
func (o *CommitDimension) HasMinimumUsers() bool {
	if o != nil && !IsNil(o.MinimumUsers) {
		return true
	}

	return false
}

// SetMinimumUsers gets a reference to the given int32 and assigns it to the MinimumUsers field.
func (o *CommitDimension) SetMinimumUsers(v int32) {
	o.MinimumUsers = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommitDimension) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommitDimension) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommitDimension) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CommitDimension) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CommitDimension) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *CommitDimension) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CommitDimension) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CommitDimension) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *CommitDimension) SetRate(v float32) {
	o.Rate = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *CommitDimension) GetTerm() string {
	if o == nil || IsNil(o.Term) {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetTermOk() (*string, bool) {
	if o == nil || IsNil(o.Term) {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *CommitDimension) HasTerm() bool {
	if o != nil && !IsNil(o.Term) {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *CommitDimension) SetTerm(v string) {
	o.Term = &v
}

// GetTermEndTime returns the TermEndTime field value if set, zero value otherwise.
func (o *CommitDimension) GetTermEndTime() string {
	if o == nil || IsNil(o.TermEndTime) {
		var ret string
		return ret
	}
	return *o.TermEndTime
}

// GetTermEndTimeOk returns a tuple with the TermEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetTermEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TermEndTime) {
		return nil, false
	}
	return o.TermEndTime, true
}

// HasTermEndTime returns a boolean if a field has been set.
func (o *CommitDimension) HasTermEndTime() bool {
	if o != nil && !IsNil(o.TermEndTime) {
		return true
	}

	return false
}

// SetTermEndTime gets a reference to the given string and assigns it to the TermEndTime field.
func (o *CommitDimension) SetTermEndTime(v string) {
	o.TermEndTime = &v
}

// GetTimeUnit returns the TimeUnit field value if set, zero value otherwise.
func (o *CommitDimension) GetTimeUnit() TimeUnit {
	if o == nil || IsNil(o.TimeUnit) {
		var ret TimeUnit
		return ret
	}
	return *o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetTimeUnitOk() (*TimeUnit, bool) {
	if o == nil || IsNil(o.TimeUnit) {
		return nil, false
	}
	return o.TimeUnit, true
}

// HasTimeUnit returns a boolean if a field has been set.
func (o *CommitDimension) HasTimeUnit() bool {
	if o != nil && !IsNil(o.TimeUnit) {
		return true
	}

	return false
}

// SetTimeUnit gets a reference to the given TimeUnit and assigns it to the TimeUnit field.
func (o *CommitDimension) SetTimeUnit(v TimeUnit) {
	o.TimeUnit = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CommitDimension) GetType() CommitDimensionType {
	if o == nil || IsNil(o.Type) {
		var ret CommitDimensionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetTypeOk() (*CommitDimensionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CommitDimension) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CommitDimensionType and assigns it to the Type field.
func (o *CommitDimension) SetType(v CommitDimensionType) {
	o.Type = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *CommitDimension) GetTypes() []string {
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDimension) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *CommitDimension) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *CommitDimension) SetTypes(v []string) {
	o.Types = v
}

func (o CommitDimension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitDimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsUserCreated) {
		toSerialize["isUserCreated"] = o.IsUserCreated
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.MaximumUsers) {
		toSerialize["maximumUsers"] = o.MaximumUsers
	}
	if !IsNil(o.MinimumUsers) {
		toSerialize["minimumUsers"] = o.MinimumUsers
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Term) {
		toSerialize["term"] = o.Term
	}
	if !IsNil(o.TermEndTime) {
		toSerialize["termEndTime"] = o.TermEndTime
	}
	if !IsNil(o.TimeUnit) {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	return toSerialize, nil
}

type NullableCommitDimension struct {
	value *CommitDimension
	isSet bool
}

func (v NullableCommitDimension) Get() *CommitDimension {
	return v.value
}

func (v *NullableCommitDimension) Set(val *CommitDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitDimension(val *CommitDimension) *NullableCommitDimension {
	return &NullableCommitDimension{value: val, isSet: true}
}

func (v NullableCommitDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
