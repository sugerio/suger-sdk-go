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

// checks if the AwsAceIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsAceIntegration{}

// AwsAceIntegration struct for AwsAceIntegration
type AwsAceIntegration struct {
	Credential *AwsIntegrationCredential `json:"credential,omitempty"`
	// Enable assume IAM role to access the client's AWS ACE S3 bucket. If false, Suger will use the AwsIntegrationCredential to access the client's AWS ACE S3 bucket.
	EnableAssumeIamRole *bool `json:"enableAssumeIamRole,omitempty"`
	// The AWS IAM role for Suger service to assume to access the client's AWS ACE S3 bucket. Applicable only if EnableAssumeIamRole is true.
	IamRoleArn *string `json:"iamRoleArn,omitempty"`
	// The partner ID of the ISV/Seller in AWS Partner Network.
	PartnerId *string `json:"partnerId,omitempty"`
	// Is the integration paused.
	Paused *bool `json:"paused,omitempty"`
	// The Name of the S3 bucket for AWS APN Customer Engagement Program (ACE) to sync the leads & opportunities.
	S3BucketName *string `json:"s3BucketName,omitempty"`
	// The region of the S3 bucket for AWS APN Customer Engagement Program (ACE) to sync the leads & opportunities.
	S3BucketRegion *string `json:"s3BucketRegion,omitempty"`
	// The secret key used to store the AwsIntegrationCredential in AWS Secret manager. for internal usage only.
	SecretKey *string `json:"secretKey,omitempty"`
}

// NewAwsAceIntegration instantiates a new AwsAceIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsAceIntegration() *AwsAceIntegration {
	this := AwsAceIntegration{}
	return &this
}

// NewAwsAceIntegrationWithDefaults instantiates a new AwsAceIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsAceIntegrationWithDefaults() *AwsAceIntegration {
	this := AwsAceIntegration{}
	return &this
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetCredential() AwsIntegrationCredential {
	if o == nil || IsNil(o.Credential) {
		var ret AwsIntegrationCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetCredentialOk() (*AwsIntegrationCredential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given AwsIntegrationCredential and assigns it to the Credential field.
func (o *AwsAceIntegration) SetCredential(v AwsIntegrationCredential) {
	o.Credential = &v
}

// GetEnableAssumeIamRole returns the EnableAssumeIamRole field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetEnableAssumeIamRole() bool {
	if o == nil || IsNil(o.EnableAssumeIamRole) {
		var ret bool
		return ret
	}
	return *o.EnableAssumeIamRole
}

// GetEnableAssumeIamRoleOk returns a tuple with the EnableAssumeIamRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetEnableAssumeIamRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAssumeIamRole) {
		return nil, false
	}
	return o.EnableAssumeIamRole, true
}

// HasEnableAssumeIamRole returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasEnableAssumeIamRole() bool {
	if o != nil && !IsNil(o.EnableAssumeIamRole) {
		return true
	}

	return false
}

// SetEnableAssumeIamRole gets a reference to the given bool and assigns it to the EnableAssumeIamRole field.
func (o *AwsAceIntegration) SetEnableAssumeIamRole(v bool) {
	o.EnableAssumeIamRole = &v
}

// GetIamRoleArn returns the IamRoleArn field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetIamRoleArn() string {
	if o == nil || IsNil(o.IamRoleArn) {
		var ret string
		return ret
	}
	return *o.IamRoleArn
}

// GetIamRoleArnOk returns a tuple with the IamRoleArn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetIamRoleArnOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleArn) {
		return nil, false
	}
	return o.IamRoleArn, true
}

// HasIamRoleArn returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasIamRoleArn() bool {
	if o != nil && !IsNil(o.IamRoleArn) {
		return true
	}

	return false
}

