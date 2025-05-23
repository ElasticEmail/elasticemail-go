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

// TrackingValidationStatus Status of ValidDomain to determine how often tracking validation should be performed.
type TrackingValidationStatus string

// List of TrackingValidationStatus
const (
	VALIDATED TrackingValidationStatus = "Validated"
	NOT_VALIDATED TrackingValidationStatus = "NotValidated"
	INVALID TrackingValidationStatus = "Invalid"
	BROKEN TrackingValidationStatus = "Broken"
)

// All allowed values of TrackingValidationStatus enum
var AllowedTrackingValidationStatusEnumValues = []TrackingValidationStatus{
	"Validated",
	"NotValidated",
	"Invalid",
	"Broken",
}

func (v *TrackingValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrackingValidationStatus(value)
	for _, existing := range AllowedTrackingValidationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrackingValidationStatus", value)
}

// NewTrackingValidationStatusFromValue returns a pointer to a valid TrackingValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrackingValidationStatusFromValue(v string) (*TrackingValidationStatus, error) {
	ev := TrackingValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrackingValidationStatus: valid values are %v", v, AllowedTrackingValidationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrackingValidationStatus) IsValid() bool {
	for _, existing := range AllowedTrackingValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrackingValidationStatus value
func (v TrackingValidationStatus) Ptr() *TrackingValidationStatus {
	return &v
}

type NullableTrackingValidationStatus struct {
	value *TrackingValidationStatus
	isSet bool
}

func (v NullableTrackingValidationStatus) Get() *TrackingValidationStatus {
	return v.value
}

func (v *NullableTrackingValidationStatus) Set(val *TrackingValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingValidationStatus(val *TrackingValidationStatus) *NullableTrackingValidationStatus {
	return &NullableTrackingValidationStatus{value: val, isSet: true}
}

func (v NullableTrackingValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrackingValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

