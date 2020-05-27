# MainApiSupplierCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allergens** | [**[]CatalogAllergen**](catalog.Allergen.md) | List of allergens for the item, if any. | [optional] 
**Availability** | **string** | Availability status of the item. Defaults to in_stock. | [optional] 
**Currency** | **string** | Currency code for the price. In ISO 4217 three-letter format. Defaults to GBP. | [optional] 
**Description** | **string** | Product description | [optional] 
**Id** | **int32** | REKKI&#39;s ID to uniquely identify the catalog item (for REKKI internal reference). | [optional] 
**Name** | **string** | Item name as would be defined on the customer&#39;s product list. | [optional] 
**OrderCutoffTimes** | [**CatalogOrderCutoffTimes**](catalog.OrderCutoffTimes.md) |  | [optional] 
**ProductCode** | **string** | Product code for the item that maps to the supplier&#39;s catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com | [optional] 
**ReplacementProducts** | **[]string** | List of product codes for alternative items when this item is not available. | [optional] 
**Seasonality** | [**[]CatalogSeasonality**](catalog.Seasonality.md) | List of date ranges when the item is in-season. | [optional] 
**UnitsPrices** | [**[]MainApiSupplierCatalogUnit**](main.APISupplierCatalogUnit.md) | List of units and their prices that the item can be ordered in. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


