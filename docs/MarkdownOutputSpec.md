# MarkdownOutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**MainContentOnly** | Pointer to **bool** |  | [optional] [default to true]
**Type** | **string** |  | 

## Methods

### NewMarkdownOutputSpec

`func NewMarkdownOutputSpec(type_ string, ) *MarkdownOutputSpec`

NewMarkdownOutputSpec instantiates a new MarkdownOutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkdownOutputSpecWithDefaults

`func NewMarkdownOutputSpecWithDefaults() *MarkdownOutputSpec`

NewMarkdownOutputSpecWithDefaults instantiates a new MarkdownOutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarkdownOutputSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarkdownOutputSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarkdownOutputSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MarkdownOutputSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MarkdownOutputSpec) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MarkdownOutputSpec) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMainContentOnly

`func (o *MarkdownOutputSpec) GetMainContentOnly() bool`

GetMainContentOnly returns the MainContentOnly field if non-nil, zero value otherwise.

### GetMainContentOnlyOk

`func (o *MarkdownOutputSpec) GetMainContentOnlyOk() (*bool, bool)`

GetMainContentOnlyOk returns a tuple with the MainContentOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContentOnly

`func (o *MarkdownOutputSpec) SetMainContentOnly(v bool)`

SetMainContentOnly sets MainContentOnly field to given value.

### HasMainContentOnly

`func (o *MarkdownOutputSpec) HasMainContentOnly() bool`

HasMainContentOnly returns a boolean if a field has been set.

### GetType

`func (o *MarkdownOutputSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarkdownOutputSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarkdownOutputSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


