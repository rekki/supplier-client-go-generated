# \GeneralApi

All URIs are relative to *https://api.rekki.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLogMessage**](GeneralApi.md#PostLogMessage) | **Post** /integration/v1/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint
[**PostLogMessageV3**](GeneralApi.md#PostLogMessageV3) | **Post** /integration/v3/log | Post a log message for the supplier for internal debugging. There is no need to handle response from this endpoint



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

