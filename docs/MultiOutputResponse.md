# MultiOutputResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outputs** | **map[string]map[string]interface{}** |  | 
**RenderTimeMs** | **int64** |  | 
**Url** | **string** |  | 

## Methods

### NewMultiOutputResponse

`func NewMultiOutputResponse(outputs map[string]map[string]interface{}, renderTimeMs int64, url string, ) *MultiOutputResponse`

NewMultiOutputResponse instantiates a new MultiOutputResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiOutputResponseWithDefaults

`func NewMultiOutputResponseWithDefaults() *MultiOutputResponse`

NewMultiOutputResponseWithDefaults instantiates a new MultiOutputResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutputs

`func (o *MultiOutputResponse) GetOutputs() map[string]map[string]interface{}`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *MultiOutputResponse) GetOutputsOk() (*map[string]map[string]interface{}, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *MultiOutputResponse) SetOutputs(v map[string]map[string]interface{})`

SetOutputs sets Outputs field to given value.


### GetRenderTimeMs

`func (o *MultiOutputResponse) GetRenderTimeMs() int64`

GetRenderTimeMs returns the RenderTimeMs field if non-nil, zero value otherwise.

### GetRenderTimeMsOk

`func (o *MultiOutputResponse) GetRenderTimeMsOk() (*int64, bool)`

GetRenderTimeMsOk returns a tuple with the RenderTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderTimeMs

`func (o *MultiOutputResponse) SetRenderTimeMs(v int64)`

SetRenderTimeMs sets RenderTimeMs field to given value.


### GetUrl

`func (o *MultiOutputResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MultiOutputResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MultiOutputResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


