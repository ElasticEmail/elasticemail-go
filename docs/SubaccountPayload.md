# SubaccountPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Proper email address. | 
**Password** | **string** | Current password. | 
**SendActivation** | Pointer to **bool** | True, if you want to send activation email to this Account to confirm the creation of a new SubAccount. Otherwise, false (SubAccount will immediately be Active). | [optional] 
**Settings** | Pointer to [**SubaccountSettingsInfoPayload**](SubaccountSettingsInfoPayload.md) |  | [optional] 

## Methods

### NewSubaccountPayload

`func NewSubaccountPayload(email string, password string, ) *SubaccountPayload`

NewSubaccountPayload instantiates a new SubaccountPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountPayloadWithDefaults

`func NewSubaccountPayloadWithDefaults() *SubaccountPayload`

NewSubaccountPayloadWithDefaults instantiates a new SubaccountPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubaccountPayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubaccountPayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubaccountPayload) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *SubaccountPayload) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SubaccountPayload) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SubaccountPayload) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSendActivation

`func (o *SubaccountPayload) GetSendActivation() bool`

GetSendActivation returns the SendActivation field if non-nil, zero value otherwise.

### GetSendActivationOk

`func (o *SubaccountPayload) GetSendActivationOk() (*bool, bool)`

GetSendActivationOk returns a tuple with the SendActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendActivation

`func (o *SubaccountPayload) SetSendActivation(v bool)`

SetSendActivation sets SendActivation field to given value.

### HasSendActivation

`func (o *SubaccountPayload) HasSendActivation() bool`

HasSendActivation returns a boolean if a field has been set.

### GetSettings

`func (o *SubaccountPayload) GetSettings() SubaccountSettingsInfoPayload`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SubaccountPayload) GetSettingsOk() (*SubaccountSettingsInfoPayload, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SubaccountPayload) SetSettings(v SubaccountSettingsInfoPayload)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SubaccountPayload) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


