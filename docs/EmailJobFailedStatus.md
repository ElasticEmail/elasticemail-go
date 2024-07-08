# EmailJobFailedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **int32** | RFC Error code | [optional] 
**Category** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailJobFailedStatus

`func NewEmailJobFailedStatus() *EmailJobFailedStatus`

NewEmailJobFailedStatus instantiates a new EmailJobFailedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailJobFailedStatusWithDefaults

`func NewEmailJobFailedStatusWithDefaults() *EmailJobFailedStatus`

NewEmailJobFailedStatusWithDefaults instantiates a new EmailJobFailedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EmailJobFailedStatus) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmailJobFailedStatus) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmailJobFailedStatus) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *EmailJobFailedStatus) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetError

`func (o *EmailJobFailedStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EmailJobFailedStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EmailJobFailedStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EmailJobFailedStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *EmailJobFailedStatus) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *EmailJobFailedStatus) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *EmailJobFailedStatus) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *EmailJobFailedStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetCategory

`func (o *EmailJobFailedStatus) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *EmailJobFailedStatus) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *EmailJobFailedStatus) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *EmailJobFailedStatus) HasCategory() bool`

HasCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


