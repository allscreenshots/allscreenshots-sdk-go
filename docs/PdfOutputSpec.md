# PdfOutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Landscape** | Pointer to **bool** |  | [optional] [default to false]
**PrintBackground** | Pointer to **bool** |  | [optional] [default to true]
**Type** | **string** |  | 

## Methods

### NewPdfOutputSpec

`func NewPdfOutputSpec(type_ string, ) *PdfOutputSpec`

NewPdfOutputSpec instantiates a new PdfOutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdfOutputSpecWithDefaults

`func NewPdfOutputSpecWithDefaults() *PdfOutputSpec`

NewPdfOutputSpecWithDefaults instantiates a new PdfOutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *PdfOutputSpec) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PdfOutputSpec) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PdfOutputSpec) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *PdfOutputSpec) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *PdfOutputSpec) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *PdfOutputSpec) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetId

`func (o *PdfOutputSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PdfOutputSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PdfOutputSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PdfOutputSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PdfOutputSpec) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PdfOutputSpec) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLandscape

`func (o *PdfOutputSpec) GetLandscape() bool`

GetLandscape returns the Landscape field if non-nil, zero value otherwise.

### GetLandscapeOk

`func (o *PdfOutputSpec) GetLandscapeOk() (*bool, bool)`

GetLandscapeOk returns a tuple with the Landscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscape

`func (o *PdfOutputSpec) SetLandscape(v bool)`

SetLandscape sets Landscape field to given value.

### HasLandscape

`func (o *PdfOutputSpec) HasLandscape() bool`

HasLandscape returns a boolean if a field has been set.

### GetPrintBackground

`func (o *PdfOutputSpec) GetPrintBackground() bool`

GetPrintBackground returns the PrintBackground field if non-nil, zero value otherwise.

### GetPrintBackgroundOk

`func (o *PdfOutputSpec) GetPrintBackgroundOk() (*bool, bool)`

GetPrintBackgroundOk returns a tuple with the PrintBackground field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintBackground

`func (o *PdfOutputSpec) SetPrintBackground(v bool)`

SetPrintBackground sets PrintBackground field to given value.

### HasPrintBackground

`func (o *PdfOutputSpec) HasPrintBackground() bool`

HasPrintBackground returns a boolean if a field has been set.

### GetType

`func (o *PdfOutputSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PdfOutputSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PdfOutputSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


