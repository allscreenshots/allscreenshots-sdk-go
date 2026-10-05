# UpdateScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertOnFailure** | Pointer to **NullableBool** |  | [optional] 
**AutoPauseAfterFailures** | Pointer to **NullableInt32** |  | [optional] 
**Destinations** | Pointer to [**[]DeliveryDestination**](DeliveryDestination.md) |  | [optional] 
**DiffThreshold** | Pointer to **NullableFloat64** |  | [optional] 
**EndsAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OnlyOnChange** | Pointer to **NullableBool** |  | [optional] 
**Options** | Pointer to [**NullableScheduleScreenshotOptions**](ScheduleScreenshotOptions.md) |  | [optional] 
**RetentionDays** | Pointer to **NullableInt32** |  | [optional] 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**StartsAt** | Pointer to **NullableTime** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**WebhookSecret** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateScheduleRequest

`func NewUpdateScheduleRequest() *UpdateScheduleRequest`

NewUpdateScheduleRequest instantiates a new UpdateScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScheduleRequestWithDefaults

`func NewUpdateScheduleRequestWithDefaults() *UpdateScheduleRequest`

NewUpdateScheduleRequestWithDefaults instantiates a new UpdateScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertOnFailure

`func (o *UpdateScheduleRequest) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *UpdateScheduleRequest) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *UpdateScheduleRequest) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *UpdateScheduleRequest) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### SetAlertOnFailureNil

`func (o *UpdateScheduleRequest) SetAlertOnFailureNil(b bool)`

 SetAlertOnFailureNil sets the value for AlertOnFailure to be an explicit nil

### UnsetAlertOnFailure
`func (o *UpdateScheduleRequest) UnsetAlertOnFailure()`

UnsetAlertOnFailure ensures that no value is present for AlertOnFailure, not even an explicit nil
### GetAutoPauseAfterFailures

`func (o *UpdateScheduleRequest) GetAutoPauseAfterFailures() int32`

GetAutoPauseAfterFailures returns the AutoPauseAfterFailures field if non-nil, zero value otherwise.

### GetAutoPauseAfterFailuresOk

`func (o *UpdateScheduleRequest) GetAutoPauseAfterFailuresOk() (*int32, bool)`

GetAutoPauseAfterFailuresOk returns a tuple with the AutoPauseAfterFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPauseAfterFailures

`func (o *UpdateScheduleRequest) SetAutoPauseAfterFailures(v int32)`

SetAutoPauseAfterFailures sets AutoPauseAfterFailures field to given value.

### HasAutoPauseAfterFailures

`func (o *UpdateScheduleRequest) HasAutoPauseAfterFailures() bool`

HasAutoPauseAfterFailures returns a boolean if a field has been set.

### SetAutoPauseAfterFailuresNil

`func (o *UpdateScheduleRequest) SetAutoPauseAfterFailuresNil(b bool)`

 SetAutoPauseAfterFailuresNil sets the value for AutoPauseAfterFailures to be an explicit nil

### UnsetAutoPauseAfterFailures
`func (o *UpdateScheduleRequest) UnsetAutoPauseAfterFailures()`

UnsetAutoPauseAfterFailures ensures that no value is present for AutoPauseAfterFailures, not even an explicit nil
### GetDestinations

`func (o *UpdateScheduleRequest) GetDestinations() []DeliveryDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *UpdateScheduleRequest) GetDestinationsOk() (*[]DeliveryDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *UpdateScheduleRequest) SetDestinations(v []DeliveryDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *UpdateScheduleRequest) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### SetDestinationsNil

`func (o *UpdateScheduleRequest) SetDestinationsNil(b bool)`

 SetDestinationsNil sets the value for Destinations to be an explicit nil

### UnsetDestinations
`func (o *UpdateScheduleRequest) UnsetDestinations()`

UnsetDestinations ensures that no value is present for Destinations, not even an explicit nil
### GetDiffThreshold

`func (o *UpdateScheduleRequest) GetDiffThreshold() float64`

GetDiffThreshold returns the DiffThreshold field if non-nil, zero value otherwise.

### GetDiffThresholdOk

`func (o *UpdateScheduleRequest) GetDiffThresholdOk() (*float64, bool)`

GetDiffThresholdOk returns a tuple with the DiffThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffThreshold

