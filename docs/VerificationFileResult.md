# VerificationFileResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationID** | Pointer to **string** | Identifier of this verification result | [optional] 
**Filename** | Pointer to **string** | Origin file name | [optional] 
**VerificationStatus** | Pointer to [**VerificationStatus**](VerificationStatus.md) |  | [optional] [default to VERIFICATIONSTATUS_PROCESSING]
**FileUploadResult** | Pointer to [**FileUploadResult**](FileUploadResult.md) |  | [optional] 
**DateAdded** | Pointer to **time.Time** | Date of creation in YYYY-MM-DDThh:ii:ss format | [optional] 
**Source** | Pointer to **string** | Origin file extension | [optional] 

## Methods

### NewVerificationFileResult

`func NewVerificationFileResult() *VerificationFileResult`

NewVerificationFileResult instantiates a new VerificationFileResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationFileResultWithDefaults

`func NewVerificationFileResultWithDefaults() *VerificationFileResult`

NewVerificationFileResultWithDefaults instantiates a new VerificationFileResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationID

`func (o *VerificationFileResult) GetVerificationID() string`

GetVerificationID returns the VerificationID field if non-nil, zero value otherwise.

### GetVerificationIDOk

`func (o *VerificationFileResult) GetVerificationIDOk() (*string, bool)`

GetVerificationIDOk returns a tuple with the VerificationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationID

`func (o *VerificationFileResult) SetVerificationID(v string)`

SetVerificationID sets VerificationID field to given value.

### HasVerificationID

`func (o *VerificationFileResult) HasVerificationID() bool`

HasVerificationID returns a boolean if a field has been set.

### GetFilename

`func (o *VerificationFileResult) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *VerificationFileResult) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *VerificationFileResult) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *VerificationFileResult) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *VerificationFileResult) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *VerificationFileResult) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *VerificationFileResult) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *VerificationFileResult) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetFileUploadResult

`func (o *VerificationFileResult) GetFileUploadResult() FileUploadResult`

GetFileUploadResult returns the FileUploadResult field if non-nil, zero value otherwise.

### GetFileUploadResultOk

`func (o *VerificationFileResult) GetFileUploadResultOk() (*FileUploadResult, bool)`

GetFileUploadResultOk returns a tuple with the FileUploadResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadResult

`func (o *VerificationFileResult) SetFileUploadResult(v FileUploadResult)`

SetFileUploadResult sets FileUploadResult field to given value.

### HasFileUploadResult

`func (o *VerificationFileResult) HasFileUploadResult() bool`

HasFileUploadResult returns a boolean if a field has been set.

### GetDateAdded

`func (o *VerificationFileResult) GetDateAdded() time.Time`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *VerificationFileResult) GetDateAddedOk() (*time.Time, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *VerificationFileResult) SetDateAdded(v time.Time)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *VerificationFileResult) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetSource

`func (o *VerificationFileResult) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VerificationFileResult) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VerificationFileResult) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *VerificationFileResult) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


