# ScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertOnFailure** | **bool** |  | 
**AutoPauseAfterFailures** | Pointer to **NullableInt32** |  | [optional] 
**ConsecutiveFailures** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Destinations** | Pointer to [**[]DeliveryDestination**](DeliveryDestination.md) |  | [optional] 
**DiffThreshold** | Pointer to **NullableFloat64** |  | [optional] 
**EndsAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionCount** | **int32** |  | 
**FailureCount** | **int32** |  | 
**Id** | **string** |  | 
**LastExecutedAt** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**NextExecutionAt** | Pointer to **NullableTime** |  | [optional] 
**OnlyOnChange** | **bool** |  | 
**Options** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**RetentionDays** | **int32** |  | 
**Schedule** | **string** |  | 
**ScheduleDescription** | **string** |  | 
**StartsAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | **string** |  | 
**SuccessCount** | **int32** |  | 
**Timezone** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Url** | **string** |  | 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleResponse

`func NewScheduleResponse(alertOnFailure bool, consecutiveFailures int32, createdAt time.Time, executionCount int32, failureCount int32, id string, name string, onlyOnChange bool, retentionDays int32, schedule string, scheduleDescription string, status string, successCount int32, timezone string, updatedAt time.Time, url string, ) *ScheduleResponse`

NewScheduleResponse instantiates a new ScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleResponseWithDefaults

`func NewScheduleResponseWithDefaults() *ScheduleResponse`

NewScheduleResponseWithDefaults instantiates a new ScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertOnFailure

`func (o *ScheduleResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *ScheduleResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *ScheduleResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.


### GetAutoPauseAfterFailures

`func (o *ScheduleResponse) GetAutoPauseAfterFailures() int32`

GetAutoPauseAfterFailures returns the AutoPauseAfterFailures field if non-nil, zero value otherwise.

### GetAutoPauseAfterFailuresOk

`func (o *ScheduleResponse) GetAutoPauseAfterFailuresOk() (*int32, bool)`

GetAutoPauseAfterFailuresOk returns a tuple with the AutoPauseAfterFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPauseAfterFailures

`func (o *ScheduleResponse) SetAutoPauseAfterFailures(v int32)`

SetAutoPauseAfterFailures sets AutoPauseAfterFailures field to given value.

### HasAutoPauseAfterFailures

`func (o *ScheduleResponse) HasAutoPauseAfterFailures() bool`

HasAutoPauseAfterFailures returns a boolean if a field has been set.

### SetAutoPauseAfterFailuresNil

`func (o *ScheduleResponse) SetAutoPauseAfterFailuresNil(b bool)`

 SetAutoPauseAfterFailuresNil sets the value for AutoPauseAfterFailures to be an explicit nil

### UnsetAutoPauseAfterFailures
`func (o *ScheduleResponse) UnsetAutoPauseAfterFailures()`

UnsetAutoPauseAfterFailures ensures that no value is present for AutoPauseAfterFailures, not even an explicit nil
### GetConsecutiveFailures

`func (o *ScheduleResponse) GetConsecutiveFailures() int32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *ScheduleResponse) GetConsecutiveFailuresOk() (*int32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *ScheduleResponse) SetConsecutiveFailures(v int32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetCreatedAt

`func (o *ScheduleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDestinations

`func (o *ScheduleResponse) GetDestinations() []DeliveryDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ScheduleResponse) GetDestinationsOk() (*[]DeliveryDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ScheduleResponse) SetDestinations(v []DeliveryDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *ScheduleResponse) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### SetDestinationsNil

`func (o *ScheduleResponse) SetDestinationsNil(b bool)`

 SetDestinationsNil sets the value for Destinations to be an explicit nil

### UnsetDestinations
`func (o *ScheduleResponse) UnsetDestinations()`

UnsetDestinations ensures that no value is present for Destinations, not even an explicit nil
### GetDiffThreshold

`func (o *ScheduleResponse) GetDiffThreshold() float64`

GetDiffThreshold returns the DiffThreshold field if non-nil, zero value otherwise.

### GetDiffThresholdOk

`func (o *ScheduleResponse) GetDiffThresholdOk() (*float64, bool)`

GetDiffThresholdOk returns a tuple with the DiffThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffThreshold

`func (o *ScheduleResponse) SetDiffThreshold(v float64)`

SetDiffThreshold sets DiffThreshold field to given value.

### HasDiffThreshold

`func (o *ScheduleResponse) HasDiffThreshold() bool`

HasDiffThreshold returns a boolean if a field has been set.

### SetDiffThresholdNil

`func (o *ScheduleResponse) SetDiffThresholdNil(b bool)`

 SetDiffThresholdNil sets the value for DiffThreshold to be an explicit nil

### UnsetDiffThreshold
`func (o *ScheduleResponse) UnsetDiffThreshold()`

UnsetDiffThreshold ensures that no value is present for DiffThreshold, not even an explicit nil
### GetEndsAt

`func (o *ScheduleResponse) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *ScheduleResponse) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *ScheduleResponse) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *ScheduleResponse) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *ScheduleResponse) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *ScheduleResponse) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetExecutionCount

`func (o *ScheduleResponse) GetExecutionCount() int32`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *ScheduleResponse) GetExecutionCountOk() (*int32, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *ScheduleResponse) SetExecutionCount(v int32)`

