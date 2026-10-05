# ScheduleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedules** | [**[]ScheduleResponse**](ScheduleResponse.md) |  | 
**Total** | **int32** |  | 

## Methods

### NewScheduleListResponse

`func NewScheduleListResponse(schedules []ScheduleResponse, total int32, ) *ScheduleListResponse`

NewScheduleListResponse instantiates a new ScheduleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleListResponseWithDefaults

`func NewScheduleListResponseWithDefaults() *ScheduleListResponse`

NewScheduleListResponseWithDefaults instantiates a new ScheduleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedules

`func (o *ScheduleListResponse) GetSchedules() []ScheduleResponse`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ScheduleListResponse) GetSchedulesOk() (*[]ScheduleResponse, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ScheduleListResponse) SetSchedules(v []ScheduleResponse)`

SetSchedules sets Schedules field to given value.


### GetTotal

`func (o *ScheduleListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ScheduleListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ScheduleListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


