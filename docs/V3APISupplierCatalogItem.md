# V3APISupplierCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allergens** | Pointer to [**[]V3Allergens**](v3.Allergens.md) | List of allergens for the item, if any. | [optional] 
**Availability** | Pointer to **string** | Availability status of the item. Defaults to in_stock. | [optional] 
**Currency** | Pointer to **string** | Currency code for the price. In ISO 4217 three-letter format. Defaults to GBP. | [optional] 
**Description** | Pointer to **string** | Product description | [optional] 
**Id** | Pointer to **int32** | REKKI&#39;s ID to uniquely identify the catalog item (for REKKI internal reference). | [optional] 
**Name** | Pointer to **string** | Item name as would be defined on the customer&#39;s product list. | [optional] 
**OrderCutoffTimes** | Pointer to [**V3OrderCutoffTimes**](v3.OrderCutoffTimes.md) |  | [optional] 
**ProductCode** | Pointer to **string** | Product code for the item that maps to the supplier&#39;s catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com | [optional] 
**ReplacementProducts** | Pointer to **[]string** | List of product codes for alternative items when this item is not available. | [optional] 
**Seasonality** | Pointer to [**[]V3Seasonality**](v3.Seasonality.md) | List of date ranges when the item is in-season. | [optional] 
**UnitsPrices** | Pointer to [**[]V3APISupplierCatalogUnit**](v3.APISupplierCatalogUnit.md) | List of units and their prices that the item can be ordered in. | [optional] 

## Methods

### NewV3APISupplierCatalogItem

`func NewV3APISupplierCatalogItem() *V3APISupplierCatalogItem`

NewV3APISupplierCatalogItem instantiates a new V3APISupplierCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3APISupplierCatalogItemWithDefaults

`func NewV3APISupplierCatalogItemWithDefaults() *V3APISupplierCatalogItem`

NewV3APISupplierCatalogItemWithDefaults instantiates a new V3APISupplierCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllergens

`func (o *V3APISupplierCatalogItem) GetAllergens() []V3Allergens`

GetAllergens returns the Allergens field if non-nil, zero value otherwise.

### GetAllergensOk

`func (o *V3APISupplierCatalogItem) GetAllergensOk() (*[]V3Allergens, bool)`

GetAllergensOk returns a tuple with the Allergens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllergens

`func (o *V3APISupplierCatalogItem) SetAllergens(v []V3Allergens)`

SetAllergens sets Allergens field to given value.

### HasAllergens

`func (o *V3APISupplierCatalogItem) HasAllergens() bool`

HasAllergens returns a boolean if a field has been set.

### GetAvailability

`func (o *V3APISupplierCatalogItem) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *V3APISupplierCatalogItem) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *V3APISupplierCatalogItem) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *V3APISupplierCatalogItem) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCurrency

`func (o *V3APISupplierCatalogItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *V3APISupplierCatalogItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *V3APISupplierCatalogItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *V3APISupplierCatalogItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *V3APISupplierCatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V3APISupplierCatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V3APISupplierCatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V3APISupplierCatalogItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *V3APISupplierCatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V3APISupplierCatalogItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V3APISupplierCatalogItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V3APISupplierCatalogItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V3APISupplierCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V3APISupplierCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V3APISupplierCatalogItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V3APISupplierCatalogItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderCutoffTimes

`func (o *V3APISupplierCatalogItem) GetOrderCutoffTimes() V3OrderCutoffTimes`

GetOrderCutoffTimes returns the OrderCutoffTimes field if non-nil, zero value otherwise.

### GetOrderCutoffTimesOk

`func (o *V3APISupplierCatalogItem) GetOrderCutoffTimesOk() (*V3OrderCutoffTimes, bool)`

GetOrderCutoffTimesOk returns a tuple with the OrderCutoffTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCutoffTimes

`func (o *V3APISupplierCatalogItem) SetOrderCutoffTimes(v V3OrderCutoffTimes)`

SetOrderCutoffTimes sets OrderCutoffTimes field to given value.

### HasOrderCutoffTimes

`func (o *V3APISupplierCatalogItem) HasOrderCutoffTimes() bool`

HasOrderCutoffTimes returns a boolean if a field has been set.

### GetProductCode

`func (o *V3APISupplierCatalogItem) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *V3APISupplierCatalogItem) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *V3APISupplierCatalogItem) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *V3APISupplierCatalogItem) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetReplacementProducts

`func (o *V3APISupplierCatalogItem) GetReplacementProducts() []string`

GetReplacementProducts returns the ReplacementProducts field if non-nil, zero value otherwise.

### GetReplacementProductsOk

`func (o *V3APISupplierCatalogItem) GetReplacementProductsOk() (*[]string, bool)`

GetReplacementProductsOk returns a tuple with the ReplacementProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementProducts

`func (o *V3APISupplierCatalogItem) SetReplacementProducts(v []string)`

SetReplacementProducts sets ReplacementProducts field to given value.

### HasReplacementProducts

`func (o *V3APISupplierCatalogItem) HasReplacementProducts() bool`

HasReplacementProducts returns a boolean if a field has been set.

### GetSeasonality

`func (o *V3APISupplierCatalogItem) GetSeasonality() []V3Seasonality`

GetSeasonality returns the Seasonality field if non-nil, zero value otherwise.

### GetSeasonalityOk

`func (o *V3APISupplierCatalogItem) GetSeasonalityOk() (*[]V3Seasonality, bool)`

GetSeasonalityOk returns a tuple with the Seasonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonality

`func (o *V3APISupplierCatalogItem) SetSeasonality(v []V3Seasonality)`

SetSeasonality sets Seasonality field to given value.

### HasSeasonality

`func (o *V3APISupplierCatalogItem) HasSeasonality() bool`

HasSeasonality returns a boolean if a field has been set.

### GetUnitsPrices

`func (o *V3APISupplierCatalogItem) GetUnitsPrices() []V3APISupplierCatalogUnit`

GetUnitsPrices returns the UnitsPrices field if non-nil, zero value otherwise.

### GetUnitsPricesOk

`func (o *V3APISupplierCatalogItem) GetUnitsPricesOk() (*[]V3APISupplierCatalogUnit, bool)`

GetUnitsPricesOk returns a tuple with the UnitsPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsPrices

`func (o *V3APISupplierCatalogItem) SetUnitsPrices(v []V3APISupplierCatalogUnit)`

SetUnitsPrices sets UnitsPrices field to given value.

### HasUnitsPrices

`func (o *V3APISupplierCatalogItem) HasUnitsPrices() bool`

HasUnitsPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


