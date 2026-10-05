# LayoutPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanvasHeight** | **int32** |  | 
**CanvasWidth** | **int32** |  | 
**Layout** | **string** |  | 
**Metadata** | **map[string]map[string]interface{}** |  | 
**Placements** | [**[]PlacementPreview**](PlacementPreview.md) |  | 
**ResolvedLayout** | **string** |  | 

## Methods

### NewLayoutPreviewResponse

`func NewLayoutPreviewResponse(canvasHeight int32, canvasWidth int32, layout string, metadata map[string]map[string]interface{}, placements []PlacementPreview, resolvedLayout string, ) *LayoutPreviewResponse`

NewLayoutPreviewResponse instantiates a new LayoutPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutPreviewResponseWithDefaults

`func NewLayoutPreviewResponseWithDefaults() *LayoutPreviewResponse`

NewLayoutPreviewResponseWithDefaults instantiates a new LayoutPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanvasHeight

`func (o *LayoutPreviewResponse) GetCanvasHeight() int32`

GetCanvasHeight returns the CanvasHeight field if non-nil, zero value otherwise.

### GetCanvasHeightOk

`func (o *LayoutPreviewResponse) GetCanvasHeightOk() (*int32, bool)`

GetCanvasHeightOk returns a tuple with the CanvasHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasHeight

`func (o *LayoutPreviewResponse) SetCanvasHeight(v int32)`

SetCanvasHeight sets CanvasHeight field to given value.


### GetCanvasWidth

`func (o *LayoutPreviewResponse) GetCanvasWidth() int32`

GetCanvasWidth returns the CanvasWidth field if non-nil, zero value otherwise.

### GetCanvasWidthOk

`func (o *LayoutPreviewResponse) GetCanvasWidthOk() (*int32, bool)`

GetCanvasWidthOk returns a tuple with the CanvasWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasWidth

`func (o *LayoutPreviewResponse) SetCanvasWidth(v int32)`

SetCanvasWidth sets CanvasWidth field to given value.


### GetLayout

`func (o *LayoutPreviewResponse) GetLayout() string`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *LayoutPreviewResponse) GetLayoutOk() (*string, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *LayoutPreviewResponse) SetLayout(v string)`

SetLayout sets Layout field to given value.


### GetMetadata

`func (o *LayoutPreviewResponse) GetMetadata() map[string]map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LayoutPreviewResponse) GetMetadataOk() (*map[string]map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LayoutPreviewResponse) SetMetadata(v map[string]map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetPlacements

`func (o *LayoutPreviewResponse) GetPlacements() []PlacementPreview`

GetPlacements returns the Placements field if non-nil, zero value otherwise.

### GetPlacementsOk

`func (o *LayoutPreviewResponse) GetPlacementsOk() (*[]PlacementPreview, bool)`

GetPlacementsOk returns a tuple with the Placements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacements

`func (o *LayoutPreviewResponse) SetPlacements(v []PlacementPreview)`

SetPlacements sets Placements field to given value.


### GetResolvedLayout

`func (o *LayoutPreviewResponse) GetResolvedLayout() string`

GetResolvedLayout returns the ResolvedLayout field if non-nil, zero value otherwise.

### GetResolvedLayoutOk

`func (o *LayoutPreviewResponse) GetResolvedLayoutOk() (*string, bool)`

GetResolvedLayoutOk returns a tuple with the ResolvedLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedLayout

`func (o *LayoutPreviewResponse) SetResolvedLayout(v string)`

SetResolvedLayout sets ResolvedLayout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


