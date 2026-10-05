# ComposeJobCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**JobId** | **string** |  | 
**Status** | **string** |  | 
**StatusUrl** | **string** |  | 

## Methods

### NewComposeJobCreatedResponse

`func NewComposeJobCreatedResponse(createdAt time.Time, jobId string, status string, statusUrl string, ) *ComposeJobCreatedResponse`

NewComposeJobCreatedResponse instantiates a new ComposeJobCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeJobCreatedResponseWithDefaults

`func NewComposeJobCreatedResponseWithDefaults() *ComposeJobCreatedResponse`

NewComposeJobCreatedResponseWithDefaults instantiates a new ComposeJobCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ComposeJobCreatedResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ComposeJobCreatedResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ComposeJobCreatedResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetJobId

`func (o *ComposeJobCreatedResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ComposeJobCreatedResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ComposeJobCreatedResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetStatus

`func (o *ComposeJobCreatedResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComposeJobCreatedResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComposeJobCreatedResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusUrl

`func (o *ComposeJobCreatedResponse) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *ComposeJobCreatedResponse) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *ComposeJobCreatedResponse) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


