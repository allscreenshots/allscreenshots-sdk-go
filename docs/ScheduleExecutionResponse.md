# ScheduleExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changed** | Pointer to **NullableBool** |  | [optional] 
**Deliveries** | Pointer to **[]map[string]map[string]interface{}** |  | [optional] 
**DiffScore** | Pointer to **NullableFloat64** |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ExecutedAt** | **time.Time** |  | 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**FileSize** | Pointer to **NullableInt64** |  | [optional] 
**Id** | **string** |  | 
**Outputs** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**RenderTimeMs** | Pointer to **NullableInt64** |  | [optional] 
**ResultUrl** | Pointer to **NullableString** |  | [optional] 
**RunId** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**StorageUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleExecutionResponse

`func NewScheduleExecutionResponse(executedAt time.Time, id string, status string, ) *ScheduleExecutionResponse`

NewScheduleExecutionResponse instantiates a new ScheduleExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleExecutionResponseWithDefaults

`func NewScheduleExecutionResponseWithDefaults() *ScheduleExecutionResponse`

NewScheduleExecutionResponseWithDefaults instantiates a new ScheduleExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChanged

`func (o *ScheduleExecutionResponse) GetChanged() bool`

GetChanged returns the Changed field if non-nil, zero value otherwise.

### GetChangedOk

`func (o *ScheduleExecutionResponse) GetChangedOk() (*bool, bool)`

GetChangedOk returns a tuple with the Changed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanged

`func (o *ScheduleExecutionResponse) SetChanged(v bool)`

SetChanged sets Changed field to given value.

### HasChanged

`func (o *ScheduleExecutionResponse) HasChanged() bool`

HasChanged returns a boolean if a field has been set.

### SetChangedNil

`func (o *ScheduleExecutionResponse) SetChangedNil(b bool)`

 SetChangedNil sets the value for Changed to be an explicit nil

### UnsetChanged
`func (o *ScheduleExecutionResponse) UnsetChanged()`

UnsetChanged ensures that no value is present for Changed, not even an explicit nil
### GetDeliveries

`func (o *ScheduleExecutionResponse) GetDeliveries() []map[string]map[string]interface{}`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *ScheduleExecutionResponse) GetDeliveriesOk() (*[]map[string]map[string]interface{}, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *ScheduleExecutionResponse) SetDeliveries(v []map[string]map[string]interface{})`

SetDeliveries sets Deliveries field to given value.

### HasDeliveries

`func (o *ScheduleExecutionResponse) HasDeliveries() bool`

HasDeliveries returns a boolean if a field has been set.

### SetDeliveriesNil

`func (o *ScheduleExecutionResponse) SetDeliveriesNil(b bool)`

 SetDeliveriesNil sets the value for Deliveries to be an explicit nil

### UnsetDeliveries
`func (o *ScheduleExecutionResponse) UnsetDeliveries()`

UnsetDeliveries ensures that no value is present for Deliveries, not even an explicit nil
### GetDiffScore

`func (o *ScheduleExecutionResponse) GetDiffScore() float64`

GetDiffScore returns the DiffScore field if non-nil, zero value otherwise.

### GetDiffScoreOk

`func (o *ScheduleExecutionResponse) GetDiffScoreOk() (*float64, bool)`

GetDiffScoreOk returns a tuple with the DiffScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffScore

`func (o *ScheduleExecutionResponse) SetDiffScore(v float64)`

SetDiffScore sets DiffScore field to given value.

### HasDiffScore

`func (o *ScheduleExecutionResponse) HasDiffScore() bool`

HasDiffScore returns a boolean if a field has been set.

### SetDiffScoreNil

`func (o *ScheduleExecutionResponse) SetDiffScoreNil(b bool)`

 SetDiffScoreNil sets the value for DiffScore to be an explicit nil

### UnsetDiffScore
`func (o *ScheduleExecutionResponse) UnsetDiffScore()`

UnsetDiffScore ensures that no value is present for DiffScore, not even an explicit nil
### GetErrorCode

`func (o *ScheduleExecutionResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ScheduleExecutionResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ScheduleExecutionResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ScheduleExecutionResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ScheduleExecutionResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ScheduleExecutionResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *ScheduleExecutionResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ScheduleExecutionResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ScheduleExecutionResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ScheduleExecutionResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ScheduleExecutionResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ScheduleExecutionResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetExecutedAt

`func (o *ScheduleExecutionResponse) GetExecutedAt() time.Time`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *ScheduleExecutionResponse) GetExecutedAtOk() (*time.Time, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *ScheduleExecutionResponse) SetExecutedAt(v time.Time)`

SetExecutedAt sets ExecutedAt field to given value.


### GetExpiresAt

`func (o *ScheduleExecutionResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ScheduleExecutionResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ScheduleExecutionResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ScheduleExecutionResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *ScheduleExecutionResponse) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *ScheduleExecutionResponse) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetFileSize

