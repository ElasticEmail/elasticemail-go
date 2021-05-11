# Suppression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Proper email address. | [optional] 
**FriendlyErrorMessage** | Pointer to **string** | RFC error message | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | SMTP Error code | [optional] 
**DateUpdated** | Pointer to **NullableTime** | Last change date | [optional] 

## Methods

### NewSuppression

`func NewSuppression() *Suppression`

NewSuppression instantiates a new Suppression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuppressionWithDefaults

`func NewSuppressionWithDefaults() *Suppression`

NewSuppressionWithDefaults instantiates a new Suppression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Suppression) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Suppression) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Suppression) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Suppression) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFriendlyErrorMessage

`func (o *Suppression) GetFriendlyErrorMessage() string`

GetFriendlyErrorMessage returns the FriendlyErrorMessage field if non-nil, zero value otherwise.

### GetFriendlyErrorMessageOk

`func (o *Suppression) GetFriendlyErrorMessageOk() (*string, bool)`

GetFriendlyErrorMessageOk returns a tuple with the FriendlyErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyErrorMessage

`func (o *Suppression) SetFriendlyErrorMessage(v string)`

SetFriendlyErrorMessage sets FriendlyErrorMessage field to given value.

### HasFriendlyErrorMessage

`func (o *Suppression) HasFriendlyErrorMessage() bool`

HasFriendlyErrorMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *Suppression) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Suppression) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Suppression) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Suppression) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *Suppression) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *Suppression) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetDateUpdated

`func (o *Suppression) GetDateUpdated() time.Time`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *Suppression) GetDateUpdatedOk() (*time.Time, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *Suppression) SetDateUpdated(v time.Time)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *Suppression) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### SetDateUpdatedNil

`func (o *Suppression) SetDateUpdatedNil(b bool)`

 SetDateUpdatedNil sets the value for DateUpdated to be an explicit nil

### UnsetDateUpdated
`func (o *Suppression) UnsetDateUpdated()`

UnsetDateUpdated ensures that no value is present for DateUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


