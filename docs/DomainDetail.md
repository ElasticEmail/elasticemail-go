# DomainDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Name of selected domain. | [optional] 
**DefaultDomain** | Pointer to **bool** | True, if domain is used as default. Otherwise, false, | [optional] 
**Spf** | Pointer to **bool** | True, if SPF record is verified | [optional] 
**Dkim** | Pointer to **bool** | True, if DKIM record is verified | [optional] 
**MX** | Pointer to **bool** | True, if MX record is verified | [optional] 
**DMARC** | Pointer to **bool** |  | [optional] 
**IsRewriteDomainValid** | Pointer to **bool** | True, if tracking CNAME record is verified | [optional] 
**Verify** | Pointer to **bool** | True, if DKIM, SPF, or tracking are still to be verified | [optional] 
**Type** | Pointer to [**TrackingType**](TrackingType.md) |  | [optional] [default to TRACKINGTYPE_NONE]
**TrackingStatus** | Pointer to [**TrackingValidationStatus**](TrackingValidationStatus.md) |  | [optional] [default to TRACKINGVALIDATIONSTATUS_VALIDATED]
**CertificateStatus** | Pointer to [**CertificateValidationStatus**](CertificateValidationStatus.md) |  | [optional] [default to CERTIFICATEVALIDATIONSTATUS_ERROR_OCCURED]
**CertificateValidationError** | Pointer to **string** |  | [optional] 
**TrackingTypeUserRequest** | Pointer to [**TrackingType**](TrackingType.md) |  | [optional] [default to TRACKINGTYPE_NONE]
**VERP** | Pointer to **bool** |  | [optional] 
**CustomBouncesDomain** | Pointer to **string** |  | [optional] 
**IsCustomBouncesDomainDefault** | Pointer to **bool** |  | [optional] 
**IsMarkedForDeletion** | Pointer to **bool** |  | [optional] 
**Ownership** | Pointer to [**DomainOwner**](DomainOwner.md) |  | [optional] [default to DOMAINOWNER_CURRENT]

## Methods

### NewDomainDetail

`func NewDomainDetail() *DomainDetail`

NewDomainDetail instantiates a new DomainDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDetailWithDefaults

`func NewDomainDetailWithDefaults() *DomainDetail`

NewDomainDetailWithDefaults instantiates a new DomainDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainDetail) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDefaultDomain

`func (o *DomainDetail) GetDefaultDomain() bool`

GetDefaultDomain returns the DefaultDomain field if non-nil, zero value otherwise.

### GetDefaultDomainOk

`func (o *DomainDetail) GetDefaultDomainOk() (*bool, bool)`

GetDefaultDomainOk returns a tuple with the DefaultDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomain

`func (o *DomainDetail) SetDefaultDomain(v bool)`

SetDefaultDomain sets DefaultDomain field to given value.

### HasDefaultDomain

`func (o *DomainDetail) HasDefaultDomain() bool`

HasDefaultDomain returns a boolean if a field has been set.

### GetSpf

`func (o *DomainDetail) GetSpf() bool`

GetSpf returns the Spf field if non-nil, zero value otherwise.

### GetSpfOk

`func (o *DomainDetail) GetSpfOk() (*bool, bool)`

GetSpfOk returns a tuple with the Spf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpf

`func (o *DomainDetail) SetSpf(v bool)`

SetSpf sets Spf field to given value.

### HasSpf

`func (o *DomainDetail) HasSpf() bool`

HasSpf returns a boolean if a field has been set.

### GetDkim

`func (o *DomainDetail) GetDkim() bool`

GetDkim returns the Dkim field if non-nil, zero value otherwise.

### GetDkimOk

`func (o *DomainDetail) GetDkimOk() (*bool, bool)`

GetDkimOk returns a tuple with the Dkim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim

`func (o *DomainDetail) SetDkim(v bool)`

SetDkim sets Dkim field to given value.

### HasDkim

`func (o *DomainDetail) HasDkim() bool`

HasDkim returns a boolean if a field has been set.

### GetMX

`func (o *DomainDetail) GetMX() bool`

GetMX returns the MX field if non-nil, zero value otherwise.

### GetMXOk

`func (o *DomainDetail) GetMXOk() (*bool, bool)`

