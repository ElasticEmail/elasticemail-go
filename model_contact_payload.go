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

// checks if the ContactPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactPayload{}

// ContactPayload struct for ContactPayload
type ContactPayload struct {
	// Proper email address.
	Email string `json:"Email"`
	Status *ContactStatus `json:"Status,omitempty"`
	// First name.
	FirstName *string `json:"FirstName,omitempty"`
	// Last name.
	LastName *string `json:"LastName,omitempty"`
	// A key-value collection of custom contact fields which can be used in the system. Only already existing custom fields will be saved.
	CustomFields *map[string]string `json:"CustomFields,omitempty"`
	Consent *ConsentData `json:"Consent,omitempty"`
}

// NewContactPayload instantiates a new ContactPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactPayload(email string) *ContactPayload {
	this := ContactPayload{}
	this.Email = email
	var status ContactStatus = CONTACTSTATUS_TRANSACTIONAL
	this.Status = &status
	return &this
}

// NewContactPayloadWithDefaults instantiates a new ContactPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactPayloadWithDefaults() *ContactPayload {
	this := ContactPayload{}
	var status ContactStatus = CONTACTSTATUS_TRANSACTIONAL
	this.Status = &status
	return &this
}

// GetEmail returns the Email field value
func (o *ContactPayload) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ContactPayload) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ContactPayload) SetEmail(v string) {
	o.Email = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ContactPayload) GetStatus() ContactStatus {
	if o == nil || IsNil(o.Status) {
		var ret ContactStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactPayload) GetStatusOk() (*ContactStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ContactPayload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ContactStatus and assigns it to the Status field.
func (o *ContactPayload) SetStatus(v ContactStatus) {
	o.Status = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ContactPayload) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactPayload) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ContactPayload) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ContactPayload) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ContactPayload) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactPayload) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ContactPayload) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ContactPayload) SetLastName(v string) {
	o.LastName = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ContactPayload) GetCustomFields() map[string]string {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]string
		return ret
	}
	return *o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactPayload) GetCustomFieldsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ContactPayload) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]string and assigns it to the CustomFields field.
func (o *ContactPayload) SetCustomFields(v map[string]string) {
	o.CustomFields = &v
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *ContactPayload) GetConsent() ConsentData {
	if o == nil || IsNil(o.Consent) {
		var ret ConsentData
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactPayload) GetConsentOk() (*ConsentData, bool) {
	if o == nil || IsNil(o.Consent) {
		return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *ContactPayload) HasConsent() bool {
	if o != nil && !IsNil(o.Consent) {
		return true
	}

	return false
}

// SetConsent gets a reference to the given ConsentData and assigns it to the Consent field.
func (o *ContactPayload) SetConsent(v ConsentData) {
	o.Consent = &v
}

func (o ContactPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Email"] = o.Email
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.FirstName) {
		toSerialize["FirstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["LastName"] = o.LastName
	}
	if !IsNil(o.CustomFields) {
		toSerialize["CustomFields"] = o.CustomFields
	}
	if !IsNil(o.Consent) {
		toSerialize["Consent"] = o.Consent
	}
	return toSerialize, nil
}

type NullableContactPayload struct {
	value *ContactPayload
	isSet bool
}

func (v NullableContactPayload) Get() *ContactPayload {
	return v.value
}

func (v *NullableContactPayload) Set(val *ContactPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableContactPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableContactPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactPayload(val *ContactPayload) *NullableContactPayload {
	return &NullableContactPayload{value: val, isSet: true}
}

func (v NullableContactPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


