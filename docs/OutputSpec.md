# OutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** |  | [optional] 
**FullPage** | Pointer to **bool** |  | [optional] [default to false]
**Id** | Pointer to **NullableString** |  | [optional] 
**Quality** | Pointer to **int32** |  | [optional] [default to 80]
**Selector** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Landscape** | Pointer to **bool** |  | [optional] [default to false]
**PrintBackground** | Pointer to **bool** |  | [optional] [default to true]
**Clean** | Pointer to **bool** |  | [optional] [default to true]
**MainContentOnly** | Pointer to **bool** |  | [optional] [default to true]
**Schema** | [**map[string]JsonFieldSpec**](JsonFieldSpec.md) |  | 

## Methods

### NewOutputSpec

`func NewOutputSpec(type_ string, schema map[string]JsonFieldSpec, ) *OutputSpec`

NewOutputSpec instantiates a new OutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSpecWithDefaults

`func NewOutputSpecWithDefaults() *OutputSpec`

NewOutputSpecWithDefaults instantiates a new OutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *OutputSpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OutputSpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OutputSpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *OutputSpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *OutputSpec) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *OutputSpec) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetFullPage

`func (o *OutputSpec) GetFullPage() bool`

GetFullPage returns the FullPage field if non-nil, zero value otherwise.

### GetFullPageOk

`func (o *OutputSpec) GetFullPageOk() (*bool, bool)`

GetFullPageOk returns a tuple with the FullPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullPage

`func (o *OutputSpec) SetFullPage(v bool)`

SetFullPage sets FullPage field to given value.

### HasFullPage

`func (o *OutputSpec) HasFullPage() bool`

HasFullPage returns a boolean if a field has been set.

### GetId

`func (o *OutputSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutputSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OutputSpec) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OutputSpec) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetQuality

`func (o *OutputSpec) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *OutputSpec) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *OutputSpec) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *OutputSpec) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetSelector

`func (o *OutputSpec) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *OutputSpec) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *OutputSpec) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *OutputSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### SetSelectorNil

`func (o *OutputSpec) SetSelectorNil(b bool)`

 SetSelectorNil sets the value for Selector to be an explicit nil

### UnsetSelector
`func (o *OutputSpec) UnsetSelector()`

UnsetSelector ensures that no value is present for Selector, not even an explicit nil
### GetType

`func (o *OutputSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputSpec) SetType(v string)`

SetType sets Type field to given value.


### GetLandscape

`func (o *OutputSpec) GetLandscape() bool`

GetLandscape returns the Landscape field if non-nil, zero value otherwise.

### GetLandscapeOk

`func (o *OutputSpec) GetLandscapeOk() (*bool, bool)`

GetLandscapeOk returns a tuple with the Landscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscape

`func (o *OutputSpec) SetLandscape(v bool)`

SetLandscape sets Landscape field to given value.

### HasLandscape

`func (o *OutputSpec) HasLandscape() bool`

HasLandscape returns a boolean if a field has been set.

### GetPrintBackground

`func (o *OutputSpec) GetPrintBackground() bool`

GetPrintBackground returns the PrintBackground field if non-nil, zero value otherwise.

### GetPrintBackgroundOk

`func (o *OutputSpec) GetPrintBackgroundOk() (*bool, bool)`

GetPrintBackgroundOk returns a tuple with the PrintBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBackground

`func (o *OutputSpec) SetPrintBackground(v bool)`

SetPrintBackground sets PrintBackground field to given value.

### HasPrintBackground

`func (o *OutputSpec) HasPrintBackground() bool`

HasPrintBackground returns a boolean if a field has been set.

### GetClean

`func (o *OutputSpec) GetClean() bool`

GetClean returns the Clean field if non-nil, zero value otherwise.

### GetCleanOk

`func (o *OutputSpec) GetCleanOk() (*bool, bool)`

GetCleanOk returns a tuple with the Clean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClean

`func (o *OutputSpec) SetClean(v bool)`

SetClean sets Clean field to given value.

### HasClean

`func (o *OutputSpec) HasClean() bool`

HasClean returns a boolean if a field has been set.

### GetMainContentOnly

`func (o *OutputSpec) GetMainContentOnly() bool`

GetMainContentOnly returns the MainContentOnly field if non-nil, zero value otherwise.

### GetMainContentOnlyOk

`func (o *OutputSpec) GetMainContentOnlyOk() (*bool, bool)`

GetMainContentOnlyOk returns a tuple with the MainContentOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainContentOnly

`func (o *OutputSpec) SetMainContentOnly(v bool)`

SetMainContentOnly sets MainContentOnly field to given value.

### HasMainContentOnly

`func (o *OutputSpec) HasMainContentOnly() bool`

HasMainContentOnly returns a boolean if a field has been set.

### GetSchema

`func (o *OutputSpec) GetSchema() map[string]JsonFieldSpec`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OutputSpec) GetSchemaOk() (*map[string]JsonFieldSpec, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OutputSpec) SetSchema(v map[string]JsonFieldSpec)`

SetSchema sets Schema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


