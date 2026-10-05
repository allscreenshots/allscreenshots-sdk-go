# AsyncJobCreatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Status** | **string** |  | 
**StatusUrl** | **string** |  | 

## Methods

### NewAsyncJobCreatedResponse

`func NewAsyncJobCreatedResponse(createdAt time.Time, id string, status string, statusUrl string, ) *AsyncJobCreatedResponse`

NewAsyncJobCreatedResponse instantiates a new AsyncJobCreatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncJobCreatedResponseWithDefaults

`func NewAsyncJobCreatedResponseWithDefaults() *AsyncJobCreatedResponse`

NewAsyncJobCreatedResponseWithDefaults instantiates a new AsyncJobCreatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AsyncJobCreatedResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AsyncJobCreatedResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AsyncJobCreatedResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *AsyncJobCreatedResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AsyncJobCreatedResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AsyncJobCreatedResponse) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *AsyncJobCreatedResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AsyncJobCreatedResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AsyncJobCreatedResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusUrl

`func (o *AsyncJobCreatedResponse) GetStatusUrl() string`

GetStatusUrl returns the StatusUrl field if non-nil, zero value otherwise.

### GetStatusUrlOk

`func (o *AsyncJobCreatedResponse) GetStatusUrlOk() (*string, bool)`

GetStatusUrlOk returns a tuple with the StatusUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUrl

`func (o *AsyncJobCreatedResponse) SetStatusUrl(v string)`

SetStatusUrl sets StatusUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


