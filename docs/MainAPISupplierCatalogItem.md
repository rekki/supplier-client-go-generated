# MainAPISupplierCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allergens** | Pointer to [**[]CatalogAllergen**](catalog.Allergen.md) | List of allergens for the item, if any. | [optional] 
**Availability** | Pointer to **string** | Availability status of the item. Defaults to in_stock. | [optional] 
**Currency** | Pointer to **string** | Currency code for the price. In ISO 4217 three-letter format. Defaults to GBP. | [optional] 
**Description** | Pointer to **string** | Product description | [optional] 
**Id** | Pointer to **int32** | REKKI&#39;s ID to uniquely identify the catalog item (for REKKI internal reference). | [optional] 
**Name** | Pointer to **string** | Item name as would be defined on the customer&#39;s product list. | [optional] 
**OrderCutoffTimes** | Pointer to [**CatalogOrderCutoffTimes**](catalog.OrderCutoffTimes.md) |  | [optional] 
**ProductCode** | Pointer to **string** | Product code for the item that maps to the supplier&#39;s catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com | [optional] 
**ReplacementProducts** | Pointer to **[]string** | List of product codes for alternative items when this item is not available. | [optional] 
**Seasonality** | Pointer to [**[]CatalogSeasonality**](catalog.Seasonality.md) | List of date ranges when the item is in-season. | [optional] 
**UnitsPrices** | Pointer to [**[]MainAPISupplierCatalogUnit**](main.APISupplierCatalogUnit.md) | List of units and their prices that the item can be ordered in. | [optional] 

## Methods

### NewMainAPISupplierCatalogItem

`func NewMainAPISupplierCatalogItem() *MainAPISupplierCatalogItem`

NewMainAPISupplierCatalogItem instantiates a new MainAPISupplierCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMainAPISupplierCatalogItemWithDefaults

`func NewMainAPISupplierCatalogItemWithDefaults() *MainAPISupplierCatalogItem`

NewMainAPISupplierCatalogItemWithDefaults instantiates a new MainAPISupplierCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllergens

`func (o *MainAPISupplierCatalogItem) GetAllergens() []CatalogAllergen`

GetAllergens returns the Allergens field if non-nil, zero value otherwise.

### GetAllergensOk

`func (o *MainAPISupplierCatalogItem) GetAllergensOk() (*[]CatalogAllergen, bool)`

GetAllergensOk returns a tuple with the Allergens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllergens

`func (o *MainAPISupplierCatalogItem) SetAllergens(v []CatalogAllergen)`

SetAllergens sets Allergens field to given value.

### HasAllergens

`func (o *MainAPISupplierCatalogItem) HasAllergens() bool`

HasAllergens returns a boolean if a field has been set.

### GetAvailability

`func (o *MainAPISupplierCatalogItem) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MainAPISupplierCatalogItem) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *MainAPISupplierCatalogItem) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *MainAPISupplierCatalogItem) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetCurrency

`func (o *MainAPISupplierCatalogItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MainAPISupplierCatalogItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MainAPISupplierCatalogItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *MainAPISupplierCatalogItem) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *MainAPISupplierCatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MainAPISupplierCatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MainAPISupplierCatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MainAPISupplierCatalogItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *MainAPISupplierCatalogItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MainAPISupplierCatalogItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MainAPISupplierCatalogItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MainAPISupplierCatalogItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MainAPISupplierCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MainAPISupplierCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MainAPISupplierCatalogItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MainAPISupplierCatalogItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrderCutoffTimes

`func (o *MainAPISupplierCatalogItem) GetOrderCutoffTimes() CatalogOrderCutoffTimes`

GetOrderCutoffTimes returns the OrderCutoffTimes field if non-nil, zero value otherwise.

### GetOrderCutoffTimesOk

`func (o *MainAPISupplierCatalogItem) GetOrderCutoffTimesOk() (*CatalogOrderCutoffTimes, bool)`

GetOrderCutoffTimesOk returns a tuple with the OrderCutoffTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCutoffTimes

`func (o *MainAPISupplierCatalogItem) SetOrderCutoffTimes(v CatalogOrderCutoffTimes)`

SetOrderCutoffTimes sets OrderCutoffTimes field to given value.

### HasOrderCutoffTimes

`func (o *MainAPISupplierCatalogItem) HasOrderCutoffTimes() bool`

HasOrderCutoffTimes returns a boolean if a field has been set.

### GetProductCode

`func (o *MainAPISupplierCatalogItem) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *MainAPISupplierCatalogItem) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *MainAPISupplierCatalogItem) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.

### HasProductCode

`func (o *MainAPISupplierCatalogItem) HasProductCode() bool`

HasProductCode returns a boolean if a field has been set.

### GetReplacementProducts

`func (o *MainAPISupplierCatalogItem) GetReplacementProducts() []string`

GetReplacementProducts returns the ReplacementProducts field if non-nil, zero value otherwise.

### GetReplacementProductsOk

`func (o *MainAPISupplierCatalogItem) GetReplacementProductsOk() (*[]string, bool)`

GetReplacementProductsOk returns a tuple with the ReplacementProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementProducts

`func (o *MainAPISupplierCatalogItem) SetReplacementProducts(v []string)`

SetReplacementProducts sets ReplacementProducts field to given value.

### HasReplacementProducts

`func (o *MainAPISupplierCatalogItem) HasReplacementProducts() bool`

HasReplacementProducts returns a boolean if a field has been set.

### GetSeasonality

`func (o *MainAPISupplierCatalogItem) GetSeasonality() []CatalogSeasonality`

GetSeasonality returns the Seasonality field if non-nil, zero value otherwise.

### GetSeasonalityOk

`func (o *MainAPISupplierCatalogItem) GetSeasonalityOk() (*[]CatalogSeasonality, bool)`

GetSeasonalityOk returns a tuple with the Seasonality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonality

`func (o *MainAPISupplierCatalogItem) SetSeasonality(v []CatalogSeasonality)`

SetSeasonality sets Seasonality field to given value.

### HasSeasonality

`func (o *MainAPISupplierCatalogItem) HasSeasonality() bool`

HasSeasonality returns a boolean if a field has been set.

### GetUnitsPrices

`func (o *MainAPISupplierCatalogItem) GetUnitsPrices() []MainAPISupplierCatalogUnit`

GetUnitsPrices returns the UnitsPrices field if non-nil, zero value otherwise.

### GetUnitsPricesOk

`func (o *MainAPISupplierCatalogItem) GetUnitsPricesOk() (*[]MainAPISupplierCatalogUnit, bool)`

GetUnitsPricesOk returns a tuple with the UnitsPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsPrices

`func (o *MainAPISupplierCatalogItem) SetUnitsPrices(v []MainAPISupplierCatalogUnit)`

SetUnitsPrices sets UnitsPrices field to given value.

### HasUnitsPrices

`func (o *MainAPISupplierCatalogItem) HasUnitsPrices() bool`

HasUnitsPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


