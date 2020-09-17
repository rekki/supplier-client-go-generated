# MainSetErrorOrderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | Pointer to **int32** | Number of attempts made to push the integration | [optional] 
**Error** | Pointer to **string** | error message | [optional] 
**Order** | Pointer to [**IntegrationOrder**](integration.Order.md) |  | [optional] 

## Methods

### NewMainSetErrorOrderInput

`func NewMainSetErrorOrderInput() *MainSetErrorOrderInput`

NewMainSetErrorOrderInput instantiates a new MainSetErrorOrderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainSetErrorOrderInputWithDefaults

`func NewMainSetErrorOrderInputWithDefaults() *MainSetErrorOrderInput`

NewMainSetErrorOrderInputWithDefaults instantiates a new MainSetErrorOrderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *MainSetErrorOrderInput) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *MainSetErrorOrderInput) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *MainSetErrorOrderInput) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *MainSetErrorOrderInput) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.

### GetError

`func (o *MainSetErrorOrderInput) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MainSetErrorOrderInput) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MainSetErrorOrderInput) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MainSetErrorOrderInput) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOrder

`func (o *MainSetErrorOrderInput) GetOrder() IntegrationOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MainSetErrorOrderInput) GetOrderOk() (*IntegrationOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MainSetErrorOrderInput) SetOrder(v IntegrationOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MainSetErrorOrderInput) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


