# WebhookDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**OnlyOnChange** | Pointer to **NullableBool** |  | [optional] 
**Secret** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewWebhookDestination

`func NewWebhookDestination(type_ string, url string, ) *WebhookDestination`

NewWebhookDestination instantiates a new WebhookDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDestinationWithDefaults

`func NewWebhookDestinationWithDefaults() *WebhookDestination`

NewWebhookDestinationWithDefaults instantiates a new WebhookDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDestination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookDestination) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WebhookDestination) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebhookDestination) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOnlyOnChange

`func (o *WebhookDestination) GetOnlyOnChange() bool`

GetOnlyOnChange returns the OnlyOnChange field if non-nil, zero value otherwise.

### GetOnlyOnChangeOk

`func (o *WebhookDestination) GetOnlyOnChangeOk() (*bool, bool)`

GetOnlyOnChangeOk returns a tuple with the OnlyOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyOnChange

`func (o *WebhookDestination) SetOnlyOnChange(v bool)`

SetOnlyOnChange sets OnlyOnChange field to given value.

### HasOnlyOnChange

`func (o *WebhookDestination) HasOnlyOnChange() bool`

HasOnlyOnChange returns a boolean if a field has been set.

### SetOnlyOnChangeNil

`func (o *WebhookDestination) SetOnlyOnChangeNil(b bool)`

 SetOnlyOnChangeNil sets the value for OnlyOnChange to be an explicit nil

### UnsetOnlyOnChange
`func (o *WebhookDestination) UnsetOnlyOnChange()`

UnsetOnlyOnChange ensures that no value is present for OnlyOnChange, not even an explicit nil
### GetSecret

`func (o *WebhookDestination) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookDestination) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookDestination) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookDestination) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *WebhookDestination) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *WebhookDestination) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetType

`func (o *WebhookDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookDestination) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *WebhookDestination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookDestination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookDestination) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


