# MainOrderListOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]IntegrationOrder**](integration.Order.md) | list of not integrated orders | [optional] 

## Methods

### NewMainOrderListOutput

`func NewMainOrderListOutput() *MainOrderListOutput`

NewMainOrderListOutput instantiates a new MainOrderListOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainOrderListOutputWithDefaults

`func NewMainOrderListOutputWithDefaults() *MainOrderListOutput`

NewMainOrderListOutputWithDefaults instantiates a new MainOrderListOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *MainOrderListOutput) GetOrders() []IntegrationOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *MainOrderListOutput) GetOrdersOk() (*[]IntegrationOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *MainOrderListOutput) SetOrders(v []IntegrationOrder)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *MainOrderListOutput) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