`func (o *ScheduleExecutionResponse) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *ScheduleExecutionResponse) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *ScheduleExecutionResponse) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *ScheduleExecutionResponse) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### SetFileSizeNil

`func (o *ScheduleExecutionResponse) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *ScheduleExecutionResponse) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetId

`func (o *ScheduleExecutionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleExecutionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleExecutionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetOutputs

`func (o *ScheduleExecutionResponse) GetOutputs() map[string]map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *ScheduleExecutionResponse) GetOutputsOk() (*map[string]map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *ScheduleExecutionResponse) SetOutputs(v map[string]map[string]interface{})`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *ScheduleExecutionResponse) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *ScheduleExecutionResponse) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *ScheduleExecutionResponse) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetRenderTimeMs

`func (o *ScheduleExecutionResponse) GetRenderTimeMs() int64`

GetRenderTimeMs returns the RenderTimeMs field if non-nil, zero value otherwise.

### GetRenderTimeMsOk

`func (o *ScheduleExecutionResponse) GetRenderTimeMsOk() (*int64, bool)`

GetRenderTimeMsOk returns a tuple with the RenderTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTimeMs

`func (o *ScheduleExecutionResponse) SetRenderTimeMs(v int64)`

SetRenderTimeMs sets RenderTimeMs field to given value.

### HasRenderTimeMs

`func (o *ScheduleExecutionResponse) HasRenderTimeMs() bool`

HasRenderTimeMs returns a boolean if a field has been set.

### SetRenderTimeMsNil

`func (o *ScheduleExecutionResponse) SetRenderTimeMsNil(b bool)`

 SetRenderTimeMsNil sets the value for RenderTimeMs to be an explicit nil

### UnsetRenderTimeMs
`func (o *ScheduleExecutionResponse) UnsetRenderTimeMs()`

UnsetRenderTimeMs ensures that no value is present for RenderTimeMs, not even an explicit nil
### GetResultUrl

`func (o *ScheduleExecutionResponse) GetResultUrl() string`

GetResultUrl returns the ResultUrl field if non-nil, zero value otherwise.

### GetResultUrlOk

`func (o *ScheduleExecutionResponse) GetResultUrlOk() (*string, bool)`

GetResultUrlOk returns a tuple with the ResultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrl

`func (o *ScheduleExecutionResponse) SetResultUrl(v string)`

SetResultUrl sets ResultUrl field to given value.

### HasResultUrl

`func (o *ScheduleExecutionResponse) HasResultUrl() bool`

HasResultUrl returns a boolean if a field has been set.

### SetResultUrlNil

`func (o *ScheduleExecutionResponse) SetResultUrlNil(b bool)`

 SetResultUrlNil sets the value for ResultUrl to be an explicit nil

### UnsetResultUrl
`func (o *ScheduleExecutionResponse) UnsetResultUrl()`

UnsetResultUrl ensures that no value is present for ResultUrl, not even an explicit nil
### GetRunId

`func (o *ScheduleExecutionResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ScheduleExecutionResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ScheduleExecutionResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ScheduleExecutionResponse) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *ScheduleExecutionResponse) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *ScheduleExecutionResponse) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetStatus

`func (o *ScheduleExecutionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleExecutionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleExecutionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStorageUrl

`func (o *ScheduleExecutionResponse) GetStorageUrl() string`

GetStorageUrl returns the StorageUrl field if non-nil, zero value otherwise.

### GetStorageUrlOk

`func (o *ScheduleExecutionResponse) GetStorageUrlOk() (*string, bool)`

GetStorageUrlOk returns a tuple with the StorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUrl

`func (o *ScheduleExecutionResponse) SetStorageUrl(v string)`

SetStorageUrl sets StorageUrl field to given value.

### HasStorageUrl

`func (o *ScheduleExecutionResponse) HasStorageUrl() bool`

HasStorageUrl returns a boolean if a field has been set.

### SetStorageUrlNil

`func (o *ScheduleExecutionResponse) SetStorageUrlNil(b bool)`

 SetStorageUrlNil sets the value for StorageUrl to be an explicit nil

### UnsetStorageUrl
`func (o *ScheduleExecutionResponse) UnsetStorageUrl()`

UnsetStorageUrl ensures that no value is present for StorageUrl, not even an explicit nil
### GetUrl

`func (o *ScheduleExecutionResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ScheduleExecutionResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ScheduleExecutionResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ScheduleExecutionResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ScheduleExecutionResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ScheduleExecutionResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