GetMXOk returns a tuple with the MX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMX

`func (o *DomainDetail) SetMX(v bool)`

SetMX sets MX field to given value.

### HasMX

`func (o *DomainDetail) HasMX() bool`

HasMX returns a boolean if a field has been set.

### GetDMARC

`func (o *DomainDetail) GetDMARC() bool`

GetDMARC returns the DMARC field if non-nil, zero value otherwise.

### GetDMARCOk

`func (o *DomainDetail) GetDMARCOk() (*bool, bool)`

GetDMARCOk returns a tuple with the DMARC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMARC

`func (o *DomainDetail) SetDMARC(v bool)`

SetDMARC sets DMARC field to given value.

### HasDMARC

`func (o *DomainDetail) HasDMARC() bool`

HasDMARC returns a boolean if a field has been set.

### GetIsRewriteDomainValid

`func (o *DomainDetail) GetIsRewriteDomainValid() bool`

GetIsRewriteDomainValid returns the IsRewriteDomainValid field if non-nil, zero value otherwise.

### GetIsRewriteDomainValidOk

`func (o *DomainDetail) GetIsRewriteDomainValidOk() (*bool, bool)`

GetIsRewriteDomainValidOk returns a tuple with the IsRewriteDomainValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRewriteDomainValid

`func (o *DomainDetail) SetIsRewriteDomainValid(v bool)`

SetIsRewriteDomainValid sets IsRewriteDomainValid field to given value.

### HasIsRewriteDomainValid

`func (o *DomainDetail) HasIsRewriteDomainValid() bool`

HasIsRewriteDomainValid returns a boolean if a field has been set.

### GetVerify

