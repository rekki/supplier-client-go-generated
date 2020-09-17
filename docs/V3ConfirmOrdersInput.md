# V3ConfirmOrdersInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to **[]string** | Array of References of the orders to confirm, required. Order refs are discoverable when listing orders. | [optional] 

## Methods

### NewV3ConfirmOrdersInput

`func NewV3ConfirmOrdersInput() *V3ConfirmOrdersInput`

NewV3ConfirmOrdersInput instantiates a new V3ConfirmOrdersInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3ConfirmOrdersInputWithDefaults

`func NewV3ConfirmOrdersInputWithDefaults() *V3ConfirmOrdersInput`

NewV3ConfirmOrdersInputWithDefaults instantiates a new V3ConfirmOrdersInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *V3ConfirmOrdersInput) GetOrders() []string`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *V3ConfirmOrdersInput) GetOrdersOk() (*[]string, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *V3ConfirmOrdersInput) SetOrders(v []string)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *V3ConfirmOrdersInput) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


