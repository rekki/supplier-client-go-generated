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

> MainSetIntegrateOrdersInput ConfirmOrders(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Confirm a pending order by its reference code.



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
    input := openapiclient.main.SetIntegrateOrdersInput{Orders: []map[string]interface{}{123)} // MainSetIntegrateOrdersInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.ConfirmOrders(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ConfirmOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmOrders`: MainSetIntegrateOrdersInput
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ConfirmOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainSetIntegrateOrdersInput**](MainSetIntegrateOrdersInput.md) | Payload | 

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

> V3ConfirmOrdersInput ConfirmOrdersV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Confirm a pending order by its reference code.



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
    input := openapiclient.v3.ConfirmOrdersInput{Orders: []string{"Orders_example")} // V3ConfirmOrdersInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.ConfirmOrdersV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ConfirmOrdersV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmOrdersV3`: V3ConfirmOrdersInput
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ConfirmOrdersV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmOrdersV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3ConfirmOrdersInput**](V3ConfirmOrdersInput.md) | Payload | 

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

> MainOrderListOutput ListNotIntegratedOrders(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Lists all orders placed for the supplier that were placed through REKKI and not marked as integrated.

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
    input := openapiclient.main.OrderListInput{Since: 123} // MainOrderListInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.ListNotIntegratedOrders(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListNotIntegratedOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotIntegratedOrders`: MainOrderListOutput
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListNotIntegratedOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotIntegratedOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainOrderListInput**](MainOrderListInput.md) | Payload | 

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

> MainOrderListOutput ListOrdersBySupplier(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

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
    input := openapiclient.main.OrderListInput{Since: 123} // MainOrderListInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.ListOrdersBySupplier(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrdersBySupplier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrdersBySupplier`: MainOrderListOutput
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrdersBySupplier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersBySupplierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainOrderListInput**](MainOrderListInput.md) | Payload | 

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

> V3OrderListOutput ListOrdersBySupplierV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

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
    input := openapiclient.v3.OrderListInput{Since: "Since_example", SkipIntegrated: false} // V3OrderListInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.ListOrdersBySupplierV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrdersBySupplierV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrdersBySupplierV3`: V3OrderListOutput
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrdersBySupplierV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersBySupplierV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3OrderListInput**](V3OrderListInput.md) | Payload | 

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

> MainSuccessConfirmation MarkIntegrationError(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Report failure to integrate an order

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
    input := openapiclient.main.SetErrorOrderInput{Attempts: 123, Error: "Error_example", Order: openapiclient.integration.Order{ConfirmedAt: "ConfirmedAt_example", ContactInfo: "ContactInfo_example", ContactName: "ContactName_example", CustomerAccountNo: "CustomerAccountNo_example", DeliveryAddress: "DeliveryAddress_example", DeliveryOn: "DeliveryOn_example", InsertedAt: "InsertedAt_example", InsertedAtTs: 123, Items: []IntegrationOrderItem{openapiclient.integration.OrderItem{Id: "Id_example", Name: "Name_example", Price: "Price_example", PriceCents: 123, ProductCode: "ProductCode_example", Quantity: 123, Spec: "Spec_example", Units: "Units_example"}), LocationName: "LocationName_example", Notes: "Notes_example", Reference: "Reference_example", SupplierNotes: "SupplierNotes_example"}} // MainSetErrorOrderInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.MarkIntegrationError(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.MarkIntegrationError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkIntegrationError`: MainSuccessConfirmation
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.MarkIntegrationError`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkIntegrationErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainSetErrorOrderInput**](MainSetErrorOrderInput.md) | Payload | 

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

> V3SuccessConfirmation MarkIntegrationErrorV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Report failure to integrate an order

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
    input := openapiclient.v3.SetErrorOrderInput{Orders: []V3OrderIntegrationError{openapiclient.v3.OrderIntegrationError{Attempts: 123, Error: "Error_example", Reference: "Reference_example"})} // V3SetErrorOrderInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.MarkIntegrationErrorV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.MarkIntegrationErrorV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkIntegrationErrorV3`: V3SuccessConfirmation
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.MarkIntegrationErrorV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkIntegrationErrorV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3SetErrorOrderInput**](V3SetErrorOrderInput.md) | Payload | 

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

> MainUpdateSuccess MarkOrdersIntegrated(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Mark orders as integrated

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
    input := openapiclient.main.SetIntegrateOrdersInput{Orders: []map[string]interface{}{123)} // MainSetIntegrateOrdersInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.MarkOrdersIntegrated(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.MarkOrdersIntegrated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkOrdersIntegrated`: MainUpdateSuccess
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.MarkOrdersIntegrated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkOrdersIntegratedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainSetIntegrateOrdersInput**](MainSetIntegrateOrdersInput.md) | Payload | 

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

> V3UpdateSuccess MarkOrdersIntegratedV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Mark orders as integrated

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
    input := openapiclient.v3.SetIntegratedOrdersInput{Orders: []string{"Orders_example")} // V3SetIntegratedOrdersInput | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrdersApi.MarkOrdersIntegratedV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.MarkOrdersIntegratedV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkOrdersIntegratedV3`: V3UpdateSuccess
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.MarkOrdersIntegratedV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkOrdersIntegratedV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3SetIntegratedOrdersInput**](V3SetIntegratedOrdersInput.md) | Payload | 

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

