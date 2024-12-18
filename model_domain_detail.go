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

// checks if the DomainDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainDetail{}

// DomainDetail Domain data, with information about domain records.
type DomainDetail struct {
	// Name of selected domain.
	Domain *string `json:"Domain,omitempty"`
	// True, if domain is used as default. Otherwise, false,
	DefaultDomain *bool `json:"DefaultDomain,omitempty"`
	// True, if SPF record is verified
	Spf *bool `json:"Spf,omitempty"`
	// True, if DKIM record is verified
	Dkim *bool `json:"Dkim,omitempty"`
	// True, if MX record is verified
	MX *bool `json:"MX,omitempty"`
	DMARC *bool `json:"DMARC,omitempty"`
	// True, if tracking CNAME record is verified
	IsRewriteDomainValid *bool `json:"IsRewriteDomainValid,omitempty"`
	// True, if DKIM, SPF, or tracking are still to be verified
	Verify *bool `json:"Verify,omitempty"`
	Type *TrackingType `json:"Type,omitempty"`
	TrackingStatus *TrackingValidationStatus `json:"TrackingStatus,omitempty"`
	CertificateStatus *CertificateValidationStatus `json:"CertificateStatus,omitempty"`
	CertificateValidationError *string `json:"CertificateValidationError,omitempty"`
	TrackingTypeUserRequest *TrackingType `json:"TrackingTypeUserRequest,omitempty"`
	VERP *bool `json:"VERP,omitempty"`
	CustomBouncesDomain *string `json:"CustomBouncesDomain,omitempty"`
	IsCustomBouncesDomainDefault *bool `json:"IsCustomBouncesDomainDefault,omitempty"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion,omitempty"`
	Ownership *DomainOwner `json:"Ownership,omitempty"`
}

// NewDomainDetail instantiates a new DomainDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainDetail() *DomainDetail {
	this := DomainDetail{}
	var type_ TrackingType = NONE
	this.Type = &type_
	var trackingStatus TrackingValidationStatus = VALIDATED
	this.TrackingStatus = &trackingStatus
	var certificateStatus CertificateValidationStatus = ERROR_OCCURED
	this.CertificateStatus = &certificateStatus
	var trackingTypeUserRequest TrackingType = NONE
	this.TrackingTypeUserRequest = &trackingTypeUserRequest
	var ownership DomainOwner = CURRENT
	this.Ownership = &ownership
	return &this
}

// NewDomainDetailWithDefaults instantiates a new DomainDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainDetailWithDefaults() *DomainDetail {
	this := DomainDetail{}
	var type_ TrackingType = NONE
	this.Type = &type_
	var trackingStatus TrackingValidationStatus = VALIDATED
	this.TrackingStatus = &trackingStatus
	var certificateStatus CertificateValidationStatus = ERROR_OCCURED
	this.CertificateStatus = &certificateStatus
	var trackingTypeUserRequest TrackingType = NONE
	this.TrackingTypeUserRequest = &trackingTypeUserRequest
	var ownership DomainOwner = CURRENT
	this.Ownership = &ownership
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *DomainDetail) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *DomainDetail) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *DomainDetail) SetDomain(v string) {
	o.Domain = &v
}

// GetDefaultDomain returns the DefaultDomain field value if set, zero value otherwise.
func (o *DomainDetail) GetDefaultDomain() bool {
	if o == nil || IsNil(o.DefaultDomain) {
		var ret bool
		return ret
	}
	return *o.DefaultDomain
}

// GetDefaultDomainOk returns a tuple with the DefaultDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDefaultDomainOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultDomain) {
		return nil, false
	}
	return o.DefaultDomain, true
}

// HasDefaultDomain returns a boolean if a field has been set.
func (o *DomainDetail) HasDefaultDomain() bool {
	if o != nil && !IsNil(o.DefaultDomain) {
		return true
	}

	return false
}

// SetDefaultDomain gets a reference to the given bool and assigns it to the DefaultDomain field.
func (o *DomainDetail) SetDefaultDomain(v bool) {
	o.DefaultDomain = &v
}