// SetIamRoleArn gets a reference to the given string and assigns it to the IamRoleArn field.
func (o *AwsAceIntegration) SetIamRoleArn(v string) {
	o.IamRoleArn = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *AwsAceIntegration) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetPaused returns the Paused field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetPaused() bool {
	if o == nil || IsNil(o.Paused) {
		var ret bool
		return ret
	}
	return *o.Paused
}

// GetPausedOk returns a tuple with the Paused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.Paused) {
		return nil, false
	}
	return o.Paused, true
}

// HasPaused returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasPaused() bool {
	if o != nil && !IsNil(o.Paused) {
		return true
	}

	return false
}

// SetPaused gets a reference to the given bool and assigns it to the Paused field.
func (o *AwsAceIntegration) SetPaused(v bool) {
	o.Paused = &v
}

// GetS3BucketName returns the S3BucketName field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetS3BucketName() string {
	if o == nil || IsNil(o.S3BucketName) {
		var ret string
		return ret
	}
	return *o.S3BucketName
}

// GetS3BucketNameOk returns a tuple with the S3BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetS3BucketNameOk() (*string, bool) {
	if o == nil || IsNil(o.S3BucketName) {
		return nil, false
	}
	return o.S3BucketName, true
}

// HasS3BucketName returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasS3BucketName() bool {
	if o != nil && !IsNil(o.S3BucketName) {
		return true
	}

	return false
}

// SetS3BucketName gets a reference to the given string and assigns it to the S3BucketName field.
func (o *AwsAceIntegration) SetS3BucketName(v string) {
	o.S3BucketName = &v
}

// GetS3BucketRegion returns the S3BucketRegion field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetS3BucketRegion() string {
	if o == nil || IsNil(o.S3BucketRegion) {
		var ret string
		return ret
	}
	return *o.S3BucketRegion
}

// GetS3BucketRegionOk returns a tuple with the S3BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetS3BucketRegionOk() (*string, bool) {
	if o == nil || IsNil(o.S3BucketRegion) {
		return nil, false
	}
	return o.S3BucketRegion, true
}

// HasS3BucketRegion returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasS3BucketRegion() bool {
	if o != nil && !IsNil(o.S3BucketRegion) {
		return true
	}

	return false
}

// SetS3BucketRegion gets a reference to the given string and assigns it to the S3BucketRegion field.
func (o *AwsAceIntegration) SetS3BucketRegion(v string) {
	o.S3BucketRegion = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *AwsAceIntegration) GetSecretKey() string {
	if o == nil || IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAceIntegration) GetSecretKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *AwsAceIntegration) HasSecretKey() bool {
	if o != nil && !IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *AwsAceIntegration) SetSecretKey(v string) {
	o.SecretKey = &v
}

func (o AwsAceIntegration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsAceIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	if !IsNil(o.EnableAssumeIamRole) {
		toSerialize["enableAssumeIamRole"] = o.EnableAssumeIamRole
	}
	if !IsNil(o.IamRoleArn) {
		toSerialize["iamRoleArn"] = o.IamRoleArn
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partnerId"] = o.PartnerId
	}
	if !IsNil(o.Paused) {
		toSerialize["paused"] = o.Paused
	}
	if !IsNil(o.S3BucketName) {
		toSerialize["s3BucketName"] = o.S3BucketName
	}
	if !IsNil(o.S3BucketRegion) {
		toSerialize["s3BucketRegion"] = o.S3BucketRegion
	}
	if !IsNil(o.SecretKey) {
		toSerialize["secretKey"] = o.SecretKey
	}
	return toSerialize, nil
}

type NullableAwsAceIntegration struct {
	value *AwsAceIntegration
	isSet bool
}

func (v NullableAwsAceIntegration) Get() *AwsAceIntegration {
	return v.value
}

func (v *NullableAwsAceIntegration) Set(val *AwsAceIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsAceIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsAceIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsAceIntegration(val *AwsAceIntegration) *NullableAwsAceIntegration {
	return &NullableAwsAceIntegration{value: val, isSet: true}
}

func (v NullableAwsAceIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsAceIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
