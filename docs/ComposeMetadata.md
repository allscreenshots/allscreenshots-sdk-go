# ComposeMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanvasHeight** | **int32** |  | 
**CanvasWidth** | **int32** |  | 
**Captures** | [**[]CaptureInfo**](CaptureInfo.md) |  | 
**LayoutType** | **string** |  | 
**TotalCaptures** | **int32** |  | 

## Methods

### NewComposeMetadata

`func NewComposeMetadata(canvasHeight int32, canvasWidth int32, captures []CaptureInfo, layoutType string, totalCaptures int32, ) *ComposeMetadata`

NewComposeMetadata instantiates a new ComposeMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeMetadataWithDefaults

`func NewComposeMetadataWithDefaults() *ComposeMetadata`

NewComposeMetadataWithDefaults instantiates a new ComposeMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanvasHeight

`func (o *ComposeMetadata) GetCanvasHeight() int32`

GetCanvasHeight returns the CanvasHeight field if non-nil, zero value otherwise.

### GetCanvasHeightOk

`func (o *ComposeMetadata) GetCanvasHeightOk() (*int32, bool)`

GetCanvasHeightOk returns a tuple with the CanvasHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasHeight

`func (o *ComposeMetadata) SetCanvasHeight(v int32)`

SetCanvasHeight sets CanvasHeight field to given value.


### GetCanvasWidth

`func (o *ComposeMetadata) GetCanvasWidth() int32`

GetCanvasWidth returns the CanvasWidth field if non-nil, zero value otherwise.

### GetCanvasWidthOk

`func (o *ComposeMetadata) GetCanvasWidthOk() (*int32, bool)`

GetCanvasWidthOk returns a tuple with the CanvasWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasWidth

`func (o *ComposeMetadata) SetCanvasWidth(v int32)`

SetCanvasWidth sets CanvasWidth field to given value.


### GetCaptures

`func (o *ComposeMetadata) GetCaptures() []CaptureInfo`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *ComposeMetadata) GetCapturesOk() (*[]CaptureInfo, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *ComposeMetadata) SetCaptures(v []CaptureInfo)`

SetCaptures sets Captures field to given value.


### GetLayoutType

`func (o *ComposeMetadata) GetLayoutType() string`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *ComposeMetadata) GetLayoutTypeOk() (*string, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *ComposeMetadata) SetLayoutType(v string)`

SetLayoutType sets LayoutType field to given value.


### GetTotalCaptures

`func (o *ComposeMetadata) GetTotalCaptures() int32`

GetTotalCaptures returns the TotalCaptures field if non-nil, zero value otherwise.

### GetTotalCapturesOk

`func (o *ComposeMetadata) GetTotalCapturesOk() (*int32, bool)`

GetTotalCapturesOk returns a tuple with the TotalCaptures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCaptures

`func (o *ComposeMetadata) SetTotalCaptures(v int32)`

SetTotalCaptures sets TotalCaptures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


