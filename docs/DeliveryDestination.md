# DeliveryDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**OnlyOnChange** | Pointer to **NullableBool** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**To** | **[]string** |  | 
**Type** | **string** |  | 
**Secret** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewDeliveryDestination

`func NewDeliveryDestination(to []string, type_ string, url string, ) *DeliveryDestination`

NewDeliveryDestination instantiates a new DeliveryDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryDestinationWithDefaults

`func NewDeliveryDestinationWithDefaults() *DeliveryDestination`

NewDeliveryDestinationWithDefaults instantiates a new DeliveryDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeliveryDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeliveryDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeliveryDestination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeliveryDestination) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeliveryDestination) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeliveryDestination) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOnlyOnChange

`func (o *DeliveryDestination) GetOnlyOnChange() bool`

GetOnlyOnChange returns the OnlyOnChange field if non-nil, zero value otherwise.

### GetOnlyOnChangeOk

`func (o *DeliveryDestination) GetOnlyOnChangeOk() (*bool, bool)`

GetOnlyOnChangeOk returns a tuple with the OnlyOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyOnChange

`func (o *DeliveryDestination) SetOnlyOnChange(v bool)`

SetOnlyOnChange sets OnlyOnChange field to given value.

### HasOnlyOnChange

`func (o *DeliveryDestination) HasOnlyOnChange() bool`

HasOnlyOnChange returns a boolean if a field has been set.

### SetOnlyOnChangeNil

`func (o *DeliveryDestination) SetOnlyOnChangeNil(b bool)`

 SetOnlyOnChangeNil sets the value for OnlyOnChange to be an explicit nil

### UnsetOnlyOnChange
`func (o *DeliveryDestination) UnsetOnlyOnChange()`

UnsetOnlyOnChange ensures that no value is present for OnlyOnChange, not even an explicit nil
### GetSubject

`func (o *DeliveryDestination) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *DeliveryDestination) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *DeliveryDestination) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *DeliveryDestination) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *DeliveryDestination) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *DeliveryDestination) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTo

`func (o *DeliveryDestination) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *DeliveryDestination) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *DeliveryDestination) SetTo(v []string)`

SetTo sets To field to given value.


### GetType

`func (o *DeliveryDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeliveryDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeliveryDestination) SetType(v string)`

SetType sets Type field to given value.


### GetSecret

`func (o *DeliveryDestination) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DeliveryDestination) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DeliveryDestination) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DeliveryDestination) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *DeliveryDestination) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *DeliveryDestination) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetUrl

`func (o *DeliveryDestination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeliveryDestination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeliveryDestination) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