// GetSpf returns the Spf field value if set, zero value otherwise.
func (o *DomainDetail) GetSpf() bool {
	if o == nil || IsNil(o.Spf) {
		var ret bool
		return ret
	}
	return *o.Spf
}

// GetSpfOk returns a tuple with the Spf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetSpfOk() (*bool, bool) {
	if o == nil || IsNil(o.Spf) {
		return nil, false
	}
	return o.Spf, true
}

// HasSpf returns a boolean if a field has been set.
func (o *DomainDetail) HasSpf() bool {
	if o != nil && !IsNil(o.Spf) {
		return true
	}

	return false
}

// SetSpf gets a reference to the given bool and assigns it to the Spf field.
func (o *DomainDetail) SetSpf(v bool) {
	o.Spf = &v
}

// GetDkim returns the Dkim field value if set, zero value otherwise.
func (o *DomainDetail) GetDkim() bool {
	if o == nil || IsNil(o.Dkim) {
		var ret bool
		return ret
	}
	return *o.Dkim
}

// GetDkimOk returns a tuple with the Dkim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDkimOk() (*bool, bool) {
	if o == nil || IsNil(o.Dkim) {
		return nil, false
	}
	return o.Dkim, true
}

// HasDkim returns a boolean if a field has been set.
func (o *DomainDetail) HasDkim() bool {
	if o != nil && !IsNil(o.Dkim) {
		return true
	}

	return false
}

// SetDkim gets a reference to the given bool and assigns it to the Dkim field.
func (o *DomainDetail) SetDkim(v bool) {
	o.Dkim = &v
}

// GetMX returns the MX field value if set, zero value otherwise.
func (o *DomainDetail) GetMX() bool {
	if o == nil || IsNil(o.MX) {
		var ret bool
		return ret
	}
	return *o.MX
}

// GetMXOk returns a tuple with the MX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetMXOk() (*bool, bool) {
	if o == nil || IsNil(o.MX) {
		return nil, false
	}
	return o.MX, true
}

// HasMX returns a boolean if a field has been set.
func (o *DomainDetail) HasMX() bool {
	if o != nil && !IsNil(o.MX) {
		return true
	}

	return false
}

// SetMX gets a reference to the given bool and assigns it to the MX field.
func (o *DomainDetail) SetMX(v bool) {
	o.MX = &v
}

// GetDMARC returns the DMARC field value if set, zero value otherwise.
func (o *DomainDetail) GetDMARC() bool {
	if o == nil || IsNil(o.DMARC) {
		var ret bool
		return ret
	}
	return *o.DMARC
}

// GetDMARCOk returns a tuple with the DMARC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetDMARCOk() (*bool, bool) {
	if o == nil || IsNil(o.DMARC) {
		return nil, false
	}
	return o.DMARC, true
}

// HasDMARC returns a boolean if a field has been set.
func (o *DomainDetail) HasDMARC() bool {
	if o != nil && !IsNil(o.DMARC) {
		return true
	}

	return false
}

// SetDMARC gets a reference to the given bool and assigns it to the DMARC field.
func (o *DomainDetail) SetDMARC(v bool) {
	o.DMARC = &v
}

// GetIsRewriteDomainValid returns the IsRewriteDomainValid field value if set, zero value otherwise.
func (o *DomainDetail) GetIsRewriteDomainValid() bool {
	if o == nil || IsNil(o.IsRewriteDomainValid) {
		var ret bool
		return ret
	}
	return *o.IsRewriteDomainValid
}

// GetIsRewriteDomainValidOk returns a tuple with the IsRewriteDomainValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetIsRewriteDomainValidOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRewriteDomainValid) {
		return nil, false
	}
	return o.IsRewriteDomainValid, true
}

// HasIsRewriteDomainValid returns a boolean if a field has been set.
func (o *DomainDetail) HasIsRewriteDomainValid() bool {
	if o != nil && !IsNil(o.IsRewriteDomainValid) {
		return true
	}

	return false
}

