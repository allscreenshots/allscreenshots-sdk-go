# EmailDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**OnlyOnChange** | Pointer to **NullableBool** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**To** | **[]string** |  | 
**Type** | **string** |  | 

## Methods

### NewEmailDestination

`func NewEmailDestination(to []string, type_ string, ) *EmailDestination`

NewEmailDestination instantiates a new EmailDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailDestinationWithDefaults

`func NewEmailDestinationWithDefaults() *EmailDestination`

NewEmailDestinationWithDefaults instantiates a new EmailDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailDestination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailDestination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailDestination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailDestination) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EmailDestination) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EmailDestination) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetOnlyOnChange

`func (o *EmailDestination) GetOnlyOnChange() bool`

GetOnlyOnChange returns the OnlyOnChange field if non-nil, zero value otherwise.

### GetOnlyOnChangeOk

`func (o *EmailDestination) GetOnlyOnChangeOk() (*bool, bool)`

GetOnlyOnChangeOk returns a tuple with the OnlyOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyOnChange

`func (o *EmailDestination) SetOnlyOnChange(v bool)`

SetOnlyOnChange sets OnlyOnChange field to given value.

### HasOnlyOnChange

`func (o *EmailDestination) HasOnlyOnChange() bool`

HasOnlyOnChange returns a boolean if a field has been set.

### SetOnlyOnChangeNil

`func (o *EmailDestination) SetOnlyOnChangeNil(b bool)`

 SetOnlyOnChangeNil sets the value for OnlyOnChange to be an explicit nil

### UnsetOnlyOnChange
`func (o *EmailDestination) UnsetOnlyOnChange()`

UnsetOnlyOnChange ensures that no value is present for OnlyOnChange, not even an explicit nil
### GetSubject

`func (o *EmailDestination) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *EmailDestination) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *EmailDestination) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *EmailDestination) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *EmailDestination) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *EmailDestination) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTo

`func (o *EmailDestination) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailDestination) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailDestination) SetTo(v []string)`

SetTo sets To field to given value.


### GetType

`func (o *EmailDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmailDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmailDestination) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


