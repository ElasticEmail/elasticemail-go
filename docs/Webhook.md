# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookID** | Pointer to **string** | Public webhook ID | [optional] 
**Name** | Pointer to **string** | Filename | [optional] 
**DateCreated** | Pointer to **NullableTime** | Creation date. | [optional] 
**DateUpdated** | Pointer to **NullableTime** | Last change date | [optional] 
**URL** | Pointer to **string** | URL of notification. | [optional] 
**NotifyOncePerEmail** | Pointer to **bool** |  | [optional] 
**NotificationForSent** | Pointer to **bool** |  | [optional] 
**NotificationForOpened** | Pointer to **bool** |  | [optional] 
**NotificationForClicked** | Pointer to **bool** |  | [optional] 
**NotificationForUnsubscribed** | Pointer to **bool** |  | [optional] 
**NotificationForAbuseReport** | Pointer to **bool** |  | [optional] 
**NotificationForError** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewWebhook

`func NewWebhook() *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookID

`func (o *Webhook) GetWebhookID() string`

GetWebhookID returns the WebhookID field if non-nil, zero value otherwise.

### GetWebhookIDOk

`func (o *Webhook) GetWebhookIDOk() (*string, bool)`

GetWebhookIDOk returns a tuple with the WebhookID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookID

`func (o *Webhook) SetWebhookID(v string)`

SetWebhookID sets WebhookID field to given value.

### HasWebhookID

`func (o *Webhook) HasWebhookID() bool`

HasWebhookID returns a boolean if a field has been set.

### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Webhook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDateCreated

`func (o *Webhook) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Webhook) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Webhook) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Webhook) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *Webhook) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *Webhook) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetDateUpdated

`func (o *Webhook) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *Webhook) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *Webhook) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *Webhook) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *Webhook) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *Webhook) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil
### GetURL

`func (o *Webhook) GetURL() string`

GetURL returns the URL field if non-nil, zero value otherwise.

### GetURLOk

`func (o *Webhook) GetURLOk() (*string, bool)`

GetURLOk returns a tuple with the URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetURL

`func (o *Webhook) SetURL(v string)`

SetURL sets URL field to given value.

### HasURL

`func (o *Webhook) HasURL() bool`

HasURL returns a boolean if a field has been set.

### GetNotifyOncePerEmail

`func (o *Webhook) GetNotifyOncePerEmail() bool`

GetNotifyOncePerEmail returns the NotifyOncePerEmail field if non-nil, zero value otherwise.

### GetNotifyOncePerEmailOk

`func (o *Webhook) GetNotifyOncePerEmailOk() (*bool, bool)`

GetNotifyOncePerEmailOk returns a tuple with the NotifyOncePerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOncePerEmail

`func (o *Webhook) SetNotifyOncePerEmail(v bool)`

SetNotifyOncePerEmail sets NotifyOncePerEmail field to given value.

### HasNotifyOncePerEmail

`func (o *Webhook) HasNotifyOncePerEmail() bool`

HasNotifyOncePerEmail returns a boolean if a field has been set.

### GetNotificationForSent

`func (o *Webhook) GetNotificationForSent() bool`

GetNotificationForSent returns the NotificationForSent field if non-nil, zero value otherwise.

### GetNotificationForSentOk

`func (o *Webhook) GetNotificationForSentOk() (*bool, bool)`

GetNotificationForSentOk returns a tuple with the NotificationForSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForSent

`func (o *Webhook) SetNotificationForSent(v bool)`

SetNotificationForSent sets NotificationForSent field to given value.

### HasNotificationForSent

`func (o *Webhook) HasNotificationForSent() bool`

HasNotificationForSent returns a boolean if a field has been set.

### GetNotificationForOpened

`func (o *Webhook) GetNotificationForOpened() bool`

GetNotificationForOpened returns the NotificationForOpened field if non-nil, zero value otherwise.

### GetNotificationForOpenedOk

`func (o *Webhook) GetNotificationForOpenedOk() (*bool, bool)`

GetNotificationForOpenedOk returns a tuple with the NotificationForOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForOpened

`func (o *Webhook) SetNotificationForOpened(v bool)`

SetNotificationForOpened sets NotificationForOpened field to given value.

### HasNotificationForOpened

`func (o *Webhook) HasNotificationForOpened() bool`

HasNotificationForOpened returns a boolean if a field has been set.

### GetNotificationForClicked

`func (o *Webhook) GetNotificationForClicked() bool`

GetNotificationForClicked returns the NotificationForClicked field if non-nil, zero value otherwise.

### GetNotificationForClickedOk

`func (o *Webhook) GetNotificationForClickedOk() (*bool, bool)`

GetNotificationForClickedOk returns a tuple with the NotificationForClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForClicked

`func (o *Webhook) SetNotificationForClicked(v bool)`

SetNotificationForClicked sets NotificationForClicked field to given value.

### HasNotificationForClicked

`func (o *Webhook) HasNotificationForClicked() bool`

HasNotificationForClicked returns a boolean if a field has been set.

### GetNotificationForUnsubscribed

`func (o *Webhook) GetNotificationForUnsubscribed() bool`

GetNotificationForUnsubscribed returns the NotificationForUnsubscribed field if non-nil, zero value otherwise.

### GetNotificationForUnsubscribedOk

`func (o *Webhook) GetNotificationForUnsubscribedOk() (*bool, bool)`

GetNotificationForUnsubscribedOk returns a tuple with the NotificationForUnsubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForUnsubscribed

`func (o *Webhook) SetNotificationForUnsubscribed(v bool)`

SetNotificationForUnsubscribed sets NotificationForUnsubscribed field to given value.

### HasNotificationForUnsubscribed

`func (o *Webhook) HasNotificationForUnsubscribed() bool`

HasNotificationForUnsubscribed returns a boolean if a field has been set.

### GetNotificationForAbuseReport

`func (o *Webhook) GetNotificationForAbuseReport() bool`

GetNotificationForAbuseReport returns the NotificationForAbuseReport field if non-nil, zero value otherwise.

### GetNotificationForAbuseReportOk

`func (o *Webhook) GetNotificationForAbuseReportOk() (*bool, bool)`

GetNotificationForAbuseReportOk returns a tuple with the NotificationForAbuseReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForAbuseReport

`func (o *Webhook) SetNotificationForAbuseReport(v bool)`

SetNotificationForAbuseReport sets NotificationForAbuseReport field to given value.

### HasNotificationForAbuseReport

`func (o *Webhook) HasNotificationForAbuseReport() bool`

HasNotificationForAbuseReport returns a boolean if a field has been set.

### GetNotificationForError

`func (o *Webhook) GetNotificationForError() bool`

GetNotificationForError returns the NotificationForError field if non-nil, zero value otherwise.

### GetNotificationForErrorOk

`func (o *Webhook) GetNotificationForErrorOk() (*bool, bool)`

GetNotificationForErrorOk returns a tuple with the NotificationForError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForError

`func (o *Webhook) SetNotificationForError(v bool)`

SetNotificationForError sets NotificationForError field to given value.

### HasNotificationForError

`func (o *Webhook) HasNotificationForError() bool`

HasNotificationForError returns a boolean if a field has been set.

### GetIsEnabled

`func (o *Webhook) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *Webhook) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *Webhook) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *Webhook) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


