# CreateScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertOnFailure** | Pointer to **bool** |  | [optional] [default to false]
**AutoPauseAfterFailures** | Pointer to **NullableInt32** |  | [optional] 
**Destinations** | Pointer to [**[]DeliveryDestination**](DeliveryDestination.md) |  | [optional] 
**DiffThreshold** | Pointer to **NullableFloat64** |  | [optional] 
**EndsAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**OnlyOnChange** | Pointer to **bool** |  | [optional] [default to false]
**Options** | Pointer to [**NullableScheduleScreenshotOptions**](ScheduleScreenshotOptions.md) |  | [optional] 
**RetentionDays** | Pointer to **int32** |  | [optional] [default to 30]
**Schedule** | **string** |  | 
**StartsAt** | Pointer to **NullableTime** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] [default to "UTC"]
**Url** | **string** |  | 
**WebhookSecret** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateScheduleRequest

`func NewCreateScheduleRequest(name string, schedule string, url string, ) *CreateScheduleRequest`

NewCreateScheduleRequest instantiates a new CreateScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduleRequestWithDefaults

`func NewCreateScheduleRequestWithDefaults() *CreateScheduleRequest`

NewCreateScheduleRequestWithDefaults instantiates a new CreateScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertOnFailure

`func (o *CreateScheduleRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *CreateScheduleRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *CreateScheduleRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *CreateScheduleRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetAutoPauseAfterFailures

`func (o *CreateScheduleRequest) GetAutoPauseAfterFailures() int32`

GetAutoPauseAfterFailures returns the AutoPauseAfterFailures field if non-nil, zero value otherwise.

### GetAutoPauseAfterFailuresOk

`func (o *CreateScheduleRequest) GetAutoPauseAfterFailuresOk() (*int32, bool)`

GetAutoPauseAfterFailuresOk returns a tuple with the AutoPauseAfterFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPauseAfterFailures

`func (o *CreateScheduleRequest) SetAutoPauseAfterFailures(v int32)`

SetAutoPauseAfterFailures sets AutoPauseAfterFailures field to given value.

### HasAutoPauseAfterFailures

`func (o *CreateScheduleRequest) HasAutoPauseAfterFailures() bool`

HasAutoPauseAfterFailures returns a boolean if a field has been set.

### SetAutoPauseAfterFailuresNil

`func (o *CreateScheduleRequest) SetAutoPauseAfterFailuresNil(b bool)`

 SetAutoPauseAfterFailuresNil sets the value for AutoPauseAfterFailures to be an explicit nil

### UnsetAutoPauseAfterFailures
`func (o *CreateScheduleRequest) UnsetAutoPauseAfterFailures()`

UnsetAutoPauseAfterFailures ensures that no value is present for AutoPauseAfterFailures, not even an explicit nil
### GetDestinations

`func (o *CreateScheduleRequest) GetDestinations() []DeliveryDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *CreateScheduleRequest) GetDestinationsOk() (*[]DeliveryDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *CreateScheduleRequest) SetDestinations(v []DeliveryDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *CreateScheduleRequest) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### SetDestinationsNil

`func (o *CreateScheduleRequest) SetDestinationsNil(b bool)`

 SetDestinationsNil sets the value for Destinations to be an explicit nil

### UnsetDestinations
`func (o *CreateScheduleRequest) UnsetDestinations()`

UnsetDestinations ensures that no value is present for Destinations, not even an explicit nil
### GetDiffThreshold

`func (o *CreateScheduleRequest) GetDiffThreshold() float64`

GetDiffThreshold returns the DiffThreshold field if non-nil, zero value otherwise.

### GetDiffThresholdOk

`func (o *CreateScheduleRequest) GetDiffThresholdOk() (*float64, bool)`

GetDiffThresholdOk returns a tuple with the DiffThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffThreshold

`func (o *CreateScheduleRequest) SetDiffThreshold(v float64)`

SetDiffThreshold sets DiffThreshold field to given value.

### HasDiffThreshold

`func (o *CreateScheduleRequest) HasDiffThreshold() bool`

HasDiffThreshold returns a boolean if a field has been set.

### SetDiffThresholdNil

`func (o *CreateScheduleRequest) SetDiffThresholdNil(b bool)`

 SetDiffThresholdNil sets the value for DiffThreshold to be an explicit nil

### UnsetDiffThreshold
`func (o *CreateScheduleRequest) UnsetDiffThreshold()`

UnsetDiffThreshold ensures that no value is present for DiffThreshold, not even an explicit nil
### GetEndsAt

`func (o *CreateScheduleRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateScheduleRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateScheduleRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *CreateScheduleRequest) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *CreateScheduleRequest) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *CreateScheduleRequest) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetName

`func (o *CreateScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateScheduleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnlyOnChange

`func (o *CreateScheduleRequest) GetOnlyOnChange() bool`

GetOnlyOnChange returns the OnlyOnChange field if non-nil, zero value otherwise.

### GetOnlyOnChangeOk

`func (o *CreateScheduleRequest) GetOnlyOnChangeOk() (*bool, bool)`

GetOnlyOnChangeOk returns a tuple with the OnlyOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyOnChange

`func (o *CreateScheduleRequest) SetOnlyOnChange(v bool)`

SetOnlyOnChange sets OnlyOnChange field to given value.

### HasOnlyOnChange

`func (o *CreateScheduleRequest) HasOnlyOnChange() bool`

HasOnlyOnChange returns a boolean if a field has been set.

### GetOptions

`func (o *CreateScheduleRequest) GetOptions() ScheduleScreenshotOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateScheduleRequest) GetOptionsOk() (*ScheduleScreenshotOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateScheduleRequest) SetOptions(v ScheduleScreenshotOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateScheduleRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *CreateScheduleRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *CreateScheduleRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetRetentionDays

`func (o *CreateScheduleRequest) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *CreateScheduleRequest) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *CreateScheduleRequest) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *CreateScheduleRequest) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetSchedule

`func (o *CreateScheduleRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateScheduleRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateScheduleRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetStartsAt

`func (o *CreateScheduleRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CreateScheduleRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CreateScheduleRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *CreateScheduleRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *CreateScheduleRequest) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *CreateScheduleRequest) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetTimezone

`func (o *CreateScheduleRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreateScheduleRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreateScheduleRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreateScheduleRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUrl

`func (o *CreateScheduleRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateScheduleRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateScheduleRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebhookSecret

`func (o *CreateScheduleRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *CreateScheduleRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *CreateScheduleRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *CreateScheduleRequest) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *CreateScheduleRequest) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *CreateScheduleRequest) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetWebhookUrl

`func (o *CreateScheduleRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateScheduleRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateScheduleRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateScheduleRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *CreateScheduleRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *CreateScheduleRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


