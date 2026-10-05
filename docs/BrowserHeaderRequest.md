# BrowserHeaderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Origin** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewBrowserHeaderRequest

`func NewBrowserHeaderRequest(name string, ) *BrowserHeaderRequest`

NewBrowserHeaderRequest instantiates a new BrowserHeaderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserHeaderRequestWithDefaults

`func NewBrowserHeaderRequestWithDefaults() *BrowserHeaderRequest`

NewBrowserHeaderRequestWithDefaults instantiates a new BrowserHeaderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BrowserHeaderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrowserHeaderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrowserHeaderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrigin

`func (o *BrowserHeaderRequest) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *BrowserHeaderRequest) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *BrowserHeaderRequest) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *BrowserHeaderRequest) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *BrowserHeaderRequest) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *BrowserHeaderRequest) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetValue

`func (o *BrowserHeaderRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BrowserHeaderRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BrowserHeaderRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BrowserHeaderRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


