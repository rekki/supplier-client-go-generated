# \OrdersApi

All URIs are relative to *https://api.rekki.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmOrders**](OrdersApi.md#ConfirmOrders) | **Post** /integration/v1/orders/confirm | Confirm a pending order by its reference code.
[**ConfirmOrdersV3**](OrdersApi.md#ConfirmOrdersV3) | **Post** /integration/v3/orders/confirm | Confirm a pending order by its reference code.
[**ListNotIntegratedOrders**](OrdersApi.md#ListNotIntegratedOrders) | **Post** /integration/v1/orders/list_not_integrated | Lists all orders placed for the supplier that were placed through REKKI and not marked as integrated.
[**ListOrdersBySupplier**](OrdersApi.md#ListOrdersBySupplier) | **Post** /integration/v1/orders/list | Lists all orders placed for the supplier that were placed through REKKI.
[**ListOrdersBySupplierV3**](OrdersApi.md#ListOrdersBySupplierV3) | **Post** /integration/v3/orders/list | Lists all orders placed for the supplier that were placed through REKKI.
[**MarkIntegrationError**](OrdersApi.md#MarkIntegrationError) | **Post** /integration/v1/orders/set_error | Report failure to integrate an order
[**MarkIntegrationErrorV3**](OrdersApi.md#MarkIntegrationErrorV3) | **Post** /integration/v3/orders/set_error | Report failure to integrate an order
[**MarkOrdersIntegrated**](OrdersApi.md#MarkOrdersIntegrated) | **Post** /integration/v1/orders/set_integrated | Mark orders as integrated
[**MarkOrdersIntegratedV3**](OrdersApi.md#MarkOrdersIntegratedV3) | **Post** /integration/v3/orders/set_integrated | Mark orders as integrated



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

