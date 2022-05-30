/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ElasticEmail

import (
	"encoding/json"
	"fmt"
)

// ContactSource the model 'ContactSource'
type ContactSource string

// List of ContactSource
const (
	DELIVERY_API ContactSource = "DeliveryApi"
	MANUAL_INPUT ContactSource = "ManualInput"
	FILE_UPLOAD ContactSource = "FileUpload"
	WEB_FORM ContactSource = "WebForm"
	CONTACT_API ContactSource = "ContactApi"
	VERIFICATION_API ContactSource = "VerificationApi"
	FILE_VERIFICATION_API ContactSource = "FileVerificationApi"
)

// All allowed values of ContactSource enum
var AllowedContactSourceEnumValues = []ContactSource{
	"DeliveryApi",
	"ManualInput",
	"FileUpload",
	"WebForm",
	"ContactApi",
	"VerificationApi",
	"FileVerificationApi",
}

func (v *ContactSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContactSource(value)
	for _, existing := range AllowedContactSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContactSource", value)
}

// NewContactSourceFromValue returns a pointer to a valid ContactSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContactSourceFromValue(v string) (*ContactSource, error) {
	ev := ContactSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContactSource: valid values are %v", v, AllowedContactSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContactSource) IsValid() bool {
	for _, existing := range AllowedContactSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContactSource value
func (v ContactSource) Ptr() *ContactSource {
	return &v
}

type NullableContactSource struct {
	value *ContactSource
	isSet bool
}

func (v NullableContactSource) Get() *ContactSource {
	return v.value
}

func (v *NullableContactSource) Set(val *ContactSource) {
	v.value = val
	v.isSet = true
}

func (v NullableContactSource) IsSet() bool {
	return v.isSet
}

func (v *NullableContactSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactSource(val *ContactSource) *NullableContactSource {
	return &NullableContactSource{value: val, isSet: true}
}

func (v NullableContactSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

