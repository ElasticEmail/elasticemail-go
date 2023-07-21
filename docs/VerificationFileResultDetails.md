# VerificationFileResultDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationResult** | Pointer to [**[]EmailValidationResult**](EmailValidationResult.md) | Verification result&#39;s details | [optional] 
**VerificationID** | Pointer to **string** | Identifier of this verification result | [optional] 
**Filename** | Pointer to **string** | Origin file name | [optional] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] [default to VERIFICATIONSTATUS_PROCESSING]
**FileUploadResult** | Pointer to [**FileUploadResult**](FileUploadResult.md) |  | [optional] 
**DateAdded** | Pointer to **time.Time** | Date of creation in YYYY-MM-DDThh:ii:ss format | [optional] 
**Source** | Pointer to **string** | Origin file extension | [optional] 

## Methods

### NewVerificationFileResultDetails

`func NewVerificationFileResultDetails() *VerificationFileResultDetails`

NewVerificationFileResultDetails instantiates a new VerificationFileResultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationFileResultDetailsWithDefaults

`func NewVerificationFileResultDetailsWithDefaults() *VerificationFileResultDetails`

NewVerificationFileResultDetailsWithDefaults instantiates a new VerificationFileResultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationResult

`func (o *VerificationFileResultDetails) GetVerificationResult() []EmailValidationResult`

GetVerificationResult returns the VerificationResult field if non-nil, zero value otherwise.

### GetVerificationResultOk

`func (o *VerificationFileResultDetails) GetVerificationResultOk() (*[]EmailValidationResult, bool)`

GetVerificationResultOk returns a tuple with the VerificationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationResult

`func (o *VerificationFileResultDetails) SetVerificationResult(v []EmailValidationResult)`

SetVerificationResult sets VerificationResult field to given value.

### HasVerificationResult

`func (o *VerificationFileResultDetails) HasVerificationResult() bool`

HasVerificationResult returns a boolean if a field has been set.

### GetVerificationID

`func (o *VerificationFileResultDetails) GetVerificationID() string`

GetVerificationID returns the VerificationID field if non-nil, zero value otherwise.

### GetVerificationIDOk

`func (o *VerificationFileResultDetails) GetVerificationIDOk() (*string, bool)`

GetVerificationIDOk returns a tuple with the VerificationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationID

`func (o *VerificationFileResultDetails) SetVerificationID(v string)`

SetVerificationID sets VerificationID field to given value.

### HasVerificationID

`func (o *VerificationFileResultDetails) HasVerificationID() bool`

HasVerificationID returns a boolean if a field has been set.

### GetFilename

`func (o *VerificationFileResultDetails) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VerificationFileResultDetails) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VerificationFileResultDetails) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *VerificationFileResultDetails) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *VerificationFileResultDetails) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *VerificationFileResultDetails) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *VerificationFileResultDetails) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *VerificationFileResultDetails) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetFileUploadResult

`func (o *VerificationFileResultDetails) GetFileUploadResult() FileUploadResult`

GetFileUploadResult returns the FileUploadResult field if non-nil, zero value otherwise.

### GetFileUploadResultOk

`func (o *VerificationFileResultDetails) GetFileUploadResultOk() (*FileUploadResult, bool)`

GetFileUploadResultOk returns a tuple with the FileUploadResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadResult

`func (o *VerificationFileResultDetails) SetFileUploadResult(v FileUploadResult)`

SetFileUploadResult sets FileUploadResult field to given value.

### HasFileUploadResult

`func (o *VerificationFileResultDetails) HasFileUploadResult() bool`

HasFileUploadResult returns a boolean if a field has been set.

### GetDateAdded

`func (o *VerificationFileResultDetails) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *VerificationFileResultDetails) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *VerificationFileResultDetails) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *VerificationFileResultDetails) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSource

`func (o *VerificationFileResultDetails) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VerificationFileResultDetails) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VerificationFileResultDetails) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VerificationFileResultDetails) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