// SetIsRewriteDomainValid gets a reference to the given bool and assigns it to the IsRewriteDomainValid field.
func (o *DomainDetail) SetIsRewriteDomainValid(v bool) {
	o.IsRewriteDomainValid = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *DomainDetail) GetVerify() bool {
	if o == nil || IsNil(o.Verify) {
		var ret bool
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetVerifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Verify) {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *DomainDetail) HasVerify() bool {
	if o != nil && !IsNil(o.Verify) {
		return true
	}

	return false
}

// SetVerify gets a reference to the given bool and assigns it to the Verify field.
func (o *DomainDetail) SetVerify(v bool) {
	o.Verify = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DomainDetail) GetType() TrackingType {
	if o == nil || IsNil(o.Type) {
		var ret TrackingType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetTypeOk() (*TrackingType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DomainDetail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TrackingType and assigns it to the Type field.
func (o *DomainDetail) SetType(v TrackingType) {
	o.Type = &v
}

// GetTrackingStatus returns the TrackingStatus field value if set, zero value otherwise.
func (o *DomainDetail) GetTrackingStatus() TrackingValidationStatus {
	if o == nil || IsNil(o.TrackingStatus) {
		var ret TrackingValidationStatus
		return ret
	}
	return *o.TrackingStatus
}

// GetTrackingStatusOk returns a tuple with the TrackingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetTrackingStatusOk() (*TrackingValidationStatus, bool) {
	if o == nil || IsNil(o.TrackingStatus) {
		return nil, false
	}
	return o.TrackingStatus, true
}

// HasTrackingStatus returns a boolean if a field has been set.
func (o *DomainDetail) HasTrackingStatus() bool {
	if o != nil && !IsNil(o.TrackingStatus) {
		return true
	}

	return false
}

// SetTrackingStatus gets a reference to the given TrackingValidationStatus and assigns it to the TrackingStatus field.
func (o *DomainDetail) SetTrackingStatus(v TrackingValidationStatus) {
	o.TrackingStatus = &v
}

// GetCertificateStatus returns the CertificateStatus field value if set, zero value otherwise.
func (o *DomainDetail) GetCertificateStatus() CertificateValidationStatus {
	if o == nil || IsNil(o.CertificateStatus) {
		var ret CertificateValidationStatus
		return ret
	}
	return *o.CertificateStatus
}

// GetCertificateStatusOk returns a tuple with the CertificateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetCertificateStatusOk() (*CertificateValidationStatus, bool) {
	if o == nil || IsNil(o.CertificateStatus) {
		return nil, false
	}
	return o.CertificateStatus, true
}

// HasCertificateStatus returns a boolean if a field has been set.
func (o *DomainDetail) HasCertificateStatus() bool {
	if o != nil && !IsNil(o.CertificateStatus) {
		return true
	}

	return false
}

// SetCertificateStatus gets a reference to the given CertificateValidationStatus and assigns it to the CertificateStatus field.
func (o *DomainDetail) SetCertificateStatus(v CertificateValidationStatus) {
	o.CertificateStatus = &v
}

// GetCertificateValidationError returns the CertificateValidationError field value if set, zero value otherwise.
func (o *DomainDetail) GetCertificateValidationError() string {
	if o == nil || IsNil(o.CertificateValidationError) {
		var ret string
		return ret
	}
	return *o.CertificateValidationError
}

// GetCertificateValidationErrorOk returns a tuple with the CertificateValidationError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetCertificateValidationErrorOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateValidationError) {
		return nil, false
	}
	return o.CertificateValidationError, true
}

// HasCertificateValidationError returns a boolean if a field has been set.
func (o *DomainDetail) HasCertificateValidationError() bool {
	if o != nil && !IsNil(o.CertificateValidationError) {
		return true
	}

	return false
}

// SetCertificateValidationError gets a reference to the given string and assigns it to the CertificateValidationError field.
func (o *DomainDetail) SetCertificateValidationError(v string) {
	o.CertificateValidationError = &v
}

// GetTrackingTypeUserRequest returns the TrackingTypeUserRequest field value if set, zero value otherwise.
func (o *DomainDetail) GetTrackingTypeUserRequest() TrackingType {
	if o == nil || IsNil(o.TrackingTypeUserRequest) {
		var ret TrackingType
		return ret
	}
	return *o.TrackingTypeUserRequest
}

