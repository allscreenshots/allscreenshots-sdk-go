# BulkJobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedJobs** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**FailedJobs** | **int32** |  | 
**Id** | **string** |  | 
**Progress** | **int32** |  | 
**Status** | **string** |  | 
**TotalJobs** | **int32** |  | 

## Methods

### NewBulkJobSummary

`func NewBulkJobSummary(completedJobs int32, createdAt time.Time, failedJobs int32, id string, progress int32, status string, totalJobs int32, ) *BulkJobSummary`

NewBulkJobSummary instantiates a new BulkJobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobSummaryWithDefaults

`func NewBulkJobSummaryWithDefaults() *BulkJobSummary`

NewBulkJobSummaryWithDefaults instantiates a new BulkJobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BulkJobSummary) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BulkJobSummary) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BulkJobSummary) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BulkJobSummary) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BulkJobSummary) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BulkJobSummary) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletedJobs

`func (o *BulkJobSummary) GetCompletedJobs() int32`

GetCompletedJobs returns the CompletedJobs field if non-nil, zero value otherwise.

### GetCompletedJobsOk

`func (o *BulkJobSummary) GetCompletedJobsOk() (*int32, bool)`

GetCompletedJobsOk returns a tuple with the CompletedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedJobs

`func (o *BulkJobSummary) SetCompletedJobs(v int32)`

SetCompletedJobs sets CompletedJobs field to given value.


### GetCreatedAt

`func (o *BulkJobSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkJobSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkJobSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedJobs

`func (o *BulkJobSummary) GetFailedJobs() int32`

GetFailedJobs returns the FailedJobs field if non-nil, zero value otherwise.

### GetFailedJobsOk

`func (o *BulkJobSummary) GetFailedJobsOk() (*int32, bool)`

GetFailedJobsOk returns a tuple with the FailedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobs

`func (o *BulkJobSummary) SetFailedJobs(v int32)`

SetFailedJobs sets FailedJobs field to given value.


### GetId

`func (o *BulkJobSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobSummary) SetId(v string)`

SetId sets Id field to given value.


### GetProgress

`func (o *BulkJobSummary) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BulkJobSummary) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BulkJobSummary) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetStatus

`func (o *BulkJobSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkJobSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkJobSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalJobs

`func (o *BulkJobSummary) GetTotalJobs() int32`

GetTotalJobs returns the TotalJobs field if non-nil, zero value otherwise.

### GetTotalJobsOk

`func (o *BulkJobSummary) GetTotalJobsOk() (*int32, bool)`

GetTotalJobsOk returns a tuple with the TotalJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalJobs

`func (o *BulkJobSummary) SetTotalJobs(v int32)`

SetTotalJobs sets TotalJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


