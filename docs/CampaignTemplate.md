# CampaignTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Poolname** | Pointer to **string** | Name of your custom IP Pool to be used in the sending process | [optional] 
**From** | Pointer to **string** | Your e-mail with an optional name (e.g.: John Doe &lt;email@domain.com&gt;) | [optional] 
**ReplyTo** | Pointer to **string** | To what address should the recipients reply to (e.g. John Doe &lt;email@domain.com&gt;) | [optional] 
**Subject** | Pointer to **string** | Default subject of email. | [optional] 
**TemplateName** | Pointer to **string** | Name of template. | [optional] 
**AttachFiles** | Pointer to **[]string** | Names of previously uploaded files that should be sent as downloadable attachments | [optional] 
**Utm** | Pointer to [**Utm**](Utm.md) |  | [optional] 

## Methods

### NewCampaignTemplate

`func NewCampaignTemplate() *CampaignTemplate`

NewCampaignTemplate instantiates a new CampaignTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignTemplateWithDefaults

`func NewCampaignTemplateWithDefaults() *CampaignTemplate`

NewCampaignTemplateWithDefaults instantiates a new CampaignTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolname

`func (o *CampaignTemplate) GetPoolname() string`

GetPoolname returns the Poolname field if non-nil, zero value otherwise.

### GetPoolnameOk

`func (o *CampaignTemplate) GetPoolnameOk() (*string, bool)`

GetPoolnameOk returns a tuple with the Poolname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolname

`func (o *CampaignTemplate) SetPoolname(v string)`

SetPoolname sets Poolname field to given value.

### HasPoolname

`func (o *CampaignTemplate) HasPoolname() bool`

HasPoolname returns a boolean if a field has been set.

### GetFrom

`func (o *CampaignTemplate) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CampaignTemplate) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CampaignTemplate) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CampaignTemplate) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *CampaignTemplate) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CampaignTemplate) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CampaignTemplate) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CampaignTemplate) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSubject

`func (o *CampaignTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CampaignTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CampaignTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CampaignTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTemplateName

`func (o *CampaignTemplate) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *CampaignTemplate) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *CampaignTemplate) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *CampaignTemplate) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetAttachFiles

`func (o *CampaignTemplate) GetAttachFiles() []string`

GetAttachFiles returns the AttachFiles field if non-nil, zero value otherwise.

### GetAttachFilesOk

`func (o *CampaignTemplate) GetAttachFilesOk() (*[]string, bool)`

GetAttachFilesOk returns a tuple with the AttachFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachFiles

`func (o *CampaignTemplate) SetAttachFiles(v []string)`

SetAttachFiles sets AttachFiles field to given value.

### HasAttachFiles

`func (o *CampaignTemplate) HasAttachFiles() bool`

HasAttachFiles returns a boolean if a field has been set.

### GetUtm

`func (o *CampaignTemplate) GetUtm() Utm`

GetUtm returns the Utm field if non-nil, zero value otherwise.

### GetUtmOk

`func (o *CampaignTemplate) GetUtmOk() (*Utm, bool)`

GetUtmOk returns a tuple with the Utm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtm

`func (o *CampaignTemplate) SetUtm(v Utm)`

SetUtm sets Utm field to given value.

### HasUtm

`func (o *CampaignTemplate) HasUtm() bool`

HasUtm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


