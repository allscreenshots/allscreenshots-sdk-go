# BrowserStorageOriginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStorage** | Pointer to **map[string]string** |  | [optional] 
**Origin** | **string** |  | 
**SessionStorage** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBrowserStorageOriginRequest

`func NewBrowserStorageOriginRequest(origin string, ) *BrowserStorageOriginRequest`

NewBrowserStorageOriginRequest instantiates a new BrowserStorageOriginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserStorageOriginRequestWithDefaults

`func NewBrowserStorageOriginRequestWithDefaults() *BrowserStorageOriginRequest`

NewBrowserStorageOriginRequestWithDefaults instantiates a new BrowserStorageOriginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStorage

`func (o *BrowserStorageOriginRequest) GetLocalStorage() map[string]string`

GetLocalStorage returns the LocalStorage field if non-nil, zero value otherwise.

### GetLocalStorageOk

`func (o *BrowserStorageOriginRequest) GetLocalStorageOk() (*map[string]string, bool)`

GetLocalStorageOk returns a tuple with the LocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorage

`func (o *BrowserStorageOriginRequest) SetLocalStorage(v map[string]string)`

SetLocalStorage sets LocalStorage field to given value.

### HasLocalStorage

`func (o *BrowserStorageOriginRequest) HasLocalStorage() bool`

HasLocalStorage returns a boolean if a field has been set.

### GetOrigin

`func (o *BrowserStorageOriginRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BrowserStorageOriginRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BrowserStorageOriginRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetSessionStorage

`func (o *BrowserStorageOriginRequest) GetSessionStorage() map[string]string`

GetSessionStorage returns the SessionStorage field if non-nil, zero value otherwise.

### GetSessionStorageOk

`func (o *BrowserStorageOriginRequest) GetSessionStorageOk() (*map[string]string, bool)`

GetSessionStorageOk returns a tuple with the SessionStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionStorage

`func (o *BrowserStorageOriginRequest) SetSessionStorage(v map[string]string)`

SetSessionStorage sets SessionStorage field to given value.

### HasSessionStorage

`func (o *BrowserStorageOriginRequest) HasSessionStorage() bool`

HasSessionStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


