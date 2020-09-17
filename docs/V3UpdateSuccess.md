# V3UpdateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **int32** | number of items actually updated | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewV3UpdateSuccess

`func NewV3UpdateSuccess() *V3UpdateSuccess`

NewV3UpdateSuccess instantiates a new V3UpdateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3UpdateSuccessWithDefaults

`func NewV3UpdateSuccessWithDefaults() *V3UpdateSuccess`

NewV3UpdateSuccessWithDefaults instantiates a new V3UpdateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *V3UpdateSuccess) GetAffected() int32`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *V3UpdateSuccess) GetAffectedOk() (*int32, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *V3UpdateSuccess) SetAffected(v int32)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *V3UpdateSuccess) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetSuccess

`func (o *V3UpdateSuccess) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *V3UpdateSuccess) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *V3UpdateSuccess) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *V3UpdateSuccess) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


