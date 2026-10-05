# BulkJobInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**ResultUrl** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewBulkJobInfo

`func NewBulkJobInfo(id string, status string, url string, ) *BulkJobInfo`

NewBulkJobInfo instantiates a new BulkJobInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkJobInfoWithDefaults

`func NewBulkJobInfoWithDefaults() *BulkJobInfo`

NewBulkJobInfoWithDefaults instantiates a new BulkJobInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *BulkJobInfo) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *BulkJobInfo) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *BulkJobInfo) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *BulkJobInfo) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *BulkJobInfo) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *BulkJobInfo) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorMessage

`func (o *BulkJobInfo) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *BulkJobInfo) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *BulkJobInfo) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *BulkJobInfo) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *BulkJobInfo) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *BulkJobInfo) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *BulkJobInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkJobInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkJobInfo) SetId(v string)`

SetId sets Id field to given value.


### GetResultUrl

`func (o *BulkJobInfo) GetResultUrl() string`

GetResultUrl returns the ResultUrl field if non-nil, zero value otherwise.

### GetResultUrlOk

`func (o *BulkJobInfo) GetResultUrlOk() (*string, bool)`

GetResultUrlOk returns a tuple with the ResultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultUrl

`func (o *BulkJobInfo) SetResultUrl(v string)`

SetResultUrl sets ResultUrl field to given value.

### HasResultUrl

`func (o *BulkJobInfo) HasResultUrl() bool`

HasResultUrl returns a boolean if a field has been set.

### SetResultUrlNil

`func (o *BulkJobInfo) SetResultUrlNil(b bool)`

 SetResultUrlNil sets the value for ResultUrl to be an explicit nil

### UnsetResultUrl
`func (o *BulkJobInfo) UnsetResultUrl()`

UnsetResultUrl ensures that no value is present for ResultUrl, not even an explicit nil
### GetStatus

`func (o *BulkJobInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkJobInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkJobInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUrl

`func (o *BulkJobInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BulkJobInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BulkJobInfo) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


