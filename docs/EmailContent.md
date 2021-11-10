# EmailContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to [**[]BodyPart**](BodyPart.md) | List of e-mail body parts, with user-provided MIME types (text/html, text/plain etc) | [optional] 
**Merge** | Pointer to **map[string]string** | A key-value collection of custom merge fields, shared between recipients. Should be used in e-mail body like so: {firstname}, {lastname} etc. | [optional] 
**Attachments** | Pointer to [**[]MessageAttachment**](MessageAttachment.md) | Attachments provided by sending binary data | [optional] 
**Headers** | Pointer to **map[string]string** | A key-value collection of custom e-mail headers. | [optional] 
**Postback** | Pointer to **string** | Postback header. | [optional] 
**EnvelopeFrom** | Pointer to **string** | E-mail with an optional name to be used as the envelope from address (e.g.: John Doe &lt;email@domain.com&gt;) | [optional] 
**From** | Pointer to **string** | Your e-mail with an optional name (e.g.: John Doe &lt;email@domain.com&gt;) | [optional] 
**ReplyTo** | Pointer to **string** | To what address should the recipients reply to (e.g. John Doe &lt;email@domain.com&gt;) | [optional] 
**Subject** | Pointer to **string** | Default subject of email. | [optional] 
**TemplateName** | Pointer to **string** | Name of template. | [optional] 
**AttachFiles** | Pointer to **[]string** | Names of previously uploaded files that should be sent as downloadable attachments | [optional] 
**Utm** | Pointer to [**Utm**](Utm.md) |  | [optional] 

## Methods

### NewEmailContent

`func NewEmailContent() *EmailContent`

NewEmailContent instantiates a new EmailContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailContentWithDefaults

`func NewEmailContentWithDefaults() *EmailContent`

NewEmailContentWithDefaults instantiates a new EmailContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *EmailContent) GetBody() []BodyPart`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EmailContent) GetBodyOk() (*[]BodyPart, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EmailContent) SetBody(v []BodyPart)`

SetBody sets Body field to given value.

### HasBody

`func (o *EmailContent) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMerge

`func (o *EmailContent) GetMerge() map[string]string`

GetMerge returns the Merge field if non-nil, zero value otherwise.

### GetMergeOk

`func (o *EmailContent) GetMergeOk() (*map[string]string, bool)`

GetMergeOk returns a tuple with the Merge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerge

`func (o *EmailContent) SetMerge(v map[string]string)`

SetMerge sets Merge field to given value.

### HasMerge

`func (o *EmailContent) HasMerge() bool`

HasMerge returns a boolean if a field has been set.

### GetAttachments

`func (o *EmailContent) GetAttachments() []MessageAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *EmailContent) GetAttachmentsOk() (*[]MessageAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *EmailContent) SetAttachments(v []MessageAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *EmailContent) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetHeaders

`func (o *EmailContent) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *EmailContent) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *EmailContent) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *EmailContent) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPostback

`func (o *EmailContent) GetPostback() string`

GetPostback returns the Postback field if non-nil, zero value otherwise.

### GetPostbackOk

`func (o *EmailContent) GetPostbackOk() (*string, bool)`

GetPostbackOk returns a tuple with the Postback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostback

`func (o *EmailContent) SetPostback(v string)`

SetPostback sets Postback field to given value.

### HasPostback

`func (o *EmailContent) HasPostback() bool`

HasPostback returns a boolean if a field has been set.

### GetEnvelopeFrom

`func (o *EmailContent) GetEnvelopeFrom() string`

GetEnvelopeFrom returns the EnvelopeFrom field if non-nil, zero value otherwise.

### GetEnvelopeFromOk

`func (o *EmailContent) GetEnvelopeFromOk() (*string, bool)`

GetEnvelopeFromOk returns a tuple with the EnvelopeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvelopeFrom

`func (o *EmailContent) SetEnvelopeFrom(v string)`

SetEnvelopeFrom sets EnvelopeFrom field to given value.

### HasEnvelopeFrom

`func (o *EmailContent) HasEnvelopeFrom() bool`

HasEnvelopeFrom returns a boolean if a field has been set.

### GetFrom

`func (o *EmailContent) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EmailContent) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EmailContent) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EmailContent) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *EmailContent) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *EmailContent) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *EmailContent) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *EmailContent) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *EmailContent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailContent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailContent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailContent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplateName

`func (o *EmailContent) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *EmailContent) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *EmailContent) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *EmailContent) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetAttachFiles

`func (o *EmailContent) GetAttachFiles() []string`

GetAttachFiles returns the AttachFiles field if non-nil, zero value otherwise.

### GetAttachFilesOk

`func (o *EmailContent) GetAttachFilesOk() (*[]string, bool)`

GetAttachFilesOk returns a tuple with the AttachFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachFiles

`func (o *EmailContent) SetAttachFiles(v []string)`

SetAttachFiles sets AttachFiles field to given value.

### HasAttachFiles

`func (o *EmailContent) HasAttachFiles() bool`

HasAttachFiles returns a boolean if a field has been set.

### GetUtm

`func (o *EmailContent) GetUtm() Utm`

GetUtm returns the Utm field if non-nil, zero value otherwise.

### GetUtmOk

`func (o *EmailContent) GetUtmOk() (*Utm, bool)`

GetUtmOk returns a tuple with the Utm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtm

`func (o *EmailContent) SetUtm(v Utm)`

SetUtm sets Utm field to given value.

### HasUtm

`func (o *EmailContent) HasUtm() bool`

HasUtm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


