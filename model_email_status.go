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
	"time"
)

// EmailStatus Status information of the specified email
type EmailStatus struct {
	// Email address this email was sent from.
	From *string `json:"From,omitempty"`
	// Email address this email was sent to.
	To *string `json:"To,omitempty"`
	// Date the email was submitted.
	Date *time.Time `json:"Date,omitempty"`
	Status *LogJobStatus `json:"Status,omitempty"`
	// Name of email's status
	StatusName *string `json:"StatusName,omitempty"`
	// Date of last status change.
	StatusChangeDate *time.Time `json:"StatusChangeDate,omitempty"`
	// Date when the email was sent
	DateSent *time.Time `json:"DateSent,omitempty"`
	// Date when the email changed the status to 'opened'
	DateOpened NullableTime `json:"DateOpened,omitempty"`
	// Date when the email changed the status to 'clicked'
	DateClicked NullableTime `json:"DateClicked,omitempty"`
	// Detailed error or bounced message.
	ErrorMessage *string `json:"ErrorMessage,omitempty"`
	// ID number of transaction
	TransactionID *string `json:"TransactionID,omitempty"`
	// Envelope from address
	EnvelopeFrom *string `json:"EnvelopeFrom,omitempty"`
}

// NewEmailStatus instantiates a new EmailStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailStatus() *EmailStatus {
	this := EmailStatus{}
	var status LogJobStatus = ALL
	this.Status = &status
	return &this
}

