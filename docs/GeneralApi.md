# \GeneralApi

All URIs are relative to *https://api.rekki.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLogMessage**](GeneralApi.md#PostLogMessage) | **Post** /integration/v1/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint
[**PostLogMessageV3**](GeneralApi.md#PostLogMessageV3) | **Post** /integration/v3/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint



## PostLogMessage

> string PostLogMessage(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint

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
    input := openapiclient.main.LogMessage{Level: "Level_example", Message: "Message_example"} // MainLogMessage | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneralApi.PostLogMessage(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralApi.PostLogMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLogMessage`: string
    fmt.Fprintf(os.Stdout, "Response from `GeneralApi.PostLogMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLogMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**MainLogMessage**](MainLogMessage.md) | Payload | 

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

> string PostLogMessageV3(ctx).XREKKIAuthorizationType(xREKKIAuthorizationType).Input(input).Execute()

Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint

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
    input := openapiclient.v3.LogMessage{Level: "Level_example", Message: "Message_example"} // V3LogMessage | Payload

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneralApi.PostLogMessageV3(context.Background(), xREKKIAuthorizationType, input).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralApi.PostLogMessageV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostLogMessageV3`: string
    fmt.Fprintf(os.Stdout, "Response from `GeneralApi.PostLogMessageV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLogMessageV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xREKKIAuthorizationType** | **string** | Required header | 
 **input** | [**V3LogMessage**](V3LogMessage.md) | Payload | 

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

