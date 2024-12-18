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

// checks if the EmailValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailValidationResult{}

// EmailValidationResult struct for EmailValidationResult
type EmailValidationResult struct {
	// Local part of an email
	Account *string `json:"Account,omitempty"`
	// Name of selected domain.
	Domain *string `json:"Domain,omitempty"`
	// Full email address that was verified
	Email *string `json:"Email,omitempty"`
	// Suggested spelling if a possible mistake was found
	SuggestedSpelling *string `json:"SuggestedSpelling,omitempty"`
	// Does the email have a temporary domain
	Disposable *bool `json:"Disposable,omitempty"`
	// Is an email a role email (e.g. info@, noreply@ etc.)
	Role *bool `json:"Role,omitempty"`
	// All detected issues
	Reason *string `json:"Reason,omitempty"`
	// Added date
	DateAdded *time.Time `json:"DateAdded,omitempty"`
	Result *EmailValidationStatus `json:"Result,omitempty"`
	// Predicted score
	PredictedScore *float32 `json:"PredictedScore,omitempty"`
	PredictedStatus *EmailPredictedValidationStatus `json:"PredictedStatus,omitempty"`
}

// NewEmailValidationResult instantiates a new EmailValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailValidationResult() *EmailValidationResult {
	this := EmailValidationResult{}
	var result EmailValidationStatus = NONE
	this.Result = &result
	var predictedStatus EmailPredictedValidationStatus = NONE
	this.PredictedStatus = &predictedStatus
	return &this
}

// NewEmailValidationResultWithDefaults instantiates a new EmailValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailValidationResultWithDefaults() *EmailValidationResult {
	this := EmailValidationResult{}
	var result EmailValidationStatus = NONE
	this.Result = &result
	var predictedStatus EmailPredictedValidationStatus = NONE
	this.PredictedStatus = &predictedStatus
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *EmailValidationResult) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *EmailValidationResult) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *EmailValidationResult) SetAccount(v string) {
	o.Account = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *EmailValidationResult) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *EmailValidationResult) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *EmailValidationResult) SetDomain(v string) {
	o.Domain = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmailValidationResult) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmailValidationResult) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmailValidationResult) SetEmail(v string) {
	o.Email = &v
}

// GetSuggestedSpelling returns the SuggestedSpelling field value if set, zero value otherwise.
func (o *EmailValidationResult) GetSuggestedSpelling() string {
	if o == nil || IsNil(o.SuggestedSpelling) {
		var ret string
		return ret
	}
	return *o.SuggestedSpelling
}

// GetSuggestedSpellingOk returns a tuple with the SuggestedSpelling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetSuggestedSpellingOk() (*string, bool) {
	if o == nil || IsNil(o.SuggestedSpelling) {
		return nil, false
	}
	return o.SuggestedSpelling, true
}

// HasSuggestedSpelling returns a boolean if a field has been set.
func (o *EmailValidationResult) HasSuggestedSpelling() bool {
	if o != nil && !IsNil(o.SuggestedSpelling) {
		return true
	}

	return false
}

// SetSuggestedSpelling gets a reference to the given string and assigns it to the SuggestedSpelling field.
func (o *EmailValidationResult) SetSuggestedSpelling(v string) {
	o.SuggestedSpelling = &v
}

// GetDisposable returns the Disposable field value if set, zero value otherwise.
func (o *EmailValidationResult) GetDisposable() bool {
	if o == nil || IsNil(o.Disposable) {
		var ret bool
		return ret
	}
	return *o.Disposable
}

// GetDisposableOk returns a tuple with the Disposable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetDisposableOk() (*bool, bool) {
	if o == nil || IsNil(o.Disposable) {
		return nil, false
	}
	return o.Disposable, true
}

// HasDisposable returns a boolean if a field has been set.
func (o *EmailValidationResult) HasDisposable() bool {
	if o != nil && !IsNil(o.Disposable) {
		return true
	}

	return false
}

