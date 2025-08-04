# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to [**[]CampaignTemplate**](CampaignTemplate.md) | Campaign&#39;s email content. Provide multiple items to send an A/X Split Campaign | [optional] 
**Name** | **string** | Campaign name | 
**Status** | Pointer to [**CampaignStatus**](CampaignStatus.md) |  | [optional] [default to CAMPAIGNSTATUS_DELETED]
**Recipients** | [**CampaignRecipient**](CampaignRecipient.md) |  | 
**Options** | Pointer to [**CampaignOptions**](CampaignOptions.md) |  | [optional] 

## Methods

### NewCampaign

`func NewCampaign(name string, recipients CampaignRecipient, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *Campaign) GetContent() []CampaignTemplate`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Campaign) GetContentOk() (*[]CampaignTemplate, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Campaign) SetContent(v []CampaignTemplate)`

SetContent sets Content field to given value.

### HasContent

`func (o *Campaign) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Campaign) GetStatus() CampaignStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*CampaignStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v CampaignStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Campaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRecipients

`func (o *Campaign) GetRecipients() CampaignRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *Campaign) GetRecipientsOk() (*CampaignRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *Campaign) SetRecipients(v CampaignRecipient)`

SetRecipients sets Recipients field to given value.


### GetOptions

`func (o *Campaign) GetOptions() CampaignOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Campaign) GetOptionsOk() (*CampaignOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Campaign) SetOptions(v CampaignOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Campaign) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


