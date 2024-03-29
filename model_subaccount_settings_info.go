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
)

// checks if the SubaccountSettingsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountSettingsInfo{}

// SubaccountSettingsInfo SubAccount settings
type SubaccountSettingsInfo struct {
	Email *SubaccountEmailSettings `json:"Email,omitempty"`
}

// NewSubaccountSettingsInfo instantiates a new SubaccountSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountSettingsInfo() *SubaccountSettingsInfo {
	this := SubaccountSettingsInfo{}
	return &this
}

// NewSubaccountSettingsInfoWithDefaults instantiates a new SubaccountSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountSettingsInfoWithDefaults() *SubaccountSettingsInfo {
	this := SubaccountSettingsInfo{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubaccountSettingsInfo) GetEmail() SubaccountEmailSettings {
	if o == nil || isNil(o.Email) {
		var ret SubaccountEmailSettings
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountSettingsInfo) GetEmailOk() (*SubaccountEmailSettings, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubaccountSettingsInfo) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given SubaccountEmailSettings and assigns it to the Email field.
func (o *SubaccountSettingsInfo) SetEmail(v SubaccountEmailSettings) {
	o.Email = &v
}

func (o SubaccountSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountSettingsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	return toSerialize, nil
}

type NullableSubaccountSettingsInfo struct {
	value *SubaccountSettingsInfo
	isSet bool
}

func (v NullableSubaccountSettingsInfo) Get() *SubaccountSettingsInfo {
	return v.value
}

func (v *NullableSubaccountSettingsInfo) Set(val *SubaccountSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountSettingsInfo(val *SubaccountSettingsInfo) *NullableSubaccountSettingsInfo {
	return &NullableSubaccountSettingsInfo{value: val, isSet: true}
}

func (v NullableSubaccountSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


