# ChannelLogStatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelName** | Pointer to **string** | Channel name | [optional] 
**Recipients** | Pointer to **int64** | Number of recipients | [optional] 
**EmailTotal** | Pointer to **int64** | Number of emails | [optional] 
**SmsTotal** | Pointer to **int64** | Number of SMS | [optional] 
**Delivered** | Pointer to **int64** | Number of delivered messages | [optional] 
**Bounced** | Pointer to **int64** | Number of bounced messages | [optional] 
**InProgress** | Pointer to **int64** | Number of messages in progress | [optional] 
**Opened** | Pointer to **int64** | Number of opened messages | [optional] 
**Clicked** | Pointer to **int64** | Number of clicked messages | [optional] 
**Unsubscribed** | Pointer to **int64** | Number of unsubscribed messages | [optional] 
**Complaints** | Pointer to **int64** | Number of complaint messages | [optional] 
**Inbound** | Pointer to **int64** | Number of inbound messages | [optional] 
**ManualCancel** | Pointer to **int64** | Number of manually cancelled messages | [optional] 
**NotDelivered** | Pointer to **int64** | Number of messages flagged with &#39;Not Delivered&#39; | [optional] 

## Methods

### NewChannelLogStatusSummary

`func NewChannelLogStatusSummary() *ChannelLogStatusSummary`

NewChannelLogStatusSummary instantiates a new ChannelLogStatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelLogStatusSummaryWithDefaults

`func NewChannelLogStatusSummaryWithDefaults() *ChannelLogStatusSummary`

NewChannelLogStatusSummaryWithDefaults instantiates a new ChannelLogStatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelName

`func (o *ChannelLogStatusSummary) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ChannelLogStatusSummary) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ChannelLogStatusSummary) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ChannelLogStatusSummary) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetRecipients

`func (o *ChannelLogStatusSummary) GetRecipients() int64`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ChannelLogStatusSummary) GetRecipientsOk() (*int64, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ChannelLogStatusSummary) SetRecipients(v int64)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *ChannelLogStatusSummary) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetEmailTotal

`func (o *ChannelLogStatusSummary) GetEmailTotal() int64`

GetEmailTotal returns the EmailTotal field if non-nil, zero value otherwise.

### GetEmailTotalOk

`func (o *ChannelLogStatusSummary) GetEmailTotalOk() (*int64, bool)`

GetEmailTotalOk returns a tuple with the EmailTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTotal

`func (o *ChannelLogStatusSummary) SetEmailTotal(v int64)`

SetEmailTotal sets EmailTotal field to given value.

### HasEmailTotal

`func (o *ChannelLogStatusSummary) HasEmailTotal() bool`

HasEmailTotal returns a boolean if a field has been set.

### GetSmsTotal

`func (o *ChannelLogStatusSummary) GetSmsTotal() int64`

GetSmsTotal returns the SmsTotal field if non-nil, zero value otherwise.

### GetSmsTotalOk

`func (o *ChannelLogStatusSummary) GetSmsTotalOk() (*int64, bool)`

GetSmsTotalOk returns a tuple with the SmsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsTotal

`func (o *ChannelLogStatusSummary) SetSmsTotal(v int64)`

SetSmsTotal sets SmsTotal field to given value.

### HasSmsTotal

`func (o *ChannelLogStatusSummary) HasSmsTotal() bool`

HasSmsTotal returns a boolean if a field has been set.

### GetDelivered

`func (o *ChannelLogStatusSummary) GetDelivered() int64`

GetDelivered returns the Delivered field if non-nil, zero value otherwise.

### GetDeliveredOk

`func (o *ChannelLogStatusSummary) GetDeliveredOk() (*int64, bool)`

GetDeliveredOk returns a tuple with the Delivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivered

`func (o *ChannelLogStatusSummary) SetDelivered(v int64)`

SetDelivered sets Delivered field to given value.

### HasDelivered

`func (o *ChannelLogStatusSummary) HasDelivered() bool`

HasDelivered returns a boolean if a field has been set.

### GetBounced

`func (o *ChannelLogStatusSummary) GetBounced() int64`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *ChannelLogStatusSummary) GetBouncedOk() (*int64, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *ChannelLogStatusSummary) SetBounced(v int64)`

SetBounced sets Bounced field to given value.

### HasBounced

`func (o *ChannelLogStatusSummary) HasBounced() bool`