// GetTrackingTypeUserRequestOk returns a tuple with the TrackingTypeUserRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetTrackingTypeUserRequestOk() (*TrackingType, bool) {
	if o == nil || IsNil(o.TrackingTypeUserRequest) {
		return nil, false
	}
	return o.TrackingTypeUserRequest, true
}

// HasTrackingTypeUserRequest returns a boolean if a field has been set.
func (o *DomainDetail) HasTrackingTypeUserRequest() bool {
	if o != nil && !IsNil(o.TrackingTypeUserRequest) {
		return true
	}

	return false
}

// SetTrackingTypeUserRequest gets a reference to the given TrackingType and assigns it to the TrackingTypeUserRequest field.
func (o *DomainDetail) SetTrackingTypeUserRequest(v TrackingType) {
	o.TrackingTypeUserRequest = &v
}

// GetVERP returns the VERP field value if set, zero value otherwise.
func (o *DomainDetail) GetVERP() bool {
	if o == nil || IsNil(o.VERP) {
		var ret bool
		return ret
	}
	return *o.VERP
}

// GetVERPOk returns a tuple with the VERP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetVERPOk() (*bool, bool) {
	if o == nil || IsNil(o.VERP) {
		return nil, false
	}
	return o.VERP, true
}

// HasVERP returns a boolean if a field has been set.
func (o *DomainDetail) HasVERP() bool {
	if o != nil && !IsNil(o.VERP) {
		return true
	}

	return false
}

// SetVERP gets a reference to the given bool and assigns it to the VERP field.
func (o *DomainDetail) SetVERP(v bool) {
	o.VERP = &v
}

// GetCustomBouncesDomain returns the CustomBouncesDomain field value if set, zero value otherwise.
func (o *DomainDetail) GetCustomBouncesDomain() string {
	if o == nil || IsNil(o.CustomBouncesDomain) {
		var ret string
		return ret
	}
	return *o.CustomBouncesDomain
}

// GetCustomBouncesDomainOk returns a tuple with the CustomBouncesDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetCustomBouncesDomainOk() (*string, bool) {
	if o == nil || IsNil(o.CustomBouncesDomain) {
		return nil, false
	}
	return o.CustomBouncesDomain, true
}

// HasCustomBouncesDomain returns a boolean if a field has been set.
func (o *DomainDetail) HasCustomBouncesDomain() bool {
	if o != nil && !IsNil(o.CustomBouncesDomain) {
		return true
	}

	return false
}

// SetCustomBouncesDomain gets a reference to the given string and assigns it to the CustomBouncesDomain field.
func (o *DomainDetail) SetCustomBouncesDomain(v string) {
	o.CustomBouncesDomain = &v
}

// GetIsCustomBouncesDomainDefault returns the IsCustomBouncesDomainDefault field value if set, zero value otherwise.
func (o *DomainDetail) GetIsCustomBouncesDomainDefault() bool {
	if o == nil || IsNil(o.IsCustomBouncesDomainDefault) {
		var ret bool
		return ret
	}
	return *o.IsCustomBouncesDomainDefault
}

// GetIsCustomBouncesDomainDefaultOk returns a tuple with the IsCustomBouncesDomainDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetIsCustomBouncesDomainDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCustomBouncesDomainDefault) {
		return nil, false
	}
	return o.IsCustomBouncesDomainDefault, true
}

// HasIsCustomBouncesDomainDefault returns a boolean if a field has been set.
func (o *DomainDetail) HasIsCustomBouncesDomainDefault() bool {
	if o != nil && !IsNil(o.IsCustomBouncesDomainDefault) {
		return true
	}

	return false
}

// SetIsCustomBouncesDomainDefault gets a reference to the given bool and assigns it to the IsCustomBouncesDomainDefault field.
func (o *DomainDetail) SetIsCustomBouncesDomainDefault(v bool) {
	o.IsCustomBouncesDomainDefault = &v
}

// GetIsMarkedForDeletion returns the IsMarkedForDeletion field value if set, zero value otherwise.
func (o *DomainDetail) GetIsMarkedForDeletion() bool {
	if o == nil || IsNil(o.IsMarkedForDeletion) {
		var ret bool
		return ret
	}
	return *o.IsMarkedForDeletion
}

