# HtmlOutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clean** | Pointer to **bool** |  | [optional] [default to true]
**Id** | Pointer to **NullableString** |  | [optional] 
**MainContentOnly** | Pointer to **bool** |  | [optional] [default to false]
**Type** | **string** |  | 

## Methods

### NewHtmlOutputSpec

`func NewHtmlOutputSpec(type_ string, ) *HtmlOutputSpec`

NewHtmlOutputSpec instantiates a new HtmlOutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHtmlOutputSpecWithDefaults

`func NewHtmlOutputSpecWithDefaults() *HtmlOutputSpec`

NewHtmlOutputSpecWithDefaults instantiates a new HtmlOutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClean

`func (o *HtmlOutputSpec) GetClean() bool`

GetClean returns the Clean field if non-nil, zero value otherwise.

### GetCleanOk

`func (o *HtmlOutputSpec) GetCleanOk() (*bool, bool)`

GetCleanOk returns a tuple with the Clean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClean

`func (o *HtmlOutputSpec) SetClean(v bool)`

SetClean sets Clean field to given value.

### HasClean

`func (o *HtmlOutputSpec) HasClean() bool`

HasClean returns a boolean if a field has been set.

### GetId

`func (o *HtmlOutputSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HtmlOutputSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HtmlOutputSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HtmlOutputSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *HtmlOutputSpec) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *HtmlOutputSpec) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMainContentOnly

`func (o *HtmlOutputSpec) GetMainContentOnly() bool`

GetMainContentOnly returns the MainContentOnly field if non-nil, zero value otherwise.

### GetMainContentOnlyOk

`func (o *HtmlOutputSpec) GetMainContentOnlyOk() (*bool, bool)`

GetMainContentOnlyOk returns a tuple with the MainContentOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContentOnly

`func (o *HtmlOutputSpec) SetMainContentOnly(v bool)`

SetMainContentOnly sets MainContentOnly field to given value.

### HasMainContentOnly

`func (o *HtmlOutputSpec) HasMainContentOnly() bool`

HasMainContentOnly returns a boolean if a field has been set.

### GetType

`func (o *HtmlOutputSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HtmlOutputSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HtmlOutputSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


