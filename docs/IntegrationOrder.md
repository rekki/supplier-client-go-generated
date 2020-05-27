# IntegrationOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmedAt** | **string** | the time at which the supplier confirmed the order (via email, or via the REKKI supplier app, or via the REKKI API) | [optional] 
**ContactInfo** | **string** | the phone number or email address for the person who placed the order | [optional] 
**ContactName** | **string** | the full name of the person who placed the order | [optional] 
**CustomerAccountNo** | **string** | the account number for customer within the supplier system, this can be setup in REKKI supplier app ( https://supplier.rekki.com ). | [optional] 
**DeliveryAddress** | **string** | delivery address for this specific order (address, postcode) | [optional] 
**DeliveryOn** | **string** | expected delivery date (when users place orders they specify for which day it is supposed to be delivered) | [optional] 
**InsertedAt** | **string** | when was the order created by the customer (ISO date time) | [optional] 
**InsertedAtTs** | **int32** | when was the order created by the customer (unix timestamp in seconds) | [optional] 
**Items** | [**[]IntegrationOrderItem**](integration.OrderItem.md) | items requests in this order | [optional] 
**LocationName** | **string** | the name of the location that placed the order, can be NULL | [optional] 
**Notes** | **string** | defined by the user at the moment of making an order and usually refer to that specific order (e.g. \&quot;please send fresher tomatoes\&quot;) | [optional] 
**Reference** | **string** | REKKI&#39;s order reference | [optional] 
**SupplierNotes** | **string** | notes define by the user for the supplier, usually being common across orders (e.g.: \&quot;please use the side entrance for delivery\&quot;) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


