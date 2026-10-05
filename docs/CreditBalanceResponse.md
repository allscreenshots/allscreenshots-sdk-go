# CreditBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowanceType** | **string** |  | 
**Balance** | **int32** |  | 
**ResetsAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreditBalanceResponse

`func NewCreditBalanceResponse(allowanceType string, balance int32, ) *CreditBalanceResponse`

NewCreditBalanceResponse instantiates a new CreditBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditBalanceResponseWithDefaults

`func NewCreditBalanceResponseWithDefaults() *CreditBalanceResponse`

NewCreditBalanceResponseWithDefaults instantiates a new CreditBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowanceType

`func (o *CreditBalanceResponse) GetAllowanceType() string`

GetAllowanceType returns the AllowanceType field if non-nil, zero value otherwise.

### GetAllowanceTypeOk

`func (o *CreditBalanceResponse) GetAllowanceTypeOk() (*string, bool)`

GetAllowanceTypeOk returns a tuple with the AllowanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowanceType

`func (o *CreditBalanceResponse) SetAllowanceType(v string)`

SetAllowanceType sets AllowanceType field to given value.


### GetBalance

`func (o *CreditBalanceResponse) GetBalance() int32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CreditBalanceResponse) GetBalanceOk() (*int32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CreditBalanceResponse) SetBalance(v int32)`

SetBalance sets Balance field to given value.


### GetResetsAt

`func (o *CreditBalanceResponse) GetResetsAt() string`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *CreditBalanceResponse) GetResetsAtOk() (*string, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *CreditBalanceResponse) SetResetsAt(v string)`

SetResetsAt sets ResetsAt field to given value.

### HasResetsAt

`func (o *CreditBalanceResponse) HasResetsAt() bool`

HasResetsAt returns a boolean if a field has been set.

### SetResetsAtNil

`func (o *CreditBalanceResponse) SetResetsAtNil(b bool)`

 SetResetsAtNil sets the value for ResetsAt to be an explicit nil

### UnsetResetsAt
`func (o *CreditBalanceResponse) UnsetResetsAt()`

UnsetResetsAt ensures that no value is present for ResetsAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


