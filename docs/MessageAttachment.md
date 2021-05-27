# MessageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryContent** | Pointer to **string** | File&#39;s content as byte array (or a Base64 string) | [optional] 
**Name** | Pointer to **string** | Display name of the file | [optional] 
**ContentType** | Pointer to **string** | MIME content type | [optional] 

## Methods

### NewMessageAttachment

`func NewMessageAttachment() *MessageAttachment`

NewMessageAttachment instantiates a new MessageAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAttachmentWithDefaults

`func NewMessageAttachmentWithDefaults() *MessageAttachment`

NewMessageAttachmentWithDefaults instantiates a new MessageAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryContent

`func (o *MessageAttachment) GetBinaryContent() string`

GetBinaryContent returns the BinaryContent field if non-nil, zero value otherwise.

### GetBinaryContentOk

`func (o *MessageAttachment) GetBinaryContentOk() (*string, bool)`

GetBinaryContentOk returns a tuple with the BinaryContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryContent

`func (o *MessageAttachment) SetBinaryContent(v string)`

SetBinaryContent sets BinaryContent field to given value.

### HasBinaryContent

`func (o *MessageAttachment) HasBinaryContent() bool`

HasBinaryContent returns a boolean if a field has been set.

### GetName

`func (o *MessageAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MessageAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MessageAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MessageAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *MessageAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MessageAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MessageAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MessageAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


