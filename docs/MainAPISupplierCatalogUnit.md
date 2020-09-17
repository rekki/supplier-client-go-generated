# MainAPISupplierCatalogUnit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceCents** | Pointer to **int32** | The order price in cents for the item per unit. For example, a currency of GBP with unit 5L and price 850 means a 5L item can be ordered for £8.50. | [optional] 
**StockCount** | Pointer to **int32** | The number of items in stock for the related unit. | [optional] 
**Unit** | Pointer to **string** | A unit that the item can be ordered in. | [optional] 

## Methods

### NewMainAPISupplierCatalogUnit

`func NewMainAPISupplierCatalogUnit() *MainAPISupplierCatalogUnit`

NewMainAPISupplierCatalogUnit instantiates a new MainAPISupplierCatalogUnit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainAPISupplierCatalogUnitWithDefaults

`func NewMainAPISupplierCatalogUnitWithDefaults() *MainAPISupplierCatalogUnit`

NewMainAPISupplierCatalogUnitWithDefaults instantiates a new MainAPISupplierCatalogUnit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceCents

`func (o *MainAPISupplierCatalogUnit) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *MainAPISupplierCatalogUnit) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *MainAPISupplierCatalogUnit) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *MainAPISupplierCatalogUnit) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### GetStockCount

`func (o *MainAPISupplierCatalogUnit) GetStockCount() int32`

GetStockCount returns the StockCount field if non-nil, zero value otherwise.

### GetStockCountOk

`func (o *MainAPISupplierCatalogUnit) GetStockCountOk() (*int32, bool)`

GetStockCountOk returns a tuple with the StockCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockCount

`func (o *MainAPISupplierCatalogUnit) SetStockCount(v int32)`

SetStockCount sets StockCount field to given value.

### HasStockCount

`func (o *MainAPISupplierCatalogUnit) HasStockCount() bool`

HasStockCount returns a boolean if a field has been set.

### GetUnit

`func (o *MainAPISupplierCatalogUnit) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MainAPISupplierCatalogUnit) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MainAPISupplierCatalogUnit) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MainAPISupplierCatalogUnit) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


