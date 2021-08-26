# SubaccountEmailCreditsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credits** | **int32** | Positive or negative value; this will be added or subtracted from Subaccount&#39;s current email Credits pool. | 
**Notes** | Pointer to **string** | Note to append to this credits change, for history. | [optional] 

## Methods

### NewSubaccountEmailCreditsPayload

`func NewSubaccountEmailCreditsPayload(credits int32, ) *SubaccountEmailCreditsPayload`

NewSubaccountEmailCreditsPayload instantiates a new SubaccountEmailCreditsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubaccountEmailCreditsPayloadWithDefaults

`func NewSubaccountEmailCreditsPayloadWithDefaults() *SubaccountEmailCreditsPayload`

NewSubaccountEmailCreditsPayloadWithDefaults instantiates a new SubaccountEmailCreditsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredits

`func (o *SubaccountEmailCreditsPayload) GetCredits() int32`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *SubaccountEmailCreditsPayload) GetCreditsOk() (*int32, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *SubaccountEmailCreditsPayload) SetCredits(v int32)`

SetCredits sets Credits field to given value.


### GetNotes

`func (o *SubaccountEmailCreditsPayload) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SubaccountEmailCreditsPayload) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SubaccountEmailCreditsPayload) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SubaccountEmailCreditsPayload) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


