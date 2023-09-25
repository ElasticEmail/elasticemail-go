# MessageAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryContent** | **string** | File&#39;s content as byte array (or a Base64 string) | 
**Name** | **string** | Display name of the file | 
**ContentType** | Pointer to **string** | MIME content type | [optional] 
**Size** | Pointer to **int32** | Size of your attachment (in bytes). | [optional] 

## Methods

### NewMessageAttachment

`func NewMessageAttachment(binaryContent string, name string, ) *MessageAttachment`

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

### GetSize

`func (o *MessageAttachment) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MessageAttachment) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MessageAttachment) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MessageAttachment) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


