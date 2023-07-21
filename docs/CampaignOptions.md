# CampaignOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryOptimization** | Pointer to [**DeliveryOptimizationType**](DeliveryOptimizationType.md) |  | [optional] [default to DELIVERYOPTIMIZATIONTYPE_NONE]
**TrackOpens** | Pointer to **NullableBool** | Should the opens be tracked? If no value has been provided, Account&#39;s default setting will be used. | [optional] 
**TrackClicks** | Pointer to **NullableBool** | Should the clicks be tracked? If no value has been provided, Account&#39;s default setting will be used. | [optional] 
**ScheduleFor** | Pointer to **NullableTime** | Date when this Campaign is scheduled to be sent on | [optional] 
**SplitOptions** | Pointer to [**SplitOptions**](SplitOptions.md) |  | [optional] 

## Methods

### NewCampaignOptions

`func NewCampaignOptions() *CampaignOptions`

NewCampaignOptions instantiates a new CampaignOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignOptionsWithDefaults

`func NewCampaignOptionsWithDefaults() *CampaignOptions`

NewCampaignOptionsWithDefaults instantiates a new CampaignOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryOptimization

`func (o *CampaignOptions) GetDeliveryOptimization() DeliveryOptimizationType`

GetDeliveryOptimization returns the DeliveryOptimization field if non-nil, zero value otherwise.

### GetDeliveryOptimizationOk

`func (o *CampaignOptions) GetDeliveryOptimizationOk() (*DeliveryOptimizationType, bool)`

GetDeliveryOptimizationOk returns a tuple with the DeliveryOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOptimization

`func (o *CampaignOptions) SetDeliveryOptimization(v DeliveryOptimizationType)`

SetDeliveryOptimization sets DeliveryOptimization field to given value.

### HasDeliveryOptimization

`func (o *CampaignOptions) HasDeliveryOptimization() bool`

HasDeliveryOptimization returns a boolean if a field has been set.

### GetTrackOpens

`func (o *CampaignOptions) GetTrackOpens() bool`

GetTrackOpens returns the TrackOpens field if non-nil, zero value otherwise.

### GetTrackOpensOk

`func (o *CampaignOptions) GetTrackOpensOk() (*bool, bool)`

GetTrackOpensOk returns a tuple with the TrackOpens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackOpens

`func (o *CampaignOptions) SetTrackOpens(v bool)`

SetTrackOpens sets TrackOpens field to given value.

### HasTrackOpens

`func (o *CampaignOptions) HasTrackOpens() bool`

HasTrackOpens returns a boolean if a field has been set.

### SetTrackOpensNil

`func (o *CampaignOptions) SetTrackOpensNil(b bool)`

 SetTrackOpensNil sets the value for TrackOpens to be an explicit nil

### UnsetTrackOpens
`func (o *CampaignOptions) UnsetTrackOpens()`

UnsetTrackOpens ensures that no value is present for TrackOpens, not even an explicit nil
### GetTrackClicks

`func (o *CampaignOptions) GetTrackClicks() bool`

GetTrackClicks returns the TrackClicks field if non-nil, zero value otherwise.

### GetTrackClicksOk

`func (o *CampaignOptions) GetTrackClicksOk() (*bool, bool)`

GetTrackClicksOk returns a tuple with the TrackClicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackClicks

`func (o *CampaignOptions) SetTrackClicks(v bool)`

SetTrackClicks sets TrackClicks field to given value.

### HasTrackClicks

`func (o *CampaignOptions) HasTrackClicks() bool`

HasTrackClicks returns a boolean if a field has been set.

### SetTrackClicksNil

`func (o *CampaignOptions) SetTrackClicksNil(b bool)`

 SetTrackClicksNil sets the value for TrackClicks to be an explicit nil

### UnsetTrackClicks
`func (o *CampaignOptions) UnsetTrackClicks()`

UnsetTrackClicks ensures that no value is present for TrackClicks, not even an explicit nil
### GetScheduleFor

`func (o *CampaignOptions) GetScheduleFor() time.Time`

GetScheduleFor returns the ScheduleFor field if non-nil, zero value otherwise.

### GetScheduleForOk

`func (o *CampaignOptions) GetScheduleForOk() (*time.Time, bool)`

GetScheduleForOk returns a tuple with the ScheduleFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleFor

`func (o *CampaignOptions) SetScheduleFor(v time.Time)`

SetScheduleFor sets ScheduleFor field to given value.

### HasScheduleFor

`func (o *CampaignOptions) HasScheduleFor() bool`

HasScheduleFor returns a boolean if a field has been set.

### SetScheduleForNil

`func (o *CampaignOptions) SetScheduleForNil(b bool)`

 SetScheduleForNil sets the value for ScheduleFor to be an explicit nil

### UnsetScheduleFor
`func (o *CampaignOptions) UnsetScheduleFor()`

UnsetScheduleFor ensures that no value is present for ScheduleFor, not even an explicit nil
### GetSplitOptions

`func (o *CampaignOptions) GetSplitOptions() SplitOptions`

GetSplitOptions returns the SplitOptions field if non-nil, zero value otherwise.

### GetSplitOptionsOk

`func (o *CampaignOptions) GetSplitOptionsOk() (*SplitOptions, bool)`

GetSplitOptionsOk returns a tuple with the SplitOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitOptions

`func (o *CampaignOptions) SetSplitOptions(v SplitOptions)`

SetSplitOptions sets SplitOptions field to given value.

### HasSplitOptions

`func (o *CampaignOptions) HasSplitOptions() bool`

HasSplitOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