SetExecutionCount sets ExecutionCount field to given value.


### GetFailureCount

`func (o *ScheduleResponse) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *ScheduleResponse) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *ScheduleResponse) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.


### GetId

`func (o *ScheduleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastExecutedAt

`func (o *ScheduleResponse) GetLastExecutedAt() time.Time`

GetLastExecutedAt returns the LastExecutedAt field if non-nil, zero value otherwise.

### GetLastExecutedAtOk

`func (o *ScheduleResponse) GetLastExecutedAtOk() (*time.Time, bool)`

GetLastExecutedAtOk returns a tuple with the LastExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutedAt

`func (o *ScheduleResponse) SetLastExecutedAt(v time.Time)`

SetLastExecutedAt sets LastExecutedAt field to given value.

### HasLastExecutedAt

`func (o *ScheduleResponse) HasLastExecutedAt() bool`

HasLastExecutedAt returns a boolean if a field has been set.

### SetLastExecutedAtNil

`func (o *ScheduleResponse) SetLastExecutedAtNil(b bool)`

 SetLastExecutedAtNil sets the value for LastExecutedAt to be an explicit nil

### UnsetLastExecutedAt
`func (o *ScheduleResponse) UnsetLastExecutedAt()`

UnsetLastExecutedAt ensures that no value is present for LastExecutedAt, not even an explicit nil
### GetName

`func (o *ScheduleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecutionAt

`func (o *ScheduleResponse) GetNextExecutionAt() time.Time`

GetNextExecutionAt returns the NextExecutionAt field if non-nil, zero value otherwise.

### GetNextExecutionAtOk

`func (o *ScheduleResponse) GetNextExecutionAtOk() (*time.Time, bool)`

GetNextExecutionAtOk returns a tuple with the NextExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionAt

`func (o *ScheduleResponse) SetNextExecutionAt(v time.Time)`

SetNextExecutionAt sets NextExecutionAt field to given value.

### HasNextExecutionAt

`func (o *ScheduleResponse) HasNextExecutionAt() bool`

HasNextExecutionAt returns a boolean if a field has been set.

### SetNextExecutionAtNil

`func (o *ScheduleResponse) SetNextExecutionAtNil(b bool)`

 SetNextExecutionAtNil sets the value for NextExecutionAt to be an explicit nil

### UnsetNextExecutionAt
`func (o *ScheduleResponse) UnsetNextExecutionAt()`

UnsetNextExecutionAt ensures that no value is present for NextExecutionAt, not even an explicit nil
### GetOnlyOnChange

`func (o *ScheduleResponse) GetOnlyOnChange() bool`

GetOnlyOnChange returns the OnlyOnChange field if non-nil, zero value otherwise.

### GetOnlyOnChangeOk

`func (o *ScheduleResponse) GetOnlyOnChangeOk() (*bool, bool)`

GetOnlyOnChangeOk returns a tuple with the OnlyOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyOnChange

`func (o *ScheduleResponse) SetOnlyOnChange(v bool)`

SetOnlyOnChange sets OnlyOnChange field to given value.


### GetOptions

`func (o *ScheduleResponse) GetOptions() map[string]map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ScheduleResponse) GetOptionsOk() (*map[string]map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ScheduleResponse) SetOptions(v map[string]map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ScheduleResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ScheduleResponse) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ScheduleResponse) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetRetentionDays

`func (o *ScheduleResponse) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *ScheduleResponse) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *ScheduleResponse) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetSchedule

`func (o *ScheduleResponse) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduleResponse) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduleResponse) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetScheduleDescription

`func (o *ScheduleResponse) GetScheduleDescription() string`

GetScheduleDescription returns the ScheduleDescription field if non-nil, zero value otherwise.

### GetScheduleDescriptionOk

`func (o *ScheduleResponse) GetScheduleDescriptionOk() (*string, bool)`

GetScheduleDescriptionOk returns a tuple with the ScheduleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleDescription

`func (o *ScheduleResponse) SetScheduleDescription(v string)`

SetScheduleDescription sets ScheduleDescription field to given value.


### GetStartsAt

`func (o *ScheduleResponse) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ScheduleResponse) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ScheduleResponse) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *ScheduleResponse) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *ScheduleResponse) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *ScheduleResponse) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetStatus

`func (o *ScheduleResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSuccessCount

`func (o *ScheduleResponse) GetSuccessCount() int32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *ScheduleResponse) GetSuccessCountOk() (*int32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *ScheduleResponse) SetSuccessCount(v int32)`

SetSuccessCount sets SuccessCount field to given value.


### GetTimezone

`func (o *ScheduleResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ScheduleResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ScheduleResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetUpdatedAt

`func (o *ScheduleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ScheduleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ScheduleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *ScheduleResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScheduleResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScheduleResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebhookUrl

`func (o *ScheduleResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ScheduleResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ScheduleResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ScheduleResponse) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *ScheduleResponse) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *ScheduleResponse) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


