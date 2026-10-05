# JsonFieldSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **NullableString** |  | [optional] 
**Multiple** | Pointer to **bool** |  | [optional] [default to false]
**Selector** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "text"]

## Methods

### NewJsonFieldSpec

`func NewJsonFieldSpec(selector string, ) *JsonFieldSpec`

NewJsonFieldSpec instantiates a new JsonFieldSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonFieldSpecWithDefaults

`func NewJsonFieldSpecWithDefaults() *JsonFieldSpec`

NewJsonFieldSpecWithDefaults instantiates a new JsonFieldSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *JsonFieldSpec) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *JsonFieldSpec) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *JsonFieldSpec) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *JsonFieldSpec) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### SetAttributeNil

`func (o *JsonFieldSpec) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *JsonFieldSpec) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetMultiple

`func (o *JsonFieldSpec) GetMultiple() bool`

GetMultiple returns the Multiple field if non-nil, zero value otherwise.

### GetMultipleOk

`func (o *JsonFieldSpec) GetMultipleOk() (*bool, bool)`

GetMultipleOk returns a tuple with the Multiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiple

`func (o *JsonFieldSpec) SetMultiple(v bool)`

SetMultiple sets Multiple field to given value.

### HasMultiple

`func (o *JsonFieldSpec) HasMultiple() bool`

HasMultiple returns a boolean if a field has been set.

### GetSelector

`func (o *JsonFieldSpec) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *JsonFieldSpec) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *JsonFieldSpec) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetType

`func (o *JsonFieldSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonFieldSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonFieldSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JsonFieldSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