// SetDisposable gets a reference to the given bool and assigns it to the Disposable field.
func (o *EmailValidationResult) SetDisposable(v bool) {
	o.Disposable = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *EmailValidationResult) GetRole() bool {
	if o == nil || IsNil(o.Role) {
		var ret bool
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *EmailValidationResult) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given bool and assigns it to the Role field.
func (o *EmailValidationResult) SetRole(v bool) {
	o.Role = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *EmailValidationResult) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *EmailValidationResult) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *EmailValidationResult) SetReason(v string) {
	o.Reason = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *EmailValidationResult) GetDateAdded() time.Time {
	if o == nil || IsNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *EmailValidationResult) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *EmailValidationResult) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *EmailValidationResult) GetResult() EmailValidationStatus {
	if o == nil || IsNil(o.Result) {
		var ret EmailValidationStatus
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetResultOk() (*EmailValidationStatus, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *EmailValidationResult) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given EmailValidationStatus and assigns it to the Result field.
func (o *EmailValidationResult) SetResult(v EmailValidationStatus) {
	o.Result = &v
}

// GetPredictedScore returns the PredictedScore field value if set, zero value otherwise.
func (o *EmailValidationResult) GetPredictedScore() float32 {
	if o == nil || IsNil(o.PredictedScore) {
		var ret float32
		return ret
	}
	return *o.PredictedScore
}

// GetPredictedScoreOk returns a tuple with the PredictedScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetPredictedScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.PredictedScore) {
		return nil, false
	}
	return o.PredictedScore, true
}

// HasPredictedScore returns a boolean if a field has been set.
func (o *EmailValidationResult) HasPredictedScore() bool {
	if o != nil && !IsNil(o.PredictedScore) {
		return true
	}

	return false
}

// SetPredictedScore gets a reference to the given float32 and assigns it to the PredictedScore field.
func (o *EmailValidationResult) SetPredictedScore(v float32) {
	o.PredictedScore = &v
}

// GetPredictedStatus returns the PredictedStatus field value if set, zero value otherwise.
func (o *EmailValidationResult) GetPredictedStatus() EmailPredictedValidationStatus {
	if o == nil || IsNil(o.PredictedStatus) {
		var ret EmailPredictedValidationStatus
		return ret
	}
	return *o.PredictedStatus
}

// GetPredictedStatusOk returns a tuple with the PredictedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailValidationResult) GetPredictedStatusOk() (*EmailPredictedValidationStatus, bool) {
	if o == nil || IsNil(o.PredictedStatus) {
		return nil, false
	}
	return o.PredictedStatus, true
}

// HasPredictedStatus returns a boolean if a field has been set.
func (o *EmailValidationResult) HasPredictedStatus() bool {
	if o != nil && !IsNil(o.PredictedStatus) {
		return true
	}

	return false
}

// SetPredictedStatus gets a reference to the given EmailPredictedValidationStatus and assigns it to the PredictedStatus field.
func (o *EmailValidationResult) SetPredictedStatus(v EmailPredictedValidationStatus) {
	o.PredictedStatus = &v
}

func (o EmailValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["Account"] = o.Account
	}
	if !IsNil(o.Domain) {
		toSerialize["Domain"] = o.Domain
	}
	if !IsNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	if !IsNil(o.SuggestedSpelling) {
		toSerialize["SuggestedSpelling"] = o.SuggestedSpelling
	}
	if !IsNil(o.Disposable) {
		toSerialize["Disposable"] = o.Disposable
	}
	if !IsNil(o.Role) {
		toSerialize["Role"] = o.Role
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	if !IsNil(o.DateAdded) {
		toSerialize["DateAdded"] = o.DateAdded
	}
	if !IsNil(o.Result) {
		toSerialize["Result"] = o.Result
	}
	if !IsNil(o.PredictedScore) {
		toSerialize["PredictedScore"] = o.PredictedScore
	}
	if !IsNil(o.PredictedStatus) {
		toSerialize["PredictedStatus"] = o.PredictedStatus
	}
	return toSerialize, nil
}

type NullableEmailValidationResult struct {
	value *EmailValidationResult
	isSet bool
}

func (v NullableEmailValidationResult) Get() *EmailValidationResult {
	return v.value
}

func (v *NullableEmailValidationResult) Set(val *EmailValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailValidationResult(val *EmailValidationResult) *NullableEmailValidationResult {
	return &NullableEmailValidationResult{value: val, isSet: true}
}

func (v NullableEmailValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


