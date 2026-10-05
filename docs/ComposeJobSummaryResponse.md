# ComposeJobSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedCaptures** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**FailedCaptures** | **int32** |  | 
**JobId** | **string** |  | 
**LayoutType** | **string** |  | 
**Progress** | **int32** |  | 
**Status** | **string** |  | 
**TotalCaptures** | **int32** |  | 

## Methods

### NewComposeJobSummaryResponse

`func NewComposeJobSummaryResponse(completedCaptures int32, createdAt time.Time, failedCaptures int32, jobId string, layoutType string, progress int32, status string, totalCaptures int32, ) *ComposeJobSummaryResponse`

NewComposeJobSummaryResponse instantiates a new ComposeJobSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeJobSummaryResponseWithDefaults

`func NewComposeJobSummaryResponseWithDefaults() *ComposeJobSummaryResponse`

NewComposeJobSummaryResponseWithDefaults instantiates a new ComposeJobSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *ComposeJobSummaryResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ComposeJobSummaryResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ComposeJobSummaryResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ComposeJobSummaryResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ComposeJobSummaryResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ComposeJobSummaryResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletedCaptures

`func (o *ComposeJobSummaryResponse) GetCompletedCaptures() int32`

GetCompletedCaptures returns the CompletedCaptures field if non-nil, zero value otherwise.

### GetCompletedCapturesOk

`func (o *ComposeJobSummaryResponse) GetCompletedCapturesOk() (*int32, bool)`

GetCompletedCapturesOk returns a tuple with the CompletedCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCaptures

`func (o *ComposeJobSummaryResponse) SetCompletedCaptures(v int32)`

SetCompletedCaptures sets CompletedCaptures field to given value.


### GetCreatedAt

`func (o *ComposeJobSummaryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComposeJobSummaryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComposeJobSummaryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedCaptures

`func (o *ComposeJobSummaryResponse) GetFailedCaptures() int32`

GetFailedCaptures returns the FailedCaptures field if non-nil, zero value otherwise.

### GetFailedCapturesOk

`func (o *ComposeJobSummaryResponse) GetFailedCapturesOk() (*int32, bool)`

GetFailedCapturesOk returns a tuple with the FailedCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCaptures

`func (o *ComposeJobSummaryResponse) SetFailedCaptures(v int32)`

SetFailedCaptures sets FailedCaptures field to given value.


### GetJobId

`func (o *ComposeJobSummaryResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ComposeJobSummaryResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ComposeJobSummaryResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetLayoutType

`func (o *ComposeJobSummaryResponse) GetLayoutType() string`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *ComposeJobSummaryResponse) GetLayoutTypeOk() (*string, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *ComposeJobSummaryResponse) SetLayoutType(v string)`

SetLayoutType sets LayoutType field to given value.


### GetProgress

`func (o *ComposeJobSummaryResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ComposeJobSummaryResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ComposeJobSummaryResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetStatus

`func (o *ComposeJobSummaryResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComposeJobSummaryResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComposeJobSummaryResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalCaptures

`func (o *ComposeJobSummaryResponse) GetTotalCaptures() int32`

GetTotalCaptures returns the TotalCaptures field if non-nil, zero value otherwise.

### GetTotalCapturesOk

`func (o *ComposeJobSummaryResponse) GetTotalCapturesOk() (*int32, bool)`

GetTotalCapturesOk returns a tuple with the TotalCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCaptures

`func (o *ComposeJobSummaryResponse) SetTotalCaptures(v int32)`

SetTotalCaptures sets TotalCaptures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


