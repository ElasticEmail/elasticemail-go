# EmailMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | [**[]EmailRecipient**](EmailRecipient.md) | List of recipients | 
**Content** | Pointer to [**EmailContent**](EmailContent.md) |  | [optional] 
**Options** | Pointer to [**Options**](Options.md) |  | [optional] 

## Methods

### NewEmailMessageData

`func NewEmailMessageData(recipients []EmailRecipient, ) *EmailMessageData`

NewEmailMessageData instantiates a new EmailMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailMessageDataWithDefaults

`func NewEmailMessageDataWithDefaults() *EmailMessageData`

NewEmailMessageDataWithDefaults instantiates a new EmailMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *EmailMessageData) GetRecipients() []EmailRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailMessageData) GetRecipientsOk() (*[]EmailRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailMessageData) SetRecipients(v []EmailRecipient)`

SetRecipients sets Recipients field to given value.


### GetContent

`func (o *EmailMessageData) GetContent() EmailContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmailMessageData) GetContentOk() (*EmailContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmailMessageData) SetContent(v EmailContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *EmailMessageData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetOptions

`func (o *EmailMessageData) GetOptions() Options`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EmailMessageData) GetOptionsOk() (*Options, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EmailMessageData) SetOptions(v Options)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EmailMessageData) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


