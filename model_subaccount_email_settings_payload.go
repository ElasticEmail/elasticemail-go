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
)

// SubaccountEmailSettingsPayload Settings related to sending emails
type SubaccountEmailSettingsPayload struct {
	// True, if Account needs credits to send emails. Otherwise, false
	RequiresEmailCredits *bool `json:"RequiresEmailCredits,omitempty"`
	// Maximum size of email including attachments in MB's
	EmailSizeLimit *int32 `json:"EmailSizeLimit,omitempty"`
	// Amount of emails Account can send daily
	DailySendLimit *int32 `json:"DailySendLimit,omitempty"`
	// Maximum number of contacts the Account can have. 0 means that parent account's limit is used.
	MaxContacts *int32 `json:"MaxContacts,omitempty"`
	// Can the SubAccount purchase Private IP for themselves
	EnablePrivateIPPurchase *bool `json:"EnablePrivateIPPurchase,omitempty"`
	// Name of your custom IP Pool to be used in the sending process
	PoolName *string `json:"PoolName,omitempty"`
	ValidSenderDomainOnly NullableBool `json:"ValidSenderDomainOnly,omitempty"`
}

// NewSubaccountEmailSettingsPayload instantiates a new SubaccountEmailSettingsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountEmailSettingsPayload() *SubaccountEmailSettingsPayload {
	this := SubaccountEmailSettingsPayload{}
	return &this
}

// NewSubaccountEmailSettingsPayloadWithDefaults instantiates a new SubaccountEmailSettingsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountEmailSettingsPayloadWithDefaults() *SubaccountEmailSettingsPayload {
	this := SubaccountEmailSettingsPayload{}
	return &this
}

// GetRequiresEmailCredits returns the RequiresEmailCredits field value if set, zero value otherwise.
func (o *SubaccountEmailSettingsPayload) GetRequiresEmailCredits() bool {
	if o == nil || o.RequiresEmailCredits == nil {
		var ret bool
		return ret
	}
	return *o.RequiresEmailCredits
}

// GetRequiresEmailCreditsOk returns a tuple with the RequiresEmailCredits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettingsPayload) GetRequiresEmailCreditsOk() (*bool, bool) {
	if o == nil || o.RequiresEmailCredits == nil {
		return nil, false
	}
	return o.RequiresEmailCredits, true
}

// HasRequiresEmailCredits returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasRequiresEmailCredits() bool {
	if o != nil && o.RequiresEmailCredits != nil {
		return true
	}

	return false
}

// SetRequiresEmailCredits gets a reference to the given bool and assigns it to the RequiresEmailCredits field.
func (o *SubaccountEmailSettingsPayload) SetRequiresEmailCredits(v bool) {
	o.RequiresEmailCredits = &v
}

// GetEmailSizeLimit returns the EmailSizeLimit field value if set, zero value otherwise.
func (o *SubaccountEmailSettingsPayload) GetEmailSizeLimit() int32 {
	if o == nil || o.EmailSizeLimit == nil {
		var ret int32
		return ret
	}
	return *o.EmailSizeLimit
}

// GetEmailSizeLimitOk returns a tuple with the EmailSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettingsPayload) GetEmailSizeLimitOk() (*int32, bool) {
	if o == nil || o.EmailSizeLimit == nil {
		return nil, false
	}
	return o.EmailSizeLimit, true
}

// HasEmailSizeLimit returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasEmailSizeLimit() bool {
	if o != nil && o.EmailSizeLimit != nil {
		return true
	}

	return false
}

// SetEmailSizeLimit gets a reference to the given int32 and assigns it to the EmailSizeLimit field.
func (o *SubaccountEmailSettingsPayload) SetEmailSizeLimit(v int32) {
	o.EmailSizeLimit = &v
}

// GetDailySendLimit returns the DailySendLimit field value if set, zero value otherwise.
func (o *SubaccountEmailSettingsPayload) GetDailySendLimit() int32 {
	if o == nil || o.DailySendLimit == nil {
		var ret int32
		return ret
	}
	return *o.DailySendLimit
}

// GetDailySendLimitOk returns a tuple with the DailySendLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettingsPayload) GetDailySendLimitOk() (*int32, bool) {
	if o == nil || o.DailySendLimit == nil {
		return nil, false
	}
	return o.DailySendLimit, true
}

// HasDailySendLimit returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasDailySendLimit() bool {
	if o != nil && o.DailySendLimit != nil {
		return true
	}

	return false
}

// SetDailySendLimit gets a reference to the given int32 and assigns it to the DailySendLimit field.
func (o *SubaccountEmailSettingsPayload) SetDailySendLimit(v int32) {
	o.DailySendLimit = &v
}

// GetMaxContacts returns the MaxContacts field value if set, zero value otherwise.
func (o *SubaccountEmailSettingsPayload) GetMaxContacts() int32 {
	if o == nil || o.MaxContacts == nil {
		var ret int32
		return ret
	}
	return *o.MaxContacts
}

// GetMaxContactsOk returns a tuple with the MaxContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettingsPayload) GetMaxContactsOk() (*int32, bool) {
	if o == nil || o.MaxContacts == nil {
		return nil, false
	}
	return o.MaxContacts, true
}

