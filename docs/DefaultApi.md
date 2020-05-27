# \DefaultApi

All URIs are relative to *https://external-supplier-api.feat.eu-west-2.rekki.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmOrders**](DefaultApi.md#ConfirmOrders) | **Post** /integration/v1/orders/confirm | Confirm a pending order by its reference code.
[**ConfirmOrdersV3**](DefaultApi.md#ConfirmOrdersV3) | **Post** /integration/v3/orders/confirm | Confirm a pending order by its reference code.
[**DeleteCatalogItem**](DefaultApi.md#DeleteCatalogItem) | **Delete** /integration/v1/catalog/items/{id} | Delete an item from catalog
[**DeleteCatalogItemsV3**](DefaultApi.md#DeleteCatalogItemsV3) | **Post** /integration/v3/catalog/items/delete | Delete items from catalog
[**GetCatalogItem**](DefaultApi.md#GetCatalogItem) | **Get** /integration/v1/catalog/items/{id} | Lists all orders placed for the supplier that were placed through REKKI.
[**GetCatalogItemV3**](DefaultApi.md#GetCatalogItemV3) | **Get** /integration/v3/catalog/items/{id} | Fetch a specific catalog item by its Id.
[**GetCatalogItems**](DefaultApi.md#GetCatalogItems) | **Get** /integration/v1/catalog/items | Get catalog items for authenticated supplier
[**GetCatalogItemsV3**](DefaultApi.md#GetCatalogItemsV3) | **Get** /integration/v3/catalog/items | Get catalog items for authenticated supplier
[**ListNotIntegratedOrders**](DefaultApi.md#ListNotIntegratedOrders) | **Post** /integration/v1/orders/list_not_integrated | Lists all orders placed for the supplier that were placed through REKKI and not marked as integrated.
[**ListOrdersBySupplier**](DefaultApi.md#ListOrdersBySupplier) | **Post** /integration/v1/orders/list | Lists all orders placed for the supplier that were placed through REKKI.
[**ListOrdersBySupplierV3**](DefaultApi.md#ListOrdersBySupplierV3) | **Post** /integration/v3/orders/list | Lists all orders placed for the supplier that were placed through REKKI.
[**MarkIntegrationError**](DefaultApi.md#MarkIntegrationError) | **Post** /integration/v1/orders/set_error | Report failure to integrate an order
[**MarkIntegrationErrorV3**](DefaultApi.md#MarkIntegrationErrorV3) | **Post** /integration/v3/orders/set_error | Report failure to integrate an order
[**MarkOrdersIntegrated**](DefaultApi.md#MarkOrdersIntegrated) | **Post** /integration/v1/orders/set_integrated | Mark orders as integrated
[**MarkOrdersIntegratedV3**](DefaultApi.md#MarkOrdersIntegratedV3) | **Post** /integration/v3/orders/set_integrated | Mark orders as integrated
[**PostLogMessage**](DefaultApi.md#PostLogMessage) | **Post** /integration/v1/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint
[**PostLogMessageV3**](DefaultApi.md#PostLogMessageV3) | **Post** /integration/v3/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint
[**ReplaceCatalog**](DefaultApi.md#ReplaceCatalog) | **Post** /integration/v1/catalog/replace | Drop all existing items from the catalog and upload new ones
[**ReplaceCatalogV3**](DefaultApi.md#ReplaceCatalogV3) | **Post** /integration/v3/catalog/replace | Drop all existing items from the catalog and upload new ones
[**UpdateCatalogItem**](DefaultApi.md#UpdateCatalogItem) | **Post** /integration/v1/catalog/items | Creates or Updates an item on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code
[**UpdateCatalogItemAvailability**](DefaultApi.md#UpdateCatalogItemAvailability) | **Post** /integration/v2/catalog/items/availability | Update availability status for one of the items in the catalog
[**UpdateCatalogItemAvailabilityV3**](DefaultApi.md#UpdateCatalogItemAvailabilityV3) | **Post** /integration/v3/catalog/items/availability | Update availability status for one of the items in the catalog
[**UpdateCatalogItemsV3**](DefaultApi.md#UpdateCatalogItemsV3) | **Post** /integration/v3/catalog/items | Creates or Updates multiple items on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code



## ConfirmOrders

> MainSetIntegrateOrdersInput ConfirmOrders(ctx, xREKKIAuthorizationType, input)

Confirm a pending order by its reference code.

Notifies the buyer that the order has been acknowledged.  Status:` 200 OK` Body: `{ success: true}`  Status: `400 Conflict` Body: `{\"error\":\"Order already confirmed\",\"order_id\":...}`  Status: `400 Not Found` Body: `{\"error\":\"Order not found\",\"order_id\":...}`  in errors order_id denotes the order that failed to be confirmed  **the processing stops at first error** 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainSetIntegrateOrdersInput**](MainSetIntegrateOrdersInput.md)| Payload | 

### Return type

[**MainSetIntegrateOrdersInput**](main.SetIntegrateOrdersInput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmOrdersV3

> V3ConfirmOrdersInput ConfirmOrdersV3(ctx, xREKKIAuthorizationType, input)

Confirm a pending order by its reference code.

Notifies the buyer that the order has been acknowledged.  Status:` 200 OK` Body: `{ success: true}`  Status: `400 Conflict` Body: `{\"error\":\"Order already confirmed\",\"order_id\":...}`  Status: `400 Not Found` Body: `{\"error\":\"Order not found\",\"order_id\":...}`  in errors order_id denotes the order that failed to be confirmed  **the processing stops at first error** 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3ConfirmOrdersInput**](V3ConfirmOrdersInput.md)| Payload | 

### Return type

[**V3ConfirmOrdersInput**](v3.ConfirmOrdersInput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListNotIntegratedOrders

> MainOrderListOutput ListNotIntegratedOrders(ctx, xREKKIAuthorizationType, input)

Lists all orders placed for the supplier that were placed through REKKI and not marked as integrated.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainOrderListInput**](MainOrderListInput.md)| Payload | 

### Return type

[**MainOrderListOutput**](main.OrderListOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrdersBySupplier

> MainOrderListOutput ListOrdersBySupplier(ctx, xREKKIAuthorizationType, input)

Lists all orders placed for the supplier that were placed through REKKI.

Orders are limited to max 30 days old (i.e. timestamp must be within 30 days).  We recommend polling for orders by setting the new request timestamp to the time of the last successful request.  After you start using the API, you should request orders since last received order's inserted_at_ts, since the API returns orders created >= of the requested timestamp, you will always get at order from which you took the timestamp in the response. This will be explained again in the provided example.  Keep in mind that since you can have more than one order per since, you must not do since: last_order.inserted_at_ts + 1, but keep the last order you received's reference and ignore the duplicate.  After you start using the API, you should request orders since last received order's inserted_at_ts, since the API returns orders created >= of the requested timestamp, you will **always** get at order from which you took the timestamp in the response. This will be explained again in the provided example.  Keep in mind that since you can have more than one order per since, you must not do since: last_order.inserted_at_ts + 1, but keep the last order you received's reference and ignore the duplicate.  ## Exaple usage  In this JavaScript example, all orders are retrieved. Then it keeps pulling for new orders since the last order, every hour.    ```   const fetch = require(\"node-fetch\");    const sleep = function sleep(ms) {     return new Promise(resolve => setTimeout(resolve, ms));   };    const fetch_orders = async function(token, since) {     let r = await fetch(       \"https://api.rekki.com/api/catalog/integration/list_orders_by_supplier\",       {         method: \"POST\",         headers: {           Authorization: \"Bearer \" + token,           \"X-REKKI-Authorization-Type\": \"supplier_api_token\",           \"Content-Type\": \"application/json\",           Accept: \"application/json\"         },         body: JSON.stringify({ since })       }     );     return await r.json();   };    const poll = async function(token, last_rekki_order_time) {     let last_order_reference = undefined;      while (true) {       console.log(\"requesting orders since \" + last_rekki_order_time);       let response = await fetch_orders(token, last_rekki_order_time);        for (let order of response.orders) {         if (order.reference == last_order_reference) {           // here is where we are ignoring the order we           // took the inserted_at_ts from           // but since we can have more orders in the same inserted_at_ts           // you can't just do since: inserted_at_ts+1           continue;         }         if (order.inserted_at_ts >= last_rekki_order_time) {           last_rekki_order_time = order.inserted_at_ts;           last_order_reference = order.reference;         }          // process(order)         console.log(order);       }       await sleep(3600 * 1000); // wait 1 hour     }   };    poll(\"XXXXXXX-XXXX-XXXX-XXXXX-XXXXXXXXXXXX\", parseInt((+new Date() /1000) - 3600 * 24 * 30));   ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainOrderListInput**](MainOrderListInput.md)| Payload | 

### Return type

[**MainOrderListOutput**](main.OrderListOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrdersBySupplierV3

> V3OrderListOutput ListOrdersBySupplierV3(ctx, xREKKIAuthorizationType, input)

Lists all orders placed for the supplier that were placed through REKKI.

**UPDATE TO REFLECT V3 CHANGES**  Orders are limited to max 30 days old (i.e. timestamp must be within 30 days).  We recommend polling for orders by setting the new request timestamp to the time of the last successful request.  After you start using the API, you should request orders since last received order's inserted_at_ts, since the API returns orders created >= of the requested timestamp, you will always get at order from which you took the timestamp in the response. This will be explained again in the provided example.  Keep in mind that since you can have more than one order per since, you must not do since: last_order.inserted_at_ts + 1, but keep the last order you received's reference and ignore the duplicate.  After you start using the API, you should request orders since last received order's inserted_at_ts, since the API returns orders created >= of the requested timestamp, you will **always** get at order from which you took the timestamp in the response. This will be explained again in the provided example.  Keep in mind that since you can have more than one order per since, you must not do since: last_order.inserted_at_ts + 1, but keep the last order you received's reference and ignore the duplicate.  ## Exaple usage  In this JavaScript example, all orders are retrieved. Then it keeps pulling for new orders since the last order, every hour.    ```   const fetch = require(\"node-fetch\");    const sleep = function sleep(ms) {     return new Promise(resolve => setTimeout(resolve, ms));   };    const fetch_orders = async function(token, since) {     let r = await fetch(       \"https://api.rekki.com/api/catalog/integration/list_orders_by_supplier\",       {         method: \"POST\",         headers: {           Authorization: \"Bearer \" + token,           \"X-REKKI-Authorization-Type\": \"supplier_api_token\",           \"Content-Type\": \"application/json\",           Accept: \"application/json\"         },         body: JSON.stringify({ since })       }     );     return await r.json();   };    const poll = async function(token, last_rekki_order_time) {     let last_order_reference = undefined;      while (true) {       console.log(\"requesting orders since \" + last_rekki_order_time);       let response = await fetch_orders(token, last_rekki_order_time.toISOString());        for (let order of response.orders) {         if (order.reference == last_order_reference) {           // here is where we are ignoring the order we           // took the inserted_at from           // but since we can have more orders in the same inserted_at           // you can't just do since: inserted_at + 1 second           continue;         }         if (+new Date(order.inserted_at) >= +last_rekki_order_time) {           last_rekki_order_time = order.inserted_at;           last_order_reference = order.reference;         }          // process(order)         console.log(order);       }       await sleep(3600 * 1000); // wait 1 hour     }   };    let startDate = new Date()   startDate.setDate(startDate.getDate() - 30)   poll(\"XXXXXXX-XXXX-XXXX-XXXXX-XXXXXXXXXXXX\", startDate)   ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3OrderListInput**](V3OrderListInput.md)| Payload | 

### Return type

[**V3OrderListOutput**](v3.OrderListOutput.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkIntegrationError

> MainSuccessConfirmation MarkIntegrationError(ctx, xREKKIAuthorizationType, input)

Report failure to integrate an order

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainSetErrorOrderInput**](MainSetErrorOrderInput.md)| Payload | 

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


## MarkIntegrationErrorV3

> V3SuccessConfirmation MarkIntegrationErrorV3(ctx, xREKKIAuthorizationType, input)

Report failure to integrate an order

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3SetErrorOrderInput**](V3SetErrorOrderInput.md)| Payload | 

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


## MarkOrdersIntegrated

> MainUpdateSuccess MarkOrdersIntegrated(ctx, xREKKIAuthorizationType, input)

Mark orders as integrated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainSetIntegrateOrdersInput**](MainSetIntegrateOrdersInput.md)| Payload | 

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


## MarkOrdersIntegratedV3

> V3UpdateSuccess MarkOrdersIntegratedV3(ctx, xREKKIAuthorizationType, input)

Mark orders as integrated

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3SetIntegratedOrdersInput**](V3SetIntegratedOrdersInput.md)| Payload | 

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


## PostLogMessage

> string PostLogMessage(ctx, xREKKIAuthorizationType, input)

Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**MainLogMessage**](MainLogMessage.md)| Payload | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLogMessageV3

> string PostLogMessageV3(ctx, xREKKIAuthorizationType, input)

Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xREKKIAuthorizationType** | **string**| Required header | 
**input** | [**V3LogMessage**](V3LogMessage.md)| Payload | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
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

