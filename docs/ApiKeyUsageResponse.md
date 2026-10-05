# ApiKeyUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**[]ApiKeyUsageDetail**](ApiKeyUsageDetail.md) |  | 
**PeriodEnd** | **string** |  | 
**PeriodStart** | **string** |  | 

## Methods

### NewApiKeyUsageResponse

`func NewApiKeyUsageResponse(keys []ApiKeyUsageDetail, periodEnd string, periodStart string, ) *ApiKeyUsageResponse`

NewApiKeyUsageResponse instantiates a new ApiKeyUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyUsageResponseWithDefaults

`func NewApiKeyUsageResponseWithDefaults() *ApiKeyUsageResponse`

NewApiKeyUsageResponseWithDefaults instantiates a new ApiKeyUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ApiKeyUsageResponse) GetKeys() []ApiKeyUsageDetail`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ApiKeyUsageResponse) GetKeysOk() (*[]ApiKeyUsageDetail, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ApiKeyUsageResponse) SetKeys(v []ApiKeyUsageDetail)`

SetKeys sets Keys field to given value.


### GetPeriodEnd

`func (o *ApiKeyUsageResponse) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *ApiKeyUsageResponse) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *ApiKeyUsageResponse) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *ApiKeyUsageResponse) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *ApiKeyUsageResponse) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *ApiKeyUsageResponse) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