`func (o *DomainDetail) GetVerify() bool`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *DomainDetail) GetVerifyOk() (*bool, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *DomainDetail) SetVerify(v bool)`

SetVerify sets Verify field to given value.

### HasVerify

`func (o *DomainDetail) HasVerify() bool`

HasVerify returns a boolean if a field has been set.

### GetType

`func (o *DomainDetail) GetType() TrackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainDetail) GetTypeOk() (*TrackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainDetail) SetType(v TrackingType)`

SetType sets Type field to given value.

### HasType

`func (o *DomainDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTrackingStatus

`func (o *DomainDetail) GetTrackingStatus() TrackingValidationStatus`

GetTrackingStatus returns the TrackingStatus field if non-nil, zero value otherwise.

### GetTrackingStatusOk

`func (o *DomainDetail) GetTrackingStatusOk() (*TrackingValidationStatus, bool)`

GetTrackingStatusOk returns a tuple with the TrackingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingStatus

`func (o *DomainDetail) SetTrackingStatus(v TrackingValidationStatus)`

SetTrackingStatus sets TrackingStatus field to given value.

### HasTrackingStatus

`func (o *DomainDetail) HasTrackingStatus() bool`

HasTrackingStatus returns a boolean if a field has been set.

### GetCertificateStatus

`func (o *DomainDetail) GetCertificateStatus() CertificateValidationStatus`

GetCertificateStatus returns the CertificateStatus field if non-nil, zero value otherwise.

### GetCertificateStatusOk

`func (o *DomainDetail) GetCertificateStatusOk() (*CertificateValidationStatus, bool)`

GetCertificateStatusOk returns a tuple with the CertificateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStatus

`func (o *DomainDetail) SetCertificateStatus(v CertificateValidationStatus)`

SetCertificateStatus sets CertificateStatus field to given value.

### HasCertificateStatus

`func (o *DomainDetail) HasCertificateStatus() bool`

HasCertificateStatus returns a boolean if a field has been set.

### GetCertificateValidationError

`func (o *DomainDetail) GetCertificateValidationError() string`

GetCertificateValidationError returns the CertificateValidationError field if non-nil, zero value otherwise.

### GetCertificateValidationErrorOk

`func (o *DomainDetail) GetCertificateValidationErrorOk() (*string, bool)`

GetCertificateValidationErrorOk returns a tuple with the CertificateValidationError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateValidationError

`func (o *DomainDetail) SetCertificateValidationError(v string)`

SetCertificateValidationError sets CertificateValidationError field to given value.

### HasCertificateValidationError

`func (o *DomainDetail) HasCertificateValidationError() bool`

HasCertificateValidationError returns a boolean if a field has been set.

### GetTrackingTypeUserRequest

`func (o *DomainDetail) GetTrackingTypeUserRequest() TrackingType`

GetTrackingTypeUserRequest returns the TrackingTypeUserRequest field if non-nil, zero value otherwise.

### GetTrackingTypeUserRequestOk

`func (o *DomainDetail) GetTrackingTypeUserRequestOk() (*TrackingType, bool)`

GetTrackingTypeUserRequestOk returns a tuple with the TrackingTypeUserRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingTypeUserRequest

`func (o *DomainDetail) SetTrackingTypeUserRequest(v TrackingType)`

SetTrackingTypeUserRequest sets TrackingTypeUserRequest field to given value.

### HasTrackingTypeUserRequest

`func (o *DomainDetail) HasTrackingTypeUserRequest() bool`

HasTrackingTypeUserRequest returns a boolean if a field has been set.

### GetVERP

`func (o *DomainDetail) GetVERP() bool`

GetVERP returns the VERP field if non-nil, zero value otherwise.

### GetVERPOk

`func (o *DomainDetail) GetVERPOk() (*bool, bool)`

GetVERPOk returns a tuple with the VERP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVERP

`func (o *DomainDetail) SetVERP(v bool)`

SetVERP sets VERP field to given value.

### HasVERP

`func (o *DomainDetail) HasVERP() bool`

HasVERP returns a boolean if a field has been set.

### GetCustomBouncesDomain

`func (o *DomainDetail) GetCustomBouncesDomain() string`

GetCustomBouncesDomain returns the CustomBouncesDomain field if non-nil, zero value otherwise.

### GetCustomBouncesDomainOk

`func (o *DomainDetail) GetCustomBouncesDomainOk() (*string, bool)`

GetCustomBouncesDomainOk returns a tuple with the CustomBouncesDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomBouncesDomain

`func (o *DomainDetail) SetCustomBouncesDomain(v string)`

SetCustomBouncesDomain sets CustomBouncesDomain field to given value.

### HasCustomBouncesDomain

`func (o *DomainDetail) HasCustomBouncesDomain() bool`

HasCustomBouncesDomain returns a boolean if a field has been set.

### GetIsCustomBouncesDomainDefault

`func (o *DomainDetail) GetIsCustomBouncesDomainDefault() bool`

GetIsCustomBouncesDomainDefault returns the IsCustomBouncesDomainDefault field if non-nil, zero value otherwise.

### GetIsCustomBouncesDomainDefaultOk

`func (o *DomainDetail) GetIsCustomBouncesDomainDefaultOk() (*bool, bool)`

GetIsCustomBouncesDomainDefaultOk returns a tuple with the IsCustomBouncesDomainDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomBouncesDomainDefault

`func (o *DomainDetail) SetIsCustomBouncesDomainDefault(v bool)`

SetIsCustomBouncesDomainDefault sets IsCustomBouncesDomainDefault field to given value.

### HasIsCustomBouncesDomainDefault

`func (o *DomainDetail) HasIsCustomBouncesDomainDefault() bool`

HasIsCustomBouncesDomainDefault returns a boolean if a field has been set.

### GetIsMarkedForDeletion

`func (o *DomainDetail) GetIsMarkedForDeletion() bool`

GetIsMarkedForDeletion returns the IsMarkedForDeletion field if non-nil, zero value otherwise.

### GetIsMarkedForDeletionOk

`func (o *DomainDetail) GetIsMarkedForDeletionOk() (*bool, bool)`

GetIsMarkedForDeletionOk returns a tuple with the IsMarkedForDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMarkedForDeletion

`func (o *DomainDetail) SetIsMarkedForDeletion(v bool)`

SetIsMarkedForDeletion sets IsMarkedForDeletion field to given value.

### HasIsMarkedForDeletion

`func (o *DomainDetail) HasIsMarkedForDeletion() bool`

HasIsMarkedForDeletion returns a boolean if a field has been set.

### GetOwnership

`func (o *DomainDetail) GetOwnership() DomainOwner`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *DomainDetail) GetOwnershipOk() (*DomainOwner, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *DomainDetail) SetOwnership(v DomainOwner)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *DomainDetail) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