// GetIsMarkedForDeletionOk returns a tuple with the IsMarkedForDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetIsMarkedForDeletionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMarkedForDeletion) {
		return nil, false
	}
	return o.IsMarkedForDeletion, true
}

// HasIsMarkedForDeletion returns a boolean if a field has been set.
func (o *DomainDetail) HasIsMarkedForDeletion() bool {
	if o != nil && !IsNil(o.IsMarkedForDeletion) {
		return true
	}

	return false
}

// SetIsMarkedForDeletion gets a reference to the given bool and assigns it to the IsMarkedForDeletion field.
func (o *DomainDetail) SetIsMarkedForDeletion(v bool) {
	o.IsMarkedForDeletion = &v
}

// GetOwnership returns the Ownership field value if set, zero value otherwise.
func (o *DomainDetail) GetOwnership() DomainOwner {
	if o == nil || IsNil(o.Ownership) {
		var ret DomainOwner
		return ret
	}
	return *o.Ownership
}

// GetOwnershipOk returns a tuple with the Ownership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDetail) GetOwnershipOk() (*DomainOwner, bool) {
	if o == nil || IsNil(o.Ownership) {
		return nil, false
	}
	return o.Ownership, true
}

// HasOwnership returns a boolean if a field has been set.
func (o *DomainDetail) HasOwnership() bool {
	if o != nil && !IsNil(o.Ownership) {
		return true
	}

	return false
}

// SetOwnership gets a reference to the given DomainOwner and assigns it to the Ownership field.
func (o *DomainDetail) SetOwnership(v DomainOwner) {
	o.Ownership = &v
}

func (o DomainDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["Domain"] = o.Domain
	}
	if !IsNil(o.DefaultDomain) {
		toSerialize["DefaultDomain"] = o.DefaultDomain
	}
	if !IsNil(o.Spf) {
		toSerialize["Spf"] = o.Spf
	}
	if !IsNil(o.Dkim) {
		toSerialize["Dkim"] = o.Dkim
	}
	if !IsNil(o.MX) {
		toSerialize["MX"] = o.MX
	}
	if !IsNil(o.DMARC) {
		toSerialize["DMARC"] = o.DMARC
	}
	if !IsNil(o.IsRewriteDomainValid) {
		toSerialize["IsRewriteDomainValid"] = o.IsRewriteDomainValid
	}
	if !IsNil(o.Verify) {
		toSerialize["Verify"] = o.Verify
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.TrackingStatus) {
		toSerialize["TrackingStatus"] = o.TrackingStatus
	}
	if !IsNil(o.CertificateStatus) {
		toSerialize["CertificateStatus"] = o.CertificateStatus
	}
	if !IsNil(o.CertificateValidationError) {
		toSerialize["CertificateValidationError"] = o.CertificateValidationError
	}
	if !IsNil(o.TrackingTypeUserRequest) {
		toSerialize["TrackingTypeUserRequest"] = o.TrackingTypeUserRequest
	}
	if !IsNil(o.VERP) {
		toSerialize["VERP"] = o.VERP
	}
	if !IsNil(o.CustomBouncesDomain) {
		toSerialize["CustomBouncesDomain"] = o.CustomBouncesDomain
	}
	if !IsNil(o.IsCustomBouncesDomainDefault) {
		toSerialize["IsCustomBouncesDomainDefault"] = o.IsCustomBouncesDomainDefault
	}
	if !IsNil(o.IsMarkedForDeletion) {
		toSerialize["IsMarkedForDeletion"] = o.IsMarkedForDeletion
	}
	if !IsNil(o.Ownership) {
		toSerialize["Ownership"] = o.Ownership
	}
	return toSerialize, nil
}

type NullableDomainDetail struct {
	value *DomainDetail
	isSet bool
}

func (v NullableDomainDetail) Get() *DomainDetail {
	return v.value
}

func (v *NullableDomainDetail) Set(val *DomainDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainDetail(val *DomainDetail) *NullableDomainDetail {
	return &NullableDomainDetail{value: val, isSet: true}
}

func (v NullableDomainDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


