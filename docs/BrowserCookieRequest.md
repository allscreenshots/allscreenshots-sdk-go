# BrowserCookieRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **NullableString** |  | [optional] 
**Expires** | Pointer to **NullableFloat64** |  | [optional] 
**HttpOnly** | Pointer to **NullableBool** |  | [optional] 
**Name** | **string** |  | 
**Path** | Pointer to **NullableString** |  | [optional] [default to "/"]
**SameSite** | Pointer to **NullableString** |  | [optional] 
**Secure** | Pointer to **NullableBool** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewBrowserCookieRequest

`func NewBrowserCookieRequest(name string, ) *BrowserCookieRequest`

NewBrowserCookieRequest instantiates a new BrowserCookieRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrowserCookieRequestWithDefaults

`func NewBrowserCookieRequestWithDefaults() *BrowserCookieRequest`

NewBrowserCookieRequestWithDefaults instantiates a new BrowserCookieRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *BrowserCookieRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BrowserCookieRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BrowserCookieRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *BrowserCookieRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *BrowserCookieRequest) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *BrowserCookieRequest) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetExpires

`func (o *BrowserCookieRequest) GetExpires() float64`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *BrowserCookieRequest) GetExpiresOk() (*float64, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *BrowserCookieRequest) SetExpires(v float64)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *BrowserCookieRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *BrowserCookieRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *BrowserCookieRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetHttpOnly

`func (o *BrowserCookieRequest) GetHttpOnly() bool`

GetHttpOnly returns the HttpOnly field if non-nil, zero value otherwise.

### GetHttpOnlyOk

`func (o *BrowserCookieRequest) GetHttpOnlyOk() (*bool, bool)`

GetHttpOnlyOk returns a tuple with the HttpOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpOnly

`func (o *BrowserCookieRequest) SetHttpOnly(v bool)`

SetHttpOnly sets HttpOnly field to given value.

### HasHttpOnly

`func (o *BrowserCookieRequest) HasHttpOnly() bool`

HasHttpOnly returns a boolean if a field has been set.

### SetHttpOnlyNil

`func (o *BrowserCookieRequest) SetHttpOnlyNil(b bool)`

 SetHttpOnlyNil sets the value for HttpOnly to be an explicit nil

### UnsetHttpOnly
`func (o *BrowserCookieRequest) UnsetHttpOnly()`

UnsetHttpOnly ensures that no value is present for HttpOnly, not even an explicit nil
### GetName

`func (o *BrowserCookieRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrowserCookieRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrowserCookieRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *BrowserCookieRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BrowserCookieRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BrowserCookieRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BrowserCookieRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BrowserCookieRequest) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BrowserCookieRequest) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSameSite

`func (o *BrowserCookieRequest) GetSameSite() string`

GetSameSite returns the SameSite field if non-nil, zero value otherwise.

### GetSameSiteOk

`func (o *BrowserCookieRequest) GetSameSiteOk() (*string, bool)`

GetSameSiteOk returns a tuple with the SameSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameSite

`func (o *BrowserCookieRequest) SetSameSite(v string)`

SetSameSite sets SameSite field to given value.

### HasSameSite

`func (o *BrowserCookieRequest) HasSameSite() bool`

HasSameSite returns a boolean if a field has been set.

### SetSameSiteNil

`func (o *BrowserCookieRequest) SetSameSiteNil(b bool)`

 SetSameSiteNil sets the value for SameSite to be an explicit nil

### UnsetSameSite
`func (o *BrowserCookieRequest) UnsetSameSite()`

UnsetSameSite ensures that no value is present for SameSite, not even an explicit nil
### GetSecure

`func (o *BrowserCookieRequest) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *BrowserCookieRequest) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *BrowserCookieRequest) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *BrowserCookieRequest) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### SetSecureNil

`func (o *BrowserCookieRequest) SetSecureNil(b bool)`

 SetSecureNil sets the value for Secure to be an explicit nil

### UnsetSecure
`func (o *BrowserCookieRequest) UnsetSecure()`

UnsetSecure ensures that no value is present for Secure, not even an explicit nil
### GetValue

`func (o *BrowserCookieRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BrowserCookieRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BrowserCookieRequest) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BrowserCookieRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


