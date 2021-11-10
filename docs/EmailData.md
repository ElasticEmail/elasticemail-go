# EmailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preview** | Pointer to [**EmailView**](EmailView.md) |  | [optional] 
**Attachments** | Pointer to [**[]FileInfo**](FileInfo.md) | Attachments sent with the email | [optional] 
**Status** | Pointer to [**EmailStatus**](EmailStatus.md) |  | [optional] 

## Methods

### NewEmailData

`func NewEmailData() *EmailData`

NewEmailData instantiates a new EmailData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDataWithDefaults

`func NewEmailDataWithDefaults() *EmailData`

NewEmailDataWithDefaults instantiates a new EmailData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreview

`func (o *EmailData) GetPreview() EmailView`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *EmailData) GetPreviewOk() (*EmailView, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *EmailData) SetPreview(v EmailView)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *EmailData) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetAttachments

`func (o *EmailData) GetAttachments() []FileInfo`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailData) GetAttachmentsOk() (*[]FileInfo, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailData) SetAttachments(v []FileInfo)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailData) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetStatus

`func (o *EmailData) GetStatus() EmailStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EmailData) GetStatusOk() (*EmailStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EmailData) SetStatus(v EmailStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EmailData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


