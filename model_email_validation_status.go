/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://app.elasticemail.com/marketing/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// EmailValidationStatus the model 'EmailValidationStatus'
type EmailValidationStatus string

// List of EmailValidationStatus
const (
	EMAILVALIDATIONSTATUS_NONE EmailValidationStatus = "None"
	EMAILVALIDATIONSTATUS_VALID EmailValidationStatus = "Valid"
	EMAILVALIDATIONSTATUS_UNKNOWN EmailValidationStatus = "Unknown"
	EMAILVALIDATIONSTATUS_RISKY EmailValidationStatus = "Risky"
	EMAILVALIDATIONSTATUS_INVALID EmailValidationStatus = "Invalid"
)

// All allowed values of EmailValidationStatus enum
var AllowedEmailValidationStatusEnumValues = []EmailValidationStatus{
	"None",
	"Valid",
	"Unknown",
	"Risky",
	"Invalid",
}

func (v *EmailValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmailValidationStatus(value)
	for _, existing := range AllowedEmailValidationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmailValidationStatus", value)
}

// NewEmailValidationStatusFromValue returns a pointer to a valid EmailValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmailValidationStatusFromValue(v string) (*EmailValidationStatus, error) {
	ev := EmailValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmailValidationStatus: valid values are %v", v, AllowedEmailValidationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmailValidationStatus) IsValid() bool {
	for _, existing := range AllowedEmailValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmailValidationStatus value
func (v EmailValidationStatus) Ptr() *EmailValidationStatus {
	return &v
}

type NullableEmailValidationStatus struct {
	value *EmailValidationStatus
	isSet bool
}

func (v NullableEmailValidationStatus) Get() *EmailValidationStatus {
	return v.value
}

func (v *NullableEmailValidationStatus) Set(val *EmailValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailValidationStatus(val *EmailValidationStatus) *NullableEmailValidationStatus {
	return &NullableEmailValidationStatus{value: val, isSet: true}
}

func (v NullableEmailValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

