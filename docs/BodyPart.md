# BodyPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | [**BodyContentType**](BodyContentType.md) |  | [default to BODYCONTENTTYPE_HTML]
**Content** | Pointer to **string** | Actual content of the body part | [optional] 
**Charset** | Pointer to **string** | Text value of charset encoding for example: iso-8859-1, windows-1251, utf-8, us-ascii, windows-1250 and more... | [optional] 

## Methods

### NewBodyPart

`func NewBodyPart(contentType BodyContentType, ) *BodyPart`

NewBodyPart instantiates a new BodyPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBodyPartWithDefaults

`func NewBodyPartWithDefaults() *BodyPart`

NewBodyPartWithDefaults instantiates a new BodyPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *BodyPart) GetContentType() BodyContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BodyPart) GetContentTypeOk() (*BodyContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BodyPart) SetContentType(v BodyContentType)`

SetContentType sets ContentType field to given value.


### GetContent

`func (o *BodyPart) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BodyPart) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BodyPart) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BodyPart) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCharset

`func (o *BodyPart) GetCharset() string`

GetCharset returns the Charset field if non-nil, zero value otherwise.

### GetCharsetOk

`func (o *BodyPart) GetCharsetOk() (*string, bool)`

GetCharsetOk returns a tuple with the Charset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharset

`func (o *BodyPart) SetCharset(v string)`

SetCharset sets Charset field to given value.

### HasCharset

`func (o *BodyPart) HasCharset() bool`

HasCharset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


