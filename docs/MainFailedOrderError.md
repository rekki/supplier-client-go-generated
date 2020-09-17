# MainFailedOrderError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMainFailedOrderError

`func NewMainFailedOrderError() *MainFailedOrderError`

NewMainFailedOrderError instantiates a new MainFailedOrderError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainFailedOrderErrorWithDefaults

`func NewMainFailedOrderErrorWithDefaults() *MainFailedOrderError`

NewMainFailedOrderErrorWithDefaults instantiates a new MainFailedOrderError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MainFailedOrderError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MainFailedOrderError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MainFailedOrderError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MainFailedOrderError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOrderId

`func (o *MainFailedOrderError) GetOrderId() map[string]interface{}`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *MainFailedOrderError) GetOrderIdOk() (*map[string]interface{}, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *MainFailedOrderError) SetOrderId(v map[string]interface{})`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *MainFailedOrderError) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