// NewEmailStatusWithDefaults instantiates a new EmailStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailStatusWithDefaults() *EmailStatus {
	this := EmailStatus{}
	var status LogJobStatus = ALL
	this.Status = &status
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *EmailStatus) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *EmailStatus) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *EmailStatus) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *EmailStatus) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *EmailStatus) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *EmailStatus) SetTo(v string) {
	o.To = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *EmailStatus) GetDate() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetDateOk() (*time.Time, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *EmailStatus) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *EmailStatus) SetDate(v time.Time) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EmailStatus) GetStatus() LogJobStatus {
	if o == nil || o.Status == nil {
		var ret LogJobStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetStatusOk() (*LogJobStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EmailStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given LogJobStatus and assigns it to the Status field.
func (o *EmailStatus) SetStatus(v LogJobStatus) {
	o.Status = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *EmailStatus) GetStatusName() string {
	if o == nil || o.StatusName == nil {
		var ret string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetStatusNameOk() (*string, bool) {
	if o == nil || o.StatusName == nil {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *EmailStatus) HasStatusName() bool {
	if o != nil && o.StatusName != nil {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given string and assigns it to the StatusName field.
func (o *EmailStatus) SetStatusName(v string) {
	o.StatusName = &v
}

// GetStatusChangeDate returns the StatusChangeDate field value if set, zero value otherwise.
func (o *EmailStatus) GetStatusChangeDate() time.Time {
	if o == nil || o.StatusChangeDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StatusChangeDate
}

// GetStatusChangeDateOk returns a tuple with the StatusChangeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetStatusChangeDateOk() (*time.Time, bool) {
	if o == nil || o.StatusChangeDate == nil {
		return nil, false
	}
	return o.StatusChangeDate, true
}

// HasStatusChangeDate returns a boolean if a field has been set.
func (o *EmailStatus) HasStatusChangeDate() bool {
	if o != nil && o.StatusChangeDate != nil {
		return true
	}

	return false
}

// SetStatusChangeDate gets a reference to the given time.Time and assigns it to the StatusChangeDate field.
func (o *EmailStatus) SetStatusChangeDate(v time.Time) {
	o.StatusChangeDate = &v
}

// GetDateSent returns the DateSent field value if set, zero value otherwise.
func (o *EmailStatus) GetDateSent() time.Time {
	if o == nil || o.DateSent == nil {
		var ret time.Time
		return ret
	}
	return *o.DateSent
}

// GetDateSentOk returns a tuple with the DateSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetDateSentOk() (*time.Time, bool) {
	if o == nil || o.DateSent == nil {
		return nil, false
	}
	return o.DateSent, true
}

// HasDateSent returns a boolean if a field has been set.
func (o *EmailStatus) HasDateSent() bool {
	if o != nil && o.DateSent != nil {
		return true
	}

	return false
}

// SetDateSent gets a reference to the given time.Time and assigns it to the DateSent field.
func (o *EmailStatus) SetDateSent(v time.Time) {
	o.DateSent = &v
}

// GetDateOpened returns the DateOpened field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailStatus) GetDateOpened() time.Time {
	if o == nil || o.DateOpened.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DateOpened.Get()
}

// GetDateOpenedOk returns a tuple with the DateOpened field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailStatus) GetDateOpenedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateOpened.Get(), o.DateOpened.IsSet()
}

// HasDateOpened returns a boolean if a field has been set.
func (o *EmailStatus) HasDateOpened() bool {
	if o != nil && o.DateOpened.IsSet() {
		return true
	}

	return false
}

// SetDateOpened gets a reference to the given NullableTime and assigns it to the DateOpened field.
func (o *EmailStatus) SetDateOpened(v time.Time) {
	o.DateOpened.Set(&v)
}
// SetDateOpenedNil sets the value for DateOpened to be an explicit nil
func (o *EmailStatus) SetDateOpenedNil() {
	o.DateOpened.Set(nil)
}

// UnsetDateOpened ensures that no value is present for DateOpened, not even an explicit nil
func (o *EmailStatus) UnsetDateOpened() {
	o.DateOpened.Unset()
}

// GetDateClicked returns the DateClicked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmailStatus) GetDateClicked() time.Time {
	if o == nil || o.DateClicked.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DateClicked.Get()
}

// GetDateClickedOk returns a tuple with the DateClicked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmailStatus) GetDateClickedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateClicked.Get(), o.DateClicked.IsSet()
}

// HasDateClicked returns a boolean if a field has been set.
func (o *EmailStatus) HasDateClicked() bool {
	if o != nil && o.DateClicked.IsSet() {
		return true
	}

	return false
}

// SetDateClicked gets a reference to the given NullableTime and assigns it to the DateClicked field.
func (o *EmailStatus) SetDateClicked(v time.Time) {
	o.DateClicked.Set(&v)
}
// SetDateClickedNil sets the value for DateClicked to be an explicit nil
func (o *EmailStatus) SetDateClickedNil() {
	o.DateClicked.Set(nil)
}

// UnsetDateClicked ensures that no value is present for DateClicked, not even an explicit nil
func (o *EmailStatus) UnsetDateClicked() {
	o.DateClicked.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *EmailStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EmailStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *EmailStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetTransactionID returns the TransactionID field value if set, zero value otherwise.
func (o *EmailStatus) GetTransactionID() string {
	if o == nil || o.TransactionID == nil {
		var ret string
		return ret
	}
	return *o.TransactionID
}

// GetTransactionIDOk returns a tuple with the TransactionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetTransactionIDOk() (*string, bool) {
	if o == nil || o.TransactionID == nil {
		return nil, false
	}
	return o.TransactionID, true
}

// HasTransactionID returns a boolean if a field has been set.
func (o *EmailStatus) HasTransactionID() bool {
	if o != nil && o.TransactionID != nil {
		return true
	}

	return false
}

// SetTransactionID gets a reference to the given string and assigns it to the TransactionID field.
func (o *EmailStatus) SetTransactionID(v string) {
	o.TransactionID = &v
}

// GetEnvelopeFrom returns the EnvelopeFrom field value if set, zero value otherwise.
func (o *EmailStatus) GetEnvelopeFrom() string {
	if o == nil || o.EnvelopeFrom == nil {
		var ret string
		return ret
	}
	return *o.EnvelopeFrom
}

// GetEnvelopeFromOk returns a tuple with the EnvelopeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailStatus) GetEnvelopeFromOk() (*string, bool) {
	if o == nil || o.EnvelopeFrom == nil {
		return nil, false
	}
	return o.EnvelopeFrom, true
}

// HasEnvelopeFrom returns a boolean if a field has been set.
func (o *EmailStatus) HasEnvelopeFrom() bool {
	if o != nil && o.EnvelopeFrom != nil {
		return true
	}

	return false
}

// SetEnvelopeFrom gets a reference to the given string and assigns it to the EnvelopeFrom field.
func (o *EmailStatus) SetEnvelopeFrom(v string) {
	o.EnvelopeFrom = &v
}

func (o EmailStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.From != nil {
		toSerialize["From"] = o.From
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}
	if o.Date != nil {
		toSerialize["Date"] = o.Date
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusName != nil {
		toSerialize["StatusName"] = o.StatusName
	}
	if o.StatusChangeDate != nil {
		toSerialize["StatusChangeDate"] = o.StatusChangeDate
	}
	if o.DateSent != nil {
		toSerialize["DateSent"] = o.DateSent
	}
	if o.DateOpened.IsSet() {
		toSerialize["DateOpened"] = o.DateOpened.Get()
	}
	if o.DateClicked.IsSet() {
		toSerialize["DateClicked"] = o.DateClicked.Get()
	}
	if o.ErrorMessage != nil {
		toSerialize["ErrorMessage"] = o.ErrorMessage
	}
	if o.TransactionID != nil {
		toSerialize["TransactionID"] = o.TransactionID
	}
	if o.EnvelopeFrom != nil {
		toSerialize["EnvelopeFrom"] = o.EnvelopeFrom
	}
	return json.Marshal(toSerialize)
}

type NullableEmailStatus struct {
	value *EmailStatus
	isSet bool
}

func (v NullableEmailStatus) Get() *EmailStatus {
	return v.value
}

func (v *NullableEmailStatus) Set(val *EmailStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailStatus(val *EmailStatus) *NullableEmailStatus {
	return &NullableEmailStatus{value: val, isSet: true}
}

func (v NullableEmailStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


