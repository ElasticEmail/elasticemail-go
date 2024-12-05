# DomainData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationLog** | Pointer to **string** | Domain validation results - when domain has been running through validation process | [optional] 
**Domain** | Pointer to **string** | Name of selected domain. | [optional] 
**DefaultDomain** | Pointer to **bool** | True, if domain is used as default. Otherwise, false, | [optional] 
**Spf** | Pointer to **bool** | True, if SPF record is verified | [optional] 
**Dkim** | Pointer to **bool** | True, if DKIM record is verified | [optional] 
**MX** | Pointer to **bool** | True, if MX record is verified | [optional] 
**DMARC** | Pointer to **bool** |  | [optional] 
**IsRewriteDomainValid** | Pointer to **bool** | True, if tracking CNAME record is verified | [optional] 
**Verify** | Pointer to **bool** | True, if DKIM, SPF, or tracking are still to be verified | [optional] 
**Type** | Pointer to [**TrackingType**](TrackingType.md) |  | [optional] [default to NONE]
**TrackingStatus** | Pointer to [**TrackingValidationStatus**](TrackingValidationStatus.md) |  | [optional] [default to VALIDATED]
**CertificateStatus** | Pointer to [**CertificateValidationStatus**](CertificateValidationStatus.md) |  | [optional] [default to ERROR_OCCURED]
**CertificateValidationError** | Pointer to **string** |  | [optional] 
**TrackingTypeUserRequest** | Pointer to [**TrackingType**](TrackingType.md) |  | [optional] [default to NONE]
**VERP** | Pointer to **bool** |  | [optional] 
**CustomBouncesDomain** | Pointer to **string** |  | [optional] 
**IsCustomBouncesDomainDefault** | Pointer to **bool** |  | [optional] 
**IsMarkedForDeletion** | Pointer to **bool** |  | [optional] 
**Ownership** | Pointer to [**DomainOwner**](DomainOwner.md) |  | [optional] [default to CURRENT]

## Methods

### NewDomainData

`func NewDomainData() *DomainData`

NewDomainData instantiates a new DomainData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDataWithDefaults

`func NewDomainDataWithDefaults() *DomainData`

NewDomainDataWithDefaults instantiates a new DomainData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationLog

`func (o *DomainData) GetValidationLog() string`

GetValidationLog returns the ValidationLog field if non-nil, zero value otherwise.

### GetValidationLogOk

`func (o *DomainData) GetValidationLogOk() (*string, bool)`

GetValidationLogOk returns a tuple with the ValidationLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationLog

`func (o *DomainData) SetValidationLog(v string)`

SetValidationLog sets ValidationLog field to given value.

### HasValidationLog

`func (o *DomainData) HasValidationLog() bool`

HasValidationLog returns a boolean if a field has been set.

### GetDomain

`func (o *DomainData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDefaultDomain

`func (o *DomainData) GetDefaultDomain() bool`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *DomainData) GetDefaultDomainOk() (*bool, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *DomainData) SetDefaultDomain(v bool)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *DomainData) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetSpf

`func (o *DomainData) GetSpf() bool`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *DomainData) GetSpfOk() (*bool, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *DomainData) SetSpf(v bool)`

SetSpf sets Spf field to given value.

### HasSpf

`func (o *DomainData) HasSpf() bool`

HasSpf returns a boolean if a field has been set.

### GetDkim

`func (o *DomainData) GetDkim() bool`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *DomainData) GetDkimOk() (*bool, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *DomainData) SetDkim(v bool)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *DomainData) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetMX

`func (o *DomainData) GetMX() bool`

GetMX returns the MX field if non-nil, zero value otherwise.

### GetMXOk

`func (o *DomainData) GetMXOk() (*bool, bool)`

GetMXOk returns a tuple with the MX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMX

`func (o *DomainData) SetMX(v bool)`

SetMX sets MX field to given value.

### HasMX

`func (o *DomainData) HasMX() bool`

HasMX returns a boolean if a field has been set.

### GetDMARC

`func (o *DomainData) GetDMARC() bool`

GetDMARC returns the DMARC field if non-nil, zero value otherwise.

### GetDMARCOk

`func (o *DomainData) GetDMARCOk() (*bool, bool)`

GetDMARCOk returns a tuple with the DMARC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMARC

`func (o *DomainData) SetDMARC(v bool)`

SetDMARC sets DMARC field to given value.

### HasDMARC

`func (o *DomainData) HasDMARC() bool`

HasDMARC returns a boolean if a field has been set.

### GetIsRewriteDomainValid

`func (o *DomainData) GetIsRewriteDomainValid() bool`

GetIsRewriteDomainValid returns the IsRewriteDomainValid field if non-nil, zero value otherwise.

### GetIsRewriteDomainValidOk

`func (o *DomainData) GetIsRewriteDomainValidOk() (*bool, bool)`

GetIsRewriteDomainValidOk returns a tuple with the IsRewriteDomainValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRewriteDomainValid

`func (o *DomainData) SetIsRewriteDomainValid(v bool)`

SetIsRewriteDomainValid sets IsRewriteDomainValid field to given value.

### HasIsRewriteDomainValid

`func (o *DomainData) HasIsRewriteDomainValid() bool`

HasIsRewriteDomainValid returns a boolean if a field has been set.

### GetVerify

