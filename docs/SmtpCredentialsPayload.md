# SmtpCredentialsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Credential for ease of reference. It must be a valid email address. | 
**Expires** | Pointer to **NullableTime** | Date this SmtpCredential expires. | [optional] 
**RestrictAccessToIPRange** | Pointer to **[]string** | Which IPs can use this SmtpCredential | [optional] 
**Subaccount** | Pointer to **string** | Email of the subaccount for which this SmtpCredential should be created | [optional] 

## Methods

### NewSmtpCredentialsPayload

`func NewSmtpCredentialsPayload(name string, ) *SmtpCredentialsPayload`

NewSmtpCredentialsPayload instantiates a new SmtpCredentialsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpCredentialsPayloadWithDefaults

`func NewSmtpCredentialsPayloadWithDefaults() *SmtpCredentialsPayload`

NewSmtpCredentialsPayloadWithDefaults instantiates a new SmtpCredentialsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SmtpCredentialsPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmtpCredentialsPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmtpCredentialsPayload) SetName(v string)`

SetName sets Name field to given value.


### GetExpires

`func (o *SmtpCredentialsPayload) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SmtpCredentialsPayload) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SmtpCredentialsPayload) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SmtpCredentialsPayload) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *SmtpCredentialsPayload) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *SmtpCredentialsPayload) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetRestrictAccessToIPRange

`func (o *SmtpCredentialsPayload) GetRestrictAccessToIPRange() []string`

GetRestrictAccessToIPRange returns the RestrictAccessToIPRange field if non-nil, zero value otherwise.

### GetRestrictAccessToIPRangeOk

`func (o *SmtpCredentialsPayload) GetRestrictAccessToIPRangeOk() (*[]string, bool)`

GetRestrictAccessToIPRangeOk returns a tuple with the RestrictAccessToIPRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictAccessToIPRange

`func (o *SmtpCredentialsPayload) SetRestrictAccessToIPRange(v []string)`

SetRestrictAccessToIPRange sets RestrictAccessToIPRange field to given value.

### HasRestrictAccessToIPRange

`func (o *SmtpCredentialsPayload) HasRestrictAccessToIPRange() bool`

HasRestrictAccessToIPRange returns a boolean if a field has been set.

### GetSubaccount

`func (o *SmtpCredentialsPayload) GetSubaccount() string`

GetSubaccount returns the Subaccount field if non-nil, zero value otherwise.

### GetSubaccountOk

`func (o *SmtpCredentialsPayload) GetSubaccountOk() (*string, bool)`

GetSubaccountOk returns a tuple with the Subaccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaccount

`func (o *SmtpCredentialsPayload) SetSubaccount(v string)`

SetSubaccount sets Subaccount field to given value.

### HasSubaccount

`func (o *SmtpCredentialsPayload) HasSubaccount() bool`

HasSubaccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


