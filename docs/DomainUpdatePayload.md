# DomainUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStatus** | Pointer to [**CertificateValidationStatus**](CertificateValidationStatus.md) |  | [optional] [default to CERTIFICATEVALIDATIONSTATUS_ERROR_OCCURED]
**VERP** | Pointer to **bool** |  | [optional] 
**CustomBouncesDomain** | Pointer to **string** |  | [optional] 
**IsCustomBouncesDomainDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewDomainUpdatePayload

`func NewDomainUpdatePayload() *DomainUpdatePayload`

NewDomainUpdatePayload instantiates a new DomainUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUpdatePayloadWithDefaults

`func NewDomainUpdatePayloadWithDefaults() *DomainUpdatePayload`

NewDomainUpdatePayloadWithDefaults instantiates a new DomainUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStatus

`func (o *DomainUpdatePayload) GetCertificateStatus() CertificateValidationStatus`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *DomainUpdatePayload) GetCertificateStatusOk() (*CertificateValidationStatus, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *DomainUpdatePayload) SetCertificateStatus(v CertificateValidationStatus)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *DomainUpdatePayload) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetVERP

`func (o *DomainUpdatePayload) GetVERP() bool`

GetVERP returns the VERP field if non-nil, zero value otherwise.

### GetVERPOk

`func (o *DomainUpdatePayload) GetVERPOk() (*bool, bool)`

GetVERPOk returns a tuple with the VERP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVERP

`func (o *DomainUpdatePayload) SetVERP(v bool)`

SetVERP sets VERP field to given value.

### HasVERP

`func (o *DomainUpdatePayload) HasVERP() bool`

HasVERP returns a boolean if a field has been set.

### GetCustomBouncesDomain

`func (o *DomainUpdatePayload) GetCustomBouncesDomain() string`

GetCustomBouncesDomain returns the CustomBouncesDomain field if non-nil, zero value otherwise.

### GetCustomBouncesDomainOk

`func (o *DomainUpdatePayload) GetCustomBouncesDomainOk() (*string, bool)`

GetCustomBouncesDomainOk returns a tuple with the CustomBouncesDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBouncesDomain

`func (o *DomainUpdatePayload) SetCustomBouncesDomain(v string)`

SetCustomBouncesDomain sets CustomBouncesDomain field to given value.

### HasCustomBouncesDomain

`func (o *DomainUpdatePayload) HasCustomBouncesDomain() bool`

HasCustomBouncesDomain returns a boolean if a field has been set.

### GetIsCustomBouncesDomainDefault

`func (o *DomainUpdatePayload) GetIsCustomBouncesDomainDefault() bool`

GetIsCustomBouncesDomainDefault returns the IsCustomBouncesDomainDefault field if non-nil, zero value otherwise.

### GetIsCustomBouncesDomainDefaultOk

`func (o *DomainUpdatePayload) GetIsCustomBouncesDomainDefaultOk() (*bool, bool)`

GetIsCustomBouncesDomainDefaultOk returns a tuple with the IsCustomBouncesDomainDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomBouncesDomainDefault

`func (o *DomainUpdatePayload) SetIsCustomBouncesDomainDefault(v bool)`

SetIsCustomBouncesDomainDefault sets IsCustomBouncesDomainDefault field to given value.

### HasIsCustomBouncesDomainDefault

`func (o *DomainUpdatePayload) HasIsCustomBouncesDomainDefault() bool`

HasIsCustomBouncesDomainDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