`func (o *DomainData) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *DomainData) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *DomainData) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *DomainData) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetType

`func (o *DomainData) GetType() TrackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainData) GetTypeOk() (*TrackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainData) SetType(v TrackingType)`

SetType sets Type field to given value.

### HasType

`func (o *DomainData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTrackingStatus

`func (o *DomainData) GetTrackingStatus() TrackingValidationStatus`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *DomainData) GetTrackingStatusOk() (*TrackingValidationStatus, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *DomainData) SetTrackingStatus(v TrackingValidationStatus)`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *DomainData) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### GetCertificateStatus

`func (o *DomainData) GetCertificateStatus() CertificateValidationStatus`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *DomainData) GetCertificateStatusOk() (*CertificateValidationStatus, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *DomainData) SetCertificateStatus(v CertificateValidationStatus)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *DomainData) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetCertificateValidationError

`func (o *DomainData) GetCertificateValidationError() string`

GetCertificateValidationError returns the CertificateValidationError field if non-nil, zero value otherwise.

### GetCertificateValidationErrorOk

`func (o *DomainData) GetCertificateValidationErrorOk() (*string, bool)`

GetCertificateValidationErrorOk returns a tuple with the CertificateValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidationError

`func (o *DomainData) SetCertificateValidationError(v string)`

SetCertificateValidationError sets CertificateValidationError field to given value.

### HasCertificateValidationError

`func (o *DomainData) HasCertificateValidationError() bool`

HasCertificateValidationError returns a boolean if a field has been set.

### GetTrackingTypeUserRequest

`func (o *DomainData) GetTrackingTypeUserRequest() TrackingType`

GetTrackingTypeUserRequest returns the TrackingTypeUserRequest field if non-nil, zero value otherwise.

### GetTrackingTypeUserRequestOk

`func (o *DomainData) GetTrackingTypeUserRequestOk() (*TrackingType, bool)`

GetTrackingTypeUserRequestOk returns a tuple with the TrackingTypeUserRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingTypeUserRequest

`func (o *DomainData) SetTrackingTypeUserRequest(v TrackingType)`

SetTrackingTypeUserRequest sets TrackingTypeUserRequest field to given value.

### HasTrackingTypeUserRequest

`func (o *DomainData) HasTrackingTypeUserRequest() bool`

HasTrackingTypeUserRequest returns a boolean if a field has been set.

### GetVERP

`func (o *DomainData) GetVERP() bool`

GetVERP returns the VERP field if non-nil, zero value otherwise.

### GetVERPOk

`func (o *DomainData) GetVERPOk() (*bool, bool)`

GetVERPOk returns a tuple with the VERP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVERP

`func (o *DomainData) SetVERP(v bool)`

SetVERP sets VERP field to given value.

### HasVERP

`func (o *DomainData) HasVERP() bool`

HasVERP returns a boolean if a field has been set.

### GetCustomBouncesDomain

`func (o *DomainData) GetCustomBouncesDomain() string`

GetCustomBouncesDomain returns the CustomBouncesDomain field if non-nil, zero value otherwise.

### GetCustomBouncesDomainOk

`func (o *DomainData) GetCustomBouncesDomainOk() (*string, bool)`

GetCustomBouncesDomainOk returns a tuple with the CustomBouncesDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBouncesDomain

`func (o *DomainData) SetCustomBouncesDomain(v string)`

SetCustomBouncesDomain sets CustomBouncesDomain field to given value.

### HasCustomBouncesDomain

`func (o *DomainData) HasCustomBouncesDomain() bool`

HasCustomBouncesDomain returns a boolean if a field has been set.

### GetIsCustomBouncesDomainDefault

`func (o *DomainData) GetIsCustomBouncesDomainDefault() bool`

GetIsCustomBouncesDomainDefault returns the IsCustomBouncesDomainDefault field if non-nil, zero value otherwise.

### GetIsCustomBouncesDomainDefaultOk

`func (o *DomainData) GetIsCustomBouncesDomainDefaultOk() (*bool, bool)`

GetIsCustomBouncesDomainDefaultOk returns a tuple with the IsCustomBouncesDomainDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomBouncesDomainDefault

`func (o *DomainData) SetIsCustomBouncesDomainDefault(v bool)`

SetIsCustomBouncesDomainDefault sets IsCustomBouncesDomainDefault field to given value.

### HasIsCustomBouncesDomainDefault

`func (o *DomainData) HasIsCustomBouncesDomainDefault() bool`

HasIsCustomBouncesDomainDefault returns a boolean if a field has been set.

### GetIsMarkedForDeletion

`func (o *DomainData) GetIsMarkedForDeletion() bool`

GetIsMarkedForDeletion returns the IsMarkedForDeletion field if non-nil, zero value otherwise.

### GetIsMarkedForDeletionOk

`func (o *DomainData) GetIsMarkedForDeletionOk() (*bool, bool)`

GetIsMarkedForDeletionOk returns a tuple with the IsMarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedForDeletion

`func (o *DomainData) SetIsMarkedForDeletion(v bool)`

SetIsMarkedForDeletion sets IsMarkedForDeletion field to given value.

### HasIsMarkedForDeletion

`func (o *DomainData) HasIsMarkedForDeletion() bool`

HasIsMarkedForDeletion returns a boolean if a field has been set.

### GetOwnership

`func (o *DomainData) GetOwnership() DomainOwner`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *DomainData) GetOwnershipOk() (*DomainOwner, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *DomainData) SetOwnership(v DomainOwner)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *DomainData) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


