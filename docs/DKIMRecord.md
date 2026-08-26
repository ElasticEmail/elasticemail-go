# DKIMRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Selector** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**RecordValue** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** | Name of selected domain. | [optional] 

## Methods

### NewDKIMRecord

`func NewDKIMRecord() *DKIMRecord`

NewDKIMRecord instantiates a new DKIMRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDKIMRecordWithDefaults

`func NewDKIMRecordWithDefaults() *DKIMRecord`

NewDKIMRecordWithDefaults instantiates a new DKIMRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelector

`func (o *DKIMRecord) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *DKIMRecord) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *DKIMRecord) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *DKIMRecord) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetPublicKey

`func (o *DKIMRecord) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DKIMRecord) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DKIMRecord) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DKIMRecord) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetHostName

`func (o *DKIMRecord) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *DKIMRecord) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *DKIMRecord) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *DKIMRecord) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetRecordValue

`func (o *DKIMRecord) GetRecordValue() string`

GetRecordValue returns the RecordValue field if non-nil, zero value otherwise.

### GetRecordValueOk

`func (o *DKIMRecord) GetRecordValueOk() (*string, bool)`

GetRecordValueOk returns a tuple with the RecordValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordValue

`func (o *DKIMRecord) SetRecordValue(v string)`

SetRecordValue sets RecordValue field to given value.

### HasRecordValue

`func (o *DKIMRecord) HasRecordValue() bool`

HasRecordValue returns a boolean if a field has been set.

### GetDomain

`func (o *DKIMRecord) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DKIMRecord) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DKIMRecord) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DKIMRecord) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


