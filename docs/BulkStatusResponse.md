# BulkStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedJobs** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**FailedJobs** | **int32** |  | 
**Id** | **string** |  | 
**Jobs** | [**[]BulkJobDetailInfo**](BulkJobDetailInfo.md) |  | 
**Progress** | **int32** |  | 
**Status** | **string** |  | 
**TotalJobs** | **int32** |  | 

## Methods

### NewBulkStatusResponse

`func NewBulkStatusResponse(completedJobs int32, createdAt time.Time, failedJobs int32, id string, jobs []BulkJobDetailInfo, progress int32, status string, totalJobs int32, ) *BulkStatusResponse`

NewBulkStatusResponse instantiates a new BulkStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkStatusResponseWithDefaults

`func NewBulkStatusResponseWithDefaults() *BulkStatusResponse`

NewBulkStatusResponseWithDefaults instantiates a new BulkStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BulkStatusResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BulkStatusResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BulkStatusResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BulkStatusResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BulkStatusResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BulkStatusResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletedJobs

`func (o *BulkStatusResponse) GetCompletedJobs() int32`

GetCompletedJobs returns the CompletedJobs field if non-nil, zero value otherwise.

### GetCompletedJobsOk

`func (o *BulkStatusResponse) GetCompletedJobsOk() (*int32, bool)`

GetCompletedJobsOk returns a tuple with the CompletedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedJobs

`func (o *BulkStatusResponse) SetCompletedJobs(v int32)`

SetCompletedJobs sets CompletedJobs field to given value.


### GetCreatedAt

`func (o *BulkStatusResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkStatusResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkStatusResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedJobs

`func (o *BulkStatusResponse) GetFailedJobs() int32`

GetFailedJobs returns the FailedJobs field if non-nil, zero value otherwise.

### GetFailedJobsOk

`func (o *BulkStatusResponse) GetFailedJobsOk() (*int32, bool)`

GetFailedJobsOk returns a tuple with the FailedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobs

`func (o *BulkStatusResponse) SetFailedJobs(v int32)`

SetFailedJobs sets FailedJobs field to given value.


### GetId

`func (o *BulkStatusResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkStatusResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkStatusResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJobs

`func (o *BulkStatusResponse) GetJobs() []BulkJobDetailInfo`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BulkStatusResponse) GetJobsOk() (*[]BulkJobDetailInfo, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BulkStatusResponse) SetJobs(v []BulkJobDetailInfo)`

SetJobs sets Jobs field to given value.


### GetProgress

`func (o *BulkStatusResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BulkStatusResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BulkStatusResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetStatus

`func (o *BulkStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalJobs

`func (o *BulkStatusResponse) GetTotalJobs() int32`

GetTotalJobs returns the TotalJobs field if non-nil, zero value otherwise.

### GetTotalJobsOk

`func (o *BulkStatusResponse) GetTotalJobsOk() (*int32, bool)`

GetTotalJobsOk returns a tuple with the TotalJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalJobs

`func (o *BulkStatusResponse) SetTotalJobs(v int32)`

SetTotalJobs sets TotalJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


