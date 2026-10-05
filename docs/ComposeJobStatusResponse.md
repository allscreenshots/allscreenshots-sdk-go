# ComposeJobStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedCaptures** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**JobId** | **string** |  | 
**Progress** | **int32** |  | 
**Result** | Pointer to [**NullableComposeResponse**](ComposeResponse.md) |  | [optional] 
**Status** | **string** |  | 
**TotalCaptures** | **int32** |  | 

## Methods

### NewComposeJobStatusResponse

`func NewComposeJobStatusResponse(completedCaptures int32, createdAt time.Time, jobId string, progress int32, status string, totalCaptures int32, ) *ComposeJobStatusResponse`

NewComposeJobStatusResponse instantiates a new ComposeJobStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeJobStatusResponseWithDefaults

`func NewComposeJobStatusResponseWithDefaults() *ComposeJobStatusResponse`

NewComposeJobStatusResponseWithDefaults instantiates a new ComposeJobStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *ComposeJobStatusResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ComposeJobStatusResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ComposeJobStatusResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ComposeJobStatusResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ComposeJobStatusResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ComposeJobStatusResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletedCaptures

`func (o *ComposeJobStatusResponse) GetCompletedCaptures() int32`

GetCompletedCaptures returns the CompletedCaptures field if non-nil, zero value otherwise.

### GetCompletedCapturesOk

`func (o *ComposeJobStatusResponse) GetCompletedCapturesOk() (*int32, bool)`

GetCompletedCapturesOk returns a tuple with the CompletedCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCaptures

`func (o *ComposeJobStatusResponse) SetCompletedCaptures(v int32)`

SetCompletedCaptures sets CompletedCaptures field to given value.


### GetCreatedAt

`func (o *ComposeJobStatusResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComposeJobStatusResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComposeJobStatusResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorCode

`func (o *ComposeJobStatusResponse) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ComposeJobStatusResponse) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ComposeJobStatusResponse) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ComposeJobStatusResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ComposeJobStatusResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ComposeJobStatusResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *ComposeJobStatusResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ComposeJobStatusResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ComposeJobStatusResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ComposeJobStatusResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ComposeJobStatusResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ComposeJobStatusResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetJobId

`func (o *ComposeJobStatusResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ComposeJobStatusResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ComposeJobStatusResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetProgress

`func (o *ComposeJobStatusResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ComposeJobStatusResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ComposeJobStatusResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetResult

`func (o *ComposeJobStatusResponse) GetResult() ComposeResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ComposeJobStatusResponse) GetResultOk() (*ComposeResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ComposeJobStatusResponse) SetResult(v ComposeResponse)`

SetResult sets Result field to given value.

### HasResult

`func (o *ComposeJobStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *ComposeJobStatusResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ComposeJobStatusResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetStatus

`func (o *ComposeJobStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComposeJobStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComposeJobStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalCaptures

`func (o *ComposeJobStatusResponse) GetTotalCaptures() int32`

GetTotalCaptures returns the TotalCaptures field if non-nil, zero value otherwise.

### GetTotalCapturesOk

`func (o *ComposeJobStatusResponse) GetTotalCapturesOk() (*int32, bool)`

GetTotalCapturesOk returns a tuple with the TotalCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCaptures

`func (o *ComposeJobStatusResponse) SetTotalCaptures(v int32)`

SetTotalCaptures sets TotalCaptures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


