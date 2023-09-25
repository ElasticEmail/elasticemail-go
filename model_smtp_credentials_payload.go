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
	"time"
)

// checks if the SmtpCredentialsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmtpCredentialsPayload{}

// SmtpCredentialsPayload Create new SMTP Credentials
type SmtpCredentialsPayload struct {
	// Name of the Credential for ease of reference. It must be a valid email address.
	Name string `json:"Name"`
	// Date this SmtpCredential expires.
	Expires NullableTime `json:"Expires,omitempty"`
	// Which IPs can use this SmtpCredential
	RestrictAccessToIPRange []string `json:"RestrictAccessToIPRange,omitempty"`
	// Email of the subaccount for which this SmtpCredential should be created
	Subaccount *string `json:"Subaccount,omitempty"`
}

// NewSmtpCredentialsPayload instantiates a new SmtpCredentialsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmtpCredentialsPayload(name string) *SmtpCredentialsPayload {
	this := SmtpCredentialsPayload{}
	this.Name = name
	return &this
}

// NewSmtpCredentialsPayloadWithDefaults instantiates a new SmtpCredentialsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmtpCredentialsPayloadWithDefaults() *SmtpCredentialsPayload {
	this := SmtpCredentialsPayload{}
	return &this
}

// GetName returns the Name field value
func (o *SmtpCredentialsPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SmtpCredentialsPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SmtpCredentialsPayload) SetName(v string) {
	o.Name = v
}

// GetExpires returns the Expires field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmtpCredentialsPayload) GetExpires() time.Time {
	if o == nil || isNil(o.Expires.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expires.Get()
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmtpCredentialsPayload) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expires.Get(), o.Expires.IsSet()
}

// HasExpires returns a boolean if a field has been set.
func (o *SmtpCredentialsPayload) HasExpires() bool {
	if o != nil && o.Expires.IsSet() {
		return true
	}

	return false
}

// SetExpires gets a reference to the given NullableTime and assigns it to the Expires field.
func (o *SmtpCredentialsPayload) SetExpires(v time.Time) {
	o.Expires.Set(&v)
}
// SetExpiresNil sets the value for Expires to be an explicit nil
func (o *SmtpCredentialsPayload) SetExpiresNil() {
	o.Expires.Set(nil)
}

// UnsetExpires ensures that no value is present for Expires, not even an explicit nil
func (o *SmtpCredentialsPayload) UnsetExpires() {
	o.Expires.Unset()
}

// GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field value if set, zero value otherwise.
func (o *SmtpCredentialsPayload) GetRestrictAccessToIPRange() []string {
	if o == nil || isNil(o.RestrictAccessToIPRange) {
		var ret []string
		return ret
	}
	return o.RestrictAccessToIPRange
}

// GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpCredentialsPayload) GetRestrictAccessToIPRangeOk() ([]string, bool) {
	if o == nil || isNil(o.RestrictAccessToIPRange) {
		return nil, false
	}
	return o.RestrictAccessToIPRange, true
}

// HasRestrictAccessToIPRange returns a boolean if a field has been set.
func (o *SmtpCredentialsPayload) HasRestrictAccessToIPRange() bool {
	if o != nil && !isNil(o.RestrictAccessToIPRange) {
		return true
	}

	return false
}

// SetRestrictAccessToIPRange gets a reference to the given []string and assigns it to the RestrictAccessToIPRange field.
func (o *SmtpCredentialsPayload) SetRestrictAccessToIPRange(v []string) {
	o.RestrictAccessToIPRange = v
}

// GetSubaccount returns the Subaccount field value if set, zero value otherwise.
func (o *SmtpCredentialsPayload) GetSubaccount() string {
	if o == nil || isNil(o.Subaccount) {
		var ret string
		return ret
	}
	return *o.Subaccount
}

// GetSubaccountOk returns a tuple with the Subaccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmtpCredentialsPayload) GetSubaccountOk() (*string, bool) {
	if o == nil || isNil(o.Subaccount) {
		return nil, false
	}
	return o.Subaccount, true
}

// HasSubaccount returns a boolean if a field has been set.
func (o *SmtpCredentialsPayload) HasSubaccount() bool {
	if o != nil && !isNil(o.Subaccount) {
		return true
	}

	return false
}

// SetSubaccount gets a reference to the given string and assigns it to the Subaccount field.
func (o *SmtpCredentialsPayload) SetSubaccount(v string) {
	o.Subaccount = &v
}

func (o SmtpCredentialsPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmtpCredentialsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if o.Expires.IsSet() {
		toSerialize["Expires"] = o.Expires.Get()
	}
	if !isNil(o.RestrictAccessToIPRange) {
		toSerialize["RestrictAccessToIPRange"] = o.RestrictAccessToIPRange
	}
	if !isNil(o.Subaccount) {
		toSerialize["Subaccount"] = o.Subaccount
	}
	return toSerialize, nil
}

type NullableSmtpCredentialsPayload struct {
	value *SmtpCredentialsPayload
	isSet bool
}

func (v NullableSmtpCredentialsPayload) Get() *SmtpCredentialsPayload {
	return v.value
}

func (v *NullableSmtpCredentialsPayload) Set(val *SmtpCredentialsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSmtpCredentialsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSmtpCredentialsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmtpCredentialsPayload(val *SmtpCredentialsPayload) *NullableSmtpCredentialsPayload {
	return &NullableSmtpCredentialsPayload{value: val, isSet: true}
}

func (v NullableSmtpCredentialsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmtpCredentialsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


