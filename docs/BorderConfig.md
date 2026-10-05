# BorderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** |  | [optional] [default to "#cccccc"]
**Radius** | Pointer to **int32** |  | [optional] [default to 0]
**Width** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewBorderConfig

`func NewBorderConfig() *BorderConfig`

NewBorderConfig instantiates a new BorderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBorderConfigWithDefaults

`func NewBorderConfigWithDefaults() *BorderConfig`

NewBorderConfigWithDefaults instantiates a new BorderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *BorderConfig) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BorderConfig) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BorderConfig) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *BorderConfig) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetRadius

`func (o *BorderConfig) GetRadius() int32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *BorderConfig) GetRadiusOk() (*int32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *BorderConfig) SetRadius(v int32)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *BorderConfig) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetWidth

`func (o *BorderConfig) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *BorderConfig) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *BorderConfig) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *BorderConfig) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


