# MainSetIntegrateOrdersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]map[string]interface{}** | Array of References of the orders to confirm, required. Order refs are discoverable when listing orders. | [optional] 

## Methods

### NewMainSetIntegrateOrdersInput

`func NewMainSetIntegrateOrdersInput() *MainSetIntegrateOrdersInput`

NewMainSetIntegrateOrdersInput instantiates a new MainSetIntegrateOrdersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainSetIntegrateOrdersInputWithDefaults

`func NewMainSetIntegrateOrdersInputWithDefaults() *MainSetIntegrateOrdersInput`

NewMainSetIntegrateOrdersInputWithDefaults instantiates a new MainSetIntegrateOrdersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *MainSetIntegrateOrdersInput) GetOrders() []map[string]interface{}`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MainSetIntegrateOrdersInput) GetOrdersOk() (*[]map[string]interface{}, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MainSetIntegrateOrdersInput) SetOrders(v []map[string]interface{})`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MainSetIntegrateOrdersInput) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


