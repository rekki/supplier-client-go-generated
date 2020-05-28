# \CatalogApi

All URIs are relative to *https://api.rekki.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCatalogItem**](CatalogApi.md#DeleteCatalogItem) | **Delete** /integration/v1/catalog/items/{id} | Delete an item from catalog
[**DeleteCatalogItemsV3**](CatalogApi.md#DeleteCatalogItemsV3) | **Post** /integration/v3/catalog/items/delete | Delete items from catalog
[**GetCatalogItem**](CatalogApi.md#GetCatalogItem) | **Get** /integration/v1/catalog/items/{id} | Lists all orders placed for the supplier that were placed through REKKI.
[**GetCatalogItemV3**](CatalogApi.md#GetCatalogItemV3) | **Get** /integration/v3/catalog/items/{id} | Fetch a specific catalog item by its Id.
[**GetCatalogItems**](CatalogApi.md#GetCatalogItems) | **Get** /integration/v1/catalog/items | Get catalog items for authenticated supplier
[**GetCatalogItemsV3**](CatalogApi.md#GetCatalogItemsV3) | **Get** /integration/v3/catalog/items | Get catalog items for authenticated supplier
[**ReplaceCatalog**](CatalogApi.md#ReplaceCatalog) | **Post** /integration/v1/catalog/replace | Drop all existing items from the catalog and upload new ones
[**ReplaceCatalogV3**](CatalogApi.md#ReplaceCatalogV3) | **Post** /integration/v3/catalog/replace | Drop all existing items from the catalog and upload new ones
[**UpdateCatalogItem**](CatalogApi.md#UpdateCatalogItem) | **Post** /integration/v1/catalog/items | Creates or Updates an item on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code
[**UpdateCatalogItemAvailability**](CatalogApi.md#UpdateCatalogItemAvailability) | **Post** /integration/v2/catalog/items/availability | Update availability status for one of the items in the catalog
[**UpdateCatalogItemAvailabilityV3**](CatalogApi.md#UpdateCatalogItemAvailabilityV3) | **Post** /integration/v3/catalog/items/availability | Update availability status for one of the items in the catalog
[**UpdateCatalogItemsV3**](CatalogApi.md#UpdateCatalogItemsV3) | **Post** /integration/v3/catalog/items | Creates or Updates multiple items on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code



## DeleteCatalogItem

> MainSuccessConfirmation DeleteCatalogItem(ctx, xREKKIAuthorizationType, id)

Delete an item from catalog

Delete an item from your catalog by its unique ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**id** | **string**| ID of the item to retrieve. Item IDs are discoverable when listing items. | 

### Return type

[**MainSuccessConfirmation**](main.SuccessConfirmation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCatalogItemsV3

> V3SuccessConfirmation DeleteCatalogItemsV3(ctx, xREKKIAuthorizationType, input)

Delete items from catalog

Delete items from the catalog by its unique IDs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3DeleteCatalogItemsInput**](V3DeleteCatalogItemsInput.md)| Payload | 

### Return type

[**V3SuccessConfirmation**](v3.SuccessConfirmation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItem

> MainApiSupplierCatalogItem GetCatalogItem(ctx, xREKKIAuthorizationType, id)

Lists all orders placed for the supplier that were placed through REKKI.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**id** | **string**| ID of the item to retrieve. Item IDs are discoverable when listing items. | 

### Return type

[**MainApiSupplierCatalogItem**](main.APISupplierCatalogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemV3

> V3ApiSupplierCatalogItem GetCatalogItemV3(ctx, xREKKIAuthorizationType, id)

Fetch a specific catalog item by its Id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**id** | **string**| ID of the item to retrieve. Item IDs are discoverable when listing items. | 

### Return type

[**V3ApiSupplierCatalogItem**](v3.APISupplierCatalogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItems

> MainCatalogItems GetCatalogItems(ctx, xREKKIAuthorizationType)

Get catalog items for authenticated supplier

Lists all your catalog items. Sorted by creation date, with the oldest appearing first. No input parameters. Options for pagination and sorting direction may be introduced later. Response is a JSON object with a data property that contains catalog items of the authenticated supplier.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 

### Return type

[**MainCatalogItems**](main.CatalogItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemsV3

> V3CatalogItems GetCatalogItemsV3(ctx, xREKKIAuthorizationType)

Get catalog items for authenticated supplier

Lists all your catalog items. Sorted by creation date, with the oldest appearing first. No input parameters. Options for pagination and sorting direction may be introduced later. Response is a JSON object with a data property that contains catalog items of the authenticated supplier.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 

### Return type

[**V3CatalogItems**](v3.CatalogItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCatalog

> MainSuccessConfirmation ReplaceCatalog(ctx, xREKKIAuthorizationType, input)

Drop all existing items from the catalog and upload new ones

### Parameters  - **`id`** REKKI's ID to uniquely identify the catalog item (for REKKI internal reference). If `id` is specified the item will be update, if not it will attempt to insert it. - **`product_code`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Product code for the item that maps to the supplier's catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com - **`name`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Item name as would be defined on the customer's product list. - **`currency`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is GBP</span> Currency code for the price. In [ISO 4217][] three-letter format. - **`units_prices`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> List of units and their prices that the item can be ordered in. - **`units_prices.unit`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> A unit that the item can be ordered in. - **`units_prices.price_cents`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is 0</span> The order price in cents for the item per unit. For example, a currency of GBP with unit 5L and price 850 means a 5L item can be ordered for £8.50. - **`units_prices.stock_count`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> The number of items in stock for the related unit. - **`availability`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is \"in_stock\"</span> Availability status of the item. Can be \"in_stock\", \"out_of_stock\", or \"discontinued\". - **`description`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Short description of the item. - **`allergens`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of allergens for the item, if any. - **`allergens.type`**  <span style=\"font-size: 12px; font-weight: 500;\">required when allergens is given</span> Type of allergy. For example \"contains peanuts\" or \"may contain peanuts\". - **`allergens.symptoms`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of symptoms for the allergy. - **`order_cutoff_times`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Cutt-off times are the minimum amount of time before delivery when the item can still be ordered. - **`order_cutoff_times.{mon,tue,wed,thu,fri,sat,sun}`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day. - **`replacement_products`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of product codes for alternative items when this item is not available. - **`seasonality`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of date ranges when the item is in-season. - **`seasonality.start_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>   The start date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`.   - **`seasonality.end_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>     The end date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainReplaceCatalogInput**](MainReplaceCatalogInput.md)| Payload | 

### Return type

[**MainSuccessConfirmation**](main.SuccessConfirmation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCatalogV3

> V3SuccessConfirmation ReplaceCatalogV3(ctx, xREKKIAuthorizationType, input)

Drop all existing items from the catalog and upload new ones

### Parameters  - **`id`** REKKI's ID to uniquely identify the catalog item (for REKKI internal reference). If `id` is specified the item will be update, if not it will attempt to insert it. - **`product_code`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Product code for the item that maps to the supplier's catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com - **`name`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Item name as would be defined on the customer's product list. - **`currency`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is GBP</span> Currency code for the price. In [ISO 4217][] three-letter format. - **`units_prices`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> List of units and their prices that the item can be ordered in. - **`units_prices.unit`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> A unit that the item can be ordered in. - **`units_prices.price_cents`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is 0</span> The order price in cents for the item per unit. For example, a currency of GBP with unit 5L and price 850 means a 5L item can be ordered for £8.50. - **`units_prices.stock_count`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> The number of items in stock for the related unit. - **`availability`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is \"in_stock\"</span> Availability status of the item. Can be \"in_stock\", \"out_of_stock\", or \"discontinued\". - **`description`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Short description of the item. - **`allergens`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of allergens for the item, if any. - **`allergens.type`**  <span style=\"font-size: 12px; font-weight: 500;\">required when allergens is given</span> Type of allergy. For example \"contains peanuts\" or \"may contain peanuts\". - **`allergens.symptoms`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of symptoms for the allergy. - **`order_cutoff_times`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Cutt-off times are the minimum amount of time before delivery when the item can still be ordered. - **`order_cutoff_times.{mon,tue,wed,thu,fri,sat,sun}`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day. - **`replacement_products`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of product codes for alternative items when this item is not available. - **`seasonality`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of date ranges when the item is in-season. - **`seasonality.start_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>   The start date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`.   - **`seasonality.end_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>     The end date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3ReplaceCatalogInput**](V3ReplaceCatalogInput.md)| Payload | 

### Return type

[**V3SuccessConfirmation**](v3.SuccessConfirmation.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItem

> MainApiSupplierCatalogItem UpdateCatalogItem(ctx, xREKKIAuthorizationType, input)

Creates or Updates an item on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code

### Parameters  - **`id`** REKKI's ID to uniquely identify the catalog item (for REKKI internal reference). If `id` is specified the item will be update, if not it will attempt to insert it. - **`product_code`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Product code for the item that maps to the supplier's catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com - **`name`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Item name as would be defined on the customer's product list. - **`currency`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is GBP</span> Currency code for the price. In [ISO 4217][] three-letter format. - **`units_prices`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> List of units and their prices that the item can be ordered in. - **`units_prices.unit`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> A unit that the item can be ordered in. - **`units_prices.price_cents`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is 0</span> The order price in cents for the item per unit. For example, a currency of GBP with unit 5L and price 850 means a 5L item can be ordered for £8.50. - **`units_prices.stock_count`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> The number of items in stock for the related unit. - **`availability`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is \"in_stock\"</span> Availability status of the item. Can be \"in_stock\", \"out_of_stock\", or \"discontinued\". - **`description`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Short description of the item. - **`allergens`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of allergens for the item, if any. - **`allergens.type`**  <span style=\"font-size: 12px; font-weight: 500;\">required when allergens is given</span> Type of allergy. For example \"contains peanuts\" or \"may contain peanuts\". - **`allergens.symptoms`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of symptoms for the allergy. - **`order_cutoff_times`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Cutt-off times are the minimum amount of time before delivery when the item can still be ordered. - **`order_cutoff_times.{mon,tue,wed,thu,fri,sat,sun}`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day. - **`replacement_products`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of product codes for alternative items when this item is not available. - **`seasonality`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of date ranges when the item is in-season. - **`seasonality.start_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>   The start date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`.   - **`seasonality.end_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>     The end date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainApiSupplierCatalogItem**](MainApiSupplierCatalogItem.md)| Payload | 

### Return type

[**MainApiSupplierCatalogItem**](main.APISupplierCatalogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItemAvailability

> MainUpdateSuccess UpdateCatalogItemAvailability(ctx, xREKKIAuthorizationType, input)

Update availability status for one of the items in the catalog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainSetStockStatusInput**](MainSetStockStatusInput.md)| Payload | 

### Return type

[**MainUpdateSuccess**](main.UpdateSuccess.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItemAvailabilityV3

> V3UpdateSuccess UpdateCatalogItemAvailabilityV3(ctx, xREKKIAuthorizationType, input)

Update availability status for one of the items in the catalog

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3SetStockStatusInput**](V3SetStockStatusInput.md)| Payload | 

### Return type

[**V3UpdateSuccess**](v3.UpdateSuccess.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItemsV3

> V3CatalogItems UpdateCatalogItemsV3(ctx, xREKKIAuthorizationType, input)

Creates or Updates multiple items on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code

### Parameters  - **`id`** REKKI's ID to uniquely identify the catalog item (for REKKI internal reference). If `id` is specified the item will be update, if not it will attempt to insert it. - **`product_code`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Product code for the item that maps to the supplier's catalog. Suppliers can modify the product code for future orders at https://supplier.rekki.com - **`name`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> Item name as would be defined on the customer's product list. - **`currency`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is GBP</span> Currency code for the price. In [ISO 4217][] three-letter format. - **`units_prices`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> List of units and their prices that the item can be ordered in. - **`units_prices.unit`**  <span style=\"font-size: 12px; font-weight: 500;\">required</span> A unit that the item can be ordered in. - **`units_prices.price_cents`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is 0</span> The order price in cents for the item per unit. For example, a currency of GBP with unit 5L and price 850 means a 5L item can be ordered for £8.50. - **`units_prices.stock_count`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> The number of items in stock for the related unit. - **`availability`**  <span style=\"font-size: 12px; font-weight: 500;\">optional, default is \"in_stock\"</span> Availability status of the item. Can be \"in_stock\", \"out_of_stock\", or \"discontinued\". - **`description`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Short description of the item. - **`allergens`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of allergens for the item, if any. - **`allergens.type`**  <span style=\"font-size: 12px; font-weight: 500;\">required when allergens is given</span> Type of allergy. For example \"contains peanuts\" or \"may contain peanuts\". - **`allergens.symptoms`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of symptoms for the allergy. - **`order_cutoff_times`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Cutt-off times are the minimum amount of time before delivery when the item can still be ordered. - **`order_cutoff_times.{mon,tue,wed,thu,fri,sat,sun}`**  <span style=\"font-size: 12px; font-weight: 500;\">optional</span> Minimum amount of time, in hours, that an item needs to be ordered in advance of delivery for the given day. - **`replacement_products`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of product codes for alternative items when this item is not available. - **`seasonality`**   <span style=\"font-size: 12px; font-weight: 500;\">optional</span> List of date ranges when the item is in-season. - **`seasonality.start_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>   The start date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`.   - **`seasonality.end_date`**  <span style=\"font-size: 12px; font-weight: 500;\">required when seasonality is given</span>     The end date when the item is in season. In [ISO 8601][] calendar date format `YYYY-MM-DD`. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3UpdateCatalogInput**](V3UpdateCatalogInput.md)| Payload | 

### Return type

[**V3CatalogItems**](v3.CatalogItems.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

