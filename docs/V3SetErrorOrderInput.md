# V3SetErrorOrderInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]V3OrderIntegrationError**](v3.OrderIntegrationError.md) | list of orders failed to integrate, required | [optional] 

## Methods

### NewV3SetErrorOrderInput

`func NewV3SetErrorOrderInput() *V3SetErrorOrderInput`

NewV3SetErrorOrderInput instantiates a new V3SetErrorOrderInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SetErrorOrderInputWithDefaults

`func NewV3SetErrorOrderInputWithDefaults() *V3SetErrorOrderInput`

NewV3SetErrorOrderInputWithDefaults instantiates a new V3SetErrorOrderInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *V3SetErrorOrderInput) GetOrders() []V3OrderIntegrationError`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *V3SetErrorOrderInput) GetOrdersOk() (*[]V3OrderIntegrationError, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *V3SetErrorOrderInput) SetOrders(v []V3OrderIntegrationError)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *V3SetErrorOrderInput) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