HasBounced returns a boolean if a field has been set.

### GetInProgress

`func (o *ChannelLogStatusSummary) GetInProgress() int64`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *ChannelLogStatusSummary) GetInProgressOk() (*int64, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *ChannelLogStatusSummary) SetInProgress(v int64)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *ChannelLogStatusSummary) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetOpened

`func (o *ChannelLogStatusSummary) GetOpened() int64`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *ChannelLogStatusSummary) GetOpenedOk() (*int64, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *ChannelLogStatusSummary) SetOpened(v int64)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *ChannelLogStatusSummary) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetClicked

`func (o *ChannelLogStatusSummary) GetClicked() int64`

GetClicked returns the Clicked field if non-nil, zero value otherwise.

### GetClickedOk

`func (o *ChannelLogStatusSummary) GetClickedOk() (*int64, bool)`

GetClickedOk returns a tuple with the Clicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicked

`func (o *ChannelLogStatusSummary) SetClicked(v int64)`

SetClicked sets Clicked field to given value.

### HasClicked

`func (o *ChannelLogStatusSummary) HasClicked() bool`

HasClicked returns a boolean if a field has been set.

### GetUnsubscribed

`func (o *ChannelLogStatusSummary) GetUnsubscribed() int64`

GetUnsubscribed returns the Unsubscribed field if non-nil, zero value otherwise.

### GetUnsubscribedOk

`func (o *ChannelLogStatusSummary) GetUnsubscribedOk() (*int64, bool)`

GetUnsubscribedOk returns a tuple with the Unsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribed

`func (o *ChannelLogStatusSummary) SetUnsubscribed(v int64)`

SetUnsubscribed sets Unsubscribed field to given value.

### HasUnsubscribed

`func (o *ChannelLogStatusSummary) HasUnsubscribed() bool`

HasUnsubscribed returns a boolean if a field has been set.

### GetComplaints

`func (o *ChannelLogStatusSummary) GetComplaints() int64`

GetComplaints returns the Complaints field if non-nil, zero value otherwise.

### GetComplaintsOk

`func (o *ChannelLogStatusSummary) GetComplaintsOk() (*int64, bool)`

GetComplaintsOk returns a tuple with the Complaints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplaints

`func (o *ChannelLogStatusSummary) SetComplaints(v int64)`

SetComplaints sets Complaints field to given value.

### HasComplaints

`func (o *ChannelLogStatusSummary) HasComplaints() bool`

HasComplaints returns a boolean if a field has been set.

### GetInbound

`func (o *ChannelLogStatusSummary) GetInbound() int64`

GetInbound returns the Inbound field if non-nil, zero value otherwise.

### GetInboundOk

`func (o *ChannelLogStatusSummary) GetInboundOk() (*int64, bool)`

GetInboundOk returns a tuple with the Inbound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbound

`func (o *ChannelLogStatusSummary) SetInbound(v int64)`

SetInbound sets Inbound field to given value.

### HasInbound

`func (o *ChannelLogStatusSummary) HasInbound() bool`

HasInbound returns a boolean if a field has been set.

### GetManualCancel

`func (o *ChannelLogStatusSummary) GetManualCancel() int64`

GetManualCancel returns the ManualCancel field if non-nil, zero value otherwise.

### GetManualCancelOk

`func (o *ChannelLogStatusSummary) GetManualCancelOk() (*int64, bool)`

GetManualCancelOk returns a tuple with the ManualCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualCancel

`func (o *ChannelLogStatusSummary) SetManualCancel(v int64)`

SetManualCancel sets ManualCancel field to given value.

### HasManualCancel

`func (o *ChannelLogStatusSummary) HasManualCancel() bool`

HasManualCancel returns a boolean if a field has been set.

### GetNotDelivered

`func (o *ChannelLogStatusSummary) GetNotDelivered() int64`

GetNotDelivered returns the NotDelivered field if non-nil, zero value otherwise.

### GetNotDeliveredOk

`func (o *ChannelLogStatusSummary) GetNotDeliveredOk() (*int64, bool)`

GetNotDeliveredOk returns a tuple with the NotDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotDelivered

`func (o *ChannelLogStatusSummary) SetNotDelivered(v int64)`

SetNotDelivered sets NotDelivered field to given value.

### HasNotDelivered

`func (o *ChannelLogStatusSummary) HasNotDelivered() bool`

HasNotDelivered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


