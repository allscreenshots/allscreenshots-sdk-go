# QuotaStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bandwidth** | [**BandwidthQuotaResponse**](BandwidthQuotaResponse.md) |  | 
**PeriodEnds** | Pointer to **NullableString** |  | [optional] 
**Screenshots** | [**QuotaDetailResponse**](QuotaDetailResponse.md) |  | 
**Tier** | **string** |  | 

## Methods

### NewQuotaStatusResponse

`func NewQuotaStatusResponse(bandwidth BandwidthQuotaResponse, screenshots QuotaDetailResponse, tier string, ) *QuotaStatusResponse`

NewQuotaStatusResponse instantiates a new QuotaStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaStatusResponseWithDefaults

`func NewQuotaStatusResponseWithDefaults() *QuotaStatusResponse`

NewQuotaStatusResponseWithDefaults instantiates a new QuotaStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidth

`func (o *QuotaStatusResponse) GetBandwidth() BandwidthQuotaResponse`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *QuotaStatusResponse) GetBandwidthOk() (*BandwidthQuotaResponse, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *QuotaStatusResponse) SetBandwidth(v BandwidthQuotaResponse)`

SetBandwidth sets Bandwidth field to given value.


### GetPeriodEnds

`func (o *QuotaStatusResponse) GetPeriodEnds() string`

GetPeriodEnds returns the PeriodEnds field if non-nil, zero value otherwise.

### GetPeriodEndsOk

`func (o *QuotaStatusResponse) GetPeriodEndsOk() (*string, bool)`

GetPeriodEndsOk returns a tuple with the PeriodEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnds

`func (o *QuotaStatusResponse) SetPeriodEnds(v string)`

SetPeriodEnds sets PeriodEnds field to given value.

### HasPeriodEnds

`func (o *QuotaStatusResponse) HasPeriodEnds() bool`

HasPeriodEnds returns a boolean if a field has been set.

### SetPeriodEndsNil

`func (o *QuotaStatusResponse) SetPeriodEndsNil(b bool)`

 SetPeriodEndsNil sets the value for PeriodEnds to be an explicit nil

### UnsetPeriodEnds
`func (o *QuotaStatusResponse) UnsetPeriodEnds()`

UnsetPeriodEnds ensures that no value is present for PeriodEnds, not even an explicit nil
### GetScreenshots

`func (o *QuotaStatusResponse) GetScreenshots() QuotaDetailResponse`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *QuotaStatusResponse) GetScreenshotsOk() (*QuotaDetailResponse, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *QuotaStatusResponse) SetScreenshots(v QuotaDetailResponse)`

SetScreenshots sets Screenshots field to given value.


### GetTier

`func (o *QuotaStatusResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *QuotaStatusResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *QuotaStatusResponse) SetTier(v string)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


