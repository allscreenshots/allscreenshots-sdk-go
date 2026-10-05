# UsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credits** | [**CreditBalanceResponse**](CreditBalanceResponse.md) |  | 
**CurrentPeriod** | [**PeriodUsageResponse**](PeriodUsageResponse.md) |  | 
**History** | [**[]PeriodUsageResponse**](PeriodUsageResponse.md) |  | 
**Organization** | [**UsageOrganizationResponse**](UsageOrganizationResponse.md) |  | 
**Quota** | [**QuotaResponse**](QuotaResponse.md) |  | 
**Tier** | **string** |  | 
**Totals** | [**TotalsResponse**](TotalsResponse.md) |  | 

## Methods

### NewUsageResponse

`func NewUsageResponse(credits CreditBalanceResponse, currentPeriod PeriodUsageResponse, history []PeriodUsageResponse, organization UsageOrganizationResponse, quota QuotaResponse, tier string, totals TotalsResponse, ) *UsageResponse`

NewUsageResponse instantiates a new UsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageResponseWithDefaults

`func NewUsageResponseWithDefaults() *UsageResponse`

NewUsageResponseWithDefaults instantiates a new UsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredits

`func (o *UsageResponse) GetCredits() CreditBalanceResponse`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *UsageResponse) GetCreditsOk() (*CreditBalanceResponse, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *UsageResponse) SetCredits(v CreditBalanceResponse)`

SetCredits sets Credits field to given value.


### GetCurrentPeriod

`func (o *UsageResponse) GetCurrentPeriod() PeriodUsageResponse`

GetCurrentPeriod returns the CurrentPeriod field if non-nil, zero value otherwise.

### GetCurrentPeriodOk

`func (o *UsageResponse) GetCurrentPeriodOk() (*PeriodUsageResponse, bool)`

GetCurrentPeriodOk returns a tuple with the CurrentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriod

`func (o *UsageResponse) SetCurrentPeriod(v PeriodUsageResponse)`

SetCurrentPeriod sets CurrentPeriod field to given value.


### GetHistory

`func (o *UsageResponse) GetHistory() []PeriodUsageResponse`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *UsageResponse) GetHistoryOk() (*[]PeriodUsageResponse, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *UsageResponse) SetHistory(v []PeriodUsageResponse)`

SetHistory sets History field to given value.


### GetOrganization

`func (o *UsageResponse) GetOrganization() UsageOrganizationResponse`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UsageResponse) GetOrganizationOk() (*UsageOrganizationResponse, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UsageResponse) SetOrganization(v UsageOrganizationResponse)`

SetOrganization sets Organization field to given value.


### GetQuota

`func (o *UsageResponse) GetQuota() QuotaResponse`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *UsageResponse) GetQuotaOk() (*QuotaResponse, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *UsageResponse) SetQuota(v QuotaResponse)`

SetQuota sets Quota field to given value.


### GetTier

`func (o *UsageResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *UsageResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *UsageResponse) SetTier(v string)`

SetTier sets Tier field to given value.


### GetTotals

`func (o *UsageResponse) GetTotals() TotalsResponse`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *UsageResponse) GetTotalsOk() (*TotalsResponse, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *UsageResponse) SetTotals(v TotalsResponse)`

SetTotals sets Totals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


