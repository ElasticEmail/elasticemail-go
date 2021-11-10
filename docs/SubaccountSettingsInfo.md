# SubaccountSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**SubaccountEmailSettings**](SubaccountEmailSettings.md) |  | [optional] 

## Methods

### NewSubaccountSettingsInfo

`func NewSubaccountSettingsInfo() *SubaccountSettingsInfo`

NewSubaccountSettingsInfo instantiates a new SubaccountSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountSettingsInfoWithDefaults

`func NewSubaccountSettingsInfoWithDefaults() *SubaccountSettingsInfo`

NewSubaccountSettingsInfoWithDefaults instantiates a new SubaccountSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SubaccountSettingsInfo) GetEmail() SubaccountEmailSettings`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubaccountSettingsInfo) GetEmailOk() (*SubaccountEmailSettings, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubaccountSettingsInfo) SetEmail(v SubaccountEmailSettings)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubaccountSettingsInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


