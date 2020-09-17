# V3SetIntegratedOrdersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | Array of References of the orders to mark as integrated, required. Order refs are discoverable when listing orders. | [optional] 

## Methods

### NewV3SetIntegratedOrdersInput

`func NewV3SetIntegratedOrdersInput() *V3SetIntegratedOrdersInput`

NewV3SetIntegratedOrdersInput instantiates a new V3SetIntegratedOrdersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3SetIntegratedOrdersInputWithDefaults

`func NewV3SetIntegratedOrdersInputWithDefaults() *V3SetIntegratedOrdersInput`

NewV3SetIntegratedOrdersInputWithDefaults instantiates a new V3SetIntegratedOrdersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *V3SetIntegratedOrdersInput) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *V3SetIntegratedOrdersInput) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *V3SetIntegratedOrdersInput) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *V3SetIntegratedOrdersInput) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


