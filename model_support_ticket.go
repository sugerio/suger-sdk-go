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

// checks if the SupportTicket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportTicket{}

// SupportTicket struct for SupportTicket
type SupportTicket struct {
	Attachments    []SupportTicketAttachment `json:"attachments,omitempty"`
	CloseTime      *string                   `json:"closeTime,omitempty"`
	Comments       []SupportTicketComment    `json:"comments,omitempty"`
	CreationTime   *string                   `json:"creationTime,omitempty"`
	Creator        *string                   `json:"creator,omitempty"`
	Description    *string                   `json:"description,omitempty"`
	DueTime        *string                   `json:"dueTime,omitempty"`
	Id             *string                   `json:"id,omitempty"`
	Name           *string                   `json:"name,omitempty"`
	OrganizationId *string                   `json:"organizationId,omitempty"`
	Priority       *SupportTicketPriority    `json:"priority,omitempty"`
	Status         *SupportTicketStatus      `json:"status,omitempty"`
	Watchers       []string                  `json:"watchers,omitempty"`
}

// NewSupportTicket instantiates a new SupportTicket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportTicket() *SupportTicket {
	this := SupportTicket{}
	return &this
}

// NewSupportTicketWithDefaults instantiates a new SupportTicket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportTicketWithDefaults() *SupportTicket {
	this := SupportTicket{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *SupportTicket) GetAttachments() []SupportTicketAttachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []SupportTicketAttachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetAttachmentsOk() ([]SupportTicketAttachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *SupportTicket) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []SupportTicketAttachment and assigns it to the Attachments field.
func (o *SupportTicket) SetAttachments(v []SupportTicketAttachment) {
	o.Attachments = v
}

// GetCloseTime returns the CloseTime field value if set, zero value otherwise.
func (o *SupportTicket) GetCloseTime() string {
	if o == nil || IsNil(o.CloseTime) {
		var ret string
		return ret
	}
	return *o.CloseTime
}

// GetCloseTimeOk returns a tuple with the CloseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetCloseTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CloseTime) {
		return nil, false
	}
	return o.CloseTime, true
}

// HasCloseTime returns a boolean if a field has been set.
func (o *SupportTicket) HasCloseTime() bool {
	if o != nil && !IsNil(o.CloseTime) {
		return true
	}

	return false
}

// SetCloseTime gets a reference to the given string and assigns it to the CloseTime field.
func (o *SupportTicket) SetCloseTime(v string) {
	o.CloseTime = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *SupportTicket) GetComments() []SupportTicketComment {
	if o == nil || IsNil(o.Comments) {
		var ret []SupportTicketComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetCommentsOk() ([]SupportTicketComment, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *SupportTicket) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given []SupportTicketComment and assigns it to the Comments field.
func (o *SupportTicket) SetComments(v []SupportTicketComment) {
	o.Comments = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *SupportTicket) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime) {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetCreationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *SupportTicket) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *SupportTicket) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *SupportTicket) GetCreator() string {
	if o == nil || IsNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *SupportTicket) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *SupportTicket) SetCreator(v string) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SupportTicket) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SupportTicket) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SupportTicket) SetDescription(v string) {
	o.Description = &v
}

// GetDueTime returns the DueTime field value if set, zero value otherwise.
func (o *SupportTicket) GetDueTime() string {
	if o == nil || IsNil(o.DueTime) {
		var ret string
		return ret
	}
	return *o.DueTime
}

// GetDueTimeOk returns a tuple with the DueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetDueTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DueTime) {
		return nil, false
	}
	return o.DueTime, true
}

// HasDueTime returns a boolean if a field has been set.
func (o *SupportTicket) HasDueTime() bool {
	if o != nil && !IsNil(o.DueTime) {
		return true
	}

	return false
}

// SetDueTime gets a reference to the given string and assigns it to the DueTime field.
func (o *SupportTicket) SetDueTime(v string) {
	o.DueTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SupportTicket) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SupportTicket) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SupportTicket) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SupportTicket) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SupportTicket) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SupportTicket) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SupportTicket) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SupportTicket) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SupportTicket) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SupportTicket) GetPriority() SupportTicketPriority {
	if o == nil || IsNil(o.Priority) {
		var ret SupportTicketPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetPriorityOk() (*SupportTicketPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SupportTicket) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given SupportTicketPriority and assigns it to the Priority field.
func (o *SupportTicket) SetPriority(v SupportTicketPriority) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SupportTicket) GetStatus() SupportTicketStatus {
	if o == nil || IsNil(o.Status) {
		var ret SupportTicketStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetStatusOk() (*SupportTicketStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SupportTicket) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SupportTicketStatus and assigns it to the Status field.
func (o *SupportTicket) SetStatus(v SupportTicketStatus) {
	o.Status = &v
}

// GetWatchers returns the Watchers field value if set, zero value otherwise.
func (o *SupportTicket) GetWatchers() []string {
	if o == nil || IsNil(o.Watchers) {
		var ret []string
		return ret
	}
	return o.Watchers
}

// GetWatchersOk returns a tuple with the Watchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportTicket) GetWatchersOk() ([]string, bool) {
	if o == nil || IsNil(o.Watchers) {
		return nil, false
	}
	return o.Watchers, true
}

// HasWatchers returns a boolean if a field has been set.
func (o *SupportTicket) HasWatchers() bool {
	if o != nil && !IsNil(o.Watchers) {
		return true
	}

	return false
}

// SetWatchers gets a reference to the given []string and assigns it to the Watchers field.
func (o *SupportTicket) SetWatchers(v []string) {
	o.Watchers = v
}

func (o SupportTicket) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportTicket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.CloseTime) {
		toSerialize["closeTime"] = o.CloseTime
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DueTime) {
		toSerialize["dueTime"] = o.DueTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Watchers) {
		toSerialize["watchers"] = o.Watchers
	}
	return toSerialize, nil
}

type NullableSupportTicket struct {
	value *SupportTicket
	isSet bool
}

func (v NullableSupportTicket) Get() *SupportTicket {
	return v.value
}

func (v *NullableSupportTicket) Set(val *SupportTicket) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportTicket) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportTicket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportTicket(val *SupportTicket) *NullableSupportTicket {
	return &NullableSupportTicket{value: val, isSet: true}
}

func (v NullableSupportTicket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportTicket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
