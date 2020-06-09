# Go API client for client

The base URL for all API endpoints is https://api.rekki.com

Api key value contains of word Bearer together with api key that you can get from integrations@rekki.com


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./client"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.rekki.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CatalogApi* | [**DeleteCatalogItem**](docs/CatalogApi.md#deletecatalogitem) | **Delete** /integration/v1/catalog/items/{id} | Delete an item from catalog
*CatalogApi* | [**DeleteCatalogItemsV3**](docs/CatalogApi.md#deletecatalogitemsv3) | **Post** /integration/v3/catalog/items/delete | Delete items from catalog
*CatalogApi* | [**GetCatalogItem**](docs/CatalogApi.md#getcatalogitem) | **Get** /integration/v1/catalog/items/{id} | Lists all orders placed for the supplier that were placed through REKKI.
*CatalogApi* | [**GetCatalogItemV3**](docs/CatalogApi.md#getcatalogitemv3) | **Get** /integration/v3/catalog/items/{id} | Fetch a specific catalog item by its Id.
*CatalogApi* | [**GetCatalogItems**](docs/CatalogApi.md#getcatalogitems) | **Get** /integration/v1/catalog/items | Get catalog items for authenticated supplier
*CatalogApi* | [**GetCatalogItemsV3**](docs/CatalogApi.md#getcatalogitemsv3) | **Get** /integration/v3/catalog/items | Get catalog items for authenticated supplier
*CatalogApi* | [**ReplaceCatalog**](docs/CatalogApi.md#replacecatalog) | **Post** /integration/v1/catalog/replace | Drop all existing items from the catalog and upload new ones
*CatalogApi* | [**ReplaceCatalogV3**](docs/CatalogApi.md#replacecatalogv3) | **Post** /integration/v3/catalog/replace | Drop all existing items from the catalog and upload new ones
*CatalogApi* | [**UpdateCatalogItem**](docs/CatalogApi.md#updatecatalogitem) | **Post** /integration/v1/catalog/items | Creates or Updates an item on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code
*CatalogApi* | [**UpdateCatalogItemAvailability**](docs/CatalogApi.md#updatecatalogitemavailability) | **Post** /integration/v2/catalog/items/availability | Update availability status for one of the items in the catalog
*CatalogApi* | [**UpdateCatalogItemAvailabilityV3**](docs/CatalogApi.md#updatecatalogitemavailabilityv3) | **Post** /integration/v3/catalog/items/availability | Update availability status for one of the items in the catalog
*CatalogApi* | [**UpdateCatalogItemsV3**](docs/CatalogApi.md#updatecatalogitemsv3) | **Post** /integration/v3/catalog/items | Creates or Updates multiple items on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code
*GeneralApi* | [**PostLogMessage**](docs/GeneralApi.md#postlogmessage) | **Post** /integration/v1/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint
*GeneralApi* | [**PostLogMessageV3**](docs/GeneralApi.md#postlogmessagev3) | **Post** /integration/v3/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint
*OrdersApi* | [**ConfirmOrders**](docs/OrdersApi.md#confirmorders) | **Post** /integration/v1/orders/confirm | Confirm a pending order by its reference code.
*OrdersApi* | [**ConfirmOrdersV3**](docs/OrdersApi.md#confirmordersv3) | **Post** /integration/v3/orders/confirm | Confirm a pending order by its reference code.
*OrdersApi* | [**ListNotIntegratedOrders**](docs/OrdersApi.md#listnotintegratedorders) | **Post** /integration/v1/orders/list_not_integrated | Lists all orders placed for the supplier that were placed through REKKI and not marked as integrated.
*OrdersApi* | [**ListOrdersBySupplier**](docs/OrdersApi.md#listordersbysupplier) | **Post** /integration/v1/orders/list | Lists all orders placed for the supplier that were placed through REKKI.
*OrdersApi* | [**ListOrdersBySupplierV3**](docs/OrdersApi.md#listordersbysupplierv3) | **Post** /integration/v3/orders/list | Lists all orders placed for the supplier that were placed through REKKI.
*OrdersApi* | [**MarkIntegrationError**](docs/OrdersApi.md#markintegrationerror) | **Post** /integration/v1/orders/set_error | Report failure to integrate an order
*OrdersApi* | [**MarkIntegrationErrorV3**](docs/OrdersApi.md#markintegrationerrorv3) | **Post** /integration/v3/orders/set_error | Report failure to integrate an order
*OrdersApi* | [**MarkOrdersIntegrated**](docs/OrdersApi.md#markordersintegrated) | **Post** /integration/v1/orders/set_integrated | Mark orders as integrated
*OrdersApi* | [**MarkOrdersIntegratedV3**](docs/OrdersApi.md#markordersintegratedv3) | **Post** /integration/v3/orders/set_integrated | Mark orders as integrated


## Documentation For Models

 - [CatalogAllergen](docs/CatalogAllergen.md)
 - [CatalogOrderCutoffTimes](docs/CatalogOrderCutoffTimes.md)
 - [CatalogSeasonality](docs/CatalogSeasonality.md)
 - [IntegrationOrder](docs/IntegrationOrder.md)
 - [IntegrationOrderItem](docs/IntegrationOrderItem.md)
 - [MainApiSupplierCatalogItem](docs/MainApiSupplierCatalogItem.md)
 - [MainApiSupplierCatalogUnit](docs/MainApiSupplierCatalogUnit.md)
 - [MainCatalogItems](docs/MainCatalogItems.md)
 - [MainFailedOrderError](docs/MainFailedOrderError.md)
 - [MainGenericErrorResponse](docs/MainGenericErrorResponse.md)
 - [MainLogMessage](docs/MainLogMessage.md)
 - [MainOrderListInput](docs/MainOrderListInput.md)
 - [MainOrderListOutput](docs/MainOrderListOutput.md)
 - [MainReplaceCatalogInput](docs/MainReplaceCatalogInput.md)
 - [MainSetErrorOrderInput](docs/MainSetErrorOrderInput.md)
 - [MainSetIntegrateOrdersInput](docs/MainSetIntegrateOrdersInput.md)
 - [MainSetStockStatusInput](docs/MainSetStockStatusInput.md)
 - [MainStockItemState](docs/MainStockItemState.md)
 - [MainSuccessConfirmation](docs/MainSuccessConfirmation.md)
 - [MainUpdateSuccess](docs/MainUpdateSuccess.md)
 - [V3Allergens](docs/V3Allergens.md)
 - [V3ApiSupplierCatalogItem](docs/V3ApiSupplierCatalogItem.md)
 - [V3ApiSupplierCatalogUnit](docs/V3ApiSupplierCatalogUnit.md)
 - [V3AvailabilityError](docs/V3AvailabilityError.md)
 - [V3AvailabilityErrorResponse](docs/V3AvailabilityErrorResponse.md)
 - [V3CatalogItems](docs/V3CatalogItems.md)
 - [V3ConfirmOrdersInput](docs/V3ConfirmOrdersInput.md)
 - [V3DeleteCatalogItemsInput](docs/V3DeleteCatalogItemsInput.md)
 - [V3DeliveryAddress](docs/V3DeliveryAddress.md)
 - [V3FailedOrderError](docs/V3FailedOrderError.md)
 - [V3GenericErrorResponse](docs/V3GenericErrorResponse.md)
 - [V3LogMessage](docs/V3LogMessage.md)
 - [V3Order](docs/V3Order.md)
 - [V3OrderCutoffTimes](docs/V3OrderCutoffTimes.md)
 - [V3OrderIntegrationError](docs/V3OrderIntegrationError.md)
 - [V3OrderListInput](docs/V3OrderListInput.md)
 - [V3OrderListOutput](docs/V3OrderListOutput.md)
 - [V3ReplaceCatalogInput](docs/V3ReplaceCatalogInput.md)
 - [V3Seasonality](docs/V3Seasonality.md)
 - [V3SetErrorOrderInput](docs/V3SetErrorOrderInput.md)
 - [V3SetIntegratedOrdersInput](docs/V3SetIntegratedOrdersInput.md)
 - [V3SetStockStatusInput](docs/V3SetStockStatusInput.md)
 - [V3StockItemState](docs/V3StockItemState.md)
 - [V3SuccessConfirmation](docs/V3SuccessConfirmation.md)
 - [V3UpdateCatalogInput](docs/V3UpdateCatalogInput.md)
 - [V3UpdateSuccess](docs/V3UpdateSuccess.md)


## Documentation For Authorization



## ApiKeyAuth

- **Type**: API key

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
    Key: "APIKEY",
    Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```



## Author

integrations@rekki.com

