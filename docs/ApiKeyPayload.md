# ApiKeyPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the ApiKey for ease of reference. | 
**AccessLevel** | [**[]AccessLevel**](AccessLevel.md) | Access level or permission to be assigned to this ApiKey. | 
**Expires** | Pointer to **NullableTime** | Date this ApiKey expires. | [optional] 
**RestrictAccessToIPRange** | Pointer to **[]string** | Which IPs can use this ApiKey | [optional] 
**Subaccount** | Pointer to **string** | Email of the subaccount for which this ApiKey should be created | [optional] 

## Methods

### NewApiKeyPayload

`func NewApiKeyPayload(name string, accessLevel []AccessLevel, ) *ApiKeyPayload`

NewApiKeyPayload instantiates a new ApiKeyPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyPayloadWithDefaults

`func NewApiKeyPayloadWithDefaults() *ApiKeyPayload`

NewApiKeyPayloadWithDefaults instantiates a new ApiKeyPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiKeyPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyPayload) SetName(v string)`

SetName sets Name field to given value.


### GetAccessLevel

`func (o *ApiKeyPayload) GetAccessLevel() []AccessLevel`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ApiKeyPayload) GetAccessLevelOk() (*[]AccessLevel, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ApiKeyPayload) SetAccessLevel(v []AccessLevel)`

SetAccessLevel sets AccessLevel field to given value.


### GetExpires

`func (o *ApiKeyPayload) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *ApiKeyPayload) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *ApiKeyPayload) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *ApiKeyPayload) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *ApiKeyPayload) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *ApiKeyPayload) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetRestrictAccessToIPRange

`func (o *ApiKeyPayload) GetRestrictAccessToIPRange() []string`

GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field if non-nil, zero value otherwise.

### GetRestrictAccessToIPRangeOk

`func (o *ApiKeyPayload) GetRestrictAccessToIPRangeOk() (*[]string, bool)`

GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAccessToIPRange

`func (o *ApiKeyPayload) SetRestrictAccessToIPRange(v []string)`

SetRestrictAccessToIPRange sets RestrictAccessToIPRange field to given value.

### HasRestrictAccessToIPRange

`func (o *ApiKeyPayload) HasRestrictAccessToIPRange() bool`

HasRestrictAccessToIPRange returns a boolean if a field has been set.

### GetSubaccount

`func (o *ApiKeyPayload) GetSubaccount() string`

GetSubaccount returns the Subaccount field if non-nil, zero value otherwise.

### GetSubaccountOk

`func (o *ApiKeyPayload) GetSubaccountOk() (*string, bool)`

GetSubaccountOk returns a tuple with the Subaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaccount

`func (o *ApiKeyPayload) SetSubaccount(v string)`

SetSubaccount sets Subaccount field to given value.

### HasSubaccount

`func (o *ApiKeyPayload) HasSubaccount() bool`

HasSubaccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


