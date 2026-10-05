# BulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedJobs** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**FailedJobs** | **int32** |  | 
**Id** | **string** |  | 
**Jobs** | [**[]BulkJobInfo**](BulkJobInfo.md) |  | 
**Progress** | **int32** |  | 
**Status** | **string** |  | 
**TotalJobs** | **int32** |  | 

## Methods

### NewBulkResponse

`func NewBulkResponse(completedJobs int32, createdAt time.Time, failedJobs int32, id string, jobs []BulkJobInfo, progress int32, status string, totalJobs int32, ) *BulkResponse`

NewBulkResponse instantiates a new BulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResponseWithDefaults

`func NewBulkResponseWithDefaults() *BulkResponse`

NewBulkResponseWithDefaults instantiates a new BulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedAt

`func (o *BulkResponse) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BulkResponse) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BulkResponse) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BulkResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BulkResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BulkResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCompletedJobs

`func (o *BulkResponse) GetCompletedJobs() int32`

GetCompletedJobs returns the CompletedJobs field if non-nil, zero value otherwise.

### GetCompletedJobsOk

`func (o *BulkResponse) GetCompletedJobsOk() (*int32, bool)`

GetCompletedJobsOk returns a tuple with the CompletedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedJobs

`func (o *BulkResponse) SetCompletedJobs(v int32)`

SetCompletedJobs sets CompletedJobs field to given value.


### GetCreatedAt

`func (o *BulkResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BulkResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BulkResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedJobs

`func (o *BulkResponse) GetFailedJobs() int32`

GetFailedJobs returns the FailedJobs field if non-nil, zero value otherwise.

### GetFailedJobsOk

`func (o *BulkResponse) GetFailedJobsOk() (*int32, bool)`

GetFailedJobsOk returns a tuple with the FailedJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedJobs

`func (o *BulkResponse) SetFailedJobs(v int32)`

SetFailedJobs sets FailedJobs field to given value.


### GetId

`func (o *BulkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJobs

`func (o *BulkResponse) GetJobs() []BulkJobInfo`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *BulkResponse) GetJobsOk() (*[]BulkJobInfo, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *BulkResponse) SetJobs(v []BulkJobInfo)`

SetJobs sets Jobs field to given value.


### GetProgress

`func (o *BulkResponse) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BulkResponse) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BulkResponse) SetProgress(v int32)`

SetProgress sets Progress field to given value.


### GetStatus

`func (o *BulkResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalJobs

`func (o *BulkResponse) GetTotalJobs() int32`

GetTotalJobs returns the TotalJobs field if non-nil, zero value otherwise.

### GetTotalJobsOk

`func (o *BulkResponse) GetTotalJobsOk() (*int32, bool)`

GetTotalJobsOk returns a tuple with the TotalJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalJobs

`func (o *BulkResponse) SetTotalJobs(v int32)`

SetTotalJobs sets TotalJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


