# MergeEmailPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeFile** | [**MessageAttachment**](MessageAttachment.md) |  | 
**Content** | Pointer to [**EmailContent**](EmailContent.md) |  | [optional] 
**Options** | Pointer to [**Options**](Options.md) |  | [optional] 

## Methods

### NewMergeEmailPayload

`func NewMergeEmailPayload(mergeFile MessageAttachment, ) *MergeEmailPayload`

NewMergeEmailPayload instantiates a new MergeEmailPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeEmailPayloadWithDefaults

`func NewMergeEmailPayloadWithDefaults() *MergeEmailPayload`

NewMergeEmailPayloadWithDefaults instantiates a new MergeEmailPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeFile

`func (o *MergeEmailPayload) GetMergeFile() MessageAttachment`

GetMergeFile returns the MergeFile field if non-nil, zero value otherwise.

### GetMergeFileOk

`func (o *MergeEmailPayload) GetMergeFileOk() (*MessageAttachment, bool)`

GetMergeFileOk returns a tuple with the MergeFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeFile

`func (o *MergeEmailPayload) SetMergeFile(v MessageAttachment)`

SetMergeFile sets MergeFile field to given value.


### GetContent

`func (o *MergeEmailPayload) GetContent() EmailContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MergeEmailPayload) GetContentOk() (*EmailContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MergeEmailPayload) SetContent(v EmailContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *MergeEmailPayload) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetOptions

`func (o *MergeEmailPayload) GetOptions() Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *MergeEmailPayload) GetOptionsOk() (*Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *MergeEmailPayload) SetOptions(v Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *MergeEmailPayload) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


