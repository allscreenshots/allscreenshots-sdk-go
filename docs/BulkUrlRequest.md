# BulkUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**NullableBulkUrlOptions**](BulkUrlOptions.md) |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewBulkUrlRequest

`func NewBulkUrlRequest(url string, ) *BulkUrlRequest`

NewBulkUrlRequest instantiates a new BulkUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUrlRequestWithDefaults

`func NewBulkUrlRequestWithDefaults() *BulkUrlRequest`

NewBulkUrlRequestWithDefaults instantiates a new BulkUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *BulkUrlRequest) GetOptions() BulkUrlOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BulkUrlRequest) GetOptionsOk() (*BulkUrlOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BulkUrlRequest) SetOptions(v BulkUrlOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *BulkUrlRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *BulkUrlRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *BulkUrlRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetUrl

`func (o *BulkUrlRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BulkUrlRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BulkUrlRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


