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
	"bytes"
	"fmt"
)

// checks if the SubaccountPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountPayload{}

// SubaccountPayload New SubAccount payload
type SubaccountPayload struct {
	// Proper email address.
	Email string `json:"Email"`
	// Current password.
	Password string `json:"Password"`
	// True, if you want to send activation email to this Account to confirm the creation of a new SubAccount. Otherwise, false (SubAccount will immediately be Active).
	SendActivation *bool `json:"SendActivation,omitempty"`
	Settings *SubaccountSettingsInfoPayload `json:"Settings,omitempty"`
}

type _SubaccountPayload SubaccountPayload

// NewSubaccountPayload instantiates a new SubaccountPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountPayload(email string, password string) *SubaccountPayload {
	this := SubaccountPayload{}
	this.Email = email
	this.Password = password
	return &this
}

// NewSubaccountPayloadWithDefaults instantiates a new SubaccountPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountPayloadWithDefaults() *SubaccountPayload {
	this := SubaccountPayload{}
	return &this
}

// GetEmail returns the Email field value
func (o *SubaccountPayload) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SubaccountPayload) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SubaccountPayload) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *SubaccountPayload) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *SubaccountPayload) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *SubaccountPayload) SetPassword(v string) {
	o.Password = v
}

// GetSendActivation returns the SendActivation field value if set, zero value otherwise.
func (o *SubaccountPayload) GetSendActivation() bool {
	if o == nil || IsNil(o.SendActivation) {
		var ret bool
		return ret
	}
	return *o.SendActivation
}

// GetSendActivationOk returns a tuple with the SendActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountPayload) GetSendActivationOk() (*bool, bool) {
	if o == nil || IsNil(o.SendActivation) {
		return nil, false
	}
	return o.SendActivation, true
}

// HasSendActivation returns a boolean if a field has been set.
func (o *SubaccountPayload) HasSendActivation() bool {
	if o != nil && !IsNil(o.SendActivation) {
		return true
	}

	return false
}

// SetSendActivation gets a reference to the given bool and assigns it to the SendActivation field.
func (o *SubaccountPayload) SetSendActivation(v bool) {
	o.SendActivation = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SubaccountPayload) GetSettings() SubaccountSettingsInfoPayload {
	if o == nil || IsNil(o.Settings) {
		var ret SubaccountSettingsInfoPayload
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountPayload) GetSettingsOk() (*SubaccountSettingsInfoPayload, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SubaccountPayload) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SubaccountSettingsInfoPayload and assigns it to the Settings field.
func (o *SubaccountPayload) SetSettings(v SubaccountSettingsInfoPayload) {
	o.Settings = &v
}

func (o SubaccountPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Email"] = o.Email
	toSerialize["Password"] = o.Password
	if !IsNil(o.SendActivation) {
		toSerialize["SendActivation"] = o.SendActivation
	}
	if !IsNil(o.Settings) {
		toSerialize["Settings"] = o.Settings
	}
	return toSerialize, nil
}

func (o *SubaccountPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Email",
		"Password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubaccountPayload := _SubaccountPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubaccountPayload)

	if err != nil {
		return err
	}

	*o = SubaccountPayload(varSubaccountPayload)

	return err
}

type NullableSubaccountPayload struct {
	value *SubaccountPayload
	isSet bool
}

func (v NullableSubaccountPayload) Get() *SubaccountPayload {
	return v.value
}

func (v *NullableSubaccountPayload) Set(val *SubaccountPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountPayload(val *SubaccountPayload) *NullableSubaccountPayload {
	return &NullableSubaccountPayload{value: val, isSet: true}
}

func (v NullableSubaccountPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


