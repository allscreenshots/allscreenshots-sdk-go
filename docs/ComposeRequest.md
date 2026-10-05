# ComposeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Async** | Pointer to **bool** |  | [optional] 
**Captures** | Pointer to [**[]CaptureItem**](CaptureItem.md) |  | [optional] 
**CapturesMode** | Pointer to **bool** |  | [optional] 
**Defaults** | Pointer to [**NullableCaptureDefaults**](CaptureDefaults.md) |  | [optional] 
**Output** | Pointer to [**ComposeOutputConfig**](ComposeOutputConfig.md) |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Variants** | Pointer to [**[]VariantConfig**](VariantConfig.md) |  | [optional] 
**VariantsMode** | Pointer to **bool** |  | [optional] 
**WebhookSecret** | Pointer to **NullableString** |  | [optional] 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewComposeRequest

`func NewComposeRequest() *ComposeRequest`

NewComposeRequest instantiates a new ComposeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComposeRequestWithDefaults

`func NewComposeRequestWithDefaults() *ComposeRequest`

NewComposeRequestWithDefaults instantiates a new ComposeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsync

`func (o *ComposeRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *ComposeRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *ComposeRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *ComposeRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### GetCaptures

`func (o *ComposeRequest) GetCaptures() []CaptureItem`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *ComposeRequest) GetCapturesOk() (*[]CaptureItem, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *ComposeRequest) SetCaptures(v []CaptureItem)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *ComposeRequest) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### SetCapturesNil

`func (o *ComposeRequest) SetCapturesNil(b bool)`

 SetCapturesNil sets the value for Captures to be an explicit nil

### UnsetCaptures
`func (o *ComposeRequest) UnsetCaptures()`

UnsetCaptures ensures that no value is present for Captures, not even an explicit nil
### GetCapturesMode

`func (o *ComposeRequest) GetCapturesMode() bool`

GetCapturesMode returns the CapturesMode field if non-nil, zero value otherwise.

### GetCapturesModeOk

`func (o *ComposeRequest) GetCapturesModeOk() (*bool, bool)`

GetCapturesModeOk returns a tuple with the CapturesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapturesMode

`func (o *ComposeRequest) SetCapturesMode(v bool)`

SetCapturesMode sets CapturesMode field to given value.

### HasCapturesMode

`func (o *ComposeRequest) HasCapturesMode() bool`

HasCapturesMode returns a boolean if a field has been set.

### GetDefaults

`func (o *ComposeRequest) GetDefaults() CaptureDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *ComposeRequest) GetDefaultsOk() (*CaptureDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *ComposeRequest) SetDefaults(v CaptureDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *ComposeRequest) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *ComposeRequest) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *ComposeRequest) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetOutput

`func (o *ComposeRequest) GetOutput() ComposeOutputConfig`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ComposeRequest) GetOutputOk() (*ComposeOutputConfig, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ComposeRequest) SetOutput(v ComposeOutputConfig)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ComposeRequest) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetUrl

`func (o *ComposeRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ComposeRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ComposeRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ComposeRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ComposeRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ComposeRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVariants

`func (o *ComposeRequest) GetVariants() []VariantConfig`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ComposeRequest) GetVariantsOk() (*[]VariantConfig, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ComposeRequest) SetVariants(v []VariantConfig)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ComposeRequest) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### SetVariantsNil

`func (o *ComposeRequest) SetVariantsNil(b bool)`

 SetVariantsNil sets the value for Variants to be an explicit nil

### UnsetVariants
`func (o *ComposeRequest) UnsetVariants()`

UnsetVariants ensures that no value is present for Variants, not even an explicit nil
### GetVariantsMode

`func (o *ComposeRequest) GetVariantsMode() bool`

GetVariantsMode returns the VariantsMode field if non-nil, zero value otherwise.

### GetVariantsModeOk

`func (o *ComposeRequest) GetVariantsModeOk() (*bool, bool)`

GetVariantsModeOk returns a tuple with the VariantsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantsMode

`func (o *ComposeRequest) SetVariantsMode(v bool)`

SetVariantsMode sets VariantsMode field to given value.

### HasVariantsMode

`func (o *ComposeRequest) HasVariantsMode() bool`

HasVariantsMode returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *ComposeRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *ComposeRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *ComposeRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *ComposeRequest) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *ComposeRequest) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *ComposeRequest) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetWebhookUrl

`func (o *ComposeRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *ComposeRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *ComposeRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *ComposeRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *ComposeRequest) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *ComposeRequest) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


