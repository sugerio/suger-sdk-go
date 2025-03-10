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

// checks if the CommitRevenueDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitRevenueDetail{}

// CommitRevenueDetail struct for CommitRevenueDetail
type CommitRevenueDetail struct {
	// The total amount of the commit revenue.
	Amount      *float32 `json:"amount,omitempty"`
	Description *string  `json:"description,omitempty"`
	Expression  *string  `json:"expression,omitempty"`
	Key         *string  `json:"key,omitempty"`
	Name        *string  `json:"name,omitempty"`
	// The quantity of the commit dimension. Default is 1.
	Quantity *int32 `json:"quantity,omitempty"`
	// The unit price of the commit dimension.
	Rate *float32 `json:"rate,omitempty"`
}

// NewCommitRevenueDetail instantiates a new CommitRevenueDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitRevenueDetail() *CommitRevenueDetail {
	this := CommitRevenueDetail{}
	return &this
}

// NewCommitRevenueDetailWithDefaults instantiates a new CommitRevenueDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitRevenueDetailWithDefaults() *CommitRevenueDetail {
	this := CommitRevenueDetail{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *CommitRevenueDetail) SetAmount(v float32) {
	o.Amount = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommitRevenueDetail) SetDescription(v string) {
	o.Description = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CommitRevenueDetail) SetExpression(v string) {
	o.Expression = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CommitRevenueDetail) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommitRevenueDetail) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *CommitRevenueDetail) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *CommitRevenueDetail) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitRevenueDetail) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *CommitRevenueDetail) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *CommitRevenueDetail) SetRate(v float32) {
	o.Rate = &v
}

func (o CommitRevenueDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitRevenueDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
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
	return toSerialize, nil
}

type NullableCommitRevenueDetail struct {
	value *CommitRevenueDetail
	isSet bool
}

func (v NullableCommitRevenueDetail) Get() *CommitRevenueDetail {
	return v.value
}

func (v *NullableCommitRevenueDetail) Set(val *CommitRevenueDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitRevenueDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitRevenueDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitRevenueDetail(val *CommitRevenueDetail) *NullableCommitRevenueDetail {
	return &NullableCommitRevenueDetail{value: val, isSet: true}
}

func (v NullableCommitRevenueDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitRevenueDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
