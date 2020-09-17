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

> MainSuccessConfirmation DeleteCatalogItem(ctx, id).XREKKIAuthorizationType(xREKKIAuthorizationType).Execute()

Delete an item from catalog



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    id := "id_example" // string | ID of the item to retrieve. Item IDs are discoverable when listing items.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.DeleteCatalogItem(context.Background(), xREKKIAuthorizationType, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.DeleteCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCatalogItem`: MainSuccessConfirmation
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.DeleteCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the item to retrieve. Item IDs are discoverable when listing items. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 


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

> V3SuccessConfirmation DeleteCatalogItemsV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Delete items from catalog



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.v3.DeleteCatalogItemsInput{Items: []int32{123)} // V3DeleteCatalogItemsInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.DeleteCatalogItemsV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.DeleteCatalogItemsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCatalogItemsV3`: V3SuccessConfirmation
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.DeleteCatalogItemsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCatalogItemsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3DeleteCatalogItemsInput**](V3DeleteCatalogItemsInput.md) | Payload | 

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

> MainAPISupplierCatalogItem GetCatalogItem(ctx, id).XREKKIAuthorizationType(xREKKIAuthorizationType).Execute()

Lists all orders placed for the supplier that were placed through REKKI.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    id := "id_example" // string | ID of the item to retrieve. Item IDs are discoverable when listing items.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogItem(context.Background(), xREKKIAuthorizationType, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogItem`: MainAPISupplierCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the item to retrieve. Item IDs are discoverable when listing items. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 


### Return type

