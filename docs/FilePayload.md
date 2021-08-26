# FilePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryContent** | **string** | Content of the file sent as binary data | 
**Name** | Pointer to **string** | Filename | [optional] 
**ContentType** | Pointer to **string** | Type of file&#39;s content (e.g. image/jpeg) | [optional] 

## Methods

### NewFilePayload

`func NewFilePayload(binaryContent string, ) *FilePayload`

NewFilePayload instantiates a new FilePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePayloadWithDefaults

`func NewFilePayloadWithDefaults() *FilePayload`

NewFilePayloadWithDefaults instantiates a new FilePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryContent

`func (o *FilePayload) GetBinaryContent() string`

GetBinaryContent returns the BinaryContent field if non-nil, zero value otherwise.

### GetBinaryContentOk

`func (o *FilePayload) GetBinaryContentOk() (*string, bool)`

GetBinaryContentOk returns a tuple with the BinaryContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryContent

`func (o *FilePayload) SetBinaryContent(v string)`

SetBinaryContent sets BinaryContent field to given value.


### GetName

`func (o *FilePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FilePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FilePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FilePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContentType

`func (o *FilePayload) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FilePayload) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FilePayload) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FilePayload) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


