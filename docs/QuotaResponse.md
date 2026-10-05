# QuotaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | [**BandwidthQuotaResponse**](BandwidthQuotaResponse.md) |  | 
**Screenshots** | [**QuotaDetailResponse**](QuotaDetailResponse.md) |  | 

## Methods

### NewQuotaResponse

`func NewQuotaResponse(bandwidth BandwidthQuotaResponse, screenshots QuotaDetailResponse, ) *QuotaResponse`

NewQuotaResponse instantiates a new QuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaResponseWithDefaults

`func NewQuotaResponseWithDefaults() *QuotaResponse`

NewQuotaResponseWithDefaults instantiates a new QuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *QuotaResponse) GetBandwidth() BandwidthQuotaResponse`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *QuotaResponse) GetBandwidthOk() (*BandwidthQuotaResponse, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *QuotaResponse) SetBandwidth(v BandwidthQuotaResponse)`

SetBandwidth sets Bandwidth field to given value.


### GetScreenshots

`func (o *QuotaResponse) GetScreenshots() QuotaDetailResponse`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *QuotaResponse) GetScreenshotsOk() (*QuotaDetailResponse, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *QuotaResponse) SetScreenshots(v QuotaDetailResponse)`

SetScreenshots sets Screenshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


