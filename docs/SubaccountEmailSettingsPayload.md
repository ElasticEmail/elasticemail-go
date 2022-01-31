# SubaccountEmailSettingsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiresEmailCredits** | Pointer to **bool** | True, if Account needs credits to send emails. Otherwise, false | [optional] 
**EmailSizeLimit** | Pointer to **int32** | Maximum size of email including attachments in MB&#39;s | [optional] 
**DailySendLimit** | Pointer to **int32** | Amount of emails Account can send daily | [optional] 
**MaxContacts** | Pointer to **int32** | Maximum number of contacts the Account can have. 0 means that parent account&#39;s limit is used. | [optional] 
**EnablePrivateIPPurchase** | Pointer to **bool** | Can the SubAccount purchase Private IP for themselves | [optional] 
**PoolName** | Pointer to **string** | Name of your custom IP Pool to be used in the sending process | [optional] 
**ValidSenderDomainOnly** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSubaccountEmailSettingsPayload

`func NewSubaccountEmailSettingsPayload() *SubaccountEmailSettingsPayload`

NewSubaccountEmailSettingsPayload instantiates a new SubaccountEmailSettingsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountEmailSettingsPayloadWithDefaults

`func NewSubaccountEmailSettingsPayloadWithDefaults() *SubaccountEmailSettingsPayload`

NewSubaccountEmailSettingsPayloadWithDefaults instantiates a new SubaccountEmailSettingsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiresEmailCredits

`func (o *SubaccountEmailSettingsPayload) GetRequiresEmailCredits() bool`

GetRequiresEmailCredits returns the RequiresEmailCredits field if non-nil, zero value otherwise.

### GetRequiresEmailCreditsOk

`func (o *SubaccountEmailSettingsPayload) GetRequiresEmailCreditsOk() (*bool, bool)`

GetRequiresEmailCreditsOk returns a tuple with the RequiresEmailCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresEmailCredits

`func (o *SubaccountEmailSettingsPayload) SetRequiresEmailCredits(v bool)`

SetRequiresEmailCredits sets RequiresEmailCredits field to given value.

### HasRequiresEmailCredits

`func (o *SubaccountEmailSettingsPayload) HasRequiresEmailCredits() bool`

HasRequiresEmailCredits returns a boolean if a field has been set.

### GetEmailSizeLimit

`func (o *SubaccountEmailSettingsPayload) GetEmailSizeLimit() int32`

GetEmailSizeLimit returns the EmailSizeLimit field if non-nil, zero value otherwise.

### GetEmailSizeLimitOk

`func (o *SubaccountEmailSettingsPayload) GetEmailSizeLimitOk() (*int32, bool)`

GetEmailSizeLimitOk returns a tuple with the EmailSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSizeLimit

`func (o *SubaccountEmailSettingsPayload) SetEmailSizeLimit(v int32)`

SetEmailSizeLimit sets EmailSizeLimit field to given value.

### HasEmailSizeLimit

`func (o *SubaccountEmailSettingsPayload) HasEmailSizeLimit() bool`

HasEmailSizeLimit returns a boolean if a field has been set.

### GetDailySendLimit

`func (o *SubaccountEmailSettingsPayload) GetDailySendLimit() int32`

GetDailySendLimit returns the DailySendLimit field if non-nil, zero value otherwise.

### GetDailySendLimitOk

`func (o *SubaccountEmailSettingsPayload) GetDailySendLimitOk() (*int32, bool)`

GetDailySendLimitOk returns a tuple with the DailySendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySendLimit

`func (o *SubaccountEmailSettingsPayload) SetDailySendLimit(v int32)`

SetDailySendLimit sets DailySendLimit field to given value.

### HasDailySendLimit

`func (o *SubaccountEmailSettingsPayload) HasDailySendLimit() bool`

HasDailySendLimit returns a boolean if a field has been set.

### GetMaxContacts

`func (o *SubaccountEmailSettingsPayload) GetMaxContacts() int32`

GetMaxContacts returns the MaxContacts field if non-nil, zero value otherwise.

### GetMaxContactsOk

`func (o *SubaccountEmailSettingsPayload) GetMaxContactsOk() (*int32, bool)`

GetMaxContactsOk returns a tuple with the MaxContacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContacts

`func (o *SubaccountEmailSettingsPayload) SetMaxContacts(v int32)`

SetMaxContacts sets MaxContacts field to given value.

### HasMaxContacts

`func (o *SubaccountEmailSettingsPayload) HasMaxContacts() bool`

HasMaxContacts returns a boolean if a field has been set.

### GetEnablePrivateIPPurchase

`func (o *SubaccountEmailSettingsPayload) GetEnablePrivateIPPurchase() bool`

GetEnablePrivateIPPurchase returns the EnablePrivateIPPurchase field if non-nil, zero value otherwise.

### GetEnablePrivateIPPurchaseOk

`func (o *SubaccountEmailSettingsPayload) GetEnablePrivateIPPurchaseOk() (*bool, bool)`

GetEnablePrivateIPPurchaseOk returns a tuple with the EnablePrivateIPPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivateIPPurchase

`func (o *SubaccountEmailSettingsPayload) SetEnablePrivateIPPurchase(v bool)`

SetEnablePrivateIPPurchase sets EnablePrivateIPPurchase field to given value.

### HasEnablePrivateIPPurchase

`func (o *SubaccountEmailSettingsPayload) HasEnablePrivateIPPurchase() bool`

HasEnablePrivateIPPurchase returns a boolean if a field has been set.

### GetPoolName

`func (o *SubaccountEmailSettingsPayload) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *SubaccountEmailSettingsPayload) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *SubaccountEmailSettingsPayload) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *SubaccountEmailSettingsPayload) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetValidSenderDomainOnly

`func (o *SubaccountEmailSettingsPayload) GetValidSenderDomainOnly() bool`

GetValidSenderDomainOnly returns the ValidSenderDomainOnly field if non-nil, zero value otherwise.

### GetValidSenderDomainOnlyOk

`func (o *SubaccountEmailSettingsPayload) GetValidSenderDomainOnlyOk() (*bool, bool)`

GetValidSenderDomainOnlyOk returns a tuple with the ValidSenderDomainOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidSenderDomainOnly

`func (o *SubaccountEmailSettingsPayload) SetValidSenderDomainOnly(v bool)`

SetValidSenderDomainOnly sets ValidSenderDomainOnly field to given value.

### HasValidSenderDomainOnly

`func (o *SubaccountEmailSettingsPayload) HasValidSenderDomainOnly() bool`

HasValidSenderDomainOnly returns a boolean if a field has been set.

### SetValidSenderDomainOnlyNil

`func (o *SubaccountEmailSettingsPayload) SetValidSenderDomainOnlyNil(b bool)`

 SetValidSenderDomainOnlyNil sets the value for ValidSenderDomainOnly to be an explicit nil

### UnsetValidSenderDomainOnly
`func (o *SubaccountEmailSettingsPayload) UnsetValidSenderDomainOnly()`

UnsetValidSenderDomainOnly ensures that no value is present for ValidSenderDomainOnly, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


