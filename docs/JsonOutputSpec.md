# JsonOutputSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Schema** | [**map[string]JsonFieldSpec**](JsonFieldSpec.md) |  | 
**Type** | **string** |  | 

## Methods

### NewJsonOutputSpec

`func NewJsonOutputSpec(schema map[string]JsonFieldSpec, type_ string, ) *JsonOutputSpec`

NewJsonOutputSpec instantiates a new JsonOutputSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonOutputSpecWithDefaults

`func NewJsonOutputSpecWithDefaults() *JsonOutputSpec`

NewJsonOutputSpecWithDefaults instantiates a new JsonOutputSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonOutputSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonOutputSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonOutputSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JsonOutputSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *JsonOutputSpec) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *JsonOutputSpec) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSchema

`func (o *JsonOutputSpec) GetSchema() map[string]JsonFieldSpec`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *JsonOutputSpec) GetSchemaOk() (*map[string]JsonFieldSpec, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *JsonOutputSpec) SetSchema(v map[string]JsonFieldSpec)`

SetSchema sets Schema field to given value.


### GetType

`func (o *JsonOutputSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonOutputSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonOutputSpec) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


