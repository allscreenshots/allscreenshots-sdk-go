# ApiKeyUsageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | Pointer to **NullableString** |  | [optional] 
**CurrentPeriod** | [**CaptureUsageCounts**](CaptureUsageCounts.md) |  | 
**KeyPrefix** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Totals** | [**CaptureUsageCounts**](CaptureUsageCounts.md) |  | 

## Methods

### NewApiKeyUsageDetail

`func NewApiKeyUsageDetail(currentPeriod CaptureUsageCounts, name string, totals CaptureUsageCounts, ) *ApiKeyUsageDetail`

NewApiKeyUsageDetail instantiates a new ApiKeyUsageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyUsageDetailWithDefaults

`func NewApiKeyUsageDetailWithDefaults() *ApiKeyUsageDetail`

NewApiKeyUsageDetailWithDefaults instantiates a new ApiKeyUsageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *ApiKeyUsageDetail) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *ApiKeyUsageDetail) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *ApiKeyUsageDetail) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.

### HasApiKeyId

`func (o *ApiKeyUsageDetail) HasApiKeyId() bool`

HasApiKeyId returns a boolean if a field has been set.

### SetApiKeyIdNil

`func (o *ApiKeyUsageDetail) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *ApiKeyUsageDetail) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetCurrentPeriod

`func (o *ApiKeyUsageDetail) GetCurrentPeriod() CaptureUsageCounts`

GetCurrentPeriod returns the CurrentPeriod field if non-nil, zero value otherwise.

### GetCurrentPeriodOk

`func (o *ApiKeyUsageDetail) GetCurrentPeriodOk() (*CaptureUsageCounts, bool)`

GetCurrentPeriodOk returns a tuple with the CurrentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriod

`func (o *ApiKeyUsageDetail) SetCurrentPeriod(v CaptureUsageCounts)`

SetCurrentPeriod sets CurrentPeriod field to given value.


### GetKeyPrefix

`func (o *ApiKeyUsageDetail) GetKeyPrefix() string`

GetKeyPrefix returns the KeyPrefix field if non-nil, zero value otherwise.

### GetKeyPrefixOk

`func (o *ApiKeyUsageDetail) GetKeyPrefixOk() (*string, bool)`

GetKeyPrefixOk returns a tuple with the KeyPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPrefix

`func (o *ApiKeyUsageDetail) SetKeyPrefix(v string)`

SetKeyPrefix sets KeyPrefix field to given value.

### HasKeyPrefix

`func (o *ApiKeyUsageDetail) HasKeyPrefix() bool`

HasKeyPrefix returns a boolean if a field has been set.

### SetKeyPrefixNil

`func (o *ApiKeyUsageDetail) SetKeyPrefixNil(b bool)`

 SetKeyPrefixNil sets the value for KeyPrefix to be an explicit nil

### UnsetKeyPrefix
`func (o *ApiKeyUsageDetail) UnsetKeyPrefix()`

UnsetKeyPrefix ensures that no value is present for KeyPrefix, not even an explicit nil
### GetName

`func (o *ApiKeyUsageDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKeyUsageDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKeyUsageDetail) SetName(v string)`

SetName sets Name field to given value.


### GetTotals

`func (o *ApiKeyUsageDetail) GetTotals() CaptureUsageCounts`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *ApiKeyUsageDetail) GetTotalsOk() (*CaptureUsageCounts, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *ApiKeyUsageDetail) SetTotals(v CaptureUsageCounts)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


