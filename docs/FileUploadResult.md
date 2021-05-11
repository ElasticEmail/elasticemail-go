# FileUploadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailsCount** | Pointer to **int32** | How many unique emails were detected the file | [optional] 
**DuplicatedEmailsCount** | Pointer to **int32** | How many email duplicates were detected | [optional] 

## Methods

### NewFileUploadResult

`func NewFileUploadResult() *FileUploadResult`

NewFileUploadResult instantiates a new FileUploadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileUploadResultWithDefaults

`func NewFileUploadResultWithDefaults() *FileUploadResult`

NewFileUploadResultWithDefaults instantiates a new FileUploadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailsCount

`func (o *FileUploadResult) GetEmailsCount() int32`

GetEmailsCount returns the EmailsCount field if non-nil, zero value otherwise.

### GetEmailsCountOk

`func (o *FileUploadResult) GetEmailsCountOk() (*int32, bool)`

GetEmailsCountOk returns a tuple with the EmailsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsCount

`func (o *FileUploadResult) SetEmailsCount(v int32)`

SetEmailsCount sets EmailsCount field to given value.

### HasEmailsCount

`func (o *FileUploadResult) HasEmailsCount() bool`

HasEmailsCount returns a boolean if a field has been set.

### GetDuplicatedEmailsCount

`func (o *FileUploadResult) GetDuplicatedEmailsCount() int32`

GetDuplicatedEmailsCount returns the DuplicatedEmailsCount field if non-nil, zero value otherwise.

### GetDuplicatedEmailsCountOk

`func (o *FileUploadResult) GetDuplicatedEmailsCountOk() (*int32, bool)`

GetDuplicatedEmailsCountOk returns a tuple with the DuplicatedEmailsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicatedEmailsCount

`func (o *FileUploadResult) SetDuplicatedEmailsCount(v int32)`

SetDuplicatedEmailsCount sets DuplicatedEmailsCount field to given value.

### HasDuplicatedEmailsCount

`func (o *FileUploadResult) HasDuplicatedEmailsCount() bool`

HasDuplicatedEmailsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


