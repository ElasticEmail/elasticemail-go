# WebhookUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Filename | [optional] 
**URL** | Pointer to **string** | URL of notification. | [optional] 
**NotifyOncePerEmail** | Pointer to **NullableBool** |  | [optional] 
**NotificationForSent** | Pointer to **NullableBool** |  | [optional] 
**NotificationForOpened** | Pointer to **NullableBool** |  | [optional] 
**NotificationForClicked** | Pointer to **NullableBool** |  | [optional] 
**NotificationForUnsubscribed** | Pointer to **NullableBool** |  | [optional] 
**NotificationForAbuseReport** | Pointer to **NullableBool** |  | [optional] 
**NotificationForError** | Pointer to **NullableBool** |  | [optional] 
**IsEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewWebhookUpdatePayload

`func NewWebhookUpdatePayload() *WebhookUpdatePayload`

NewWebhookUpdatePayload instantiates a new WebhookUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdatePayloadWithDefaults

`func NewWebhookUpdatePayloadWithDefaults() *WebhookUpdatePayload`

NewWebhookUpdatePayloadWithDefaults instantiates a new WebhookUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookUpdatePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookUpdatePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookUpdatePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookUpdatePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetURL

`func (o *WebhookUpdatePayload) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *WebhookUpdatePayload) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *WebhookUpdatePayload) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *WebhookUpdatePayload) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetNotifyOncePerEmail

`func (o *WebhookUpdatePayload) GetNotifyOncePerEmail() bool`

GetNotifyOncePerEmail returns the NotifyOncePerEmail field if non-nil, zero value otherwise.

### GetNotifyOncePerEmailOk

`func (o *WebhookUpdatePayload) GetNotifyOncePerEmailOk() (*bool, bool)`

GetNotifyOncePerEmailOk returns a tuple with the NotifyOncePerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOncePerEmail

`func (o *WebhookUpdatePayload) SetNotifyOncePerEmail(v bool)`

SetNotifyOncePerEmail sets NotifyOncePerEmail field to given value.

### HasNotifyOncePerEmail

`func (o *WebhookUpdatePayload) HasNotifyOncePerEmail() bool`

HasNotifyOncePerEmail returns a boolean if a field has been set.

### SetNotifyOncePerEmailNil

`func (o *WebhookUpdatePayload) SetNotifyOncePerEmailNil(b bool)`

 SetNotifyOncePerEmailNil sets the value for NotifyOncePerEmail to be an explicit nil

### UnsetNotifyOncePerEmail
`func (o *WebhookUpdatePayload) UnsetNotifyOncePerEmail()`

UnsetNotifyOncePerEmail ensures that no value is present for NotifyOncePerEmail, not even an explicit nil
### GetNotificationForSent

`func (o *WebhookUpdatePayload) GetNotificationForSent() bool`

GetNotificationForSent returns the NotificationForSent field if non-nil, zero value otherwise.

### GetNotificationForSentOk

`func (o *WebhookUpdatePayload) GetNotificationForSentOk() (*bool, bool)`

GetNotificationForSentOk returns a tuple with the NotificationForSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForSent

`func (o *WebhookUpdatePayload) SetNotificationForSent(v bool)`

SetNotificationForSent sets NotificationForSent field to given value.

### HasNotificationForSent

`func (o *WebhookUpdatePayload) HasNotificationForSent() bool`

HasNotificationForSent returns a boolean if a field has been set.

### SetNotificationForSentNil

`func (o *WebhookUpdatePayload) SetNotificationForSentNil(b bool)`

 SetNotificationForSentNil sets the value for NotificationForSent to be an explicit nil

### UnsetNotificationForSent
`func (o *WebhookUpdatePayload) UnsetNotificationForSent()`

UnsetNotificationForSent ensures that no value is present for NotificationForSent, not even an explicit nil
### GetNotificationForOpened

`func (o *WebhookUpdatePayload) GetNotificationForOpened() bool`

GetNotificationForOpened returns the NotificationForOpened field if non-nil, zero value otherwise.

### GetNotificationForOpenedOk

`func (o *WebhookUpdatePayload) GetNotificationForOpenedOk() (*bool, bool)`

GetNotificationForOpenedOk returns a tuple with the NotificationForOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForOpened

`func (o *WebhookUpdatePayload) SetNotificationForOpened(v bool)`

SetNotificationForOpened sets NotificationForOpened field to given value.

### HasNotificationForOpened

`func (o *WebhookUpdatePayload) HasNotificationForOpened() bool`

HasNotificationForOpened returns a boolean if a field has been set.

### SetNotificationForOpenedNil

`func (o *WebhookUpdatePayload) SetNotificationForOpenedNil(b bool)`

 SetNotificationForOpenedNil sets the value for NotificationForOpened to be an explicit nil

### UnsetNotificationForOpened
`func (o *WebhookUpdatePayload) UnsetNotificationForOpened()`

UnsetNotificationForOpened ensures that no value is present for NotificationForOpened, not even an explicit nil
### GetNotificationForClicked

`func (o *WebhookUpdatePayload) GetNotificationForClicked() bool`

GetNotificationForClicked returns the NotificationForClicked field if non-nil, zero value otherwise.

### GetNotificationForClickedOk

`func (o *WebhookUpdatePayload) GetNotificationForClickedOk() (*bool, bool)`

GetNotificationForClickedOk returns a tuple with the NotificationForClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForClicked

`func (o *WebhookUpdatePayload) SetNotificationForClicked(v bool)`

SetNotificationForClicked sets NotificationForClicked field to given value.

### HasNotificationForClicked

`func (o *WebhookUpdatePayload) HasNotificationForClicked() bool`

HasNotificationForClicked returns a boolean if a field has been set.

### SetNotificationForClickedNil

`func (o *WebhookUpdatePayload) SetNotificationForClickedNil(b bool)`

 SetNotificationForClickedNil sets the value for NotificationForClicked to be an explicit nil

### UnsetNotificationForClicked
`func (o *WebhookUpdatePayload) UnsetNotificationForClicked()`

UnsetNotificationForClicked ensures that no value is present for NotificationForClicked, not even an explicit nil
### GetNotificationForUnsubscribed

`func (o *WebhookUpdatePayload) GetNotificationForUnsubscribed() bool`

GetNotificationForUnsubscribed returns the NotificationForUnsubscribed field if non-nil, zero value otherwise.

### GetNotificationForUnsubscribedOk

`func (o *WebhookUpdatePayload) GetNotificationForUnsubscribedOk() (*bool, bool)`

GetNotificationForUnsubscribedOk returns a tuple with the NotificationForUnsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForUnsubscribed

`func (o *WebhookUpdatePayload) SetNotificationForUnsubscribed(v bool)`

SetNotificationForUnsubscribed sets NotificationForUnsubscribed field to given value.

### HasNotificationForUnsubscribed

`func (o *WebhookUpdatePayload) HasNotificationForUnsubscribed() bool`

HasNotificationForUnsubscribed returns a boolean if a field has been set.

### SetNotificationForUnsubscribedNil

`func (o *WebhookUpdatePayload) SetNotificationForUnsubscribedNil(b bool)`

 SetNotificationForUnsubscribedNil sets the value for NotificationForUnsubscribed to be an explicit nil

### UnsetNotificationForUnsubscribed
`func (o *WebhookUpdatePayload) UnsetNotificationForUnsubscribed()`

UnsetNotificationForUnsubscribed ensures that no value is present for NotificationForUnsubscribed, not even an explicit nil
### GetNotificationForAbuseReport

`func (o *WebhookUpdatePayload) GetNotificationForAbuseReport() bool`

GetNotificationForAbuseReport returns the NotificationForAbuseReport field if non-nil, zero value otherwise.

### GetNotificationForAbuseReportOk

`func (o *WebhookUpdatePayload) GetNotificationForAbuseReportOk() (*bool, bool)`

GetNotificationForAbuseReportOk returns a tuple with the NotificationForAbuseReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForAbuseReport

`func (o *WebhookUpdatePayload) SetNotificationForAbuseReport(v bool)`

SetNotificationForAbuseReport sets NotificationForAbuseReport field to given value.

### HasNotificationForAbuseReport

`func (o *WebhookUpdatePayload) HasNotificationForAbuseReport() bool`

HasNotificationForAbuseReport returns a boolean if a field has been set.

### SetNotificationForAbuseReportNil

`func (o *WebhookUpdatePayload) SetNotificationForAbuseReportNil(b bool)`

 SetNotificationForAbuseReportNil sets the value for NotificationForAbuseReport to be an explicit nil

### UnsetNotificationForAbuseReport
`func (o *WebhookUpdatePayload) UnsetNotificationForAbuseReport()`

UnsetNotificationForAbuseReport ensures that no value is present for NotificationForAbuseReport, not even an explicit nil
### GetNotificationForError

`func (o *WebhookUpdatePayload) GetNotificationForError() bool`

GetNotificationForError returns the NotificationForError field if non-nil, zero value otherwise.

### GetNotificationForErrorOk

`func (o *WebhookUpdatePayload) GetNotificationForErrorOk() (*bool, bool)`

GetNotificationForErrorOk returns a tuple with the NotificationForError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForError

`func (o *WebhookUpdatePayload) SetNotificationForError(v bool)`

SetNotificationForError sets NotificationForError field to given value.

### HasNotificationForError

`func (o *WebhookUpdatePayload) HasNotificationForError() bool`

HasNotificationForError returns a boolean if a field has been set.

### SetNotificationForErrorNil

`func (o *WebhookUpdatePayload) SetNotificationForErrorNil(b bool)`

 SetNotificationForErrorNil sets the value for NotificationForError to be an explicit nil

### UnsetNotificationForError
`func (o *WebhookUpdatePayload) UnsetNotificationForError()`

UnsetNotificationForError ensures that no value is present for NotificationForError, not even an explicit nil
### GetIsEnabled

`func (o *WebhookUpdatePayload) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WebhookUpdatePayload) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WebhookUpdatePayload) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *WebhookUpdatePayload) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### SetIsEnabledNil

`func (o *WebhookUpdatePayload) SetIsEnabledNil(b bool)`

 SetIsEnabledNil sets the value for IsEnabled to be an explicit nil

### UnsetIsEnabled
`func (o *WebhookUpdatePayload) UnsetIsEnabled()`

UnsetIsEnabled ensures that no value is present for IsEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