`func (o *UpdateScheduleRequest) SetDiffThreshold(v float64)`

SetDiffThreshold sets DiffThreshold field to given value.

### HasDiffThreshold

`func (o *UpdateScheduleRequest) HasDiffThreshold() bool`

HasDiffThreshold returns a boolean if a field has been set.

### SetDiffThresholdNil

`func (o *UpdateScheduleRequest) SetDiffThresholdNil(b bool)`

 SetDiffThresholdNil sets the value for DiffThreshold to be an explicit nil

### UnsetDiffThreshold
`func (o *UpdateScheduleRequest) UnsetDiffThreshold()`

UnsetDiffThreshold ensures that no value is present for DiffThreshold, not even an explicit nil
### GetEndsAt

`func (o *UpdateScheduleRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *UpdateScheduleRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *UpdateScheduleRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *UpdateScheduleRequest) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *UpdateScheduleRequest) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *UpdateScheduleRequest) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetName

`func (o *UpdateScheduleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateScheduleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateScheduleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateScheduleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateScheduleRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateScheduleRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOnlyOnChange

`func (o *UpdateScheduleRequest) GetOnlyOnChange() bool`

GetOnlyOnChange returns the OnlyOnChange field if non-nil, zero value otherwise.

### GetOnlyOnChangeOk

`func (o *UpdateScheduleRequest) GetOnlyOnChangeOk() (*bool, bool)`

GetOnlyOnChangeOk returns a tuple with the OnlyOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyOnChange

`func (o *UpdateScheduleRequest) SetOnlyOnChange(v bool)`

SetOnlyOnChange sets OnlyOnChange field to given value.

### HasOnlyOnChange

`func (o *UpdateScheduleRequest) HasOnlyOnChange() bool`

HasOnlyOnChange returns a boolean if a field has been set.

### SetOnlyOnChangeNil

`func (o *UpdateScheduleRequest) SetOnlyOnChangeNil(b bool)`

 SetOnlyOnChangeNil sets the value for OnlyOnChange to be an explicit nil

### UnsetOnlyOnChange
`func (o *UpdateScheduleRequest) UnsetOnlyOnChange()`

UnsetOnlyOnChange ensures that no value is present for OnlyOnChange, not even an explicit nil
### GetOptions

`func (o *UpdateScheduleRequest) GetOptions() ScheduleScreenshotOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateScheduleRequest) GetOptionsOk() (*ScheduleScreenshotOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateScheduleRequest) SetOptions(v ScheduleScreenshotOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateScheduleRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *UpdateScheduleRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *UpdateScheduleRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetRetentionDays

`func (o *UpdateScheduleRequest) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *UpdateScheduleRequest) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *UpdateScheduleRequest) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *UpdateScheduleRequest) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### SetRetentionDaysNil

`func (o *UpdateScheduleRequest) SetRetentionDaysNil(b bool)`

 SetRetentionDaysNil sets the value for RetentionDays to be an explicit nil

### UnsetRetentionDays
`func (o *UpdateScheduleRequest) UnsetRetentionDays()`

UnsetRetentionDays ensures that no value is present for RetentionDays, not even an explicit nil
### GetSchedule

`func (o *UpdateScheduleRequest) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateScheduleRequest) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateScheduleRequest) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateScheduleRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *UpdateScheduleRequest) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *UpdateScheduleRequest) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetStartsAt

`func (o *UpdateScheduleRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateScheduleRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateScheduleRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateScheduleRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *UpdateScheduleRequest) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *UpdateScheduleRequest) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetTimezone

`func (o *UpdateScheduleRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateScheduleRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateScheduleRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateScheduleRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *UpdateScheduleRequest) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *UpdateScheduleRequest) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetUrl

`func (o *UpdateScheduleRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateScheduleRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateScheduleRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateScheduleRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *UpdateScheduleRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *UpdateScheduleRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetWebhookSecret

`func (o *UpdateScheduleRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *UpdateScheduleRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *UpdateScheduleRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *UpdateScheduleRequest) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *UpdateScheduleRequest) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *UpdateScheduleRequest) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetWebhookUrl

`func (o *UpdateScheduleRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateScheduleRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateScheduleRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateScheduleRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *UpdateScheduleRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *UpdateScheduleRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


