# MainStockItemState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | **string** | only possible values are &#x60;in_stock&#x60; or &#x60;out_of_stock&#x60; | 
**ProductCode** | **string** |  | 

## Methods

### NewMainStockItemState

`func NewMainStockItemState(availability string, productCode string, ) *MainStockItemState`

NewMainStockItemState instantiates a new MainStockItemState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainStockItemStateWithDefaults

`func NewMainStockItemStateWithDefaults() *MainStockItemState`

NewMainStockItemStateWithDefaults instantiates a new MainStockItemState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *MainStockItemState) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MainStockItemState) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *MainStockItemState) SetAvailability(v string)`

SetAvailability sets Availability field to given value.


### GetProductCode

`func (o *MainStockItemState) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *MainStockItemState) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *MainStockItemState) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


