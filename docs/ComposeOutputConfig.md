# ComposeOutputConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to **string** |  | [optional] [default to "center"]
**Background** | Pointer to **string** |  | [optional] [default to "#ffffff"]
**Border** | Pointer to [**NullableBorderConfig**](BorderConfig.md) |  | [optional] 
**Columns** | Pointer to **NullableInt32** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "png"]
**Labels** | Pointer to [**NullableLabelConfig**](LabelConfig.md) |  | [optional] 
**Layout** | Pointer to **string** |  | [optional] [default to "AUTO"]
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**Padding** | Pointer to **int32** |  | [optional] [default to 0]
**Quality** | Pointer to **int32** |  | [optional] [default to 90]
**Shadow** | Pointer to [**NullableShadowConfig**](ShadowConfig.md) |  | [optional] 
**Spacing** | Pointer to **int32** |  | [optional] [default to 10]
**ThumbnailWidth** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewComposeOutputConfig

`func NewComposeOutputConfig() *ComposeOutputConfig`

NewComposeOutputConfig instantiates a new ComposeOutputConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeOutputConfigWithDefaults

`func NewComposeOutputConfigWithDefaults() *ComposeOutputConfig`

NewComposeOutputConfigWithDefaults instantiates a new ComposeOutputConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *ComposeOutputConfig) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *ComposeOutputConfig) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *ComposeOutputConfig) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *ComposeOutputConfig) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetBackground

`func (o *ComposeOutputConfig) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *ComposeOutputConfig) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *ComposeOutputConfig) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *ComposeOutputConfig) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBorder

`func (o *ComposeOutputConfig) GetBorder() BorderConfig`

GetBorder returns the Border field if non-nil, zero value otherwise.

### GetBorderOk

`func (o *ComposeOutputConfig) GetBorderOk() (*BorderConfig, bool)`

GetBorderOk returns a tuple with the Border field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorder

`func (o *ComposeOutputConfig) SetBorder(v BorderConfig)`

SetBorder sets Border field to given value.

### HasBorder

`func (o *ComposeOutputConfig) HasBorder() bool`

HasBorder returns a boolean if a field has been set.

### SetBorderNil

`func (o *ComposeOutputConfig) SetBorderNil(b bool)`

 SetBorderNil sets the value for Border to be an explicit nil

### UnsetBorder
`func (o *ComposeOutputConfig) UnsetBorder()`

UnsetBorder ensures that no value is present for Border, not even an explicit nil
### GetColumns

`func (o *ComposeOutputConfig) GetColumns() int32`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *ComposeOutputConfig) GetColumnsOk() (*int32, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *ComposeOutputConfig) SetColumns(v int32)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *ComposeOutputConfig) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *ComposeOutputConfig) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *ComposeOutputConfig) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetFormat

`func (o *ComposeOutputConfig) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ComposeOutputConfig) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ComposeOutputConfig) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ComposeOutputConfig) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetLabels

`func (o *ComposeOutputConfig) GetLabels() LabelConfig`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ComposeOutputConfig) GetLabelsOk() (*LabelConfig, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ComposeOutputConfig) SetLabels(v LabelConfig)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ComposeOutputConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ComposeOutputConfig) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ComposeOutputConfig) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetLayout

`func (o *ComposeOutputConfig) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ComposeOutputConfig) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ComposeOutputConfig) SetLayout(v string)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ComposeOutputConfig) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetMaxHeight

`func (o *ComposeOutputConfig) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ComposeOutputConfig) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ComposeOutputConfig) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ComposeOutputConfig) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *ComposeOutputConfig) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *ComposeOutputConfig) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetMaxWidth

`func (o *ComposeOutputConfig) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *ComposeOutputConfig) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *ComposeOutputConfig) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *ComposeOutputConfig) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *ComposeOutputConfig) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *ComposeOutputConfig) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetPadding

`func (o *ComposeOutputConfig) GetPadding() int32`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *ComposeOutputConfig) GetPaddingOk() (*int32, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *ComposeOutputConfig) SetPadding(v int32)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *ComposeOutputConfig) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetQuality

`func (o *ComposeOutputConfig) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ComposeOutputConfig) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ComposeOutputConfig) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ComposeOutputConfig) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetShadow

`func (o *ComposeOutputConfig) GetShadow() ShadowConfig`

GetShadow returns the Shadow field if non-nil, zero value otherwise.

### GetShadowOk

`func (o *ComposeOutputConfig) GetShadowOk() (*ShadowConfig, bool)`

GetShadowOk returns a tuple with the Shadow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShadow

`func (o *ComposeOutputConfig) SetShadow(v ShadowConfig)`

SetShadow sets Shadow field to given value.

### HasShadow

`func (o *ComposeOutputConfig) HasShadow() bool`

HasShadow returns a boolean if a field has been set.

### SetShadowNil

`func (o *ComposeOutputConfig) SetShadowNil(b bool)`

 SetShadowNil sets the value for Shadow to be an explicit nil

### UnsetShadow
`func (o *ComposeOutputConfig) UnsetShadow()`

UnsetShadow ensures that no value is present for Shadow, not even an explicit nil
### GetSpacing

`func (o *ComposeOutputConfig) GetSpacing() int32`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *ComposeOutputConfig) GetSpacingOk() (*int32, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *ComposeOutputConfig) SetSpacing(v int32)`

SetSpacing sets Spacing field to given value.

### HasSpacing

`func (o *ComposeOutputConfig) HasSpacing() bool`

HasSpacing returns a boolean if a field has been set.

### GetThumbnailWidth

`func (o *ComposeOutputConfig) GetThumbnailWidth() int32`

GetThumbnailWidth returns the ThumbnailWidth field if non-nil, zero value otherwise.

### GetThumbnailWidthOk

`func (o *ComposeOutputConfig) GetThumbnailWidthOk() (*int32, bool)`

GetThumbnailWidthOk returns a tuple with the ThumbnailWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailWidth

`func (o *ComposeOutputConfig) SetThumbnailWidth(v int32)`

SetThumbnailWidth sets ThumbnailWidth field to given value.

### HasThumbnailWidth

`func (o *ComposeOutputConfig) HasThumbnailWidth() bool`

HasThumbnailWidth returns a boolean if a field has been set.

### SetThumbnailWidthNil

`func (o *ComposeOutputConfig) SetThumbnailWidthNil(b bool)`

 SetThumbnailWidthNil sets the value for ThumbnailWidth to be an explicit nil

### UnsetThumbnailWidth
`func (o *ComposeOutputConfig) UnsetThumbnailWidth()`

UnsetThumbnailWidth ensures that no value is present for ThumbnailWidth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


