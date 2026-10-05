# ScheduleHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | [**[]ScheduleExecutionResponse**](ScheduleExecutionResponse.md) |  | 
**ScheduleId** | **string** |  | 
**TotalExecutions** | **int64** |  | 

## Methods

### NewScheduleHistoryResponse

`func NewScheduleHistoryResponse(executions []ScheduleExecutionResponse, scheduleId string, totalExecutions int64, ) *ScheduleHistoryResponse`

NewScheduleHistoryResponse instantiates a new ScheduleHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleHistoryResponseWithDefaults

`func NewScheduleHistoryResponseWithDefaults() *ScheduleHistoryResponse`

NewScheduleHistoryResponseWithDefaults instantiates a new ScheduleHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *ScheduleHistoryResponse) GetExecutions() []ScheduleExecutionResponse`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *ScheduleHistoryResponse) GetExecutionsOk() (*[]ScheduleExecutionResponse, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *ScheduleHistoryResponse) SetExecutions(v []ScheduleExecutionResponse)`

SetExecutions sets Executions field to given value.


### GetScheduleId

`func (o *ScheduleHistoryResponse) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ScheduleHistoryResponse) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ScheduleHistoryResponse) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.


### GetTotalExecutions

`func (o *ScheduleHistoryResponse) GetTotalExecutions() int64`

GetTotalExecutions returns the TotalExecutions field if non-nil, zero value otherwise.

### GetTotalExecutionsOk

`func (o *ScheduleHistoryResponse) GetTotalExecutionsOk() (*int64, bool)`

GetTotalExecutionsOk returns a tuple with the TotalExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExecutions

`func (o *ScheduleHistoryResponse) SetTotalExecutions(v int64)`

SetTotalExecutions sets TotalExecutions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


