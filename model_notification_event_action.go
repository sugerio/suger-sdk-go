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
	"fmt"
)

// NotificationEventAction the model 'NotificationEventAction'
type NotificationEventAction string

// List of NotificationEventAction
const (
	NotificationEventAction_UNKNOWN                     NotificationEventAction = ""
	NotificationEventAction_ABNORMAL_ALERT              NotificationEventAction = "ABNORMAL_ALERT"
	NotificationEventAction_ACCEPT                      NotificationEventAction = "ACCEPT"
	NotificationEventAction_ADD                         NotificationEventAction = "ADD"
	NotificationEventAction_APPROVE                     NotificationEventAction = "APPROVE"
	NotificationEventAction_CANCEL                      NotificationEventAction = "CANCEL"
	NotificationEventAction_CLOSE                       NotificationEventAction = "CLOSE"
	NotificationEventAction_CREATE                      NotificationEventAction = "CREATE"
	NotificationEventAction_DELETE                      NotificationEventAction = "DELETE"
	NotificationEventAction_DISBURSE                    NotificationEventAction = "DISBURSE"
	NotificationEventAction_END_SOON                    NotificationEventAction = "END_SOON"
	NotificationEventAction_EXPIRE                      NotificationEventAction = "EXPIRE"
	NotificationEventAction_EXPIRE_SOON                 NotificationEventAction = "EXPIRE_SOON"
	NotificationEventAction_MERGE                       NotificationEventAction = "MERGE"
	NotificationEventAction_METER                       NotificationEventAction = "METER"
	NotificationEventAction_NEW_CLIENT_SIGNUP           NotificationEventAction = "NEW_CLIENT_SIGNUP"
	NotificationEventAction_NOTIFY                      NotificationEventAction = "NOTIFY"
	NotificationEventAction_NOTIFY_CONTACTS             NotificationEventAction = "NOTIFY_CONTACTS"
	NotificationEventAction_OPEN_EMAIL                  NotificationEventAction = "OPEN_EMAIL"
	NotificationEventAction_PENDING_CANCEL              NotificationEventAction = "PENDING_CANCEL"
	NotificationEventAction_PENDING_ACCEPTANCE          NotificationEventAction = "PENDING_ACCEPTANCE"
	NotificationEventAction_REINSTATE                   NotificationEventAction = "REINSTATE"
	NotificationEventAction_REJECT                      NotificationEventAction = "REJECT"
	NotificationEventAction_REOPEN                      NotificationEventAction = "REOPEN"
	NotificationEventAction_CHARGE                      NotificationEventAction = "CHARGE"
	NotificationEventAction_REFUND                      NotificationEventAction = "REFUND"
	NotificationEventAction_ISSUE                       NotificationEventAction = "ISSUE"
	NotificationEventAction_ROTATE_SECRET               NotificationEventAction = "ROTATE_SECRET"
	NotificationEventAction_SUSPEND                     NotificationEventAction = "SUSPEND"
	NotificationEventAction_TEST                        NotificationEventAction = "TEST"
	NotificationEventAction_UPDATE                      NotificationEventAction = "UPDATE"
	NotificationEventAction_ACE_ENGAGEMENT_SCORE_UPDATE NotificationEventAction = "ACE_ENGAGEMENT_SCORE_UPDATE"
	NotificationEventAction_ACE_SALES_REP_UPDATE        NotificationEventAction = "ACE_SALES_REP_UPDATE"
	NotificationEventAction_ACE_CUSTOMER_EMAIL_UPDATE   NotificationEventAction = "ACE_CUSTOMER_EMAIL_UPDATE"
	NotificationEventAction_SUBMIT_APPROVAL_REQUEST     NotificationEventAction = "SUBMIT_APPROVAL_REQUEST"
	NotificationEventAction_REVIEW_APPROVAL_REQUEST     NotificationEventAction = "REVIEW_APPROVAL_REQUEST"
	NotificationEventAction_COMPLETE                    NotificationEventAction = "COMPLETE"
	NotificationEventAction_FAIL                        NotificationEventAction = "FAIL"
	NotificationEventAction_WEBHOOK                     NotificationEventAction = "WEBHOOK"
)

// All allowed values of NotificationEventAction enum
var AllowedNotificationEventActionEnumValues = []NotificationEventAction{
	"",
	"ABNORMAL_ALERT",
	"ACCEPT",
	"ADD",
	"APPROVE",
	"CANCEL",
	"CLOSE",
	"CREATE",
	"DELETE",
	"DISBURSE",
	"END_SOON",
	"EXPIRE",
	"EXPIRE_SOON",
	"MERGE",
	"METER",
	"NEW_CLIENT_SIGNUP",
	"NOTIFY",
	"NOTIFY_CONTACTS",
	"OPEN_EMAIL",
	"PENDING_CANCEL",
	"PENDING_ACCEPTANCE",
	"REINSTATE",
	"REJECT",
	"REOPEN",
	"CHARGE",
	"REFUND",
	"ISSUE",
	"ROTATE_SECRET",
	"SUSPEND",
	"TEST",
	"UPDATE",
	"ACE_ENGAGEMENT_SCORE_UPDATE",
	"ACE_SALES_REP_UPDATE",
	"ACE_CUSTOMER_EMAIL_UPDATE",
	"SUBMIT_APPROVAL_REQUEST",
	"REVIEW_APPROVAL_REQUEST",
	"COMPLETE",
	"FAIL",
	"WEBHOOK",
}

func (v *NotificationEventAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationEventAction(value)
	for _, existing := range AllowedNotificationEventActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationEventAction", value)
}

// NewNotificationEventActionFromValue returns a pointer to a valid NotificationEventAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationEventActionFromValue(v string) (*NotificationEventAction, error) {
	ev := NotificationEventAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationEventAction: valid values are %v", v, AllowedNotificationEventActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationEventAction) IsValid() bool {
	for _, existing := range AllowedNotificationEventActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationEventAction value
func (v NotificationEventAction) Ptr() *NotificationEventAction {
	return &v
}

type NullableNotificationEventAction struct {
	value *NotificationEventAction
	isSet bool
}

func (v NullableNotificationEventAction) Get() *NotificationEventAction {
	return v.value
}

func (v *NullableNotificationEventAction) Set(val *NotificationEventAction) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationEventAction) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationEventAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationEventAction(val *NotificationEventAction) *NullableNotificationEventAction {
	return &NullableNotificationEventAction{value: val, isSet: true}
}

func (v NullableNotificationEventAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationEventAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
