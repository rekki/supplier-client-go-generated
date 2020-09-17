# V3AvailabilityError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**V3StockItemState**](v3.StockItemState.md) |  | [optional] 

## Methods

### NewV3AvailabilityError

`func NewV3AvailabilityError() *V3AvailabilityError`

NewV3AvailabilityError instantiates a new V3AvailabilityError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3AvailabilityErrorWithDefaults

`func NewV3AvailabilityErrorWithDefaults() *V3AvailabilityError`

NewV3AvailabilityErrorWithDefaults instantiates a new V3AvailabilityError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V3AvailabilityError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V3AvailabilityError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V3AvailabilityError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V3AvailabilityError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetItem

`func (o *V3AvailabilityError) GetItem() V3StockItemState`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *V3AvailabilityError) GetItemOk() (*V3StockItemState, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *V3AvailabilityError) SetItem(v V3StockItemState)`

SetItem sets Item field to given value.

### HasItem

`func (o *V3AvailabilityError) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


