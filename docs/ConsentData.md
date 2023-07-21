# ConsentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsentIP** | Pointer to **string** | IP address of consent to send this contact(s) your email. If not provided your current public IP address is used for consent. | [optional] 
**ConsentDate** | Pointer to **NullableTime** | Date of consent to send this contact(s) your email. If not provided current date is used for consent. | [optional] 
**ConsentTracking** | Pointer to [**ConsentTracking**](ConsentTracking.md) |  | [optional] [default to CONSENTTRACKING_UNKNOWN]

## Methods

### NewConsentData

`func NewConsentData() *ConsentData`

NewConsentData instantiates a new ConsentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentDataWithDefaults

`func NewConsentDataWithDefaults() *ConsentData`

NewConsentDataWithDefaults instantiates a new ConsentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsentIP

`func (o *ConsentData) GetConsentIP() string`

GetConsentIP returns the ConsentIP field if non-nil, zero value otherwise.

### GetConsentIPOk

`func (o *ConsentData) GetConsentIPOk() (*string, bool)`

GetConsentIPOk returns a tuple with the ConsentIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentIP

`func (o *ConsentData) SetConsentIP(v string)`

SetConsentIP sets ConsentIP field to given value.

### HasConsentIP

`func (o *ConsentData) HasConsentIP() bool`

HasConsentIP returns a boolean if a field has been set.

### GetConsentDate

`func (o *ConsentData) GetConsentDate() time.Time`

GetConsentDate returns the ConsentDate field if non-nil, zero value otherwise.

### GetConsentDateOk

`func (o *ConsentData) GetConsentDateOk() (*time.Time, bool)`

GetConsentDateOk returns a tuple with the ConsentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentDate

`func (o *ConsentData) SetConsentDate(v time.Time)`

SetConsentDate sets ConsentDate field to given value.

### HasConsentDate

`func (o *ConsentData) HasConsentDate() bool`

HasConsentDate returns a boolean if a field has been set.

### SetConsentDateNil

`func (o *ConsentData) SetConsentDateNil(b bool)`

 SetConsentDateNil sets the value for ConsentDate to be an explicit nil

### UnsetConsentDate
`func (o *ConsentData) UnsetConsentDate()`

UnsetConsentDate ensures that no value is present for ConsentDate, not even an explicit nil
### GetConsentTracking

`func (o *ConsentData) GetConsentTracking() ConsentTracking`

GetConsentTracking returns the ConsentTracking field if non-nil, zero value otherwise.

### GetConsentTrackingOk

`func (o *ConsentData) GetConsentTrackingOk() (*ConsentTracking, bool)`

GetConsentTrackingOk returns a tuple with the ConsentTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentTracking

`func (o *ConsentData) SetConsentTracking(v ConsentTracking)`

SetConsentTracking sets ConsentTracking field to given value.

### HasConsentTracking

`func (o *ConsentData) HasConsentTracking() bool`

HasConsentTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


