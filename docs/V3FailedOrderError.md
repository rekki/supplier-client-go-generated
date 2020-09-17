# V3FailedOrderError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 

## Methods

### NewV3FailedOrderError

`func NewV3FailedOrderError() *V3FailedOrderError`

NewV3FailedOrderError instantiates a new V3FailedOrderError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3FailedOrderErrorWithDefaults

`func NewV3FailedOrderErrorWithDefaults() *V3FailedOrderError`

NewV3FailedOrderErrorWithDefaults instantiates a new V3FailedOrderError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V3FailedOrderError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V3FailedOrderError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V3FailedOrderError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V3FailedOrderError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOrderId

`func (o *V3FailedOrderError) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *V3FailedOrderError) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *V3FailedOrderError) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *V3FailedOrderError) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