[**MainAPISupplierCatalogItem**](main.APISupplierCatalogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItemV3

> V3APISupplierCatalogItem GetCatalogItemV3(ctx, id).XREKKIAuthorizationType(xREKKIAuthorizationType).Execute()

Fetch a specific catalog item by its Id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    id := "id_example" // string | ID of the item to retrieve. Item IDs are discoverable when listing items.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogItemV3(context.Background(), xREKKIAuthorizationType, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogItemV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogItemV3`: V3APISupplierCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogItemV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the item to retrieve. Item IDs are discoverable when listing items. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 


### Return type

[**V3APISupplierCatalogItem**](v3.APISupplierCatalogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogItems

> MainCatalogItems GetCatalogItems(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Execute()

Get catalog items for authenticated supplier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogItems(context.Background(), xREKKIAuthorizationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogItems`: MainCatalogItems
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 

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

> V3CatalogItems GetCatalogItemsV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Execute()

Get catalog items for authenticated supplier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.GetCatalogItemsV3(context.Background(), xREKKIAuthorizationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogItemsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogItemsV3`: V3CatalogItems
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogItemsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogItemsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 

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

> MainSuccessConfirmation ReplaceCatalog(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Drop all existing items from the catalog and upload new ones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.main.ReplaceCatalogInput{Data: []MainAPISupplierCatalogItem{openapiclient.main.APISupplierCatalogItem{Allergens: []CatalogAllergen{openapiclient.catalog.Allergen{Symptoms: []string{"Symptoms_example"), Type: "Type_example"}), Availability: "Availability_example", Currency: "Currency_example", Description: "Description_example", Id: 123, Name: "Name_example", OrderCutoffTimes: openapiclient.catalog.OrderCutoffTimes{Fri: 123, Mon: 123, Thu: 123, Tue: 123, Wed: 123}, ProductCode: "ProductCode_example", ReplacementProducts: []string{"ReplacementProducts_example"), Seasonality: []CatalogSeasonality{openapiclient.catalog.Seasonality{EndDate: "EndDate_example", StartDate: "StartDate_example"}), UnitsPrices: []MainAPISupplierCatalogUnit{openapiclient.main.APISupplierCatalogUnit{PriceCents: 123, StockCount: 123, Unit: "Unit_example"})})} // MainReplaceCatalogInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ReplaceCatalog(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ReplaceCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCatalog`: MainSuccessConfirmation
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ReplaceCatalog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainReplaceCatalogInput**](MainReplaceCatalogInput.md) | Payload | 

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

> V3SuccessConfirmation ReplaceCatalogV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Drop all existing items from the catalog and upload new ones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.v3.ReplaceCatalogInput{Data: []V3APISupplierCatalogItem{openapiclient.v3.APISupplierCatalogItem{Allergens: []V3Allergens{openapiclient.v3.Allergens{Symptoms: []string{"Symptoms_example"), Type: "Type_example"}), Availability: "Availability_example", Currency: "Currency_example", Description: "Description_example", Id: 123, Name: "Name_example", OrderCutoffTimes: openapiclient.v3.OrderCutoffTimes{Fri: 123, Mon: 123, Thu: 123, Tue: 123, Wed: 123}, ProductCode: "ProductCode_example", ReplacementProducts: []string{"ReplacementProducts_example"), Seasonality: []V3Seasonality{openapiclient.v3.Seasonality{EndDate: "EndDate_example", StartDate: "StartDate_example"}), UnitsPrices: []V3APISupplierCatalogUnit{openapiclient.v3.APISupplierCatalogUnit{PriceCents: 123, StockCount: 123, Unit: "Unit_example"})})} // V3ReplaceCatalogInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.ReplaceCatalogV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.ReplaceCatalogV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceCatalogV3`: V3SuccessConfirmation
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.ReplaceCatalogV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCatalogV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3ReplaceCatalogInput**](V3ReplaceCatalogInput.md) | Payload | 

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

> MainAPISupplierCatalogItem UpdateCatalogItem(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Creates or Updates an item on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.main.APISupplierCatalogItem{Allergens: []CatalogAllergen{openapiclient.catalog.Allergen{Symptoms: []string{"Symptoms_example"), Type: "Type_example"}), Availability: "Availability_example", Currency: "Currency_example", Description: "Description_example", Id: 123, Name: "Name_example", OrderCutoffTimes: openapiclient.catalog.OrderCutoffTimes{Fri: 123, Mon: 123, Thu: 123, Tue: 123, Wed: 123}, ProductCode: "ProductCode_example", ReplacementProducts: []string{"ReplacementProducts_example"), Seasonality: []CatalogSeasonality{openapiclient.catalog.Seasonality{EndDate: "EndDate_example", StartDate: "StartDate_example"}), UnitsPrices: []MainAPISupplierCatalogUnit{openapiclient.main.APISupplierCatalogUnit{PriceCents: 123, StockCount: 123, Unit: "Unit_example"})} // MainAPISupplierCatalogItem | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.UpdateCatalogItem(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.UpdateCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogItem`: MainAPISupplierCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.UpdateCatalogItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainAPISupplierCatalogItem**](MainAPISupplierCatalogItem.md) | Payload | 

### Return type

[**MainAPISupplierCatalogItem**](main.APISupplierCatalogItem.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCatalogItemAvailability

> MainUpdateSuccess UpdateCatalogItemAvailability(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Update availability status for one of the items in the catalog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.main.SetStockStatusInput{Items: []MainStockItemState{openapiclient.main.StockItemState{Availability: "Availability_example", ProductCode: "ProductCode_example"})} // MainSetStockStatusInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.UpdateCatalogItemAvailability(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.UpdateCatalogItemAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogItemAvailability`: MainUpdateSuccess
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.UpdateCatalogItemAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainSetStockStatusInput**](MainSetStockStatusInput.md) | Payload | 

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

> V3UpdateSuccess UpdateCatalogItemAvailabilityV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Update availability status for one of the items in the catalog

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.v3.SetStockStatusInput{Items: []V3StockItemState{openapiclient.v3.StockItemState{Availability: "Availability_example", ProductCode: "ProductCode_example"})} // V3SetStockStatusInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.UpdateCatalogItemAvailabilityV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.UpdateCatalogItemAvailabilityV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogItemAvailabilityV3`: V3UpdateSuccess
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.UpdateCatalogItemAvailabilityV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemAvailabilityV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3SetStockStatusInput**](V3SetStockStatusInput.md) | Payload | 

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

> V3CatalogItems UpdateCatalogItemsV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Creates or Updates multiple items on your catalog. If item with this product_code already exists, you can update this item. Item is looked up by product code



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xREKKIAuthorizationType := "xREKKIAuthorizationType_example" // string | Required header
    input := openapiclient.v3.UpdateCatalogInput{Data: []V3APISupplierCatalogItem{openapiclient.v3.APISupplierCatalogItem{Allergens: []V3Allergens{openapiclient.v3.Allergens{Symptoms: []string{"Symptoms_example"), Type: "Type_example"}), Availability: "Availability_example", Currency: "Currency_example", Description: "Description_example", Id: 123, Name: "Name_example", OrderCutoffTimes: openapiclient.v3.OrderCutoffTimes{Fri: 123, Mon: 123, Thu: 123, Tue: 123, Wed: 123}, ProductCode: "ProductCode_example", ReplacementProducts: []string{"ReplacementProducts_example"), Seasonality: []V3Seasonality{openapiclient.v3.Seasonality{EndDate: "EndDate_example", StartDate: "StartDate_example"}), UnitsPrices: []V3APISupplierCatalogUnit{openapiclient.v3.APISupplierCatalogUnit{PriceCents: 123, StockCount: 123, Unit: "Unit_example"})})} // V3UpdateCatalogInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.UpdateCatalogItemsV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.UpdateCatalogItemsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCatalogItemsV3`: V3CatalogItems
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.UpdateCatalogItemsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCatalogItemsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3UpdateCatalogInput**](V3UpdateCatalogInput.md) | Payload | 

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

