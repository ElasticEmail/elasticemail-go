# DomainPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Name of selected domain. | [optional] 
**SetAsDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewDomainPayload

`func NewDomainPayload() *DomainPayload`

NewDomainPayload instantiates a new DomainPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPayloadWithDefaults

`func NewDomainPayloadWithDefaults() *DomainPayload`

NewDomainPayloadWithDefaults instantiates a new DomainPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainPayload) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainPayload) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainPayload) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DomainPayload) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSetAsDefault

`func (o *DomainPayload) GetSetAsDefault() bool`

GetSetAsDefault returns the SetAsDefault field if non-nil, zero value otherwise.

### GetSetAsDefaultOk

`func (o *DomainPayload) GetSetAsDefaultOk() (*bool, bool)`

GetSetAsDefaultOk returns a tuple with the SetAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAsDefault

`func (o *DomainPayload) SetSetAsDefault(v bool)`

SetSetAsDefault sets SetAsDefault field to given value.

### HasSetAsDefault

`func (o *DomainPayload) HasSetAsDefault() bool`

HasSetAsDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


