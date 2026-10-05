# ScreenshotOutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] [default to "png"]
**FullPage** | Pointer to **bool** |  | [optional] [default to false]
**Id** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to **int32** |  | [optional] [default to 80]
**Selector** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewScreenshotOutputSpec

`func NewScreenshotOutputSpec(type_ string, ) *ScreenshotOutputSpec`

NewScreenshotOutputSpec instantiates a new ScreenshotOutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScreenshotOutputSpecWithDefaults

`func NewScreenshotOutputSpecWithDefaults() *ScreenshotOutputSpec`

NewScreenshotOutputSpecWithDefaults instantiates a new ScreenshotOutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ScreenshotOutputSpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ScreenshotOutputSpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ScreenshotOutputSpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ScreenshotOutputSpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFullPage

`func (o *ScreenshotOutputSpec) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *ScreenshotOutputSpec) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *ScreenshotOutputSpec) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *ScreenshotOutputSpec) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetId

`func (o *ScreenshotOutputSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScreenshotOutputSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScreenshotOutputSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScreenshotOutputSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ScreenshotOutputSpec) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ScreenshotOutputSpec) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuality

`func (o *ScreenshotOutputSpec) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ScreenshotOutputSpec) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ScreenshotOutputSpec) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ScreenshotOutputSpec) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetSelector

`func (o *ScreenshotOutputSpec) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *ScreenshotOutputSpec) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *ScreenshotOutputSpec) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *ScreenshotOutputSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *ScreenshotOutputSpec) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *ScreenshotOutputSpec) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetType

`func (o *ScreenshotOutputSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScreenshotOutputSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScreenshotOutputSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


