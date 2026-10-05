# BrowserSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cookies** | Pointer to [**[]BrowserCookieRequest**](BrowserCookieRequest.md) |  | [optional] 
**Headers** | Pointer to [**[]BrowserHeaderRequest**](BrowserHeaderRequest.md) |  | [optional] 
**Storage** | Pointer to [**[]BrowserStorageOriginRequest**](BrowserStorageOriginRequest.md) |  | [optional] 

## Methods

### NewBrowserSessionRequest

`func NewBrowserSessionRequest() *BrowserSessionRequest`

NewBrowserSessionRequest instantiates a new BrowserSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserSessionRequestWithDefaults

`func NewBrowserSessionRequestWithDefaults() *BrowserSessionRequest`

NewBrowserSessionRequestWithDefaults instantiates a new BrowserSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCookies

`func (o *BrowserSessionRequest) GetCookies() []BrowserCookieRequest`

GetCookies returns the Cookies field if non-nil, zero value otherwise.

### GetCookiesOk

`func (o *BrowserSessionRequest) GetCookiesOk() (*[]BrowserCookieRequest, bool)`

GetCookiesOk returns a tuple with the Cookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookies

`func (o *BrowserSessionRequest) SetCookies(v []BrowserCookieRequest)`

SetCookies sets Cookies field to given value.

### HasCookies

`func (o *BrowserSessionRequest) HasCookies() bool`

HasCookies returns a boolean if a field has been set.

### GetHeaders

`func (o *BrowserSessionRequest) GetHeaders() []BrowserHeaderRequest`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BrowserSessionRequest) GetHeadersOk() (*[]BrowserHeaderRequest, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BrowserSessionRequest) SetHeaders(v []BrowserHeaderRequest)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BrowserSessionRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetStorage

`func (o *BrowserSessionRequest) GetStorage() []BrowserStorageOriginRequest`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BrowserSessionRequest) GetStorageOk() (*[]BrowserStorageOriginRequest, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BrowserSessionRequest) SetStorage(v []BrowserStorageOriginRequest)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *BrowserSessionRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


