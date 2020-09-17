# V3StockItemState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | **string** | only possible values are &#x60;in_stock&#x60; or &#x60;out_of_stock&#x60; | 
**ProductCode** | **string** |  | 

## Methods

### NewV3StockItemState

`func NewV3StockItemState(availability string, productCode string, ) *V3StockItemState`

NewV3StockItemState instantiates a new V3StockItemState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3StockItemStateWithDefaults

`func NewV3StockItemStateWithDefaults() *V3StockItemState`

NewV3StockItemStateWithDefaults instantiates a new V3StockItemState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *V3StockItemState) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *V3StockItemState) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *V3StockItemState) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetProductCode

`func (o *V3StockItemState) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *V3StockItemState) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *V3StockItemState) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


