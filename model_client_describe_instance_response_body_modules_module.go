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

// checks if the ClientDescribeInstanceResponseBodyModulesModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientDescribeInstanceResponseBodyModulesModule{}

// ClientDescribeInstanceResponseBodyModulesModule struct for ClientDescribeInstanceResponseBodyModulesModule
type ClientDescribeInstanceResponseBodyModulesModule struct {
	Code       *string                                                    `json:"Code,omitempty"`
	Id         *string                                                    `json:"Id,omitempty"`
	Name       *string                                                    `json:"Name,omitempty"`
	Properties *ClientDescribeInstanceResponseBodyModulesModuleProperties `json:"Properties,omitempty"`
}

// NewClientDescribeInstanceResponseBodyModulesModule instantiates a new ClientDescribeInstanceResponseBodyModulesModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientDescribeInstanceResponseBodyModulesModule() *ClientDescribeInstanceResponseBodyModulesModule {
	this := ClientDescribeInstanceResponseBodyModulesModule{}
	return &this
}

// NewClientDescribeInstanceResponseBodyModulesModuleWithDefaults instantiates a new ClientDescribeInstanceResponseBodyModulesModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientDescribeInstanceResponseBodyModulesModuleWithDefaults() *ClientDescribeInstanceResponseBodyModulesModule {
	this := ClientDescribeInstanceResponseBodyModulesModule{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ClientDescribeInstanceResponseBodyModulesModule) SetCode(v string) {
	o.Code = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientDescribeInstanceResponseBodyModulesModule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClientDescribeInstanceResponseBodyModulesModule) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetProperties() ClientDescribeInstanceResponseBodyModulesModuleProperties {
	if o == nil || IsNil(o.Properties) {
		var ret ClientDescribeInstanceResponseBodyModulesModuleProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) GetPropertiesOk() (*ClientDescribeInstanceResponseBodyModulesModuleProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ClientDescribeInstanceResponseBodyModulesModule) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given ClientDescribeInstanceResponseBodyModulesModuleProperties and assigns it to the Properties field.
func (o *ClientDescribeInstanceResponseBodyModulesModule) SetProperties(v ClientDescribeInstanceResponseBodyModulesModuleProperties) {
	o.Properties = &v
}

func (o ClientDescribeInstanceResponseBodyModulesModule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientDescribeInstanceResponseBodyModulesModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Properties) {
		toSerialize["Properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableClientDescribeInstanceResponseBodyModulesModule struct {
	value *ClientDescribeInstanceResponseBodyModulesModule
	isSet bool
}

func (v NullableClientDescribeInstanceResponseBodyModulesModule) Get() *ClientDescribeInstanceResponseBodyModulesModule {
	return v.value
}

func (v *NullableClientDescribeInstanceResponseBodyModulesModule) Set(val *ClientDescribeInstanceResponseBodyModulesModule) {
	v.value = val
	v.isSet = true
}

func (v NullableClientDescribeInstanceResponseBodyModulesModule) IsSet() bool {
	return v.isSet
}

func (v *NullableClientDescribeInstanceResponseBodyModulesModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientDescribeInstanceResponseBodyModulesModule(val *ClientDescribeInstanceResponseBodyModulesModule) *NullableClientDescribeInstanceResponseBodyModulesModule {
	return &NullableClientDescribeInstanceResponseBodyModulesModule{value: val, isSet: true}
}

func (v NullableClientDescribeInstanceResponseBodyModulesModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientDescribeInstanceResponseBodyModulesModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