// HasMaxContacts returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasMaxContacts() bool {
	if o != nil && o.MaxContacts != nil {
		return true
	}

	return false
}

// SetMaxContacts gets a reference to the given int32 and assigns it to the MaxContacts field.
func (o *SubaccountEmailSettingsPayload) SetMaxContacts(v int32) {
	o.MaxContacts = &v
}

// GetEnablePrivateIPPurchase returns the EnablePrivateIPPurchase field value if set, zero value otherwise.
func (o *SubaccountEmailSettingsPayload) GetEnablePrivateIPPurchase() bool {
	if o == nil || o.EnablePrivateIPPurchase == nil {
		var ret bool
		return ret
	}
	return *o.EnablePrivateIPPurchase
}

// GetEnablePrivateIPPurchaseOk returns a tuple with the EnablePrivateIPPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettingsPayload) GetEnablePrivateIPPurchaseOk() (*bool, bool) {
	if o == nil || o.EnablePrivateIPPurchase == nil {
		return nil, false
	}
	return o.EnablePrivateIPPurchase, true
}

// HasEnablePrivateIPPurchase returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasEnablePrivateIPPurchase() bool {
	if o != nil && o.EnablePrivateIPPurchase != nil {
		return true
	}

	return false
}

// SetEnablePrivateIPPurchase gets a reference to the given bool and assigns it to the EnablePrivateIPPurchase field.
func (o *SubaccountEmailSettingsPayload) SetEnablePrivateIPPurchase(v bool) {
	o.EnablePrivateIPPurchase = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *SubaccountEmailSettingsPayload) GetPoolName() string {
	if o == nil || o.PoolName == nil {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountEmailSettingsPayload) GetPoolNameOk() (*string, bool) {
	if o == nil || o.PoolName == nil {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasPoolName() bool {
	if o != nil && o.PoolName != nil {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *SubaccountEmailSettingsPayload) SetPoolName(v string) {
	o.PoolName = &v
}

// GetValidSenderDomainOnly returns the ValidSenderDomainOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubaccountEmailSettingsPayload) GetValidSenderDomainOnly() bool {
	if o == nil || o.ValidSenderDomainOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ValidSenderDomainOnly.Get()
}

// GetValidSenderDomainOnlyOk returns a tuple with the ValidSenderDomainOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubaccountEmailSettingsPayload) GetValidSenderDomainOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidSenderDomainOnly.Get(), o.ValidSenderDomainOnly.IsSet()
}

// HasValidSenderDomainOnly returns a boolean if a field has been set.
func (o *SubaccountEmailSettingsPayload) HasValidSenderDomainOnly() bool {
	if o != nil && o.ValidSenderDomainOnly.IsSet() {
		return true
	}

	return false
}

// SetValidSenderDomainOnly gets a reference to the given NullableBool and assigns it to the ValidSenderDomainOnly field.
func (o *SubaccountEmailSettingsPayload) SetValidSenderDomainOnly(v bool) {
	o.ValidSenderDomainOnly.Set(&v)
}
// SetValidSenderDomainOnlyNil sets the value for ValidSenderDomainOnly to be an explicit nil
func (o *SubaccountEmailSettingsPayload) SetValidSenderDomainOnlyNil() {
	o.ValidSenderDomainOnly.Set(nil)
}

// UnsetValidSenderDomainOnly ensures that no value is present for ValidSenderDomainOnly, not even an explicit nil
func (o *SubaccountEmailSettingsPayload) UnsetValidSenderDomainOnly() {
	o.ValidSenderDomainOnly.Unset()
}

func (o SubaccountEmailSettingsPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequiresEmailCredits != nil {
		toSerialize["RequiresEmailCredits"] = o.RequiresEmailCredits
	}
	if o.EmailSizeLimit != nil {
		toSerialize["EmailSizeLimit"] = o.EmailSizeLimit
	}
	if o.DailySendLimit != nil {
		toSerialize["DailySendLimit"] = o.DailySendLimit
	}
	if o.MaxContacts != nil {
		toSerialize["MaxContacts"] = o.MaxContacts
	}
	if o.EnablePrivateIPPurchase != nil {
		toSerialize["EnablePrivateIPPurchase"] = o.EnablePrivateIPPurchase
	}
	if o.PoolName != nil {
		toSerialize["PoolName"] = o.PoolName
	}
	if o.ValidSenderDomainOnly.IsSet() {
		toSerialize["ValidSenderDomainOnly"] = o.ValidSenderDomainOnly.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSubaccountEmailSettingsPayload struct {
	value *SubaccountEmailSettingsPayload
	isSet bool
}

func (v NullableSubaccountEmailSettingsPayload) Get() *SubaccountEmailSettingsPayload {
	return v.value
}

func (v *NullableSubaccountEmailSettingsPayload) Set(val *SubaccountEmailSettingsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountEmailSettingsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountEmailSettingsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountEmailSettingsPayload(val *SubaccountEmailSettingsPayload) *NullableSubaccountEmailSettingsPayload {
	return &NullableSubaccountEmailSettingsPayload{value: val, isSet: true}
}

func (v NullableSubaccountEmailSettingsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountEmailSettingsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


