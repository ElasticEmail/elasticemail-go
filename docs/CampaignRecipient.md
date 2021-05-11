# CampaignRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListNames** | Pointer to **[]string** | Names of lists from your Account to read recipients from | [optional] 
**SegmentNames** | Pointer to **[]string** | Names of segments from your Account to read recipients from | [optional] 

## Methods

### NewCampaignRecipient

`func NewCampaignRecipient() *CampaignRecipient`

NewCampaignRecipient instantiates a new CampaignRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignRecipientWithDefaults

`func NewCampaignRecipientWithDefaults() *CampaignRecipient`

NewCampaignRecipientWithDefaults instantiates a new CampaignRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListNames

`func (o *CampaignRecipient) GetListNames() []string`

GetListNames returns the ListNames field if non-nil, zero value otherwise.

### GetListNamesOk

`func (o *CampaignRecipient) GetListNamesOk() (*[]string, bool)`

GetListNamesOk returns a tuple with the ListNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListNames

`func (o *CampaignRecipient) SetListNames(v []string)`

SetListNames sets ListNames field to given value.

### HasListNames

`func (o *CampaignRecipient) HasListNames() bool`

HasListNames returns a boolean if a field has been set.

### GetSegmentNames

`func (o *CampaignRecipient) GetSegmentNames() []string`

GetSegmentNames returns the SegmentNames field if non-nil, zero value otherwise.

### GetSegmentNamesOk

`func (o *CampaignRecipient) GetSegmentNamesOk() (*[]string, bool)`

GetSegmentNamesOk returns a tuple with the SegmentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentNames

`func (o *CampaignRecipient) SetSegmentNames(v []string)`

SetSegmentNames sets SegmentNames field to given value.

### HasSegmentNames

`func (o *CampaignRecipient) HasSegmentNames() bool`

HasSegmentNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


